/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package ixgo

import (
	"encoding/binary"
	"go/token"
	"reflect"
	"unsafe"

	"github.com/goplus/ixgo/internal/memo"
	"golang.org/x/tools/go/ssa"
)

const interpreterMemoCacheLimit = 4096

// frameExtension holds lazily allocated optional state for a frame.
type frameExtension struct {
	memo        *interpreterMemo
	accumulator *interpreterAccumulatorState
	receiver    *interpreterReceiverState
	numberSlots []uint64
}

type interpreterAccumulatorState struct {
	memo   *interpreterAccumulatorMemo
	key    interpreterAccumulatorKey
	start  uint64
	active bool
}

type interpreterReceiverState struct {
	memo   *interpreterReceiverMemo
	key    interpreterAccumulatorKey
	fields []reflect.Value
	starts []uint64
	affine []receiverTransform
}

func (ext *frameExtension) reset() {
	ext.memo = nil
	if ext.accumulator != nil {
		ext.accumulator.reset()
	}
	if ext.receiver != nil {
		ext.receiver.reset()
	}
}

func (state *interpreterReceiverState) reset() {
	state.memo = nil
	state.key = interpreterAccumulatorKey{}
	clear(state.fields)
	state.fields = state.fields[:0]
	state.starts = state.starts[:0]
	state.affine = state.affine[:0]
}

func (state *interpreterAccumulatorState) reset() {
	state.memo = nil
	state.key = interpreterAccumulatorKey{}
	state.start = 0
	state.active = false
}

func (ext *frameExtension) resetReceiver() *interpreterReceiverState {
	if ext.receiver == nil {
		ext.receiver = &interpreterReceiverState{}
	} else {
		ext.receiver.reset()
	}
	return ext.receiver
}

func (ext *frameExtension) resetAccumulator() *interpreterAccumulatorState {
	if ext.accumulator == nil {
		ext.accumulator = &interpreterAccumulatorState{}
	} else {
		ext.accumulator.reset()
	}
	return ext.accumulator
}

func (ext *frameExtension) resetScalar() {
	ext.memo = nil
}

// Each active self-recursion tree owns its cache. Results contain only basic
// scalar values, so replay cannot expose aliases or mutable cached objects.
type interpreterMemo struct {
	values map[uint64][]value
}

// Wrap the ordinary interpreter instructions, including its Return handlers.
// This path never executes typed runner instructions.
func installInterpreterMemo(p *function) {
	p.extFlags |= extMemoScalar
	first := p.Instrs[0]
	p.Instrs[0] = func(fr *frame) {
		if fr.ext == nil {
			fr.ext = &frameExtension{}
		} else {
			fr.ext.resetScalar()
		}
		if fr.caller != nil && fr.caller.pfn == p && fr.caller.ext != nil {
			fr.ext.memo = fr.caller.ext.memo
		}
		if fr.ext.memo == nil {
			fr.ext.memo = &interpreterMemo{values: make(map[uint64][]value)}
		}
		key := memo.ScalarKey(fr.stack[p.nres]) // first scalar argument; later params must be forwarded unchanged
		if result, ok := fr.ext.memo.values[key]; ok {
			copy(fr.stack[:p.nres], result)
			fr.ipc = -1
			fr.ext.memo = nil
			return
		}
		first(fr)
	}
	for i, instr := range p.ssaInstrs {
		if _, ok := instr.(*ssa.Return); !ok {
			continue
		}
		ret := p.Instrs[i]
		p.Instrs[i] = func(fr *frame) {
			ret(fr)
			if fr.ipc == -1 && fr.ext != nil && fr.ext.memo != nil {
				key := memo.ScalarKey(fr.stack[p.nres])
				if len(fr.ext.memo.values) < interpreterMemoCacheLimit {
					fr.ext.memo.values[key] = append([]value(nil), fr.stack[:p.nres]...)
				}
				fr.ext.memo = nil
			}
		}
	}
}

type interpreterAccumulatorKey struct {
	fn     *ssa.Function
	inputs string
}

type interpreterAccumulatorMemo struct {
	values   map[interpreterAccumulatorKey]uint64
	disabled bool
}

func installInterpreterAccumulator(p *function, output int, state bool) {
	p.extFlags |= extMemoAccumulator
	first := p.Instrs[0]
	p.Instrs[0] = func(fr *frame) {
		if fr.ext == nil {
			fr.ext = &frameExtension{}
		}
		accumulator := fr.ext.resetAccumulator()
		if fr.caller != nil && fr.caller.ext != nil && fr.caller.ext.accumulator != nil {
			accumulator.memo = fr.caller.ext.accumulator.memo
		}
		if accumulator.memo == nil {
			accumulator.memo = &interpreterAccumulatorMemo{values: make(map[interpreterAccumulatorKey]uint64)}
		}
		cache := accumulator.memo
		if !cache.disabled {
			out := reflect.ValueOf(fr.stack[output])
			valid := !out.IsNil()
			var keys []byte
			for index := 0; index < p.narg; index++ {
				if index == output && !state {
					continue
				}
				input := reflect.ValueOf(fr.stack[index])
				if input.Kind() == reflect.Pointer {
					if input.IsNil() {
						valid = false
						continue
					}
					if index != output && !out.IsNil() {
						addr, other := out.Pointer(), input.Pointer()
						if (addr >= other && addr-other < input.Type().Elem().Size()) || (other >= addr && other-addr < out.Type().Elem().Size()) {
							cache.disabled = true
						}
					}
					input = input.Elem()
				}
				keys = binary.LittleEndian.AppendUint64(keys, memo.ValueBits(input))
			}
			// Nil arguments retain the original load and panic order.
			if valid && !cache.disabled {
				key := interpreterAccumulatorKey{fn: p.Fn, inputs: string(keys)}
				if delta, ok := cache.values[key]; ok {
					if state {
						memo.SetScalarBits(out.Elem(), delta)
					} else {
						memo.SetIntegerBits(out.Elem(), memo.IntegerBits(out.Elem())+delta)
					}
					fr.ipc = -1
					fr.ext.accumulator.memo = nil
					return
				}
				fr.ext.accumulator.key, fr.ext.accumulator.start, fr.ext.accumulator.active = key, memo.ValueBits(out.Elem()), true
			}
		}
		first(fr)
	}
	for i, instr := range p.ssaInstrs {
		if _, ok := instr.(*ssa.Return); !ok {
			continue
		}
		ret := p.Instrs[i]
		p.Instrs[i] = func(fr *frame) {
			ret(fr)
			if fr.ipc == -1 && fr.ext != nil && fr.ext.accumulator.memo != nil {
				if fr.ext.accumulator.active && !fr.ext.accumulator.memo.disabled {
					end := memo.ValueBits(reflect.ValueOf(fr.stack[output]).Elem())
					if len(fr.ext.accumulator.memo.values) < interpreterMemoCacheLimit {
						if state {
							fr.ext.accumulator.memo.values[fr.ext.accumulator.key] = end
						} else {
							fr.ext.accumulator.memo.values[fr.ext.accumulator.key] = end - fr.ext.accumulator.start
						}
					}
				}
				fr.ext.accumulator.memo = nil
				fr.ext.accumulator.active = false
			}
		}
	}
}

// Receiver caches are confined to compatible closed graphs and one receiver.
// Saturated caches retain existing hits but execute new states without storage.
const interpreterReceiverCacheLimit = 4096

type interpreterReceiverResult struct {
	fields []uint64
	affine []receiverTransform
	values []value
}
type interpreterReceiverMemo struct {
	plan     *MemoReceiverPlan
	receiver any
	values   map[interpreterAccumulatorKey]interpreterReceiverResult
	hits     uint64
	disabled bool
}

// Affine coefficients or bitwise keep/flip masks; replay narrows to field width.
type receiverTransform struct{ a, b uint64 }

func composeReceiverParent(fr *frame, transforms []receiverTransform) {
	parent := fr.caller
	if parent == nil || parent.ext == nil || fr.ext == nil ||
		parent.ext.receiver == nil || fr.ext.receiver == nil ||
		parent.ext.receiver.memo == nil || fr.ext.receiver.memo == nil ||
		fr.ext.receiver.memo.receiver != parent.ext.receiver.memo.receiver {
		return
	}
	for i, field := range fr.ext.receiver.fields {
		if !fr.ext.receiver.memo.plan.Receiver.Affine[i] && !fr.ext.receiver.memo.plan.Receiver.Bitwise[i] {
			continue
		}
		for j, target := range parent.ext.receiver.fields {
			if parent.ext.receiver.memo.plan.Receiver.Additive[j] || field.UnsafeAddr() != target.UnsafeAddr() ||
				fr.ext.receiver.memo.plan.Receiver.Affine[i] != parent.ext.receiver.memo.plan.Receiver.Affine[j] ||
				fr.ext.receiver.memo.plan.Receiver.Bitwise[i] != parent.ext.receiver.memo.plan.Receiver.Bitwise[j] {
				continue
			}
			t, prev := transforms[i], parent.ext.receiver.affine[j]
			if fr.ext.receiver.memo.plan.Receiver.Affine[i] && parent.ext.receiver.memo.plan.Receiver.Affine[j] {
				parent.ext.receiver.affine[j] = receiverTransform{t.a * prev.a, t.a*prev.b + t.b}
			} else if fr.ext.receiver.memo.plan.Receiver.Bitwise[i] && parent.ext.receiver.memo.plan.Receiver.Bitwise[j] {
				parent.ext.receiver.affine[j] = receiverTransform{t.a & prev.a, (prev.b & t.a) ^ t.b}
			}
		}
	}
}

func installInterpreterReceiver(p *function, plan *MemoReceiverPlan) {
	p.extFlags |= extMemoReceiver
	updates := make(map[ssa.Instruction]MemoReceiverUpdate)
	for store, update := range plan.Updates {
		updates[store.Val.(ssa.Instruction)] = update
	}
	for i, instruction := range p.ssaInstrs {
		update, ok := updates[instruction]
		if !ok || plan.Receiver.Additive[update.Field] {
			continue
		}
		var operand register
		if update.Operand != nil {
			operand = p.regIndex(update.Operand)
		}
		execute := p.Instrs[i]
		p.Instrs[i] = func(fr *frame) {
			// Execute first so invalid shift counts retain the original panic.
			execute(fr)
			if fr.ext != nil && fr.ext.receiver.memo != nil && !fr.ext.receiver.memo.disabled {
				var bits uint64
				if update.Operand != nil {
					bits = memo.IntegerBits(reflect.ValueOf(fr.reg(operand)))
				}
				t := &fr.ext.receiver.affine[update.Field]
				switch update.Op {
				case token.ADD:
					t.b += bits
				case token.SUB:
					if update.Operand == nil {
						t.a, t.b = -t.a, -t.b
					} else {
						t.b -= bits
					}
				case token.SHL:
					// Counts are not narrowed to the destination integer width.
					factor := uint64(1) << bits
					t.a *= factor
					t.b *= factor
				case token.AND:
					t.a &= bits
					t.b &= bits
				case token.AND_NOT:
					t.a &^= bits
					t.b &^= bits
				case token.OR:
					t.a &^= bits
					t.b |= bits
				case token.XOR:
					if update.Operand == nil {
						bits = ^uint64(0)
					}
					t.b ^= bits
				case token.MUL:
					t.a *= bits
					t.b *= bits
				}
			}
		}
	}
	first := p.Instrs[0]
	p.Instrs[0] = func(fr *frame) {
		if fr.ext == nil {
			fr.ext = &frameExtension{}
		}
		receiverState := fr.ext.resetReceiver()
		receiver := fr.stack[p.nres]
		if fr.caller != nil && fr.caller.ext != nil && fr.caller.ext.receiver != nil {
			parent := fr.caller.ext.receiver.memo
			if parent != nil && compatibleReceiverPlan(plan, parent.plan) && receiver == parent.receiver {
				receiverState.memo = parent
			}
		}
		if receiverState.memo == nil {
			receiverState.memo = &interpreterReceiverMemo{plan: plan, receiver: receiver, values: make(map[interpreterAccumulatorKey]interpreterReceiverResult)}
		}
		cache := receiverState.memo
		var keys []byte
		valid := !cache.disabled
		recv := reflect.ValueOf(receiver)
		if recv.Kind() == reflect.Pointer {
			if recv.IsNil() {
				valid = false
			} else {
				recv = recv.Elem()
			}
		}
		if valid {
			for index, path := range plan.Receiver.Fields {
				field := recv.FieldByIndex(path)
				// Interpreted private fields have the same addressability as exported ones.
				if !field.CanSet() {
					field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
				}
				fr.ext.receiver.fields = append(fr.ext.receiver.fields, field)
				bits := memo.ValueBits(field)
				fr.ext.receiver.starts = append(fr.ext.receiver.starts, bits)
				identity := uint64(1)
				if plan.Receiver.Bitwise[index] {
					identity = ^uint64(0)
				}
				fr.ext.receiver.affine = append(fr.ext.receiver.affine, receiverTransform{a: identity})
				if !plan.Receiver.Affine[index] && !plan.Receiver.Bitwise[index] {
					keys = binary.LittleEndian.AppendUint64(keys, bits)
				}
			}
		}
		for index := 1; index < p.narg; index++ {
			arg := reflect.ValueOf(fr.stack[p.nres+index])
			if arg.Kind() == reflect.Pointer {
				if arg.IsNil() {
					valid = false
					continue
				}
				// Read-only inputs may alias each other, but never receiver storage.
				r := reflect.ValueOf(receiver)
				if r.Kind() == reflect.Pointer && !r.IsNil() {
					a, b := arg.Pointer(), r.Pointer()
					if (a >= b && a-b < r.Type().Elem().Size()) || (b >= a && b-a < arg.Type().Elem().Size()) {
						cache.disabled = true
						valid = false
					}
				}
				arg = arg.Elem()
			}
			keys = binary.LittleEndian.AppendUint64(keys, memo.ValueBits(arg))
		}
		if valid {
			fr.ext.receiver.key = interpreterAccumulatorKey{fn: p.Fn, inputs: string(keys)}
			if result, ok := cache.values[fr.ext.receiver.key]; ok {
				for i, field := range fr.ext.receiver.fields {
					if plan.Receiver.Affine[i] {
						t := result.affine[i]
						memo.SetIntegerBits(field, t.a*memo.IntegerBits(field)+t.b)
					} else if plan.Receiver.Bitwise[i] {
						t := result.affine[i]
						memo.SetIntegerBits(field, (memo.IntegerBits(field)&t.a)^t.b)
					} else {
						memo.SetScalarBits(field, result.fields[i])
					}
				}
				composeReceiverParent(fr, result.affine)
				copy(fr.stack[:p.nres], result.values)
				cache.hits++
				fr.ipc = -1
				return
			}
		} else {
			fr.ext.receiver.fields = nil
			fr.ext.receiver.starts = nil
			fr.ext.receiver.memo = nil
		}
		first(fr)
	}
	for i, instr := range p.ssaInstrs {
		if _, ok := instr.(*ssa.Return); !ok {
			continue
		}
		ret := p.Instrs[i]
		p.Instrs[i] = func(fr *frame) {
			ret(fr)
			if fr.ipc == -1 && fr.ext != nil && fr.ext.receiver.memo != nil && !fr.ext.receiver.memo.disabled {
				result := interpreterReceiverResult{values: append([]value(nil), fr.stack[:p.nres]...), affine: append([]receiverTransform(nil), fr.ext.receiver.affine...)}
				for index, field := range fr.ext.receiver.fields {
					bits := memo.ValueBits(field)
					if plan.Receiver.Additive[index] {
						bits -= fr.ext.receiver.starts[index]
						result.affine[index] = receiverTransform{a: 1, b: bits}
					}
					result.fields = append(result.fields, bits)
				}
				composeReceiverParent(fr, result.affine)
				if len(fr.ext.receiver.memo.values) < interpreterReceiverCacheLimit {
					fr.ext.receiver.memo.values[fr.ext.receiver.key] = result
				}
			}
		}
	}
}

// MemoInstaller exposes the memoization installers available to an optional optimizer.
type MemoInstaller interface {
	InstallScalarMemo()
	InstallReceiverMemo(*MemoReceiverPlan)
	InstallAccumulatorMemo(output int, state bool)
}

// MemoPlugin is called for each eligible function when an optional memo package is registered.
type MemoPlugin func(MemoInstaller, *ssa.Function)

var memoPlugin MemoPlugin

// RegisterMemoPlugin enables an optional optimizer package.
// The plugin API is experimental and may change without notice.
func RegisterMemoPlugin(plugin MemoPlugin) { memoPlugin = plugin }

// MemoReceiverPlan describes the receiver state needed by the interpreter.
type MemoReceiverPlan struct {
	Receiver  *MemoReceiverState
	Functions map[*ssa.Function]bool
	Updates   map[*ssa.Store]MemoReceiverUpdate
}

// MemoReceiverState describes the cached numerical leaves of a receiver.
type MemoReceiverState struct {
	Fields   [][]int
	Additive []bool
	Affine   []bool
	Bitwise  []bool
}

// MemoReceiverUpdate records an isolated store that updates one cached leaf.
type MemoReceiverUpdate struct {
	Field   int
	Op      token.Token
	Operand ssa.Value
}

func compatibleReceiverPlan(p, q *MemoReceiverPlan) bool {
	if p == nil || q == nil || len(p.Functions) != len(q.Functions) || len(p.Receiver.Fields) != len(q.Receiver.Fields) {
		return false
	}
	for fn := range p.Functions {
		if !q.Functions[fn] {
			return false
		}
	}
	for i, path := range p.Receiver.Fields {
		if p.Receiver.Additive[i] != q.Receiver.Additive[i] || p.Receiver.Affine[i] != q.Receiver.Affine[i] || p.Receiver.Bitwise[i] != q.Receiver.Bitwise[i] {
			return false
		}
		other := q.Receiver.Fields[i]
		if len(path) != len(other) {
			return false
		}
		for j, n := range path {
			if n != other[j] {
				return false
			}
		}
	}
	return true
}

// InstallScalarMemo installs scalar result memoization.
func (p *function) InstallScalarMemo() { installInterpreterMemo(p) }

// InstallReceiverMemo installs receiver-state memoization for an analyzed plan.
func (p *function) InstallReceiverMemo(plan *MemoReceiverPlan) {
	installInterpreterReceiver(p, plan)
}

// InstallAccumulatorMemo installs accumulator or full-state memoization.
func (p *function) InstallAccumulatorMemo(output int, state bool) {
	installInterpreterAccumulator(p, output, state)
}

func applyMemoPlugin(p *function, fn *ssa.Function) {
	if memoPlugin != nil {
		memoPlugin(p, fn)
	}
}
