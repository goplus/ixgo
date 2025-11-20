//go:build go1.18
// +build go1.18

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
	"go/ast"
	"go/types"
	"reflect"
	"strconv"
	"strings"

	"github.com/goplus/ixgo/internal/typesalias"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/types/typeutil"
)

const (
	enabledTypeParam = true
)

func hasTypeParam(typ types.Type) bool {
	return typesalias.HasTypeParam(typ)
}

type nestedStack struct {
	targs []string
	cache []*typeutil.Map
}

func (s *nestedStack) Push(targs string, cache *typeutil.Map) {
	s.targs = append(s.targs, targs)
	s.cache = append(s.cache, cache)
}

func (s *nestedStack) Pop() (targs string, cache *typeutil.Map) {
	n := len(s.targs)
	if n >= 1 {
		targs = s.targs[n-1]
		cache = s.cache[n-1]
	}
	s.targs = s.targs[:n-1]
	s.cache = s.cache[:n-1]
	return
}

func (r *TypesRecord) typeId(typ types.Type, t reflect.Type) string {
	path := t.PkgPath()
	if path == "" {
		return t.String()
	}
	name := path + "." + t.Name()
	if named, ok := typ.(*types.Named); ok {
		if n := r.nested[named.Origin()]; n != 0 {
			name += "·" + strconv.Itoa(n)
		}
	}
	return name
}

func (r *TypesRecord) EnterInstance(fn *ssa.Function) {
	r.ninst++
	r.ncache = &typeutil.Map{}
	tp := fn.TypeParams()
	for i := 0; i < tp.Len(); i++ {
		rt, _ := r.ToType(fn.TypeArgs()[i])
		r.ncache.Set(tp.At(i), rt)
	}
	r.nstack.Push(r.fntargs, r.ncache)
	r.fntargs = r.parseFuncTypeArgs(fn)
}

func (r *TypesRecord) LeaveInstance(fn *ssa.Function) {
	r.ninst--
	r.fntargs, r.ncache = r.nstack.Pop()
}

func (r *TypesRecord) parseFuncTypeArgs(fn *ssa.Function) (targs string) {
	typeargs := fn.TypeArgs()
	if len(typeargs) == 0 {
		return
	}
	var args []string
	for _, typ := range typeargs {
		rt, _ := r.ToType(typ)
		args = append(args, r.typeId(typ, rt))
	}
	return strings.Join(args, ",")
}

func (r *TypesRecord) extractNamed(named *types.Named, totype bool) (pkgpath string, name string, typeargs bool, nested bool) {
	obj := named.Obj()
	if pkg := obj.Pkg(); pkg != nil {
		if pkg.Name() == "main" {
			pkgpath = "main"
		} else {
			pkgpath = pkg.Path()
		}
	}
	if r.fntargs != "" && r.nested[named.Origin()] != 0 {
		nested = true
	}
	name = obj.Name()
	var ids string = r.fntargs
	if args := named.TypeArgs(); args != nil {
		typeargs = true
		var targs []string
		for i := 0; i < args.Len(); i++ {
			if totype {
				typ := args.At(i)
				rt, _ := r.ToType(typ)
				targs = append(targs, r.typeId(typ, rt))
			} else {
				targs = append(targs, args.At(i).String())
			}
		}
		if ids != "" {
			ids += ";"
		}
		ids += strings.Join(targs, ",")
	}
	if ids != "" {
		name += "[" + ids + "]"
	}
	return
}

func (r *TypesRecord) LookupReflect(typ types.Type) (rt reflect.Type, ok bool, nested bool) {
	rt, ok = r.loader.LookupReflect(typ)
	if !ok {
		return r.lookupCache(typ)
	}
	return
}

func (r *TypesRecord) lookupCache(typ types.Type) (rt reflect.Type, ok bool, nested bool) {
	if r.ninst > 0 {
		if rt := r.ncache.At(typ); rt != nil {
			return rt.(reflect.Type), true, true
		}
		n := len(r.nstack.cache)
		for i := n; i > 0; i-- {
			if rt := r.nstack.cache[i-1].At(typ); rt != nil {
				return rt.(reflect.Type), true, true
			}
		}
	}
	if rt := r.tcache.At(typ); rt != nil {
		return rt.(reflect.Type), true, false
	}
	return
}

func (r *TypesLoader) hasTypeArgs(rt reflect.Type) bool {
	switch rt.Kind() {
	case reflect.Pointer:
		return r.hasTypeArgs(rt.Elem())
	case reflect.Slice:
		return r.hasTypeArgs(rt.Elem())
	case reflect.Array:
		return r.hasTypeArgs(rt.Elem())
	case reflect.Chan:
		return r.hasTypeArgs(rt.Elem())
	case reflect.Map:
		return r.hasTypeArgs(rt.Key()) || r.hasTypeArgs(rt.Elem())
	case reflect.Struct:
		if pkgPath := rt.PkgPath(); pkgPath != "" {
			if pkg, ok := r.packages[pkgPath]; ok && pkg.Complete() {
				name := rt.Name()
				var typeArgs string
				if n := strings.Index(name, "["); n != -1 {
					if end := strings.LastIndex(name, "]"); end != -1 {
						typeArgs = name[n+1 : end]
						name = name[:n]
					}
				}
				if obj := pkg.Scope().Lookup(name); obj != nil && len(typeArgs) > 0 {
					return true
				}
			}
		}
	}
	return false
}

func newTypesInfo() *types.Info {
	return &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Scopes:     make(map[ast.Node]*types.Scope),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Instances:  make(map[*ast.Ident]types.Instance),
	}
}
