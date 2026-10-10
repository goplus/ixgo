package ixgo

import (
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/ssa"
)

// skipUnusedArrayDeref reports whether v is the unused *array load that
// x/tools SSA emits as the range operand immediately before a rangeindex
// loop. The Go spec does not evaluate a range expression when at most one
// iteration variable is present and len(x) is constant, so ranging over nil
// *[N]T must not panic. See Go issues 72844 and 73476.
//
// Only the trailing unused deref is skipped. An earlier `_ = *p` in the same
// block still panics, matching gc.
func skipUnusedArrayDeref(v *ssa.UnOp) bool {
	if v.Op != token.MUL {
		return false
	}
	block := v.Block()
	// x/tools names the range loop block "rangeindex.*"; interp.go keys off
	// block comments the same way. If the name changes, this just stops firing.
	if block == nil || len(block.Succs) != 1 || !strings.HasPrefix(block.Succs[0].Comment, "rangeindex.") {
		return false
	}
	refs, ok := nonDebugReferrers(v)
	if !ok || len(refs) != 0 {
		return false
	}
	if _, ok := types.Unalias(v.Type()).Underlying().(*types.Array); !ok {
		return false
	}
	return isTrailingRangeArrayDeref(v)
}

// isTrailingRangeArrayDeref reports whether v is immediately followed by a
// jump to rangeindex, optionally with the rangeindex Alloc/Store prologue
// that x/tools emits before blockopt.
func isTrailingRangeArrayDeref(v *ssa.UnOp) bool {
	block := v.Block()
	seen := false
	for _, instr := range block.Instrs {
		if !seen {
			if instr == v {
				seen = true
			}
			continue
		}
		switch instr.(type) {
		case *ssa.DebugRef, *ssa.Alloc, *ssa.Store:
			continue
		case *ssa.Jump:
			return true
		default:
			return false
		}
	}
	return false
}

// arrayPointerOperandHasEffectAfter reports whether v contains a function call
// or channel receive that the spec still requires evaluating as part of the
// range operand.
func arrayPointerOperandHasEffectAfter(v ssa.Value, after token.Pos, seen map[ssa.Value]bool) bool {
	if v == nil {
		return false
	}
	if seen == nil {
		seen = make(map[ssa.Value]bool)
	}
	if seen[v] {
		return false
	}
	seen[v] = true

	instr, ok := v.(ssa.Instruction)
	if !ok {
		return false
	}
	if pos := instr.Pos(); after.IsValid() && pos.IsValid() && pos <= after {
		// SSA eliminates local assignments. Do not mistake a call that produced
		// the assigned value for a call contained in the range operand itself.
		// after is the '*' position; embedded calls in *f() sit after it.
		return false
	}
	switch v := v.(type) {
	case *ssa.Call:
		return true
	case *ssa.UnOp:
		if v.Op == token.ARROW {
			return true
		}
	}
	for _, operand := range instr.Operands(nil) {
		if operand != nil && arrayPointerOperandHasEffectAfter(*operand, after, seen) {
			return true
		}
	}
	return false
}

// nonDebugReferrers omits DebugRef so GlobalDebug does not keep the unused
// range deref alive.
func nonDebugReferrers(v ssa.Value) (refs []ssa.Instruction, available bool) {
	all := v.Referrers()
	if all == nil {
		return nil, false
	}
	for _, ref := range *all {
		if _, ok := ref.(*ssa.DebugRef); !ok {
			refs = append(refs, ref)
		}
	}
	return refs, true
}
