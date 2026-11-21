package ixgo

import (
	"fmt"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"golang.org/x/tools/go/ssa"
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

func testGCLivenessModes(t *testing.T, run func(*testing.T, Mode)) {
	t.Helper()
	for _, tc := range []struct {
		name string
		mode Mode
	}{{"closure", 0}, {"execution_plan", EnableGCExecutionPlan}} {
		t.Run(tc.name, func(t *testing.T) { run(t, tc.mode) })
	}
}

func TestGCLivenessPrecomputed(t *testing.T) {
	testGCLivenessModes(t, testGCLivenessPrecomputed)
}

func testGCLivenessPrecomputed(t *testing.T, planMode Mode) {
	pfn := loadGCLivenessFunction(t, ExperimentalSupportGC|planMode)
	if pfn.gcPlan != nil {
		t.Fatal("eagerly materialized function retained GC plan")
	}
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
	testGCLivenessModes(t, testGCLivenessDetectsRuntimeGCValue)
}

func testGCLivenessDetectsRuntimeGCValue(t *testing.T, planMode Mode) {
	const source = `package main

import "runtime"

var collect = runtime.GC

func work() { collect() }
`
	interp, err := NewContext(ExperimentalSupportGC|planMode).LoadInterp("main.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	if !interp.gcEnabled {
		t.Fatal("runtime.GC function value was not detected")
	}

	interp, err = NewContext(ExperimentalSupportGC|planMode).LoadInterp("main.go", "package main\nfunc work() {}\n")
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	if interp.gcEnabled {
		t.Fatal("GC liveness enabled without a runtime.GC reference")
	}
}

func TestGCLivenessLargeFunctionIsLazy(t *testing.T) {
	testGCLivenessModes(t, testGCLivenessLargeFunctionIsLazy)
}

func testGCLivenessLargeFunctionIsLazy(t *testing.T, planMode Mode) {
	const additions = maxEagerGCLivenessPCs + 100
	var source strings.Builder
	source.WriteString(`package main

import "runtime"

func newWork(captured *int) func(*int, int) int {
	return func(parameter *int, n int) int {
		runtime.GC()
`)
	for i := 0; i < additions; i++ {
		fmt.Fprintf(&source, "n += %d\n", i)
	}
	source.WriteString(`
		n += *parameter + *captured
		runtime.GC()
		return n
	}
}

func Work(n int) int {
	parameter, captured := new(int), new(int)
	*parameter, *captured = n, n + 1
	return newWork(captured)(parameter, n)
}
`)
	interp, err := NewContext(ExperimentalSupportGC|planMode).LoadInterp("main.go", source.String())
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	fn := interp.mainpkg.Func("newWork").AnonFuncs[0]
	pfn := interp.funcs[fn]
	if len(pfn.ssaInstrs) <= maxEagerGCLivenessPCs {
		t.Fatalf("instructions = %d, want more than %d", len(pfn.ssaInstrs), maxEagerGCLivenessPCs)
	}
	if pfn.gcReady == nil || pfn.gcReady[0].Load() || (pfn.gcPlan != nil) != (planMode != 0) || pfn.gcRegs[0] != nil {
		t.Fatal("large function GC registers were eagerly materialized")
	}
	parameterRegs := []register{pfn.regIndex(fn.Params[0]), pfn.regIndex(fn.FreeVars[0])}
	if planMode != 0 {
		pfn.functionCompilation = functionCompilation{}
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

	// Both parameter and captured references must survive the first GC. The
	// second GC must release them, including on repeated calls using cached
	// liveness information for this large function.
	for _, n := range []int{3, 11, 42} {
		got, err := interp.RunFunc("Work", n)
		if err != nil {
			t.Fatal(err)
		}
		if want := 3*n + 1 + additions*(additions-1)/2; got != want {
			t.Fatalf("Work(%d) = %v, want %d", n, got, want)
		}
	}
	lastGC := len(pfn.gcReady) - 1
	for lastGC > 0 && !pfn.gcReady[lastGC].Load() {
		lastGC--
	}
	if lastGC == 0 {
		t.Fatal("large function liveness was not computed during GC")
	}
	for i, value := range []ssa.Value{fn.Params[0], fn.FreeVars[0]} {
		reg := parameterRegs[i]
		if !slices.Contains(pfn.gcRegsForPC(lastGC), reg) {
			t.Errorf("dead %s register was retained after the second GC", value.Name())
		}
	}
}

func TestGCLivenessPlanPCMapping(t *testing.T) {
	pfn := loadGCLivenessFunction(t, ExperimentalSupportGC)
	pfn.gcPlan = pfn.compileGCLiveness()
	for pc, instr := range pfn.ssaInstrs {
		point := pfn.gcPlan.points[pc]
		if point.block < 0 {
			t.Fatalf("missing GC plan point at PC %d", pc)
		}
		if got := pfn.Fn.Blocks[point.block].Instrs[point.instr]; got != instr {
			t.Fatalf("GC plan PC %d maps to %v, want %v", pc, got, instr)
		}
	}
}

func BenchmarkFrameGC(b *testing.B) {
	for _, tc := range []struct {
		name string
		mode Mode
	}{{"closure", 0}, {"execution_plan", EnableGCExecutionPlan}} {
		b.Run(tc.name, func(b *testing.B) { benchmarkFrameGCMode(b, tc.mode) })
	}
}

func benchmarkFrameGCMode(b *testing.B, mode Mode) {
	pfn := loadGCLivenessFunction(b, ExperimentalSupportGC|mode)
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
		ipc:   int32(pc + 1),
	}
	b.Run("precomputed", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			fr.gc()
		}
	})
	if mode == 0 {
		b.Run("lazy-plan-query", func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				_ = pfn.gcRegsLazy(pc)
			}
		})
	} else {
		plan := pfn.compileGCLiveness()
		b.Run("plan-materialization", func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				_ = plan.regsAt(pc)
			}
		})
	}
}

func TestKeepLiveClosureBindingsWorklist(t *testing.T) {
	// wrap captures inner's MakeClosure result; inner captures f.
	// Only wrap starts live, so inner must be enqueued when its result is revived.
	pfn := &function{functionCompilation: functionCompilation{
		closureCaptures: []closureCapture{
			{result: 10, binds: []int{5}},
			{result: 5, binds: []int{1}},
		},
		closureByResult: map[int]int{10: 0, 5: 1},
	}}
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

func TestGCLivenessLoopDefinitionsAndPhiEdges(t *testing.T) {
	testGCLivenessModes(t, testGCLivenessLoopDefinitionsAndPhiEdges)
}

func testGCLivenessLoopDefinitionsAndPhiEdges(t *testing.T, planMode Mode) {
	const source = `package main
import "runtime"
func consume(p *int) { if *p != 42 { panic("value") } }
func loop(ch chan *int) {
	for p := range ch {
		consume(p)
		p = nil
		runtime.GC()
	}
}
func phi(first, second *int, choose bool) {
	p := first
	if choose {
		runtime.GC()
	} else {
		p = second
		runtime.GC()
	}
	consume(p)
}
`
	interp, err := NewContext(ExperimentalSupportGC|planMode).LoadInterp("main.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	loop := interp.funcs[interp.mainpkg.Func("loop")]
	var received ssa.Value
	for _, block := range loop.Fn.Blocks {
		for _, instr := range block.Instrs {
			if extract, ok := instr.(*ssa.Extract); ok && extract.Index == 0 {
				received = extract
			}
		}
	}
	if received == nil {
		t.Fatal("missing received pointer")
	}
	checked := 0
	for pc, instr := range loop.ssaInstrs {
		if call, ok := instr.(*ssa.Call); ok && isRuntimeGCFunction(call.Common().StaticCallee()) {
			checked++
			if !slices.Contains(loop.gcRegsForPC(pc), loop.regIndex(received)) {
				t.Error("previous loop iteration's received pointer is retained")
			}
		}
	}
	phi := interp.funcs[interp.mainpkg.Func("phi")]
	for pc, instr := range phi.ssaInstrs {
		call, ok := instr.(*ssa.Call)
		if !ok || !isRuntimeGCFunction(call.Common().StaticCallee()) {
			continue
		}
		checked++
		liveParam := 0
		if strings.Contains(instr.Block().Comment, "else") {
			liveParam = 1
		}
		for i, param := range phi.Fn.Params[:2] {
			dead := slices.Contains(phi.gcRegsForPC(pc), phi.regIndex(param))
			if dead == (i == liveParam) {
				t.Errorf("%s: parameter %s dead = %v", instr.Block(), param.Name(), dead)
			}
		}
	}
	if checked != 3 {
		t.Fatalf("checked %d GC points, want 3", checked)
	}
}

func TestGCLivenessCacheConcurrent(t *testing.T) {
	testGCLivenessModes(t, testGCLivenessCacheConcurrentMode)
}

func testGCLivenessCacheConcurrentMode(t *testing.T, mode Mode) {
	pfn := loadGCLivenessFunction(t, ExperimentalSupportGC|mode)
	if mode != 0 {
		pfn.gcPlan = pfn.compileGCLiveness()
	}
	want := make([][]register, len(pfn.gcRegs))
	for pc, regs := range pfn.gcRegs {
		want[pc] = slices.Clone(regs)
		slices.Sort(want[pc])
	}
	// Execution-plan queries must also work after compiler state is released.
	if mode != 0 {
		pfn.functionCompilation = functionCompilation{}
	} else {
		pfn.gcLiveOnce = sync.Once{}
		pfn.closureCaptures = nil
		pfn.closureByResult = nil
	}
	pfn.gcRegs = make([][]register, len(want))
	pfn.gcReady = make([]atomic.Bool, len(want))
	check := func() {
		var wg sync.WaitGroup
		for worker := 0; worker < 16; worker++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for pc := range want {
					got := slices.Clone(pfn.gcRegsForPC(pc))
					slices.Sort(got)
					if !slices.Equal(got, want[pc]) {
						t.Errorf("GC point %d: dead registers = %v, want %v", pc, got, want[pc])
					}
				}
			}()
		}
		wg.Wait()
	}
	check()
	for pc := range want {
		if !pfn.gcReady[pc].Load() {
			t.Errorf("GC point %d was not cached", pc)
		}
	}
	check()
}

func TestGCLivenessRecoverNamedResult(t *testing.T) {
	testGCLivenessModes(t, testGCLivenessRecoverNamedResult)
}

func testGCLivenessRecoverNamedResult(t *testing.T, planMode Mode) {
	const source = `package main
import "runtime"
func work() (result *int) {
	defer func() { recover() }()
	result = new(int)
	*result = 42
	runtime.GC()
	panic("resume in recover block")
}
func Work() int { return *work() }
`
	interp, err := NewContext(ExperimentalSupportGC|planMode).LoadInterp("main.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	pfn := interp.funcs[interp.mainpkg.Func("work")]
	plan := pfn.compileGCLiveness()
	if pfn.Fn.Recover == nil || len(plan.recover) == 0 {
		t.Fatal("missing recover block liveness")
	}
	for pc, instr := range pfn.ssaInstrs {
		call, ok := instr.(*ssa.Call)
		if !ok || !isRuntimeGCFunction(call.Common().StaticCallee()) {
			continue
		}
		for _, reg := range plan.recover {
			if slices.Contains(pfn.gcRegsForPC(pc), reg) {
				t.Errorf("recover register %d is dead at GC point %d", reg, pc)
			}
		}
	}
	for range 3 {
		got, err := interp.RunFunc("Work")
		if err != nil {
			t.Fatal(err)
		}
		if got != 42 {
			t.Fatalf("Work() = %v, want 42", got)
		}
	}
}

func TestGCLivenessRecoverRootsAfterDefinitions(t *testing.T) {
	plan := &gcLivenessPlan{
		registers: []register{0, 1}, recover: []register{0},
		points: []gcLivenessPoint{{block: 0, instr: 0}},
		blocks: []gcLivenessBlock{{transfers: []gcLivenessTransfer{{def: -1}, {def: 0}, {def: 1}}}},
	}
	if got := plan.regsAt(0); !slices.Equal(got, []register{1}) {
		t.Fatalf("dead registers = %v; want [1] with recover root 0 retained", got)
	}
}

func TestGCLivenessModesEquivalent(t *testing.T) {
	ctx := NewContext(ExperimentalSupportGC | EnableGCExecutionPlan)
	interp, err := ctx.LoadInterp("equivalence.go", `package main
import "runtime"
func work(n int) (result *int) {
 defer func() { recover(); if result != nil { *result += 1 } }()
 result = new(int)
 inner := func() int { return *result }
 outer := func() int { return inner() }
 for i := 0; i < n; i++ {
  if i%2 == 0 { *result += outer()+i } else { *result += i }
  runtime.GC()
 }
 if n < 0 { panic("recover") }
 return result
}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for fn, p := range interp.funcs {
		if !p.hasGCRegs {
			continue
		}
		plan := p.compileGCLiveness()
		for pc := range p.ssaInstrs {
			want := p.gcRegsLazy(pc)
			got := plan.regsAt(pc)
			slices.Sort(want)
			slices.Sort(got)
			if !slices.Equal(got, want) {
				t.Fatalf("%s PC %d: plan %v, SSA %v", fn, pc, got, want)
			}
		}
	}
}

func TestGCExecutionPlanModeWithoutGC(t *testing.T) {
	p := loadGCLivenessFunction(t, EnableGCExecutionPlan)
	if p.gcRegs != nil || p.gcPlan != nil {
		t.Fatal("plan mode enabled GC implicitly")
	}
}
