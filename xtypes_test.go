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
