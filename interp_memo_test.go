package ixgo

import (
	"math"
	"os"
	"reflect"
	"testing"
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
func TestInterpreterMemo(t *testing.T) {
	for _, name := range []string{"fibish", "fibish_closure", "fibish_float32", "fibish_float64"} {
		t.Run(name, func(t *testing.T) {
			src, err := os.ReadFile("testdata/fibish/" + name + ".go")
			if err != nil {
				t.Fatal(err)
			}
			baseline, i := loadMainMemoPair(t, src)
			p := i.funcs[i.mainpkg.Func("f")]
			if p.memoFlags&interpreterMemoScalar == 0 {
				t.Fatal("memo not enabled")
			}
			if name == "fibish_closure" { // Exercise the unchanged forwarded function argument via main.
				for _, i := range []*Interp{baseline, i} {
					if _, err := i.RunFunc("main"); err != nil {
						t.Fatal(err)
					}
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

				a, err := i.RunFunc("f", x)
				if err != nil {
					t.Fatal(err)
				}
				b, err := baseline.RunFunc("f", x)
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
func receiverFunctions(i *Interp) []*function {
	var out []*function
	for fn, p := range i.funcs {
		if fn.Package() == i.mainpkg && fn.Signature.Recv() != nil {
			out = append(out, p)
		}
	}
	return out
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
