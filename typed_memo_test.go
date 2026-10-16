package ixgo

import (
	"math"
	"os"
	"strings"
	"testing"
)

func TestTypedMemoIntegerOverflowTable(t *testing.T) {
	m := newTypedMemo[int]()
	for i := 0; i < 100; i++ {
		m.store(i, typedResults[int]{values: [maxTypedResults]int{i + 1}, n: 1})
	}
	for i := 0; i < 100; i++ {
		v, ok := m.lookup(i)
		if !ok || v.values[0] != i+1 {
			t.Fatal(i, v, ok)
		}
	}
	if _, ok := m.lookup(101); ok {
		t.Fatal("unexpected cache hit")
	}
}

func TestTypedMemoFloatBits(t *testing.T) {
	checkFloatMemoBits[float32](t)
	checkFloatMemoBits[float64](t)
}
func checkFloatMemoBits[T TypedScalar](t *testing.T) {
	m := newTypedMemo[T]()
	// Put signed zeros and distinct NaNs in both tiers.
	for tier := 0; tier < 2; tier++ {
		m = newTypedMemo[T]()
		if tier == 1 {
			for i := 1; i <= smallTypedMemoSize; i++ {
				m.store(T(i), typedResults[T]{n: 1})
			}
		}
		values := []T{T(0), T(math.Copysign(0, -1)), T(math.Inf(1)), T(math.Inf(-1))}
		if m.kind.String() == "float32" {
			values = append(values, T(math.Float32frombits(0x7fc00001)), T(math.Float32frombits(0x7fc00002)))
		} else {
			values = append(values, T(math.Float64frombits(0x7ff8000000000001)), T(math.Float64frombits(0x7ff8000000000002)))
		}
		if m.key(values[0]) == m.key(values[1]) || m.key(values[4]) == m.key(values[5]) {
			t.Fatal("distinct bit patterns conflated")
		}
		for _, x := range values {
			m.store(x, typedResults[T]{values: [maxTypedResults]T{x}, n: 1})
		}
		for _, x := range values {
			v, ok := m.lookup(x)
			if !ok || m.key(x) != m.key(v.values[0]) {
				t.Fatalf("tier %d key %x: %v %v", tier, m.key(x), v, ok)
			}
		}
	}
}

func TestTypedMemoSignedZeroRecursion(t *testing.T) {
	for _, kind := range []string{"float32", "float64"} {
		pkg, err := NewContext(0).LoadFile("zero.go", strings.ReplaceAll(`package main
func f(x FLOAT) FLOAT { if x==0 {return x}; z:=x-x;return 1/f(z)+1/f(-z) }
`, "FLOAT", kind))
		if err != nil {
			t.Fatal(err)
		}
		if kind == "float32" {
			r := MakeTypedRunnerOf[float32](pkg.Func("f")).(*typedRunner[float32])
			if !math.IsNaN(float64(r.Run(float32(1)).(float32))) {
				t.Fatal("signed zero conflated")
			}
		} else {
			r := MakeTypedRunnerOf[float64](pkg.Func("f")).(*typedRunner[float64])
			if !math.IsNaN(r.Run(float64(1)).(float64)) {
				t.Fatal("signed zero conflated")
			}
		}
	}
}

func TestTypedMemoFloatRecursion(t *testing.T) {
	checkFloatRecursion[float32](t, "float32")
	checkFloatRecursion[float64](t, "float64")
}
func checkFloatRecursion[T TypedScalar](t *testing.T, kind string) {
	src, err := os.ReadFile("testdata/fibish/fibish_" + kind + ".go")
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := NewContext(0).LoadFile("fib.go", src)
	if err != nil {
		t.Fatal(err)
	}
	r := MakeTypedRunnerOf[T](pkg.Func("f")).(*typedRunner[T])
	for n := 0; n <= 20; n++ {
		a := r.runSelectedTyped(T(n))
		b := r.runRecursiveTyped(T(n))
		m := newTypedMemo[T]()
		for i := 0; i < int(a.n); i++ {
			if m.key(a.values[i]) != m.key(b.values[i]) {
				t.Fatal(kind, n, a, b)
			}
		}
	}
}

func BenchmarkTypedMemoFloat(b *testing.B) {
	benchmarkFloatMemo[float32](b, "float32")
	benchmarkFloatMemo[float64](b, "float64")
}
func benchmarkFloatMemo[T TypedScalar](b *testing.B, kind string) {
	src, err := os.ReadFile("testdata/fibish/fibish_" + kind + ".go")
	if err != nil {
		b.Fatal(err)
	}
	pkg, err := NewContext(0).LoadFile("fib.go", src)
	if err != nil {
		b.Fatal(err)
	}
	r := MakeTypedRunnerOf[T](pkg.Func("f")).(*typedRunner[T])
	for _, memo := range []bool{false, true} {
		name := "baseline"
		if memo {
			name = "memo"
		}
		b.Run(kind+"/"+name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				var v typedResults[T]
				if memo {
					v = r.runSelectedTyped(T(20))
				} else {
					v = r.runRecursiveTyped(T(20))
				}
				if v.n != 2 {
					b.Fatal(v)
				}
			}
		})
	}
}
