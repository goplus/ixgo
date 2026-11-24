package ixgo

import (
	"fmt"
	"testing"
)

func TestClosureBackendInstructionDescriptions(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		t.Run(fmt.Sprintf("descriptions=%t", enabled), func(t *testing.T) {
			mode := DisableMemo
			if enabled {
				mode |= EnableInstructionDescriptions
			}
			ctx := NewContext(mode)
			interp, err := ctx.LoadInterp("closure_description.go", executionDescriptionSource)
			if err != nil {
				t.Fatal(err)
			}
			defer interp.UnsafeRelease()
			for _, p := range interp.funcs {
				if !enabled && p.Code != nil {
					t.Fatal("default closure backend retained typed descriptions")
				}
			}
			for _, name := range []string{"phi", "closure", "interfaceCall", "deferred", "recovered", "goroutine"} {
				if got, err := interp.RunFunc(name); err != nil || got != 42 {
					t.Fatalf("%s: got %v, %v; want 42", name, got, err)
				}
			}
			// Description storage is optional even after compilation.
			for _, p := range interp.funcs {
				p.Code = nil
			}
			if got, err := interp.RunFunc("phi"); err != nil || got != 42 {
				t.Fatalf("without descriptions: got %v, %v", got, err)
			}
		})
	}
}

// Compilation and warm-up are excluded so these samples measure the closure
// executor with and without retained typed instruction descriptions.
func BenchmarkClosureInstructionDescriptions(b *testing.B) {
	for _, enabled := range []bool{false, true} {
		b.Run(fmt.Sprintf("descriptions=%t", enabled), func(b *testing.B) {
			for _, name := range []string{"phi", "closure", "interfaceCall", "deferred", "recovered", "goroutine"} {
				b.Run(name, func(b *testing.B) {
					mode := DisableMemo
					if enabled {
						mode |= EnableInstructionDescriptions
					}
					ctx := NewContext(mode)
					ctx.SetLeastCallForEnablePool(0)
					interp, err := ctx.LoadInterp("closure_bench.go", executionDescriptionSource)
					if err != nil {
						b.Fatal(err)
					}
					b.Cleanup(interp.UnsafeRelease)
					if got, err := interp.RunFunc(name); err != nil || got != 42 {
						b.Fatalf("got %v, %v", got, err)
					}
					b.ReportAllocs()
					b.ResetTimer()
					for b.Loop() {
						if got, err := interp.RunFunc(name); err != nil || got != 42 {
							b.Fatalf("got %v, %v", got, err)
						}
					}
				})
			}
		})
	}
}

func TestInstructionDescriptionModeIndependence(t *testing.T) {
	for _, descriptions := range []bool{false, true} {
		for _, plan := range []bool{false, true} {
			t.Run(fmt.Sprintf("descriptions=%t/plan=%t", descriptions, plan), func(t *testing.T) {
				mode := DisableMemo | ExperimentalSupportGC
				if descriptions {
					mode |= EnableInstructionDescriptions
				}
				if plan {
					mode |= EnableGCExecutionPlan
				}
				source := "package main\nimport \"runtime\"\nfunc work(n int) *int { runtime.GC(); p := new(int); *p = n\n"
				for i := 0; i < 1100; i++ {
					source += "*p += 1\n"
				}
				source += "runtime.GC(); return p }"
				interp, err := NewContext(mode).LoadInterp("mode_independence.go", source)
				if err != nil {
					t.Fatal(err)
				}
				defer interp.UnsafeRelease()
				p := interp.funcs[interp.mainpkg.Func("work")]
				if got := len(p.Code) != 0; got != descriptions {
					t.Fatalf("descriptions = %t, want %t", got, descriptions)
				}
				if got := p.gcPlan != nil; got != plan {
					t.Fatalf("GC plan = %t, want %t", got, plan)
				}
				if got, err := interp.RunFunc("work", 2); err != nil {
					t.Fatalf("work = %v, %v", got, err)
				} else if result, ok := got.(*int); !ok || *result != 1102 {
					t.Fatalf("unexpected work result: %v", got)
				}
			})
		}
	}
}
