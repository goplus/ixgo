package ixgo

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

const gcLivenessSource = `package main

import "runtime"

func work(n int) *int {
	runtime.GC()
	temporary := new(int)
	*temporary = n
	seed := *temporary
	x := new(int)
	for i := 0; i < n; i++ {
		*x += seed + i
	}
	return x
}
`

func loadGCLivenessFunction(b testing.TB, mode Mode) *function {
	b.Helper()
	interp, err := NewContext(mode).LoadInterp("main.go", gcLivenessSource)
	if err != nil {
		b.Fatal(err)
	}
	b.Cleanup(interp.UnsafeRelease)
	fn := interp.mainpkg.Func("work")
	pfn := interp.funcs[fn]
	if pfn == nil {
		b.Fatal("work function was not loaded")
	}
	return pfn
}

func TestGCLivenessPrecomputed(t *testing.T) {
	pfn := loadGCLivenessFunction(t, ExperimentalSupportGC)
	if len(pfn.gcRegs) != len(pfn.ssaInstrs) {
		t.Fatalf("GC points = %d, want %d", len(pfn.gcRegs), len(pfn.ssaInstrs))
	}
	for pc, regs := range pfn.gcRegs {
		for _, reg := range regs {
			if reg < 0 || int(reg) >= len(pfn.stack) {
				t.Fatalf("GC point %d contains invalid register %d", pc, reg)
			}
		}
	}

	pfn = loadGCLivenessFunction(t, 0)
	if pfn.gcRegs != nil {
		t.Fatalf("GC liveness generated without ExperimentalSupportGC")
	}
}

func TestGCLivenessDetectsRuntimeGCValue(t *testing.T) {
	const source = `package main

import "runtime"

var collect = runtime.GC

func work() { collect() }
`
	interp, err := NewContext(ExperimentalSupportGC).LoadInterp("main.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	if !interp.gcEnabled {
		t.Fatal("runtime.GC function value was not detected")
	}

	interp, err = NewContext(ExperimentalSupportGC).LoadInterp("main.go", "package main\nfunc work() {}\n")
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	if interp.gcEnabled {
		t.Fatal("GC liveness enabled without a runtime.GC reference")
	}
}

func TestGCLivenessLargeFunctionIsLazy(t *testing.T) {
	var source strings.Builder
	source.WriteString("package main\nimport \"runtime\"\nfunc work(n int) int {\np := new(int)\n*p = n\nruntime.GC()\n")
	for i := 0; i < maxEagerGCLivenessPCs+100; i++ {
		fmt.Fprintf(&source, "n += %d\n", i)
	}
	source.WriteString("return n + *p\n}\n")

	ctx := NewContext(ExperimentalSupportGC)
	interp, err := ctx.LoadInterp("main.go", source.String())
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	pfn := interp.funcs[interp.mainpkg.Func("work")]
	if len(pfn.ssaInstrs) <= maxEagerGCLivenessPCs {
		t.Fatalf("instructions = %d, want more than %d", len(pfn.ssaInstrs), maxEagerGCLivenessPCs)
	}
	if pfn.gcReady == nil || pfn.gcReady[0].Load() {
		t.Fatal("large function liveness was eagerly computed")
	}
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = pfn.gcRegsForPC(0)
		}()
	}
	wg.Wait()
	if !pfn.gcReady[0].Load() {
		t.Fatal("liveness was not cached after first GC")
	}
}

func BenchmarkFrameGC(b *testing.B) {
	pfn := loadGCLivenessFunction(b, ExperimentalSupportGC)
	pc := 0
	for i := range pfn.gcRegs {
		if len(pfn.gcRegs[i]) > len(pfn.gcRegs[pc]) {
			pc = i
		}
	}
	if len(pfn.gcRegs[pc]) == 0 {
		b.Fatal("benchmark function has no dead registers")
	}
	fr := &frame{
		pfn:   pfn,
		stack: append([]value(nil), pfn.stack...),
		ipc:   pc + 1,
	}
	b.Run("precomputed", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			fr.gc()
		}
	})
	b.Run("dynamic-analysis", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = pfn.gcRegsAt(pc)
		}
	})
}

func TestKeepLiveClosureBindingsWorklist(t *testing.T) {
	// wrap captures inner's MakeClosure result; inner captures f.
	// Only wrap starts live, so inner must be enqueued when its result is revived.
	pfn := &function{
		closureCaptures: []closureCapture{
			{result: 10, binds: []int{5}},
			{result: 5, binds: []int{1}},
		},
		closureByResult: map[int]int{10: 0, 5: 1},
	}
	dead := map[int]bool{1: true, 5: true, 10: false}
	pfn.keepLiveClosureBindings(dead)
	if dead[5] || dead[1] {
		t.Fatalf("nested bindings still dead: %v", dead)
	}
}

func TestInitClosureCapturesNilFn(t *testing.T) {
	pfn := &function{}
	pfn.initClosureCaptures()
	if pfn.closureCaptures == nil {
		t.Fatal("expected empty non-nil slice")
	}
	pfn.initClosureCaptures()
}
