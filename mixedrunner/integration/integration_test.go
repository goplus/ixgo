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
