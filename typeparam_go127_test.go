//go:build go1.27
// +build go1.27

package ixgo_test

import (
	"testing"

	"github.com/goplus/ixgo"
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

func TestPromotedFieldsInLiteral(t *testing.T) {
	src := `package main

type fields struct {
	count int
	name  string
}

type record struct {
	fields
	active bool
}

type genericFields[T any] struct {
	value T
}

type genericRecord[T any] struct {
	genericFields[T]
	name string
}

type object struct {
	name  string
	color string
}

type point3D struct {
	object
	x, y, z float64
}

type line struct {
	point3D
	q point3D
}

func main() {
	r := record{count: 42, name: "answer", active: true}
	if r.count != 42 || r.name != "answer" || !r.active {
		panic("promoted field literal")
	}

	gr := genericRecord[int]{value: 7, name: "generic"}
	if gr.value != 7 || gr.name != "generic" {
		panic("promoted generic field literal")
	}

	// Promoted fields may be selected through more than one embedded value.
	l := line{name: "diagonal", y: -4, z: 12.3}
	if l.name != "diagonal" || l.y != -4 || l.z != 12.3 {
		panic("nested promoted field literal")
	}
}
`
	if _, err := ixgo.RunFile("main.go", src, nil, 0); err != nil {
		t.Fatal(err)
	}
}

func TestGenericFunctionAssignmentInference(t *testing.T) {
	src := `package main

import "fmt"

func identity[T any](v T) T {
	return v
}

func pair[A, B any](a A, b B) (A, B) {
	return a, b
}

func apply(fn func(int) int, v int) int {
	return fn(v)
}

func makeIdentity() func(string) string {
	return identity
}

func genericFormatter[T any](v T) string {
	return fmt.Sprintf("value: %v", v)
}

type intFormatter func(int) string

func main() {
	var intIdentity func(int) int = identity
	if intIdentity(42) != 42 {
		panic("single type argument assignment")
	}

	var pairFn func(int, string) (int, string) = pair
	if n, s := pairFn(7, "seven"); n != 7 || s != "seven" {
		panic("multiple type argument assignment")
	}

	var pairFromInt func(int, string) (int, string) = pair[int]
	if n, s := pairFromInt(8, "eight"); n != 8 || s != "eight" {
		panic("partial type argument assignment")
	}

	if apply(identity, 9) != 9 {
		panic("function argument inference")
	}

	stringIdentity := makeIdentity()
	if stringIdentity("inferred") != "inferred" {
		panic("return assignment inference")
	}

	formatters := []intFormatter{genericFormatter}
	if formatters[0](12) != "value: 12" {
		panic("composite literal inference")
	}

	fn := intFormatter(genericFormatter)
	if fn(13) != "value: 13" {
		panic("function conversion inference")
	}

	ch := make(chan intFormatter, 1)
	ch <- genericFormatter
	if got := (<-ch)(14); got != "value: 14" {
		panic("channel send inference")
	}
}
`
	if _, err := ixgo.RunFile("main.go", src, nil, 0); err != nil {
		t.Fatal(err)
	}
}
