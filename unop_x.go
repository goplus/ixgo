/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ixgo

import (
	"reflect"

	"github.com/visualfc/xtype"
	"golang.org/x/tools/go/ssa"
)

func makeUnOpNOT(pfn *function, instr *ssa.UnOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		if kx == kindConst {
			v := !vx.(bool)
			return func(fr *frame) {
				fr.setReg(ir, v)
			}
		}
		return func(fr *frame) {
			v := !fr.reg(ix).(bool)
			fr.setReg(ir, v)
		}
	} else {
		if kx == kindConst {
			v := xtype.Not(vx)
			return func(fr *frame) {
				fr.setReg(ir, v)
			}
		}
		return func(fr *frame) {
			v := xtype.Not(fr.reg(ix))
			fr.setReg(ir, v)
		}
	}
}

func makeUnOpMUL(pfn *function, instr *ssa.UnOp) func(fr *frame) {
	// Range over *[N]T with only an index (or no variables) does not evaluate
	// *p when len is constant. x/tools still emits an unused deref as the
	// range operand; skip only that trailing load, unless the operand itself
	// has a call or receive that must panic. The skipped instruction never
	// writes a register, so debug inspection of the elided *p sees nil.
	if skipUnusedArrayDeref(instr) && !arrayPointerOperandHasEffectAfter(instr.X, instr.Pos(), nil) {
		return nil
	}
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	if kx == kindGlobal {
		v := reflect.ValueOf(vx)
		return func(fr *frame) {
			elem := v.Elem()
			if !elem.IsValid() {
				panic(fr.runtimeError(instr, "invalid memory address or nil pointer dereference"))
			}
			fr.setReg(ir, elem.Interface())
		}
	}
	return func(fr *frame) {
		elem := reflect.ValueOf(fr.reg(ix)).Elem()
		if !elem.IsValid() {
			panic(fr.runtimeError(instr, "invalid memory address or nil pointer dereference"))
		}
		fr.setReg(ir, elem.Interface())
	}
}

func makeUnOpARROW(pfn *function, instr *ssa.UnOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	typ := pfn.Interp.preToType(instr.X.Type()).Elem()
	if kx == kindGlobal {
		x := reflect.ValueOf(vx)
		if instr.CommaOk {
			return func(fr *frame) {
				v, ok := x.Recv()
				if !ok {
					v = reflect.New(typ).Elem()
				}
				fr.setReg(ir, tuple{v.Interface(), ok})
			}
		}
		return func(fr *frame) {
			v, ok := x.Recv()
			if !ok {
				v = reflect.New(typ).Elem()
			}
			fr.setReg(ir, v.Interface())
		}
	}
	if instr.CommaOk {
		return func(fr *frame) {
			x := reflect.ValueOf(fr.reg(ix))
			v, ok := x.Recv()
			if !ok {
				v = reflect.New(typ).Elem()
			}
			fr.setReg(ir, tuple{v.Interface(), ok})
		}
	}
	return func(fr *frame) {
		x := reflect.ValueOf(fr.reg(ix))
		v, ok := x.Recv()
		if !ok {
			v = reflect.New(typ).Elem()
		}
		fr.setReg(ir, v.Interface())
	}
}
