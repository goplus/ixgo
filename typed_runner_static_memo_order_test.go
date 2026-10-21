package ixgo_test

import (
	"testing"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/mixedrunner"
	"golang.org/x/tools/go/ssa"
)

// Importing mixedrunner must not enable it for unrelated interpreter tests.
func init() { ixgo.RegisterMixedRunner(nil) }

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
	t.Cleanup(func() { ixgo.RegisterMixedRunner(nil) })
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
