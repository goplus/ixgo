/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package ixgo

import (
	"math"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func loadMainMemo(t testing.TB, source []byte) *Interp {
	t.Helper()
	i, err := NewContext(0).LoadInterp("memo.go", source)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(i.UnsafeRelease)
	return i
}

func TestInterpreterMemo(t *testing.T) {
	for _, name := range []string{"fibish", "fibish_closure", "fibish_float32", "fibish_float64"} {
		t.Run(name, func(t *testing.T) {
			src, err := os.ReadFile("testdata/fibish/" + name + ".go")
			if err != nil {
				t.Fatal(err)
			}
			i := loadMainMemo(t, src)
			p := i.funcs[i.mainpkg.Func("f")]
			if !p.interpreterMemo {
				t.Fatal("memo not enabled")
			}
			if name == "fibish_closure" { // Exercise the unchanged forwarded function argument via main.
				if _, err := i.RunFunc("main"); err != nil {
					t.Fatal(err)
				}
				return
			}
			for n := 0; n <= 20; n++ {
				var x Value = n
				if name == "fibish_float32" {
					x = float32(n)
				}
				if name == "fibish_float64" {
					x = float64(n)
				}
				p.interpreterMemo = true
				a, err := i.RunFunc("f", x)
				if err != nil {
					t.Fatal(err)
				}
				p.interpreterMemo = false
				b, err := i.RunFunc("f", x)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(a, b) {
					t.Fatal(n, a, b)
				}
			}
		})
	}
}

func TestInterpreterMemoSafety(t *testing.T) {
	i := loadMainMemo(t, []byte(`package main
var count int
func effects(n int) int {count++;if n<2{return n};return effects(n-1)+effects(n-2)}
func zero(x float64) float64 {if x==0{return x};z:=x-x;return 1/zero(z)+1/zero(-z)}
func divide(n int) int {if n<1{return 1/n};return divide(n-1)+divide(n-1)}
func get() int {return count}
`))
	if i.funcs[i.mainpkg.Func("effects")].interpreterMemo {
		t.Fatal("effects memoized")
	}
	a, err := i.RunFunc("zero", float64(1))
	if err != nil || !math.IsNaN(a.(float64)) {
		t.Fatal(a, err)
	}
	for j := 0; j < 2; j++ {
		if _, err := i.RunFunc("divide", 2); err == nil {
			t.Fatal("panic missing")
		}
	}
	if _, err := i.RunFunc("effects", 5); err != nil {
		t.Fatal(err)
	}
	a, err = i.RunFunc("get")
	if err != nil || a != 15 {
		t.Fatal(a, err)
	}
}

func BenchmarkInterpreterMemo(b *testing.B) {
	for _, name := range []string{"fibish", "fibish_float32", "fibish_float64"} {
		src, err := os.ReadFile("testdata/fibish/" + name + ".go")
		if err != nil {
			b.Fatal(err)
		}
		i := loadMainMemo(b, src)
		p := i.funcs[i.mainpkg.Func("f")]
		for _, enabled := range []bool{false, true} {
			label := "baseline"
			if enabled {
				label = "memo"
			}
			b.Run(name+"/"+label, func(b *testing.B) {
				p.interpreterMemo = enabled
				var x Value = 20
				if name == "fibish_float32" {
					x = float32(20)
				}
				if name == "fibish_float64" {
					x = float64(20)
				}
				b.ReportAllocs()
				for b.Loop() {
					if _, err := i.RunFunc("f", x); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}

func TestInterpreterMemoHooksAndTracing(t *testing.T) {
	for _, tracing := range []bool{false, true} {
		mode := Mode(0)
		if tracing {
			mode = EnableTracing
		}
		ctx := NewContext(mode)
		calls := 0
		if !tracing {
			ctx.SetFuncHook(func(*FuncInfo) { calls++ }, nil)
		}
		i, err := ctx.LoadInterp("hook.go", `package main
func f(n int) int {if n<2{return n};return f(n-1)+f(n-2)}
`)
		if err != nil {
			t.Fatal(err)
		}
		if i.funcs[i.mainpkg.Func("f")].interpreterMemo {
			t.Fatal("instrumented function memoized")
		}
		if !tracing {
			v, err := i.RunFunc("f", 5)
			if err != nil || v != 5 || calls != 15 {
				t.Fatal(v, err, calls)
			}
		}
		i.UnsafeRelease()
	}
}

func TestInterpreterMemoConcurrent(t *testing.T) {
	var interps []*Interp
	for j := 0; j < 4; j++ {
		interps = append(interps, loadMainMemo(t, []byte(`package main
func f(n int) int {if n<2{return n};return f(n-1)+f(n-2)}
`)))
	}
	var wg sync.WaitGroup
	for _, i := range interps {
		wg.Add(1)
		go func(i *Interp) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				v, err := i.RunFunc("f", 20)
				if err != nil || v != 6765 {
					t.Errorf("%v %v", v, err)
				}
			}
		}(i)
	}
	wg.Wait()
}

func TestInterpreterMemoInput40(t *testing.T) {
	for _, name := range []string{"fibish", "fibish_float32", "fibish_float64"} {
		src, err := os.ReadFile("testdata/fibish/" + name + ".go")
		if err != nil {
			t.Fatal(err)
		}
		i := loadMainMemo(t, src)
		var x Value = 40
		var want Value = Tuple{39088169, 126491972}
		if name == "fibish_float32" {
			x = float32(40)
			want = Tuple{float32(39088168), float32(126491968)}
		}
		if name == "fibish_float64" {
			x = float64(40)
			want = Tuple{float64(39088169), float64(126491972)}
		}
		start := time.Now()
		v, err := i.RunFunc("f", x)
		elapsed := time.Since(start)
		if err != nil || !reflect.DeepEqual(v, want) {
			t.Fatal(name, v, err)
		}
		t.Logf("%s ordinary interpreter f(40): %s (load excluded)", name, elapsed)
	}
}

func TestInterpreterMemoDebugFallback(t *testing.T) {
	ctx := NewContext(0)
	callbacks := 0
	ctx.SetDebug(func(*DebugInfo) { callbacks++ })
	i, err := ctx.LoadInterp("debug.go", `package main
func f(n int) int {if n<2{return n};return f(n-1)+f(n-2)}
func acc(x int,a,b,p *int) {if x<2 {*p+=x;return};acc(x-1,a,b,p);acc(x-2,a,b,p)}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	for _, name := range []string{"f", "acc"} {
		p := i.funcs[i.mainpkg.Func(name)]
		if p.interpreterMemo || p.interpreterAccumulator {
			t.Fatal("debug function memoized", name)
		}
	}
	if v, err := i.RunFunc("f", 5); err != nil || v != 5 {
		t.Fatal(v, err)
	}
	if callbacks == 0 {
		t.Fatal("debug callback not called")
	}
}

func loadUglyMainMemo(t testing.TB) *Interp {
	src, err := os.ReadFile("testdata/fibish/uglyfib.go")
	if err != nil {
		t.Fatal(err)
	}
	i := loadMainMemo(t, src)
	for _, name := range []string{"f", "g", "h", "k"} {
		if !i.funcs[i.mainpkg.Func(name)].interpreterAccumulator {
			t.Fatalf("%s not optimized", name)
		}
	}
	return i
}
func setMainAccumulator(i *Interp, enabled bool) {
	for _, name := range []string{"f", "g", "h", "k"} {
		i.funcs[i.mainpkg.Func(name)].interpreterAccumulator = enabled
	}
}

func TestInterpreterAccumulator(t *testing.T) {
	i := loadUglyMainMemo(t)
	for _, name := range []string{"f", "g", "h", "k"} {
		for n := 0; n <= 16; n++ {
			for alias := 0; alias < 3; alias++ {
				if alias != 0 && n > 3 {
					continue
				}
				var got [2][3]int
				for pass, enabled := range []bool{false, true} {
					setMainAccumulator(i, enabled)
					a, b, out := n-1, n-2, int(^uint(0)>>1)-3
					p := &out
					if alias == 1 {
						a = -1
						p = &a
					}
					if alias == 2 {
						b = -1
						p = &b
					}
					if _, err := i.RunFunc(name, n, &a, &b, p); err != nil {
						t.Fatal(err)
					}
					got[pass] = [3]int{a, b, out}
				}
				if got[0] != got[1] {
					t.Fatal(name, n, alias, got)
				}
			}
		}
		setMainAccumulator(i, true)
		out := 7
		if _, err := i.RunFunc(name, 1, (*int)(nil), (*int)(nil), &out); err != nil || out != 8 {
			t.Fatal(name, out, err)
		}
		for repeat := 0; repeat < 2; repeat++ {
			if _, err := i.RunFunc(name, 3, (*int)(nil), new(int), new(int)); err == nil {
				t.Fatal("nil panic missing")
			}
		}
	}
}

func TestInterpreterAccumulatorOverlap(t *testing.T) {
	i := loadMainMemo(t, []byte(`package main
func leaf(x int64,a,b,p *int64) { *p+=1;if *a==0 {*p+=2} }
func root(x int64,a,b,p *int64) {leaf(x,a,b,p);leaf(x,a,b,p)}
`))
	for _, name := range []string{"root", "leaf"} {
		if !i.funcs[i.mainpkg.Func(name)].interpreterAccumulator {
			t.Fatal("memo missing", name)
		}
	}
	run := func(enabled bool) int64 {
		for _, name := range []string{"root", "leaf"} {
			i.funcs[i.mainpkg.Func(name)].interpreterAccumulator = enabled
		}
		storage := new([3]int64)
		p := &storage[0]
		a := (*int64)(unsafe.Add(unsafe.Pointer(p), 1))
		*p = 252
		if _, err := i.RunFunc("root", int64(0), a, &storage[2], p); err != nil {
			t.Fatal(err)
		}
		return *p
	}
	if a, b := run(false), run(true); a != b {
		t.Fatal(a, b)
	}
}

func TestInterpreterAccumulatorConcurrent(t *testing.T) {
	var interps []*Interp
	for j := 0; j < 4; j++ {
		interps = append(interps, loadUglyMainMemo(t))
	}
	var wg sync.WaitGroup
	for _, i := range interps {
		wg.Add(1)
		go func(i *Interp) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				a, b, out := 19, 18, 5
				if _, err := i.RunFunc("f", 20, &a, &b, &out); err != nil || out != 6770 {
					t.Errorf("%v %v", out, err)
				}
			}
		}(i)
	}
	wg.Wait()
}

func BenchmarkInterpreterAccumulator(b *testing.B) {
	i := loadUglyMainMemo(b)
	for _, enabled := range []bool{false, true} {
		name := "baseline"
		if enabled {
			name = "memo"
		}
		b.Run(name, func(b *testing.B) {
			setMainAccumulator(i, enabled)
			a, c, out := 19, 18, 0
			b.ReportAllocs()
			for b.Loop() {
				out = 0
				if _, err := i.RunFunc("f", 20, &a, &c, &out); err != nil || out != 6765 {
					b.Fatal(out, err)
				}
			}
		})
	}
}

func TestInterpreterAccumulatorRejectsEffects(t *testing.T) {
	for _, body := range []string{"*a += 1", "f(x-1,p,b,p)", "if a == p { x-- }"} {
		src := `package main
func f(x int,a,b,p *int) {if x<2 {*p+=x;return};` + body + `;f(x-1,a,b,p);f(x-2,a,b,p)}`
		i := loadMainMemo(t, []byte(src))
		if i.funcs[i.mainpkg.Func("f")].interpreterAccumulator {
			t.Fatal("unsafe effects accepted", body)
		}
	}
}

func TestInterpreterAccumulatorUint8Overflow(t *testing.T) {
	i := loadMainMemo(t, []byte(`package main
func f(x uint8,a,b,p *uint8) {if x<2 {*p+=x;return};f(x-1,a,b,p);f(x-2,a,b,p)}
`))
	p := i.funcs[i.mainpkg.Func("f")]
	if !p.interpreterAccumulator {
		t.Fatal("memo disabled")
	}
	var want uint8
	for _, enabled := range []bool{false, true} {
		p.interpreterAccumulator = enabled
		a, b, out := uint8(0), uint8(0), uint8(250)
		if _, err := i.RunFunc("f", uint8(12), &a, &b, &out); err != nil {
			t.Fatal(err)
		}
		if !enabled {
			want = out
		} else if out != want {
			t.Fatal(out, want)
		}
	}
	if _, err := i.RunFunc("f", uint8(0), new(uint8), new(uint8), (*uint8)(nil)); err == nil {
		t.Fatal("nil output panic missing")
	}
}

func TestInterpreterAccumulatorGeneralSignatures(t *testing.T) {
	i := loadMainMemo(t, []byte(`package main
// Two parameters, output first.
func sum(out *int,n int) {if n<2 {*out+=n;return};sum(out,n-1);sum(out,n-2)}
// Three parameters, additional scalar input participates in every leaf.
func weighted(n int,out *int,weight int) {if n<2 {*out+=n*weight;return};weighted(n-1,out,weight);weighted(n-2,out,weight)}
// Five parameters, mixed scalar types and reordered output in mutual calls.
func mixed(a *int,n int,out *int,scale int8,b *int) {if n<2 {*out+=n*int(scale)+*a+*b;return};shift(out,b,n-1,a,scale);shift(out,b,n-2,a,scale)}
func shift(out *int,b *int,n int,a *int,scale int8) {mixed(a,n,out,scale,b)}
// Conversion-free mixed widths are also supported.
func widths(out *int64,n int,weight int64) {if n<2 {*out+=weight;return};widths(out,n-1,weight);widths(out,n-2,weight)}
// Different arities within one call graph.
func outer(n int,out *int,extra int) {inner(out,n);inner(out,extra)}
func inner(out *int,n int) {*out+=n}
`))
	// Numerical conversions are supported by full-state caching.
	if !i.funcs[i.mainpkg.Func("mixed")].interpreterAccumulator {
		t.Fatal("conversion rejected")
	}
	for _, name := range []string{"sum", "weighted", "widths", "outer", "inner"} {
		if !i.funcs[i.mainpkg.Func(name)].interpreterAccumulator {
			t.Fatal("generic signature rejected", name)
		}
	}
	for _, enabled := range []bool{false, true} {
		for _, p := range i.funcs {
			if p.Fn.Package() == i.mainpkg {
				p.interpreterAccumulator = enabled
			}
		}
		out := 7
		if _, err := i.RunFunc("sum", &out, 15); err != nil || out != 617 {
			t.Fatal(out, err)
		}
		out = 7
		if _, err := i.RunFunc("weighted", 15, &out, 3); err != nil || out != 1837 {
			t.Fatal(out, err)
		}
		wide := int64(7)
		if _, err := i.RunFunc("widths", &wide, 10, int64(3)); err != nil || wide != 274 {
			t.Fatal(wide, err)
		}
		out = 7
		if _, err := i.RunFunc("outer", 5, &out, 9); err != nil || out != 21 {
			t.Fatal(out, err)
		}
	}
}

func TestInterpreterAccumulatorGeneralPointerInputs(t *testing.T) {
	i := loadMainMemo(t, []byte(`package main
func many(a *int,n int,out *int,b *int,c *int,weight int) {if n<2 {*out+=n*weight+*a+*b+*c;return};other(out,n-1,c,b,a,weight);other(out,n-2,c,b,a,weight)}
func other(out *int,n int,c,b,a *int,weight int) {many(a,n,out,b,c,weight)}
`))
	for _, name := range []string{"many", "other"} {
		if !i.funcs[i.mainpkg.Func(name)].interpreterAccumulator {
			t.Fatal(name)
		}
	}
	var want int
	for _, enabled := range []bool{false, true} {
		for _, name := range []string{"many", "other"} {
			i.funcs[i.mainpkg.Func(name)].interpreterAccumulator = enabled
		}
		a, b, c, out := 1, 2, 3, 7
		if _, err := i.RunFunc("many", &a, 12, &out, &b, &c, 4); err != nil {
			t.Fatal(err)
		}
		if !enabled {
			want = out
		} else if out != want {
			t.Fatal(out, want)
		}
	}
}
