package ixgo

import (
	"go/token"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/ssa"
)

type typedAccumulatorKey[T TypedScalar] struct {
	runner  *typedArgsRunner[T]
	x, a, b T
}

// Only integer additive effects can be replayed independently of the initial
// accumulator value. Floating-point regrouping would change rounding.
func enableTypedAccumulatorMemo[T TypedScalar](root *typedArgsRunner[T]) {
	var zero T
	kind := reflect.TypeOf(zero).Kind()
	if kind == reflect.Float32 || kind == reflect.Float64 {
		return
	}
	seen := map[*typedArgsRunner[T]]bool{}
	var check func(*typedArgsRunner[T]) bool
	check = func(r *typedArgsRunner[T]) bool {
		if seen[r] {
			return true
		}
		seen[r] = true
		fn := r.source
		if fn == nil || len(fn.Params) != 4 || len(r.resultKinds) != 0 || len(r.kinds) != 4 || r.kinds[0] != typedArgScalar {
			return false
		}
		for i := 1; i < 4; i++ {
			if r.kinds[i] != typedArgPointer {
				return false
			}
		}
		out := fn.Params[3]
		var local func(ssa.Value) bool
		local = func(v ssa.Value) bool {
			switch v := v.(type) {
			case *ssa.Alloc:
				return true
			case *ssa.IndexAddr:
				return local(v.X)
			}
			return false
		}
		// An accumulator address may only be forwarded or participate in an
		// isolated load/add/store. In particular it cannot become a recursion input.
		for _, use := range *out.Referrers() {
			switch u := use.(type) {
			case *ssa.Call:
				if len(u.Call.Args) != 4 || u.Call.Args[3] != out {
					return false
				}
				for _, a := range u.Call.Args[:3] {
					if a == out {
						return false
					}
				}
			case *ssa.UnOp:
				if u.Op != token.MUL || u.X != out || len(*u.Referrers()) != 1 {
					return false
				}
				add, ok := (*u.Referrers())[0].(*ssa.BinOp)
				if !ok || add.Op != token.ADD || len(*add.Referrers()) != 1 {
					return false
				}
				other := add.X
				if other == u {
					other = add.Y
				}
				if load, ok := other.(*ssa.UnOp); ok && load.Op == token.MUL && load.X == out {
					return false
				}

				store, ok := (*add.Referrers())[0].(*ssa.Store)
				if !ok || store.Addr != out || store.Val != add || store.Block() != u.Block() {
					return false
				}
				// No intervening call or write may alter the value read by the addition.
				active := false
				for _, instr := range u.Block().Instrs {
					if instr == u {
						active = true
						continue
					}
					if instr == store {
						break
					}
					if active {
						switch instr.(type) {
						case *ssa.Call, *ssa.Store:
							return false
						}
					}
				}
			case *ssa.Store:
				add, ok := u.Val.(*ssa.BinOp)
				if !ok || add.Op != token.ADD || u.Addr != out {
					return false
				}
				load, ok := add.X.(*ssa.UnOp)
				if !ok || load.X != out || load.Op != token.MUL {
					load, ok = add.Y.(*ssa.UnOp)
					if !ok || load.X != out || load.Op != token.MUL {
						return false
					}
				}
			default:
				return false
			}
		}
		for _, b := range fn.Blocks {
			for _, instr := range b.Instrs {
				if blockInCycle(b) {
					return false
				}
				switch u := instr.(type) {
				case *ssa.Store:
					if u.Addr != out && !local(u.Addr) {
						return false
					}
				case *ssa.Call:
					if len(u.Call.Args) != 4 || u.Call.Args[3] != out {
						return false
					}
					for _, a := range u.Call.Args[1:3] {
						if a != fn.Params[1] && a != fn.Params[2] && !local(a) {
							return false
						}
					}
				case *ssa.Phi:
					if _, ok := u.Type().Underlying().(*types.Pointer); ok {
						return false
					}
				case *ssa.ChangeType, *ssa.Convert:
					return false
				case *ssa.BinOp:
					if _, ok := u.X.Type().Underlying().(*types.Pointer); ok {
						return false
					}
				}
			}
		}
		for _, op := range r.code {
			if op.kind == typedCall && !check(op.callee) {
				return false
			}
		}
		return true
	}
	if !check(root) {
		return
	}
	for r := range seen {
		r.accumulatorMemo = true
	}
}
