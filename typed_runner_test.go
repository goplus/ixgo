package ixgo

import (
	"go/types"
	"reflect"
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
