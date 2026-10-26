package integration_test

import (
	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/memo"
	"github.com/goplus/ixgo/mixedrunner"
	_ "github.com/goplus/ixgo/mixedrunner"
	"golang.org/x/tools/go/ssa"
	"testing"
)

func TestCombinedPlugins(t *testing.T) {
	selected := make(map[string]bool)
	ixgo.RegisterMixedRunner(func(fn *ssa.Function) ixgo.TypedRunner {
		selected[fn.Name()] = true
		return mixedrunner.MakeRunner(fn)
	})
	t.Cleanup(func() { ixgo.RegisterMixedRunner(mixedrunner.MakeRunner) })
	i, err := ixgo.NewContext(0).LoadInterp("combined.go", `package main
func plain(n int) int {return n+1}
func mixed(n int,x float64) float64 {return float64(n)+x}
func fib(n int) int {if n<2{return n};return fib(n-1)+fib(n-2)}
func wrapper(n int,x float64) float64 {return float64(fib(n))+x}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	if !selected["plain"] || selected["fib"] || selected["wrapper"] || !selected["mixed"] {
		t.Fatalf("unexpected mixed selection: %v", selected)
	}
	for _, tc := range []struct {
		name string
		args []ixgo.Value
		want ixgo.Value
	}{
		{"plain", []ixgo.Value{5}, 6}, {"mixed", []ixgo.Value{5, 0.5}, 5.5},
		{"fib", []ixgo.Value{10}, 55}, {"wrapper", []ixgo.Value{10, 0.5}, 55.5},
	} {
		if got, err := i.RunFunc(tc.name, tc.args...); err != nil || got != tc.want {
			t.Fatalf("%s: %v %v", tc.name, got, err)
		}
	}
}

type staticMemoOrderRunner struct {
	ixgo.TypedArgumentsRunner
	calls *int
}

func (r staticMemoOrderRunner) RunArgs(args []ixgo.Value) ixgo.Value {
	*r.calls++
	return r.TypedArgumentsRunner.RunArgs(args)
}

func TestMixedStaticCallsAfterMemoInstallation(t *testing.T) {
	calls := 0
	ixgo.RegisterMixedRunner(func(fn *ssa.Function) ixgo.TypedRunner {
		runner := mixedrunner.MakeRunner(fn)
		if runner == nil {
			return nil
		}
		if fn.Name() == "leaf" || fn.Name() == "pair" {
			return staticMemoOrderRunner{runner.(ixgo.TypedArgumentsRunner), &calls}
		}
		return runner
	})
	t.Cleanup(func() { ixgo.RegisterMixedRunner(mixedrunner.MakeRunner) })
	ctx := ixgo.NewContext(ixgo.DisableMemo)
	pkg, err := ctx.LoadFile("static_memo_order.go", `package main
var count int
func leaf(n int) int { return n+1 }
func pair(n int) (int, int) { return n+2, n+3 }
func caller(n int) int { count++; x := leaf(n); a,b := pair(n); return x+a+b }
`)
	if err != nil {
		t.Fatal(err)
	}
	// Keep the callees reachable only through caller to force deferred installation.
	delete(pkg.Members, "leaf")
	delete(pkg.Members, "pair")
	i, err := ixgo.NewInterp(ctx, pkg)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	got, err := i.RunFunc("caller", 4)
	if err != nil || got != 18 {
		t.Fatalf("caller(4) = %v, %v; want 18", got, err)
	}
	if calls != 2 {
		t.Fatalf("real mixed static calls = %d, want 2", calls)
	}
}
