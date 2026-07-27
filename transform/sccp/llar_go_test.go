// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sccp

import (
	"bytes"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// These cases adapt SCCP-active programs from the Go compiler tests. in.go is
// the source program and out.txt is the complete package SSA after SCCP.
func TestSCCPGoGolden(t *testing.T) {
	entries, err := os.ReadDir("testdata/sccp")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		t.Run(entry.Name(), func(t *testing.T) {
			dir := filepath.Join("testdata/sccp", entry.Name())
			fn := buildSCCPFixture(t, dir)
			runSCCPPassForTest(t, fn.Pkg)
			for candidate := range ssautil.AllFunctions(fn.Prog) {
				if candidate.Pkg != fn.Pkg || candidate.Blocks == nil {
					continue
				}
				assertRewrittenSSA(t, candidate)
				assertSSAUsesDominated(t, candidate)
			}

			got := packageSSA(fn.Pkg)
			outPath := filepath.Join(dir, "out.txt")
			if os.Getenv("LLAR_UPDATE_SCCP_GOLDEN") == "1" {
				if err := os.WriteFile(outPath, got, 0666); err != nil {
					t.Fatal(err)
				}
			}
			want, err := os.ReadFile(outPath)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, want) {
				t.Fatalf("rewritten SSA does not match out.txt\n\nwant:\n%s\ngot:\n%s", want, got)
			}
		})
	}
}

func TestSCCPGoExecution(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "closure-slice", want: 3},
		{name: "empty-range"},
		{name: "interfaces"},
		{name: "switch-fallthrough"},
		{name: "variadic"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dir := filepath.Join("testdata/sccp", test.name)
			original := buildSCCPFixture(t, dir)
			if got := runSSAInt(t, original.Pkg); got != test.want {
				t.Fatalf("original target() = %d, want %d", got, test.want)
			}
			rewritten := buildSCCPFixture(t, dir)
			runSCCPPassForTest(t, rewritten.Pkg)
			if got := runSSAInt(t, rewritten.Pkg); got != test.want {
				t.Fatalf("rewritten target() = %d, want %d", got, test.want)
			}
		})
	}
}

func buildSCCPFixture(t *testing.T, dir string) *ssa.Function {
	t.Helper()
	source, err := os.ReadFile(filepath.Join(dir, "in.go"))
	if err != nil {
		t.Fatal(err)
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "in.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	pkg, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()},
		fset,
		types.NewPackage("example.com/test", "test"),
		[]*ast.File{file},
		ssa.SanityCheckFunctions,
	)
	if err != nil {
		t.Fatal(err)
	}
	fn := pkg.Func("target")
	if fn == nil {
		t.Fatal("target function not found")
	}
	return fn
}

func packageSSA(pkg *ssa.Package) []byte {
	var functions []*ssa.Function
	for fn := range ssautil.AllFunctions(pkg.Prog) {
		if fn.Pkg == pkg && fn.Blocks != nil && fn.Synthetic == "" {
			functions = append(functions, fn)
		}
	}
	sort.Slice(functions, func(i, j int) bool {
		return functions[i].String() < functions[j].String()
	})
	var out bytes.Buffer
	for _, fn := range functions {
		fn.WriteTo(&out)
	}
	return out.Bytes()
}
