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
