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
	"fmt"
	"reflect"
	"sync"
)

// TypedCallbackMaker builds a Go function value for an interpreted function of
// a specific signature. The result should be a method value whose receiver
// embeds FuncVal as its first field, so the interpreter can recover the
// payload without reflect.MakeFunc.
type TypedCallbackMaker func(FuncVal) reflect.Value

// FuncVal is the receiver payload for typed callback method values.
// Embed it as the first field of a value receiver, then register the method:
//
//	type callback struct{ ixgo.FuncVal }
//	func (c callback) Int(n int) int { return c.Call(n).(int) }
//
//	func init() {
//	    ixgo.RegisterTypedCallbackFunc(func(c ixgo.FuncVal) func(int) int {
//	        return callback{c}.Int
//	    })
//	}
//
// Interpreted functions of that type then call as ordinary Go functions.
// func() is built in. Import github.com/goplus/ixgo/typed for common func() T results.
type FuncVal struct {
	interp *Interp
	pfn    *function
	typ    reflect.Type
	env    []value
}

// Call runs the interpreted function and returns its result.
// A function with no result returns nil. Multiple results are returned as a Tuple.
func (c FuncVal) Call(args ...any) any {
	return c.interp.callFunction(c.interp.tryDeferFrame(), c.pfn, args, c.env)
}

// CallDiscard runs the interpreted function and discards its results.
func (c FuncVal) CallDiscard(args ...any) {
	c.interp.callFunctionDiscardsResult(c.interp.tryDeferFrame(), c.pfn, args, c.env)
}

func (c FuncVal) callReflect(args []reflect.Value) []reflect.Value {
	return c.interp.callFunctionByReflect(c.interp.tryDeferFrame(), c.pfn, c.typ, args, c.env)
}

type voidCallback struct {
	FuncVal
}

func (c voidCallback) call() {
	c.CallDiscard()
}

var (
	callbackVoidType = reflect.TypeFor[func()]()
	callbackVoidPC   = reflect.ValueOf(voidCallback{}.call).Pointer()
)

func makeTypedFunction(c FuncVal, typ reflect.Type) (reflect.Value, bool) {
	if typ == callbackVoidType {
		return reflect.ValueOf(voidCallback{c}.call), true
	}
	if maker, ok := lookupTypedCallback(typ); ok {
		return maker(c), true
	}
	return reflect.Value{}, false
}

type typedCallbackEntry struct {
	maker TypedCallbackMaker
	pc    uintptr
}

var (
	typedCallbackMu     sync.RWMutex
	typedCallbackMakers = map[reflect.Type]typedCallbackEntry{}
	typedCallbackPCs    = map[uintptr]struct{}{}
)

// RegisterTypedCallback registers a maker for functions of type typ.
// Later registrations for the same type replace earlier ones.
// Passing a nil maker removes the registration.
// Do not unregister or replace a signature while function values of that
// type are still in use.
//
// The maker must return a method value of type typ whose receiver embeds
// FuncVal as its first field. On gc, registration panics if the maker does
// not satisfy that contract. func() is built in and is not taken from the
// registry.
func RegisterTypedCallback(typ reflect.Type, maker TypedCallbackMaker) {
	if typ == nil || typ.Kind() != reflect.Func {
		panic("ixgo: typed callback type must be a function")
	}
	var pc uintptr
	if maker != nil {
		pc = validateTypedCallback(typ, maker)
	}
	typedCallbackMu.Lock()
	defer typedCallbackMu.Unlock()
	if old, ok := typedCallbackMakers[typ]; ok {
		delete(typedCallbackPCs, old.pc)
		delete(typedCallbackMakers, typ)
	}
	if maker == nil {
		return
	}
	if pc != 0 {
		typedCallbackPCs[pc] = struct{}{}
	}
	typedCallbackMakers[typ] = typedCallbackEntry{maker: maker, pc: pc}
}

// RegisterTypedCallbackFunc registers a typed callback for signature F.
//
//	type callback struct{ ixgo.FuncVal }
//	func (c callback) Int(n int) int { return c.Call(n).(int) }
//
//	func init() {
//	    ixgo.RegisterTypedCallbackFunc(func(c ixgo.FuncVal) func(int) int {
//	        return callback{c}.Int
//	    })
//	}
//
// After registration, run source that uses that signature:
//
//	src := `package main
//	func main() {
//		add := func(x int) int { return x + 1 }
//		println(add(41))
//	}
//	`
//	ixgo.RunFile("main.go", src, nil, 0)
//
// bind must return a method value whose receiver embeds FuncVal as its first
// field. Passing a nil bind removes the registration.
func RegisterTypedCallbackFunc[F any](bind func(FuncVal) F) {
	typ := reflect.TypeFor[F]()
	if bind == nil {
		RegisterTypedCallback(typ, nil)
		return
	}
	RegisterTypedCallback(typ, func(c FuncVal) reflect.Value {
		return reflect.ValueOf(bind(c))
	})
}

func checkTypedCallbackType(typ reflect.Type, maker TypedCallbackMaker) {
	v := maker(FuncVal{})
	if !v.IsValid() || v.Kind() != reflect.Func || v.Type() != typ {
		panic(fmt.Sprintf("ixgo: typed callback maker returned %v, want %v", v.Type(), typ))
	}
	if v.IsNil() {
		panic("ixgo: typed callback maker returned a nil function")
	}
}

func lookupTypedCallback(typ reflect.Type) (TypedCallbackMaker, bool) {
	typedCallbackMu.RLock()
	e, ok := typedCallbackMakers[typ]
	typedCallbackMu.RUnlock()
	if !ok {
		return nil, false
	}
	return e.maker, true
}

func isTypedCallbackPC(pc uintptr) bool {
	typedCallbackMu.RLock()
	_, ok := typedCallbackPCs[pc]
	typedCallbackMu.RUnlock()
	return ok
}
