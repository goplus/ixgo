package ixgo

import "testing"

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

func TestTypedMemoFloatFallback(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("memo.go", `package main
func f(x float64) float64 { if x<1 {return x};return f(x-1)+f(x-2) }
`)
	if err != nil {
		t.Fatal(err)
	}
	r := MakeTypedRunnerOf[float64](pkg.Func("f")).(*typedRunner[float64])
	if r.memoizable {
		t.Fatal("float memoization must preserve signed zero through fallback")
	}
}
