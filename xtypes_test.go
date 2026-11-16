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
	"go/types"
	"testing"
)

// Unexported methods with the same spelling from different packages are
// distinct. extractMethodSet must keep both test/a.m and test/b.m
// (Go issue24693 / ixgo#515).
func TestExtractMethodSetDistinctUnexportedMethods(t *testing.T) {
	const pkgA = `package a

type T struct{}

func (T) m() { panic("called a.m") }

type I interface{ m() }
`
	const pkgB = `package b

import "test/a"

type T struct{ a.T }

func (T) m() {}

func F1(i interface{ m() }) { i.m() }

func F2(i interface {
	m()
	a.I
}) {
	i.m()
}
`
	const src = `package main

import "test/b"

func main() {
	b.F1(b.T{})
	b.F2(b.T{})
}
`
	ctx := NewContext(0)
	if err := ctx.AddImportFile("test/a", "a.go", pkgA); err != nil {
		t.Fatal(err)
	}
	if err := ctx.AddImportFile("test/b", "b.go", pkgB); err != nil {
		t.Fatal(err)
	}
	pkg, err := ctx.LoadFile("main.go", src)
	if err != nil {
		t.Fatal(err)
	}
	interp, err := NewInterp(ctx, pkg)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()

	bpkg := pkg.Prog.ImportedPackage("test/b")
	if bpkg == nil {
		t.Fatal("missing package test/b")
	}
	named := bpkg.Type("T")
	if named == nil {
		t.Fatal("missing type test/b.T")
	}
	methods, _, _ := interp.record.extractMethodSet(named.Type())
	got := make(map[string]int, len(methods))
	for _, m := range methods {
		got[m.Obj().Id()]++
	}
	if len(methods) != 2 || got["test/a.m"] != 1 || got["test/b.m"] != 1 {
		t.Fatalf("extractMethodSet identities = %v (len=%d), want test/a.m and test/b.m once each", got, len(methods))
	}

	if err := interp.RunInit(); err != nil {
		t.Fatal(err)
	}
	if _, err := interp.RunMain(); err != nil {
		t.Fatal(err)
	}
}

func TestSetMethodsDeclarationSignature(t *testing.T) {
	const src = `package main
import "reflect"

type Base struct { N int }
func (b Base) Value(n int) int { return b.N + n }
func (b *Base) Add(n int) int { b.N += n; return b.N }
func (b Base) Sum(ns ...int) int {
 result := b.N
 for _, n := range ns { result += n }
 return result
}
type Outer struct { Base }
type PointerOuter struct { *Base }
type Nested struct { Outer }
type Box[T any] struct { V T }
func (b Box[T]) Get() T { return b.V }
func (b *Box[T]) Set(v T) { b.V = v }
type IntOuter struct { Box[int] }
type StringOuter struct { *Box[string] }

func check(v bool) { if !v { panic("method signature regression") } }
func value(v interface { Value(int) int }, n int) int { return v.Value(n) }
func add(v interface { Add(int) int }, n int) int { return v.Add(n) }
func main() {
 b := Base{N: 1}
 check(value(b, 2) == 3)
 check(add(&b, 3) == 4)
 o := Outer{Base{N: 10}}
 check(value(o, 2) == 12)
 check(add(&o, 3) == 13 && o.N == 13)
 p := PointerOuter{&Base{N: 20}}
 check(value(p, 2) == 22)
 check(add(p, 3) == 23 && p.N == 23)
 n := Nested{Outer{Base{N: 30}}}
 check(value(n, 2) == 32)
 check(add(&n, 3) == 33 && n.N == 33)
 for index, v := range []interface{}{b, o, p, n} {
  rv := reflect.ValueOf(v)
  want := []int64{4, 13, 23, 33}[index]
  check(rv.MethodByName("Value").Call([]reflect.Value{reflect.ValueOf(2)})[0].Int() == want + 2)
  check(rv.MethodByName("Sum").CallSlice([]reflect.Value{reflect.ValueOf([]int{2, 3})})[0].Int() == want + 5)
 }
 check(reflect.ValueOf(&o).MethodByName("Add").Call([]reflect.Value{reflect.ValueOf(1)})[0].Int() == 14)
 check(o.N == 14)
 check(reflect.ValueOf(o).MethodByName("Add").IsValid() == false)
 check(reflect.ValueOf(&o).MethodByName("Add").IsValid())
 check(reflect.ValueOf(p).MethodByName("Add").IsValid())
 i := IntOuter{Box[int]{V: 42}}
 var ig interface { Get() int } = i
 check(ig.Get() == 42)
 var is interface { Set(int) } = &i
 is.Set(43)
 check(i.Get() == 43)
 check(reflect.ValueOf(i).MethodByName("Get").Call(nil)[0].Int() == 43)
 s := StringOuter{&Box[string]{V: "first"}}
 var sg interface { Get() string } = s
 check(sg.Get() == "first")
 var ss interface { Set(string) } = s
 ss.Set("second")
 check(s.Get() == "second")
 check(reflect.ValueOf(s).MethodByName("Get").Call(nil)[0].String() == "second")
}
`
	for _, mode := range []Mode{0, SupportMultipleInterp, SupportMultipleInterp | CheckGopOverloadFunc} {
		t.Run(fmt.Sprintf("mode_%d", mode), func(t *testing.T) {
			ctx := NewContext(mode)
			pkg, err := ctx.LoadFile("main.go", src)
			if err != nil {
				t.Fatal(err)
			}
			for n := 0; n < 2; n++ {
				interp, err := NewInterp(ctx, pkg)
				if err != nil {
					t.Fatal(err)
				}
				func() {
					defer interp.UnsafeRelease()
					for _, name := range []string{"Base", "Outer", "PointerOuter", "Nested", "IntOuter", "StringOuter"} {
						typ := pkg.Type(name).Type()
						for _, recv := range []types.Type{typ, types.NewPointer(typ)} {
							mset := pkg.Prog.MethodSets.MethodSet(recv)
							for j := 0; j < mset.Len(); j++ {
								sel := mset.At(j)
								selectedSig := sel.Type().(*types.Signature)
								declaredSig := sel.Obj().Type().(*types.Signature)
								selectedType, selectedNested := interp.record.ToType(selectedSig)
								declaredType, declaredNested := interp.record.ToType(declaredSig)
								if selectedType != declaredType || selectedNested != declaredNested {
									t.Fatalf("%s: selected signature converts to (%v, %v), declaration to (%v, %v)", sel, selectedType, selectedNested, declaredType, declaredNested)
								}
								if isPointer(selectedSig.Recv().Type()) != isPointer(sel.Recv()) {
									t.Fatalf("%s: selected receiver pointer flag differs from signature", sel)
								}
							}
						}
					}
					if err := interp.RunInit(); err != nil {
						t.Fatal(err)
					}
					if _, err := interp.RunMain(); err != nil {
						t.Fatal(err)
					}
				}()
			}
		})
	}
}
