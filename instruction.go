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

// operandKind distinguishes non-register inputs without retaining SSA values.
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
	ExtractResults []register
}

func (callAux) instructionAux() { return }

// typedInstruction records partial execution metadata alongside the closure
// backend. Each entry corresponds to one generated closure; Phi auxiliary data
// covers the entire group. Args preserve SSA operand order and use -1 for inputs
// without a register; ArgKinds distinguishes absent, static and unallocated
// inputs. Opcode-specific lowering and static target resolution are incremental.
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

func describeInstruction(pfn *function, instr ssa.Instruction, operands []*ssa.Value) typedInstruction {
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
		aux := callAux{}
		if refs := v.Referrers(); refs != nil {
			for _, ref := range *refs {
				ex, ok := ref.(*ssa.Extract)
				if !ok {
					continue
				}
				if _, elided := pfn.elidedExtract[ex]; !elided {
					continue
				}
				if aux.ExtractResults == nil {
					aux.ExtractResults = make([]register, callResultCount(&v.Call))
					for j := range aux.ExtractResults {
						aux.ExtractResults[j] = -1
					}
				}
				if uses := ex.Referrers(); uses != nil && len(*uses) > 0 {
					aux.ExtractResults[ex.Index] = descriptionRegister(pfn, ex)
				}
			}
		}
		d.Aux = aux
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
