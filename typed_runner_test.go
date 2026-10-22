package ixgo

import (
	"go/types"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"golang.org/x/tools/go/ssa"
)

type testScalarRunner struct{}

func (testScalarRunner) Run(v Value) Value { return v.(int) + 1 }

type testArgsRunner struct{ result Value }

func (r testArgsRunner) Run(Value) Value       { panic("RunArgs required") }
func (r testArgsRunner) RunArgs([]Value) Value { return r.result }

type testBindingRunner struct {
	testArgsRunner
	accept bool
	bound  bool
}

func (r *testBindingRunner) BindTypes(resolve func(types.Type) reflect.Type) bool {
	r.bound = resolve(types.Typ[types.Int]) == reflect.TypeOf(int(0))
	return r.accept
}

func TestMixedRunnerDispatch(t *testing.T) {
	p := &function{narg: 1, nres: 1}
	if _, ok := p.runMixed([]value{2}, nil); ok {
		t.Fatal("runner absent")
	}
	p.mixedRunner = testScalarRunner{}
	if got, ok := p.runMixed([]value{2}, nil); !ok || got != 3 {
		t.Fatalf("scalar: %v, %v", got, ok)
	}
	for _, tc := range []struct {
		narg, nenv, nres int
		env              []value
	}{{0, 0, 1, nil}, {1, 1, 1, nil}, {1, 0, 1, []value{}}, {1, 0, 0, nil}} {
		p.nenv, p.nres = tc.nenv, tc.nres
		if p.canRunMixed(tc.narg, tc.env) {
			t.Fatalf("accepted incompatible call: %+v", tc)
		}
	}
	p.narg, p.nenv, p.nres = 2, 0, 2
	p.mixedRunner = testArgsRunner{tuple{4, true}}
	if got, ok := p.runMixed([]value{1, 2}, nil); !ok || !reflect.DeepEqual(got, tuple{4, true}) {
		t.Fatalf("arguments: %v, %v", got, ok)
	}
}

func TestMixedCallSpecializerResults(t *testing.T) {
	for _, want := range []tuple{nil, {7}, {7, true}} {
		p := &function{narg: 1, nres: len(want), mixedRunner: testArgsRunner{want}}
		if len(want) == 1 {
			p.mixedRunner = testArgsRunner{want[0]}
		}
		dests := make([]register, len(want))
		for j := range dests {
			dests[j] = register(j + 1)
		}
		run := makeCallSpecializerInstr(nil, p, []register{0}, dests)
		if run == nil {
			t.Fatal("missing mixed call specializer")
		}
		fr := &frame{stack: make([]value, 3)}
		fr.stack[0] = 5
		run(fr)
		for j, v := range want {
			if fr.stack[j+1] != v {
				t.Fatalf("result %d: %v, want %v", j, fr.stack[j+1], v)
			}
		}
	}
}

func TestMixedRunnerTypeBinding(t *testing.T) {
	old := mixedRunnerFactory
	t.Cleanup(func() { RegisterMixedRunner(old) })
	for _, accept := range []bool{false, true} {
		runner := &testBindingRunner{testArgsRunner: testArgsRunner{42}, accept: accept}
		RegisterMixedRunner(func(fn *ssa.Function) TypedRunner {
			if fn.Name() == "answer" {
				return runner
			}
			return nil
		})
		i, err := NewContext(DisableMemo).LoadInterp("binding.go", `package main
func answer(n int) int { return n + 1 }
`)
		if err != nil {
			t.Fatal(err)
		}
		if !runner.bound {
			t.Fatal("runner types were not bound")
		}
		got, err := i.RunFunc("answer", 2)
		i.UnsafeRelease()
		want := 3
		if accept {
			want = 42
		}
		if err != nil || got != want {
			t.Fatalf("accept=%v: %v, %v", accept, got, err)
		}
	}
}

func TestMixedReflectResults(t *testing.T) {
	for _, tc := range []struct {
		typ    reflect.Type
		result value
		want   []value
	}{
		{reflect.TypeOf(func(int) {}), nil, nil},
		{reflect.TypeOf(func(int) *int { return nil }), nil, []value{(*int)(nil)}},
		{reflect.TypeOf(func(int) (int, bool) { return 0, false }), tuple{7, true}, []value{7, true}},
	} {
		p := &function{narg: 1, nres: tc.typ.NumOut(), mixedRunner: testArgsRunner{tc.result}}
		got := new(Interp).callFunctionByReflect(nil, p, tc.typ, []reflect.Value{reflect.ValueOf(2)}, nil)
		if len(got) != len(tc.want) {
			t.Fatalf("result count %d, want %d", len(got), len(tc.want))
		}
		for j, v := range got {
			if v.Interface() != tc.want[j] {
				t.Fatalf("result %d: %v, want %v", j, v.Interface(), tc.want[j])
			}
		}
	}
}

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

func TestDeferredMixedCallInstallation(t *testing.T) {
	oldMemo, oldMixed := memoPlugin, mixedRunnerFactory
	t.Cleanup(func() {
		RegisterMemoPlugin(oldMemo)
		RegisterMixedRunner(oldMixed)
	})
	for _, tc := range []struct {
		name   string
		setup  func(*Context)
		accept bool
		memo   bool
		want   int
	}{
		{name: "installed", accept: true, want: 43},
		{name: "rejected", want: 2},
		{name: "memo", accept: true, memo: true, want: 2},
		{name: "tracing", setup: func(ctx *Context) { ctx.Mode |= EnableTracing }, accept: true, want: 2},
		{name: "evaluation", setup: func(ctx *Context) { ctx.evalCallFn = func(*Interp, *ssa.Call, ...interface{}) {} }, accept: true, want: 2},
		{name: "hooks", setup: func(ctx *Context) { ctx.SetFuncHook(func(*FuncInfo) {}, func(*FuncInfo) {}) }, accept: true, want: 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			RegisterMemoPlugin(nil)
			RegisterMixedRunner(nil)
			ctx := NewContext(0)
			if tc.setup != nil {
				tc.setup(ctx)
			}
			i, err := ctx.LoadInterp("deferred_mixed.go", `package main
func A(n int) int { if n == 0 { return 0 }; return B(n-1)+1 }
func B(n int) int { if n == 0 { return 0 }; return A(n-1)+1 }
`)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			i.funcs = make(map[*ssa.Function]*function)
			if tc.memo {
				RegisterMemoPlugin(func(p MemoInstaller, fn *ssa.Function) {
					if fn.Name() == "B" {
						p.InstallScalarMemo()
					}
				})
			}
			RegisterMixedRunner(func(fn *ssa.Function) TypedRunner {
				if tc.accept && fn.Name() == "A" {
					return testArgsRunner{42}
				}
				return nil
			})
			v := visitor{intp: i, prog: i.mainpkg.Prog, pkgs: map[*ssa.Package]bool{i.mainpkg: true}, seen: make(map[*ssa.Function]bool), base: fnBase}
			v.function(i.mainpkg.Func("A"))
			if len(v.mixedCalls) != 0 || len(v.mixedPending) != 0 {
				t.Fatal("deferred installation state retained")
			}
			caller := i.funcs[i.mainpkg.Func("B")]
			for index, instr := range caller.ssaInstrs {
				if call, ok := instr.(*ssa.Call); ok && call.Common().StaticCallee() == i.mainpkg.Func("A") {
					name := runtime.FuncForPC(reflect.ValueOf(caller.Instrs[index]).Pointer()).Name()
					if strings.Contains(name, "deferMixedCall") {
						t.Fatalf("retained deferred trampoline: %s", name)
					}
					if tc.name == "installed" && !strings.Contains(name, "makeCallSpecializerInstr") {
						t.Fatalf("instruction was not replaced directly: %s", name)
					}
				}
			}
			got, err := i.RunFunc("B", 2)
			if err != nil || got != tc.want {
				t.Fatalf("B(2) = %v, %v; want %d", got, err, tc.want)
			}
		})
	}
}

func TestDeferredMixedCallElidedResults(t *testing.T) {
	old := mixedRunnerFactory
	t.Cleanup(func() { RegisterMixedRunner(old) })
	for _, body := range []string{
		"a, b := A(n-1); return a, b",
		"a, _ := A(n-1); return a, false",
		"_, b := A(n-1); return 0, b",
		"A(n-1); return 0, false",
	} {
		t.Run(body, func(t *testing.T) {
			RegisterMixedRunner(nil)
			i, err := NewContext(DisableMemo).LoadInterp("deferred_tuple.go", `package main
func A(n int) (int, bool) { if n == 0 { return 0, false }; return B(n-1) }
func B(n int) (int, bool) { `+body+` }
`)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			i.funcs = make(map[*ssa.Function]*function)
			RegisterMixedRunner(func(fn *ssa.Function) TypedRunner {
				if fn.Name() == "A" {
					return testArgsRunner{tuple{42, true}}
				}
				return nil
			})
			v := visitor{intp: i, prog: i.mainpkg.Prog, pkgs: map[*ssa.Package]bool{i.mainpkg: true}, seen: make(map[*ssa.Function]bool), base: fnBase}
			v.function(i.mainpkg.Func("A"))
			got, err := i.RunFunc("B", 2)
			want := tuple{42, true}
			switch body {
			case "a, _ := A(n-1); return a, false":
				want[1] = false
			case "_, b := A(n-1); return 0, b":
				want[0] = 0
			case "A(n-1); return 0, false":
				want = tuple{0, false}
			}
			if err != nil || !reflect.DeepEqual(got, want) {
				t.Fatalf("B(2) = %v, %v; want %v", got, err, want)
			}
		})
	}
}

func TestRunFuncWithoutMixedArgumentAllocation(t *testing.T) {
	oldMemo, oldMixed := memoPlugin, mixedRunnerFactory
	t.Cleanup(func() {
		RegisterMemoPlugin(oldMemo)
		RegisterMixedRunner(oldMixed)
	})
	RegisterMemoPlugin(nil)
	for _, tc := range []struct {
		name    string
		factory TypedRunnerFactory
	}{
		{name: "disabled"},
		{name: "rejected", factory: func(*ssa.Function) TypedRunner { return nil }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			RegisterMixedRunner(tc.factory)
			i, err := NewContext(DisableMemo).LoadInterp("identity.go", `package main
func Identity(n int) int { return n }
`)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			allocs := testing.AllocsPerRun(1000, func() {
				got, err := i.RunFunc("Identity", 7)
				if err != nil || got != 7 {
					t.Fatalf("Identity(7) = %v, %v", got, err)
				}
			})
			if allocs > 1 {
				t.Fatalf("RunFunc allocated %g times; want only the root frame allocation", allocs)
			}
		})
	}
}

type testRetainingArgsRunner struct{ calls [][]Value }

func (*testRetainingArgsRunner) Run(Value) Value { panic("RunArgs required") }
func (r *testRetainingArgsRunner) RunArgs(args []Value) Value {
	r.calls = append(r.calls, args)
	return args
}

func TestMixedRunnerRetainedArguments(t *testing.T) {
	old := mixedRunnerFactory
	t.Cleanup(func() { RegisterMixedRunner(old) })
	runner := new(testRetainingArgsRunner)
	RegisterMixedRunner(func(fn *ssa.Function) TypedRunner {
		if fn.Name() == "Identity" {
			return runner
		}
		return nil
	})
	i, err := NewContext(DisableMemo).LoadInterp("retained_args.go", `package main
func Identity(n int) int { return n }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	args := []Value{7}
	got, err := i.RunFunc("Identity", args...)
	if err != nil {
		t.Fatal(err)
	}
	args[0] = 9
	if _, err := i.RunFunc("Identity", 11); err != nil {
		t.Fatal(err)
	}
	pfn := i.funcs[i.mainpkg.Func("Identity")]
	i.callFunctionDiscardsResult(&frame{}, pfn, args, nil)
	args[0] = 13
	runtime.GC()
	if !reflect.DeepEqual(got, []Value{7}) {
		t.Fatalf("returned arguments changed: %v", got)
	}
	if want := [][]Value{{7}, {11}, {9}}; !reflect.DeepEqual(runner.calls, want) {
		t.Fatalf("retained arguments = %v, want %v", runner.calls, want)
	}
}

func TestDeferredMixedCallPreservesTuple(t *testing.T) {
	old := mixedRunnerFactory
	t.Cleanup(func() { RegisterMixedRunner(old) })
	RegisterMixedRunner(nil)
	i, err := NewContext(DisableMemo).LoadInterp("tuple.go", `package main
func pair() (float32, int) { return 1, 2 }
func caller() int { _, n := pair(); return n }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	caller := i.funcs[i.mainpkg.Func("caller")]
	var call *ssa.Call
	for _, instr := range caller.ssaInstrs {
		if c, ok := instr.(*ssa.Call); ok {
			call = c
			break
		}
	}
	if call == nil {
		t.Fatal("missing pair call")
	}
	// Model a caller compiled through the tuple-producing path before the
	// callee's mixed runner becomes available.
	caller.elidedExtract = nil
	v := visitor{intp: i}
	v.deferMixedCall(caller, call, 0)
	if len(v.mixedCalls) != 0 {
		t.Fatal("tuple-producing call queued for direct-result replacement")
	}
	got, err := i.RunFunc("caller")
	if err != nil || got != 2 {
		t.Fatalf("caller: %v, %v", got, err)
	}
}
