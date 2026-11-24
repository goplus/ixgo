package ixgo

import (
	"go/token"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/ssa"
)

// instructionOpcode identifies an execution instruction independently of its
// SSA representation. The closure backend remains the executor while this
// description is introduced as the migration boundary toward bytecode.
type instructionOpcode uint16

const (
	opInvalid instructionOpcode = iota
	opAlloc
	opPhi
	opCall
	opBinOp
	opUnOp
	opChangeInterface
	opChangeType
	opConvert
	opMakeInterface
	opMakeClosure
	opMakeChan
	opMakeMap
	opMakeSlice
	opSlice
	opFieldAddr
	opField
	opIndexAddr
	opIndex
	opLookup
	opSelect
	opSliceToArrayPointer
	opRange
	opNext
	opTypeAssert
	opExtract
	opJump
	opIf
	opReturn
	opRunDefers
	opPanic
	opGo
	opDefer
	opSend
	opStore
	opMapUpdate
	opDebugRef
)

// instructionAux stores metadata whose shape depends on the opcode. Keeping
// it typed makes each instruction's execution inputs explicit before encoding.
type instructionAux interface {
	instructionAux()
}

type binaryOpAux struct{ Op token.Token }

func (binaryOpAux) instructionAux() { return }

type unaryOpAux struct {
	Op      token.Token
	CommaOk bool
}

func (unaryOpAux) instructionAux() { return }

// phiAux describes the parallel assignment performed by one grouped closure.
// Edges[i][j] supplies Results[i] when control arrives from Preds[j].
type phiAux struct {
	Preds   []int32
	Results []register
	Edges   [][]register
}

func (phiAux) instructionAux() { return }

type branchAux struct{ Targets []int32 }

func (branchAux) instructionAux() { return }

type returnAux struct{ Results int }

func (returnAux) instructionAux() { return }

type extractAux struct{ Index int }

func (extractAux) instructionAux() { return }

type nextAux struct{ IsString bool }

func (nextAux) instructionAux() { return }

// fieldAux.Source is the operand type, a pointer for FieldAddr. Field and
// Offset describe the field within the struct, after dereferencing Source;
// Offset is relative to the start of that struct.
type fieldAux struct {
	Index  int
	Source reflect.Type
	Field  reflect.Type
	Offset uintptr
}

func (fieldAux) instructionAux() { return }

// lookupAux records map key/element types. For a string Source, Key is the
// index operand type and Element is byte.
type lookupAux struct {
	Source  reflect.Type
	Key     reflect.Type
	Element reflect.Type
	CommaOk bool
}

func (lookupAux) instructionAux() { return }

// indexAux.Source retains fixed array lengths; slice/string lengths and index
// bounds are checked against the runtime value. Target is the produced type.
type indexAux struct {
	Source reflect.Type
	Target reflect.Type
}

func (indexAux) instructionAux() { return }

// Slice Args are source, low, high, max, including -1 for omitted bounds.
// MakeSliceCheck preserves the checks used when slicing an allocation.
type sliceAux struct {
	Source         reflect.Type
	Target         reflect.Type
	Low, High, Max bool
	MakeSliceCheck bool
}

func (sliceAux) instructionAux() { return }

// rangeAux records map iterator types. For a string Source, Key and Element
// are int and rune, respectively.
type rangeAux struct {
	Source  reflect.Type
	Key     reflect.Type
	Element reflect.Type
}

func (rangeAux) instructionAux() { return }

// Reflect types are the in-memory runtime representation. A portable type table
// remains future work; these descriptions retain no SSA instruction or value.
type typeAssertAux struct {
	Source  reflect.Type
	Target  reflect.Type
	CommaOk bool
}

func (typeAssertAux) instructionAux() { return }

// conversionAux records in-memory types, independently of the SSA values.
// Conversion behavior is selected by the instruction opcode and these types.
type conversionAux struct {
	Source reflect.Type
	Target reflect.Type
}

func (conversionAux) instructionAux() { return }

// creationAux.Target is the produced type (a pointer for Alloc). Heap is only
// meaningful for Alloc; sizes and capacities retain their SSA order in Args.
type creationAux struct {
	Target reflect.Type
	Heap   bool
}

func (creationAux) instructionAux() { return }

// selectStateAux describes one channel operation in a Select instruction.
type selectStateAux struct {
	Chan reflect.Type
	Send reflect.Type
	Dir  reflect.SelectDir
}

type selectAux struct {
	States   []selectStateAux
	Blocking bool
}

func (selectAux) instructionAux() { return }

type sendAux struct {
	Chan  reflect.Type
	Value reflect.Type
}

func (sendAux) instructionAux() { return }

type storeAux struct {
	Address reflect.Type
	Value   reflect.Type
}

func (storeAux) instructionAux() { return }

type mapUpdateAux struct {
	Map   reflect.Type
	Key   reflect.Type
	Value reflect.Type
}

func (mapUpdateAux) instructionAux() { return }

type panicAux struct{ Value reflect.Type }

func (panicAux) instructionAux() { return }

// operandKind classifies the source of an input independently from its register.
// operandFunction and operandBuiltin may still have a valid Args register; only
// operandAbsent and operandUnallocated imply Args == -1.
type operandKind uint8

const (
	operandAbsent operandKind = iota
	operandRegister
	operandFunction
	operandBuiltin
	operandUnallocated
)

// callAux describes result placement after tuple extraction elision. A nil
// ExtractResults means the call still uses its ordinary result register. A -1
// entry denotes an unused result. Static target resolution remains incremental.
type callAux struct {
	// ExtractResults is a copy of the destinations used by the call closure.
	// A nil slice means the call writes its ordinary result register.
	ExtractResults []register
}

func (callAux) instructionAux() { return }

// typedInstruction records partial execution metadata alongside the closure
// backend. Each entry corresponds to one generated closure; Phi auxiliary data
// covers the entire group; the common Result and Args fields still refer to
// the first Phi entry. Args preserve SSA operand order; ArgKinds distinguishes
// the source class independently from whether Args has a valid register.
// Opcode-specific lowering and static target resolution are incremental.
type typedInstruction struct {
	Op       instructionOpcode
	Result   register
	Args     []register
	ArgKinds []operandKind
	Aux      instructionAux
}

func descriptionRegister(pfn *function, v ssa.Value) register {
	if v != nil {
		if encoded, ok := pfn.index[v]; ok {
			return register(encoded & 0xffffff)
		}
	}
	return -1
}

func describeInstruction(pfn *function, instr ssa.Instruction, operands []*ssa.Value, elidedResults []register) typedInstruction {
	d := typedInstruction{Op: opcodeOf(instr), Result: -1}
	for _, operand := range operands {
		if operand != nil {
			r := descriptionRegister(pfn, *operand)
			kind := operandRegister
			switch (*operand).(type) {
			case nil:
				kind = operandAbsent
			case *ssa.Function:
				kind = operandFunction
			case *ssa.Builtin:
				kind = operandBuiltin
			default:
				if r < 0 {
					kind = operandUnallocated
				}
			}
			d.Args = append(d.Args, r)
			d.ArgKinds = append(d.ArgKinds, kind)
		}
	}
	if producesRegister(instr) {
		if v, ok := instr.(ssa.Value); ok {
			d.Result = descriptionRegister(pfn, v)
		}
	}
	switch v := instr.(type) {
	case *ssa.BinOp:
		d.Aux = binaryOpAux{Op: v.Op}
	case *ssa.UnOp:
		d.Aux = unaryOpAux{Op: v.Op, CommaOk: v.CommaOk}
	case *ssa.Phi:
		aux := phiAux{}
		if block := v.Block(); block != nil {
			aux.Preds = make([]int32, len(block.Preds))
			for i, pred := range block.Preds {
				aux.Preds[i] = int32(pred.Index)
			}
			for _, instr := range block.Instrs {
				phi, ok := instr.(*ssa.Phi)
				if !ok {
					break
				}
				aux.Results = append(aux.Results, descriptionRegister(pfn, phi))
				edges := make([]register, len(phi.Edges))
				for i, edge := range phi.Edges {
					edges[i] = descriptionRegister(pfn, edge)
				}
				aux.Edges = append(aux.Edges, edges)
			}
		}
		d.Aux = aux
	case *ssa.Call:
		if elidedResults != nil {
			d.Result = -1
		}
		d.Aux = callAux{ExtractResults: elidedResults}
	case *ssa.ChangeType:
		d.Aux = describeConversion(pfn, v.X, v.Type())
	case *ssa.Convert:
		d.Aux = describeConversion(pfn, v.X, v.Type())
	case *ssa.ChangeInterface:
		d.Aux = describeConversion(pfn, v.X, v.Type())
	case *ssa.MakeInterface:
		d.Aux = describeConversion(pfn, v.X, v.Type())
	case *ssa.SliceToArrayPointer:
		d.Aux = describeConversion(pfn, v.X, v.Type())
	case *ssa.Alloc:
		d.Aux = creationAux{Target: pfn.Interp.preToType(v.Type()), Heap: v.Heap}
	case *ssa.MakeSlice:
		d.Aux = creationAux{Target: pfn.Interp.preToType(v.Type())}
	case *ssa.MakeMap:
		d.Aux = creationAux{Target: pfn.Interp.preToType(v.Type())}
	case *ssa.MakeChan:
		d.Aux = creationAux{Target: pfn.Interp.preToType(v.Type())}
	case *ssa.Field:
		d.Aux = describeField(pfn, v.X, v.Field)
	case *ssa.FieldAddr:
		d.Aux = describeField(pfn, v.X, v.Field)
	case *ssa.Lookup:
		source := pfn.Interp.preToType(v.X.Type())
		aux := lookupAux{Source: source, CommaOk: v.CommaOk}
		if source.Kind() == reflect.Map {
			aux.Key, aux.Element = source.Key(), source.Elem()
		} else {
			aux.Key, aux.Element = pfn.Interp.preToType(v.Index.Type()), reflect.TypeOf(byte(0))
		}
		d.Aux = aux
	case *ssa.Index:
		d.Aux = indexAux{Source: pfn.Interp.preToType(v.X.Type()), Target: pfn.Interp.preToType(v.Type())}
	case *ssa.IndexAddr:
		d.Aux = indexAux{Source: pfn.Interp.preToType(v.X.Type()), Target: pfn.Interp.preToType(v.Type())}
	case *ssa.Slice:
		_, allocation := v.X.(*ssa.Alloc)
		d.Aux = sliceAux{Source: pfn.Interp.preToType(v.X.Type()), Target: pfn.Interp.preToType(v.Type()), Low: v.Low != nil, High: v.High != nil, Max: v.Max != nil, MakeSliceCheck: allocation}
	case *ssa.Range:
		source := pfn.Interp.preToType(v.X.Type())
		aux := rangeAux{Source: source}
		if source.Kind() == reflect.Map {
			aux.Key, aux.Element = source.Key(), source.Elem()
		} else {
			aux.Key, aux.Element = reflect.TypeOf(int(0)), reflect.TypeOf(rune(0))
		}
		d.Aux = aux
	case *ssa.Select:
		aux := selectAux{Blocking: v.Blocking, States: make([]selectStateAux, len(v.States))}
		for i, state := range v.States {
			chanType := pfn.Interp.preToType(state.Chan.Type())
			dir := reflect.SelectRecv
			if state.Dir == types.SendOnly {
				dir = reflect.SelectSend
			}
			var sendType reflect.Type
			if state.Send != nil {
				sendType = pfn.Interp.preToType(state.Send.Type())
			}
			aux.States[i] = selectStateAux{Chan: chanType, Send: sendType, Dir: dir}
		}
		d.Aux = aux
	case *ssa.Send:
		d.Aux = sendAux{Chan: pfn.Interp.preToType(v.Chan.Type()), Value: pfn.Interp.preToType(v.X.Type())}
	case *ssa.Store:
		d.Aux = storeAux{Address: pfn.Interp.preToType(v.Addr.Type()), Value: pfn.Interp.preToType(v.Val.Type())}
	case *ssa.MapUpdate:
		d.Aux = mapUpdateAux{Map: pfn.Interp.preToType(v.Map.Type()), Key: pfn.Interp.preToType(v.Key.Type()), Value: pfn.Interp.preToType(v.Value.Type())}
	case *ssa.Panic:
		d.Aux = panicAux{Value: pfn.Interp.preToType(v.X.Type())}
	case *ssa.TypeAssert:
		d.Aux = typeAssertAux{
			Source:  pfn.Interp.preToType(v.X.Type()),
			Target:  pfn.Interp.preToType(v.AssertedType),
			CommaOk: v.CommaOk,
		}
	case *ssa.Extract:
		d.Aux = extractAux{Index: v.Index}
	case *ssa.Next:
		d.Aux = nextAux{IsString: v.IsString}
	case *ssa.Return:
		d.Aux = returnAux{Results: len(v.Results)}
	case *ssa.If:
		if block := v.Block(); block != nil && len(block.Succs) >= 2 {
			d.Aux = branchAux{Targets: []int32{
				int32(block.Succs[0].Index),
				int32(block.Succs[1].Index),
			}}
		}
	case *ssa.Jump:
		if block := v.Block(); block != nil && len(block.Succs) >= 1 {
			d.Aux = branchAux{Targets: []int32{int32(block.Succs[0].Index)}}
		}
	}
	return d
}

func describeField(pfn *function, source ssa.Value, index int) fieldAux {
	typ := pfn.Interp.preToType(source.Type())
	record := typ
	if record.Kind() == reflect.Pointer {
		record = record.Elem()
	}
	field := record.Field(index)
	return fieldAux{Index: index, Source: typ, Field: field.Type, Offset: field.Offset}
}

func describeConversion(pfn *function, source ssa.Value, target types.Type) conversionAux {
	return conversionAux{
		Source: pfn.Interp.preToType(source.Type()),
		Target: pfn.Interp.preToType(target),
	}
}

func producesRegister(instr ssa.Instruction) bool {
	switch instr.(type) {
	case *ssa.Alloc, *ssa.Phi, *ssa.Call, *ssa.BinOp, *ssa.UnOp,
		*ssa.ChangeInterface, *ssa.ChangeType, *ssa.Convert, *ssa.MakeInterface,
		*ssa.MakeClosure, *ssa.MakeChan, *ssa.MakeMap, *ssa.MakeSlice, *ssa.Slice,
		*ssa.FieldAddr, *ssa.Field, *ssa.IndexAddr, *ssa.Index, *ssa.Lookup,
		*ssa.Select, *ssa.SliceToArrayPointer, *ssa.Range, *ssa.Next,
		*ssa.TypeAssert, *ssa.Extract:
		return true
	default:
		return false
	}
}

func opcodeOf(instr ssa.Instruction) instructionOpcode {
	switch instr.(type) {
	case *ssa.Alloc:
		return opAlloc
	case *ssa.Phi:
		return opPhi
	case *ssa.Call:
		return opCall
	case *ssa.BinOp:
		return opBinOp
	case *ssa.UnOp:
		return opUnOp
	case *ssa.ChangeInterface:
		return opChangeInterface
	case *ssa.ChangeType:
		return opChangeType
	case *ssa.Convert:
		return opConvert
	case *ssa.MakeInterface:
		return opMakeInterface
	case *ssa.MakeClosure:
		return opMakeClosure
	case *ssa.MakeChan:
		return opMakeChan
	case *ssa.MakeMap:
		return opMakeMap
	case *ssa.MakeSlice:
		return opMakeSlice
	case *ssa.Slice:
		return opSlice
	case *ssa.FieldAddr:
		return opFieldAddr
	case *ssa.Field:
		return opField
	case *ssa.IndexAddr:
		return opIndexAddr
	case *ssa.Index:
		return opIndex
	case *ssa.Lookup:
		return opLookup
	case *ssa.Select:
		return opSelect
	case *ssa.SliceToArrayPointer:
		return opSliceToArrayPointer
	case *ssa.Range:
		return opRange
	case *ssa.Next:
		return opNext
	case *ssa.TypeAssert:
		return opTypeAssert
	case *ssa.Extract:
		return opExtract
	case *ssa.Jump:
		return opJump
	case *ssa.If:
		return opIf
	case *ssa.Return:
		return opReturn
	case *ssa.RunDefers:
		return opRunDefers
	case *ssa.Panic:
		return opPanic
	case *ssa.Go:
		return opGo
	case *ssa.Defer:
		return opDefer
	case *ssa.Send:
		return opSend
	case *ssa.Store:
		return opStore
	case *ssa.MapUpdate:
		return opMapUpdate
	case *ssa.DebugRef:
		return opDebugRef
	default:
		return opInvalid
	}
}
