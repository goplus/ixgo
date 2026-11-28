//go:build instruction

package ixgo

import (
	"fmt"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
)

// planInstructionLayout reserves frame slots directly from SSA. It does not
// build stack instructions or depend on their elision decisions.
func planInstructionLayout(pfn *function) {
	for _, block := range pfn.Fn.Blocks {
		for _, instr := range block.Instrs {
			if value, ok := instr.(ssa.Value); ok {
				pfn.regIndex(value)
			}
			for _, operand := range instr.Operands(nil) {
				if operand != nil && *operand != nil {
					if _, builtin := (*operand).(*ssa.Builtin); !builtin {
						pfn.regIndex(*operand)
					}
				}
			}
			switch v := instr.(type) {
			case *ssa.Call:
				if builtin, ok := v.Call.Value.(*ssa.Builtin); ok && builtin.Name() == "ssa:deferstack" {
					pfn.hasDeferStack = true
				}
			case *ssa.RunDefers:
				pfn.hasRunDefers = true
			case *ssa.Defer:
				if v.DeferStack != nil {
					pfn.hasDeferStack = true
					owner := deferStackOwner(pfn.Fn)
					if owner == nil {
						panic("range-over-func defer has no enclosing source function")
					}
					if body := pfn.Interp.funcs[owner]; body != nil {
						body.needInject = true
					}
				}
			}
		}
	}
}

func compileInstruction(pfn *function, instr ssa.Instruction) (typedInstruction, func(*frame)) {
	if load, ok := instr.(*ssa.UnOp); ok && load.Op == token.MUL &&
		skipUnusedArrayDeref(load) && !arrayPointerOperandHasEffectAfter(load.X, load.Pos(), nil) {
		return typedInstruction{}, nil
	}
	// One Phi executor performs the complete parallel assignment.
	if phi, ok := instr.(*ssa.Phi); ok && phi.Block().Instrs[0] != phi {
		return typedInstruction{}, nil
	}
	if store, ok := instr.(*ssa.Store); ok {
		if addr, ok := store.Addr.(*ssa.FieldAddr); ok {
			if ptr, ok := addr.X.Type().Underlying().(*types.Pointer); ok {
				if fields, ok := ptr.Elem().Underlying().(*types.Struct); ok && fields.Field(addr.Field).Name() == "_" {
					return typedInstruction{}, nil
				}
			}
		}
		if pfn.Fn.Name() == "init" && pfn.Fn.Synthetic == "package initializer" {
			pfn.Interp.chkinit[store.Addr.String()] = true
		}
	}
	if call, ok := instr.(*ssa.Call); ok {
		if fn, ok := call.Call.Value.(*ssa.Function); ok && fn.Blocks == nil && fn.Pkg != nil && fn.Name() == "init" {
			if _, ok := staticDirectAdapter(pfn.Interp, fn); !ok {
				if _, ok := findExternFunc(pfn.Interp, fn); !ok {
					return typedInstruction{}, nil
				}
			}
		}
	}
	operands := instr.Operands(nil)
	if _, slice := instr.(*ssa.Slice); slice {
		compact := operands[:0]
		for _, operand := range operands {
			if operand != nil && *operand != nil {
				compact = append(compact, operand)
			}
		}
		operands = compact
	}
	d := describeInstruction(pfn, instr, operands, nil)
	bindRuntimeRegisters(&d)
	execute := makeInstruction(pfn.Interp, d)
	if execute == nil {
		panic(fmt.Errorf("instruction backend: unsupported %T: %s", instr, instr))
	}
	return d, execute
}
