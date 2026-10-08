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

// Package typed registers typed callbacks for func() T with common result types.
//
//	import _ "github.com/goplus/ixgo/typed"
//
// func() is built into ixgo. Signatures with parameters are not registered
// here. Register those from the host with ixgo.RegisterTypedCallbackFunc.
package typed

import "github.com/goplus/ixgo"

type result[T any] struct {
	ixgo.FuncVal
}

func (c result[T]) get() T {
	return c.Call().(T)
}

type resultError struct {
	ixgo.FuncVal
}

func (c resultError) get() error {
	v := c.Call()
	if v == nil {
		return nil
	}
	return v.(error)
}

func registerResult[T any]() {
	ixgo.RegisterTypedCallbackFunc(func(c ixgo.FuncVal) func() T {
		return result[T]{c}.get
	})
}

func init() {
	registerResult[bool]()
	registerResult[int]()
	registerResult[int8]()
	registerResult[int16]()
	registerResult[int32]()
	registerResult[int64]()
	registerResult[uint]()
	registerResult[uint8]()
	registerResult[uint16]()
	registerResult[uint32]()
	registerResult[uint64]()
	registerResult[uintptr]()
	registerResult[float32]()
	registerResult[float64]()
	registerResult[string]()
	ixgo.RegisterTypedCallbackFunc(func(c ixgo.FuncVal) func() error {
		return resultError{c}.get
	})
}
