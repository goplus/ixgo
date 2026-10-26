package integration_test

import (
	"testing"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/memo"
	"github.com/goplus/ixgo/mixedrunner"
	"golang.org/x/tools/go/ssa"
)

const modeSource = `package main
func mixed(n int, x float64) float64 { return float64(n) + x }
func fib(n int) int { if n < 2 { return n }; return fib(n-1) + fib(n-2) }
`

func TestOptimizerModes(t *testing.T) {
	for _, tc := range []struct {
		name      string
		mode      ixgo.Mode
		wantMixed bool
	}{
		{"default", 0, true},
		{"disable mixed", ixgo.DisableMixedRunner, false},
		{"disable memo", ixgo.DisableMemo, true},
		{"disable both", ixgo.DisableMemo | ixgo.DisableMixedRunner, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			selected := make(map[string]bool)
			ixgo.RegisterMixedRunner(func(fn *ssa.Function) ixgo.TypedRunner {
				selected[fn.Name()] = true
				return mixedrunner.MakeRunner(fn)
			})
			t.Cleanup(func() { ixgo.RegisterMixedRunner(mixedrunner.MakeRunner) })
			interp, err := ixgo.NewContext(tc.mode).LoadInterp("mode.go", []byte(modeSource))
			if err != nil {
				t.Fatal(err)
			}
			defer interp.UnsafeRelease()
			if selected["mixed"] != tc.wantMixed {
				t.Fatalf("mixed selected=%v, want %v", selected["mixed"], tc.wantMixed)
			}
			if got, err := interp.RunFunc("mixed", 3, 0.5); err != nil || got != 3.5 {
				t.Fatalf("mixed: %v %v", got, err)
			}
			if got, err := interp.RunFunc("fib", 8); err != nil || got != 21 {
				t.Fatalf("fib: %v %v", got, err)
			}
		})
	}
}
