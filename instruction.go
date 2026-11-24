package ixgo

import "golang.org/x/tools/go/ssa"

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

// typedInstruction is the first execution-description form. Args and Aux are
// explicit compiler output; no runtime execution path needs to inspect SSA to
// recover the register operands represented here.
type typedInstruction struct {
	Op     instructionOpcode
	Result register
	Args   []register
	Aux    int32
}

func describeInstruction(pfn *function, instr ssa.Instruction, operands []*ssa.Value) typedInstruction {
	d := typedInstruction{Op: opcodeOf(instr), Result: -1, Aux: -1}
	for _, operand := range operands {
		if operand != nil {
			d.Args = append(d.Args, pfn.regIndex(*operand))
		}
	}
	if producesRegister(instr) {
		if v, ok := instr.(ssa.Value); ok {
			d.Result = pfn.regIndex(v)
		}
	}
	return d
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
