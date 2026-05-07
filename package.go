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
	"go/constant"
	"go/token"
	"go/types"
	"log"
	"reflect"
	"sort"
	"sync"

	"github.com/goplus/ixgo/alias"
)

var (
	registerPkgs   = make(map[string]*pkgInfo)
	registerPatchs = make(map[string][]interface{})
)

// PackageList return all register packages
func PackageList() (list []string) {
	for pkg := range registerPkgs {
		list = append(list, pkg)
	}
	sort.Strings(list)
	return
}

// LookupPackage lookup register pkgs
func LookupPackage(name string) (pkg *Package, ok bool) {
	if pkg, ok := registerPkgs[name]; ok {
		return pkg.Package(), true
	}
	return
}

// RegisterPackage register pkg
func RegisterPackage(pkg string, fn func() *Package) {
	registerPkgs[pkg] = &pkgInfo{
		load: fn,
	}
}

type pkgInfo struct {
	once sync.Once
	pkg  *Package
	load func() *Package
}

func (p *pkgInfo) Package() *Package {
	p.once.Do(func() {
		p.pkg = p.load()
	})
	return p.pkg
}

// RegisterPatch register pkg with "pkg@patch"
func RegisterPatch(pkg string, src ...interface{}) {
	registerPatchs[pkg] = append(registerPatchs[pkg], src...)
}

type UntypedConst struct {
	Typ   string
	Value constant.Value
}

type Package struct {
	Interfaces    map[string]reflect.Type
	NamedTypes    map[string]reflect.Type
	AliasTypes    map[string]reflect.Type
	Vars          map[string]interface{}
	Funcs         map[string]interface{}
	TypedConsts   map[string]interface{}
	UntypedConsts map[string]UntypedConst
	Deps          map[string]string // path -> name
	Alias         map[string]alias.Type
	Name          string
	Path          string
	Source        string
	Import        func(fset *token.FileSet, pkgs map[string]*types.Package) (*types.Package, error)
}

var (
	externValues = make(map[string]reflect.Value)
)

// RegisterExternal is register external variable address or func
func RegisterExternal(key string, i interface{}) {
	if i == nil {
		delete(externValues, key)
		return
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Func, reflect.Ptr:
		externValues[key] = v
	default:
		log.Printf("register external must variable address or func. not %v\n", v.Kind())
	}
}
