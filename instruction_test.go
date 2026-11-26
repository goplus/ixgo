//go:build instruction

package ixgo

import (
	"fmt"
	"go/constant"
	"go/parser"
	"go/token"
	"go/types"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"

	"golang.org/x/tools/go/ssa"
)

// instructionCallMetadata returns the common metadata for each call aux type.
func instructionCallMetadata(aux instructionAux) callAux {
	switch aux := aux.(type) {
	case dynamicCallAux:
		return aux.callAux
	case directCallAux:
		return aux.callAux
	case externalCallAux:
		return aux.callAux
	case bodyCallAux:
		return aux.callAux
	case callAux:
		return aux
	}
	panic("missing call metadata")
}

func TestDescriptionPreservesFrameLayout(t *testing.T) {
	for _, config := range []struct{ extra, descriptions Mode }{
		{0, EnableInstructionDescriptions}, {0, EnableInstructionExecution}, {0, EnableInstructionDescriptions | EnableInstructionExecution},
		{ExperimentalSupportGC, EnableInstructionDescriptions},
		{ExperimentalSupportGC, EnableInstructionExecution}, {EnableCachedReg, EnableInstructionExecution},
		{ExperimentalSupportGC | EnableGCExecutionPlan, EnableInstructionDescriptions},
		{ExperimentalSupportGC | EnableGCExecutionPlan, EnableInstructionExecution},
	} {
		extra := config.extra
		t.Run(fmt.Sprintf("mode=%v/descriptions=%v", extra, config.descriptions), func(t *testing.T) {
			load := func(mode Mode) *Interp {
				source := strings.Replace(executionDescriptionSource, "package main", "package main\nimport \"runtime\"", 1)
				source += `
func gcPhi(n int) int { a,b:=new(int),new(int); *a=7; *b=9; p:=a; if n>0 { p=b }; runtime.GC(); return *p }
`
				i, err := NewContext(DisableMemo|extra|mode).LoadInterp("description_layout.go", source)
				if err != nil {
					t.Fatal(err)
				}
				t.Cleanup(i.UnsafeRelease)
				return i
			}
			baseline, described := load(0), load(config.descriptions)
			byName := make(map[string]*function)
			for fn, p := range baseline.funcs {
				byName[fn.String()] = p
			}
			gcFunctions := 0
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
				if len(p.Code) != len(p.Instrs) || len(p.Instrs) != len(old.Instrs) ||
					!reflect.DeepEqual(p.Blocks, old.Blocks) {
					t.Fatalf("%s execution PC alignment changed: Code=%d, closures=%d/%d, blocks=%v/%v", fn, len(p.Code), len(p.Instrs), len(old.Instrs), p.Blocks, old.Blocks)
				}
				for pc := range p.Instrs {
					wantPos := p.ssaInstrs[pc].Pos()
					if _, ok := p.ssaInstrs[pc].(*ssa.RunDefers); ok && p.Fn.Syntax() != nil {
						wantPos = p.Fn.Syntax().End()
					}
					if p.PosForPC(pc) != wantPos {
						t.Fatalf("%s source position misaligned at PC %d", fn, pc)
					}
					// Generated builtin declarations may have a different source order.
					if pos := p.PositionForPC(pc); pos.Filename == "description_layout.go" && pos != old.PositionForPC(pc) {
						t.Fatalf("%s source position changed at PC %d: %v/%v", fn, pc, pos, old.PositionForPC(pc))
					}
				}
				if len(p.gcRegs) != len(old.gcRegs) {
					t.Fatalf("%s GC PC count changed", fn)
				}
				if len(p.gcRegs) > 0 {
					gcFunctions++
				}
				for pc := range p.gcRegs {
					if !reflect.DeepEqual(p.gcRegsForPC(pc), old.gcRegsForPC(pc)) {
						t.Fatalf("%s GC registers changed at PC %d", fn, pc)
					}
				}
			}
			if extra&ExperimentalSupportGC != 0 && (!described.gcEnabled || gcFunctions == 0) {
				t.Fatal("GC checks did not run")
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
	expected := make(map[reflect.Type]instructionOpcode)
	for _, test := range tests {
		expected[reflect.TypeOf(test.instr)] = test.op
	}
	ctx := NewContext(DisableMemo | EnableInstructionDescriptions)
	ctx.SetDebug(func(*DebugInfo) {})
	source := executionDescriptionSource + `
type opcodeWide interface { A(); B() }
type opcodeNarrow interface { A() }
func opcodeNarrower(v opcodeWide) opcodeNarrow { return v }
type opcodeRecord struct { A int; B string }
func opcodeIdentity(v opcodeRecord) opcodeRecord { return v }
func opcodeField(v opcodeRecord) string { return opcodeIdentity(v).B }
func opcodeContainers(n int,s string,p *opcodeRecord) int {
 p.A=n
 m:=make(map[int]int,n);m[n]=n
 a:=make([]int,n,n+1)
 a=a[:n]
 return len(a)+len(m)+int(s[n])
}
func opcodeExtra(m map[string]int, s string, a []int, x any, c chan int) int {
 m["x"] = 1
 n,ok := m["x"]
 if ok { n++ }
 for _,v := range m { n += v }
 for _,r := range s { n += int(r) }
 if v,ok := x.(int);ok { n += v }
 p := (*[2]int)(a)
 n += p[0]
 select {case v := <-c: n += v; default:}
 return n
}
`
	i, err := ctx.LoadInterp("opcode_classification.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	seen := make(map[instructionOpcode]bool)
	for _, p := range i.funcs {
		for pc, d := range p.Code {
			instr := p.ssaInstrs[pc]
			want, ok := expected[reflect.TypeOf(instr)]
			if !ok || d.Op != want {
				t.Fatalf("%T: opcode %v, want %v", instr, d.Op, want)
			}
			seen[d.Op] = true
		}
	}
	for _, test := range tests {
		if !seen[test.op] {
			t.Errorf("%T: opcode %v not exercised", test.instr, test.op)
		}
	}
	if describeInstruction(&function{}, nil, nil, nil).Op != opInvalid {
		t.Fatal("unknown instruction must be invalid")
	}
}

func TestInstructionAuxVariants(t *testing.T) {
	variants := []instructionAux{
		binaryOpAux{},
		unaryOpAux{},
		phiAux{},
		jumpAux{},
		ifAux{},
		returnAux{},
		extractAux{},
		nextAux{},
		callAux{},
		directCallAux{},
		externalCallAux{},
		bodyCallAux{},
		callSiteAux{},
		fieldAux{},
		lookupAux{},
		typeAssertAux{},
		conversionAux{},
		creationAux{},
		selectAux{},
		sendAux{},
		storeAux{},
		mapUpdateAux{},
		panicAux{},
		closureAux{},
		indexAux{},
		sliceAux{},
		rangeAux{},
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
	fn := &ssa.Function{Signature: types.NewSignatureType(nil, nil, nil, types.NewTuple(), types.NewTuple(), false)}
	p.index[fn] = 7 | uint32(kindFunction)<<30
	if got := descriptionRegister(p, fn); got != 7 {
		t.Fatalf("register = %d", got)
	}
	if descriptionRegister(p, nil) != -1 || descriptionRegister(p, &ssa.Function{}) != -1 {
		t.Fatal("unallocated operand must use -1")
	}
	ctx := NewContext(DisableMemo)
	i, err := ctx.LoadInterp("lookup.go", "package main; var Global int")
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	p.Interp = i
	x := ssa.NewConst(constant.MakeInt64(3), types.Typ[types.Int])
	before := len(p.index)
	for _, instr := range []ssa.Instruction{
		&ssa.BinOp{Op: token.ADD, X: x, Y: x},
		&ssa.UnOp{Op: token.SUB, X: x},
	} {
		d := describeInstruction(p, instr, instr.Operands(nil), nil)
		if d.Result != -1 || len(p.index) != before || len(p.stack) != 0 {
			t.Fatal("description allocated a register")
		}
		for j, r := range d.Args {
			if r != -1 || d.ArgKinds[j] != operandUnallocated {
				t.Fatal("unallocated operand metadata")
			}
		}
		switch aux := d.Aux.(type) {
		case binaryOpAux:
			if aux.Constant != [2]bool{} {
				t.Fatal("unallocated constant read")
			}
		case unaryOpAux:
			if aux.Constant || aux.Global {
				t.Fatal("unallocated unary value read")
			}
		}
	}

	global := i.mainpkg.Members["Global"].(*ssa.Global)
	instr := &ssa.UnOp{Op: token.MUL, X: global}
	d := describeInstruction(p, instr, instr.Operands(nil), nil)
	aux := d.Aux.(unaryOpAux)
	if aux.Global || aux.Constant || d.Args[0] != -1 || len(p.index) != before || len(p.stack) != 0 {
		t.Fatal("unallocated global read or allocated")
	}
	p.index[x] = uint32(kindConst) << 30
	p.stack = []value{3}
	allocated := describeInstruction(p, &ssa.BinOp{Op: token.ADD, X: x, Y: x}, nil, nil).Aux.(binaryOpAux)
	if allocated.Constant != [2]bool{true, true} || allocated.Values != [2]value{3, 3} {
		t.Fatalf("allocated constant: %#v", allocated)
	}
	ptr := new(int)
	p.index[global] = uint32(kindGlobal)<<30 | 1
	p.stack = append(p.stack, ptr)
	allocatedGlobal := describeInstruction(p, instr, instr.Operands(nil), nil).Aux.(unaryOpAux)
	if !allocatedGlobal.Global || allocatedGlobal.Constant || allocatedGlobal.Value != ptr {
		t.Fatalf("allocated global: %#v", allocatedGlobal)
	}

	p.index[global] = uint32(kindGlobal)<<30 | 4
	if k, v := descriptionOperand(p, global); k != kindInvalid || v != nil {
		t.Fatal("out-of-range operand read")
	}
}

func TestDescriptionPhiGroup(t *testing.T) {
	for _, descriptions := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		t.Run(fmt.Sprint(descriptions), func(t *testing.T) {
			mode := DisableMemo | descriptions
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
						continue
					}
					if false {
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
				// Execute the selected closure with only the grouped description available
				// to the test oracle. Cyclic inputs must be read before any result is stored.
				initial := append([]value(nil), p.stack...)
				for k, result := range aux.Results {
					initial[result] = 100 + k
				}
				for j, pred := range aux.Preds {
					fr := &frame{pfn: &function{}, pred: pred, stack: append([]value(nil), initial...)}
					original := aux.Edges[0][j]
					if descriptions&EnableInstructionExecution != 0 {
						// The installed description executor must read the description's edges.
						for _, result := range aux.Results {
							if result != original {
								aux.Edges[0][j] = result
								break
							}
						}
					}
					want := make([]value, len(aux.Results))
					for k := range want {
						want[k] = fr.reg(aux.Edges[k][j])
					}
					p.Instrs[pc](fr)
					aux.Edges[0][j] = original
					for k, result := range aux.Results {
						if !reflect.DeepEqual(fr.reg(result), want[k]) {
							t.Fatalf("parallel assignment %d: got %v, want %v", k, fr.reg(result), want[k])
						}
					}
				}
			}
			if descriptions != 0 && groups != 1 {
				t.Fatalf("Phi group count = %d, want 1", groups)
			}
			if descriptions == 0 && len(p.Code) != 0 {
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
			aux := instructionCallMetadata(d.Aux)
			if false {
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
			aux := instructionCallMetadata(d.Aux)
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
func single(n int) int { return n+1 }
func pair(n int) (int,int) { return n,n+1 }
func use(n int) int { a,b:=pair(n); return single(a)+b }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	for fn, p := range i.funcs {
		if p.callDests != nil {
			t.Fatalf("%s retained temporary call destinations", fn)
		}
	}
	p = i.funcs[i.mainpkg.Func("use")]
	var ordinary, elided int
	for pc, instr := range p.ssaInstrs {
		call, ok := instr.(*ssa.Call)
		if !ok {
			continue
		}
		desc := p.Code[pc]
		aux := instructionCallMetadata(desc.Aux)
		switch call.Call.StaticCallee().Name() {
		case "single":
			ordinary++
			if aux.ExtractResults != nil || desc.Result < 0 || desc.Result != descriptionRegister(p, call) {
				t.Fatalf("ordinary result = %#v", desc)
			}
		case "pair":
			elided++
			if desc.Result != -1 || len(aux.ExtractResults) != 2 {
				t.Fatalf("elided result = %#v", desc)
			}
			for _, ref := range *call.Referrers() {
				if ex, ok := ref.(*ssa.Extract); ok {
					if aux.ExtractResults[ex.Index] < 0 || aux.ExtractResults[ex.Index] != descriptionRegister(p, ex) {
						t.Fatal("cleanup lost extract destination")
					}
				}
			}
		}
	}
	if ordinary != 1 || elided != 1 {
		t.Fatalf("call counts = %d/%d, want 1/1", ordinary, elided)
	}
	if got, err := i.RunFunc("use", 3); err != nil || got != 8 {
		t.Fatalf("use = %v, %v; want 8", got, err)
	}
}

func TestCallDestinationsUnsafeRelease(t *testing.T) {
	p := &function{}
	p.callDests = map[ssa.Value][]register{&ssa.Function{}: {1, 2}}
	p.UnsafeRelease()
	if p.callDests != nil {
		t.Fatal("UnsafeRelease retained temporary call destinations")
	}
}

// Snapshot of the closure executor layout before optional instruction metadata.
type closureExecutionLayout struct {
	Instrs           []func(*frame)
	Recover          []func(*frame)
	Blocks           []int
	stack            []value
	pool             *sync.Pool
	base             int
	nres, narg, nenv int
	used, cached     int32
	cacheRegs        []register
	gcRegs           [][]register
	gcReady          []atomic.Bool
	gcPlan           *gcLivenessPlan
	gcMu             sync.Mutex
	clearRanges      []stackRange
	localRefs        []register
	hasGCRegs        bool
	frameCacheSafe   bool
	childFrameReg    register
	memoFlags        uint8
	metadata         functionMetadata
}

func TestInstructionDescriptionsColdLayout(t *testing.T) {
	want := reflect.TypeOf(closureExecutionLayout{})
	got := reflect.TypeOf(functionExecution{})
	for i := 0; i < want.NumField(); i++ {
		field := want.Field(i)
		actual, ok := got.FieldByName(field.Name)
		if !ok || actual.Offset != field.Offset || actual.Type != field.Type {
			t.Fatalf("closure field %s changed: got %v, want %v", field.Name, actual, field)
		}
	}
}

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
				if !enabled && p.callDests != nil {
					t.Fatal("default closure backend retained temporary call destinations")
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

func TestDescriptionValueMetadata(t *testing.T) {
	const source = `package main
 type Inner struct { A int; B string }
 type Outer struct { Padding byte; Inner }
 type Named int
 func identity(v Inner) Inner { return v }
 func field(v Inner) string { return identity(v).B }
 func address(p *Outer) int { p.A++; return p.A }
 func lookup(m map[string]int) int { return m["key"] }
 func lookupOK(m map[string]int) bool { _,ok:=m["key"];return ok }
 func assert(x any) int { return x.(int) }
 func assertOK(x any) bool { _,ok:=x.(Named);return ok }
 func assertInterface(x any) bool { _,ok:=x.(interface{ String() string });return ok }
 func run() int { p:=Outer{Inner:Inner{A:3,B:"value"}}; if field(p.Inner)!="value" { panic("field") }; return address(&p) }
 func named() bool { return assertOK(Named(1)) }
 `
	for _, mode := range []Mode{0, EnableCachedReg, ExperimentalSupportGC, ExperimentalSupportGC | EnableGCExecutionPlan} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			i, err := NewContext(DisableMemo|EnableInstructionDescriptions|mode).LoadInterp("value_metadata.go", source)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			counts := map[instructionOpcode]int{}
			for _, p := range i.funcs {
				for pc, d := range p.Code {
					switch v := p.ssaInstrs[pc].(type) {
					case *ssa.Field:
						aux, ok := d.Aux.(fieldAux)
						if !ok || aux.Index != v.Field {
							t.Fatalf("field = %#v", d)
						}
						counts[d.Op]++
					case *ssa.FieldAddr:
						aux, ok := d.Aux.(fieldAux)
						if !ok || aux.Index != v.Field {
							t.Fatalf("field address = %#v", d)
						}
						counts[d.Op]++
					case *ssa.Lookup:
						aux, ok := d.Aux.(lookupAux)
						if !ok || aux.CommaOk != v.CommaOk {
							t.Fatalf("lookup = %#v", d)
						}
						counts[d.Op]++
					case *ssa.TypeAssert:
						aux, ok := d.Aux.(typeAssertAux)
						if !ok || aux.CommaOk != v.CommaOk || aux.Source != i.preToType(v.X.Type()) || aux.Target != i.preToType(v.AssertedType) {
							t.Fatalf("assertion = %#v", d)
						}
						if aux.Source.Kind() != reflect.Interface {
							t.Fatalf("assertion source = %v, want interface", aux.Source)
						}
						switch p.Fn.Name() {
						case "assert":
							if aux.Target != reflect.TypeOf(int(0)) || aux.CommaOk {
								t.Fatalf("int assertion = %#v", aux)
							}
						case "assertInterface":
							if aux.Target.Kind() != reflect.Interface || aux.Target.NumMethod() != 1 || !aux.CommaOk {
								t.Fatalf("interface assertion = %#v", aux)
							}
							method := aux.Target.Method(0)
							if method.Name != "String" || method.Type.NumIn() != 0 || method.Type.NumOut() != 1 || method.Type.Out(0) != reflect.TypeOf("") {
								t.Fatalf("interface method = %v, want String() string", method)
							}
						}
						counts[d.Op]++
					default:
						continue
					}
					if d.Result < 0 || len(d.Args) == 0 || d.Args[0] < 0 {
						t.Fatalf("missing register operands: %#v", d)
					}
				}
			}
			for _, op := range []instructionOpcode{opField, opFieldAddr, opLookup, opTypeAssert} {
				if counts[op] == 0 {
					t.Fatalf("missing opcode %d", op)
				}
			}
			cases := []struct {
				name string
				args []value
				want value
			}{
				{"run", nil, 4}, {"named", nil, true},
				{"lookup", []value{map[string]int{"key": 7}}, 7},
				{"lookup", []value{map[string]int(nil)}, 0},
				{"lookupOK", []value{map[string]int{}}, false},
				{"lookupOK", []value{map[string]int{"key": 0}}, true},
				{"assert", []value{42}, 42},
				{"assertOK", []value{1}, false},
				{"assertInterface", []value{1}, false},
			}
			for _, c := range cases {
				got, err := i.RunFunc(c.name, c.args...)
				if err != nil || !reflect.DeepEqual(got, c.want) {
					t.Fatalf("%s = %v, %v; want %v", c.name, got, err, c.want)
				}
			}
			if _, err := i.RunFunc("assert", "wrong"); err == nil {
				t.Fatal("failed assertion should return error")
			}
		})
	}
}

func TestDescriptionTypeMetadata(t *testing.T) {
	const source = `package main
 type Named int
 type Wide interface { A(); B() }
 type Narrow interface { A() }
 func narrow(x Wide) Narrow { return x }
 func named(n int) Named { return Named(n) }
 func number(n int) int64 { return int64(n) }
 func box(n int) any { return n }
 func bytes(s string) []byte { return []byte(s) }
 func text(b []byte) string { return string(b) }
 func pointer(s []int) *[2]int { return (*[2]int)(s) }
 func heap(n int) *int { p:=new(int);*p=n;return p }
 func local(n int) int { p:=new(int);*p=n;return *p }
 func slice(n int) int { s:=make([]int,n,n+1);return cap(s) }
 func mapping(n int) int { m:=make(map[string]int,n);m["x"]=n;return m["x"] }
 func channel(n int) int { c:=make(chan int,n);return cap(c) }
 `
	for _, mode := range []Mode{0, EnableCachedReg, ExperimentalSupportGC, ExperimentalSupportGC | EnableGCExecutionPlan} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			i, err := NewContext(DisableMemo|EnableInstructionDescriptions|mode).LoadInterp("type_metadata.go", source)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			counts := map[instructionOpcode]int{}
			for _, p := range i.funcs {
				for pc, d := range p.Code {
					switch aux := d.Aux.(type) {
					case conversionAux:
						if aux.Source == nil || aux.Target == nil {
							t.Fatalf("missing conversion types: %#v", d)
						}
						switch p.Fn.Name() {
						case "number":
							if aux.Source != reflect.TypeOf(0) || aux.Target != reflect.TypeOf(int64(0)) {
								t.Fatalf("numeric types = %#v", aux)
							}
						case "named":
							if aux.Source.Kind() != reflect.Int || aux.Target.Kind() != reflect.Int {
								t.Fatalf("named types = %#v", aux)
							}
						case "box":
							if aux.Source != reflect.TypeOf(0) || aux.Target.Kind() != reflect.Interface {
								t.Fatalf("box types = %#v", aux)
							}
						case "narrow":
							if aux.Source.Kind() != reflect.Interface || aux.Source.NumMethod() != 2 || aux.Target.Kind() != reflect.Interface || aux.Target.NumMethod() != 1 {
								t.Fatalf("interface types = %#v", aux)
							}
						case "bytes":
							if aux.Source != reflect.TypeOf("") || aux.Target != reflect.TypeOf([]byte(nil)) {
								t.Fatalf("byte conversion = %#v", aux)
							}
						case "text":
							if aux.Source != reflect.TypeOf([]byte(nil)) || aux.Target != reflect.TypeOf("") {
								t.Fatalf("string conversion = %#v", aux)
							}
						case "pointer":
							if aux.Source != reflect.TypeOf([]int(nil)) || aux.Target != reflect.TypeOf((*[2]int)(nil)) {
								t.Fatalf("array pointer conversion = %#v", aux)
							}
						}
					case creationAux:
						if aux.Target == nil {
							t.Fatal("missing creation type")
						}
						switch v := p.ssaInstrs[pc].(type) {
						case *ssa.Alloc:
							if aux.Target.Kind() != reflect.Pointer || aux.Heap != v.Heap {
								t.Fatalf("allocation = %#v", aux)
							}
						case *ssa.MakeSlice:
							if aux.Target != reflect.TypeOf([]int(nil)) || len(d.Args) != 2 {
								t.Fatalf("slice = %#v", d)
							}
						case *ssa.MakeMap:
							if aux.Target != reflect.TypeOf(map[string]int(nil)) || len(d.Args) != 1 {
								t.Fatalf("map = %#v", d)
							}
						case *ssa.MakeChan:
							if aux.Target != reflect.TypeOf((chan int)(nil)) || len(d.Args) != 1 {
								t.Fatalf("chan = %#v", d)
							}
						default:
							t.Fatalf("unexpected creation opcode %T", v)
						}
					default:
						continue
					}
					counts[d.Op]++
					if d.Result < 0 {
						t.Fatalf("missing result: %#v", d)
					}
				}
			}
			for _, op := range []instructionOpcode{opChangeType, opConvert, opChangeInterface, opMakeInterface, opAlloc, opMakeSlice, opMakeMap, opMakeChan, opSliceToArrayPointer} {
				if counts[op] == 0 {
					t.Fatalf("missing metadata opcode %d", op)
				}
			}
			cases := []struct {
				name string
				args []value
				want value
			}{
				{"number", []value{3}, int64(3)}, {"box", []value{3}, 3},
				{"bytes", []value{"ab"}, []byte("ab")}, {"text", []value{[]byte(nil)}, ""},
				{"local", []value{4}, 4}, {"slice", []value{0}, 1}, {"slice", []value{3}, 4},
				{"mapping", []value{0}, 0}, {"channel", []value{0}, 0}, {"channel", []value{2}, 2},
			}
			for _, c := range cases {
				got, err := i.RunFunc(c.name, c.args...)
				if err != nil || !reflect.DeepEqual(got, c.want) {
					t.Fatalf("%s = %v, %v; want %v", c.name, got, err, c.want)
				}
			}
			if got, err := i.RunFunc("heap", 7); err != nil || *got.(*int) != 7 {
				t.Fatalf("heap = %v, %v", got, err)
			}
			if got, err := i.RunFunc("pointer", []int{1, 2}); err != nil || *got.(*[2]int) != [2]int{1, 2} {
				t.Fatalf("pointer = %v, %v", got, err)
			}
			for _, name := range []string{"slice", "channel"} {
				if _, err := i.RunFunc(name, -1); err == nil {
					t.Fatalf("%s negative size should fail", name)
				}
			}
			if _, err := i.RunFunc("pointer", []int{1}); err == nil {
				t.Fatal("short slice should fail")
			}
		})
	}
}

func TestDescriptionBoundsMetadata(t *testing.T) {
	const source = `package main
 type Record struct { Pad byte; Value int64 }
 type Named []int
 func record(p *Record) int64 { return p.Value }
 func array(a [3]int, n int) int { return a[n] }
 func index(s []int,n int) int { return s[n] }
 func text(s string,n int) byte { return s[n] }
 func all(s []int) []int { return s[:] }
 func full(s []int,a,b,c int) []int { return s[a:b:c] }
 func named(s Named) Named { return s[:] }
 func allocated(n int) []int { a := [3]int{1,2,3}; return a[:n] }
 func ranges(m map[string]int,s string) int { n:=0;for _,v:=range m {n+=v};for _,r:=range s {n+=int(r)};return n }
 func lookup(m map[string]int) int { return m["x"] }
 `
	for _, mode := range []Mode{0, EnableCachedReg, ExperimentalSupportGC, ExperimentalSupportGC | EnableGCExecutionPlan} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			i, err := NewContext(DisableMemo|EnableInstructionDescriptions|mode).LoadInterp("bounds_metadata.go", source)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			var fields, indexes, slices, ranges, lookups int
			for _, p := range i.funcs {
				for _, d := range p.Code {
					switch aux := d.Aux.(type) {
					case fieldAux:
						fields++
						typ := aux.Source
						if typ.Kind() == reflect.Pointer {
							typ = typ.Elem()
						}
						f := typ.Field(aux.Index)
						if aux.Field != f.Type || aux.Offset != f.Offset {
							t.Fatal("field layout mismatch")
						}
						if f.Name == "Value" && (aux.Field != reflect.TypeOf(int64(0)) || aux.Offset == 0) {
							t.Fatal("missing padded field layout")
						}
					case indexAux:
						indexes++
						typ := aux.Source
						if typ.Kind() == reflect.Pointer {
							typ = typ.Elem()
						}
						if typ.Kind() == reflect.Array && typ.Len() != 3 {
							t.Fatal("array length mismatch")
						}
						target := aux.Target
						if target.Kind() == reflect.Pointer {
							target = target.Elem()
						}
						element := reflect.TypeOf(byte(0))
						if typ.Kind() != reflect.String {
							element = typ.Elem()
						}
						if target != element {
							t.Fatal("index element mismatch")
						}
					case sliceAux:
						slices++
						if len(d.Args) != 4 {
							t.Fatal("slice operand positions changed")
						}
						for j, present := range []bool{aux.Low, aux.High, aux.Max} {
							if present != (d.ArgKinds[j+1] != operandAbsent) {
								t.Fatal("slice omitted boundary mismatch")
							}
						}
						if p.Fn.Name() == "all" && (aux.Low || aux.High || aux.Max) {
							t.Fatal("default boundaries present")
						}
						if p.Fn.Name() == "full" && (!aux.Low || !aux.High || !aux.Max) {
							t.Fatal("three-index boundaries missing")
						}
						if p.Fn.Name() == "allocated" && !aux.MakeSliceCheck {
							t.Fatal("allocation check missing")
						}
						if aux.Target.Kind() != reflect.Slice {
							t.Fatal("slice target missing")
						}
					case rangeAux:
						ranges++
						if aux.Source.Kind() == reflect.Map {
							if aux.Key != reflect.TypeOf("") || aux.Element != reflect.TypeOf(0) {
								t.Fatal("map iterator types")
							}
						} else if aux.Source.Kind() != reflect.String || aux.Key != reflect.TypeOf(0) || aux.Element != reflect.TypeOf(rune(0)) {
							t.Fatal("string iterator types")
						}
					case lookupAux:
						lookups++
						if aux.Source.Kind() == reflect.Map {
							if aux.Key != reflect.TypeOf("") || aux.Element != reflect.TypeOf(0) {
								t.Fatal("map lookup types")
							}
						} else if aux.Source.Kind() != reflect.String || aux.Element != reflect.TypeOf(byte(0)) {
							t.Fatal("string lookup types")
						}
					}
				}
			}
			if fields == 0 || indexes == 0 || slices == 0 || ranges != 2 || lookups == 0 {
				t.Fatalf("missing metadata: %d/%d/%d/%d/%d", fields, indexes, slices, ranges, lookups)
			}
			cases := []struct {
				name string
				args []value
				want value
			}{
				{"array", []value{[3]int{1, 2, 3}, 2}, 3}, {"index", []value{[]int{4}, 0}, 4}, {"text", []value{"ab", 1}, byte('b')},
				{"all", []value{[]int(nil)}, []int(nil)}, {"full", []value{[]int{1, 2, 3}, 1, 2, 2}, []int{2}},
				{"allocated", []value{2}, []int{1, 2}}, {"ranges", []value{map[string]int{"x": 2}, "a"}, 99},
			}
			for _, c := range cases {
				got, err := i.RunFunc(c.name, c.args...)
				if err != nil || !reflect.DeepEqual(got, c.want) {
					t.Fatalf("%s = %v, %v; want %v", c.name, got, err, c.want)
				}
			}
			failures := []struct {
				name string
				args []value
			}{
				{"index", []value{[]int(nil), 0}}, {"index", []value{[]int{1}, -1}}, {"array", []value{[3]int{}, 3}},
				{"text", []value{"", 0}}, {"full", []value{[]int{1}, 0, 1, 2}}, {"full", []value{[]int{1}, 1, 0, 1}}, {"allocated", []value{4}},
			}
			for _, c := range failures {
				if _, err := i.RunFunc(c.name, c.args...); err == nil {
					t.Fatalf("%s invalid bounds should fail", c.name)
				}
			}
		})
	}
}

func TestDescriptionStringLookup(t *testing.T) {
	i, err := NewContext(DisableMemo|EnableInstructionDescriptions).LoadInterp("string_lookup.go", "package main; func f() {}")
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	p := i.funcs[i.mainpkg.Func("f")]
	// Older SSA producers may use Lookup for string indexing.
	instr := &ssa.Lookup{X: ssa.NewConst(constant.MakeString("ab"), types.Typ[types.String]), Index: ssa.NewConst(constant.MakeInt64(1), types.Typ[types.Int])}
	d := describeInstruction(p, instr, instr.Operands(nil), nil)
	aux := d.Aux.(lookupAux)
	if aux.Source != reflect.TypeOf("") || aux.Key != reflect.TypeOf(0) || aux.Element != reflect.TypeOf(byte(0)) || aux.CommaOk {
		t.Fatalf("string lookup = %#v", aux)
	}
}

func TestDescriptionSideEffectMetadata(t *testing.T) {
	const source = `package main
func selectRecv(c chan int) int { select { case v := <-c: return v; default: return 0 } }
func selectSend(c chan int, v int) { select { case c <- v: default: } }
func send(c chan int, v int) { c <- v }
func store(p *int, v int) { *p = v }
func update(m map[string]int, v int) { m["x"] = v }
func panicf(x any) { panic(x) }
`
	for _, mode := range []Mode{0, EnableCachedReg, ExperimentalSupportGC, ExperimentalSupportGC | EnableGCExecutionPlan} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			i, err := NewContext(DisableMemo|EnableInstructionDescriptions|mode).LoadInterp("side_effect_metadata.go", source)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			counts := map[instructionOpcode]int{}
			var sendStates, recvStates int
			for _, p := range i.funcs {
				for pc, d := range p.Code {
					switch v := p.ssaInstrs[pc].(type) {
					case *ssa.Select:
						aux, ok := d.Aux.(selectAux)
						if !ok || len(aux.States) != len(v.States) || aux.Blocking != v.Blocking {
							t.Fatalf("select = %#v", d)
						}
						for j, state := range v.States {
							wantDir := reflect.SelectRecv
							if state.Dir == types.SendOnly {
								wantDir = reflect.SelectSend
							}
							if aux.States[j].Chan != i.preToType(state.Chan.Type()) || aux.States[j].Dir != wantDir {
								t.Fatalf("select state = %#v", aux.States[j])
							}
							if state.Send == nil && aux.States[j].Send != nil {
								t.Fatal("receive has send type")
							}
							if state.Send != nil && aux.States[j].Send != i.preToType(state.Send.Type()) {
								t.Fatal("send type mismatch")
							}
							if wantDir == reflect.SelectSend {
								if aux.States[j].Send == nil {
									t.Fatal("missing select send value type")
								}
								sendStates++
							} else {
								recvStates++
							}
						}
					case *ssa.Send:
						aux, ok := d.Aux.(sendAux)
						if !ok || aux.Chan != i.preToType(v.Chan.Type()) || aux.Value != i.preToType(v.X.Type()) {
							t.Fatalf("send = %#v", d)
						}
					case *ssa.Store:
						aux, ok := d.Aux.(storeAux)
						if !ok || aux.Address != i.preToType(v.Addr.Type()) || aux.Value != i.preToType(v.Val.Type()) {
							t.Fatalf("store = %#v", d)
						}
					case *ssa.MapUpdate:
						aux, ok := d.Aux.(mapUpdateAux)
						if !ok || aux.Map != i.preToType(v.Map.Type()) || aux.Key != i.preToType(v.Key.Type()) || aux.Value != i.preToType(v.Value.Type()) {
							t.Fatalf("map update = %#v", d)
						}
					case *ssa.Panic:
						aux, ok := d.Aux.(panicAux)
						if !ok || aux.Value != i.preToType(v.X.Type()) {
							t.Fatalf("panic = %#v", d)
						}
					default:
						continue
					}
					counts[d.Op]++
				}
			}
			for _, op := range []instructionOpcode{opSelect, opSend, opStore, opMapUpdate, opPanic} {
				if counts[op] == 0 {
					t.Fatalf("missing metadata opcode %d", op)
				}
			}
			if sendStates == 0 || recvStates == 0 {
				t.Fatalf("select directions not covered: send=%d receive=%d", sendStates, recvStates)
			}
			c := make(chan int, 1)
			if _, err := i.RunFunc("send", c, 7); err != nil {
				t.Fatal(err)
			}
			if got := <-c; got != 7 {
				t.Fatalf("sent %d", got)
			}
			m := map[string]int{}
			if _, err := i.RunFunc("update", m, 9); err != nil {
				t.Fatal(err)
			}
			if m["x"] != 9 {
				t.Fatalf("map update = %v", m)
			}
			v := 0
			if _, err := i.RunFunc("store", &v, 5); err != nil {
				t.Fatal(err)
			}
			if v != 5 {
				t.Fatalf("store = %d", v)
			}
			if got, err := i.RunFunc("selectRecv", make(chan int)); err != nil || got != 0 {
				t.Fatalf("select recv = %v, %v", got, err)
			}
			if _, err := i.RunFunc("selectSend", c, 11); err != nil {
				t.Fatal(err)
			}
			select {
			case got := <-c:
				if got != 11 {
					t.Fatalf("select sent %d, want 11", got)
				}
			default:
				t.Fatal("select send did not enqueue value")
			}
			// A full buffer must take the default arm without replacing its value.
			c <- 12
			if _, err := i.RunFunc("selectSend", c, 13); err != nil {
				t.Fatal(err)
			}
			if got := <-c; got != 12 {
				t.Fatalf("default select replaced buffered value with %d", got)
			}
			c <- 14
			if got, err := i.RunFunc("selectRecv", c); err != nil || got != 14 {
				t.Fatalf("ready select receive = %v, %v", got, err)
			}
			if _, err := i.RunFunc("panicf", "side-effect panic"); err == nil || !strings.Contains(err.Error(), "side-effect panic") {
				t.Fatalf("panic result = %v", err)
			}
		})
	}
}

func TestDescriptionClosureMetadata(t *testing.T) {
	const source = `package main
func noCapture(n int) int { f:=func(x int) int { return x+1 }; return f(n) }
func captures(a,b int) int { f:=func(x int) int { return x+a+b }; return f(3) }
func nested(a int) func(int) int { return func(x int) int { return x+a } }
func top() func(int) int { return func(x int) int { return x+1 } }
type Counter struct { Base int }
func (c Counter) Add(n int) int { return c.Base+n }
func method(n int) func(int) int { c:=Counter{Base:n};return c.Add }
`
	i, err := NewContext(DisableMemo|EnableInstructionDescriptions).LoadInterp("closure_metadata.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	var one, two, methods int
	for _, p := range i.funcs {
		for pc, d := range p.Code {
			v, ok := p.ssaInstrs[pc].(*ssa.MakeClosure)
			if !ok {
				continue
			}
			aux, ok := d.Aux.(closureAux)
			if !ok || aux.Signature == nil {
				t.Fatalf("closure metadata = %#v", d)
			}
			fn, ok := v.Fn.(*ssa.Function)
			if !ok || fn.Signature == nil {
				t.Fatalf("closure target = %T", v.Fn)
			}
			if aux.Signature != i.preToType(v.Fn.Type()) || aux.Signature.NumIn() != fn.Signature.Params().Len() || aux.Signature.NumOut() != fn.Signature.Results().Len() {
				t.Fatalf("closure signature = %#v", aux)
			}
			if len(aux.Bindings) != len(v.Bindings) || len(aux.Bindings) != len(fn.FreeVars) {
				t.Fatalf("closure bindings = %#v", aux)
			}
			for j, binding := range v.Bindings {
				if aux.Bindings[j] < 0 || aux.Bindings[j] != p.regIndex(binding) {
					t.Fatalf("binding %d = %v", j, aux.Bindings[j])
				}
			}
			if p.Fn.Name() == "method" {
				methods++
			}
			switch len(aux.Bindings) {
			case 1:
				one++
			case 2:
				two++
			}
		}
	}
	if one == 0 || two == 0 || methods != 1 {
		t.Fatalf("capture cases = one:%d two:%d methods:%d", one, two, methods)
	}
	// No-capture literals use static SSA functions, so these cases test runtime
	// behavior without exercising MakeClosure metadata.
	for _, c := range []struct {
		name string
		args []value
		want value
	}{
		{"noCapture", []value{4}, 5}, {"captures", []value{2, 4}, 9},
	} {
		if got, err := i.RunFunc(c.name, c.args...); err != nil || got != c.want {
			t.Fatalf("%s = %v, %v; want %v", c.name, got, err, c.want)
		}
	}
	for _, c := range []struct {
		name string
		args []value
		want int
	}{
		{"nested", []value{7}, 10}, {"top", nil, 4}, {"method", []value{8}, 11},
	} {
		got, err := i.RunFunc(c.name, c.args...)
		if err != nil {
			t.Fatal(err)
		}
		fn, ok := got.(func(int) int)
		if !ok {
			t.Fatalf("%s returned %T", c.name, got)
		}
		if result := fn(3); result != c.want {
			t.Fatalf("%s(3) = %d, want %d", c.name, result, c.want)
		}
	}
}

func TestDescriptionCallSiteMetadata(t *testing.T) {
	const source = `package main
 type Value int
 func (v Value) Add(n int) int { return int(v)+n }
 type Adder interface { Add(int) int }
 func static(n int) int { return n+1 }
 func dynamic(f func(int)int,n int)int { return f(n) }
 func invoke(v Adder,n int)int {return v.Add(n)}
 func ordinary()int {return static(2)+invoke(Value(4),3)+dynamic(static,1)}
 func direct()int {return Value(4).Add(1)}
 func bound()int {v:=Value(4);f:=v.Add;return f(3)}
 func async()int { c:=make(chan int,1);go func(){c<-8}();return <-c }
 func delayed()(r int){defer func(){r++}();return 4}
 func builtinLen()int {s:=[]int{1,2};return len(s)}
 func deferredBuiltin()int {c:=make(chan int);defer close(c);return 1}
 `
	for _, mode := range []Mode{0, EnableCachedReg, ExperimentalSupportGC, ExperimentalSupportGC | EnableGCExecutionPlan} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			i, err := NewContext(DisableMemo|EnableInstructionDescriptions|mode).LoadInterp("call_site.go", source)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			kinds := map[callKind]int{}
			ops := map[instructionOpcode]int{}
			var directCalls, boundCalls int
			for _, p := range i.funcs {
				for pc, d := range p.Code {
					var call *ssa.CallCommon
					var aux callSiteAux
					switch v := p.ssaInstrs[pc].(type) {
					case *ssa.Call:
						call = &v.Call
						base := instructionCallMetadata(d.Aux)
						aux = base.callSiteAux
					case *ssa.Go:
						call = &v.Call
						base := instructionCallMetadata(d.Aux)
						aux = base.callSiteAux
					case *ssa.Defer:
						call = &v.Call
						base := instructionCallMetadata(d.Aux)
						aux = base.callSiteAux
					default:
						continue
					}
					kinds[aux.Kind]++
					ops[d.Op]++
					if aux.Target != descriptionRegister(p, call.Value) || len(aux.Arguments) != len(call.Args) {
						t.Fatalf("call register layout = %#v", aux)
					}
					for j, arg := range call.Args {
						if aux.Arguments[j] != p.regIndex(arg) {
							t.Fatal("argument register mismatch")
						}
					}
					if b, ok := call.Value.(*ssa.Builtin); ok {
						if aux.Kind != callBuiltin || aux.Name != b.Name() || aux.Signature != nil {
							t.Fatalf("builtin metadata = %#v", aux)
						}
					} else {
						if aux.Signature == nil || aux.Signature.Kind() != reflect.Func {
							t.Fatalf("function signature = %#v", aux)
						}
						switch {
						case call.IsInvoke():
							if aux.Kind != callInvoke || aux.Name != call.Method.Id() || aux.Target < 0 {
								t.Fatalf("invoke = %#v", aux)
							}
						default:
							if aux.Signature.NumIn() != len(aux.Arguments) {
								t.Fatalf("%s: signature %v has %d inputs for %d arguments", p.Fn, aux.Signature, aux.Signature.NumIn(), len(aux.Arguments))
							}
							if p.Fn.Name() == "direct" {
								directCalls++
								if aux.Kind != callStatic || aux.Signature.NumIn() != 2 {
									t.Fatalf("direct method = %#v", aux)
								}
							}
							if p.Fn.Name() == "bound" {
								boundCalls++
								if aux.Kind != callDynamic || aux.Signature.NumIn() != 1 {
									t.Fatalf("bound method = %#v", aux)
								}
							}
							_, static := call.Value.(*ssa.Function)
							want := callDynamic
							if static {
								want = callStatic
							}
							if aux.Kind != want || aux.Name != "" {
								t.Fatalf("call kind = %#v", aux)
							}
						}
						if p.Fn.Name() == "dynamic" && (aux.Kind != callDynamic || aux.Signature != reflect.TypeOf(func(int) int { return 0 })) {
							t.Fatal("dynamic signature mismatch")
						}
					}
				}
			}
			if directCalls != 1 || boundCalls != 1 {
				t.Fatalf("method call coverage direct=%d bound=%d", directCalls, boundCalls)
			}
			for _, kind := range []callKind{callStatic, callDynamic, callBuiltin, callInvoke} {
				if kinds[kind] == 0 {
					t.Fatalf("missing kind %d", kind)
				}
			}
			for _, op := range []instructionOpcode{opCall, opGo, opDefer} {
				if ops[op] == 0 {
					t.Fatalf("missing call opcode %d", op)
				}
			}
			for _, c := range []struct {
				name string
				want int
			}{{"direct", 5}, {"bound", 7}, {"ordinary", 12}, {"async", 8}, {"delayed", 5}, {"builtinLen", 2}, {"deferredBuiltin", 1}} {
				if got, err := i.RunFunc(c.name); err != nil || got != c.want {
					t.Fatalf("%s = %v, %v; want %d", c.name, got, err, c.want)
				}
			}
		})
	}
}

func TestControlDescriptionExecution(t *testing.T) {
	const source = `package main
 import "runtime"
 type flag bool
 func branch(v flag) int { if v { return 7 }; return 9 }
 func named(n int) int { return branch(flag(n > 0)) }
 func constant(n int) int { for { if true { break } }; return n }
 func choosePointer(n int) int { a,b:=new(int),new(int); *a=7; *b=9; p:=a; if n>0 { p=b }; runtime.GC(); return *p }
 func swap(n int) int { a,b:=1,2; for j:=0;j<n;j++ { a,b=b,a }; return a*10+b }
 func choose(n int) int { x:=3; if n>0 { x=4 } else if n<0 { x=5 };return x }
 `
	for _, descriptionMode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		for _, extra := range []Mode{0, EnableCachedReg, ExperimentalSupportGC | EnableGCExecutionPlan} {
			t.Run(fmt.Sprintf("descriptions=%v/mode=%v", descriptionMode, extra), func(t *testing.T) {
				mode := DisableMemo | extra | descriptionMode
				i, err := NewContext(mode).LoadInterp("control_description.go", source)
				if err != nil {
					t.Fatal(err)
				}
				defer i.UnsafeRelease()
				gcFunctions := 0
				for _, p := range i.funcs {
					if len(p.gcRegs) > 0 {
						gcFunctions++
					}
					if descriptionMode != 0 {
						for pc, d := range p.Code {
							if d.Op != opJump && d.Op != opIf && d.Op != opPhi {
								continue
							}
							executor := makeInstruction(nil, d)
							if executor == nil {
								t.Fatalf("pc %d: control builder mismatch", pc)
							}
							if descriptionMode&EnableInstructionExecution != 0 {
								name := runtime.FuncForPC(reflect.ValueOf(p.Instrs[pc]).Pointer()).Name()
								helper := "makeInstruction"
								if got, want := strings.Contains(name, helper), descriptionMode&EnableInstructionExecution != 0; got != want {
									t.Fatalf("pc %d: installed %s, description executor=%t, want %t", pc, name, got, want)
								}
							}
						}
					}
					if descriptionMode == 0 && len(p.Code) != 0 {
						t.Fatal("default path generated descriptions")
					}
					if descriptionMode != 0 && len(p.Code) != len(p.Instrs) {
						t.Fatal("description PC count mismatch")
					}
					p.functionCompilation = functionCompilation{}
				}
				if extra&ExperimentalSupportGC != 0 && (!i.gcEnabled || gcFunctions == 0) {
					t.Fatal("GC plan was not built")
				}
				for _, c := range []struct {
					name string
					arg  int
					want int
				}{
					{"swap", 0, 12}, {"swap", 1, 21}, {"swap", 10, 12}, {"swap", 11, 21},
					{"choose", 0, 3}, {"choose", 1, 4}, {"choose", -1, 5},
					{"named", 0, 9}, {"named", 1, 7}, {"constant", 42, 42},
					{"choosePointer", 0, 7}, {"choosePointer", 1, 9},
				} {
					if got, err := i.RunFunc(c.name, c.arg); err != nil || got != c.want {
						t.Fatalf("%s(%d)=%v,%v;want %d", c.name, c.arg, got, err, c.want)
					}
				}
			})
		}
	}
}

func TestControlDescriptionBranchAndPhi(t *testing.T) {
	// A named bool must use the runtime boolean accessor.
	type flag bool
	for _, condition := range []value{true, false, flag(true), flag(false)} {
		d := typedInstruction{Op: opIf, Args: []register{0}, Aux: ifAux{Targets: [2]int32{1, 2}}}
		execute := makeInstruction(nil, d)
		p := &function{}
		p.Blocks = []int{0, 4, 8}
		fr := &frame{pfn: p, stack: []value{condition}}
		execute(fr)
		want := int32(8)
		if reflect.ValueOf(condition).Bool() {
			want = 4
		}
		if fr.ipc != want {
			t.Fatalf("condition %v: pc=%d want %d", condition, fr.ipc, want)
		}
	}
	jump := makeInstruction(nil, typedInstruction{Op: opJump, Aux: jumpAux{Target: 1}})
	if false {
		t.Fatal("missing Jump executor")
	}
	p := &function{}
	p.Blocks = []int{0, 4}
	jumpFrame := &frame{pfn: p, block: 0}
	jump(jumpFrame)
	if jumpFrame.ipc != 4 || jumpFrame.block != 1 || jumpFrame.pred != 0 {
		t.Fatalf("Jump state: pc=%d block=%d pred=%d", jumpFrame.ipc, jumpFrame.block, jumpFrame.pred)
	}
	d := typedInstruction{Op: opPhi, Aux: phiAux{Preds: []int32{1}, Results: []register{0, 1}, Edges: [][]register{{1}, {0}}}}
	execute := makeInstruction(nil, d)
	fr := &frame{pred: 1, stack: []value{1, 2}}
	execute(fr)
	if fr.reg(0) != 2 || fr.reg(1) != 1 {
		t.Fatal("Phi assignment is not parallel")
	}
	fr.pred = 99
	execute(fr)
	if fr.reg(0) != 2 || fr.reg(1) != 1 {
		t.Fatal("unknown predecessor changed registers")
	}
}

func BenchmarkControlDescription(b *testing.B) {
	for _, descriptions := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		mode := DisableMemo | descriptions
		b.Run(fmt.Sprintf("descriptions=%v", descriptions), func(b *testing.B) {
			b.Run("execute", func(b *testing.B) {
				ctx := NewContext(mode)
				ctx.SetLeastCallForEnablePool(0)
				i, err := ctx.LoadInterp("control_bench.go", executionDescriptionSource)
				if err != nil {
					b.Fatal(err)
				}
				b.Cleanup(i.UnsafeRelease)
				b.ReportAllocs()
				b.ResetTimer()
				for b.Loop() {
					if got, err := i.RunFunc("phi"); err != nil || got != 42 {
						b.Fatalf("phi = %v, %v", got, err)
					}
				}
			})
			b.Run("load", func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					i, err := NewContext(mode).LoadInterp("control_bench.go", executionDescriptionSource)
					if err != nil {
						b.Fatal(err)
					}
					i.UnsafeRelease()
				}
			})
		})
	}
}

func TestControlDescriptionIfSpecialization(t *testing.T) {
	for _, want := range []bool{false, true} {
		instr := &ssa.If{Cond: ssa.NewConst(constant.MakeBool(want), types.Typ[types.Bool])}
		d := describeInstruction(&function{}, instr, instr.Operands(nil), nil)
		aux := d.Aux.(ifAux)
		if !aux.Constant || aux.Condition != want || !aux.BasicBool {
			t.Fatalf("constant metadata: %#v", aux)
		}
		aux.Targets = [2]int32{1, 2}
		d.Aux = aux
		execute := makeInstruction(nil, d)
		p := &function{}
		p.Blocks = []int{0, 4, 8}
		fr := &frame{pfn: p}
		execute(fr)
		expected := int32(8)
		if want {
			expected = 4
		}
		if fr.ipc != expected {
			t.Fatalf("constant %t: PC %d, want %d", want, fr.ipc, expected)
		}
	}
	for _, basic := range []bool{false, true} {
		for _, condition := range []bool{false, true} {
			d := typedInstruction{Op: opIf, Args: []register{0}, Aux: ifAux{Targets: [2]int32{1, 2}, BasicBool: basic}}
			execute := makeInstruction(nil, d)
			p := &function{}
			p.Blocks = []int{0, 4, 8}
			fr := &frame{pfn: p, stack: []value{condition}}
			execute(fr)
			want := int32(8)
			if condition {
				want = 4
			}
			if fr.ipc != want {
				t.Fatalf("basic=%t condition=%t: PC %d want %d", basic, condition, fr.ipc, want)
			}
		}
	}
}

func BenchmarkControlDescriptionBranch(b *testing.B) {
	const source = `package main
 type flag bool
 func basic(v bool) int { if v { return 7 }; return 9 }
 func named(v flag) int { if v { return 7 }; return 9 }
 func loop() int { n:=0; for j:=0;j<100;j++ { if j&1==0 { n+=1 } else { n+=2 } }; return n }
 `
	for _, descriptions := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		mode := DisableMemo | descriptions
		for _, name := range []string{"basic", "named", "loop"} {
			b.Run(fmt.Sprintf("descriptions=%v/%s", descriptions, name), func(b *testing.B) {
				ctx := NewContext(mode)
				ctx.SetLeastCallForEnablePool(0)
				i, err := ctx.LoadInterp("branch_bench.go", source)
				if err != nil {
					b.Fatal(err)
				}
				b.Cleanup(i.UnsafeRelease)
				args := []value{true}
				want := 7
				if name == "loop" {
					args = nil
					want = 150
				}
				b.ReportAllocs()
				b.ResetTimer()
				for b.Loop() {
					if got, err := i.RunFunc(name, args...); err != nil || got != want {
						b.Fatalf("%v %v", got, err)
					}
				}
			})
		}
	}
}

func TestControlDescriptionWithoutBlock(t *testing.T) {
	jump := describeInstruction(&function{}, &ssa.Jump{}, nil, nil)
	if aux, ok := jump.Aux.(jumpAux); jump.Op != opJump || !ok || aux.Target != 0 {
		t.Fatalf("Jump without block: %#v", jump)
	}
	branch := describeInstruction(&function{}, &ssa.If{Cond: ssa.NewConst(constant.MakeBool(true), types.Typ[types.Bool])}, nil, nil)
	if aux, ok := branch.Aux.(ifAux); branch.Op != opIf || !ok || aux.Targets != [2]int32{} || !aux.Constant || !aux.Condition {
		t.Fatalf("If without block: %#v", branch)
	}
}

func TestValueDescriptionExecution(t *testing.T) {
	const source = `package main
 type number int
 type flag bool
 func calc(x int) int { return (x+3)*(10-x)-(-x) }
 func left(x int) int { return 4*x+7-x }
 func negate(x bool) bool { return !x }
 func named(x int) int { n:=number(x); return int(n+3) }
 func namedBool(x bool) bool { return bool(!flag(x)) }
 func divide(x int) int { return 12/x }
 func shift(x int) int { return x<<2 }
 func compare(x int) bool { return x<3 }
 func mixed(x int, n uint8) int { return x<<n }
 `
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		for _, extra := range []Mode{0, EnableCachedReg, ExperimentalSupportGC | EnableGCExecutionPlan} {
			t.Run(fmt.Sprintf("mode=%v/extra=%v", mode, extra), func(t *testing.T) {
				i, err := NewContext(DisableMemo|mode|extra).LoadInterp("value_execution.go", source)
				if err != nil {
					t.Fatal(err)
				}
				defer i.UnsafeRelease()
				built := 0
				fallback := 0
				for _, p := range i.funcs {
					for pc, d := range p.Code {
						if d.Op != opBinOp && d.Op != opUnOp {
							continue
						}
						switch aux := d.Aux.(type) {
						case binaryOpAux:
							if aux.Type == nil || aux.RightType == nil {
								t.Fatal("missing binary operand types")
							}
							for j, constant := range aux.Constant {
								if constant && !reflect.DeepEqual(aux.Values[j], p.stack[d.Args[j]]) {
									t.Fatal("constant value mismatch")
								}
							}
							if aux.Op == token.LSS && aux.Type != reflect.TypeFor[int]() {
								t.Fatal("comparison records result type instead of operand type")
							}
						case unaryOpAux:
							if aux.Type == nil {
								t.Fatal("missing unary operand type")
							}
						}
						supported := makeInstruction(nil, d)
						name := runtime.FuncForPC(reflect.ValueOf(p.Instrs[pc]).Pointer()).Name()
						if got, want := (strings.Contains(name, "makeInstruction") || strings.Contains(name, "makeBinaryInstruction") || strings.Contains(name, "makeUnaryInstruction")), supported != nil && mode&EnableInstructionExecution != 0; got != want {
							t.Fatalf("pc %d executor %s, want described=%t", pc, name, want)
						}
						if supported != nil {
							built++
						} else {
							fallback++
						}
					}
					p.functionCompilation = functionCompilation{}
				}
				if mode != 0 && built == 0 {
					t.Fatalf("supported=%d fallback=%d", built, fallback)
				}
				for _, c := range []struct {
					name      string
					arg, want value
				}{
					{"calc", 0, 30}, {"calc", 5, 45}, {"calc", -3, -3}, {"left", 2, 13}, {"negate", true, false}, {"negate", false, true},
					{"named", 4, 7}, {"namedBool", true, false}, {"divide", 3, 4}, {"shift", 3, 12}, {"compare", 2, true}, {"compare", 4, false},
				} {
					if got, err := i.RunFunc(c.name, c.arg); err != nil || got != c.want {
						t.Fatalf("%s(%v)=%v,%v want %v", c.name, c.arg, got, err, c.want)
					}
				}
			})
		}
	}
}

func TestValueDescriptionConstants(t *testing.T) {
	for _, op := range []token.Token{token.ADD, token.SUB, token.MUL} {
		for _, constants := range [][2]bool{{false, false}, {true, false}, {false, true}, {true, true}} {
			aux := binaryOpAux{Op: op, Type: reflect.TypeFor[int](), Constant: constants, Values: [2]value{7, 3}}
			d := typedInstruction{Op: opBinOp, Result: 2, Args: []register{0, 1}, Aux: aux}
			execute := makeInstruction(nil, d)
			if execute == nil {
				t.Fatal("missing executor")
			}
			x, y := 11, 5
			if constants[0] {
				x = 7
			}
			if constants[1] {
				y = 3
			}
			want := x + y
			if op == token.SUB {
				want = x - y
			}
			if op == token.MUL {
				want = x * y
			}
			fr := &frame{stack: []value{11, 5, nil}}
			execute(fr)
			if fr.reg(2) != want {
				t.Fatalf("%v constants=%v got %v want %v", op, constants, fr.reg(2), want)
			}
		}
	}
	for _, c := range []struct {
		op          token.Token
		typ         reflect.Type
		input, want value
	}{{token.SUB, reflect.TypeFor[int](), 7, -7}, {token.NOT, reflect.TypeFor[bool](), true, false}} {
		execute := makeInstruction(nil, typedInstruction{Op: opUnOp, Result: 0, Args: []register{-1}, Aux: unaryOpAux{Op: c.op, Type: c.typ, Constant: true, Value: c.input}})
		if execute == nil {
			t.Fatal("missing unary executor")
		}
		fr := &frame{stack: []value{nil}}
		execute(fr)
		if fr.reg(0) != c.want {
			t.Fatalf("got %v want %v", fr.reg(0), c.want)
		}
	}
}

func TestValueDescriptionIntOverflow(t *testing.T) {
	max := int(^uint(0) >> 1)
	min := -max - 1
	for _, c := range []struct {
		op         token.Token
		x, y, want int
	}{{token.ADD, max, 1, min}, {token.SUB, min, 1, max}, {token.MUL, max, 2, -2}} {
		execute := makeInstruction(nil, typedInstruction{Op: opBinOp, Result: 2, Args: []register{0, 1}, Aux: binaryOpAux{Op: c.op, Type: reflect.TypeFor[int]()}})
		fr := &frame{stack: []value{c.x, c.y, nil}}
		execute(fr)
		if fr.reg(2) != c.want {
			t.Fatalf("%v: %v want %v", c.op, fr.reg(2), c.want)
		}
	}
	execute := makeInstruction(nil, typedInstruction{Op: opUnOp, Result: 1, Args: []register{0}, Aux: unaryOpAux{Op: token.SUB, Type: reflect.TypeFor[int]()}})
	fr := &frame{stack: []value{min, nil}}
	execute(fr)
	if fr.reg(1) != min {
		t.Fatalf("negate min: %v", fr.reg(1))
	}
}

func BenchmarkValueDescription(b *testing.B) {
	const source = `package main
 func calc(x int) int { for j:=0;j<100;j++ { x=x*3+7-x }; return x }
 `
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution} {
		b.Run(fmt.Sprint(mode), func(b *testing.B) {
			ctx := NewContext(DisableMemo | mode)
			ctx.SetLeastCallForEnablePool(0)
			i, err := ctx.LoadInterp("value_bench.go", source)
			if err != nil {
				b.Fatal(err)
			}
			b.Cleanup(i.UnsafeRelease)
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				if _, err := i.RunFunc("calc", 1); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func TestBinaryDescriptionMatrix(t *testing.T) {
	groups := []struct {
		types []string
		ops   []string
	}{
		{[]string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"}, []string{"+", "-", "*", "/", "%", "&", "|", "^", "&^", "<", "<=", ">", ">=", "==", "!="}},
		{[]string{"float32", "float64"}, []string{"+", "-", "*", "/", "<", "<=", ">", ">=", "==", "!="}},
		{[]string{"complex64", "complex128"}, []string{"+", "-", "*", "/", "==", "!="}},
		{[]string{"string"}, []string{"+", "<", "<=", ">", ">=", "==", "!="}},
		{[]string{"bool"}, []string{"==", "!="}},
	}
	var source strings.Builder
	source.WriteString("package main\n")
	var names []string
	for _, g := range groups {
		for _, typ := range g.types {
			for _, named := range []bool{false, true} {
				tname := typ
				if named {
					tname = "named" + typ
					fmt.Fprintf(&source, "type %s %s\n", tname, typ)
				}
				fmt.Fprintf(&source, "func id%s(v %s) %s {return v}\n", tname, tname, tname)
				x, y := "7", "3"
				if typ == "string" {
					x, y = `"ab"`, `"cd"`
				}
				if typ == "bool" {
					x, y = "true", "false"
				}
				for _, op := range g.ops {
					for _, shape := range []string{"x OP y", "X OP y", "x OP Y", "X OP Y"} {
						name := fmt.Sprintf("binary%d", len(names))
						names = append(names, name)
						expr := strings.NewReplacer("OP", op, "X", tname+"("+x+")", "Y", tname+"("+y+")").Replace(shape)
						fmt.Fprintf(&source, "func %s() interface{} { x,y:=id%s(%s(%s)),id%s(%s(%s)); _,_=x,y;return %s }\n", name, tname, tname, x, tname, tname, y, expr)
					}
				}
			}
		}
	}
	// Exercise every left/right shift-kind pair, including named operand types.
	ints := groups[0].types
	for _, left := range ints {
		for _, right := range ints {
			for _, op := range []string{"<<", ">>"} {
				for _, shape := range []string{"x OP y", "X OP y", "x OP Y", "X OP Y"} {
					name := fmt.Sprintf("binary%d", len(names))
					names = append(names, name)
					expr := strings.NewReplacer("OP", op, "X", "named"+left+"(7)", "Y", "named"+right+"(2)").Replace(shape)
					fmt.Fprintf(&source, "func %s() interface{} {x,y:=idnamed%s(7),idnamed%s(2);_,_=x,y;return %s}\n", name, left, right, expr)
				}
			}
		}
	}
	load := func(mode Mode) *Interp {
		i, err := NewContext(DisableMemo|mode).LoadInterp("binary_matrix.go", source.String())
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(i.UnsafeRelease)
		return i
	}
	baseline := load(0)
	want := make([]string, len(names))
	for j, name := range names {
		v, err := baseline.RunFunc(name)
		if err != nil {
			t.Fatal(err)
		}
		want[j] = fmt.Sprint(v)
	}
	for _, mode := range []Mode{EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			i := load(mode)
			count := 0
			for _, p := range i.funcs {
				for pc, d := range p.Code {
					if d.Op == opBinOp {
						count++
						execute := makeInstruction(nil, d)
						if execute == nil {
							t.Fatalf("unsupported BinOp at %d", pc)
						}
						name := runtime.FuncForPC(reflect.ValueOf(p.Instrs[pc]).Pointer()).Name()
						if strings.Contains(name, "makeBinaryInstruction") != (mode&EnableInstructionExecution != 0) {
							t.Fatalf("wrong executor %s", name)
						}
					}
				}
				p.functionCompilation = functionCompilation{}
			}
			if count == 0 {
				t.Fatal("no binary instructions tested")
			}
			for j, name := range names {
				v, err := i.RunFunc(name)
				if err != nil || fmt.Sprint(v) != want[j] {
					t.Fatalf("%s=%v,%v want %s", name, v, err, want[j])
				}
			}
		})
	}
}

func TestBinaryDescriptionRuntimeEdges(t *testing.T) {
	const source = `package main
 type A [2]int
 type S struct{ X int }
 func edges() bool {
  a,b:=A{1,2},[2]int{1,2};s,t:=S{3},S{3};x:=1;p,q:=&x,&x
  ch:=make(chan int);var r <-chan int=ch
  var sl []int;var m map[int]int;var f func();var i interface{};var j interface{}
  return a==b && s==t && p==q && ch==r && sl==nil && m==nil && f==nil && i==j && !(a!=b) && !(ch!=r)
 }
 func zero() (ok bool) { defer func(){ok=recover()!=nil}(); x:=0; _=7/x;return }
 func negative(n int) (ok bool) { defer func(){ok=recover()!=nil}();x:=7;_=x<<n;return }
 func ifacePanic() (ok bool) {defer func(){ok=recover()!=nil}();var x interface{}=[]int{1};var y interface{}=[]int{1};_=x==y;return}
 `
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		i, err := NewContext(DisableMemo|mode).LoadInterp("binary_edges.go", source)
		if err != nil {
			t.Fatal(err)
		}
		for _, p := range i.funcs {
			p.functionCompilation = functionCompilation{}
		}
		for _, name := range []string{"edges", "zero", "negative", "ifacePanic"} {
			args := []value(nil)
			if name == "negative" {
				args = []value{-1}
			}
			if got, err := i.RunFunc(name, args...); err != nil || got != true {
				t.Fatalf("mode %v %s=%v,%v", mode, name, got, err)
			}
		}
		i.UnsafeRelease()
	}
}

func TestUnaryDescriptionExecution(t *testing.T) {
	const source = `package main
 import "runtime"
 type N int16
 type B bool
 var global=7
 var channel=make(chan int,1)
 func arithmetic() bool {a:=N(3);b:=B(true);return -a==N(-3) && ^a==N(-4) && !b==B(false)}
 func deref() int {p:=new(int);*p=5;return *p+global}
 func nilPointer() (ok bool) {defer func(){ok=recover()!=nil}();var p *int;return *p==0}
 func receive() bool {ch:=make(chan N,1);ch<-N(9);v,ok:=<-ch;close(ch);zero,open:=<-ch;runtime.GC();return v==9&&ok&&zero==0&&!open}
 func single() int {ch:=make(chan int,1);ch<-7;v:=<-ch;close(ch);return v+<-ch}
 func globalReceive() int {channel<-11;return <-channel}
 func skipped() int {var p *[3]int;n:=0;for j:=range *p {n+=j};return n}
 `
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		for _, extra := range []Mode{0, EnableCachedReg, ExperimentalSupportGC | EnableGCExecutionPlan} {
			t.Run(fmt.Sprintf("%v/%v", mode, extra), func(t *testing.T) {
				i, err := NewContext(DisableMemo|mode|extra).LoadInterp("unary_execution.go", source)
				if err != nil {
					t.Fatal(err)
				}
				defer i.UnsafeRelease()
				if _, err := i.RunFunc("init"); err != nil {
					t.Fatal(err)
				}
				for _, p := range i.funcs {
					for pc, d := range p.Code {
						if d.Op == opUnOp {
							execute := makeInstruction(nil, d)
							if execute == nil {
								t.Fatal("unsupported unary")
							}
							name := runtime.FuncForPC(reflect.ValueOf(p.Instrs[pc]).Pointer()).Name()
							if strings.Contains(name, "makeUnaryInstruction") != (mode&EnableInstructionExecution != 0) {
								t.Fatalf("wrong executor %s", name)
							}
						}
					}
					p.functionCompilation = functionCompilation{}
				}
				for _, c := range []struct {
					name string
					want value
				}{{"arithmetic", true}, {"deref", 12}, {"nilPointer", true}, {"receive", true}, {"single", 7}, {"globalReceive", 11}, {"skipped", 3}} {
					if got, err := i.RunFunc(c.name); err != nil || got != c.want {
						t.Fatalf("%s=%v,%v want %v", c.name, got, err, c.want)
					}
				}
			})
		}
	}
}

func TestUnaryDescriptionTypeMatrix(t *testing.T) {
	var source strings.Builder
	source.WriteString("package main\n")
	var names []string
	for _, typ := range []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "float32", "float64", "complex64", "complex128", "bool"} {
		for _, named := range []bool{false, true} {
			tname := typ
			if named {
				tname = "N" + typ
				fmt.Fprintf(&source, "type %s %s\n", tname, typ)
			}
			fmt.Fprintf(&source, "func id%s(x %s) %s {return x}\n", tname, tname, tname)
			ops := []string{"-"}
			v := "3"
			switch {
			case typ == "bool":
				ops = []string{"!"}
				v = "true"
			case strings.HasPrefix(typ, "int") || strings.HasPrefix(typ, "uint"):
				ops = append(ops, "^")
			}
			for _, op := range ops {
				for _, constant := range []bool{false, true} {
					name := fmt.Sprintf("unary%d", len(names))
					names = append(names, name)
					expr := "x"
					if constant {
						expr = tname + "(" + v + ")"
					}
					// Unsigned constant negation is rejected by the language; dynamic negation wraps.
					if constant && strings.HasPrefix(typ, "uint") && op == "-" {
						names = names[:len(names)-1]
						continue
					}
					fmt.Fprintf(&source, "func %s() interface{} {x:=id%s(%s);_=x;return %s%s}\n", name, tname, v, op, expr)
				}
			}
		}
	}
	var want []string
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		i, err := NewContext(DisableMemo|mode).LoadInterp("unary_matrix.go", source.String())
		if err != nil {
			t.Fatal(err)
		}
		for _, p := range i.funcs {
			p.functionCompilation = functionCompilation{}
		}
		for j, name := range names {
			v, err := i.RunFunc(name)
			if err != nil {
				t.Fatal(err)
			}
			got := fmt.Sprint(v)
			if mode == 0 {
				want = append(want, got)
			} else if got != want[j] {
				t.Fatalf("%s=%v want %s", name, v, want[j])
			}
		}
		i.UnsafeRelease()
	}
}

// Exercise builder-only constant specializations that SSA normally folds away.
func TestBinaryDescriptionConstantSpecializations(t *testing.T) {
	groups := []struct{ types, ops []string }{
		{[]string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"}, []string{"+", "-", "*", "/", "%", "&", "|", "^", "&^", "<<", ">>", "<", "<=", ">", ">=", "==", "!="}},
		{[]string{"float32", "float64"}, []string{"+", "-", "*", "/", "<", "<=", ">", ">=", "==", "!="}},
		{[]string{"complex64", "complex128"}, []string{"+", "-", "*", "/", "==", "!="}},
		{[]string{"string"}, []string{"+", "<", "<=", ">", ">=", "==", "!="}},
		{[]string{"bool"}, []string{"==", "!="}},
	}
	var source strings.Builder
	source.WriteString("package main\n")
	for _, g := range groups {
		for _, typ := range g.types {
			for _, named := range []bool{false, true} {
				tname := typ
				if named {
					tname = "C" + typ
					fmt.Fprintf(&source, "type %s %s\n", tname, typ)
				}
				for j, op := range g.ops {
					fmt.Fprintf(&source, "func f%s%d(x,y %s) interface{} {return x %s y}\n", tname, j, tname, op)
				}
			}
		}
	}
	integers := groups[0].types
	for _, left := range integers {
		for _, right := range integers {
			for _, op := range []string{"<<", ">>"} {
				fmt.Fprintf(&source, "func shift%s%s%d(x C%s,y C%s) interface{} {return x %s y}\n", left, right, len(source.String()), left, right, op)
				fmt.Fprintf(&source, "func basic%s%s%d(x %s,y %s) interface{} {return x %s y}\n", left, right, len(source.String()), left, right, op)
			}
		}
	}

	i, err := NewContext(DisableMemo|EnableInstructionDescriptions).LoadInterp("binary_constants.go", source.String())
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	operand := func(typ reflect.Type, n int) value {
		v := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v.SetInt(int64(n))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			v.SetUint(uint64(n))
		case reflect.Float32, reflect.Float64:
			v.SetFloat(float64(n))
		case reflect.Complex64, reflect.Complex128:
			v.SetComplex(complex(float64(n), 1))
		case reflect.String:
			v.SetString(fmt.Sprint(n))
		case reflect.Bool:
			v.SetBool(n == 7)
		default:
			t.Fatalf("unsupported %v", typ)
		}
		return v.Interface()
	}
	for _, p := range i.funcs {
		for pc, d := range p.Code {
			if d.Op != opBinOp {
				continue
			}
			aux := d.Aux.(binaryOpAux)
			x, y := operand(aux.Type, 7), operand(aux.RightType, 3)
			stack := append([]value(nil), p.stack...)
			stack[d.Args[0]], stack[d.Args[1]] = x, y
			fr := &frame{pfn: p, stack: stack}
			p.Instrs[pc](fr)
			want := fr.reg(d.Result)
			for _, c := range [][2]bool{{false, false}, {true, false}, {false, true}, {true, true}} {
				aux.Constant = c
				aux.Values = [2]value{x, y}
				d.Aux = aux
				execute := makeInstruction(nil, d)
				if execute == nil {
					t.Fatal("missing binary executor")
				}
				fr.stack = append([]value(nil), stack...)
				execute(fr)
				if !reflect.DeepEqual(fr.reg(d.Result), want) {
					t.Fatalf("%v %v constants=%v: %v want %v", aux.Type, aux.Op, c, fr.reg(d.Result), want)
				}
			}
		}
	}
}

func TestUnaryDescriptionConstantSpecializations(t *testing.T) {
	var source strings.Builder
	source.WriteString("package main\n")
	for _, typ := range []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "float32", "float64", "complex64", "complex128", "bool"} {
		for _, named := range []bool{false, true} {
			name := typ
			if named {
				name = "U" + typ
				fmt.Fprintf(&source, "type %s %s\n", name, typ)
			}
			ops := []string{"-"}
			if typ == "bool" {
				ops = []string{"!"}
			} else if strings.HasPrefix(typ, "int") || strings.HasPrefix(typ, "uint") {
				ops = append(ops, "^")
			}
			for j, op := range ops {
				fmt.Fprintf(&source, "func unary%s%d(x %s) interface{} {return %sx}\n", name, j, name, op)
			}
		}
	}
	i, err := NewContext(DisableMemo|EnableInstructionDescriptions).LoadInterp("unary_constants.go", source.String())
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	for _, p := range i.funcs {
		for pc, d := range p.Code {
			if d.Op != opUnOp || (d.Aux.(unaryOpAux).Op != token.SUB && d.Aux.(unaryOpAux).Op != token.XOR && d.Aux.(unaryOpAux).Op != token.NOT) {
				continue
			}
			aux := d.Aux.(unaryOpAux)
			v := reflect.New(aux.Type).Elem()
			switch aux.Type.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v.SetInt(3)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				v.SetUint(3)
			case reflect.Float32, reflect.Float64:
				v.SetFloat(3.5)
			case reflect.Complex64, reflect.Complex128:
				v.SetComplex(3 + 2i)
			case reflect.Bool:
				v.SetBool(true)
			default:
				t.Fatal(aux.Type)
			}
			stack := append([]value(nil), p.stack...)
			stack[d.Args[0]] = v.Interface()
			fr := &frame{pfn: p, stack: stack}
			p.Instrs[pc](fr)
			want := fr.reg(d.Result)
			for _, c := range []bool{false, true} {
				aux.Constant = c
				aux.Value = v.Interface()
				d.Aux = aux
				execute := makeInstruction(nil, d)
				if execute == nil {
					t.Fatal("missing unary executor")
				}
				fr.stack = append([]value(nil), stack...)
				execute(fr)
				if !reflect.DeepEqual(fr.reg(d.Result), want) {
					t.Fatalf("%v %v constant=%t got %v want %v", aux.Type, aux.Op, c, fr.reg(d.Result), want)
				}
			}
		}
	}
}

func TestUnaryDescriptionGlobalReceive(t *testing.T) {
	for _, global := range []bool{false, true} {
		for _, comma := range []bool{false, true} {
			for _, closed := range []bool{false, true} {
				ch := make(chan int, 1)
				if closed {
					close(ch)
				} else {
					ch <- 9
				}
				d := typedInstruction{Op: opUnOp, Result: 1, Args: []register{0}, Aux: unaryOpAux{Op: token.ARROW, Type: reflect.TypeFor[chan int](), Global: global, CommaOk: comma, Value: ch}}
				execute := makeInstruction(nil, d)
				if execute == nil {
					t.Fatal("missing receive")
				}
				p := &function{}
				p.Interp = &Interp{}
				fr := &frame{pfn: p, stack: []value{ch, nil}}
				execute(fr)
				var want value = 9
				if closed {
					want = 0
				}
				if comma {
					want = tuple{want, !closed}
				}
				if !reflect.DeepEqual(fr.reg(1), want) {
					t.Fatalf("global=%t comma=%t closed=%t: %v want %v", global, comma, closed, fr.reg(1), want)
				}
			}
		}
	}
}

func TestEmptyPhiDescription(t *testing.T) {
	for _, aux := range []phiAux{{}, {Results: []register{0}}, {Edges: [][]register{{0}}}} {
		execute := makeInstruction(nil, typedInstruction{Op: opPhi, Aux: aux})
		if execute == nil {
			t.Fatal("empty Phi returned nil executor")
		}
		fr := &frame{stack: []value{1}}
		execute(fr)
		if fr.phiVals != nil || fr.reg(0) != 1 {
			t.Fatalf("empty Phi changed frame: %#v", fr)
		}
	}
}

func TestPhiDirectAssignment(t *testing.T) {
	for _, c := range phiDirectAssignmentCases {
		if len(c.results) == 0 || len(c.edges) == 0 {
			continue
		}
		t.Run(c.name, func(t *testing.T) {
			preds := make([]int32, len(c.edges[0]))
			for j := range preds {
				preds[j] = int32(j + 1)
			}
			execute := makeInstruction(nil, typedInstruction{Op: opPhi, Aux: phiAux{Results: c.results, Edges: c.edges, Preds: preds}})
			unknown := &frame{pred: -1, stack: []value{10, 20, 30, 40}}
			execute(unknown)
			if !reflect.DeepEqual(unknown.stack, []value{10, 20, 30, 40}) || unknown.phiVals != nil {
				t.Fatal("unknown predecessor changed frame")
			}
			for j, pred := range preds {
				initial := []value{10, 20, 30, 40}
				fr := &frame{pred: pred, stack: append([]value(nil), initial...)}
				execute(fr)
				for k, result := range c.results {
					if fr.reg(result) != initial[c.edges[k][j]] {
						t.Fatalf("pred %d result %d=%v want %v", pred, result, fr.reg(result), initial[c.edges[k][j]])
					}
				}
				if c.safe && fr.phiVals != nil {
					t.Fatal("direct assignment allocated scratch")
				}
			}
		})
	}
}

func BenchmarkPhiAssignment(b *testing.B) {
	const source = `package main
 func single(n int) int {x:=1;if n>0{x=2};return x}
 func independent(n int) int {a,c:=1,2;for j:=0;j<n;j++{a=a+1;c=c+2};return a+c}
 func swap(n int) int {a,c:=1,2;for j:=0;j<n;j++{a,c=c,a};return a*10+c}
 func rotate(n int) int {a,c,d:=1,2,3;for j:=0;j<n;j++{a,c,d=c,d,a};return a*100+c*10+d}
 `
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution} {
		for _, c := range []struct {
			name    string
			n, want int
		}{{"single", 100, 2}, {"independent", 100, 303}, {"swap", 100, 12}, {"rotate", 99, 123}} {
			b.Run(fmt.Sprintf("mode=%v/%s", mode, c.name), func(b *testing.B) {
				ctx := NewContext(DisableMemo | mode)
				ctx.SetLeastCallForEnablePool(0)
				i, err := ctx.LoadInterp("phi_assignment.go", source)
				if err != nil {
					b.Fatal(err)
				}
				b.Cleanup(i.UnsafeRelease)
				if got, err := i.RunFunc(c.name, c.n); err != nil || got != c.want {
					b.Fatal(got, err)
				}
				b.ReportAllocs()
				b.ResetTimer()
				for b.Loop() {
					if got, err := i.RunFunc(c.name, c.n); err != nil || got != c.want {
						b.Fatal(got, err)
					}
				}
			})
		}
	}
}

func TestTypedInstructionDescription(t *testing.T) {
	ctx := NewContext(DisableMemo | EnableInstructionDescriptions)
	interp, err := ctx.LoadInterp("typed_instruction.go", `package main
func branch(n int) int {
	if n > 0 { return n + 1 }
	return n - 1
}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	p := interp.funcs[interp.mainpkg.Func("branch")]
	if p == nil {
		t.Fatal("missing branch function")
	}
	if len(p.Code) != len(p.Instrs) || len(p.Code) != len(p.ssaInstrs) {
		t.Fatalf("typed code lengths = %d/%d/%d", len(p.Code), len(p.Instrs), len(p.ssaInstrs))
	}
	if len(p.Code) == 0 {
		t.Fatal("empty typed code")
	}
	var sawBranch bool
	for pc, desc := range p.Code {
		if desc.Op == opIf || desc.Op == opJump {
			sawBranch = true
			switch desc.Op {
			case opIf:
				if _, ok := desc.Aux.(ifAux); !ok {
					t.Fatalf("missing If targets at pc %d", pc)
				}
			case opJump:
				if _, ok := desc.Aux.(jumpAux); !ok {
					t.Fatalf("missing Jump target at pc %d", pc)
				}
			}
		}
		if desc.Op == opBinOp {
			aux, ok := desc.Aux.(binaryOpAux)
			if !ok || aux.Op == token.ILLEGAL {
				t.Fatalf("missing binary operator at pc %d", pc)
			}
		}
		if desc.Op == opInvalid {
			t.Fatalf("invalid opcode at pc %d", pc)
		}
		_ = p.PosForPC(pc) // Some synthetic instructions have no source position.
	}
	if !sawBranch {
		t.Fatal("missing branch instruction")
	}
	if len(p.Blocks) == 0 || p.Blocks[0] != 0 {
		t.Fatalf("block entries = %v", p.Blocks)
	}
	if got := interp.callFunction(&frame{}, p, []value{2}, nil); got != 3 {
		t.Fatalf("branch result = %v, want 3", got)
	}
}

func TestInstructionFunctionReferences(t *testing.T) {
	const source = `package main
 import "strings"
 func first(x int) int {return x+1}
 func second(x int) int {return x+2}
 func recursive(n int) int {if n==0{return 0};return 1+recursive(n-1)}
 func closures(x int) (func(int) int,func(int) int) {return func(y int) int{return x+y},func(y int) int{return x-y}}
 func run() int {a,b:=closures(7);if strings.ToUpper("ab")!="AB"{return -1};return first(1)+second(1)+recursive(3)+a(2)+b(2)}
 `
	for _, mode := range []Mode{EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			i, err := NewContext(DisableMemo|mode).LoadInterp("function_refs.go", source)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			bodies := make(map[string]*function)
			var closures []closureAux
			var external reflect.Value
			for _, p := range i.funcs {
				for pc, d := range p.Code {
					switch aux := d.Aux.(type) {
					case closureAux:
						if aux.Body == nil {
							t.Fatal("missing closure body")
						}
						if aux.Body != i.funcs[p.ssaInstrs[pc].(*ssa.MakeClosure).Fn.(*ssa.Function)] {
							t.Fatal("wrong closure body")
						}
						closures = append(closures, aux)
					case bodyCallAux:
						call := p.ssaInstrs[pc].(*ssa.Call)
						fn := call.Call.Value.(*ssa.Function)
						if aux.Body == nil || aux.Body != i.funcs[fn] {
							t.Fatalf("wrong target %s", fn)
						}
						bodies[fn.Name()] = aux.Body
					case externalCallAux:
						call := p.ssaInstrs[pc].(*ssa.Call)
						if call.Call.StaticCallee().Name() == "ToUpper" {
							external = aux.External
						}
					}
				}
			}
			if bodies["first"] == nil || bodies["second"] == nil || bodies["first"] == bodies["second"] || bodies["recursive"] == nil {
				t.Fatal("function identity lost")
			}
			if len(closures) != 2 || closures[0].Signature != closures[1].Signature || closures[0].Body == closures[1].Body {
				t.Fatal("closure identity lost")
			}
			if !external.IsValid() {
				t.Fatal("missing external target")
			}
			for _, p := range i.funcs {
				p.functionCompilation = functionCompilation{}
			}
			for name, want := range map[string]int{"first": 5, "second": 6, "recursive": 4} {
				got := bodies[name].callFunctionByReflect(reflect.TypeFor[func(int) int](), []reflect.Value{reflect.ValueOf(4)}, nil)
				if len(got) != 1 || got[0].Interface() != want {
					t.Fatalf("%s returned %v", name, got)
				}
			}
			if got := external.Call([]reflect.Value{reflect.ValueOf("ab")})[0].String(); got != "AB" {
				t.Fatal(got)
			}
			for j, aux := range closures {
				capture := 7
				fn := aux.Body.makeFunction(aux.Signature, []value{&capture}).Interface().(func(int) int)
				want := 9
				if j == 1 {
					want = 5
				}
				if got := fn(2); got != want {
					t.Fatalf("closure %d=%d want %d", j, got, want)
				}
			}
			if got, err := i.RunFunc("run"); err != nil || got != 22 {
				t.Fatalf("run=%v,%v", got, err)
			}
		})
	}
}

func TestDescriptionStaticCallExecution(t *testing.T) {
	const source = `package main
 type counter int
 func (c counter) add(x int) int {return int(c)+x}
 func none(p *int) {*p+=1}
 func one(x int) int {return x+1}
 func pair(x int) (int,int) {return x,x+1}
 func triple(x int) (int,int,int) {return x,x+1,x+2}
 func recovered() (v int) {defer func(){if recover()!=nil{v=7}}();panic("test")}
 func deferred(x int) (v int) {defer func(){v++}();return x}
 func factory(x int) func(int) int {return func(y int) int {x+=y;return x}}
 func run() int {n:=0;none(&n);a,b:=pair(2);_,c,_:=triple(3);f:=factory(5);return n+one(1)+a+b+c+recovered()+deferred(3)+counter(2).add(3)+f(2)+f(1)}
 `
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		for _, extra := range []Mode{0, EnableCachedReg, ExperimentalSupportGC | EnableGCExecutionPlan} {
			t.Run(fmt.Sprintf("%v/%v", mode, extra), func(t *testing.T) {
				i, err := NewContext(DisableMemo|mode|extra).LoadInterp("static_calls.go", source)
				if err != nil {
					t.Fatal(err)
				}
				defer i.UnsafeRelease()
				calls, closures := 0, 0
				for _, p := range i.funcs {
					for pc, d := range p.Code {
						if d.Op != opCall && d.Op != opMakeClosure {
							continue
						}
						supported := false
						switch aux := d.Aux.(type) {
						case dynamicCallAux:
							supported = aux.Target >= 0
						case bodyCallAux:
							supported = aux.Body != nil
							if supported {
								calls++
							}
						case closureAux:
							supported = aux.Body != nil
							if supported {
								closures++
							}
						}
						if !supported {
							continue
						}
						// Rebuild to verify that descriptions are sufficient for construction.
						if makeInstruction(i, d) == nil {
							t.Fatal("missing description executor")
						}
						name := runtime.FuncForPC(reflect.ValueOf(p.Instrs[pc]).Pointer()).Name()
						// Call specializers and dynamic-call helpers are shared by both paths.
						_, specialized := d.Aux.(dynamicCallAux)
						if aux, ok := d.Aux.(bodyCallAux); ok && aux.Body.Recover == nil {
							specialized = makeCallSpecializerInstr(i, aux.Body, aux.Arguments, aux.ExtractResults) != nil
						}
						if strings.Contains(name, "makeInstruction") != (mode&EnableInstructionExecution != 0) && !specialized {
							t.Fatalf("unexpected executor %s", name)
						}
					}
					p.functionCompilation = functionCompilation{}
				}
				if mode != 0 && (calls == 0 || closures == 0) {
					t.Fatal("missing call/closure coverage")
				}
				if got, err := i.RunFunc("run"); err != nil || got != 43 {
					t.Fatalf("run=%v,%v want 43", got, err)
				}
			})
		}
	}
}

func TestDescriptionPreservesDirectCall(t *testing.T) {
	const pkgPath = "example.com/instructiondirect"
	calls := 0
	RegisterDirectCalls(pkgPath, map[string]DirectCallAdapter{"Value": func(c DirectCallContext) { calls++; c.SetResult(42) }})
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		ctx := NewContext(DisableMemo | mode)
		file, err := parser.ParseFile(ctx.FileSet, "direct.go", `package instructiondirect
 func Value() int {return 7}
 func Plain() int {return 1}
 func Run() int {return Value()+Plain()}`, 0)
		if err != nil {
			t.Fatal(err)
		}
		pkg, err := ctx.LoadAstFile(pkgPath, file)
		if err != nil {
			t.Fatal(err)
		}
		i, err := NewInterp(ctx, pkg)
		if err != nil {
			t.Fatal(err)
		}
		before := calls
		if got, err := i.RunFunc("Run"); err != nil || got != 43 || calls != before+1 {
			t.Fatalf("mode %v=%v,%v calls=%d", mode, got, err, calls-before)
		}
		for _, p := range i.funcs {
			if p.directCall != nil || p.directAdapter != nil {
				t.Fatal("lowering call state retained")
			}
			for pc, d := range p.Code {
				if d.Op == opCall {
					aux := instructionCallMetadata(d.Aux)
					wantDirect := p.ssaInstrs[pc].(*ssa.Call).Call.StaticCallee().Name() == "Value"
					if _, direct := d.Aux.(directCallAux); direct != wantDirect {
						t.Fatal("wrong call aux type")
					}
					if aux.Kind == callStatic && makeInstruction(i, d) == nil {
						t.Fatal("incorrect direct adapter selection")
					}
					if wantDirect {
						fr := &frame{pfn: &function{Interp: i}, stack: make([]value, len(p.stack))}
						makeInstruction(i, d)(fr)
						if got := fr.reg(d.Result); got != 42 {
							t.Fatalf("direct description result=%v", got)
						}
					}
				}
			}
		}
		// Description reads the lowering decision even when an adapter is registered.
		p := i.funcs[pkg.Func("Run")]
		for _, block := range p.Fn.Blocks {
			for _, instr := range block.Instrs {
				call, ok := instr.(*ssa.Call)
				if !ok || call.Call.StaticCallee().Name() != "Value" {
					continue
				}
				if _, direct := describeCall(p, &call.Call, nil, nil).(directCallAux); direct {
					t.Fatal("description resolved adapter without lowering decision")
				}
				p.directCall = &call.Call
				p.directAdapter = func(DirectCallContext) {}
				if _, direct := describeCall(p, &call.Call, nil, nil).(directCallAux); !direct {
					t.Fatal("description lost lowering decision")
				}
				aux := describeCall(p, &call.Call, nil, call).(directCallAux)
				if aux.DeferStack != descriptionRegister(p, call) || aux.DeferStack < 0 {
					t.Fatal("direct adapter lost defer-stack register")
				}
				p.directCall = nil
				p.directAdapter = nil
			}
		}
		i.UnsafeRelease()
	}
}

func TestDescriptionDirectCallResults(t *testing.T) {
	const path = "example.com/instructiondirectresults"
	RegisterDirectCalls(path, map[string]DirectCallAdapter{
		"Store": func(c DirectCallContext) { *DirectCallArg[*int](c, 0) = 5 },
		"Pair":  func(c DirectCallContext) { c.SetResult(tuple{20, 22}) },
		"Sum": func(c DirectCallContext) {
			x := DirectCallArg[int](c, 0)
			for _, n := range DirectCallArg[[]int](c, 1) {
				x += n
			}
			c.SetResult(x)
		},
	})
	const source = `package directresults
 func Store(p *int) { *p = -1 }
 func Pair() (int,int) { return -1,-1 }
 func Sum(x int, rest ...int) int { return -1 }
 func Run() int { x:=0; Store(&x); a,b:=Pair(); return x+a+b+Sum(20,10,12) }`
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		for _, extra := range []Mode{0, EnableCachedReg, ExperimentalSupportGC | EnableGCExecutionPlan} {
			t.Run(fmt.Sprintf("mode=%v/extra=%v", mode, extra), func(t *testing.T) {
				ctx := NewContext(DisableMemo | mode | extra)
				file, err := parser.ParseFile(ctx.FileSet, "direct_results.go", source, 0)
				if err != nil {
					t.Fatal(err)
				}
				pkg, err := ctx.LoadAstFile(path, file)
				if err != nil {
					t.Fatal(err)
				}
				i, err := NewInterp(ctx, pkg)
				if err != nil {
					t.Fatal(err)
				}
				defer i.UnsafeRelease()
				for _, p := range i.funcs {
					p.functionCompilation = functionCompilation{}
				}
				if got, err := i.RunFunc("Run"); err != nil || got != 89 {
					t.Fatalf("result=%v,%v want 89", got, err)
				}
			})
		}
	}
}

func TestDescriptionDirectCallExecution(t *testing.T) {
	for _, test := range []struct {
		name    string
		args    []value
		adapter DirectCallAdapter
		want    value
	}{
		{"zero", []value{new(int)}, func(c DirectCallContext) { *DirectCallArg[*int](c, 0) = 42 }, nil},
		{"one", []value{20, 22}, func(c DirectCallContext) { c.SetResult(DirectCallArg[int](c, 0) + DirectCallArg[int](c, 1)) }, 42},
		{"tuple", []value{20}, func(c DirectCallContext) { x := DirectCallArg[int](c, 0); c.SetResult(tuple{x, x + 2}) }, tuple{20, 22}},
		{"receiver", []value{new(int), 42}, func(c DirectCallContext) {
			p := DirectCallArg[*int](c, 0)
			*p = DirectCallArg[int](c, 1)
			c.SetResult(*p)
		}, 42},
		{"variadic", []value{20, []int{10, 12}}, func(c DirectCallContext) {
			x := DirectCallArg[int](c, 0)
			for _, n := range DirectCallArg[[]int](c, 1) {
				x += n
			}
			c.SetResult(x)
		}, 42},
	} {
		t.Run(test.name, func(t *testing.T) {
			args := make([]register, len(test.args))
			for n := range args {
				args[n] = register(n)
			}
			result := register(len(args))
			d := typedInstruction{Op: opCall, Result: result, Aux: directCallAux{callAux: callAux{callSiteAux: callSiteAux{Kind: callStatic, Arguments: args}}, Adapter: test.adapter}}
			interp := &Interp{}
			run := makeInstruction(interp, d)
			if run == nil {
				t.Fatal("missing direct-call executor")
			}
			// No interpreted body, external function, or compilation state is needed.
			p := &function{}
			fr := &frame{pfn: p, stack: append(append([]value(nil), test.args...), nil), deferid: 7}
			run(fr)
			if got := fr.reg(result); !reflect.DeepEqual(got, test.want) {
				t.Fatalf("result=%v want %v", got, test.want)
			}
			if test.name == "zero" && *test.args[0].(*int) != 42 {
				t.Fatal("void adapter was not called")
			}
			if got, ok := interp.deferMap.Load(uint64(7)); !ok || got != fr {
				t.Fatal("defer frame not tracked")
			}
		})
	}
	t.Run("panic", func(t *testing.T) {
		d := typedInstruction{Op: opCall, Aux: directCallAux{callAux: callAux{callSiteAux: callSiteAux{Kind: callStatic}}, Adapter: func(DirectCallContext) { panic("direct panic") }}}
		defer func() {
			if got := recover(); got != "direct panic" {
				t.Fatalf("panic=%v", got)
			}
		}()
		makeInstruction(&Interp{}, d)(&frame{pfn: &function{}})
	})
}

func TestDescriptionExternalCallExecution(t *testing.T) {
	for _, withFrame := range []bool{false, true} {
		t.Run(fmt.Sprintf("frame=%v", withFrame), func(t *testing.T) {
			external := reflect.ValueOf(func(x int) (int, int) { return x, x + 2 })
			if withFrame {
				external = reflect.ValueOf(func(fr *frame, x int) (int, int) {
					if fr.deferid != 7 {
						t.Fatal("missing caller frame")
					}
					return x, x + 2
				})
			}
			for _, extracted := range []bool{false, true} {
				aux := externalCallAux{callAux: callAux{callSiteAux: callSiteAux{Kind: callStatic, Arguments: []register{0}}}, External: external}
				result := register(1)
				if extracted {
					aux.ExtractResults = []register{1, 2}
					result = -1
				}
				interp := &Interp{}
				run := makeInstruction(interp, typedInstruction{Op: opCall, Result: result, Aux: aux})
				if run == nil {
					t.Fatal("missing external executor")
				}
				fr := &frame{pfn: &function{}, stack: []value{20, nil, nil}, deferid: 7}
				run(fr)
				if extracted {
					if fr.reg(1) != 20 || fr.reg(2) != 22 {
						t.Fatal("incorrect extracted results")
					}
				} else if !reflect.DeepEqual(fr.reg(1), tuple{20, 22}) {
					t.Fatalf("tuple=%v", fr.reg(1))
				}
				if got, ok := interp.deferMap.Load(uint64(7)); !ok || got != fr {
					t.Fatal("missing defer frame")
				}
			}
		})
	}
	for _, aux := range []instructionAux{directCallAux{}, externalCallAux{}, bodyCallAux{}} {
		if makeInstruction(nil, typedInstruction{Op: opCall, Aux: aux}) != nil {
			t.Fatal("unresolved target executor")
		}
	}
}

func TestDescriptionCallFallbacks(t *testing.T) {
	if makeInstruction(nil, typedInstruction{Op: opMakeClosure, Aux: closureAux{}}) != nil {
		t.Fatal("unresolved closure executor")
	}
	for _, kind := range []callKind{callStatic, callDynamic, callBuiltin, callInvoke} {
		if makeInstruction(nil, typedInstruction{Op: opCall, Aux: callAux{callSiteAux: callSiteAux{Kind: kind}}}) != nil {
			t.Fatal("unresolved call executor")
		}
	}
	ctx := NewContext(DisableMemo | EnableInstructionExecution)
	i, err := ctx.LoadInterp("call_recover_shapes.go", `package main
 func zero(p *int){defer func(){*p+=1}()}
 func many(x int)(int,int){defer func(){}();return x,x+1}
 func run()int{x:=0;zero(&x);a,b:=many(2);return x+a+b}`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	if got, err := i.RunFunc("run"); err != nil || got != 6 {
		t.Fatalf("run=%v,%v", got, err)
	}
	fn := &ssa.Function{Signature: types.NewSignatureType(nil, nil, nil, types.NewTuple(), types.NewTuple(), false)}
	if _, ok := describeCall(i.funcs[i.mainpkg.Func("run")], &ssa.CallCommon{Value: fn}, nil, nil).(externalCallAux); ok {
		t.Fatal("unallocated external resolved")
	}
}

func TestDescriptionGoDeferExecution(t *testing.T) {
	const source = `package godefer
 import "sync"
 import "strings"
 type counter struct { n int }
 func (c *counter) Add(n int) { c.n += n }
 func send(ch chan int, n int) { ch <- n }
 func recovered() (r int) { defer func(){if recover()!=nil {r=7}}(); panic("test") }
 func deferred() (r int) {
 c:=&counter{}; n:=2
 defer func(){r=c.n}()
 defer c.Add(n)
 n=9
 defer c.Add(3)
 defer strings.ToUpper("unused")
 return
 }
 func Run() int {
 ch:=make(chan int,1); n:=4
 go send(ch,n); n=20
 x:=<-ch
 c:=&counter{}; defer c.Add(10)
 var wg sync.WaitGroup; wg.Add(1); go wg.Done(); wg.Wait()
 captured:=6; go func(){ch<-captured}(); y:=<-ch
 return x+y+deferred()+recovered()
 }`
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		for _, extra := range []Mode{0, EnableCachedReg, ExperimentalSupportGC | EnableGCExecutionPlan} {
			t.Run(fmt.Sprintf("mode=%v/extra=%v", mode, extra), func(t *testing.T) {
				ctx := NewContext(DisableMemo | mode | extra)
				i, err := ctx.LoadInterp("go_defer.go", source)
				if err != nil {
					t.Fatal(err)
				}
				defer i.UnsafeRelease()
				supported := 0
				for _, p := range i.funcs {
					for _, d := range p.Code {
						if d.Op != opGo && d.Op != opDefer {
							continue
						}
						switch d.Aux.(type) {
						case bodyCallAux, externalCallAux:
							if makeInstruction(i, d) == nil {
								t.Fatal("missing static go/defer executor")
							}
							supported++
						}
					}
					p.functionCompilation = functionCompilation{}
				}
				if mode != 0 && supported == 0 {
					t.Fatal("no static go/defer descriptions")
				}
				for n := 0; n < 3; n++ {
					if got, err := i.RunFunc("Run"); err != nil || got != 22 {
						t.Fatalf("Run=%v,%v want 22", got, err)
					}
				}
			})
		}
	}
}

func TestDescriptionDeferStackExecution(t *testing.T) {
	interp := &Interp{}
	for _, external := range []bool{false, true} {
		t.Run(fmt.Sprintf("external=%v", external), func(t *testing.T) {
			base := callAux{callSiteAux: callSiteAux{Kind: callStatic, DeferStack: 0, Arguments: []register{1}}}
			body := &function{}
			var aux instructionAux = bodyCallAux{callAux: base, Body: body}
			ext := reflect.ValueOf(func(int) {})
			if external {
				aux = externalCallAux{callAux: base, External: ext}
			}
			run := makeInstruction(interp, typedInstruction{Op: opDefer, Aux: aux})
			var head *_defer
			fr := &frame{stack: []value{&head, 42}}
			run(fr)
			fr.stack[1] = 99
			if fr._defer != nil || head == nil || len(head.args) != 1 || head.args[0] != 42 {
				t.Fatal("incorrect defer stack or argument snapshot")
			}
			if external {
				if got, ok := head.fn.(reflect.Value); !ok || got.Pointer() != ext.Pointer() {
					t.Fatal("wrong external target")
				}
			} else if head.fn != body {
				t.Fatal("wrong body target")
			}
		})
	}
	for _, op := range []instructionOpcode{opGo, opDefer} {
		for _, aux := range []instructionAux{callAux{}, bodyCallAux{}, externalCallAux{}} {
			if makeInstruction(interp, typedInstruction{Op: op, Aux: aux}) != nil {
				t.Fatal("unresolved go/defer should retain lowering")
			}
		}
	}
}

func TestDescriptionRangeDefer(t *testing.T) {
	const source = `package rangedefer
 import "strings"
 type adder struct{ p *int }
 func (a adder) Add(n int){*a.p=(*a.p*10+n)%100000000}
 type adding interface{Add(int)}
 func seq(yield func(int)bool){for n:=1;n<=2;n++{if !yield(n){return}}}
 func add(p *int,n int){*p=(*p*10+n)%100000000}
 func Run()(r int){
 var a adding=adder{&r}; f:=add
 for n:=range seq {
 defer add(&r,n)
 defer adder{&r}.Add(n+2)
 defer a.Add(n+4)
 defer f(&r,n+6)
 defer func(x int){r=(r*10+x)%100000000}(n+8)
 defer strings.ToUpper("ab")
 defer close(make(chan int))
 }
 if r!=0 {panic("defer ran in yield")}
 return
 }`
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		for _, extra := range []Mode{0, EnableCachedReg, ExperimentalSupportGC | EnableGCExecutionPlan} {
			t.Run(fmt.Sprintf("mode=%v/extra=%v", mode, extra), func(t *testing.T) {
				i, err := NewContext(DisableMemo|mode|extra).LoadInterp("range_defer.go", source)
				if err != nil {
					t.Fatal(err)
				}
				defer i.UnsafeRelease()
				kinds := map[callKind]bool{}
				external, body := false, false
				for _, p := range i.funcs {
					for pc, d := range p.Code {
						instr, ok := p.ssaInstrs[pc].(*ssa.Defer)
						if !ok || instr.DeferStack == nil {
							continue
						}
						base := instructionCallMetadata(d.Aux)
						if base.DeferStack < 0 || base.DeferStack != descriptionRegister(p, instr.DeferStack) {
							t.Fatal("missing enclosing defer stack")
						}
						kinds[base.Kind] = true
						switch d.Aux.(type) {
						case externalCallAux:
							external = true
						case bodyCallAux:
							body = true
						}
						// Missing token must retain lowering for every target kind.
						encoded, exists := p.index[instr.DeferStack]
						if !exists {
							t.Fatal("unallocated token")
						}
						delete(p.index, instr.DeferStack)
						missing := describeInstruction(p, instr, nil, nil)
						p.index[instr.DeferStack] = encoded
						if _, ok := missing.Aux.(callAux); !ok || makeInstruction(i, missing) != nil {
							t.Fatal("missing token replaced lowering")
						}
					}
				}
				if mode != 0 && (!kinds[callBuiltin] || !kinds[callStatic] || !kinds[callDynamic] || !kinds[callInvoke] || !external || !body) {
					t.Fatalf("missing target kinds %v external=%v body=%v", kinds, external, body)
				}
				for _, p := range i.funcs {
					p.functionCompilation = functionCompilation{}
				}
				if got, err := i.RunFunc("Run"); err != nil || got != 64297531 {
					t.Fatalf("Run=%v,%v", got, err)
				}
			})
		}
	}
}

func TestDescriptionBoxedDeferStack(t *testing.T) {
	var head *_defer
	ext := reflect.ValueOf(func(int) {})
	aux := externalCallAux{callAux: callAux{callSiteAux: callSiteAux{Kind: callStatic, DeferStack: 0, Arguments: []register{1}}}, External: ext}
	run := makeInstruction(&Interp{}, typedInstruction{Op: opDefer, Aux: aux})
	fr := &frame{stack: []value{unsafe.Pointer(&head), 42}}
	run(fr)
	if head == nil || head.args[0] != 42 || fr._defer != nil {
		t.Fatal("boxed token used wrong stack")
	}
	fr.stack[0] = nil
	defer func() {
		if got := recover(); got != "range-over-func defer: missing deferstack token" {
			t.Fatalf("panic=%v", got)
		}
	}()
	run(fr)
}

func TestDescriptionDynamicCalls(t *testing.T) {
	const source = `package dynamic
 import "strings"
 var F func(int)(int,int)
 func factory(n int)func(int)(int,int){return func(x int)(int,int){n+=x;return n,n+1}}
 func call(f func(int)(int,int),n int)(int,int){return f(n)}
 func recovered(f func())(r int){defer func(){if recover()!=nil{r=7}}();f();return}
 var D func(int)
 func deferred()(r int){D=func(n int){r=r*10+n};f:=D;defer f(2);f=func(n int){r=99};return 1}
 func launch(f func(int),n int){go f(n)}
 func Run()int{
 F=factory(10);a,b:=call(F,2);c,_:=F(3)
 ch:=make(chan int,1);send:=func(n int){ch<-n};launch(send,4);send=func(int){panic("wrong target")};x:=<-ch
 ext:=strings.ToUpper
 return a+b+c+x+len(ext("ab"))+deferred()+recovered(nil)
 }`
	for _, mode := range []Mode{0, EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		for _, extra := range []Mode{0, EnableCachedReg, ExperimentalSupportGC | EnableGCExecutionPlan, DisableDynamicFuncCallAnalysis} {
			t.Run(fmt.Sprintf("mode=%v/extra=%v", mode, extra), func(t *testing.T) {
				i, err := NewContext(DisableMemo|mode|extra).LoadInterp("dynamic.go", source)
				if err != nil {
					t.Fatal(err)
				}
				defer i.UnsafeRelease()
				ops := map[instructionOpcode]bool{}
				for _, p := range i.funcs {
					for _, d := range p.Code {
						if _, ok := d.Aux.(dynamicCallAux); ok {
							if makeInstruction(i, d) == nil {
								t.Fatal("missing dynamic executor")
							}
							ops[d.Op] = true
						}
					}
					p.functionCompilation = functionCompilation{}
				}
				if mode != 0 && (!ops[opCall] || !ops[opGo] || !ops[opDefer]) {
					t.Fatalf("missing dynamic ops %v", ops)
				}
				for n := 0; n < 3; n++ {
					if got, err := i.RunFunc("Run"); err != nil || got != 65 {
						t.Fatalf("Run=%v,%v want 65", got, err)
					}
				}
			})
		}
	}
	for _, op := range []instructionOpcode{opCall, opGo, opDefer} {
		if makeInstruction(nil, typedInstruction{Op: op, Aux: dynamicCallAux{callAux: callAux{callSiteAux: callSiteAux{Target: -1}}}}) != nil {
			t.Fatal("unresolved dynamic target")
		}
	}
}

func TestDescriptionPreservesImmediateClosures(t *testing.T) {
	const source = `package immediate
 func Run()(r int){
 x:=1
 r=func(n int)int{return x+n}(2)
 defer func(){r+=x}()
 ch:=make(chan int,1);go func(){ch<-x}();r+=<-ch
 return
 }`
	i, err := NewContext(DisableMemo|EnableInstructionExecution).LoadInterp("immediate.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	seen := map[instructionOpcode]bool{}
	for _, p := range i.funcs {
		for pc, d := range p.Code {
			var call *ssa.CallCommon
			switch instr := p.ssaInstrs[pc].(type) {
			case *ssa.Call:
				call = &instr.Call
			case *ssa.Go:
				call = &instr.Call
			case *ssa.Defer:
				call = &instr.Call
			default:
				continue
			}
			if _, ok := call.Value.(*ssa.MakeClosure); !ok {
				continue
			}
			if _, ok := d.Aux.(callAux); !ok || makeInstruction(i, d) != nil {
				t.Fatal("immediate closure replaced dedicated lowering")
			}
			seen[d.Op] = true
		}
		p.functionCompilation = functionCompilation{}
	}
	if !seen[opCall] || !seen[opGo] || !seen[opDefer] {
		t.Fatalf("missing immediate ops %v", seen)
	}
	if got, err := i.RunFunc("Run"); err != nil || got != 5 {
		t.Fatalf("Run=%v,%v", got, err)
	}
}
