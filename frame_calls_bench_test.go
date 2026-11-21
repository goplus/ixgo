package ixgo

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func BenchmarkFrameCalls(b *testing.B) {
	for _, tc := range []struct {
		name string
		args []Value
		want Value
	}{
		{"empty", nil, nil},
		{"add", []Value{41}, 42},
		{"arithmetic", []Value{1000}, 499500},
		{"pairLoop", []Value{1000}, 1000000},
		{"pair2Loop", []Value{1000}, 999000},
		{"fibish", []Value{20}, Tuple{2584, 8362}},
	} {
		b.Run(tc.name, func(b *testing.B) {
			ctx := NewContext(0)
			ctx.SetLeastCallForEnablePool(0)
			interp, err := ctx.LoadInterp("frame_bench.go", frameCallsSource)
			if err != nil {
				b.Fatal(err)
			}
			b.Cleanup(interp.UnsafeRelease)
			if got, err := interp.RunFunc(tc.name, tc.args...); err != nil || !reflect.DeepEqual(got, tc.want) {
				b.Fatalf("got %v, %v; want %v", got, err, tc.want)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				if _, err := interp.RunFunc(tc.name, tc.args...); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

const frameCallsSource = `package main
func empty() {}
func add(n int) int { return n + 1 }
func pair(n int) (int, int) { return n, n + 1 }
func arithmetic(n int) int {
	s := 0
	for i := 0; i < n; i++ { s += i }
	return s
}
func pairLoop(n int) int {
	s := 0
	for i := 0; i < n; i++ {
		a, b := pair(i)
		s += a + b
	}
	return s
}
func pair2Loop(n int) int {
	s := 0
	for i := 0; i < n; i++ {
		a, b := pair2(i, i + 1)
		s += a + b
	}
	return s
}
func pair2(a, b int) (int, int) { return a + b, a - b }
func fibish(x int) (int, int) {
	if x < 3 { return 0, x }
	a, b := fibish(x - 2)
	c, d := fibish(x - 1)
	return a + d, b + c
}
`

// Compare identical workloads with a baseline commit using GOMAXPROCS=1,
// -benchmem -benchtime=200ms -count=10, then analyze with benchstat.
// BenchmarkExecutionDescription covers the call paths changed by execution
// descriptions. DisableMemo keeps the loop workload on the interpreter backend.
func BenchmarkExecutionDescription(b *testing.B) {
	for _, name := range []string{"phi", "closure", "interfaceCall", "deferred", "recovered", "goroutine"} {
		b.Run(name, func(b *testing.B) {
			ctx := NewContext(DisableMemo)
			ctx.SetLeastCallForEnablePool(0)
			interp, err := ctx.LoadInterp("execution_bench.go", executionDescriptionSource)
			if err != nil {
				b.Fatal(err)
			}
			b.Cleanup(interp.UnsafeRelease)
			if got, err := interp.RunFunc(name); err != nil || got != 42 {
				b.Fatalf("got %v, %v; want 42", got, err)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				if got, err := interp.RunFunc(name); err != nil || got != 42 {
					b.Fatalf("got %v, %v; want 42", got, err)
				}
			}
		})
	}
}

const executionDescriptionSource = `package main
type counter int
func (c counter) Value() int { return int(c) }
func phi() int {
 a, b := 0, 1
 for i := 0; i < 10; i++ { a, b = b, a+b }
 return a-13
}
func closure() int {
 n := 41
 f := func() int { return n+1 }
 return f()
}
func interfaceCall() int {
 var v interface { Value() int } = counter(42)
 return v.Value()
}
func deferred() (n int) { defer func() { n++ }(); return 41 }
func recovered() (n int) { defer func() { recover(); n = 42 }(); panic("bench") }
func goroutine() int {
 ch := make(chan int, 1)
 go func(n int) { defer close(ch); ch <- n }(42)
 return <-ch
}
`

// Loading is measured separately for lazy SSA analysis and load-time GC plans.
// The large case crosses the lazy GC materialization threshold.
func BenchmarkExecutionDescriptionLoad(b *testing.B) {
	for _, size := range []struct {
		name      string
		additions int
	}{{"small", 8}, {"large", 1100}} {
		var source strings.Builder
		source.WriteString("package main\nimport \"runtime\"\nfunc work(n int) *int { runtime.GC(); p := new(int); *p = n\n")
		for i := 0; i < size.additions; i++ {
			fmt.Fprintf(&source, "*p += %d\n", i)
		}
		source.WriteString("runtime.GC(); return p }\n")
		code := source.String()
		for _, gc := range []bool{false, true} {
			for _, plan := range []bool{false, true} {
				if plan && !gc {
					continue
				}
				name := fmt.Sprintf("%s/gc=%t", size.name, gc)
				if plan {
					name += "/execution_plan"
				}
				b.Run(name, func(b *testing.B) {
					mode := DisableMemo
					if gc {
						mode |= ExperimentalSupportGC
					}
					if plan {
						mode |= EnableGCExecutionPlan
					}
					b.ReportAllocs()
					for b.Loop() {
						interp, err := NewContext(mode).LoadInterp("load_bench.go", code)
						if err != nil {
							b.Fatal(err)
						}
						interp.UnsafeRelease()
					}
				})
			}
		}
	}
}
