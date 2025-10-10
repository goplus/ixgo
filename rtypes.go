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
	"go/constant"
	"go/importer"
	"go/token"
	"go/types"
	"log"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"golang.org/x/tools/go/types/typeutil"
)

var (
	basicTypeNames = make(map[string]*types.Basic)
)

var (
	typesDummyStruct    = types.NewStruct(nil, nil)
	typesDummySig       = types.NewSignature(nil, nil, nil, false)
	typesDummySlice     = types.NewSlice(typesDummyStruct)
	typesError          = types.Universe.Lookup("error").Type()
	typesEmptyInterface = types.NewInterfaceType(nil, nil)
)

var (
	tyEmptyInterface = reflect.TypeOf((*interface{})(nil)).Elem()
	tyErrorInterface = reflect.TypeOf((*error)(nil)).Elem()
)

func init() {
	for i := types.Invalid; i <= types.UntypedNil; i++ {
		typ := types.Typ[i]
		basicTypeNames[typ.String()] = typ
	}
}

type TypesLoader struct {
	importer  types.Importer
	ctx       *Context
	tcache    *typeutil.Map
	curpkg    *Package
	packages  map[string]*types.Package
	installed map[string]*Package
	pkgloads  map[string]func() error
	rcache    map[reflect.Type]types.Type
	mode      Mode
}

// NewTypesLoader install package and readonly
func NewTypesLoader(ctx *Context, mode Mode) Loader {
	r := &TypesLoader{
		packages:  make(map[string]*types.Package),
		installed: make(map[string]*Package),
		pkgloads:  make(map[string]func() error),
		rcache:    make(map[reflect.Type]types.Type),
		tcache:    &typeutil.Map{},
		ctx:       ctx,
		mode:      mode,
	}
	r.packages["unsafe"] = types.Unsafe
	r.rcache[tyErrorInterface] = typesError
	r.rcache[tyEmptyInterface] = typesEmptyInterface
	r.importer = importer.Default()
	return r
}

func (r *TypesLoader) SetImport(path string, pkg *types.Package, load func() error) error {
	r.packages[path] = pkg
	if load != nil {
		r.pkgloads[path] = load
	}
	return nil
}

func (r *TypesLoader) Installed(path string) (pkg *Package, ok bool) {
	pkg, ok = r.installed[path]
	return
}

func (r *TypesLoader) Packages() (pkgs []*types.Package) {
	for _, pkg := range r.packages {
		pkgs = append(pkgs, pkg)
	}
	return
}

func (r *TypesLoader) LookupPackage(pkgpath string) (*types.Package, bool) {
	pkg, ok := r.packages[pkgpath]
	return pkg, ok
}

func lookupReflectByNamed(p *Package, name string) (reflect.Type, bool) {
	if rt, ok := p.NamedTypes[name]; ok {
		return rt, true
	}
	if rt, ok := p.AliasTypes[name]; ok {
		return rt, true
	}
	if rt, ok := p.Interfaces[name]; ok {
		return rt, true
	}
	return nil, false
}

func (r *TypesLoader) lookupRelfect(typ types.Type) (reflect.Type, bool) {
	var star bool
	if t, ok := typ.(*types.Pointer); ok {
		star = true
		typ = t.Elem()
	}
	if named, ok := typ.(*types.Named); ok {
		if pkg := named.Obj().Pkg(); pkg != nil {
			if p, ok := r.installed[pkg.Path()]; ok {
				if rt, ok := lookupReflectByNamed(p, named.Obj().Name()); ok {
					if star {
						rt = reflect.PtrTo(rt)
					}
					return rt, true
				}
			}
		}
	}
	return nil, false
}

func (r *TypesLoader) LookupReflect(typ types.Type) (reflect.Type, bool) {
	if rt := r.tcache.At(typ); rt != nil {
		return rt.(reflect.Type), true
	}
	if rt, ok := r.lookupRelfect(typ); ok {
		r.tcache.Set(typ, rt)
		return rt, ok
	}
	return nil, false
}

func (r *TypesLoader) LookupTypes(typ reflect.Type) (types.Type, bool) {
	t, ok := r.rcache[typ]
	return t, ok
}

func (r *TypesLoader) Import(path string) (*types.Package, error) {
	if p, ok := r.packages[path]; ok {
		if !p.Complete() {
			if load, ok := r.pkgloads[path]; ok {
				load()
			}
			if pkg, ok := registerPkgs[path]; ok {
				r.installed[path] = pkg
			}
		}
		return p, nil
	}
	pkg, ok := registerPkgs[path]
	if !ok {
		return nil, fmt.Errorf("not found package %v", path)
	}
	p := types.NewPackage(pkg.Path, pkg.Name)
	r.packages[path] = p
	for dep := range pkg.Deps {
		r.Import(dep)
	}
	if len(pkg.Source) > 0 {
		tp, ok := r.ctx.pkgs[pkg.Path]
		if !ok {
			var err error
			tp, err = r.ctx.addImportFile(pkg.Path, pkg.Name+".go", pkg.Source)
			if err != nil {
				return nil, err
			}
		}
		if err := tp.Load(); err != nil {
			log.Printf("ixgo warning: load pkg %v source error: %v\n", pkg.Path, err)
			return nil, err
		}
		tp.Register = true
		r.packages[path] = tp.Package
		r.installed[path] = pkg
		return tp.Package, nil
	}
	if err := r.installPackage(pkg); err != nil {
		return nil, err
	}
	var list []*types.Package
	for dep := range pkg.Deps {
		if p, ok := r.packages[dep]; ok {
			list = append(list, p)
		}
	}
	p.SetImports(list)
	p.MarkComplete()
	return p, nil
}

func (r *TypesLoader) installPackage(pkg *Package) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		r.curpkg = nil
	}()
	r.curpkg = pkg
	r.installed[pkg.Path] = pkg
	p, ok := r.packages[pkg.Path]
	if !ok {
		p = types.NewPackage(pkg.Path, pkg.Name)
		r.packages[pkg.Path] = p
	}
	for name, typ := range pkg.Interfaces {
		r.InsertInterface(p, name, typ)
	}
	for name, typ := range pkg.NamedTypes {
		if typ.Kind() == reflect.Struct {
			r.InsertNamedType(p, name, typ)
		}
	}
	for name, typ := range pkg.NamedTypes {
		if typ.Kind() != reflect.Struct {
			r.InsertNamedType(p, name, typ)
		}
	}
	for name, typ := range pkg.AliasTypes {
		r.InsertAlias(p, name, typ)
	}
	for name, fn := range pkg.Funcs {
		r.InsertFunc(p, name, fn)
	}
	for name, v := range pkg.Vars {
		r.InsertVar(p, name, v.Elem())
	}
	for name, c := range pkg.TypedConsts {
		r.InsertTypedConst(p, name, c)
	}
	for name, c := range pkg.UntypedConsts {
		r.InsertUntypedConst(p, name, c)
	}
	return
}

func (r *TypesLoader) InsertInterface(p *types.Package, name string, rt reflect.Type) {
	r.ToType(rt)
}

func (r *TypesLoader) InsertNamedType(p *types.Package, name string, rt reflect.Type) {
	r.ToType(rt)
}

func (r *TypesLoader) InsertAlias(p *types.Package, name string, rt reflect.Type) {
	typ := r.ToType(rt)
	p.Scope().Insert(types.NewTypeName(token.NoPos, p, name, typ))
}

func (r *TypesLoader) InsertFunc(p *types.Package, name string, v reflect.Value) {
	typ := r.ToType(v.Type())
	p.Scope().Insert(types.NewFunc(token.NoPos, p, name, typ.(*types.Signature)))
}

func (r *TypesLoader) InsertVar(p *types.Package, name string, v reflect.Value) {
	typ := r.ToType(v.Type())
	p.Scope().Insert(types.NewVar(token.NoPos, p, name, typ))
}

func (r *TypesLoader) InsertConst(p *types.Package, name string, typ types.Type, c constant.Value) {
	p.Scope().Insert(types.NewConst(token.NoPos, p, name, typ, c))
}

func splitPath(path string) (pkg string, name string, ok bool) {
	pos := strings.LastIndex(path, ".")
	if pos == -1 {
		return path, "", false
	}
	return path[:pos], path[pos+1:], true
}

func (r *TypesLoader) parserNamed(path string) (*types.Package, string) {
	if pkg, name, ok := splitPath(path); ok {
		if p := r.GetPackage(pkg); p != nil {
			return p, name
		}
	}
	panic(fmt.Errorf("parse path failed: %v", path))
}

func (r *TypesLoader) LookupType(typ string) types.Type {
	if t, ok := basicTypeNames[typ]; ok {
		return t
	}
	p, name := r.parserNamed(typ)
	return p.Scope().Lookup(name).Type()
}

func (r *TypesLoader) InsertTypedConst(p *types.Package, name string, v TypedConst) {
	typ := r.ToType(v.Typ)
	r.InsertConst(p, name, typ, v.Value)
}

func (r *TypesLoader) InsertUntypedConst(p *types.Package, name string, v UntypedConst) {
	var typ types.Type
	if t, ok := basicTypeNames[v.Typ]; ok {
		typ = t
	} else {
		typ = r.LookupType(v.Typ)
	}
	r.InsertConst(p, name, typ, v.Value)
}

func (r *TypesLoader) GetPackage(pkg string) *types.Package {
	if pkg == "" {
		return nil
	}
	if p, ok := r.packages[pkg]; ok {
		return p
	}
	var name string
	if r.curpkg != nil {
		name = r.curpkg.Deps[pkg]
	}
	if name == "" {
		pkgs := strings.Split(pkg, "/")
		name = pkgs[len(pkgs)-1]
	}
	p := types.NewPackage(pkg, name)
	r.packages[pkg] = p
	return p
}

func toTypeChanDir(dir reflect.ChanDir) types.ChanDir {
	switch dir {
	case reflect.RecvDir:
		return types.RecvOnly
	case reflect.SendDir:
		return types.SendOnly
	case reflect.BothDir:
		return types.SendRecv
	}
	panic("unreachable")
}

func (r *TypesLoader) Insert(v reflect.Value) {
	typ := r.ToType(v.Type())
	if v.Kind() == reflect.Func {
		name := runtime.FuncForPC(v.Pointer()).Name()
		names := strings.Split(name, ".")
		pkg := r.GetPackage(names[0])
		pkg.Scope().Insert(types.NewFunc(token.NoPos, pkg, names[1], typ.(*types.Signature)))
	}
}

func (r *TypesLoader) toMethod(pkg *types.Package, recv *types.Var, inoff int, rt reflect.Type) *types.Signature {
	numIn := rt.NumIn()
	numOut := rt.NumOut()
	in := make([]*types.Var, numIn-inoff)
	out := make([]*types.Var, numOut)
	for i := inoff; i < numIn; i++ {
		it := r.ToType(rt.In(i))
		in[i-inoff] = types.NewVar(token.NoPos, pkg, "", it)
	}
	for i := 0; i < numOut; i++ {
		it := r.ToType(rt.Out(i))
		out[i] = types.NewVar(token.NoPos, pkg, "", it)
	}
	return types.NewSignature(recv, types.NewTuple(in...), types.NewTuple(out...), rt.IsVariadic())
}

func (r *TypesLoader) toFunc(pkg *types.Package, rt reflect.Type) *types.Signature {
	numIn := rt.NumIn()
	numOut := rt.NumOut()
	in := make([]*types.Var, numIn)
	out := make([]*types.Var, numOut)
	// mock type
	variadic := rt.IsVariadic()
	if variadic {
		for i := 0; i < numIn-1; i++ {
			in[i] = types.NewVar(token.NoPos, pkg, "", typesDummyStruct)
		}
		in[numIn-1] = types.NewVar(token.NoPos, pkg, "", typesDummySlice)
	} else {
		for i := 0; i < numIn; i++ {
			in[i] = types.NewVar(token.NoPos, pkg, "", typesDummyStruct)
		}
	}
	for i := 0; i < numOut; i++ {
		out[i] = types.NewVar(token.NoPos, pkg, "", typesDummyStruct)
	}
	typ := types.NewSignature(nil, types.NewTuple(in...), types.NewTuple(out...), variadic)
	r.rcache[rt] = typ
	r.tcache.Set(typ, rt)
	// real type
	for i := 0; i < numIn; i++ {
		it := r.ToType(rt.In(i))
		in[i] = types.NewVar(token.NoPos, pkg, "", it)
	}
	for i := 0; i < numOut; i++ {
		it := r.ToType(rt.Out(i))
		out[i] = types.NewVar(token.NoPos, pkg, "", it)
	}
	return typ
}

func (r *TypesLoader) ToType(rt reflect.Type) types.Type {
	if t, ok := r.rcache[rt]; ok {
		return t
	}
	var isNamed bool
	var pkgPath string
	// check complete pkg named type
	if pkgPath = rt.PkgPath(); pkgPath != "" {
		if pkg, ok := r.packages[pkgPath]; ok && pkg.Complete() {
			if obj := pkg.Scope().Lookup(rt.Name()); obj != nil {
				typ := obj.Type()
				r.rcache[rt] = typ
				return typ
			}
		}
		isNamed = true
	}
	var typ types.Type
	var fields []*types.Var
	var imethods []*types.Func
	kind := rt.Kind()
	switch kind {
	case reflect.Invalid:
		typ = types.Typ[types.Invalid]
	case reflect.Bool:
		typ = types.Typ[types.Bool]
	case reflect.Int:
		typ = types.Typ[types.Int]
	case reflect.Int8:
		typ = types.Typ[types.Int8]
	case reflect.Int16:
		typ = types.Typ[types.Int16]
	case reflect.Int32:
		typ = types.Typ[types.Int32]
	case reflect.Int64:
		typ = types.Typ[types.Int64]
	case reflect.Uint:
		typ = types.Typ[types.Uint]
	case reflect.Uint8:
		typ = types.Typ[types.Uint8]
	case reflect.Uint16:
		typ = types.Typ[types.Uint16]
	case reflect.Uint32:
		typ = types.Typ[types.Uint32]
	case reflect.Uint64:
		typ = types.Typ[types.Uint64]
	case reflect.Uintptr:
		typ = types.Typ[types.Uintptr]
	case reflect.Float32:
		typ = types.Typ[types.Float32]
	case reflect.Float64:
		typ = types.Typ[types.Float64]
	case reflect.Complex64:
		typ = types.Typ[types.Complex64]
	case reflect.Complex128:
		typ = types.Typ[types.Complex128]
	case reflect.Array:
		elem := r.ToType(rt.Elem())
		typ = types.NewArray(elem, int64(rt.Len()))
	case reflect.Chan:
		elem := r.ToType(rt.Elem())
		dir := toTypeChanDir(rt.ChanDir())
		typ = types.NewChan(dir, elem)
	case reflect.Func:
		if !isNamed {
			typ = r.toMethod(nil, nil, 0, rt)
		} else {
			typ = typesDummySig
		}
	case reflect.Interface:
		n := rt.NumMethod()
		imethods = make([]*types.Func, n)
		for i := 0; i < n; i++ {
			im := rt.Method(i)
			sig := typesDummySig
			pkg := r.GetPackage(im.PkgPath)
			imethods[i] = types.NewFunc(token.NoPos, pkg, im.Name, sig)
		}
		typ = types.NewInterfaceType(imethods, nil)
	case reflect.Map:
		key := r.ToType(rt.Key())
		elem := r.ToType(rt.Elem())
		typ = types.NewMap(key, elem)
	case reflect.Ptr:
		elem := r.ToType(rt.Elem())
		typ = types.NewPointer(elem)
	case reflect.Slice:
		elem := r.ToType(rt.Elem())
		typ = types.NewSlice(elem)
	case reflect.String:
		typ = types.Typ[types.String]
	case reflect.Struct:
		n := rt.NumField()
		fields = make([]*types.Var, n)
		tags := make([]string, n)
		pkg := r.GetPackage(rt.PkgPath())
		for i := 0; i < n; i++ {
			f := rt.Field(i)
			ft := types.Typ[types.UnsafePointer] //r.ToType(f.Type)
			fields[i] = types.NewVar(token.NoPos, pkg, f.Name, ft)
			tags[i] = string(f.Tag)
		}
		typ = types.NewStruct(fields, tags)
	case reflect.UnsafePointer:
		typ = types.Typ[types.UnsafePointer]
	default:
		panic("unreachable")
	}
	var named *types.Named
	if isNamed {
		pkg := r.GetPackage(pkgPath)
		obj := types.NewTypeName(token.NoPos, pkg, rt.Name(), nil)
		named = types.NewNamed(obj, typ, nil)
		typ = named
		pkg.Scope().Insert(obj)
	}
	r.rcache[rt] = typ
	r.tcache.Set(typ, rt)
	if kind == reflect.Struct {
		n := rt.NumField()
		pkg := r.GetPackage(pkgPath)
		for i := 0; i < n; i++ {
			f := rt.Field(i)
			if enabledTypeParam && r.hasTypeArgs(f.Type) {
				continue
			}
			ft := r.ToType(f.Type)
			fields[i] = types.NewField(token.NoPos, pkg, f.Name, ft, f.Anonymous)
		}
	} else if kind == reflect.Interface {
		n := rt.NumMethod()
		pkg := r.GetPackage(rt.PkgPath())
		recv := types.NewVar(token.NoPos, pkg, "", typ)
		for i := 0; i < n; i++ {
			im := rt.Method(i)
			pkg := r.GetPackage(im.PkgPath)
			sig := r.toMethod(pkg, recv, 0, im.Type)
			(*object)(unsafe.Pointer(imethods[i])).typ = sig
		}
		typ.Underlying().(*types.Interface).Complete()
	}
	if named != nil {
		switch kind {
		case reflect.Func:
			named.SetUnderlying(r.toMethod(named.Obj().Pkg(), nil, 0, rt))
		}
		if kind != reflect.Interface {
			pkg := named.Obj().Pkg()
			skip := make(map[string]bool)
			recv := types.NewVar(token.NoPos, pkg, "", typ)
			for _, im := range allMethodX(rt) {
				var sig *types.Signature
				if im.Type != nil {
					sig = r.toMethod(pkg, recv, 1, im.Type)
				} else {
					sig = r.toMethod(pkg, recv, 0, tyUnusedFunc)
				}
				skip[im.Name] = true
				named.AddMethod(types.NewFunc(token.NoPos, pkg, im.Name, sig))
			}
			prt := reflect.PtrTo(rt)
			ptyp := r.ToType(prt)
			precv := types.NewVar(token.NoPos, pkg, "", ptyp)
			for _, im := range allMethodX(prt) {
				if skip[im.Name] {
					continue
				}
				var sig *types.Signature
				if im.Type != nil {
					sig = r.toMethod(pkg, precv, 1, im.Type)
				} else {
					sig = r.toMethod(pkg, precv, 0, tyUnusedFunc)
				}
				named.AddMethod(types.NewFunc(token.NoPos, pkg, im.Name, sig))
			}
		}
	}
	return typ
}

// go/types.object
type object struct {
	parent    *types.Scope
	pos       token.Pos
	pkg       *types.Package
	name      string
	typ       types.Type
	order_    uint32
	color_    uint32
	scopePos_ token.Pos
}
