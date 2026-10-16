/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package ixgo

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/goplus/ixgo/internal/memo"
)

func TestInterpreterNumericState(t *testing.T) {
	for _, kind := range []string{"int64", "float32", "float64"} {
		operations := []string{"*p -= x", "*p *= x", "*p /= x", "*p = -*p + x", "if *p > x { *p -= x } else { *p *= x }"}
		if kind == "int64" {
			operations = append(operations, "*p %= x", "*p ^= x", "*p &= x", "*p |= x", "*p &^= x", "*p <<= uint(x)", "*p >>= uint(x)")
		}
		for _, op := range operations {
			t.Run(kind+"/"+op, func(t *testing.T) {
				src := fmt.Sprintf(`package main
func f(n int,x %[1]s,p *%[1]s) {if n==0 {%[2]s;return};*p=7;f(n-1,x,p);*p=7;f(n-1,x,p)}
`, kind, op)
				i := loadMainMemo(t, []byte(src))
				p := i.funcs[i.mainpkg.Func("f")]
				if !p.interpreterAccumulator || memo.AnalyzeState(p.Fn) == nil {
					t.Fatal("state cache missing")
				}
				typ := reflect.TypeOf(int64(0))
				if kind == "float32" {
					typ = reflect.TypeOf(float32(0))
				}
				if kind == "float64" {
					typ = reflect.TypeOf(float64(0))
				}
				x := reflect.New(typ).Elem()
				if kind == "int64" {
					x.SetInt(3)
				} else {
					x.SetFloat(3)
				}
				var want uint64
				for _, enabled := range []bool{false, true} {
					p.interpreterAccumulator = enabled
					out := reflect.New(typ)
					memo.SetScalarBits(out.Elem(), memo.ValueBits(x))
					if _, err := i.RunFunc("f", 8, x.Interface(), out.Interface()); err != nil {
						t.Fatal(err)
					}
					got := memo.ValueBits(out.Elem())
					if !enabled {
						want = got
					} else if got != want {
						t.Fatalf("bits %x != %x", got, want)
					}
				}
			})
		}
	}
}

func TestInterpreterStateFloatSpecials(t *testing.T) {
	i := loadMainMemo(t, []byte(`package main
func leaf(x float64,p *float64) {*p /= x}
func root(x float64,p *float64) {leaf(x,p);*p = -0.0;leaf(x,p)}
`))
	for _, x := range []float64{0, math.Copysign(0, -1), math.Inf(1), math.Inf(-1), math.Float64frombits(0x7ff8000000000001)} {
		var want uint64
		for _, enabled := range []bool{false, true} {
			for _, name := range []string{"root", "leaf"} {
				i.funcs[i.mainpkg.Func(name)].interpreterAccumulator = enabled
			}
			out := math.Copysign(0, -1)
			if _, err := i.RunFunc("root", x, &out); err != nil {
				t.Fatal(err)
			}
			got := math.Float64bits(out)
			if !enabled {
				want = got
			} else if got != want {
				t.Fatalf("special %x: %x != %x", math.Float64bits(x), got, want)
			}
		}
	}
}

func TestInterpreterStatePanic(t *testing.T) {
	for _, op := range []string{"*p /= x", "*p %= x", "*p <<= x"} {
		i := loadMainMemo(t, []byte(fmt.Sprintf(`package main
func f(x int,p *int) {%s}
`, op)))
		if !i.funcs[i.mainpkg.Func("f")].interpreterAccumulator {
			t.Fatal(op)
		}
		x := 0
		if op == "*p <<= x" {
			x = -1
		}
		for j := 0; j < 2; j++ {
			out := 7
			if _, err := i.RunFunc("f", x, &out); err == nil {
				t.Fatal("panic missing", op)
			}
		}
	}
}

func BenchmarkInterpreterState(b *testing.B) {
	for _, kind := range []string{"int64", "float64"} {
		i := loadMainMemo(b, []byte(fmt.Sprintf(`package main
func f(n int,x %[1]s,p *%[1]s) {if n==0 {*p *= x;return};*p=7;f(n-1,x,p);*p=7;f(n-1,x,p)}
`, kind)))
		p := i.funcs[i.mainpkg.Func("f")]
		for _, enabled := range []bool{false, true} {
			name := "baseline"
			if enabled {
				name = "state"
			}
			b.Run(kind+"/"+name, func(b *testing.B) {
				p.interpreterAccumulator = enabled
				var x Value = int64(3)
				var out Value = new(int64)
				if kind == "float64" {
					x = float64(3)
					out = new(float64)
				}
				b.ReportAllocs()
				for b.Loop() {
					if _, err := i.RunFunc("f", 16, x, out); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}

func TestInterpreterStateAliasing(t *testing.T) {
	i := loadMainMemo(t, []byte(`package main
func f(n int,a,p *int) {if n==0 {*p *= *a;return};f(n-1,a,p);f(n-1,a,p)}
`))
	p := i.funcs[i.mainpkg.Func("f")]
	if !p.interpreterAccumulator {
		t.Fatal("state cache missing")
	}
	var want int
	for _, enabled := range []bool{false, true} {
		p.interpreterAccumulator = enabled
		out := 2
		if _, err := i.RunFunc("f", 2, &out, &out); err != nil {
			t.Fatal(err)
		}
		if !enabled {
			want = out
		} else if out != want {
			t.Fatal(out, want)
		}
	}
}
