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

package repl

import (
	"fmt"
	"go/token"
	"os/exec"
	"reflect"
	"regexp"
	"strings"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/internal/constant"
	"golang.org/x/tools/go/ssa"
)

const (
	// ContinuePrompt - the current code statement is not completed.
	ContinuePrompt string = "... "
	// NormalPrompt - start of a code statement.
	NormalPrompt string = ">>> "
)

type UI interface {
	SetPrompt(prompt string)
	Printf(format string, a ...interface{})
}

type REPL struct {
	*ixgo.Repl
	term UI
	more string
}

func NewREPL(mode ixgo.Mode) *REPL {
	r := &REPL{}
	ixgo.RegisterCustomBuiltin("__ixgo_repl_info__", func(v interface{}) {
		r.Printf("%v %T\n", v, v)
	})
	ctx := ixgo.NewContext(mode)
	r.Repl = ixgo.NewRepl(ctx)
	return r
}

func (r *REPL) SetUI(term UI) {
	r.term = term
	term.SetPrompt(NormalPrompt)
}

func (r *REPL) SetNormal() {
	r.more = ""
	r.SetPrompt(NormalPrompt)
}

func (r *REPL) SetPrompt(prompt string) {
	if r.term != nil {
		r.term.SetPrompt(prompt)
	}
}

func (r *REPL) IsNormal() bool {
	return r.more == ""
}

func (r *REPL) TryDump(expr string) bool {
	i := r.Interp()
	if i != nil {
		if m, v, ok := i.GetSymbol(expr); ok {
			switch p := m.(type) {
			case *ssa.NamedConst:
				r.Printf("const %v %v\n", constant.ExactConstant(p.Value.Value), p.Type())
			case *ssa.Global:
				e := reflect.ValueOf(v).Elem().Interface()
				r.Printf("%v %T (global var)\n", e, e)
			case *ssa.Type:
				if m.Package().Pkg.Name() == "main" {
					return false
				}
				if r.tryDumpByPkg(expr) != nil {
					r.Printf("%v %v\n", p.Type().Underlying(), v)
				}
			case *ssa.Function:
				n := p.Signature.Params().Len()
				if n == 0 || (n == 1 && p.Signature.Variadic()) {
					return false
				}
				r.Printf("%v %v\n", v, p.Type())
			}
			return true
		}
	}
	return false
}

func (r *REPL) Dump(expr string) {
	i := r.Interp()
	if i != nil {
		if m, v, ok := i.GetSymbol(expr); ok {
			switch p := m.(type) {
			case *ssa.NamedConst:
				r.Printf("const %v %v\n", constant.ExactConstant(p.Value.Value), p.Type())
			case *ssa.Global:
				e := reflect.ValueOf(v).Elem().Interface()
				r.Printf("%v %T (global var)\n", e, e)
			case *ssa.Type:
				if r.tryDumpByPkg(expr) != nil {
					r.Printf("%v %v\n", p.Type().Underlying(), v)
				}
			case *ssa.Function:
				r.Printf("%v %v\n", v, p.Type())
			}
			return
		}
	}
	_, _, err := r.Eval(fmt.Sprintf("__ixgo_repl_info__(%v)", expr))
	if err == nil {
		return
	}
	if err = r.godoc(expr); err == nil {
		return
	}
	if err = r.tryDumpByPkg(expr); err == nil {
		return
	}
	r.Printf("not found %v\n", expr)
}

func (r *REPL) tryDumpByPkg(expr string) error {
	pkg, sym, _, err := parseSymbol(expr)
	if err != nil {
		return err
	}
	if pkgPath, found := findPkg(pkg); found {
		if p, found := ixgo.LookupPackage(pkgPath); found {
			if sym == "" {
				r.Printf("%v\n", dumpPkg(p))
				return nil
			}
			if info, ok := lookupSymbol(p, sym); ok {
				r.Printf("%v\n", info)
				return nil
			}
			return fmt.Errorf("not found symbol %v.%v", pkg, sym)
		}
	}
	return fmt.Errorf("not found pkg %v", pkg)
}

func (r *REPL) godoc(expr string) error {
	gobin, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	cmd := exec.Command(gobin, "doc", expr)
	data, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	r.Printf("%v\n", string(data))
	return nil
}

func (r *REPL) Printf(_fmt string, a ...interface{}) {
	if r.term != nil {
		r.term.Printf(_fmt, a...)
	} else {
		fmt.Printf(_fmt, a...)
	}
}

var (
	regWord = regexp.MustCompile("\\w+")
)

func (r *REPL) Run(line string) error {
	var expr string
	if r.more != "" {
		if line == "" {
			r.SetNormal()
			return nil
		}
		expr = r.more + "\n" + line
	} else {
		if strings.HasPrefix(line, "?") {
			r.Dump(strings.TrimSpace(line[1:]))
			return nil
		} else if regWord.MatchString(line) {
			if r.TryDump(line) {
				return nil
			}
		}
		expr = line
	}
	tok, eval, err := r.Eval(expr)
	if err != nil {
		if checkMore(tok, err) {
			r.more += "\n" + line
			r.SetPrompt(ContinuePrompt)
			return nil
		} else {
			r.SetNormal()
		}
		return err
	}
	switch len(eval) {
	case 0:
	case 1:
		r.Printf("%v\n", eval[0])
	default:
		var info []string
		for _, v := range eval {
			info = append(info, v.String())
		}
		r.Printf("(%v)\n", strings.Join(info, ", "))
	}
	r.SetNormal()
	return nil
}

func checkMore(tok token.Token, err error) bool {
	s := err.Error()
	if strings.Contains(s, `expected `) {
		return true
	}
	return false
}
