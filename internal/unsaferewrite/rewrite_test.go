package unsaferewrite

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"strings"
	"testing"
)

func TestExternalLayoutPaths(t *testing.T) {
	const src = `package main

import "unsafe"

func ordinary() {}
func layout[T any](x T) (uintptr, uintptr) {
	ordinary()
	return unsafe.Sizeof(x), unsafe.Alignof(x)
}
`
	f, err := parser.ParseFile(token.NewFileSet(), "paths.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	c := Isolate(f)
	if c == f || c.Decls[1] != f.Decls[1] || c.Imports[0] != f.Imports[0] {
		t.Fatal("expected a private root with shared ordinary declaration and import")
	}
	old := f.Decls[2].(*ast.FuncDecl)
	newFn := c.Decls[2].(*ast.FuncDecl)
	if old == newFn || old.Body == newFn.Body || old.Body.List[0] != newFn.Body.List[0] {
		t.Fatal("unexpected function path sharing")
	}
	before := old.Body.List[1].(*ast.ReturnStmt)
	after := newFn.Body.List[1].(*ast.ReturnStmt)
	for i, expr := range before.Results {
		a := expr.(*ast.CallExpr)
		b := after.Results[i].(*ast.CallExpr)
		if a == b || a.Args[0] != b.Args[0] {
			t.Fatalf("result %d: expected private call with shared operand", i)
		}
	}
}

func BenchmarkIsolate(b *testing.B) {
	for _, count := range []int{0, 1, 100} {
		var src strings.Builder
		src.WriteString("package main\nimport \"unsafe\"\nvar _ unsafe.Pointer\n")
		for i := 0; i < 100; i++ {
			if i < count {
				fmt.Fprintf(&src, "func f%d[T any](x T) uintptr {return unsafe.Sizeof(x)}\n", i)
			} else {
				fmt.Fprintf(&src, "func f%d[T any](x T) uintptr {return 0}\n", i)
			}
		}
		file, err := parser.ParseFile(token.NewFileSet(), "paths.go", src.String(), 0)
		if err != nil {
			b.Fatal(err)
		}
		b.Run(fmt.Sprint(count), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				Isolate(file)
			}
		})
	}
}

func TestRewrite(t *testing.T) {
	const source = `package p

import u "unsafe"

func layout[T any](v T) uintptr {
	return u.Sizeof(v) + u.Alignof(v) + u.Offsetof(struct{ Field T }{}.Field)
}

var size = u.Sizeof(int64(0))
`
	for _, dot := range []bool{false, true} {
		t.Run(fmt.Sprint(dot), func(t *testing.T) {
			src := source
			if dot {
				src = strings.ReplaceAll(strings.Replace(src, "import u ", "import . ", 1), "u.", "")
			}
			fs := token.NewFileSet()
			original, err := parser.ParseFile(fs, "layout.go", src, parser.ParseComments)
			if err != nil {
				t.Fatal(err)
			}
			printFile := func(f *ast.File) string {
				var buf bytes.Buffer
				if err := printer.Fprint(&buf, fs, f); err != nil {
					t.Fatal(err)
				}
				return buf.String()
			}
			before := printFile(original)
			file := Isolate(original)
			info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue), Uses: make(map[*ast.Ident]types.Object), Defs: make(map[*ast.Ident]types.Object), Selections: make(map[*ast.SelectorExpr]*types.Selection)}
			conf := types.Config{Importer: importer.Default()}
			if _, err := conf.Check("p", fs, []*ast.File{file}, info); err != nil {
				t.Fatal(err)
			}
			Rewrite([]*ast.File{file}, info)
			count := 0
			for n := range ast.Preorder(file) {
				call, ok := n.(*ast.CallExpr)
				if !ok || info.Types[call].Value != nil {
					continue
				}
				count++
				if len(call.Args) != 2 {
					t.Fatalf("args = %d", len(call.Args))
				}
				if _, ok := info.TypeOf(call.Args[0]).(*types.Pointer); !ok {
					t.Fatalf("carrier type = %v", info.TypeOf(call.Args[0]))
				}
				if info.Types[call.Args[1]].Value == nil {
					t.Fatal("missing metadata constant")
				}
			}
			if count != 3 {
				t.Fatalf("rewritten calls = %d", count)
			}
			rewritten := printFile(file)
			Rewrite([]*ast.File{file}, info)
			if printFile(file) != rewritten {
				t.Fatal("rewrite is not idempotent")
			}
			if printFile(original) != before {
				t.Fatal("caller syntax changed")
			}
		})
	}
}
