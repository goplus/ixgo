//go:build !go1.18
// +build !go1.18

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

	"golang.org/x/tools/go/ssa"
)

const (
	enabledTypeParam = false
)

func hasTypeParam(t types.Type) bool {
	return false
}

type nestedStack struct {
}

func (r *TypesRecord) EnterInstance(fn *ssa.Function) {
}
func (r *TypesRecord) LeaveInstance(fn *ssa.Function) {
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
	name = obj.Name()
	return
}

func (r *TypesRecord) LookupReflect(typ types.Type) (rt reflect.Type, ok bool, nested bool) {
	rt, ok = r.loader.LookupReflect(typ)
	if !ok {
		if rt := r.tcache.At(typ); rt != nil {
			return rt.(reflect.Type), true, false
		}
	}
	return
}

func (r *TypesRecord) lookupCache(typ types.Type) (rt reflect.Type, ok bool, nested bool) {
	if rt := r.tcache.At(typ); rt != nil {
		return rt.(reflect.Type), true, false
	}
	return
}

func (r *TypesLoader) hasTypeArgs(rt reflect.Type) bool {
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
	}
}
