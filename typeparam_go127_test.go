//go:build go1.27
// +build go1.27

package ixgo_test

import (
	"testing"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/bytes"
)

func TestConcreteMethods(t *testing.T) {
	src := `package main

type G[P any] struct {
	x P
}

func (g G[P]) M[Q any](q Q) (P, Q) {
	return g.x, q
}

func (g *G[P]) N[Q any](q Q) (P, Q) {
	return g.x, q
}

type N struct{}

func (N) M[P any](p P) P {
	return p
}

func (*N) N[P any](p P) P {
	return p
}

func main() {
	g := G[int]{x: 42}
	if p, q := g.M[bool](true); p != 42 || !q {
		panic("generic value method")
	}
	if p, q := g.M("inferred"); p != 42 || q != "inferred" {
		panic("inferred generic value method")
	}
	if p, q := g.N[bool](true); p != 42 || !q {
		panic("generic pointer method")
	}
	if p, q := g.N("inferred"); p != 42 || q != "inferred" {
		panic("inferred generic pointer method")
	}

	valueMethod := g.M[string]
	if p, q := valueMethod("method value"); p != 42 || q != "method value" {
		panic("generic value method value")
	}
	var inferredValueMethod func(bool) (int, bool) = g.M
	if p, q := inferredValueMethod(true); p != 42 || !q {
		panic("inferred generic value method value")
	}
	pointerMethod := g.N[string]
	if p, q := pointerMethod("method value"); p != 42 || q != "method value" {
		panic("generic pointer method value")
	}
	var inferredPointerMethod func(bool) (int, bool) = g.N
	if p, q := inferredPointerMethod(true); p != 42 || !q {
		panic("inferred generic pointer method value")
	}

	valueExpr := G[int].M[string]
	if p, q := valueExpr(g, "method expression"); p != 42 || q != "method expression" {
		panic("generic value method expression")
	}
	var inferredValueExpr func(G[int], bool) (int, bool) = G[int].M
	if p, q := inferredValueExpr(g, true); p != 42 || !q {
		panic("inferred generic value method expression")
	}
	pointerExpr := (*G[int]).N[string]
	if p, q := pointerExpr(&g, "method expression"); p != 42 || q != "method expression" {
		panic("generic pointer method expression")
	}
	var inferredPointerExpr func(*G[int], bool) (int, bool) = (*G[int]).N
	if p, q := inferredPointerExpr(&g, true); p != 42 || !q {
		panic("inferred generic pointer method expression")
	}

	n := N{}
	if n.M(42) != 42 || n.N("method") != "method" {
		panic("generic methods on a non-generic type")
	}
}
`
	if _, err := ixgo.RunFile("main.go", src, nil, 0); err != nil {
		t.Fatal(err)
	}
}
