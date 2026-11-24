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

func TestInstructionAuxVariants(t *testing.T) {
	variants := []instructionAux{
		binaryOpAux{},
		unaryOpAux{},
		phiAux{},
		branchAux{},
		returnAux{},
		extractAux{},
		nextAux{},
		callAux{},
	}
	for _, aux := range variants {
		t.Run(reflect.TypeOf(aux).Name(), func(t *testing.T) {
			aux.instructionAux()
		})
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
	d := describeInstruction(p, &ssa.MakeClosure{Fn: &ssa.Function{}}, nil, nil)
	if d.Result != -1 || len(p.index) != before || len(p.stack) != 0 {
		t.Fatal("description allocated a register")
	}
}

func TestDescriptionPhiGroup(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		t.Run(fmt.Sprint(enabled), func(t *testing.T) {
			mode := DisableMemo
			if enabled {
				mode |= EnableInstructionDescriptions
			}
			i, err := NewContext(mode).LoadInterp("phi_group.go", `package main
func swap(n int) int {
 a, b := 1, 2
 for j := 0; j < n; j++ { a, b = b, a }
 return a*10+b
}
`)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			p := i.funcs[i.mainpkg.Func("swap")]
			var groups int
			for pc, desc := range p.Code {
				if desc.Op != opPhi {
					continue
				}
				groups++
				aux := desc.Aux.(phiAux)
				block := p.ssaInstrs[pc].Block()
				if len(aux.Preds) != len(block.Preds) {
					t.Fatal("predecessor count mismatch")
				}
				var phis []*ssa.Phi
				for _, instr := range block.Instrs {
					phi, ok := instr.(*ssa.Phi)
					if !ok {
						break
					}
					phis = append(phis, phi)
				}
				if len(phis) != 3 || len(aux.Results) != len(phis) || len(aux.Edges) != len(phis) {
					t.Fatalf("incomplete Phi group: %#v", aux)
				}
				for j, pred := range block.Preds {
					if aux.Preds[j] != int32(pred.Index) {
						t.Fatal("predecessor order mismatch")
					}
				}
				for k, phi := range phis {
					if aux.Results[k] < 0 || aux.Results[k] != descriptionRegister(p, phi) {
						t.Fatal("result register mismatch")
					}
					if len(aux.Edges[k]) != len(phi.Edges) {
						t.Fatal("edge count mismatch")
					}
					for j, edge := range phi.Edges {
						if aux.Edges[k][j] < 0 || aux.Edges[k][j] != descriptionRegister(p, edge) {
							t.Fatal("edge register mismatch")
						}
					}
				}
				// Execute the original closure with only the grouped description available
				// to the test oracle. Cyclic inputs must be read before any result is stored.
				initial := append([]value(nil), p.stack...)
				for k, result := range aux.Results {
					initial[result] = 100 + k
				}
				for j, pred := range aux.Preds {
					fr := &frame{pfn: p, pred: pred, stack: append([]value(nil), initial...)}
					want := make([]value, len(aux.Results))
					for k := range want {
						want[k] = fr.reg(aux.Edges[k][j])
					}
					compilation := p.functionCompilation
					p.functionCompilation = functionCompilation{}
					p.Instrs[pc](fr)
					p.functionCompilation = compilation
					for k, result := range aux.Results {
						if !reflect.DeepEqual(fr.reg(result), want[k]) {
							t.Fatalf("parallel assignment %d: got %v, want %v", k, fr.reg(result), want[k])
						}
					}
				}
			}
			if enabled && groups != 1 {
				t.Fatalf("Phi group count = %d, want 1", groups)
			}
			if !enabled && len(p.Code) != 0 {
				t.Fatal("default backend generated descriptions")
			}
			for _, n := range []int{0, 1, 2, 3, 10, 11} {
				want := 12
				if n%2 != 0 {
					want = 21
				}
				if got, err := i.RunFunc("swap", n); err != nil || got != want {
					t.Fatalf("swap(%d) = %v, %v; want %d", n, got, err, want)
				}
			}
		})
	}
}

func TestDescriptionCallResultsAndOperandKinds(t *testing.T) {
	i, err := NewContext(DisableMemo|EnableInstructionDescriptions).LoadInterp("call_metadata.go", `package main
func pair(n int) (int, int) { return n, n+1 }
func all(n int) int { a,b := pair(n); return a+b }
func partial(n int) int { _,b := pair(n); return b }
func dynamic(f func(int)(int,int), n int) int { a,b := f(n); return a+b }
func slices(s []int) int { return len(s[:]) }
func closure(n int) int { f := func() int { return n }; return f() }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	var elided, ordinary, absent, static, builtin, registers int
	for _, p := range i.funcs {
		for pc, d := range p.Code {
			if len(d.Args) != len(d.ArgKinds) {
				t.Fatal("operand classification count mismatch")
			}
			operands := p.ssaInstrs[pc].Operands(nil)
			if len(operands) != len(d.Args) {
				t.Fatalf("%s pc %d: operand count mismatch", p.Fn, pc)
			}
			for j, kind := range d.ArgKinds {
				operand := *operands[j]
				wantKind := operandRegister
				switch operand.(type) {
				case nil:
					wantKind = operandAbsent
				case *ssa.Function:
					wantKind = operandFunction
				case *ssa.Builtin:
					wantKind = operandBuiltin
				default:
					if _, allocated := p.index[operand]; !allocated {
						wantKind = operandUnallocated
					}
				}
				if kind != wantKind || d.Args[j] != descriptionRegister(p, operand) {
					t.Fatalf("%s pc %d operand %d (%T): got kind/register %v/%v, want %v/%v", p.Fn, pc, j, operand, kind, d.Args[j], wantKind, descriptionRegister(p, operand))
				}
				switch kind {
				case operandAbsent:
					absent++
					if d.Args[j] != -1 {
						t.Fatal("absent operand has register")
					}
				case operandFunction:
					static++
				case operandBuiltin:
					builtin++
				case operandRegister:
					registers++
					if d.Args[j] < 0 {
						t.Fatal("register operand missing register")
					}
				}
			}
			call, ok := p.ssaInstrs[pc].(*ssa.Call)
			if !ok {
				continue
			}
			aux, ok := d.Aux.(callAux)
			if !ok {
				t.Fatal("missing call metadata")
			}
			if aux.ExtractResults == nil {
				ordinary++
				continue
			}
			elided++
			if len(aux.ExtractResults) != callResultCount(&call.Call) {
				t.Fatal("result arity mismatch")
			}
			for _, ref := range *call.Referrers() {
				ex, ok := ref.(*ssa.Extract)
				if !ok {
					continue
				}
				if _, ok := p.elidedExtract[ex]; !ok {
					t.Fatal("unexpected extract destination")
				}
				want := register(-1)
				if uses := ex.Referrers(); uses != nil && len(*uses) > 0 {
					want = descriptionRegister(p, ex)
				}
				if aux.ExtractResults[ex.Index] != want {
					t.Fatal("extract destination mismatch")
				}
			}
		}
	}
	if elided == 0 || ordinary == 0 || absent == 0 || static == 0 || builtin == 0 || registers == 0 {
		t.Fatalf("missing metadata cases: elided=%d ordinary=%d absent=%d static=%d builtin=%d registers=%d", elided, ordinary, absent, static, builtin, registers)
	}
	for _, name := range []string{"all", "partial", "dynamic"} {
		p := i.funcs[i.mainpkg.Func(name)]
		calls := 0
		for pc, d := range p.Code {
			call, ok := p.ssaInstrs[pc].(*ssa.Call)
			if !ok {
				continue
			}
			calls++
			aux := d.Aux.(callAux)
			if len(aux.ExtractResults) != 2 {
				t.Fatalf("%s: expected two result slots, got %v", name, aux.ExtractResults)
			}
			if d.Result != -1 {
				t.Fatalf("%s: elided call result = %d, want -1", name, d.Result)
			}
			for index, dest := range aux.ExtractResults {
				unused := name == "partial" && index == 0
				if (dest < 0) != unused {
					t.Fatalf("%s result %d: register %d, unused=%t", name, index, dest, unused)
				}
			}
			if name == "dynamic" {
				if _, ok := call.Call.Value.(*ssa.Parameter); !ok || d.ArgKinds[0] != operandRegister {
					t.Fatal("dynamic call target must use a register")
				}
			} else if d.ArgKinds[0] != operandFunction {
				t.Fatalf("%s: expected static function target", name)
			}
		}
		if calls != 1 {
			t.Fatalf("%s: call count %d, want 1", name, calls)
		}
	}
	if got, err := i.RunFunc("dynamic", func(n int) (int, int) { return n, n + 1 }, 3); err != nil || got != 7 {
		t.Fatalf("dynamic = %v, %v", got, err)
	}
	if got, err := i.RunFunc("slices", []int{1, 2, 3}); err != nil || got != 3 {
		t.Fatalf("slices = %v, %v", got, err)
	}

	for _, name := range []string{"all", "partial", "closure"} {
		want := 7
		if name == "partial" {
			want = 4
		}
		if name == "closure" {
			want = 3
		}
		if got, err := i.RunFunc(name, 3); err != nil || got != want {
			t.Fatalf("%s = %v, %v", name, got, err)
		}
	}
}

func TestDescriptionUnallocatedAndUnelidedOperands(t *testing.T) {
	p := &function{}
	operand := ssa.Value(&ssa.Parameter{})
	d := describeInstruction(p, &ssa.DebugRef{}, []*ssa.Value{&operand}, nil)
	if len(d.ArgKinds) != 1 || d.ArgKinds[0] != operandUnallocated || d.Args[0] != -1 {
		t.Fatalf("unallocated operand = %#v", d)
	}
	i, err := NewContext(DisableMemo|EnableInstructionDescriptions).LoadInterp("unelided.go", `package main
func lookup(m map[int]int) int { a,b:=m[0]; if b { return a }; return 0 }
func pair() (int,int) { return 1,2 }
func use() int { a,b:=pair();return a+b }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	p = i.funcs[i.mainpkg.Func("use")]
	for pc, instr := range p.ssaInstrs {
		call, ok := instr.(*ssa.Call)
		if !ok {
			continue
		}
		saved := p.elidedExtract
		p.elidedExtract = nil
		desc := describeInstruction(p, call, call.Operands(nil), nil)
		p.elidedExtract = saved
		if desc.Aux.(callAux).ExtractResults != nil {
			t.Fatalf("pc %d: unelided call has extract destinations", pc)
		}
		return
	}
	t.Fatal("missing call")
}
