/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package memo

import (
	"go/token"
	"go/types"
	"sort"

	"golang.org/x/tools/go/ssa"
)

// receiver describes only numerical leaves accessed by a closed call graph.
// emptyreceiver is shared by ordinary functions and empty value receivers.
type receiver struct {
	Fields   [][]int
	Additive []bool
	Affine   []bool
	Bitwise  []bool
}

var emptyreceiver = &receiver{}

// receiverOf models the receiver independently of the ordinary parameters.
func receiverOf(fn *ssa.Function) *receiver {
	if fn == nil || fn.Signature.Recv() == nil {
		return emptyreceiver
	}
	t := fn.Signature.Recv().Type().Underlying()
	if s, ok := t.(*types.Struct); ok && s.NumFields() == 0 {
		return emptyreceiver
	}
	return &receiver{}
}

type receiverPlan struct {
	receiver  *receiver
	Functions map[*ssa.Function]bool
	Updates   map[*ssa.Store]receiverUpdate
}

// receiverUpdate identifies an isolated update; unary operations have no operand.
type receiverUpdate struct {
	Field   int
	Op      token.Token
	Operand ssa.Value
}

// analyzeReceiver proves a closed graph with one unchanged receiver, read-only
// scalar pointer arguments, numerical results, and no externally visible effects
// other than writes to the recorded receiver leaves.
func analyzeReceiver(root *ssa.Function) *receiverPlan {
	if root == nil || root.Signature.Recv() == nil {
		return nil
	}
	recvType := root.Signature.Recv().Type()
	empty := receiverOf(root) == emptyreceiver
	if !empty {
		ptr, ok := recvType.Underlying().(*types.Pointer)
		if !ok {
			return nil
		}
		if _, ok := ptr.Elem().Underlying().(*types.Struct); !ok {
			return nil
		}
	}
	plan := &receiverPlan{receiver: receiverOf(root), Functions: map[*ssa.Function]bool{}, Updates: map[*ssa.Store]receiverUpdate{}}
	fields := map[string][]int{}
	addresses := map[*ssa.Function]map[ssa.Value][]int{}
	var check func(*ssa.Function) bool
	check = func(fn *ssa.Function) bool {
		if plan.Functions[fn] {
			return true
		}
		if fn == nil || fn.Signature.Recv() == nil || !types.Identical(fn.Signature.Recv().Type(), recvType) || len(fn.Blocks) == 0 || len(fn.FreeVars) != 0 || fn.Recover != nil || fn.Signature.Variadic() {
			return false
		}
		plan.Functions[fn] = true
		recv := fn.Params[0]
		for _, p := range fn.Params[1:] {
			t := p.Type()
			if ptr, ok := t.Underlying().(*types.Pointer); ok {
				t = ptr.Elem()
			}
			if !numericType(t, false) {
				return false
			}
		}
		for i := 0; i < fn.Signature.Results().Len(); i++ {
			if !numericType(fn.Signature.Results().At(i).Type(), false) {
				return false
			}
		}
		paths := map[ssa.Value][]int{recv: {}}
		addresses[fn] = paths
		input := func(v ssa.Value) bool {
			for _, p := range fn.Params[1:] {
				if p == v {
					_, ok := p.Type().Underlying().(*types.Pointer)
					return ok
				}
			}
			return false
		}
		for _, b := range fn.Blocks {
			if blockInCycle(b) {
				return false
			}
			for _, instr := range b.Instrs {
				switch u := instr.(type) {
				case *ssa.FieldAddr:
					path, ok := paths[u.X]
					if !ok || empty {
						return false
					}
					path = append(append([]int(nil), path...), u.Field)
					paths[u] = path
				case *ssa.Alloc:
					if !localNumericType(u.Type().Underlying().(*types.Pointer).Elem(), false) {
						return false
					}
				case *ssa.IndexAddr:
					if !localAddress(u.X) {
						return false
					}
				case *ssa.UnOp:
					if u.Op == token.MUL {
						if path, ok := paths[u.X]; ok {
							if len(path) == 0 || !numericType(u.Type(), false) {
								return false
							}
							fields[pathKey(path)] = path
						} else if !input(u.X) && !localAddress(u.X) {
							return false
						}
					} else if u.Op != token.SUB && u.Op != token.XOR && u.Op != token.NOT {
						return false
					}
				case *ssa.Store:
					if path, ok := paths[u.Addr]; ok {
						if len(path) == 0 || !numericType(u.Val.Type(), false) {
							return false
						}
						fields[pathKey(path)] = path
					} else if !localAddress(u.Addr) {
						return false
					}
				case *ssa.Call:
					callee := u.Call.StaticCallee()
					if callee == nil || len(u.Call.Args) != len(callee.Params) || len(u.Call.Args) == 0 {
						return false
					}
					arg := u.Call.Args[0]
					if empty {
						if !types.Identical(arg.Type(), recvType) {
							return false
						}
						if arg != recv {
							if _, ok := arg.(*ssa.Const); !ok {
								return false
							}
						}
					} else if arg != recv {
						return false
					}
					for _, a := range u.Call.Args[1:] {
						if _, ok := a.Type().Underlying().(*types.Pointer); ok && !input(a) && !localAddress(a) {
							return false
						}
					}
					if !check(callee) {
						return false
					}
				case *ssa.BinOp:
					if !numericType(u.X.Type(), false) {
						return false
					}
				case *ssa.Convert:
					if !numericType(u.X.Type(), false) || !numericType(u.Type(), false) {
						return false
					}
				case *ssa.ChangeType:
					if !numericType(u.X.Type(), false) || !numericType(u.Type(), false) {
						return false
					}
				case *ssa.Phi:
					if !numericType(u.Type(), false) {
						return false
					}
				case *ssa.Extract, *ssa.If, *ssa.Jump, *ssa.Return, *ssa.DebugRef:
				default:
					return false
				}
				// receiver addresses may not escape into results, locals or comparisons.
				for _, op := range instr.Operands(nil) {
					if op == nil {
						continue
					}
					if _, ok := paths[*op]; !ok {
						continue
					}
					switch u := instr.(type) {
					case *ssa.FieldAddr:
						if u.X != *op {
							return false
						}
					case *ssa.UnOp:
						if u.Op != token.MUL {
							return false
						}
					case *ssa.Store:
						if u.Addr != *op || u.Val == *op {
							return false
						}
					case *ssa.Call:
						if *op != recv {
							return false
						}
					case *ssa.DebugRef:
					default:
						return false
					}
				}
			}
		}
		return true
	}
	if !check(root) {
		return nil
	}
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		plan.receiver.Fields = append(plan.receiver.Fields, fields[k])
		plan.receiver.Additive = append(plan.receiver.Additive, receiverAdditive(k, addresses))
		updates, affine := receiverTransformUpdates(k, addresses, false)
		bitUpdates, bitwise := receiverTransformUpdates(k, addresses, true)
		plan.receiver.Bitwise = append(plan.receiver.Bitwise, bitwise)
		if bitwise {
			updates = bitUpdates
		}
		plan.receiver.Affine = append(plan.receiver.Affine, affine)
		if affine || bitwise {
			for store, update := range updates {
				update.Field = len(plan.receiver.Fields) - 1
				plan.Updates[store] = update
			}
		}
	}
	return plan
}
func pathKey(path []int) string {
	b := make([]byte, 0, len(path)*8)
	for _, n := range path {
		for i := 0; i < 8; i++ {
			b = append(b, byte(uint64(n)>>uint(i*8)))
		}
	}
	return string(b)
}

// receiverAdditive checks every use of a leaf across the entire closed graph.
// Address identity is structural: separate FieldAddr instructions can name one leaf.
func receiverAdditive(key string, functions map[*ssa.Function]map[ssa.Value][]int) bool {
	updates, ok := receiverTransformUpdates(key, functions, false)
	if !ok {
		return false
	}
	for _, update := range updates {
		if update.Op != token.ADD && update.Op != token.SUB || update.Operand == nil {
			return false
		}
	}
	return true
}

// receiverTransformUpdates proves isolated updates in one closed transform family.
func receiverTransformUpdates(key string, functions map[*ssa.Function]map[ssa.Value][]int, bitwise bool) (map[*ssa.Store]receiverUpdate, bool) {
	updates := map[*ssa.Store]receiverUpdate{}
	written := false
	for fn, paths := range functions {
		matched := func(v ssa.Value) bool {
			path, ok := paths[v]
			return ok && pathKey(path) == key
		}
		stores := map[*ssa.Store]bool{}
		for _, block := range fn.Blocks {
			for _, instr := range block.Instrs {
				load, ok := instr.(*ssa.UnOp)
				if !ok || load.Op != token.MUL || !matched(load.X) {
					continue
				}
				if !numericType(load.Type(), true) {
					return nil, false
				}
				use := receiverSingleUse(load)
				var op token.Token
				var other ssa.Value
				var operation ssa.Value
				switch u := use.(type) {
				case *ssa.BinOp:
					op, operation = u.Op, u
					commutative := op == token.ADD || op == token.MUL
					if bitwise {
						if op != token.AND && op != token.OR && op != token.XOR && op != token.AND_NOT {
							return nil, false
						}
						commutative = op != token.AND_NOT
					} else if op != token.ADD && op != token.SUB && op != token.MUL && op != token.SHL {
						return nil, false
					}
					if u.X != load && (!commutative || u.Y != load) {
						return nil, false
					}
					other = u.X
					if other == load {
						other = u.Y
					}
					if read, ok := other.(*ssa.UnOp); ok && read.Op == token.MUL && matched(read.X) {
						return nil, false
					}
				case *ssa.UnOp:
					op, operation = u.Op, u
					if u.X != load || (bitwise && op != token.XOR) || (!bitwise && op != token.SUB) {
						return nil, false
					}
				default:
					return nil, false
				}
				store, ok := receiverSingleUse(operation).(*ssa.Store)
				if !ok || store.Val != operation || !matched(store.Addr) || store.Block() != block {
					return nil, false
				}
				active, found := false, false
				for _, between := range block.Instrs {
					if between == load {
						active = true
						continue
					}
					if !active {
						continue
					}
					if between == store {
						found = true
						break
					}
					switch between.(type) {
					case *ssa.Call, *ssa.Store:
						return nil, false
					}
				}
				if !found {
					return nil, false
				}
				stores[store] = true
				updates[store] = receiverUpdate{Op: op, Operand: other}
			}
		}
		for _, block := range fn.Blocks {
			for _, instr := range block.Instrs {
				if store, ok := instr.(*ssa.Store); ok && matched(store.Addr) {
					if !stores[store] {
						return nil, false
					}
					written = true
				}
			}
		}
	}
	return updates, written
}

func receiverSingleUse(v ssa.Value) ssa.Instruction {
	var use ssa.Instruction
	for _, ref := range *v.Referrers() {
		if _, ok := ref.(*ssa.DebugRef); ok {
			continue
		}
		if use != nil {
			return nil
		}
		use = ref
	}
	return use
}
