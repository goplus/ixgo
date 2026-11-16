package ixgo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"github.com/goplus/ixgo/internal/typesutil"
)

var genericMethodResult bool

func TestIsGenericMethod(t *testing.T) {
	src := `package methods

type Base struct{}
func (Base) Value() {}
func (*Base) Pointer() {}
type Outer struct { Base }
type PointerOuter struct { *Base }
type Generic[T any] struct{}
func (Generic[T]) Value() T { var zero T; return zero }
func (*Generic[T]) Pointer() {}

func selections(b Base, o Outer, p PointerOuter, g Generic[int]) {
	_ = b.Value
	_ = b.Pointer
	_ = o.Value
	_ = o.Pointer
	_ = p.Value
	_ = p.Pointer
	_ = g.Value
	_ = g.Pointer
	_ = Base.Value
	_ = (*Base).Pointer
	_ = Outer.Value
	_ = (*Outer).Pointer
	_ = PointerOuter.Value
	_ = PointerOuter.Pointer
	_ = Generic[int].Value
	_ = (*Generic[int]).Pointer
}
`
	if typesutil.SupportsGenericMethods {
		src += `
func (Base) Method[T any](v T) T { return v }
func (*Base) PointerMethod[T any](v T) T { return v }
func (Generic[T]) Method[U any](v U) U { return v }
func (*Generic[T]) PointerMethod[U any](v U) U { return v }
`
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "methods.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	info := &types.Info{Selections: make(map[*ast.SelectorExpr]*types.Selection)}
	pkg, err := new(types.Config).Check("methods", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}

	check := func(name string, sel *types.Selection, want bool) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			if got := isGenericMethod(sel); got != want {
				t.Fatalf("isGenericMethod(%s) = %v, want %v", sel, got, want)
			}
			if allocs := testing.AllocsPerRun(100, func() {
				genericMethodResult = isGenericMethod(sel)
			}); allocs != 0 {
				t.Fatalf("isGenericMethod allocated %g times per call", allocs)
			}
		})
	}
	for _, name := range []string{"Base", "Outer", "PointerOuter", "Generic"} {
		typ := pkg.Scope().Lookup(name).Type()
		if name == "Generic" {
			typ, err = types.Instantiate(nil, typ, []types.Type{types.Typ[types.Int]}, true)
			if err != nil {
				t.Fatal(err)
			}
		}
		for _, recv := range []types.Type{typ, types.NewPointer(typ)} {
			set := types.NewMethodSet(recv)
			for i := 0; i < set.Len(); i++ {
				sel := set.At(i)
				want := sel.Obj().Name() == "Method" || sel.Obj().Name() == "PointerMethod"
				check(recv.String()+"/"+sel.Obj().Name(), sel, want)
			}
		}
	}
	for expr, sel := range info.Selections {
		check(fset.Position(expr.Pos()).String(), sel, false)
	}
}
