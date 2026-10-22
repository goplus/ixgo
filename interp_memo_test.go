/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package ixgo

import (
	"fmt"
	"golang.org/x/tools/go/ssa"
	"math"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func testValueBits(v reflect.Value) uint64 {
	switch v.Kind() {
	case reflect.Float32:
		return uint64(math.Float32bits(v.Interface().(float32)))
	case reflect.Float64:
		return math.Float64bits(v.Float())
	default:
		if v.Kind() >= reflect.Int && v.Kind() <= reflect.Int64 {
			return uint64(v.Int())
		}
		return v.Uint()
	}
}

func testSetScalarBits(v reflect.Value, bits uint64) {
	switch v.Kind() {
	case reflect.Float32:
		v.SetFloat(float64(math.Float32frombits(uint32(bits))))
	case reflect.Float64:
		v.SetFloat(math.Float64frombits(bits))
	default:
		if v.Kind() >= reflect.Int && v.Kind() <= reflect.Int64 {
			v.SetInt(int64(bits))
		} else {
			v.SetUint(bits)
		}
	}
}

func loadMainMemo(t testing.TB, source []byte) *Interp {
	t.Helper()
	i, err := NewContext(0).LoadInterp("memo.go", source)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(i.UnsafeRelease)
	return i
}

// Loading without the plugin must not run in parallel with other interpreter loads.
func loadMainWithoutMemo(t testing.TB, source []byte) *Interp {
	t.Helper()
	i, err := func() (*Interp, error) {
		old := memoPlugin
		memoPlugin = nil
		defer func() { memoPlugin = old }()
		return NewContext(0).LoadInterp("memo.go", source)
	}()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(i.UnsafeRelease)
	for _, p := range i.funcs {
		if p.memoFlags != 0 {
			t.Fatal("memo enabled without plugin", p.Fn)
		}
	}
	return i
}

func loadMainMemoPair(t testing.TB, source []byte) (baseline, memo *Interp) {
	t.Helper()
	baseline = loadMainWithoutMemo(t, source)
	memo = loadMainMemo(t, source)
	return
}

func loadUglyMainMemoPair(t testing.TB) (baseline, memo *Interp) {
	t.Helper()
	src, err := os.ReadFile("testdata/fibish/uglyfib.go")
	if err != nil {
		t.Fatal(err)
	}
	baseline = loadMainWithoutMemo(t, src)
	memo = loadUglyMainMemo(t)
	return
}

func TestInterpreterMemo(t *testing.T) {
	for _, name := range []string{"fibish", "fibish_closure", "fibish_float32", "fibish_float64"} {
		t.Run(name, func(t *testing.T) {
			src, err := os.ReadFile("testdata/fibish/" + name + ".go")
			if err != nil {
				t.Fatal(err)
			}
			entry := "f"
			if name == "fibish_closure" {
				src = append(src, []byte("\nfunc runMemoInput(n int) (int, int) { return f(n, f) }\n")...)
				entry = "runMemoInput"
			}
			baseline, i := loadMainMemoPair(t, src)
			p := i.funcs[i.mainpkg.Func("f")]
			if p.memoFlags&interpreterMemoScalar == 0 {
				t.Fatal("memo not enabled")
			}
			for n := 0; n <= 20; n++ {
				var x Value = n
				if name == "fibish_float32" {
					x = float32(n)
				}
				if name == "fibish_float64" {
					x = float64(n)
				}

				a, err := i.RunFunc(entry, x)
				if err != nil {
					t.Fatal(err)
				}
				b, err := baseline.RunFunc(entry, x)
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
	if i.funcs[i.mainpkg.Func("effects")].memoFlags&interpreterMemoScalar != 0 {
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
		baseline, i := loadMainMemoPair(b, src)
		if i.funcs[i.mainpkg.Func("f")].memoFlags&interpreterMemoScalar == 0 {
			b.Fatal("memo not enabled")
		}
		for pass, i := range []*Interp{baseline, i} {
			label := "baseline"
			if pass == 1 {
				label = "memo"
			}
			b.Run(name+"/"+label, func(b *testing.B) {
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
		if i.funcs[i.mainpkg.Func("f")].memoFlags&interpreterMemoScalar != 0 {
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
	for _, name := range []string{"fibish", "fibish_closure", "fibish_float32", "fibish_float64"} {
		src, err := os.ReadFile("testdata/fibish/" + name + ".go")
		if err != nil {
			t.Fatal(err)
		}
		entry := "f"
		if name == "fibish_closure" {
			src = append(src, []byte("\nfunc runMemoInput(n int) (int, int) { return f(n, f) }\n")...)
			entry = "runMemoInput"
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
		v, err := i.RunFunc(entry, x)
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
		if p.memoFlags&interpreterMemoScalar != 0 || p.memoFlags&interpreterMemoAccumulator != 0 {
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
		if i.funcs[i.mainpkg.Func(name)].memoFlags&interpreterMemoAccumulator == 0 {
			t.Fatalf("%s not optimized", name)
		}
	}
	return i
}

func TestInterpreterAccumulator(t *testing.T) {
	baseline, i := loadUglyMainMemoPair(t)
	for _, name := range []string{"f", "g", "h", "k"} {
		for n := 0; n <= 16; n++ {
			for alias := 0; alias < 3; alias++ {
				if alias != 0 && n > 3 {
					continue
				}
				var got [2][3]int
				for pass, i := range []*Interp{baseline, i} {
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
	baseline, i := loadMainMemoPair(t, []byte(`package main
func leaf(x int64,a,b,p *int64) { *p+=1;if *a==0 {*p+=2} }
func root(x int64,a,b,p *int64) {leaf(x,a,b,p);leaf(x,a,b,p)}
`))
	for _, name := range []string{"root", "leaf"} {
		if i.funcs[i.mainpkg.Func(name)].memoFlags&interpreterMemoAccumulator == 0 {
			t.Fatal("memo missing", name)
		}
	}
	run := func(i *Interp) int64 {
		storage := new([3]int64)
		p := &storage[0]
		a := (*int64)(unsafe.Add(unsafe.Pointer(p), 1))
		*p = 252
		if _, err := i.RunFunc("root", int64(0), a, &storage[2], p); err != nil {
			t.Fatal(err)
		}
		return *p
	}
	if a, b := run(baseline), run(i); a != b {
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
	baseline, i := loadUglyMainMemoPair(b)
	for pass, i := range []*Interp{baseline, i} {
		name := "baseline"
		if pass == 1 {
			name = "memo"
		}
		b.Run(name, func(b *testing.B) {

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
		if i.funcs[i.mainpkg.Func("f")].memoFlags&interpreterMemoAccumulator != 0 {
			t.Fatal("unsafe effects accepted", body)
		}
	}
}

func TestInterpreterAccumulatorUint8Overflow(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
func f(x uint8,a,b,p *uint8) {if x<2 {*p+=x;return};f(x-1,a,b,p);f(x-2,a,b,p)}
`))
	p := i.funcs[i.mainpkg.Func("f")]
	if p.memoFlags&interpreterMemoAccumulator == 0 {
		t.Fatal("memo disabled")
	}
	var want uint8
	for pass, i := range []*Interp{baseline, i} {
		a, b, out := uint8(0), uint8(0), uint8(250)
		if _, err := i.RunFunc("f", uint8(12), &a, &b, &out); err != nil {
			t.Fatal(err)
		}
		if pass == 0 {
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
	baseline, i := loadMainMemoPair(t, []byte(`package main
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
	if i.funcs[i.mainpkg.Func("mixed")].memoFlags&interpreterMemoAccumulator == 0 {
		t.Fatal("conversion rejected")
	}
	for _, name := range []string{"sum", "weighted", "widths", "outer", "inner"} {
		if i.funcs[i.mainpkg.Func(name)].memoFlags&interpreterMemoAccumulator == 0 {
			t.Fatal("generic signature rejected", name)
		}
	}
	for _, i := range []*Interp{baseline, i} {
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
	baseline, i := loadMainMemoPair(t, []byte(`package main
func many(a *int,n int,out *int,b *int,c *int,weight int) {if n<2 {*out+=n*weight+*a+*b+*c;return};other(out,n-1,c,b,a,weight);other(out,n-2,c,b,a,weight)}
func other(out *int,n int,c,b,a *int,weight int) {many(a,n,out,b,c,weight)}
`))
	for _, name := range []string{"many", "other"} {
		if i.funcs[i.mainpkg.Func(name)].memoFlags&interpreterMemoAccumulator == 0 {
			t.Fatal(name)
		}
	}
	var want int
	for pass, i := range []*Interp{baseline, i} {
		a, b, c, out := 1, 2, 3, 7
		if _, err := i.RunFunc("many", &a, 12, &out, &b, &c, 4); err != nil {
			t.Fatal(err)
		}
		if pass == 0 {
			want = out
		} else if out != want {
			t.Fatal(out, want)
		}
	}
}

func TestMemoRequiresRegistration(t *testing.T) {
	src, err := os.ReadFile("testdata/fibish/fibish.go")
	if err != nil {
		t.Fatal(err)
	}
	i := loadMainWithoutMemo(t, src)
	for _, fn := range i.funcs {
		if fn.memoFlags&interpreterMemoScalar != 0 || fn.memoFlags&interpreterMemoReceiver != 0 || fn.memoFlags&interpreterMemoAccumulator != 0 {
			t.Fatal("memo enabled without plugin")
		}
	}
}

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
				baseline, i := loadMainMemoPair(t, []byte(src))
				p := i.funcs[i.mainpkg.Func("f")]
				if p.memoFlags&interpreterMemoAccumulator == 0 {
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
				for pass, i := range []*Interp{baseline, i} {
					out := reflect.New(typ)
					testSetScalarBits(out.Elem(), testValueBits(x))
					if _, err := i.RunFunc("f", 8, x.Interface(), out.Interface()); err != nil {
						t.Fatal(err)
					}
					got := testValueBits(out.Elem())
					if pass == 0 {
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
	baseline, i := loadMainMemoPair(t, []byte(`package main
func leaf(x float64,p *float64) {*p /= x}
func root(x float64,p *float64) {leaf(x,p);*p = -0.0;leaf(x,p)}
`))
	for _, x := range []float64{0, math.Copysign(0, -1), math.Inf(1), math.Inf(-1), math.Float64frombits(0x7ff8000000000001)} {
		var want uint64
		for pass, i := range []*Interp{baseline, i} {
			out := math.Copysign(0, -1)
			if _, err := i.RunFunc("root", x, &out); err != nil {
				t.Fatal(err)
			}
			got := math.Float64bits(out)
			if pass == 0 {
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
		if i.funcs[i.mainpkg.Func("f")].memoFlags&interpreterMemoAccumulator == 0 {
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
		baseline, i := loadMainMemoPair(b, []byte(fmt.Sprintf(`package main
func f(n int,x %[1]s,p *%[1]s) {if n==0 {*p *= x;return};*p=7;f(n-1,x,p);*p=7;f(n-1,x,p)}
`, kind)))
		if i.funcs[i.mainpkg.Func("f")].memoFlags&interpreterMemoAccumulator == 0 {
			b.Fatal("state cache missing")
		}
		for pass, i := range []*Interp{baseline, i} {
			name := "baseline"
			if pass == 1 {
				name = "state"
			}
			b.Run(kind+"/"+name, func(b *testing.B) {
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
	baseline, i := loadMainMemoPair(t, []byte(`package main
func f(n int,a,p *int) {if n==0 {*p *= *a;return};f(n-1,a,p);f(n-1,a,p)}
`))
	p := i.funcs[i.mainpkg.Func("f")]
	if p.memoFlags&interpreterMemoAccumulator == 0 {
		t.Fatal("state cache missing")
	}
	var want int
	for pass, i := range []*Interp{baseline, i} {
		out := 2
		if _, err := i.RunFunc("f", 2, &out, &out); err != nil {
			t.Fatal(err)
		}
		if pass == 0 {
			want = out
		} else if out != want {
			t.Fatal(out, want)
		}
	}
}

func receiverFunctions(i *Interp) []*function {
	var out []*function
	for fn, p := range i.funcs {
		if fn.Package() == i.mainpkg && fn.Signature.Recv() != nil {
			out = append(out, p)
		}
	}
	return out
}
func TestComposeReceiverParentRejectsIncompatiblePlans(t *testing.T) {
	type state struct{ Value int }
	receiver := &state{}
	field := reflect.ValueOf(receiver).Elem().Field(0)
	parentPlan := &MemoReceiverPlan{Receiver: &MemoReceiverState{
		Fields: [][]int{{0}}, Additive: []bool{false}, Affine: []bool{false}, Bitwise: []bool{true},
	}}
	childPlan := &MemoReceiverPlan{Receiver: &MemoReceiverState{
		Fields: [][]int{{0}}, Additive: []bool{false}, Affine: []bool{true}, Bitwise: []bool{false},
	}}
	parent := &frame{memoState: &interpreterMemoState{receiver: &interpreterReceiverState{
		memo: &interpreterReceiverMemo{plan: parentPlan, receiver: receiver}, fields: []reflect.Value{field}, affine: []receiverTransform{{a: 1}},
	}}}
	child := &frame{caller: parent, memoState: &interpreterMemoState{receiver: &interpreterReceiverState{
		memo: &interpreterReceiverMemo{plan: childPlan, receiver: receiver}, fields: []reflect.Value{field},
	}}}
	composeReceiverParent(child, []receiverTransform{{a: 2, b: 3}})
	if got := parent.memoState.receiver.affine[0]; got != (receiverTransform{a: 1}) {
		t.Fatalf("incompatible receiver plans composed transform: %#v", got)
	}
}

func TestInterpreterReceiverStateAndHits(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type state struct { value float64; unused *int; spare string }
 func (p *state) step(n int,x float64) float64 {if n==0 {p.value /= x;return p.value};p.value=7;a:=p.step(n-1,x);p.value=7;return a+p.step(n-1,x)}
 func run(n int,x float64)(float64,float64,int,string){u:=42;p:=&state{value:3,unused:&u,spare:"untouched"};v:=p.step(n,x);return v,p.value,*p.unused,p.spare}
 `))
	methods := receiverFunctions(i)
	if len(methods) != 1 || methods[0].memoFlags&interpreterMemoReceiver == 0 {
		t.Fatal("receiver disabled")
	}
	p := methods[0]
	var hits uint64
	for j, instr := range p.ssaInstrs {
		if _, ok := instr.(*ssa.Return); ok {
			ret := p.Instrs[j]
			p.Instrs[j] = func(fr *frame) {
				ret(fr)
				if fr.memoState != nil && fr.memoState.receiver.memo != nil {
					hits += fr.memoState.receiver.memo.hits
				}
			}
		}
	}
	for _, x := range []float64{3, 0, math.Copysign(0, -1), math.Inf(1), math.Float64frombits(0x7ff8000000000001)} {
		for n := 0; n <= 8; n++ {
			var want Tuple
			for pass, i := range []*Interp{baseline, i} {
				v, err := i.RunFunc("run", n, x)
				if err != nil {
					t.Fatal(err)
				}
				got := v.(Tuple)
				if pass == 0 {
					want = got
				} else {
					for j := 0; j < 2; j++ {
						if math.Float64bits(got[j].(float64)) != math.Float64bits(want[j].(float64)) {
							t.Fatal(n, x, got, want)
						}
					}
					if got[2] != 42 || got[3] != "untouched" {
						t.Fatal(got)
					}
				}
			}
		}
	}
	if hits == 0 {
		t.Fatal("no actual cache hits")
	}
}

func TestInterpreterReceiverSafety(t *testing.T) {
	for _, body := range []string{"count++", "println(n)", "p.ref = new(int)", "*a += 1", "if a == &p.v { n-- }", "p.v += *p.ref", "other.v++"} {
		t.Run(body, func(t *testing.T) {
			src := `package main
var count int
var other S
type S struct{v int; ref *int}
func(p *S) f(n int,a *int)int{` + body + `;if n<1{return p.v};return p.f(n-1,a)}
func run(n int)int{p:=&S{};a:=0;return p.f(n,&a)}
`
			i := loadMainMemo(t, []byte(src))
			for _, p := range receiverFunctions(i) {
				if p.memoFlags&interpreterMemoReceiver != 0 {
					t.Fatal("unsafe method enabled")
				}
			}
		})
	}
	baseline, i := loadMainMemoPair(t, []byte(`package main
type S struct{v int}
func(p *S) f(n int,a *int)int{if n==0{p.v++;return *a};return p.f(n-1,a)+p.f(n-1,a)}
func run(n int,alias bool)(int,int){p:=&S{v:2};a:=7;q:=&a;if alias{q=&p.v};v:=p.f(n,q);return v,p.v}
func nilrun(n int)int{var p *S;return p.f(n,nil)}
`))
	p := receiverFunctions(i)[0]
	if p.memoFlags&interpreterMemoReceiver == 0 {
		t.Fatal("safe method rejected")
	}
	for _, alias := range []bool{false, true} {
		for n := 0; n <= 8; n++ {
			var want Value
			for pass, i := range []*Interp{baseline, i} {
				v, err := i.RunFunc("run", n, alias)
				if err != nil {
					t.Fatal(err)
				}
				if pass == 0 {
					want = v
				} else if !reflect.DeepEqual(v, want) {
					t.Fatal(alias, n, v, want)
				}
			}
		}
	}
	for _, i := range []*Interp{baseline, i} {
		for j := 0; j < 2; j++ {
			if _, err := i.RunFunc("nilrun", 2); err == nil {
				t.Fatal("nil panic missing")
			}
		}
	}
}

func TestInterpreterReceiverHooks(t *testing.T) {
	for _, kind := range []string{"entry", "debug", "trace"} {
		t.Run(kind, func(t *testing.T) {
			mode := Mode(0)
			if kind == "trace" {
				mode = EnableTracing
			}
			ctx := NewContext(mode)
			if kind == "entry" {
				ctx.SetFuncHook(func(*FuncInfo) {}, nil)
			}
			if kind == "debug" {
				ctx.SetDebug(func(*DebugInfo) {})
			}
			i, err := ctx.LoadInterp("receiver.go", `package main
type S struct{v int}
func(p *S) f(n int)int{if n<2{return p.v};return p.f(n-1)+p.f(n-2)}
func run()int{return (&S{}).f(3)}`)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			for _, p := range receiverFunctions(i) {
				if p.memoFlags&interpreterMemoReceiver != 0 {
					t.Fatal("instrumented method enabled")
				}
			}
		})
	}
}

func TestInterpreterReceiverNilAndGraphSharing(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type S struct { v int; unused int }
 func(p *S) leaf(n int)int{if n<0{return 17};if n==0{p.v++;return p.v};return p.leaf(n-1)}
 func(p *S) outer(n int)int{if n==0{return p.leaf(0)};return p.outer(n-1)+p.leaf(n)}
 func nilok()int{var p *S;return p.leaf(-1)}
 func nilbad()int{var p *S;return p.leaf(0)}
 func run(n int)(int,int,int){p:=&S{v:3,unused:29};a:=p.outer(n);return a,p.v,p.unused}
 `))
	methods := receiverFunctions(i)
	for _, p := range methods {
		if p.memoFlags&interpreterMemoReceiver == 0 {
			t.Fatal(p.Fn)
		}
	}
	for n := 0; n <= 16; n++ {
		var want Value
		for pass, i := range []*Interp{baseline, i} {
			v, err := i.RunFunc("run", n)
			if err != nil {
				t.Fatal(err)
			}
			if pass == 0 {
				want = v
			} else if !reflect.DeepEqual(v, want) {
				t.Fatal(n, v, want)
			}
		}
	}
	for _, i := range []*Interp{baseline, i} {
		if v, err := i.RunFunc("nilok"); err != nil || v != 17 {
			t.Fatal(v, err)
		}
		if _, err := i.RunFunc("nilbad"); err == nil {
			t.Fatal("nil receiver did not panic")
		}
	}
}

func TestInterpreterReceiverFloat32AndNested(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type inner struct{ x float32 }
 type S struct{ inner inner; untouched *int }
 func(p *S) f(n int,x float32)float32{if n==0{p.inner.x=x;return x};p.inner.x=7;a:=p.f(n-1,x);p.inner.x=7;return a+p.f(n-1,x)}
 func run(n int,x float32)(float32,float32){p:=&S{};a:=p.f(n,x);return a,p.inner.x}
 `))
	p := receiverFunctions(i)[0]
	if p.memoFlags&interpreterMemoReceiver == 0 {
		t.Fatal("nested state rejected")
	}
	for _, x := range []float32{math.Float32frombits(0x80000000), math.Float32frombits(0x7fc00123), 1.25} {
		for n := 0; n <= 8; n++ {
			var want Tuple
			for pass, i := range []*Interp{baseline, i} {
				v, err := i.RunFunc("run", n, x)
				if err != nil {
					t.Fatal(err)
				}
				got := v.(Tuple)
				if pass == 0 {
					want = got
				} else {
					for j := range got {
						if math.Float32bits(got[j].(float32)) != math.Float32bits(want[j].(float32)) {
							t.Fatal(n, got, want)
						}
					}
				}
			}
		}
	}
}

func TestInterpreterReceiverConcurrent(t *testing.T) {
	var interps []*Interp
	for j := 0; j < 4; j++ {
		interps = append(interps, loadMainMemo(t, []byte(`package main
 type S struct{v int}
 func(p *S) f(n int)int{if n<2{return n+p.v};return p.f(n-1)+p.f(n-2)}
 func run(n int)int{return (&S{}).f(n)}
 `)))
	}
	var wg sync.WaitGroup
	for _, i := range interps {
		wg.Add(1)
		go func(i *Interp) {
			defer wg.Done()
			for n := 0; n < 10; n++ {
				v, err := i.RunFunc("run", 16)
				if err != nil || v != 987 {
					t.Errorf("%v %v", v, err)
				}
			}
		}(i)
	}
	wg.Wait()
}

func receiverHits(methods []*function) *uint64 {
	hits := new(uint64)
	for _, p := range methods {
		for j, instr := range p.ssaInstrs {
			if _, ok := instr.(*ssa.Return); ok {
				ret := p.Instrs[j]
				p.Instrs[j] = func(fr *frame) {
					ret(fr)
					if fr.memoState != nil && fr.memoState.receiver.memo != nil {
						*hits += fr.memoState.receiver.memo.hits
					}
				}
			}
		}
	}
	return hits
}

func TestInterpreterReceiverMixedAdditive(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type inner struct{sum int8; sub uint8}
 type S struct{inner inner;wide int64; ro int; state int; floating float64;unused string}
 func(p *S) f(n int) int {
  if n==0 {p.inner.sum=p.inner.sum+int8(p.ro);p.inner.sub-=uint8(p.ro);p.wide+=int64(p.ro);p.state*=3;p.floating+=1.25;return p.state}
  p.state=2;p.floating=4;a:=p.f(n-1);p.state=2;p.floating=4;return a+p.f(n-1)
 }
 func run(n int, x int8,y uint8,z int64,ro int)(int,int8,uint8,int64,int,int,float64,string){p:=&S{inner:inner{x,y},wide:z,ro:ro,state:5,floating:8,unused:"kept"};a:=p.f(n);return a,p.inner.sum,p.inner.sub,p.wide,p.ro,p.state,p.floating,p.unused}
 `))
	methods := receiverFunctions(i)
	if len(methods) != 1 || methods[0].memoFlags&interpreterMemoReceiver == 0 {
		t.Fatal("disabled")
	}
	hits := receiverHits(methods)
	for _, x := range []int8{0, 127, -128, -9} {
		for _, ro := range []int{3, 7} {
			for n := 0; n <= 9; n++ {
				var want Value
				for pass, i := range []*Interp{baseline, i} {
					got, err := i.RunFunc("run", n, x, uint8(2), int64(math.MaxInt64-1), ro)
					if err != nil {
						t.Fatal(err)
					}
					if pass == 0 {
						want = got
					} else if !memoResultsEqual(got, want) {
						t.Fatal(n, x, ro, got, want)
					}
				}
			}
		}
	}
	if *hits == 0 {
		t.Fatal("no mixed-state hits")
	}
}

func TestInterpreterReceiverAdditiveRejections(t *testing.T) {
	for _, body := range []string{
		`if p.v>0 {p.v+=n}`, `p.v+=n;return p.v`, `p.v+=p.v`, `p.v=n-p.v`,
		`p.v*=2`, `p.v+=p.g(0,p.v)`, `old:=p.v;p.other=1;p.v=old+n`,
		`old:=p.v;p.g(0,0);p.v=old+n`, `old:=p.v;p.other+=old;p.v=old+n`,
	} {
		t.Run(body, func(t *testing.T) {
			src := `package main
 type S struct{v int;other int}
 func(p *S) f(n int) int {if n==0 {` + body + `;return 0};return p.g(n-1,0)+p.g(n-1,0)}
 func(p *S) g(n,x int) int {p.v+=x;if n==0{return 0};return p.f(n-1)}
 func run(n,initial int)(int,int,int){p:=&S{v:initial};a:=p.f(n);return a,p.v,p.other}`
			baseline, i := loadMainMemoPair(t, []byte(src))
			methods := receiverFunctions(i)
			for _, p := range methods {
				if p.memoFlags&interpreterMemoReceiver == 0 {
					t.Fatal("full-state fallback disabled", p.Fn)
				}
			}
			for _, initial := range []int{-2, 3} {
				var want Value
				for pass, i := range []*Interp{baseline, i} {
					got, err := i.RunFunc("run", 6, initial)
					if err != nil {
						t.Fatal(err)
					}
					if pass == 0 {
						want = got
					} else if !memoResultsEqual(got, want) {
						t.Fatal(got, want)
					}
				}
			}
		})
	}
}

func TestInterpreterReceiverAffine(t *testing.T) {
	for _, kind := range []string{"int8", "uint8", "int64", "uint64"} {
		t.Run(kind, func(t *testing.T) {
			src := `package main
 type inner struct{x T; y T}
 type S struct{inner inner; state int; ro T; unused string}
 func(p *S) f(n int, factor T) int {
  if n==0 {p.inner.x*=factor;p.inner.x+=p.ro;p.inner.y-=factor;p.inner.y*=factor;return p.state}
  p.inner.x+=T(n);p.state=7;a:=p.g(n-1,factor);p.inner.x*=factor;p.state=7;return a+p.g(n-1,factor)
 }
 func(p *S) g(n int,factor T)int {p.inner.y*=factor;p.inner.y+=T(n);if n==0{return p.f(0,factor)};return p.f(n-1,factor)+p.f(n-1,factor)}
 func run(n int, initial, factor int)(int,T,T,int,T,string){p:=&S{inner:inner{T(initial),T(initial+1)},state:7,ro:3,unused:"kept"};a:=p.f(n,T(factor));return a,p.inner.x,p.inner.y,p.state,p.ro,p.unused}
 `
			baseline, i := loadMainMemoPair(t, []byte(strings.ReplaceAll(src, "T", kind)))
			methods := receiverFunctions(i)
			hits := receiverHits(methods)
			for _, p := range methods {
				if p.memoFlags&interpreterMemoReceiver == 0 {
					t.Fatal("affine proof", p.Fn)
				}
			}
			for _, initial := range []int{0, 1, -3, 127, 255, math.MaxInt64 - 1} {
				for _, factor := range []int{0, 1, -1, 2, -3, 129} {
					for n := 0; n <= 7; n++ {
						var want Value
						for pass, i := range []*Interp{baseline, i} {
							got, err := i.RunFunc("run", n, initial, factor)
							if err != nil {
								t.Fatal(err)
							}
							if pass == 0 {
								want = got
							} else if !memoResultsEqual(got, want) {
								t.Fatal(n, initial, factor, got, want)
							}
						}
					}
				}
			}
			if *hits == 0 {
				t.Fatal("no affine hits")
			}
		})
	}
}

func TestInterpreterReceiverAffineSeparateGraphs(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type S struct{x int}
 func(p *S) leaf(n,k int){p.x+=k;if n>0{p.leaf(n-1,k);p.leaf(n-1,k)}}
 func(p *S) outer(n,k int){p.x*=k;p.leaf(2,k);if n>0{p.outer(n-1,k);p.outer(n-1,k)};p.x-=n}
 func run(n,x,k int)int{p:=&S{x:x};p.outer(n,k);return p.x}
 `))
	methods := receiverFunctions(i)
	hits := receiverHits(methods)
	for _, x := range []int{0, 7, -9} {
		for _, k := range []int{0, -2, 3} {
			var want Value
			for pass, i := range []*Interp{baseline, i} {
				got, err := i.RunFunc("run", 6, x, k)
				if err != nil {
					t.Fatal(err)
				}
				if pass == 0 {
					want = got
				} else if got != want {
					t.Fatal(x, k, got, want)
				}
			}
		}
	}
	if *hits == 0 {
		t.Fatal("no separate graph hits")
	}
}

func TestInterpreterReceiverAffineRejections(t *testing.T) {
	for _, body := range []string{`p.v*=p.v`, `p.v/=k`, `if p.v>0{p.v*=k}`, `p.v*=k;return p.v`, `old:=p.v;p.other=1;p.v=old*k`, `p.v*=p.other;p.other*=p.v`} {
		t.Run(body, func(t *testing.T) {
			baseline, i := loadMainMemoPair(t, []byte(`package main
 type S struct{v int;other int}
 func(p *S) f(n,k int)int{if n==0{`+body+`;return 0};return p.f(n-1,k)+p.f(n-1,k)}
 func run(n,x,k int)(int,int,int){p:=&S{v:x,other:2};a:=p.f(n,k);return a,p.v,p.other}`))
			methods := receiverFunctions(i)
			for _, p := range methods {
				if p.memoFlags&interpreterMemoReceiver == 0 {
					t.Fatal("unsafe affine rejected", p.Fn)
				}
			}
			for _, x := range []int{0, -2, 3} {
				var want Value
				for pass, i := range []*Interp{baseline, i} {
					got, err := i.RunFunc("run", 5, x, 2)
					if err != nil {
						t.Fatal(err)
					}
					if pass == 0 {
						want = got
					} else if !memoResultsEqual(got, want) {
						t.Fatal(got, want)
					}
				}
			}
		})
	}
}

func BenchmarkInterpreterReceiverAffine(b *testing.B) {
	baseline, i := loadMainMemoPair(b, []byte(`package main
 type S struct{x int64;y uint8}
 func(p *S) f(n int,k int64){if n==0{p.x*=k;p.x+=3;p.y*=uint8(k);p.y-=7;return};p.f(n-1,k);p.x-=int64(n);p.f(n-1,k)}
 func run(n int)int64{p:=&S{x:7};p.f(n,-3);return p.x}`))
	for _, p := range receiverFunctions(i) {
		if p.memoFlags&interpreterMemoReceiver == 0 {
			b.Fatal("disabled")
		}
	}
	for pass, i := range []*Interp{baseline, i} {
		name := "off"
		if pass == 1 {
			name = "on"
		}
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				if _, err := i.RunFunc("run", 16); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func TestInterpreterReceiverAffinePointerOperand(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type Number int16
 type S struct{x Number}
 func(p *S) f(n int,k *Number){if n==0{p.x=*k*p.x;p.x=Number(n+3)+p.x;return};p.f(n-1,k);p.x-=Number(n);p.f(n-1,k)}
 func run(n,x,k int)Number{p:=&S{x:Number(x)};factor:=Number(k);p.f(n,&factor);return p.x}`))
	methods := receiverFunctions(i)
	hits := receiverHits(methods)
	for _, p := range methods {
		if p.memoFlags&interpreterMemoReceiver == 0 {
			t.Fatal("pointer factor rejected", p.Fn)
		}
	}
	for _, x := range []int{0, 7, 32767} {
		for _, k := range []int{0, -3, 129} {
			var want Value
			for pass, i := range []*Interp{baseline, i} {
				got, err := i.RunFunc("run", 9, x, k)
				if err != nil {
					t.Fatal(err)
				}
				if pass == 0 {
					want = got
				} else if !memoResultsEqual(got, want) {
					t.Fatal(x, k, got, want)
				}
			}
		}
	}
	if *hits == 0 {
		t.Fatal("no pointer factor hits")
	}
}

func TestInterpreterReceiverAffineNonrecursiveHelper(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type S struct{x int}
 func(p *S) update(k int){p.x*=k;p.x+=3}
 func(p *S) f(n,k int){p.update(k);if n>0{p.f(n-1,k);p.f(n-1,k)};p.x-=n}
 func run(n,x,k int)int{p:=&S{x:x};p.f(n,k);return p.x}`))
	methods := receiverFunctions(i)
	hits := receiverHits(methods)
	for _, p := range methods {
		if p.memoFlags&interpreterMemoReceiver == 0 {
			t.Fatal("helper tracking disabled", p.Fn)
		}
	}
	for _, x := range []int{0, 7, -2} {
		for _, k := range []int{0, -3, 2} {
			var want Value
			for pass, i := range []*Interp{baseline, i} {
				got, err := i.RunFunc("run", 7, x, k)
				if err != nil {
					t.Fatal(err)
				}
				if pass == 0 {
					want = got
				} else if got != want {
					t.Fatal(x, k, got, want)
				}
			}
		}
	}
	if *hits == 0 {
		t.Fatal("no helper graph hits")
	}
}

func TestInterpreterReceiverBitwise(t *testing.T) {
	for _, kind := range []string{"int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint", "uintptr"} {
		t.Run(kind, func(t *testing.T) {
			src := strings.ReplaceAll(`package main
 type Number BASE
 type inner struct{x Number;y Number}
 type S struct{inner inner;z Number;sum Number;state int;unused string}
 func(p *S) update(k Number){p.inner.x &= k;p.inner.x ^= Number(5);p.inner.y |= k;p.inner.y &^= Number(2);p.z = ^p.z;p.sum = -p.sum;p.sum += Number(3)}
 func(p *S) f(n int,k *Number)int{p.update(*k);if n==0{return p.state};p.state=7;a:=p.g(n-1,k);p.inner.x=Number(n)^p.inner.x;p.inner.y=Number(n)|p.inner.y;p.state=7;return a+p.g(n-1,k)}
 func(p *S) g(n int,k *Number)int{p.inner.y=Number(127)&p.inner.y;p.z ^= Number(n);if n==0{return p.f(0,k)};return p.f(n-1,k)+p.f(n-1,k)}
 func run(n int,x,k uint64)(int,Number,Number,Number,Number,int,string){p:=&S{inner:inner{Number(x),Number(x+1)},z:Number(x),sum:Number(x),state:7,unused:"kept"};mask:=Number(k);a:=p.f(n,&mask);return a,p.inner.x,p.inner.y,p.z,p.sum,p.state,p.unused}
 `, "BASE", kind)
			baseline, i := loadMainMemoPair(t, []byte(src))
			methods := receiverFunctions(i)
			hits := receiverHits(methods)
			for _, p := range methods {
				if p.memoFlags&interpreterMemoReceiver == 0 {
					t.Fatal("bitwise proof", p.Fn)
				}
			}
			for _, x := range []uint64{0, 1, 0x80, 0xff, 0x8000000000000000, 0xaaaaaaaaaaaaaaaa, math.MaxUint64} {
				for _, k := range []uint64{0, 1, 0x55, 0x8000000000000000, math.MaxUint64} {
					for _, n := range []int{0, 1, 5} {
						receiverDifferential(t, baseline, i, "run", n, x, k)
					}
				}
			}
			if *hits == 0 {
				t.Fatal("no bitwise hits")
			}
		})
	}
}

func memoResultsEqual(got, want Value) bool {
	if a, ok := got.(Tuple); ok {
		b, ok := want.(Tuple)
		if !ok || len(a) != len(b) {
			return false
		}
		for j := range a {
			if !memoResultsEqual(a[j], b[j]) {
				return false
			}
		}
		return true
	}
	// Defined numeric types have distinct identities in independent interpreters.
	a, b := reflect.ValueOf(got), reflect.ValueOf(want)
	if a.IsValid() && b.IsValid() && a.Kind() == b.Kind() {
		switch a.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
			reflect.Float32, reflect.Float64:
			return testValueBits(a) == testValueBits(b)
		}
	}
	return reflect.DeepEqual(got, want)
}

func receiverDifferential(t *testing.T, baseline, i *Interp, name string, args ...Value) {
	t.Helper()
	var want Value
	for pass, i := range []*Interp{baseline, i} {
		got, err := i.RunFunc(name, args...)
		if err != nil {
			t.Fatal(name, args, pass, err)
		}
		if pass == 0 {
			want = got
		} else if !memoResultsEqual(got, want) {
			t.Fatal(name, args, got, want)
		}
	}
}

func TestInterpreterReceiverShiftNegation(t *testing.T) {
	for _, kind := range []string{"int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint", "uintptr"} {
		t.Run(kind, func(t *testing.T) {
			src := strings.ReplaceAll(`package main
 type Number BASE
 type S struct{x Number;y Number}
 func(p *S) update(s uint64,k Number){p.x <<= s;p.x = -p.x;p.x += Number(7);p.y *= k;p.y = -p.y;p.y <<= s;p.y -= Number(3)}
 func(p *S) f(n int,s uint64,k Number){p.update(s,k);if n>0{p.f(n-1,s,k);p.x = -p.x;p.f(n-1,s,k)}}
 func run(n int,x,s uint64,k int)(Number,Number){p:=&S{x:Number(x),y:Number(x+1)};p.f(n,s,Number(k));return p.x,p.y}
 `, "BASE", kind)
			baseline, i := loadMainMemoPair(t, []byte(src))
			methods := receiverFunctions(i)
			hits := receiverHits(methods)
			for _, p := range methods {
				if p.memoFlags&interpreterMemoReceiver == 0 {
					t.Fatal("shift/negation proof", p.Fn)
				}
			}
			for _, x := range []uint64{0, 1, 127, 0x8000000000000000, math.MaxUint64} {
				for _, s := range []uint64{0, 1, 7, 8, 15, 16, 31, 32, 63, 64, 65, 256, 1 << 63, math.MaxUint64} {
					for _, k := range []int{0, -1, 2, 128} {
						receiverDifferential(t, baseline, i, "run", 4, x, s, k)
					}
				}
			}
			if *hits == 0 {
				t.Fatal("no shift/negation hits")
			}
		})
	}
}

func TestInterpreterReceiverNegativeShift(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type S struct{x int8}
 func(p *S) f(n int,s int64){if n==0{p.x <<= s;return};p.f(n-1,s);p.f(n-1,s)}
 func run(n int,s int64)int8{p:=&S{x:-7};p.f(n,s);return p.x}
 `))
	methods := receiverFunctions(i)
	for _, p := range methods {
		if p.memoFlags&interpreterMemoReceiver == 0 {
			t.Fatal("shift proof", p.Fn)
		}
	}
	for pass, i := range []*Interp{baseline, i} {
		for _, s := range []int64{-1, -256, math.MinInt64} {
			for j := 0; j < 2; j++ {
				if _, err := i.RunFunc("run", 4, s); err == nil || !strings.Contains(err.Error(), "negative shift") {
					t.Fatal("negative shift panic", pass, s, err)
				}
			}
		}
	}
}

func TestInterpreterReceiverBitwiseSeparateGraphs(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type S struct{x int8;y uint64;z int16}
 func(p *S) leaf(n int,k uint64){p.x ^= int8(k);p.y &^= k;p.z = -p.z;if n>0{p.leaf(n-1,k);p.y |= uint64(n);p.leaf(n-1,k)}}
 func(p *S) helper(k uint64){p.x = ^p.x;p.y ^= k;p.z <<= k}
 func(p *S) outer(n int,k uint64){p.x |= int8(n);p.leaf(2,k);p.y &= k;p.helper(k);if n>0{p.outer(n-1,k);p.x &^= int8(n);p.outer(n-1,k)};p.leaf(1,k);p.helper(k)}
 func run(n int,x int8,k uint64)(int8,uint64,int16){p:=&S{x:x,y:uint64(x),z:int16(x)};p.outer(n,k);return p.x,p.y,p.z}
 `))
	methods := receiverFunctions(i)
	hits := receiverHits(methods)
	for _, p := range methods {
		if p.memoFlags&interpreterMemoReceiver == 0 {
			t.Fatal("separate graph proof", p.Fn)
		}
	}
	for _, x := range []int8{0, 7, -128, -1} {
		for _, k := range []uint64{0, 1, 7, 64, math.MaxUint64} {
			receiverDifferential(t, baseline, i, "run", 6, x, k)
		}
	}
	if *hits == 0 {
		t.Fatal("no separate graph hits")
	}
}

func TestInterpreterReceiverTransformFallback(t *testing.T) {
	for _, body := range []string{
		`p.v ^= k;p.v += k`, `p.v <<= k;p.v &= k`, `p.v = ^p.v;p.v = -p.v`,
		`p.v /= k`, `p.v %= k`, `p.v >>= k`, `p.v = k &^ p.v`,
		`p.v ^= p.v`, `p.v <<= uint(p.v)`, `p.v = ^p.v;return p.v`,
		`if p.v>0{p.v &= k}`, `old:=p.v;p.other=3;p.v=old^k`,
		`old:=p.v;p.g(k);p.v=old^k`, `p.v += k;p.g(k)`,
	} {
		t.Run(body, func(t *testing.T) {
			baseline, i := loadMainMemoPair(t, []byte(`package main
 type S struct{v int;other int}
 func(p *S) g(k int)int{p.v ^= k;return 0}
 func(p *S) f(n,k int)int{if n==0{`+body+`;return 0};return p.f(n-1,k)+p.f(n-1,k)}
 func run(n,x,k int)(int,int,int){p:=&S{v:x,other:2};a:=p.f(n,k);return a,p.v,p.other}`))
			methods := receiverFunctions(i)
			for _, p := range methods {
				if p.Fn.Name() != "f" {
					continue
				}
				if p.memoFlags&interpreterMemoReceiver == 0 {
					t.Fatal("unsafe transform", p.Fn)
				}
			}
			for _, x := range []int{0, 3, 127} {
				receiverDifferential(t, baseline, i, "run", 5, x, 2)
			}
		})
	}
}

func BenchmarkInterpreterReceiverOperators(b *testing.B) {
	for _, body := range []struct{ name, code string }{
		{"bitwise", `p.x ^= k;p.x &^= 8;p.x |= 2;p.y = ^p.y;p.y &= uint8(k);p.y ^= 37`},
		{"shift_negate", `p.x <<= uint64(k)&3;p.x = -p.x;p.x += 7;p.y = -p.y;p.y <<= uint64(k)&1;p.y -= 3`},
	} {
		baseline, i := loadMainMemoPair(b, []byte(`package main
 type S struct{x int64;y uint8}
 func(p *S) update(k int64){`+body.code+`}
 func(p *S) f(n int,k int64){p.update(k);if n>0{p.f(n-1,k);p.update(int64(n));p.f(n-1,k)}}
 func run(n int)(int64,uint8){p:=&S{x:7,y:129};p.f(n,-3);return p.x,p.y}`))
		for _, p := range receiverFunctions(i) {
			if p.memoFlags&interpreterMemoReceiver == 0 {
				b.Fatal("disabled")
			}
		}
		for pass, i := range []*Interp{baseline, i} {
			name := "off"
			if pass == 1 {
				name = "on"
			}
			b.Run(body.name+"/"+name, func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for j := 0; j < b.N; j++ {
					if _, err := i.RunFunc("run", 14); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}

func TestInterpreterReceiverExecutedShiftCounts(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type Count int16
 type S struct{x int8;y uint16;count Count}
 func(p *S) f(n int,s Count){if n==0{return};p.x <<= s+Count(n);p.x = -p.x;p.y <<= p.count+Count(n);p.f(n-1,s);p.x += int8(n);p.f(n-1,s)}
 func run(n int,x int8,s,count int16)(int8,uint16,Count){p:=&S{x:x,y:uint16(x),count:Count(count)};p.f(n,Count(s));return p.x,p.y,p.count}
 `))
	methods := receiverFunctions(i)
	hits := receiverHits(methods)
	for _, p := range methods {
		if p.memoFlags&interpreterMemoReceiver == 0 {
			t.Fatal("variable-count proof", p.Fn)
		}
	}
	for _, s := range []int16{0, 1, 7, 8, 63, 64, 255, 256} {
		receiverDifferential(t, baseline, i, "run", 5, int8(-127), s, s)
	}
	// An unexecuted negative count must not be evaluated at entry.
	receiverDifferential(t, baseline, i, "run", 0, int8(-127), int16(-100), int16(-100))
	if *hits == 0 {
		t.Fatal("no variable-count hits")
	}
}

func assertReceiverFullState(t *testing.T, p *function, _ int) {
	t.Helper()
	if p.memoFlags&interpreterMemoReceiver == 0 {
		t.Fatal("receiver full-state plan missing", p.Fn)
	}
}

func receiverFloatBits(t *testing.T, baseline, i *Interp, name string, args ...Value) Tuple {
	t.Helper()
	var want Tuple
	for pass, i := range []*Interp{baseline, i} {
		v, err := i.RunFunc(name, args...)
		if err != nil {
			t.Fatal(name, pass, err)
		}
		got := v.(Tuple)
		if pass == 0 {
			want = got
			continue
		}
		for j := range got {
			a, b := testValueBits(reflect.ValueOf(got[j])), testValueBits(reflect.ValueOf(want[j]))
			if a != b {
				t.Fatalf("%s %v output %d: %x != %x", name, args, j, a, b)
			}
		}
	}
	return want
}

func TestInterpreterReceiverFloatingFullState(t *testing.T) {
	for _, base := range []string{"float32", "float64"} {
		for _, defined := range []bool{false, true} {
			name := base
			decl := "type Number = " + base
			if defined {
				name += "Defined"
				decl = "type Number " + base
			}
			t.Run(name, func(t *testing.T) {
				src := `package main
 DECL
 type S struct{v Number}
 func(p *S) f(n int,x,y Number)(Number,Number){
  if n==0{return p.v,p.v}
  old:=p.v
  p.v=x;p.f(n-1,x,y)
  p.v=y;b,_:=p.f(n-1,x,y)
  p.v=x;p.f(n-1,x,y)
  return old,b
 }
 func run(n int,x,y BASE)(BASE,BASE,BASE){p:=&S{v:Number(x)};a,b:=p.f(n,Number(x),Number(y));return BASE(a),BASE(b),BASE(p.v)}
 `
				src = strings.ReplaceAll(strings.ReplaceAll(src, "DECL", decl), "BASE", base)
				baseline, i := loadMainMemoPair(t, []byte(src))
				methods := receiverFunctions(i)
				if len(methods) != 1 {
					t.Fatal(len(methods))
				}
				assertReceiverFullState(t, methods[0], 1)
				hits := receiverHits(methods)
				bits := []uint64{0, 1, 1 << 63, 0x7ff0000000000000, 0xfff0000000000000, 0x7ff8000000000123, 0xfff8000000000456, 0x7ff0000000000123}
				if base == "float32" {
					bits = []uint64{0, 1, 1 << 31, 0x7f800000, 0xff800000, 0x7fc00123, 0xffc00456, 0x7f800123}
				}
				for j, x := range bits {
					y := bits[(j+1)%len(bits)]
					var a, b Value = math.Float64frombits(x), math.Float64frombits(y)
					if base == "float32" {
						a, b = math.Float32frombits(uint32(x)), math.Float32frombits(uint32(y))
					}
					got := receiverFloatBits(t, baseline, i, "run", 4, a, b)
					for k, want := range []uint64{x, y, x} {
						// Ordinary interpreted float32 conversions may quiet signaling NaNs.
						if base == "float32" && want&0x7f800000 == 0x7f800000 && want&0x007fffff != 0 {
							want |= 0x00400000
						}
						if v := testValueBits(reflect.ValueOf(got[k])); v != want {
							t.Fatalf("identity bits %d: %x != %x", k, v, want)
						}
					}
				}
				if *hits == 0 {
					t.Fatal("no floating state hits")
				}
			})
		}
	}
}

func TestInterpreterReceiverFloatingArithmeticOrder(t *testing.T) {
	for _, base := range []string{"float32", "float64"} {
		t.Run(base, func(t *testing.T) {
			src := strings.ReplaceAll(`package main
 type Number BASE
 type S struct{a,b Number}
 func(p *S) f(n int,x,k Number)Number{
  if n==0{p.a+=k;p.a-=k;p.b*=k;p.b/=k;return p.a}
  p.a=x;p.b=x;a:=p.f(n-1,x,k)
  p.a=x;p.b=x;return a+p.f(n-1,x,k)
 }
 func run(n int,x,k BASE)(BASE,BASE,BASE){p:=&S{a:Number(x),b:Number(x)};a:=p.f(n,Number(x),Number(k));return BASE(a),BASE(p.a),BASE(p.b)}
 `, "BASE", base)
			baseline, i := loadMainMemoPair(t, []byte(src))
			methods := receiverFunctions(i)
			if len(methods) != 1 {
				t.Fatal(len(methods))
			}
			assertReceiverFullState(t, methods[0], 2)
			hits := receiverHits(methods)
			cases := [][2]float64{{1, 1 << 54}, {1.25, 3}, {0, 0}, {math.Copysign(0, -1), 2}, {math.Inf(1), 2}, {1, math.Inf(-1)}, {math.Float64frombits(0x7ff8000000000123), 1}, {math.SmallestNonzeroFloat64, 2}}
			for j, c := range cases {
				var x, k Value = c[0], c[1]
				if base == "float32" {
					x, k = float32(c[0]), float32(c[1])
					if j == len(cases)-1 {
						x = math.SmallestNonzeroFloat32
					}
				}
				got := receiverFloatBits(t, baseline, i, "run", 5, x, k)
				if j == 0 && testValueBits(reflect.ValueOf(got[1])) != 0 {
					t.Fatal("floating updates regrouped", got)
				}
			}
			if *hits == 0 {
				t.Fatal("no mutable floating hits")
			}
		})
	}
}

func TestInterpreterReceiverFloatingAliasAndNil(t *testing.T) {
	for _, base := range []string{"float32", "float64"} {
		t.Run(base, func(t *testing.T) {
			src := strings.ReplaceAll(`package main
 type Number BASE
 type S struct{v Number}
 func(p *S) f(n int,x *Number)Number{if n<0{return 17};if n==0{p.v+=*x;return p.v};return p.f(n-1,x)+p.f(n-1,x)}
 func run(n int,alias bool)(BASE,BASE){p:=&S{v:1.25};x:=Number(0.5);q:=&x;if alias{q=&p.v};a:=p.f(n,q);return BASE(a),BASE(p.v)}
 func nilok()Number{var p *S;return p.f(-1,nil)}
 func nilrecv()Number{var p *S;x:=Number(1);return p.f(0,&x)}
 func nilarg()Number{return (&S{}).f(0,nil)}
 `, "BASE", base)
			baseline, i := loadMainMemoPair(t, []byte(src))
			methods := receiverFunctions(i)
			if len(methods) != 1 {
				t.Fatal(len(methods))
			}
			assertReceiverFullState(t, methods[0], 1)
			for _, alias := range []bool{false, true} {
				receiverFloatBits(t, baseline, i, "run", 5, alias)
			}
			for pass, i := range []*Interp{baseline, i} {
				if _, err := i.RunFunc("nilok"); err != nil {
					t.Fatal(err)
				}
				for _, name := range []string{"nilrecv", "nilarg"} {
					for j := 0; j < 2; j++ {
						if _, err := i.RunFunc(name); err == nil {
							t.Fatal("missing panic", name, pass)
						}
					}
				}
			}
		})
	}
}

func TestInterpreterMemoCacheBound(t *testing.T) {
	i := loadMainMemo(t, []byte(`package main
func f(n int) int { if n == 0 { return 0 }; return f(n-1)+1 }
`))
	p := i.funcs[i.mainpkg.Func("f")]
	if p.memoFlags&interpreterMemoScalar == 0 {
		t.Fatal("memo disabled")
	}
	peak := 0
	for j, instr := range p.ssaInstrs {
		if _, ok := instr.(*ssa.Return); ok {
			ret := p.Instrs[j]
			p.Instrs[j] = func(fr *frame) {
				cache := (*interpreterMemo)(nil)
				if fr.memoState != nil {
					cache = fr.memoState.memo
				}
				ret(fr)
				if cache != nil {
					if len(cache.values) > interpreterMemoCacheLimit {
						t.Fatal("cache exceeded limit", len(cache.values))
					}
					peak = max(peak, len(cache.values))
				}
			}
		}
	}
	if got, err := i.RunFunc("f", interpreterMemoCacheLimit+1); err != nil || got != interpreterMemoCacheLimit+1 {
		t.Fatal(got, err)
	}
	if peak != interpreterMemoCacheLimit {
		t.Fatal("cache did not saturate", peak)
	}
}

func TestInterpreterAccumulatorCacheBound(t *testing.T) {
	i := loadMainMemo(t, []byte(`package main
func f(n int, p *int) { if n == 0 { *p += 1; return }; f(n-1, p) }
`))
	p := i.funcs[i.mainpkg.Func("f")]
	if p.memoFlags&interpreterMemoAccumulator == 0 {
		t.Fatal("accumulator memo disabled")
	}
	peak := 0
	for j, instr := range p.ssaInstrs {
		if _, ok := instr.(*ssa.Return); ok {
			ret := p.Instrs[j]
			p.Instrs[j] = func(fr *frame) {
				cache := (*interpreterAccumulatorMemo)(nil)
				if fr.memoState != nil {
					cache = fr.memoState.accumulator.memo
				}
				ret(fr)
				if cache != nil {
					if len(cache.values) > interpreterMemoCacheLimit {
						t.Fatal("cache exceeded limit", len(cache.values))
					}
					peak = max(peak, len(cache.values))
				}
			}
		}
	}
	out := 0
	if _, err := i.RunFunc("f", interpreterMemoCacheLimit+1, &out); err != nil || out != 1 {
		t.Fatal(out, err)
	}
	if peak != interpreterMemoCacheLimit {
		t.Fatal("cache did not saturate", peak)
	}
}

func TestInterpreterReceiverCacheBound(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type S struct{v float64}
 func(p *S) f(n int){if n==0{p.v+=1;return};p.f(n-1);p.f(n-1)}
 func run(n int)float64{p:=&S{};p.f(n);return p.v}
 `))
	methods := receiverFunctions(i)
	p := methods[0]
	if p.memoFlags&interpreterMemoReceiver == 0 {
		t.Fatal("method disabled")
	}
	assertReceiverFullState(t, p, 1)
	var caches []*interpreterReceiverMemo
	peak := 0
	for j, instr := range p.ssaInstrs {
		if _, ok := instr.(*ssa.Return); ok {
			ret := p.Instrs[j]
			p.Instrs[j] = func(fr *frame) {
				ret(fr)
				if fr.memoState != nil {
					if cache := fr.memoState.receiver.memo; cache != nil {
						if len(cache.values) > interpreterReceiverCacheLimit {
							t.Fatal("cache exceeded limit", len(cache.values))
						}
						if len(cache.values) > peak {
							peak = len(cache.values)
						}
						if fr.caller == nil || fr.caller.pfn != p {
							caches = append(caches, cache)
						}
					}
				}
			}
		}
	}
	for j := 0; j < 2; j++ {
		receiverDifferential(t, baseline, i, "run", 13)
	}
	if peak != interpreterReceiverCacheLimit {
		t.Fatal("cache did not saturate", peak)
	}
	if len(caches) != 2 || caches[0] == caches[1] {
		t.Fatal("cache not invocation-local", len(caches))
	}
	for _, cache := range caches {
		clear(cache.values)
	}
}

func TestInterpreterReceiverSaturatedTransforms(t *testing.T) {
	baseline, i := loadMainMemoPair(t, []byte(`package main
 type S struct{serial float64;affine uint64;bits uint8}
 func(p *S) helper(n int){p.affine*=3;p.affine+=uint64(n);p.bits^=uint8(n);p.bits|=2}
 func(p *S) f(n int){p.helper(n);if n==0{p.serial+=1;return};p.f(n-1);p.helper(n);p.f(n-1)}
 func run(n int)(float64,uint64,uint8){p:=&S{affine:7,bits:9};p.f(n);return p.serial,p.affine,p.bits}
 `))
	methods := receiverFunctions(i)
	for _, p := range methods {
		if p.memoFlags&interpreterMemoReceiver == 0 {
			t.Fatal("method disabled", p.Fn)
		}
	}
	receiverDifferential(t, baseline, i, "run", 13)
}

func TestGCGlobalChannelReceive(t *testing.T) {
	for _, commaOK := range []bool{false, true} {
		for _, closed := range []bool{false, true} {
			i, err := NewContext(ExperimentalSupportGC).LoadInterp("receive.go", `package main
func receive(ch chan int) int { x, _ := <-ch; return x }
`)
			if err != nil {
				t.Fatal(err)
			}
			// This small source has no runtime.GC call to enable safepoints.
			// Enable them explicitly to exercise the static-channel GC branches.
			i.gcEnabled = true
			pfn := i.funcs[i.mainpkg.Func("receive")]
			var recv *ssa.UnOp
			var pc int
			for j, instr := range pfn.ssaInstrs {
				if op, ok := instr.(*ssa.UnOp); ok {
					recv, pc = op, j
					break
				}
			}
			if recv == nil {
				t.Fatal("missing receive")
			}
			ch := make(chan int, 1)
			want := 0
			if closed {
				close(ch)
			} else {
				ch <- 42
				want = 42
			}
			// Native channels can be embedded directly in the static register table.
			reg := pfn.regIndex(recv.X)
			pfn.index[recv.X] = uint32(kindGlobal)<<30 | uint32(reflect.Chan)<<24 | uint32(reg)
			pfn.stack[reg] = ch
			recv.CommaOk = commaOK
			fr := &frame{pfn: pfn, block: recv.Block(), ipc: int32(pc + 1), stack: append([]value(nil), pfn.stack...)}
			makeUnOpARROW(pfn, recv)(fr)
			got := fr.reg(pfn.regIndex(recv))
			if commaOK {
				v := got.(tuple)
				if v[0] != want || v[1] != !closed {
					t.Fatalf("receive: %v", v)
				}
			} else if got != want {
				t.Fatalf("receive: %v, want %d", got, want)
			}
			i.UnsafeRelease()
		}
	}
}

func TestGCFramesSkipsOnlyEmptyRoots(t *testing.T) {
	cached := new(int)
	pfn := &function{cacheRegs: []register{0}}
	fr := &frame{pfn: pfn, block: &ssa.BasicBlock{}, stack: []value{cached}, phiVals: []value{cached}}
	fr.gcFrames()
	if fr.stack[0] != nil || fr.phiVals[0] != nil {
		t.Fatal("frame without SSA roots retained hidden references")
	}
	child := &frame{pfn: &function{}, block: &ssa.BasicBlock{}, caller: fr}
	fr.stack[0], fr.phiVals[0] = cached, cached
	child.gcFrames()
	if fr.stack[0] != nil || fr.phiVals[0] != nil {
		t.Fatal("empty child frame skipped cleanup of its caller")
	}
}
