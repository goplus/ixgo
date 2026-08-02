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
	"sync"
	"sync/atomic"

	"golang.org/x/tools/go/ssa"
)

type invokeSignature struct {
	params   []reflect.Type
	results  []reflect.Type
	variadic bool
}

func makeInvokeSignature(interp *Interp, sig *types.Signature) invokeSignature {
	signature := invokeSignature{
		params:   make([]reflect.Type, sig.Params().Len()),
		results:  make([]reflect.Type, sig.Results().Len()),
		variadic: sig.Variadic(),
	}
	for i := range signature.params {
		signature.params[i] = interp.preToType(sig.Params().At(i).Type())
	}
	for i := range signature.results {
		signature.results[i] = interp.preToType(sig.Results().At(i).Type())
	}
	return signature
}

func (signature invokeSignature) matches(target, receiver reflect.Type) bool {
	if target.NumIn() != len(signature.params)+1 || target.In(0) != receiver || target.IsVariadic() != signature.variadic {
		return false
	}
	for i, param := range signature.params {
		if target.In(i+1) != param {
			return false
		}
	}
	if target.NumOut() != len(signature.results) {
		return false
	}
	for i, result := range signature.results {
		if target.Out(i) != result {
			return false
		}
	}
	return true
}

func resolveInvokeDirectCall(interp *Interp, receiver reflect.Type, method string, signature invokeSignature, reflectFunc reflect.Value) (DirectCallAdapter, bool) {
	pkgPath, key, ok := directCallMethodKey(receiver, method)
	if !ok {
		return nil, false
	}
	binding, ok := lookupDirectCallBinding(interp, pkgPath, key)
	if !ok || !signature.matches(binding.Target.Type(), receiver) || !binding.matchesTarget(reflectFunc, reflectFunc.Type()) {
		return nil, false
	}
	return binding.Adapter, true
}

// invokeTarget is the final dispatch target for one concrete receiver type.
// reflectFunc remains available when direct is set so Go and defer instructions
// can preserve eager argument capture and execute later. A zero target is a
// stable method lookup miss.
type invokeTarget struct {
	interpreted *function
	direct      DirectCallAdapter
	reflectFunc reflect.Value
}

type invokeCacheEntry struct {
	receiver reflect.Type
	target   invokeTarget
}

// invokeResolver caches interface dispatch per SSA callsite and concrete
// receiver type. Entries are immutable and safe to share between interpreted
// goroutines. The recent entry keeps monomorphic callsites off sync.Map's
// polymorphic path.
type invokeResolver struct {
	interp    *Interp
	name      string
	signature invokeSignature

	recent  atomic.Pointer[invokeCacheEntry]
	targets sync.Map // map[reflect.Type]*invokeCacheEntry
}

func newInvokeResolver(interp *Interp, call *ssa.CallCommon) *invokeResolver {
	if !call.IsInvoke() {
		return nil
	}
	return &invokeResolver{
		interp:    interp,
		name:      call.Method.Name(),
		signature: makeInvokeSignature(interp, call.Method.Type().(*types.Signature)),
	}
}

func (r *invokeResolver) resolve(receiver reflect.Type) invokeTarget {
	if entry := r.recent.Load(); entry != nil && entry.receiver == receiver {
		return entry.target
	}
	if cached, ok := r.targets.Load(receiver); ok {
		entry := cached.(*invokeCacheEntry)
		r.recent.Store(entry)
		return entry.target
	}

	target, cacheable := r.resolveUncached(receiver)
	if !cacheable {
		return target
	}
	entry := &invokeCacheEntry{receiver: receiver, target: target}
	actual, _ := r.targets.LoadOrStore(receiver, entry)
	entry = actual.(*invokeCacheEntry)
	r.recent.Store(entry)
	return entry.target
}

func (r *invokeResolver) resolveUncached(receiver reflect.Type) (invokeTarget, bool) {
	if methods, ok := r.interp.msets[receiver]; ok {
		if fn, ok := methods[r.name]; ok {
			return invokeTarget{interpreted: r.interp.funcs[fn]}, true
		}
		if reflectFunc, ok := findUserMethod(receiver, r.name); ok {
			// reflectx owns this wrapper. ResetIcall and ResetAllIcall may
			// invalidate it, so resolve it again for each invocation.
			return invokeTarget{reflectFunc: reflectFunc}, false
		}
		return invokeTarget{}, true
	}

	reflectFunc, ok := findExternMethod(receiver, r.name)
	if !ok {
		return invokeTarget{}, true
	}
	target := invokeTarget{reflectFunc: reflectFunc}
	if direct, ok := resolveInvokeDirectCall(r.interp, receiver, r.name, r.signature, reflectFunc); ok {
		target.direct = direct
	}
	return target, true
}
