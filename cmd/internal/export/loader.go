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

package export

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/constant"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"strconv"
	"strings"

	"golang.org/x/tools/go/loader"
)

type Program struct {
	prog *loader.Program
	ctx  *build.Context
	fset *token.FileSet
	imps map[*loader.PackageInfo][]*ast.ImportSpec
}

func NewProgram(ctx *build.Context) *Program {
	if ctx == nil {
		ctx = &build.Default
		ctx.BuildTags = strings.Split(flagBuildTags, ",")
	}
	return &Program{ctx: ctx, fset: token.NewFileSet(), imps: make(map[*loader.PackageInfo][]*ast.ImportSpec)}
}

func (p *Program) Load(pkgs []string) error {
	var cfg loader.Config
	cfg.Build = p.ctx
	cfg.Fset = p.fset
	if flagExportSource {
		cfg.AfterTypeCheck = p.typeCheck
	} else if flagExportCode {
		cfg.AfterTypeCheck = p.typeCheckCode
	}
	for _, pkg := range pkgs {
		cfg.Import(pkg)
	}
	iprog, err := cfg.Load()
	if err != nil {
		return fmt.Errorf("conf.Load failed: %s", err)
	}
	p.prog = iprog
	return nil
}

func (p *Program) typeCheckCode(info *loader.PackageInfo, files []*ast.File) {
	for _, file := range files {
		p.imps[info] = append(p.imps[info], file.Imports...)
	}
}

func (p *Program) typeCheck(info *loader.PackageInfo, files []*ast.File) {
	for _, file := range files {
		p.imps[info] = append(p.imps[info], file.Imports...)
		for _, decl := range file.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				if funcHasTypeParams(fn) {
					continue
				}
				if recv := recvType(fn); recv != nil && !ast.IsExported(recv.Name) {
					continue
				}
				if !ast.IsExported(fn.Name.Name) {
					continue
				}
				fn.Body = nil
			}
		}
	}
}

func recvType(fn *ast.FuncDecl) *ast.Ident {
	if fn.Recv == nil {
		return nil
	}
	if len(fn.Recv.List) != 1 {
		return nil
	}
	expr := fn.Recv.List[0].Type
retry:
	switch v := expr.(type) {
	case *ast.ParenExpr:
		expr = v.X
		goto retry
	case *ast.StarExpr:
		expr = v.X
		goto retry
	case *ast.Ident:
		return v
	}
	return nil
}

func loadProgram(path string, ctx *build.Context) (*Program, error) {
	var cfg loader.Config
	cfg.Build = ctx
	cfg.Import(path)

	iprog, err := cfg.Load()
	if err != nil {
		return nil, fmt.Errorf("conf.Load failed: %s", err)
	}
	return &Program{prog: iprog, ctx: ctx}, nil
}

func (p *Program) DumpDeps(path string) {
	pkg := p.prog.Package(path)
	for _, im := range pkg.Pkg.Imports() {
		fmt.Println(im.Path())
	}
}

func (p *Program) dumpDeps(path string, sep string) {
	pkg := p.prog.Package(path)
	for _, im := range pkg.Pkg.Imports() {
		fmt.Println(sep, im.Path())
		p.dumpDeps(im.Path(), sep+"  ")
	}
}

func (p *Program) DumpExport(path string) {
	pkg := p.prog.Package(path)
	for _, v := range pkg.Pkg.Scope().Names() {
		if token.IsExported(v) {
			fmt.Println(v)
		}
	}
}

/*
type ConstValue struct {
	Typ   string
	Value constant.Value
}

type Package struct {
	Name    string
	Path    string
	Types   []reflect.Type
	Vars    map[string]reflect.Value
	Funcs   map[string]reflect.Value
	Consts  map[string]ConstValue
	Deps    map[string]string
}

*/

type Package struct {
	Name          string
	Path          string
	Deps          []string
	NamedTypes    []string
	Interfaces    []string
	AliasTypes    []string
	Vars          []string
	Funcs         []string
	Consts        []string
	TypedConsts   []string
	UntypedConsts []string
	Links         []string
	Source        string
	usedPkg       bool
}

func (p *Package) IsEmpty() bool {
	return len(p.NamedTypes) == 0 && len(p.Interfaces) == 0 &&
		len(p.AliasTypes) == 0 && len(p.Vars) == 0 &&
		len(p.Funcs) == 0 && len(p.Consts) == 0 &&
		len(p.TypedConsts) == 0 && len(p.UntypedConsts) == 0
}

/*
func unmarshalFloat(str string) constant.Value {
	if sep := strings.IndexByte(str, '/'); sep >= 0 {
		x := constant.MakeFromLiteral(str[:sep], token.FLOAT, 0)
		y := constant.MakeFromLiteral(str[sep+1:], token.FLOAT, 0)
		return constant.BinaryOp(x, token.QUO, y)
	}
	return constant.MakeFromLiteral(str, token.FLOAT, 0)
}
*/

func (p *Program) constToLit(named string, c constant.Value) string {
	switch c.Kind() {
	case constant.Bool:
		if named != "" {
			return fmt.Sprintf("constant.MakeBool(bool(%v))", named)
		}
		return fmt.Sprintf("constant.MakeBool(%v)", constant.BoolVal(c))
	case constant.String:
		if named != "" {
			return fmt.Sprintf("constant.MakeString(string(%v))", named)
		}
		return fmt.Sprintf("constant.MakeString(%q)", constant.StringVal(c))
	case constant.Int:
		if v, ok := constant.Int64Val(c); ok {
			if named != "" {
				return fmt.Sprintf("constant.MakeInt64(int64(%v))", named)
			}
			return fmt.Sprintf("constant.MakeInt64(%v)", v)
		} else if v, ok := constant.Uint64Val(c); ok {
			if named != "" {
				return fmt.Sprintf("constant.MakeUint64(uint64(%v))", named)
			}
			return fmt.Sprintf("constant.MakeUint64(%v)", v)
		}
		return fmt.Sprintf("constant.MakeFromLiteral(%q, token.INT, 0)", c.ExactString())
	case constant.Float:
		s := c.ExactString()
		if pos := strings.IndexByte(s, '/'); pos >= 0 {
			sx := s[:pos]
			sy := s[pos+1:]
			// simplify 314/100 => 3.14
			// 80901699437494742410229341718281905886015458990288143106772431
			// 50000000000000000000000000000000000000000000000000000000000000
			if strings.HasPrefix(sy, "1") && strings.Count(sy, "0") == len(sy)-1 {
				if len(sx) == len(sy) {
					return fmt.Sprintf("constant.MakeFromLiteral(\"%v.%v\", token.FLOAT, 0)", sx[:1], sx[1:])
				} else if len(sx) == len(sy)-1 {
					return fmt.Sprintf("constant.MakeFromLiteral(\"0.%v\", token.FLOAT, 0)", sx)
				} else if len(sx) < len(sy) {
					return fmt.Sprintf("constant.MakeFromLiteral(\"%v.%ve-%v\", token.FLOAT, 0)", sx[:1], sx[1:], len(sy)-len(sx))
				}
			} else if strings.HasPrefix(sy, "5") && strings.Count(sy, "0") == len(sy)-1 {
				if len(sx) == len(sy) {
					c := constant.BinaryOp(constant.MakeFromLiteral(sx, token.INT, 0), token.MUL, constant.MakeInt64(2))
					sx = c.ExactString()
					return fmt.Sprintf("constant.MakeFromLiteral(\"%v.%v\", token.FLOAT, 0)", sx[:1], sx[1:])
				}
			} else if strings.HasPrefix(sx, "1") && strings.Count(sx, "0") == len(sx)-1 {
				// skip
			}
			x := fmt.Sprintf("constant.MakeFromLiteral(%q, token.INT, 0)", sx)
			y := fmt.Sprintf("constant.MakeFromLiteral(%q, token.INT, 0)", sy)
			return fmt.Sprintf("constant.BinaryOp(%v, token.QUO, %v)", x, y)
		}
		if pos := strings.LastIndexAny(s, "123456789"); pos != -1 {
			sx := s[:pos+1]
			return fmt.Sprintf("constant.MakeFromLiteral(\"%v.%ve+%v\", token.FLOAT, 0)", sx[:1], sx[1:], len(s)-1)
		}
		return fmt.Sprintf("constant.MakeFromLiteral(%q, token.FLOAT, 0)", s)
	case constant.Complex:
		re := p.constToLit("", constant.Real(c))
		im := p.constToLit("", constant.Imag(c))
		return fmt.Sprintf("constant.BinaryOp(%v, token.ADD, constan.MakeImag(%v))", re, im)
	default:
		panic("unreachable")
	}
}

func (p *Program) ExportSource(e *Package, info *loader.PackageInfo) error {
	pkg := info.Pkg
	pkgPath := pkg.Path()
	pkgName := pkg.Name()

	outf := new(ast.File)
	outf.Name = ast.NewIdent(pkgName)

	var specs []ast.Spec
	impls := p.imps[info]
	for _, im := range pkg.Imports() {
		spec := &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(im.Path()),
			},
		}
		for _, i := range impls {
			if i.Name == nil {
				continue
			}
			// update name
			if path, err := strconv.Unquote(i.Path.Value); err == nil && path == im.Path() {
				if spec.Name == nil || spec.Name.Name == "_" {
					spec.Name = i.Name
				}
			}
		}
		specs = append(specs, spec)
	}
	if len(specs) > 0 {
		outf.Decls = append(outf.Decls, &ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: specs,
		})
	}

	var links []string
	for _, file := range info.Files {
		outf.Imports = append(outf.Imports, file.Imports...)
		for _, decl := range file.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				if d.Tok == token.IMPORT {
					continue
				}
				if d.Tok == token.VAR {
					var skip bool
					for _, spec := range d.Specs {
						for _, name := range spec.(*ast.ValueSpec).Names {
							if name.Name == "_" {
								skip = true
								continue
							}
						}
					}
					if skip {
						continue
					}
				}
				outf.Decls = append(outf.Decls, d)
			case *ast.FuncDecl:
				outf.Decls = append(outf.Decls, d)
				if funcHasTypeParams(d) {
					continue
				}
				fnName := d.Name.Name
				if d.Recv == nil && d.Body == nil && !ast.IsExported(fnName) {
					decl := &ast.FuncDecl{}
					decl.Type = d.Type
					lcName := "_" + fnName
					decl.Name = ast.NewIdent(lcName)
					decl.Doc = &ast.CommentGroup{[]*ast.Comment{
						&ast.Comment{Text: fmt.Sprintf("//go:linkname %v %v.%v", lcName, pkgPath, d.Name)},
					}}
					var buf bytes.Buffer
					printer.Fprint(&buf, p.fset, decl)
					links = append(links, buf.String())
					e.Funcs = append(e.Funcs, fmt.Sprintf("%q : reflect.ValueOf(%v)", fnName, lcName))
				}
			}
		}
	}
	var buf bytes.Buffer
	err := printer.Fprint(&buf, p.fset, outf)
	if err != nil {
		return err
	}
	e.Links = links
	e.Source = strconv.Quote(buf.String())
	return nil
}

func (p *Program) ExportPkg(path string, sname string) (*Package, error) {
	info := p.prog.Package(path)
	if info == nil {
		return nil, fmt.Errorf("not found path %v", path)
	}
	pkg := info.Pkg
	pkgPath := pkg.Path()
	pkgName := pkg.Name()
	e := &Package{Name: pkgName, Path: pkgPath}
	pkgName = sname
	for _, v := range pkg.Imports() {
		e.Deps = append(e.Deps, fmt.Sprintf("%q: %q", v.Path(), v.Name()))
	}
	if flagExportCode {
		if err := p.ExportSource(e, info); err != nil {
			return nil, fmt.Errorf("export source for %q failed: %w", e.Path, err)
		}
		return e, nil
	}
	var foundGeneric bool
	for _, name := range pkg.Scope().Names() {
		if !token.IsExported(name) {
			continue
		}
		obj := pkg.Scope().Lookup(name)
		switch t := obj.(type) {
		case *types.Const:
			named := pkgName + "." + t.Name()
			if typ := t.Type().String(); strings.HasPrefix(typ, "untyped ") {
				e.UntypedConsts = append(e.UntypedConsts, fmt.Sprintf("%q: {%q, %v}", t.Name(), t.Type().String(), p.constToLit(named, t.Val())))
			} else {
				e.TypedConsts = append(e.TypedConsts, fmt.Sprintf("%q : {reflect.TypeOf(%v), %v}", t.Name(), pkgName+"."+t.Name(), p.constToLit(named, t.Val())))
			}
			e.usedPkg = true
		case *types.Var:
			e.Vars = append(e.Vars, fmt.Sprintf("%q : reflect.ValueOf(&%v)", t.Name(), pkgName+"."+t.Name()))
			e.usedPkg = true
		case *types.Func:
			if hasTypeParam(t.Type()) {
				if !flagExportSource && !flagExportCode {
					log.Println("skip typeparam", t)
				}
				foundGeneric = true
				continue
			}
			e.Funcs = append(e.Funcs, fmt.Sprintf("%q : reflect.ValueOf(%v)", t.Name(), pkgName+"."+t.Name()))
			e.usedPkg = true
		case *types.TypeName:
			if hasTypeParam(t.Type()) {
				if !flagExportSource && !flagExportCode {
					log.Println("skip typeparam", t)
				}
				foundGeneric = true
				continue
			}
			if t.IsAlias() {
				name := obj.Name()
				switch typ := obj.Type().(type) {
				case *types.Basic:
					e.AliasTypes = append(e.AliasTypes, fmt.Sprintf("%q: reflect.TypeOf((*%v)(nil)).Elem()", name, typ.Name()))
				// case *types.Named:
				// 	e.AliasTypes = append(e.AliasTypes, fmt.Sprintf("%q: reflect.TypeOf((*%v.%v)(nil)).Elem()", name, sname, name))
				default:
					e.AliasTypes = append(e.AliasTypes, fmt.Sprintf("%q: reflect.TypeOf((*%v.%v)(nil)).Elem()", name, sname, name))
				}
				e.usedPkg = true
				continue
			}
			typeName := t.Name()
			if types.IsInterface(t.Type()) {
				e.Interfaces = append(e.Interfaces, fmt.Sprintf("%q : reflect.TypeOf((*%v.%v)(nil)).Elem()", typeName, pkgName, typeName))
			} else {
				e.NamedTypes = append(e.NamedTypes, fmt.Sprintf("%q : reflect.TypeOf((*%v.%v)(nil)).Elem()", typeName, pkgName, typeName))
			}
			e.usedPkg = true
		default:
			log.Panicf("unreachable %v %T\n", name, t)
		}
	}
	if flagExportSource && foundGeneric {
		if err := p.ExportSource(e, info); err != nil {
			return nil, fmt.Errorf("export source for %q failed: %w", e.Path, err)
		}
	}
	return e, nil
}
