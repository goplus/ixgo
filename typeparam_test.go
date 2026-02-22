//go:build go1.18
// +build go1.18

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

package ixgo_test

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/go/token"
	_ "github.com/goplus/ixgo/pkg/path/filepath"
	_ "github.com/goplus/ixgo/pkg/reflect"
	_ "github.com/goplus/ixgo/pkg/sync/atomic"
)

func TestTypeParamNamed(t *testing.T) {
	src := `package main

import (
	"path/filepath"
	"reflect"
)

type N int
type _Nil[a any, b any] struct{}

func main() {
	var v1 _Nil[int, N]
	var v2 _Nil[filepath.WalkFunc, _Nil[int, N]]
	if s := reflect.TypeOf(v1).String(); s != "main._Nil[int,main.N]" {
		panic(s)
	}
	if s := reflect.TypeOf(v2).String(); s != "main._Nil[path/filepath.WalkFunc,main._Nil[int,main.N]]" {
		panic(s)
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNestedTypeParameterized(t *testing.T) {
	src := `// run

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This test case stress tests a number of subtle cases involving
// nested type-parameterized declarations. At a high-level, it
// declares a generic function that contains a generic type
// declaration:
//
//	func F[A intish]() {
//		type T[B intish] struct{}
//
//		// store reflect.Type tuple (A, B, F[A].T[B]) in tests
//	}
//
// It then instantiates this function with a variety of type arguments
// for A and B. Particularly tricky things like shadowed types.
//
// From this data it tests two things:
//
// 1. Given tuples (A, B, F[A].T[B]) and (A', B', F[A'].T[B']),
//    F[A].T[B] should be identical to F[A'].T[B'] iff (A, B) is
//    identical to (A', B').
//
// 2. A few of the instantiations are constructed to be identical, and
//    it tests that exactly these pairs are duplicated (by golden
//    output comparison to nested.out).
//
// In both cases, we're effectively using the compiler's existing
// runtime.Type handling (which is well tested) of type identity of A
// and B as a way to help bootstrap testing and validate its new
// runtime.Type handling of F[A].T[B].
//
// This isn't perfect, but it smoked out a handful of issues in
// gotypes2 and unified IR.

package main

import (
	"fmt"
	"reflect"
)

type test struct {
	TArgs    [2]reflect.Type
	Instance reflect.Type
}

var tests []test

type intish interface{ ~int }

type Int int
type GlobalInt = Int // allow access to global Int, even when shadowed

func F[A intish]() {
	add := func(B, T interface{}) {
		tests = append(tests, test{
			TArgs: [2]reflect.Type{
				reflect.TypeOf(A(0)),
				reflect.TypeOf(B),
			},
			Instance: reflect.TypeOf(T),
		})
	}

	type Int int

	type T[B intish] struct{}

	add(int(0), T[int]{})
	add(Int(0), T[Int]{})
	add(GlobalInt(0), T[GlobalInt]{})
	add(A(0), T[A]{}) // NOTE: intentionally dups with int and GlobalInt

	type U[_ any] int
	type V U[int]
	type W V

	add(U[int](0), T[U[int]]{})
	add(U[Int](0), T[U[Int]]{})
	add(U[GlobalInt](0), T[U[GlobalInt]]{})
	add(U[A](0), T[U[A]]{}) // NOTE: intentionally dups with U[int] and U[GlobalInt]
	add(V(0), T[V]{})
	add(W(0), T[W]{})
}

func main() {
	type Int int

	F[int]()
	F[Int]()
	F[GlobalInt]()

	type U[_ any] int
	type V U[int]
	type W V

	F[U[int]]()
	F[U[Int]]()
	F[U[GlobalInt]]()
	F[V]()
	F[W]()

	type X[A any] U[X[A]]

	F[X[int]]()
	F[X[Int]]()
	F[X[GlobalInt]]()

	for j, tj := range tests {
		for i, ti := range tests[:j+1] {
			if (ti.TArgs == tj.TArgs) != (ti.Instance == tj.Instance) {
				fmt.Printf("FAIL: %d,%d: %s, but %s\n", i, j, eq(ti.TArgs, tj.TArgs), eq(ti.Instance, tj.Instance))
			}

			// The test is constructed so we should see a few identical types.
			// See "NOTE" comments above.
			if i != j && ti.Instance == tj.Instance {
				fmt.Printf("%d,%d: %v\n", i, j, ti.Instance)
			}
		}
	}
}

func eq(a, b interface{}) string {
	op := "=="
	if a != b {
		op = "!="
	}
	return fmt.Sprintf("%v %s %v", a, op, b)
}
`
	out := `0,3: main.T[int;int]
4,7: main.T[int;main.U[int;int]]
22,23: main.T[main.Int;main.Int]
26,27: main.T[main.Int;main.U[main.Int;main.Int]]
`
	ctx := ixgo.NewContext(0)
	var buf bytes.Buffer
	ctx.RegisterExternal("fmt.Printf", func(format string, a ...interface{}) (n int, err error) {
		fmt.Fprintf(&buf, format, a...)
		return fmt.Printf(format, a...)
	})
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() != out {
		t.Fatal("error output", buf.String())
	}
}

func TestNestedTypeParams(t *testing.T) {
	src := `package main

func T[N ~int](v N) {
	type T []N
	var t T
	println(t)
}

func main() {
	T(100)
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTypeParamsRecursive(t *testing.T) {
	src := `package main

import "fmt"

func main() {
	recursive()
}	

func recursive() {
	type T int
	if got, want := recur1[T](5), T(110); got != want {
		panic(fmt.Sprintf("recur1[T](5) = %d, want = %d", got, want))
	}
}

type Integer interface {
	~int | ~int32 | ~int64
}

func recur1[T Integer](n T) T {
	if n == 0 || n == 1 {
		return T(1)
	} else {
		return n * recur2(n-1)
	}
}

func recur2[T Integer](n T) T {
	list := make([]T, n)
	for i, _ := range list {
		list[i] = T(i + 1)
	}
	var sum T
	for _, elt := range list {
		sum += elt
	}
	return sum + recur1(n-1)
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAtomicPointer(t *testing.T) {
	if runtime.Version()[:6] == "go1.18" {
		t.Skip("skip go1.18")
	}
	src := `package main

import (
	"sync/atomic"
)

func main() {
	var i atomic.Int64
	i.Store(200)
	if i.Load() != 200 {
		panic("error atomic.Int64")
	}
	var v atomic.Pointer[int]
	var n int = 200
	v.Store(&n)
	if p := v.Load(); *p != 200 {
		panic("error atomic.Pointer[int]")
	}
	var v2 atomic.Pointer[string]
	var n2 string = "hello"
	v2.Store(&n2)
	if p := v2.Load(); *p != "hello" {
		panic("error atomic.Pointer[string]")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAtomicTypeArgs(t *testing.T) {
	// type FileSet struct {
	// 	mutex sync.RWMutex         // protects the file set
	// 	base  int                  // base offset for the next file
	// 	files []*File              // list of files in the order added to the set
	// 	last  atomic.Pointer[File] // cache of last file looked up
	// }
	src := `package main

import (
	"go/token"
)

type CFG struct {
}

func (g *CFG) Format(fset *token.FileSet) {
}

func main() {
}`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}
