/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package memo_test

import (
	"testing"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/internal/memo"
)

func TestEligibility(t *testing.T) {
	pkg, err := ixgo.NewContext(0).LoadFile("check.go", `package main
var count int
func scalar(n int) int {if n<2{return n};return scalar(n-1)+scalar(n-2)}
func effect(n int) int {count++;if n<2{return n};return effect(n-1)}
func acc(x int,a,b,p *int) {if x<2 {*p+=x;return};other(x-1,a,b,p);other(x-2,a,b,p)}
func other(x int,a,b,p *int) {acc(x,a,b,p)}
func bad(x int,a,b,p *int) {*a+=1;acc(x,a,b,p)}
func floating(x float64,a,b,p *float64) {*p+=x}
`)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name        string
		scalar, acc bool
	}{
		{"scalar", true, false}, {"effect", false, false}, {"acc", false, true}, {"other", false, true}, {"bad", false, false}, {"floating", false, false},
	} {
		fn := pkg.Func(tc.name)
		if memo.ScalarEligible(fn) != tc.scalar || memo.AccumulatorEligible(fn) != tc.acc {
			t.Fatal(tc.name)
		}
	}
	if memo.ScalarEligible(nil) || memo.AccumulatorEligible(nil) {
		t.Fatal("nil function accepted")
	}
}

func TestStateEligibility(t *testing.T) {
	pkg, err := ixgo.NewContext(0).LoadFile("state.go", `package main
type Amount float64
func defined(x Amount,p *Amount) {*p *= x}
var global float64
func arithmetic(n int,x float64,p *float64) {if n==0 {*p=(*p-x)*x/x;return};arithmetic(n-1,x,p)}
func bits(x uint64,p *uint64) {*p=(*p^x)>>2}
func multiple(x int,a,b *int) {*a *= x;*b *= x}
func external(x float64,p *float64) {*p += global}
func convert(x int,p *float64) {*p *= float64(x)}
`)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"arithmetic", "bits", "convert", "defined"} {
		if memo.AnalyzeState(pkg.Func(name)) == nil {
			t.Fatal("state rejected", name)
		}
	}
	for _, name := range []string{"multiple", "external"} {
		if memo.AnalyzeState(pkg.Func(name)) != nil {
			t.Fatal("unsafe state accepted", name)
		}
	}
}
