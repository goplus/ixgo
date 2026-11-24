package ixgo

import (
	"go/token"

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

func (binaryOpAux) instructionAux() {}

type unaryOpAux struct {
	Op      token.Token
	CommaOk bool
}

func (unaryOpAux) instructionAux() {}

type phiAux struct{ Preds []int32 }

func (phiAux) instructionAux() {}

type branchAux struct{ Targets []int32 }

func (branchAux) instructionAux() {}

type returnAux struct{ Results int }

func (returnAux) instructionAux() {}

// typedInstruction is the first execution-description form. Args and Aux are
// explicit compiler output; no runtime execution path needs to inspect SSA to
// recover the register operands represented here.
type typedInstruction struct {
	Op     instructionOpcode
	Result register
	Args   []register
	Aux    instructionAux
}

func describeInstruction(pfn *function, instr ssa.Instruction, operands []*ssa.Value) typedInstruction {
	d := typedInstruction{Op: opcodeOf(instr), Result: -1}
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
	switch v := instr.(type) {
	case *ssa.BinOp:
		d.Aux = binaryOpAux{Op: v.Op}
	case *ssa.UnOp:
		d.Aux = unaryOpAux{Op: v.Op, CommaOk: v.CommaOk}
	case *ssa.Phi:
		aux := phiAux{Preds: make([]int32, 0)}
		if block := v.Block(); block != nil {
			aux.Preds = make([]int32, len(block.Preds))
			for i, pred := range block.Preds {
				aux.Preds[i] = int32(pred.Index)
			}
		}
		d.Aux = aux
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
