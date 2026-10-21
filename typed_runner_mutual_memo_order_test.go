package ixgo

import (
	"testing"

	"golang.org/x/tools/go/ssa"
)

func TestMixedRunnerIncompleteTraversal(t *testing.T) {
	oldMemo, oldMixed := memoPlugin, mixedRunnerFactory
	t.Cleanup(func() {
		RegisterMemoPlugin(oldMemo)
		RegisterMixedRunner(oldMixed)
	})
	RegisterMemoPlugin(nil)
	RegisterMixedRunner(nil)
	i, err := NewContext(0).LoadInterp("incomplete_memo_order.go", `package main
func A(n int) int { return B(n)+1 }
func B(n int) int { return n+1 }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	i.funcs = make(map[*ssa.Function]*function)
	RegisterMemoPlugin(func(p MemoInstaller, fn *ssa.Function) {
		if fn.Name() == "A" {
			panic("incomplete traversal")
		}
	})
	factoryCalls := 0
	RegisterMixedRunner(func(fn *ssa.Function) TypedRunner {
		factoryCalls++
		return testScalarRunner{}
	})
	v := visitor{intp: i, prog: i.mainpkg.Prog, pkgs: map[*ssa.Package]bool{i.mainpkg: true}, seen: make(map[*ssa.Function]bool), base: fnBase}
	func() {
		defer func() {
			if got := recover(); got != "incomplete traversal" {
				t.Fatalf("panic = %v", got)
			}
		}()
		v.function(i.mainpkg.Func("A"))
	}()
	if factoryCalls != 0 || len(v.mixedPending) != 0 || len(v.mixedCalls) != 0 {
		t.Fatalf("incomplete traversal installed %d runners, pending %d runners and %d calls", factoryCalls, len(v.mixedPending), len(v.mixedCalls))
	}
}

func TestMixedRunnerMutualMemoInstallationOrder(t *testing.T) {
	oldMemo, oldMixed := memoPlugin, mixedRunnerFactory
	t.Cleanup(func() {
		RegisterMemoPlugin(oldMemo)
		RegisterMixedRunner(oldMixed)
	})
	for _, memoName := range []string{"A", "B", ""} {
		t.Run("memo="+memoName, func(t *testing.T) {
			ctx := NewContext(0)
			pkg, err := ctx.LoadFile("mutual_memo_order.go", `package main
func A(n int) int { if n == 0 { return 0 }; return B(n-1)+1 }
func B(n int) int { if n == 0 { return 0 }; return A(n-1)+1 }
`)
			if err != nil {
				t.Fatal(err)
			}
			RegisterMemoPlugin(nil)
			RegisterMixedRunner(nil)
			i, err := NewInterp(ctx, pkg)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			// Recompile from A so map iteration cannot choose the traversal order.
			i.funcs = make(map[*ssa.Function]*function)
			var installed []string
			RegisterMemoPlugin(func(p MemoInstaller, fn *ssa.Function) {
				installed = append(installed, fn.Name())
				if fn.Name() == memoName {
					p.InstallScalarMemo()
				}
			})
			factoryCalls := 0
			RegisterMixedRunner(func(fn *ssa.Function) TypedRunner {
				factoryCalls++
				return testScalarRunner{}
			})
			v := visitor{intp: i, prog: pkg.Prog, pkgs: map[*ssa.Package]bool{pkg: true}, seen: make(map[*ssa.Function]bool), base: fnBase}
			v.function(pkg.Func("A"))
			if len(installed) != 2 || installed[0] != "B" || installed[1] != "A" {
				t.Fatalf("memo installation order: %v, want [B A]", installed)
			}
			wantCalls := 0
			if memoName == "" {
				wantCalls = 2
			}
			if factoryCalls != wantCalls {
				t.Errorf("mixed factory calls = %d, want %d", factoryCalls, wantCalls)
			}
			for _, name := range []string{"A", "B"} {
				p := i.funcs[pkg.Func(name)]
				if (p.memoFlags != 0) != (name == memoName) {
					t.Errorf("%s memo flags = %v", name, p.memoFlags)
				}
				if (p.mixedRunner != nil) != (memoName == "") {
					t.Errorf("%s mixed runner installed = %v", name, p.mixedRunner != nil)
				}
			}
			if memoName != "" {
				got, err := i.RunFunc("A", 8)
				if err != nil || got != 8 {
					t.Fatalf("A(8) = %v, %v; want 8", got, err)
				}
			}
		})
	}
}
