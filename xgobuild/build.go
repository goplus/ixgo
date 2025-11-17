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

package xgobuild

import (
	"bytes"
	"fmt"
	goast "go/ast"
	"go/types"
	"os"
	"path/filepath"

	"github.com/goplus/gogen"
	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/internal/typesalias"
	"github.com/goplus/ixgo/internal/typesutil"
	"github.com/goplus/mod/modfile"
	"github.com/goplus/xgo/ast"
	"github.com/goplus/xgo/cl"
	"github.com/goplus/xgo/parser"
	"github.com/goplus/xgo/token"
)

type Class = modfile.Class
type Project = modfile.Project
type Import = modfile.Import

const (
	// StaticLoad is a mode flag that disables dynamic Importer creation.
	// When StaticLoad is set, the Context will use only the base Loader
	// for package imports instead of creating a dynamic ixgo.Importer.
	// This mode is useful for scenarios where all packages are pre-loaded
	// or when dynamic importing is not needed.
	StaticLoad = ixgo.LastMode * 2
)

var (
	projects = make(map[string]*modfile.Project)
)

func RegisterClassFileType(ext string, class string, works []*modfile.Class, pkgPaths ...string) {
	cls := &modfile.Project{
		Ext:      ext,
		Class:    class,
		Works:    works,
		PkgPaths: pkgPaths,
	}
	if ext != "" {
		projects[ext] = cls
	}
	for _, w := range works {
		projects[w.Ext] = cls
	}
}

func RegisterProject(proj *modfile.Project) {
	if proj.Ext != "" {
		projects[proj.Ext] = proj
	}
	for _, w := range proj.Works {
		projects[w.Ext] = proj
	}
}

func init() {
	cl.SetDebug(cl.FlagNoMarkAutogen)
	ixgo.RegisterFileProcess(".xgo", BuildFile)
	ixgo.RegisterFileProcess(".gop", BuildFile)
	ixgo.RegisterFileProcess(".gox", BuildFile)
	ixgo.RegisterFileProcess(".gsh", BuildFile)
	RegisterClassFileType(".gmx", "Game", []*Class{{Ext: ".spx", Class: "Sprite"}}, "github.com/goplus/spx", "math")
	RegisterClassFileType(".gsh", "App", nil, "github.com/qiniu/x/gsh", "math")
	RegisterClassFileType("_test.gox", "App", []*Class{{Ext: "_test.gox", Class: "Case"}}, "github.com/goplus/xgo/test")
}

func BuildFile(ctx *ixgo.Context, filename string, src interface{}) (data []byte, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("compile %v failed. %v", filename, r)
		}
	}()
	c := NewContext(ctx)
	pkg, err := c.ParseFile(filename, src)
	if err != nil {
		return nil, err
	}
	return pkg.ToSource()
}

func BuildFSDir(ctx *ixgo.Context, fs parser.FileSystem, dir string) (data []byte, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("compile %v failed. %v", dir, r)
		}
	}()
	c := NewContext(ctx)
	pkg, err := c.ParseFSDir(fs, dir)
	if err != nil {
		return nil, err
	}
	return pkg.ToSource()
}

func BuildDir(ctx *ixgo.Context, dir string) (data []byte, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("compile %v failed. %v", dir, r)
		}
	}()
	c := NewContext(ctx)
	pkg, err := c.ParseDir(dir)
	if err != nil {
		return nil, err
	}
	return pkg.ToSource()
}

type Package struct {
	Fset *token.FileSet
	Pkgs []*gogen.Package
}

func (p *Package) MainPkg() *gogen.Package {
	if len(p.Pkgs) == 0 {
		return nil
	}
	for _, pkg := range p.Pkgs {
		if pkg.Types.Name() == "main" {
			return pkg
		}
	}
	return p.Pkgs[0]
}

func (p *Package) ToSource() ([]byte, error) {
	pkg := p.MainPkg()
	if pkg == nil {
		return nil, os.ErrNotExist
	}
	var buf bytes.Buffer
	if err := pkg.WriteTo(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *Package) ToAst() (*goast.File, error) {
	pkg := p.MainPkg()
	if pkg == nil {
		return nil, os.ErrNotExist
	}
	return pkg.ASTFile(), nil
}

func (p *Package) ForEachFile(fn func(pkg *gogen.Package, fname string)) {
	for _, pkg := range p.Pkgs {
		pkg.ForEachFile(func(fname string, file *gogen.File) {
			fn(pkg, fname)
		})
	}
}

type Context struct {
	Context   *ixgo.Context
	FileSet   *token.FileSet
	Importer  types.Importer
	Loader    ixgo.Loader
	pkgs      map[string]*types.Package
	resetPkgs []func()
}

func ClassKind(fname string) (isProj, ok bool) {
	ext := modfile.ClassExt(fname)
	switch ext {
	case ".gmx", ".gsh":
		return true, true
	case ".spx":
		return fname == "main.spx", true
	default:
		if c, ok := projects[ext]; ok {
			for _, w := range c.Works {
				if w.Ext == ext {
					if ext != c.Ext || fname != "main"+ext {
						return false, true
					}
					break
				}
			}
			return true, true
		}
	}
	return
}

func NewContext(ctx *ixgo.Context) *Context {
	if ctx.IsEvalMode() {
		ctx = ixgo.NewContext(0)
	}
	ctx.Mode |= ixgo.CheckGopOverloadFunc
	c := &Context{Context: ctx, FileSet: token.NewFileSet(), Loader: ctx.Loader,
		pkgs: make(map[string]*types.Package)}
	if ctx.Mode&StaticLoad == 0 {
		c.Importer = ixgo.NewImporter(ctx)
	}
	return c
}

// rewindPkgs cleans up package patches by removing temporarily added scope objects
// to restore them to their original state. It's idempotent.
func (p *Context) rewindPkgs() {
	fns := p.resetPkgs
	p.resetPkgs = nil
	for _, fn := range fns {
		fn()
	}
}

func RegisterPackagePatch(ctx *ixgo.Context, path string, src interface{}) error {
	return ctx.RegisterPatch(path, src)
}

func (c *Context) isOwnedPackage(path string) bool {
	if c.Context.SourcePackage(path+"@patch") != nil {
		return true
	}
	if pkg, ok := ixgo.LookupPackage(path); ok {
		if _, ok := pkg.UntypedConsts["GopPackage"]; ok {
			return true
		}
	}
	return false
}

func (c *Context) importPath(path string) (pkg *types.Package, err error) {
	if c.Importer == nil || c.isOwnedPackage(path) {
		pkg, err = c.Loader.Import(path)
	} else {
		pkg, err = c.Importer.Import(path)
	}
	return
}

func (c *Context) Import(path string) (*types.Package, error) {
	if pkg, ok := c.pkgs[path]; ok {
		return pkg, nil
	}
	pkg, err := c.importPath(path)
	if err != nil {
		return pkg, err
	}
	c.pkgs[path] = pkg
	if sp := c.Context.SourcePackage(path + "@patch"); sp != nil {
		sp.Importer = c
		err := sp.Load()
		if err != nil {
			return nil, err
		}
		scope := sp.Package.Scope()
		names := scope.Names()
		rscope := pkg.Scope()
		rnames := make([]string, 0, len(names))
		ralts := make([]types.Object, 0, len(names))
		for _, name := range names {
			obj := scope.Lookup(name)
			switch v := obj.(type) {
			case *types.Const:
				obj = types.NewConst(obj.Pos(), pkg, obj.Name(), obj.Type(), v.Val())
			case *types.Var:
				obj = types.NewVar(obj.Pos(), pkg, obj.Name(), obj.Type())
			case *types.Func:
				obj = types.NewFunc(obj.Pos(), pkg, obj.Name(), obj.Type().(*types.Signature))
			case *types.TypeName:
				switch typ := obj.Type().(type) {
				case *typesalias.Alias:
					obj = types.NewTypeName(obj.Pos(), pkg, obj.Name(), nil)
					typesalias.NewAlias(obj.(*types.TypeName), typesalias.Rhs(typ))
				case *types.Named:
					var methods []*types.Func
					if n := typ.NumMethods(); n > 0 {
						methods = make([]*types.Func, n)
						for i := 0; i < n; i++ {
							methods[i] = typ.Method(i)
						}
					}
					obj = types.NewTypeName(obj.Pos(), pkg, obj.Name(), nil)
					types.NewNamed(obj.(*types.TypeName), typ.Underlying(), methods)
				}
			default:
				continue
			}
			alt := rscope.Insert(obj)
			if alt == nil {
				rnames = append(rnames, name)
			} else if alt.Pkg() == sp.Package {
				typesutil.ScopeDelete(rscope, name)
				rscope.Insert(obj)
				rnames = append(rnames, name)
				ralts = append(ralts, alt)
			}
		}
		c.resetPkgs = append(c.resetPkgs, func() {
			for _, name := range rnames {
				typesutil.ScopeDelete(rscope, name)
			}
			for _, alt := range ralts {
				rscope.Insert(alt)
			}
			delete(c.pkgs, path)
		})
	}
	return pkg, nil
}

func (c *Context) ParseDir(dir string) (*Package, error) {
	pkgs, err := parser.ParseDirEx(c.FileSet, dir, parser.Config{
		ClassKind: ClassKind,
	})
	if err != nil {
		return nil, err
	}
	return c.loadPackages(dir, pkgs)
}

func (c *Context) ParseFSDir(fs parser.FileSystem, dir string) (*Package, error) {
	pkgs, err := parser.ParseFSDir(c.FileSet, fs, dir, parser.Config{
		ClassKind: ClassKind,
	})
	if err != nil {
		return nil, err
	}
	return c.loadPackages(dir, pkgs)
}

func (c *Context) ParseFile(fname string, src interface{}) (*Package, error) {
	srcDir, _ := filepath.Split(fname)
	fnameRmGox := fname
	ext := filepath.Ext(fname)
	var err error
	var isProj, isClass, isNormalGox, rmGox bool
	switch ext {
	case ".go", ".gop", ".xgo":
	case ".gox":
		isClass = true
		t := fname[:len(fname)-4]
		if c := filepath.Ext(t); c != "" {
			ext, fnameRmGox, rmGox = c, t, true
		} else {
			isNormalGox = true
		}
		fallthrough
	default:
		if !isNormalGox {
			if isProj, isClass = ClassKind(fnameRmGox); !isClass {
				if rmGox {
					return nil, fmt.Errorf("not found Go+ class by ext %q for %q", ext, fname)
				}
				return nil, nil
			}
		}
	}
	if err != nil {
		return nil, err
	}
	mode := parser.ParseComments
	if isClass {
		mode |= parser.ParseGoPlusClass
	}
	f, err := parser.ParseFile(c.FileSet, fname, src, mode)
	if err != nil {
		return nil, err
	}
	f.IsProj, f.IsClass, f.IsNormalGox = isProj, isClass, isNormalGox
	name := f.Name.Name
	pkgs := map[string]*ast.Package{
		name: &ast.Package{
			Name: name,
			Files: map[string]*ast.File{
				fname: f,
			},
		},
	}
	return c.loadPackages(srcDir, pkgs)
}

func (c *Context) loadPackages(srcDir string, apkgs map[string]*ast.Package) (*Package, error) {
	defer c.rewindPkgs()
	var pkgs []*gogen.Package
	for _, apkg := range apkgs {
		pkg, err := c.loadPackage(srcDir, apkg)
		if err != nil {
			return nil, err
		}
		pkgs = append(pkgs, pkg)
	}
	return &Package{c.FileSet, pkgs}, nil
}

func (c *Context) loadPackage(srcDir string, apkg *ast.Package) (*gogen.Package, error) {
	if c.Context.Mode&ixgo.DisableCustomBuiltin == 0 {
		if f, err := ixgo.ParseBuiltin(c.FileSet, apkg.Name); err == nil {
			apkg.GoFiles = map[string]*goast.File{"_ixgo_builtin.go": f}
		}
	}
	conf := &cl.Config{Fset: c.FileSet}
	conf.Importer = c
	conf.LookupClass = func(ext string) (c *cl.Project, ok bool) {
		c, ok = projects[ext]
		return
	}
	if c.Context.IsEvalMode() {
		conf.NoSkipConstant = true
	}
	return cl.NewPackage("", apkg, conf)
}
