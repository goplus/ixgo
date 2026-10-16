/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package ixgo

import (
	"encoding/binary"
	"reflect"

	"github.com/goplus/ixgo/internal/memo"
	"golang.org/x/tools/go/ssa"
)

// Each active self-recursion tree owns its cache. Results contain only basic
// scalar values, so replay cannot expose aliases or mutable cached objects.
type interpreterMemo struct {
	values map[uint64][]value
}

// Wrap the ordinary interpreter instructions, including its Return handlers.
// This path never executes typed runner instructions.
func installInterpreterMemo(p *function) {
	p.interpreterMemo = true
	first := p.Instrs[0]
	p.Instrs[0] = func(fr *frame) {
		fr.memo = nil
		if !p.interpreterMemo {
			first(fr)
			return
		}
		if fr.caller != nil && fr.caller.pfn == p {
			fr.memo = fr.caller.memo
		}
		if fr.memo == nil {
			fr.memo = &interpreterMemo{values: make(map[uint64][]value)}
		}
		key := memo.ScalarKey(fr.stack[p.nres])
		if result, ok := fr.memo.values[key]; ok {
			copy(fr.stack[:p.nres], result)
			fr.ipc = -1
			fr.memo = nil
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
			if fr.ipc == -1 && fr.memo != nil {
				key := memo.ScalarKey(fr.stack[p.nres])
				fr.memo.values[key] = append([]value(nil), fr.stack[:p.nres]...)
				fr.memo = nil
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

func installInterpreterAccumulator(p *function) {
	plan := memo.AnalyzeAccumulator(p.Fn)
	if plan == nil {
		return
	}
	output := plan.Output
	p.interpreterAccumulator = true
	first := p.Instrs[0]
	p.Instrs[0] = func(fr *frame) {
		fr.accMemo = nil
		fr.accActive = false
		if !p.interpreterAccumulator {
			first(fr)
			return
		}
		if fr.caller != nil {
			fr.accMemo = fr.caller.accMemo
		}
		if fr.accMemo == nil {
			fr.accMemo = &interpreterAccumulatorMemo{values: make(map[interpreterAccumulatorKey]uint64)}
		}
		cache := fr.accMemo
		if !cache.disabled {
			out := reflect.ValueOf(fr.stack[output])
			valid := !out.IsNil()
			var keys []byte
			for index := 0; index < p.narg; index++ {
				if index == output {
					continue
				}
				input := reflect.ValueOf(fr.stack[index])
				if input.Kind() == reflect.Pointer {
					if input.IsNil() {
						valid = false
						continue
					}
					if !out.IsNil() {
						addr, other := out.Pointer(), input.Pointer()
						if (addr >= other && addr-other < input.Type().Elem().Size()) || (other >= addr && other-addr < out.Type().Elem().Size()) {
							cache.disabled = true
						}
					}
					input = input.Elem()
				}
				keys = binary.LittleEndian.AppendUint64(keys, memo.IntegerBits(input))
			}
			// Nil arguments retain the original load and panic order.
			if valid && !cache.disabled {
				key := interpreterAccumulatorKey{fn: p.Fn, inputs: string(keys)}
				if delta, ok := cache.values[key]; ok {
					memo.SetIntegerBits(out.Elem(), memo.IntegerBits(out.Elem())+delta)
					fr.ipc = -1
					fr.accMemo = nil
					return
				}
				fr.accKey, fr.accStart, fr.accActive = key, memo.IntegerBits(out.Elem()), true
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
			if fr.ipc == -1 && fr.accMemo != nil {
				if fr.accActive && !fr.accMemo.disabled {
					end := memo.IntegerBits(reflect.ValueOf(fr.stack[output]).Elem())
					fr.accMemo.values[fr.accKey] = end - fr.accStart
				}
				fr.accMemo = nil
				fr.accActive = false
			}
		}
	}
}
