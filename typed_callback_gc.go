//go:build !llgo
// +build !llgo

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
	"unsafe"

	"github.com/visualfc/funcval"
)

func inspectTypedCallback(typ reflect.Type, maker TypedCallbackMaker) uintptr {
	if !funcval.IsSupport {
		checkTypedCallbackType(typ, maker)
		return 0
	}
	return validateTypedCallback(typ, maker)
}

func validateTypedCallback(typ reflect.Type, maker TypedCallbackMaker) uintptr {
	sentinel := new(Interp)
	fn := new(function)
	env := []value{sentinel}
	v := maker(FuncVal{interp: sentinel, pfn: fn, typ: typ, env: env})
	if !v.IsValid() || v.Kind() != reflect.Func || v.Type() != typ {
		panic(fmt.Sprintf("ixgo: typed callback maker returned %v, want %v", v.Type(), typ))
	}
	if v.IsNil() {
		panic("ixgo: typed callback maker returned a nil function")
	}
	fv, n := funcval.Get(v.Interface())
	if n != 0 {
		panic("ixgo: typed callback maker must return a method value")
	}
	got := typedCallbackReceiver(fv)
	if got == nil || got.interp != sentinel || got.pfn != fn || got.typ != typ || len(got.env) != 1 || got.env[0] != sentinel {
		panic("ixgo: typed callback receiver must embed ixgo.FuncVal as its first field")
	}
	return fv.Fn
}

func typedCallbackReceiver(fv *funcval.FuncVal) *FuncVal {
	if fv == nil {
		return nil
	}
	// gc ABI: FuncVal is one word, followed by the receiver.
	// The call methods must retain value receivers.
	return &(*struct {
		funcval.FuncVal
		receiver FuncVal
	})(unsafe.Pointer(fv)).receiver
}
