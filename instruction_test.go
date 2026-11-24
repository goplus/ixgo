package ixgo

import (
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/tools/go/ssa"
)

func TestDescriptionPreservesFrameLayout(t *testing.T) {
	for _, extra := range []Mode{0, ExperimentalSupportGC, EnableCachedReg, ExperimentalSupportGC | EnableGCExecutionPlan} {
		t.Run(fmt.Sprint(extra), func(t *testing.T) {
			load := func(mode Mode) *Interp {
				i, err := NewContext(DisableMemo|extra|mode).LoadInterp("description_layout.go", executionDescriptionSource)
				if err != nil {
					t.Fatal(err)
				}
				t.Cleanup(i.UnsafeRelease)
				return i
			}
			baseline, described := load(0), load(EnableInstructionDescriptions)
			byName := make(map[string]*function)
			for fn, p := range baseline.funcs {
				byName[fn.String()] = p
			}
			for fn, p := range described.funcs {
				old := byName[fn.String()]
				if old == nil {
					t.Fatalf("description added function %s", fn)
				}
				if len(p.stack) != len(old.stack) || len(p.index) != len(old.index) || p.hasGCRegs != old.hasGCRegs {
					t.Fatalf("%s layout changed: stack %d/%d, index %d/%d, GC %t/%t", fn, len(p.stack), len(old.stack), len(p.index), len(old.index), p.hasGCRegs, old.hasGCRegs)
				}
				registers := func(f *function) map[uint32]string {
					result := make(map[uint32]string)
					for v, encoded := range f.index {
						result[encoded] = reflect.TypeOf(v).String() + ":" + v.Name() + ":" + v.String()
					}
					return result
				}
				if !reflect.DeepEqual(registers(p), registers(old)) {
					t.Fatalf("%s register mapping changed: got %v want %v", fn, registers(p), registers(old))
				}
			}
			if len(described.funcs) != len(baseline.funcs) {
				t.Fatal("function count changed")
			}
		})
	}
}

func TestDescriptionTupleAndIterator(t *testing.T) {
	ctx := NewContext(DisableMemo | EnableInstructionDescriptions)
	i, err := ctx.LoadInterp("description_tuple.go", `package main
func pair(n int) (int, int) { return n, n+1 }
func work(n int) int {
 a, b := pair(n)
 for _, r := range "ab" { a += int(r) }
 for _, v := range map[int]int{1: 2} { b += v }
 return a+b
}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	var extracts, strings, maps int
	for _, p := range i.funcs {
		for pc, desc := range p.Code {
			switch v := p.ssaInstrs[pc].(type) {
			case *ssa.Extract:
				extracts++
				aux, ok := desc.Aux.(extractAux)
				if !ok || aux.Index != v.Index {
					t.Fatalf("extract metadata = %#v", desc)
				}
				if len(desc.Args) != 1 || desc.Args[0] < 0 || desc.Args[0] != descriptionRegister(p, v.Tuple) {
					t.Fatalf("extract args = %v", desc.Args)
				}
			case *ssa.Next:
				aux, ok := desc.Aux.(nextAux)
				if !ok || aux.IsString != v.IsString {
					t.Fatalf("next metadata = %#v", desc)
				}
				if len(desc.Args) != 1 || desc.Args[0] < 0 || desc.Args[0] != descriptionRegister(p, v.Iter) {
					t.Fatalf("next args = %v", desc.Args)
				}
				if v.IsString {
					strings++
				} else {
					maps++
				}
			}
		}
	}
	if extracts == 0 || strings == 0 || maps == 0 {
		t.Fatalf("missing coverage: extracts=%d strings=%d maps=%d", extracts, strings, maps)
	}
	if got, err := i.RunFunc("work", 3); err != nil || got != 204 {
		t.Fatalf("work = %v, %v; want 204", got, err)
	}
}

func TestInstructionOpcodeClassification(t *testing.T) {
	tests := []struct {
		instr ssa.Instruction
		op    instructionOpcode
	}{
		{&ssa.Alloc{}, opAlloc},
		{&ssa.Phi{}, opPhi},
		{&ssa.Call{}, opCall},
		{&ssa.BinOp{}, opBinOp},
		{&ssa.UnOp{}, opUnOp},
		{&ssa.ChangeInterface{}, opChangeInterface},
		{&ssa.ChangeType{}, opChangeType},
		{&ssa.Convert{}, opConvert},
		{&ssa.MakeInterface{}, opMakeInterface},
		{&ssa.MakeClosure{}, opMakeClosure},
		{&ssa.MakeChan{}, opMakeChan},
		{&ssa.MakeMap{}, opMakeMap},
		{&ssa.MakeSlice{}, opMakeSlice},
		{&ssa.Slice{}, opSlice},
		{&ssa.FieldAddr{}, opFieldAddr},
		{&ssa.Field{}, opField},
		{&ssa.IndexAddr{}, opIndexAddr},
		{&ssa.Index{}, opIndex},
		{&ssa.Lookup{}, opLookup},
		{&ssa.Select{}, opSelect},
		{&ssa.SliceToArrayPointer{}, opSliceToArrayPointer},
		{&ssa.Range{}, opRange},
		{&ssa.Next{}, opNext},
		{&ssa.TypeAssert{}, opTypeAssert},
		{&ssa.Extract{}, opExtract},
		{&ssa.Jump{}, opJump},
		{&ssa.If{}, opIf},
		{&ssa.Return{}, opReturn},
		{&ssa.RunDefers{}, opRunDefers},
		{&ssa.Panic{}, opPanic},
		{&ssa.Go{}, opGo},
		{&ssa.Defer{}, opDefer},
		{&ssa.Send{}, opSend},
		{&ssa.Store{}, opStore},
		{&ssa.MapUpdate{}, opMapUpdate},
		{&ssa.DebugRef{}, opDebugRef},
	}
	for _, test := range tests {
		if got := opcodeOf(test.instr); got != test.op {
			t.Fatalf("%T: opcode %v, want %v", test.instr, got, test.op)
		}
		_, isValue := test.instr.(ssa.Value)
		if got := producesRegister(test.instr); got != isValue {
			t.Fatalf("%T: produces register = %t, want %t", test.instr, got, isValue)
		}
	}
	if opcodeOf(nil) != opInvalid {
		t.Fatal("unknown instruction must be invalid")
	}
}

func TestDescriptionRegisterLookup(t *testing.T) {
	p := &function{}
	p.index = make(map[ssa.Value]uint32)
	fn := &ssa.Function{}
	p.index[fn] = 7 | uint32(kindFunction)<<30
	if got := descriptionRegister(p, fn); got != 7 {
		t.Fatalf("register = %d", got)
	}
	if descriptionRegister(p, nil) != -1 || descriptionRegister(p, &ssa.Function{}) != -1 {
		t.Fatal("unallocated operand must use -1")
	}
	before := len(p.index)
	d := describeInstruction(p, &ssa.MakeClosure{Fn: &ssa.Function{}}, nil)
	if d.Result != -1 || len(p.index) != before || len(p.stack) != 0 {
		t.Fatal("description allocated a register")
	}
}
