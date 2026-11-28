//go:build instruction

package ixgo

import (
	"fmt"

	"golang.org/x/tools/go/ssa"
)

// planInstructionLayout reserves frame slots directly from SSA. It does not
// build stack instructions or depend on their elision decisions.
func planInstructionLayout(pfn *function) {
	var buf [32]*ssa.Value
	for _, block := range pfn.Fn.Blocks {
		for _, instr := range block.Instrs {
			if value, ok := instr.(ssa.Value); ok {
				pfn.regIndex(value)
			}
			for _, operand := range instr.Operands(buf[:0]) {
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
			case *ssa.FieldAddr:
				if pfn.Interp.ctx.Mode&EnableInstructionExecution != 0 && pfn.Interp.ctx.Mode&EnableCachedReg != 0 {
					if pfn.fieldAddrCache == nil {
						pfn.fieldAddrCache = make(map[*ssa.FieldAddr]register)
					}
					pfn.fieldAddrCache[v] = pfn.newCacheReg()
				}
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
	// Destination registers must exist before call descriptions inspect referrers.
	for _, block := range pfn.Fn.Blocks {
		for _, instr := range block.Instrs {
			call, ok := instr.(*ssa.Call)
			if !ok || describeCallResults(pfn, call) == nil {
				continue
			}
			if pfn.elidedExtract == nil {
				pfn.elidedExtract = make(map[*ssa.Extract]struct{})
			}
			if refs := call.Referrers(); refs != nil {
				for _, ref := range *refs {
					if extract, ok := ref.(*ssa.Extract); ok {
						pfn.elidedExtract[extract] = struct{}{}
					}
				}
			}
			pfn.elidedTuples = append(pfn.elidedTuples, pfn.regIndex(call))
		}
	}
}

func compileInstruction(pfn *function, instr ssa.Instruction) (typedInstruction, func(*frame)) {
	if extract, ok := instr.(*ssa.Extract); ok {
		if _, elided := pfn.elidedExtract[extract]; elided {
			return typedInstruction{}, nil
		}
	}
	if load, ok := instr.(*ssa.UnOp); ok && skipArrayRangeDeref(load) {
		return typedInstruction{}, nil
	}
	if phi, ok := instr.(*ssa.Phi); ok && skipNonleadingPhi(phi) {
		return typedInstruction{}, nil
	}
	if store, ok := instr.(*ssa.Store); ok {
		if skipBlankFieldStore(store) {
			return typedInstruction{}, nil
		}
		recordPackageInitStore(pfn, store)
	}
	if call, ok := instr.(*ssa.Call); ok {
		if fn, ok := call.Call.Value.(*ssa.Function); ok && skipMissingInit(pfn.Interp, fn) {
			return typedInstruction{}, nil
		}
	}
	var buf [32]*ssa.Value
	operands := instr.Operands(buf[:0])
	if _, slice := instr.(*ssa.Slice); slice {
		compact := operands[:0]
		for _, operand := range operands {
			if operand != nil && *operand != nil {
				compact = append(compact, operand)
			}
		}
		operands = compact
	}
	var dests []register
	if call, ok := instr.(*ssa.Call); ok {
		dests = describeCallResults(pfn, call)
	}
	d := describeInstruction(pfn, instr, operands, dests)
	bindRuntimeRegisters(&d)
	execute := makeInstruction(pfn.Interp, d)
	if execute == nil {
		panic(fmt.Errorf("instruction backend: unsupported %T: %s", instr, instr))
	}
	return d, execute
}
