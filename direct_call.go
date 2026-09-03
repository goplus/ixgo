/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
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
	"go/types"
	"reflect"
	"strings"
	"sync"

	"golang.org/x/tools/go/ssa"
)

// DirectCallContext gives a generated direct-call adapter access to the
// current call's arguments and result register.
type DirectCallContext struct {
	frame  *frame
	result register
	args   []register
}

// DirectCallArg returns the argument at index using its generated static type.
// A nil interface value is converted to the zero value of T.
func DirectCallArg[T any](ctx DirectCallContext, index int) T {
	v := ctx.frame.reg(ctx.args[index])
	if v == nil {
		var zero T
		return zero
	}
	return v.(T)
}

// SetResult stores the result produced by a generated direct-call adapter.
func (ctx DirectCallContext) SetResult(result any) {
	ctx.frame.setReg(ctx.result, result)
}

// DirectCallAdapter is a qexp-generated, reflection-free call adapter.
type DirectCallAdapter func(DirectCallContext)

var directCallBindings sync.Map  // map[string]DirectCallAdapter
var directCallPackages sync.Map  // map[string]struct{}
var directCallMethods sync.Map   // map[reflect.Type]*sync.Map(method -> DirectCallAdapter)
var directCallSelectors sync.Map // map[string]*sync.Map(selector -> struct{})

// RegisterDirectCalls registers generated direct-call bindings for pkgPath.
// Registration must finish before constructing an interpreter that uses the
// package. Later registrations replace bindings with the same selector.
func RegisterDirectCalls(pkgPath string, bindings map[string]DirectCallAdapter) {
	directCallPackages.Store(pkgPath, struct{}{})
	for selector, adapter := range bindings {
		externalKey := directCallSymbol(pkgPath, selector)
		directCallBindings.Store(externalKey, adapter)
		selectors, _ := directCallSelectors.LoadOrStore(pkgPath, new(sync.Map))
		selectors.(*sync.Map).Store(selector, struct{}{})
	}
	if pkg, ok := registerPkgs[pkgPath].(*baseLoad); ok {
		registerDirectCallMethods(pkgPath, pkg.pkg)
	}
}

func hasDirectCalls(pkgPath string) bool {
	_, ok := directCallPackages.Load(pkgPath)
	return ok
}

func directCallSymbol(pkgPath, selector string) string {
	if strings.HasPrefix(selector, "(*") {
		return "(*" + pkgPath + "." + selector[2:]
	}
	if strings.HasPrefix(selector, "(") {
		if end := strings.IndexByte(selector, ')'); end > 1 {
			return "(" + pkgPath + "." + selector[1:end] + selector[end:]
		}
	}
	if receiver, method, ok := strings.Cut(selector, "."); ok {
		return "(" + pkgPath + "." + receiver + ")." + method
	}
	return pkgPath + "." + selector
}

func lookupDirectCallBinding(interp *Interp, key string) (DirectCallAdapter, bool) {
	if _, overridden := interp.ctx.override.Load(key); overridden {
		return nil, false
	}
	if _, overridden := externValues.Load(key); overridden {
		return nil, false
	}
	value, ok := directCallBindings.Load(key)
	if !ok {
		return nil, false
	}
	binding := value.(DirectCallAdapter)
	if binding == nil {
		return nil, false
	}
	return binding, true
}

func resolveStaticDirectCall(interp *Interp, fn *ssa.Function) (DirectCallAdapter, bool) {
	externalKey := fn.String()
	return lookupDirectCallBinding(interp, externalKey)
}

func resolveInvokeDirectCall(receiver reflect.Type, method string) (DirectCallAdapter, bool) {
	value, ok := directCallMethods.Load(receiver)
	if !ok {
		return nil, false
	}
	value, ok = value.(*sync.Map).Load(method)
	if !ok {
		return nil, false
	}
	adapter := value.(DirectCallAdapter)
	return adapter, adapter != nil
}

func registerDirectCallMethods(pkgPath string, pkg *Package) {
	selectors, ok := directCallSelectors.Load(pkgPath)
	if !ok {
		return
	}
	selectors.(*sync.Map).Range(func(key, _ any) bool {
		selector := key.(string)
		receiver, method, ok := registeredDirectCallMethod(pkg, selector)
		if !ok {
			return true
		}
		adapter, ok := directCallBindings.Load(directCallSymbol(pkgPath, selector))
		if !ok {
			return true
		}
		methods, _ := directCallMethods.LoadOrStore(receiver, new(sync.Map))
		methods.(*sync.Map).Store(method, adapter.(DirectCallAdapter))
		return true
	})
}

func registeredDirectCallMethod(pkg *Package, selector string) (reflect.Type, string, bool) {
	var receiverName, method string
	var pointer bool
	switch {
	case strings.HasPrefix(selector, "(*"):
		end := strings.Index(selector, ").")
		if end < 3 {
			return nil, "", false
		}
		receiverName, method, pointer = selector[2:end], selector[end+2:], true
	case strings.HasPrefix(selector, "("):
		end := strings.Index(selector, ").")
		if end < 2 {
			return nil, "", false
		}
		receiverName, method = selector[1:end], selector[end+2:]
	default:
		var ok bool
		receiverName, method, ok = strings.Cut(selector, ".")
		if !ok {
			return nil, "", false
		}
	}
	receiver, ok := pkg.NamedTypes[receiverName]
	if !ok || method == "" {
		return nil, "", false
	}
	if pointer {
		receiver = reflect.PtrTo(receiver)
	}
	return receiver, method, true
}

func makeStaticDirectCallInstr(interp *Interp, fn *ssa.Function, result register, args []register) func(*frame) {
	if !directCallTypesCompatible(interp, fn.Signature) {
		return nil
	}
	adapter, ok := resolveStaticDirectCall(interp, fn)
	if !ok {
		return nil
	}
	return func(fr *frame) {
		interp.invokeDirectCall(fr, adapter, result, args)
	}
}

// direct calls use host Go values. A source-backed named type belongs to the
// interpreter's type universe, so crossing that boundary would produce two
// reflect types with the same printed name but different identities.
func directCallTypesCompatible(interp *Interp, sig *types.Signature) bool {
	if sig.Recv() != nil && !directCallTypeCompatible(interp, sig.Recv().Type()) {
		return false
	}
	if !directCallTupleCompatible(interp, sig.Params()) || !directCallTupleCompatible(interp, sig.Results()) {
		return false
	}
	return true
}

func directCallTupleCompatible(interp *Interp, tuple *types.Tuple) bool {
	for i := 0; i < tuple.Len(); i++ {
		if !directCallTypeCompatible(interp, tuple.At(i).Type()) {
			return false
		}
	}
	return true
}

func directCallTypeCompatible(interp *Interp, typ types.Type) bool {
	switch t := typ.(type) {
	case *types.Named:
		if obj := t.Obj(); obj.Pkg() != nil && interp.ctx.SourcePackage(obj.Pkg().Path()) != nil {
			return false
		}
		return true
	case *types.Pointer:
		return directCallTypeCompatible(interp, t.Elem())
	case *types.Slice:
		return directCallTypeCompatible(interp, t.Elem())
	case *types.Array:
		return directCallTypeCompatible(interp, t.Elem())
	case *types.Map:
		return directCallTypeCompatible(interp, t.Key()) && directCallTypeCompatible(interp, t.Elem())
	case *types.Chan:
		return directCallTypeCompatible(interp, t.Elem())
	case *types.Signature:
		return directCallTypesCompatible(interp, t)
	case *types.Struct:
		for i := 0; i < t.NumFields(); i++ {
			if !directCallTypeCompatible(interp, t.Field(i).Type()) {
				return false
			}
		}
	}
	return true
}

func (i *Interp) invokeDirectCall(fr *frame, adapter DirectCallAdapter, result register, args []register) {
	i.trackDeferFrame(fr)
	adapter(DirectCallContext{frame: fr, result: result, args: args})
}

func (i *Interp) trackDeferFrame(fr *frame) {
	if fr.deferid != 0 {
		i.deferMap.Store(fr.deferid, fr)
	}
}
