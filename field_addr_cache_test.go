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

package ixgo_test

import (
	"strings"
	"testing"

	"github.com/goplus/ixgo"
)

func TestFieldAddrCache(t *testing.T) {
	const source = `package main

type record struct { value int }

func stable(r *record) *int {
	var first *int
	for i := 0; i < 4; i++ {
		addr := &r.value
		if first == nil {
			first = addr
		} else if addr != first {
			panic("cached address changed")
		}
	}
	return first
}

func main() {
	left, right := &record{}, &record{}
	for i := 0; i < 64; i++ { // also exercises pooled frames and receiver changes
		current := left
		if i&1 != 0 { current = right }
		*stable(current) = i + 1
	}
	if left.value != 63 || right.value != 64 { panic("wrong cached address") }
}
`
	ctx := ixgo.NewContext(0)
	ctx.SetLeastCallForEnablePool(0)
	if _, err := ctx.RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
}

func TestFieldAddrCacheNilReceiver(t *testing.T) {
	const source = `package main

type record struct { value int }

func main() {
	r := &record{}
	for i := 0; i < 2; i++ {
		current := r
		if i != 0 { current = nil }
		_ = &current.value
	}
}
`
	_, err := ixgo.NewContext(0).RunFile("main.go", source, nil)
	if err == nil || !strings.Contains(err.Error(), "invalid memory address or nil pointer dereference") {
		t.Fatalf("got error %v, want nil pointer dereference", err)
	}
}

func BenchmarkFieldAddrCache(b *testing.B) {
	const source = `package main

type record struct { value int }

var item record

func FieldAddr(iterations int) *int {
	r := &item
	var addr *int
	for i := 0; i < iterations; i++ {
		addr = &r.value
	}
	return addr
}

func main() {}
`
	ctx := ixgo.NewContext(0)
	ctx.SetLeastCallForEnablePool(0)
	interp, err := ctx.LoadInterp("main.go", source)
	if err != nil {
		b.Fatal(err)
	}
	b.Cleanup(interp.UnsafeRelease)

	b.ReportAllocs()
	for b.Loop() {
		if _, err := interp.RunFunc("FieldAddr", 128); err != nil {
			b.Fatal(err)
		}
	}
}
