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
)

// DirectFuncMaker builds a Go function value for an interpreted function of
// a specific signature. The result should be a method value whose receiver
// embeds DirectFuncVal as its first field, so the interpreter can recover the
// payload without reflect.MakeFunc.
type DirectFuncMaker func(DirectFuncVal) reflect.Value

// DirectFuncVal is the receiver payload for DirectFunc method values.
// Embed it as the first field of a value receiver, then register the method:
//
//	type fn struct{ ixgo.DirectFuncVal }
//	func (c fn) Int(n int) int { return c.Call(n).(int) }
//
//	func init() {
//	    ixgo.RegisterDirectFuncFor(func(c ixgo.DirectFuncVal) func(int) int {
//	        return fn{c}.Int
//	    })
//	}
//
// Interpreted functions of that type then call as ordinary Go functions.
// func() is built in. Import github.com/goplus/ixgo/directfunc for common func() T results.
type DirectFuncVal struct {
	interp *Interp
	pfn    *function
	typ    reflect.Type
	env    []value
}

// Call runs the interpreted function and returns its result.
// A function with no result returns nil. Multiple results are returned as a Tuple.
func (c DirectFuncVal) Call(args ...any) any {
	return c.interp.callFunction(c.interp.tryDeferFrame(), c.pfn, args, c.env)
}

// CallDiscard runs the interpreted function and discards its results.
func (c DirectFuncVal) CallDiscard(args ...any) {
	c.interp.callFunctionDiscardsResult(c.interp.tryDeferFrame(), c.pfn, args, c.env)
}

func (c DirectFuncVal) callReflect(args []reflect.Value) []reflect.Value {
	return c.interp.callFunctionByReflect(c.interp.tryDeferFrame(), c.pfn, c.typ, args, c.env)
}

type voidFunc struct {
	DirectFuncVal
}

func (c voidFunc) call() {
	c.CallDiscard()
}

var (
	directFuncVoidType = reflect.TypeFor[func()]()
	directFuncVoidPC   = reflect.ValueOf(voidFunc{}.call).Pointer()
)

func makeDirectFunction(c DirectFuncVal, typ reflect.Type) (reflect.Value, bool) {
	if typ == directFuncVoidType {
		return reflect.ValueOf(voidFunc{c}.call), true
	}
	if maker, ok := lookupDirectFunc(typ); ok {
		return maker(c), true
	}
	return reflect.Value{}, false
}

type directFuncEntry struct {
	maker DirectFuncMaker
	pc    uintptr
}

var (
	directFuncMakers = map[reflect.Type]directFuncEntry{}
	directFuncPCs    = map[uintptr]struct{}{}
)

// RegisterDirectFunc registers a maker for functions of type typ.
// Later registrations for the same type replace earlier ones.
// Passing a nil maker removes the registration.
// Do not unregister or replace a signature while function values of that
// type are still in use.
//
// The maker must return a method value of type typ whose receiver embeds
// DirectFuncVal as its first field. On gc, registration panics if the maker does
// not satisfy that contract. func() is built in and is not taken from the
// registry.
func RegisterDirectFunc(typ reflect.Type, maker DirectFuncMaker) {
	if typ == nil || typ.Kind() != reflect.Func {
		panic("ixgo: DirectFunc type must be a function")
	}
	var pc uintptr
	if maker != nil {
		pc = validateDirectFunc(typ, maker)
	}
	if old, ok := directFuncMakers[typ]; ok {
		delete(directFuncPCs, old.pc)
		delete(directFuncMakers, typ)
	}
	if maker == nil {
		return
	}
	if pc != 0 {
		directFuncPCs[pc] = struct{}{}
	}
	directFuncMakers[typ] = directFuncEntry{maker: maker, pc: pc}
}

// RegisterDirectFuncFor registers a DirectFunc for signature F.
//
//	type fn struct{ ixgo.DirectFuncVal }
//	func (c fn) Int(n int) int { return c.Call(n).(int) }
//
//	func init() {
//	    ixgo.RegisterDirectFuncFor(func(c ixgo.DirectFuncVal) func(int) int {
//	        return fn{c}.Int
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
// bind must return a method value whose receiver embeds DirectFuncVal as its first
// field. Passing a nil bind removes the registration.
func RegisterDirectFuncFor[F any](bind func(DirectFuncVal) F) {
	typ := reflect.TypeFor[F]()
	if bind == nil {
		RegisterDirectFunc(typ, nil)
		return
	}
	RegisterDirectFunc(typ, func(c DirectFuncVal) reflect.Value {
		return reflect.ValueOf(bind(c))
	})
}

func checkDirectFuncType(typ reflect.Type, maker DirectFuncMaker) {
	v := maker(DirectFuncVal{})
	if !v.IsValid() || v.Kind() != reflect.Func || v.Type() != typ {
		panic(fmt.Sprintf("ixgo: DirectFunc maker returned %v, want %v", v.Type(), typ))
	}
	if v.IsNil() {
		panic("ixgo: DirectFunc maker returned a nil function")
	}
}

func lookupDirectFunc(typ reflect.Type) (DirectFuncMaker, bool) {
	e, ok := directFuncMakers[typ]
	if !ok {
		return nil, false
	}
	return e.maker, true
}

func isDirectFuncPC(pc uintptr) bool {
	_, ok := directFuncPCs[pc]
	return ok
}
