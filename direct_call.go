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

var directCallBindings sync.Map // map[string]DirectCallAdapter

var directCallMethods sync.Map // map[reflect.Type]*sync.Map(method -> DirectCallAdapter)

// RegisterDirectCalls registers generated direct-call bindings for pkgPath.
// Registration must finish before constructing an interpreter that uses the
// package. Later registrations replace bindings with the same selector.
func RegisterDirectCalls(pkgPath string, bindings map[string]DirectCallAdapter) {
	for selector, adapter := range bindings {
		externalKey := directCallSymbol(pkgPath, selector)
		directCallBindings.Store(externalKey, adapter)
		if receiver, method, ok := registeredDirectCallMethod(pkgPath, selector); ok {
			value, _ := directCallMethods.LoadOrStore(receiver, new(sync.Map))
			value.(*sync.Map).Store(method, adapter)
		}
	}
}

func registeredDirectCallMethod(pkgPath, selector string) (reflect.Type, string, bool) {
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
	if receiverName == "" || method == "" {
		return nil, "", false
	}
	pkg, ok := LookupPackage(pkgPath)
	if !ok {
		return nil, "", false
	}
	receiver, ok := pkg.NamedTypes[receiverName]
	if !ok {
		return nil, "", false
	}
	if pointer {
		receiver = reflect.PtrTo(receiver)
	}
	return receiver, method, true
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

func makeStaticDirectCallInstr(interp *Interp, fn *ssa.Function, result register, args []register) func(*frame) {
	adapter, ok := resolveStaticDirectCall(interp, fn)
	if !ok {
		return nil
	}
	return func(fr *frame) {
		interp.invokeDirectCall(fr, adapter, result, args)
	}
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
