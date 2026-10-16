/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package memo

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
)

// ScalarEligible reports whether a function has pure scalar self-recursion.
func ScalarEligible(fn *ssa.Function) bool {
	if fn == nil || len(fn.Params) == 0 || len(fn.FreeVars) != 0 || fn.Recover != nil || fn.Signature.Variadic() {
		return false
	}
	typ, ok := types.Unalias(fn.Params[0].Type()).(*types.Basic)
	if !ok || typ.Info()&(types.IsInteger|types.IsFloat) == 0 {
		return false
	}
	if fn.Signature.Results().Len() == 0 {
		return false
	}
	for i := 0; i < fn.Signature.Results().Len(); i++ {
		if !types.Identical(fn.Signature.Results().At(i).Type(), typ) {
			return false
		}
	}
	recursive := false
	for _, b := range fn.Blocks {
		for _, instr := range b.Instrs {
			switch u := instr.(type) {
			case *ssa.BinOp, *ssa.Phi, *ssa.Extract, *ssa.If, *ssa.Jump, *ssa.Return, *ssa.DebugRef:
			case *ssa.UnOp:
				if u.Op != token.SUB && u.Op != token.XOR && u.Op != token.NOT {
					return false
				}
			case *ssa.Call:
				if u.Call.StaticCallee() != fn || len(u.Call.Args) != len(fn.Params) {
					return false
				}
				for i := 1; i < len(fn.Params); i++ {
					if u.Call.Args[i] != fn.Params[i] {
						return false
					}
				}
				recursive = true
			default:
				return false
			}
			// Additional parameters are permitted only as unchanged forwarded values.
			for _, operand := range instr.Operands(nil) {
				if operand == nil {
					continue
				}
				for _, param := range fn.Params[1:] {
					if *operand == param {
						if _, ok := instr.(*ssa.Call); !ok {
							return false
						}
					}
				}
			}
		}
	}
	return recursive
}

// AccumulatorPlan identifies an additive output derived from the SSA call graph.
// All other scalar values and pointed-to scalar inputs form the cache key.
type AccumulatorPlan struct{ Output int }

// AccumulatorEligible reports whether the call graph has an isolated additive output.
func AccumulatorEligible(fn *ssa.Function) bool { return AnalyzeAccumulator(fn) != nil }

// AnalyzeAccumulator proves one isolated additive output and read-only inputs.
// Parameter count, order, integer widths and recursive call signatures may vary.
func AnalyzeAccumulator(root *ssa.Function) *AccumulatorPlan {
	if root == nil {
		return nil
	}
	for output, param := range root.Params {
		ptr, ok := param.Type().Underlying().(*types.Pointer)
		if !ok || !integerType(ptr.Elem()) {
			continue
		}
		seen := map[*ssa.Function]int{}
		var check func(*ssa.Function, int) bool
		check = func(fn *ssa.Function, index int) bool {
			if previous, ok := seen[fn]; ok {
				return previous == index
			}
			if fn == nil || len(fn.Blocks) == 0 || len(fn.FreeVars) != 0 || fn.Recover != nil || fn.Signature.Variadic() || fn.Signature.Results().Len() != 0 {
				return false
			}
			if index < 0 || index >= len(fn.Params) {
				return false
			}
			seen[fn] = index
			out := fn.Params[index]
			if !types.Identical(out.Type(), param.Type()) {
				return false
			}
			for _, p := range fn.Params {
				t := p.Type()
				if ptr, ok := t.Underlying().(*types.Pointer); ok {
					t = ptr.Elem()
				}
				if !integerType(t) {
					return false
				}
			}
			local := localAddress
			inputPointer := func(v ssa.Value) bool {
				p, ok := v.(*ssa.Parameter)
				if !ok || p == out {
					return false
				}
				for _, input := range fn.Params {
					if input == p {
						_, ok := p.Type().Underlying().(*types.Pointer)
						return ok
					}
				}
				return false
			}
			// An accumulator address may only be forwarded or participate in an
			// isolated load/add/store. In particular it cannot become a recursion input.
			for _, use := range *out.Referrers() {
				switch u := use.(type) {
				case *ssa.Call:
					if !validCall(u, out) {
						return false
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
					case *ssa.Alloc:
						if !localType(u.Type().Underlying().(*types.Pointer).Elem()) {
							return false
						}
					case *ssa.UnOp:
						if u.Op == token.MUL && !inputPointer(u.X) && u.X != out && !local(u.X) {
							return false
						}
						if u.Op != token.MUL && u.Op != token.SUB && u.Op != token.XOR && u.Op != token.NOT {
							return false
						}
					case *ssa.IndexAddr:
						if !local(u.X) {
							return false
						}
					case *ssa.If, *ssa.Jump, *ssa.Return, *ssa.DebugRef:

					case *ssa.Store:
						if u.Addr != out && !local(u.Addr) {
							return false
						}
					case *ssa.Call:
						if !validCall(u, out) {
							return false
						}
						callee := u.Call.StaticCallee()
						if !check(callee, outputArgument(u, out)) {
							return false
						}
						for _, arg := range u.Call.Args {
							if _, ok := arg.Type().Underlying().(*types.Pointer); ok && arg != out && !inputPointer(arg) && !local(arg) {
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
					default:
						return false
					}
				}
			}
			return true
		}
		if check(root, output) {
			return &AccumulatorPlan{Output: output}
		}
	}
	return nil
}

func outputArgument(call *ssa.Call, out ssa.Value) int {
	index := -1
	for i, arg := range call.Call.Args {
		if arg == out {
			if index >= 0 {
				return -1
			}
			index = i
		}
	}
	return index
}
func validCall(call *ssa.Call, out ssa.Value) bool {
	fn := call.Call.StaticCallee()
	return fn != nil && len(call.Call.Args) == len(fn.Params) && outputArgument(call, out) >= 0
}
func integerType(t types.Type) bool {
	b, ok := types.Unalias(t).(*types.Basic)
	return ok && b.Info()&types.IsInteger != 0
}
func localAddress(v ssa.Value) bool {
	switch v := v.(type) {
	case *ssa.Alloc:
		return true
	case *ssa.IndexAddr:
		return localAddress(v.X)
	}
	return false
}
func localType(t types.Type) bool {
	if integerType(t) {
		return true
	}
	if a, ok := t.Underlying().(*types.Array); ok {
		return localType(a.Elem())
	}
	return false
}

func blockInCycle(start *ssa.BasicBlock) bool {
	seen := map[*ssa.BasicBlock]bool{}
	var walk func(*ssa.BasicBlock) bool
	walk = func(b *ssa.BasicBlock) bool {
		if b == start {
			return true
		}
		if seen[b] {
			return false
		}
		seen[b] = true
		for _, next := range b.Succs {
			if walk(next) {
				return true
			}
		}
		return false
	}
	for _, next := range start.Succs {
		if walk(next) {
			return true
		}
	}
	return false
}
