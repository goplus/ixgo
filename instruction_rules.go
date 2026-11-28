package ixgo

import (
	"go/types"

	"golang.org/x/tools/go/ssa"
)

func skipBlankFieldStore(store *ssa.Store) bool {
	addr, ok := store.Addr.(*ssa.FieldAddr)
	if !ok {
		return false
	}
	ptr, ok := addr.X.Type().Underlying().(*types.Pointer)
	if !ok {
		return false
	}
	fields, ok := ptr.Elem().Underlying().(*types.Struct)
	return ok && fields.Field(addr.Field).Name() == "_"
}

func skipMissingInit(interp *Interp, fn *ssa.Function) bool {
	if fn.Blocks != nil || fn.Pkg == nil || fn.Name() != "init" {
		return false
	}
	// A direct adapter can provide an initializer without a reflected function.
	if _, ok := staticDirectAdapter(interp, fn); ok {
		return false
	}
	_, ok := findExternFunc(interp, fn)
	return !ok
}

func skipArrayRangeDeref(load *ssa.UnOp) bool {
	return skipUnusedArrayDeref(load) && !arrayPointerOperandHasEffectAfter(load.X, load.Pos(), nil)
}

// One Phi executor performs the complete parallel assignment.
func skipNonleadingPhi(phi *ssa.Phi) bool {
	return phi.Block().Instrs[0] != phi
}

func recordPackageInitStore(pfn *function, store *ssa.Store) {
	if pfn.Fn.Name() == "init" && pfn.Fn.Synthetic == "package initializer" {
		pfn.Interp.chkinit[store.Addr.String()] = true
	}
}
