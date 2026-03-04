/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ixgo

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/scanner"
	"go/token"
	"go/types"
	"reflect"
	"runtime"
	"strings"

	xconst "github.com/goplus/ixgo/internal/constant"
	"golang.org/x/tools/go/ssa"
)

/*
package main

import "fmt"

const/var/type/func

func main(){}
*/

type Repl struct {
	ctx       *Context               // interp context
	pkg       *ssa.Package           // main package
	builtin   *ast.File              // builtin func
	interp    *Interp                // last interp
	fsInit    *fnState               // func init
	fsMain    *fnState               // func main
	globalMap map[string]interface{} // keep repl global map
	fileName  string                 // file.go or file.gop
	source    string                 // all source
	imports   []string               // import lines
	globals   []string               // global var/func/type
	infuncs   []string               // in main func
	lastEval  []*Eval                // last __ixgo_repl_eval__
}

type Eval struct {
	Value interface{} // constant.Value or interface{}
	Type  reflect.Type
}

func (e *Eval) String() string {
	if c, ok := e.Value.(constant.Value); ok {
		s, _ := xconst.ExactConstantEx(c, true)
		if c.Kind() == constant.Float {
			es := c.ExactString()
			return fmt.Sprintf("%v untyped %v\n(%v)", s, e.Type, es)
		} else {
			return fmt.Sprintf("%v untyped %v", s, e.Type)
		}
	}
	return fmt.Sprintf("%v %v", e.Value, e.Type)
}

func toEval(i []interface{}) (eval []*Eval) {
	for _, v := range i {
		eval = append(eval, &Eval{v, reflect.TypeOf(v)})
	}
	return
}

func toResultEval(interp *Interp, vs []interface{}, rs *types.Tuple) (eval []*Eval) {
	for i, v := range vs {
		eval = append(eval, &Eval{v, interp.toType(rs.At(i).Type())})
	}
	return
}

func NewRepl(ctx *Context) *Repl {
	r := &Repl{
		ctx:       ctx,
		globalMap: make(map[string]interface{}),
	}
	ctx.SetEvalMode(true)
	RegisterCustomBuiltin("__ixgo_repl_used__", func(v interface{}) {})
	RegisterCustomBuiltin("__ixgo_repl_eval__", func(v ...interface{}) {
		r.lastEval = toEval(v)
	})
	// reset runtime.GC to default
	ctx.RegisterExternal("runtime.GC", runtime.GC)
	ctx.evalCallFn = func(interp *Interp, call *ssa.Call, res ...interface{}) {
		if len(*call.Referrers()) != 0 {
			return
		}
		v := call.Call.Value
		if un, ok := v.(*ssa.UnOp); ok {
			v = un.X
		}
		if strings.Contains(v.Name(), "__ixgo_repl_") {
			return
		}
		r.lastEval = toResultEval(interp, res, call.Call.Signature().Results())
	}
	if f, err := ParseBuiltin(r.ctx.FileSet, "main"); err == nil {
		r.builtin = f
	}
	return r
}

func (r *Repl) SetFileName(filename string) {
	r.fileName = filename
}

func (r *Repl) Eval(expr string) (tok token.Token, dump []*Eval, err error) {
	r.lastEval = nil
	tok = r.firstToken(expr)
	err = r.eval(tok, expr)
	return tok, r.lastEval, err
}

func (r *Repl) Interp() *Interp {
	return r.interp
}

func (r *Repl) Source() string {
	return r.source
}

func (r *Repl) buildSource(expr string, tok token.Token) string {
	v0 := strings.Join(r.imports, "\n")
	v1 := strings.Join(r.globals, "\n")
	v2 := strings.Join(r.infuncs, "\n")
	switch tok {
	case token.IMPORT:
		v0 += "\n" + expr
	case token.CONST, token.VAR, token.TYPE, token.FUNC:
		v1 += "\n" + expr
	default:
		v2 += "\n" + expr
	}
	return fmt.Sprintf(`package main
%v
%v
func main() {
	%v
}
`, v0, v1, v2)
}

func (r *Repl) eval(tok token.Token, expr string) (err error) {
	var src string
	var inMain bool
	var evalConst bool
	var lastErrors []error
	switch tok {
	case token.PACKAGE:
		// skip package
		return nil
	case token.IMPORT:
		src = r.buildSource(expr, tok)
		errors, err := r.check(r.fileName, src)
		if err != nil {
			return err
		}
		if errors != nil {
			return errors[0]
		}
		r.imports = append(r.imports, expr)
		return nil
	case token.FUNC:
		src = r.buildSource(expr, tok)
		errors, err := r.check(r.fileName, src)
		if err != nil {
			// check funlit
			msg := err.Error()
			if strings.Contains(msg, errMaybeGoFunLit) || strings.Contains(msg, errMaybeGopFunLit) {
				inMain = true
				src = r.buildSource(expr, token.ILLEGAL)
				errors, err = r.check(r.fileName, src)
			}
			if err != nil {
				return err
			}
		}
		if errors != nil {
			return errors[0]
		}
	case token.CONST, token.TYPE, token.VAR:
		src = r.buildSource(expr, tok)
		errors, err := r.check(r.fileName, src)
		if err != nil {
			return err
		}
		if errors != nil {
			return errors[0]
		}
	default:
		orgExpr := expr
		switch tok {
		case token.FOR, token.IF, token.SWITCH, token.SELECT:
			expr = "func(){\n" + expr + "\n}()"
		case token.INT, token.FLOAT, token.IMAG, token.CHAR, token.STRING,
			token.ADD, token.SUB, token.NOT, token.XOR,
			token.LPAREN, token.LBRACK, token.LBRACE:
			expr = "__ixgo_repl_eval__(" + expr + ")"
		}
		inMain = true
		src = r.buildSource(expr, tok)
		errors, err := r.check(r.fileName, src)
		if err != nil {
			return err
		}
		var fixed []string
		for _, err := range errors {
			e, ok := err.(types.Error)
			if !ok {
				return e
			}
			if v, ok := isTypesDeclaredNotUsed(e.Msg); ok {
				fixed = append(fixed, "__ixgo_repl_used__(&"+v+")")
				// fixed = append(fixed, "__ixgo_repl_eval__("+v+")")
			} else if strings.HasSuffix(e.Msg, errIsNotUsed) {
				if _, ok := extractConstant([]byte(e.Msg[:len(e.Msg)-len(errIsNotUsed)])); ok {
					expr = "const __ixgo_repl_const__ = " + orgExpr
					tok = token.CONST
					evalConst = true
				} else {
					expr = "__ixgo_repl_eval__(" + orgExpr + ")"
				}
			} else if strings.HasSuffix(e.Msg, errDumpOverflows) {
				if _, ok := extractConstant([]byte(e.Msg[:len(e.Msg)-len(errDumpOverflows)])); ok {
					expr = "const __ixgo_repl_const__ = " + orgExpr
					tok = token.CONST
					evalConst = true
				} else {
					return e
				}
			} else {
				return e
			}
		}
		lastErrors = errors
		if len(fixed) != 0 {
			expr += "\n" + strings.Join(fixed, "\n")
		}
		src = r.buildSource(expr, tok)
	}
	r.pkg, err = r.ctx.LoadFile(r.fileName, src)
	if err != nil {
		if lastErrors != nil {
			return lastErrors[0]
		}
		return err
	}
	i, err := newInterp(r.ctx, r.pkg, r.globalMap)
	if err != nil {
		return err
	}
	rinit, err := r.runFunc(i, "init", r.fsInit)
	if err == nil {
		rinit.pc--
	}
	rmain, err := r.runFunc(i, "main", r.fsMain)
	if err != nil {
		return err
	}
	if evalConst {
		m := i.mainpkg.Members["__ixgo_repl_const__"]
		c, _ := m.(*ssa.NamedConst)
		r.lastEval = []*Eval{&Eval{c.Value.Value, i.toType(c.Type())}}
		return nil
	}
	r.interp = i
	r.fsInit = rinit
	r.fsMain = rmain
	for k, v := range i.globals {
		r.globalMap[k] = v
	}
	if inMain {
		r.infuncs = append(r.infuncs, expr)
	} else {
		r.globals = append(r.globals, expr)
	}
	r.source = src
	return nil
}

const (
	errIsNotUsed      = "is not used"
	errDumpOverflows  = "to __ixgo_repl_eval__ (overflows)"
	errMaybeGoFunLit  = `expected 'IDENT', found '{'`
	errMaybeGopFunLit = `expected '(',`
)

func (r *Repl) check(filename string, src interface{}) (errors []error, err error) {
	file, err := r.ctx.ParseFile(filename, src)
	if err != nil {
		return nil, err
	}
	tc := &types.Config{
		Importer:                 NewImporter(r.ctx),
		DisableUnusedImportCheck: true,
	}
	tc.Error = func(err error) {
		errors = append(errors, err)
	}
	pkg := types.NewPackage(file.Name.Name, "")
	var files []*ast.File
	if r.builtin != nil {
		files = []*ast.File{r.builtin, file}
	} else {
		files = []*ast.File{file}
	}
	types.NewChecker(tc, r.ctx.FileSet, pkg, &types.Info{}).Files(files)
	return
}

func (r *Repl) firstToken(src string) token.Token {
	var s scanner.Scanner
	file := r.ctx.FileSet.AddFile("", -1, len(src))
	s.Init(file, []byte(src), nil, 0)
	_, tok, _ := s.Scan()
	return tok
}

type fnState struct {
	fr *frame
	pc int
}

func (r *Repl) runFunc(i *Interp, fnname string, fs *fnState) (rfs *fnState, err error) {
	defer func() {
		switch p := recover().(type) {
		case nil:
		case exitPanic:
			err = ExitError(int(p))
		default:
			err = fmt.Errorf("%v", p)
		}
	}()
	fn := r.pkg.Func(fnname)
	if fn == nil {
		return nil, fmt.Errorf("no function %v", fnname)
	}
	pfn := i.funcs[fn]
	fr := pfn.allocFrame(&frame{})
	if fs != nil {
		copy(fr.stack, fs.fr.stack)
		fr.ipc = fs.pc
	}
	var pc int
	for fr.ipc != -1 {
		fn := fr.pfn.Instrs[fr.ipc]
		pc = fr.ipc
		fr.ipc++
		fn(fr)
	}
	return &fnState{fr: fr, pc: pc}, nil
}

type tokenLit struct {
	Lit string
	Pos token.Pos
	Tok token.Token
}

// extractConstant extract constant int and overflows float/complex
func extractConstant(src []byte) (constant *tokenLit, ok bool) {
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, 0)
	pos, tok, lit := s.Scan()
	if tok == token.EOF {
		return
	}
	first := &tokenLit{lit, pos, tok}
	var check int
	for {
		_, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		switch tok {
		case token.LPAREN:
			check++
		case token.RPAREN:
			check--
		case token.IDENT:
			if check == 1 && lit == "constant" {
				var nodes []*tokenLit
			loop:
				for {
					pos, tok, lit := s.Scan()
					if tok == token.EOF {
						break
					}
					switch tok {
					case token.LPAREN:
						check++
					case token.RPAREN:
						check--
						if check == 0 {
							break loop
						}
					default:
						nodes = append(nodes, &tokenLit{lit, pos, tok})
					}
				}
				switch len(nodes) {
				case 0:
					return first, true
				case 1:
					switch nodes[0].Tok {
					case token.INT:
						return nodes[0], true
					case token.FLOAT:
						return nodes[0], true
						// extract if not parse float64
						// if _, err := strconv.ParseFloat(nodes[0].Lit, 128); err != nil {
						// 	return nodes[0], true
						// }
					}
				default:
					last := nodes[len(nodes)-1]
					return &tokenLit{string(src[int(nodes[0].Pos)-1 : int(last.Pos)+len(last.Lit)]), nodes[0].Pos, token.IMAG}, true
				}
			}
		}
	}
	return
}
