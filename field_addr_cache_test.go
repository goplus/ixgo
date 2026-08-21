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

var fieldAddrModes = [...]struct {
	name string
	mode ixgo.Mode
}{
	{name: "uncached"},
	{name: "cached", mode: ixgo.EnableCachedReg},
}

func TestFieldAddrCache(t *testing.T) {
	const source = `package main

type record struct { value int }

func update(left, right *record) {
	for i := 0; i < 4; i++ {
		current := left
		if i >= 2 { current = right }
		*(&current.value) += i + 1
	}
}

func main() {
	left, right := &record{}, &record{}
	for i := 0; i < 8; i++ { update(left, right) }
	if left.value != 24 || right.value != 56 { panic("wrong cached address") }
}
`
	for _, mode := range fieldAddrModes {
		t.Run(mode.name, func(t *testing.T) {
			ctx := ixgo.NewContext(mode.mode)
			ctx.SetLeastCallForEnablePool(0)
			if _, err := ctx.RunFile("main.go", source, nil); err != nil {
				t.Fatal(err)
			}
		})
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
	for _, mode := range fieldAddrModes {
		t.Run(mode.name, func(t *testing.T) {
			_, err := ixgo.NewContext(mode.mode).RunFile("main.go", source, nil)
			if err == nil || !strings.Contains(err.Error(), "invalid memory address or nil pointer dereference") {
				t.Fatalf("got error %v, want nil pointer dereference", err)
			}
		})
	}
}

func TestNamedPointerFieldAddr(t *testing.T) {
	const source = `package main

type B *struct{ A }
type A interface{ m(B) }

type s struct{}

func (s) m(B) {}

func main() {
	var b B = new(struct{ A })
	b.A = s{}
	(*b).m(b)
}
`
	for _, mode := range fieldAddrModes {
		t.Run(mode.name, func(t *testing.T) {
			if _, err := ixgo.NewContext(mode.mode).RunFile("main.go", source, nil); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestFieldAddrCacheReleasesReceiver(t *testing.T) {
	if ixgo.IsLLGo {
		t.Skip("skip llgo")
	}
	const source = `package main

import "runtime"

type record struct { value [8]int64 }

var finalized = make(chan struct{}, 1)
var sink *[8]int64

func main() {
	receiver := new(record)
	runtime.SetFinalizer(receiver, func(*record) { finalized <- struct{}{} })
	sink = &receiver.value
	sink = nil
	for i := 0; i < 5; i++ { runtime.GC() }
	select {
	case <-finalized:
	default:
		panic("cached receiver kept object alive")
	}
}
`
	mode := ixgo.ExperimentalSupportGC | ixgo.EnableCachedReg
	if _, err := ixgo.NewContext(mode).RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkFieldAddrCache(b *testing.B) {
	const source = `package main

type record struct { value int }

var items [128]record

func Receivers(vary bool) []*record {
	receivers := make([]*record, len(items))
	for i := range receivers {
		index := 0
		if vary { index = i }
		receivers[i] = &items[index]
	}
	return receivers
}

func FieldAddr(receivers []*record) *int {
	var addr *int
	for _, receiver := range receivers {
		addr = &receiver.value
	}
	return addr
}

func main() {}
`
	for _, mode := range fieldAddrModes {
		for _, test := range []struct {
			name string
			vary bool
		}{
			{name: "hit"},
			{name: "miss", vary: true},
		} {
			b.Run(mode.name+"/"+test.name, func(b *testing.B) {
				ctx := ixgo.NewContext(mode.mode)
				ctx.SetLeastCallForEnablePool(0)
				interp, err := ctx.LoadInterp("main.go", source)
				if err != nil {
					b.Fatal(err)
				}
				b.Cleanup(interp.UnsafeRelease)
				receivers, err := interp.RunFunc("Receivers", test.vary)
				if err != nil {
					b.Fatal(err)
				}
				b.ReportAllocs()
				b.ResetTimer()
				for b.Loop() {
					if _, err := interp.RunFunc("FieldAddr", receivers); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}
