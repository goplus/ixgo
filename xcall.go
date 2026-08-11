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
	"errors"
	"reflect"
	"unsafe"

	"github.com/goplus/reflectx"
)

func fieldAddrX(v interface{}, index int) (interface{}, error) {
	x := reflect.ValueOf(v).Elem()
	if !x.IsValid() {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	return reflectx.FieldX(x, index).Addr().Interface(), nil
}

func fieldAddrAt(v interface{}, offset uintptr, typ reflect.Type) (interface{}, error) {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() || rv.IsNil() {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	return reflect.NewAt(typ, unsafe.Add(unsafe.Pointer(rv.Pointer()), offset)).Interface(), nil
}

func fieldX(v interface{}, index int) (interface{}, error) {
	x := reflect.ValueOf(v)
	for x.Kind() == reflect.Ptr {
		x = x.Elem()
	}
	if !x.IsValid() {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	return reflectx.FieldX(x, index).Interface(), nil
}

func allMethodX(typ reflect.Type) []reflect.Method {
	n := reflectx.NumMethodX(typ)
	if n == 0 {
		return nil
	}
	ms := make([]reflect.Method, n)
	for i := 0; i < n; i++ {
		ms[i] = reflectx.MethodX(typ, i)
	}
	return ms
}
