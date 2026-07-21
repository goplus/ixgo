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
	"fmt"
	"go/token"
	"go/types"
	"io"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/goplus/ixgo/load"
	"github.com/goplus/reflectx"
	"github.com/visualfc/xtype"
	"golang.org/x/tools/go/ssa"
)

/*
type Op int

const (
	OpInvalid Op = iota
	// Value-defining instructions
	OpAlloc
	OpPhi
	OpCall
	OpBinOp
	OpUnOp
	OpChangeType
	OpConvert
	OpChangeInterface
	OpSliceToArrayPointer
	OpMakeInterface
	OpMakeClosure
	OpMakeMap
	OpMakeChan
	OpMakeSlice
	OpSlice
	OpFieldAddr
	OpField
	OpIndexAddr
	OpIndex
	OpLookup
	OpSelect
	OpRange
	OpNext
	OpTypeAssert
	OpExtract
	// Instructions executed for effect
	OpJump
	OpIf
	OpReturn
	OpRunDefers
	OpPanic
	OpGo
	OpDefer
	OpSend
	OpStore
	OpMapUpdate
	OpDebugRef
)
*/

type kind int

func (k kind) isStatic() bool {
	return k == kindConst || k == kindGlobal || k == kindFunction
}

const (
	kindInvalid kind = iota
	kindConst
	kindGlobal
	kindFunction
)

type Value = value
type Tuple = tuple

type register int
type value = interface{}
type tuple []value

type closure struct {
	pfn *function
	env []value
}

type function struct {
	Interp     *Interp
	Fn         *ssa.Function                // ssa function
	Main       *ssa.BasicBlock              // Fn.Blocks[0]
	pool       *sync.Pool                   // create frame pool
	makeInstr  ssa.Instruction              // make instr check
	index      map[ssa.Value]uint32         // stack value index 32bit: kind(2) reflect.Kind(6) index(24)
	instrIndex map[ssa.Instruction][]uint32 // instr -> index
	Instrs     []func(fr *frame)            // main instrs
	Recover    []func(fr *frame)            // recover instrs
	Blocks     []int                        // block offset
	stack      []value                      // results args envs datas
	ssaInstrs  []ssa.Instruction            // org ssa instr
	base       int                          // base of interp
	nres       int                          // results count
	narg       int                          // arguments count
	nenv       int                          // closure free vars count
	used       int32                        // function used count
	cached     int32                        // enable cached by pool
}

func (p *function) UnsafeRelease() {
	p.Interp = nil
	p.Fn = nil
	p.pool = nil
	p.index = nil
	p.instrIndex = nil
	p.Instrs = nil
	p.Recover = nil
	p.Blocks = nil
	p.stack = nil
	p.ssaInstrs = nil
	p.Main = nil
}

func (p *function) initPool() {
	p.pool = &sync.Pool{}
	p.pool.New = func() interface{} {
		fr := &frame{interp: p.Interp, pfn: p, block: p.Main}
		fr.stack = append([]value{}, p.stack...)
		return fr
	}
}

func (p *function) allocFrame(caller *frame) *frame {
	var fr *frame
	if atomic.LoadInt32(&p.cached) == 1 {
		fr = p.pool.Get().(*frame)
		fr.block = p.Main
		fr._defer = nil
		fr._panic = nil
		fr.ipc = 0
		fr.pred = 0
	} else {
		if atomic.AddInt32(&p.used, 1) > int32(p.Interp.ctx.callForPool) {
			atomic.StoreInt32(&p.cached, 1)
		}
		fr = &frame{interp: p.Interp, pfn: p, block: p.Main}
		fr.stack = append([]value{}, p.stack...)
	}
	fr.caller = caller
	fr.deferid = caller.deferid
	caller.callee = fr
	return fr
}

func (p *function) deleteFrame(caller *frame, fr *frame) {
	if atomic.LoadInt32(&p.cached) == 1 {
		p.pool.Put(fr)
	} else {
		caller.callee = nil
	}
	fr = nil
}

func (p *function) InstrForPC(pc int) ssa.Instruction {
	if pc >= 0 && pc < len(p.ssaInstrs) {
		return p.ssaInstrs[pc]
	}
	return nil
}

func (p *function) PosForPC(pc int) token.Pos {
	if instr := p.InstrForPC(pc); instr != nil {
		if _, ok := instr.(*ssa.RunDefers); ok {
			return p.Fn.Syntax().End()
		}
		return instr.Pos()
	}
	return token.NoPos
}

func (p *function) PositionForPC(pc int) token.Position {
	pos := p.PosForPC(pc)
	return p.Fn.Prog.Fset.Position(pos)
}

func (p *function) regIndex3(v ssa.Value) (register, kind, value) {
	instr := p.regInstr(v)
	index := int(instr & 0xffffff)
	return register(index), kind(instr >> 30), p.stack[index]
}

func (p *function) regIndex(v ssa.Value) register {
	instr := p.regInstr(v)
	return register(instr & 0xffffff)
}

func (p *function) regInstr(v ssa.Value) uint32 {
	if i, ok := p.index[v]; ok {
		return i
	}
	var vs interface{}
	var vk kind
	switch v := v.(type) {
	case *ssa.Const:
		vs = constToValue(p.Interp, v)
		vk = kindConst
	case *ssa.Global:
		vs, _ = globalToValue(p.Interp, v)
		vk = kindGlobal
	case *ssa.Function:
		vk = kindFunction
		if v.Blocks != nil {
			typ := p.Interp.preToType(v.Type())
			pfn := p.Interp.loadFunction(v)
			vs = pfn.makeFunction(typ, nil).Interface()
		} else {
			ext, ok := findExternFunc(p.Interp, v)
			if !ok {
				if v.Name() != "init" {
					panic(fmt.Errorf("no code for function: %v", v))
				}
			} else {
				vs = ext.Interface()
			}
		}
	}
	var kind reflect.Kind
	if v != nil {
		kind = toKind(v.Type())
	}
	i := uint32(len(p.stack) | int(vk<<30) | int(kind<<24))
	p.stack = append(p.stack, vs)
	p.index[v] = i
	p.instrIndex[p.makeInstr] = append(p.instrIndex[p.makeInstr], i)
	return i
}

func findExternValue(interp *Interp, name string) (ext reflect.Value, ok bool) {
	// check override value
	ext, ok = interp.ctx.override[name]
	if !ok {
		// check extern value
		ext, ok = externValues[name]
	}
	return
}

func findExternLinkFunc(interp *Interp, link *load.Linkname) (ext reflect.Value, ok bool) {
	fullName := link.PkgPath + "." + link.Name
	// check override value
	ext, ok = interp.ctx.override[fullName]
	if ok {
		return
	}
	// check extern value
	ext, ok = externValues[fullName]
	if ok {
		return
	}
	// check install pkg
	if pkg, found := interp.installed(link.PkgPath); found {
		if recv := link.Recv; recv != "" {
			var star bool
			if recv[0] == '*' {
				star = true
				recv = recv[1:]
			}
			if typ, ok := pkg.NamedTypes[recv]; ok {
				if star {
					typ = reflect.PtrTo(typ)
				}
				if m, ok := reflectx.MethodByName(typ, link.Method); ok {
					return m.Func, true
				}
			}
			return
		}
		ext, ok = pkg.Funcs[link.Name]
	}
	return
}

func findExternVar(interp *Interp, pkgPath string, name string) (ext reflect.Value, ok bool) {
	fullName := pkgPath + "." + name
	// check override value
	ext, ok = interp.ctx.override[fullName]
	if ok {
		return
	}
	// check extern value
	ext, ok = externValues[fullName]
	if ok {
		return
	}
	// check install pkg
	if pkg, found := interp.installed(pkgPath); found {
		ext, ok = pkg.Vars[name]
	}
	return
}

func checkFuncCompatible(t1, t2 reflect.Type) bool {
	i1 := t1.NumIn()
	i2 := t2.NumIn()
	if i1 != i2 {
		return false
	}
	for i := 0; i < i1; i++ {
		if t1.In(i).Size() != t2.In(i).Size() {
			return false
		}
	}
	return true
}

func findExternFunc(interp *Interp, fn *ssa.Function) (ext reflect.Value, ok bool) {
	fnName := fn.String()
	ext, ok = findExternValue(interp, fnName)
	if ok {
		typ := interp.preToType(fn.Type())
		ftyp := ext.Type()
		if typ != ftyp && checkFuncCompatible(typ, ftyp) {
			ext = xtype.ConvertFuncValue(xtype.TypeOfType(typ), ext)
		}
		return
	}
	if fn.Pkg != nil {
		if recv := fn.Signature.Recv(); recv == nil {
			if pkg, found := LookupPackage(fn.Pkg.Pkg.Path()); found {
				ext, ok = pkg.Funcs[fn.Name()]
			}
		} else if typ, found := interp.ctx.Loader.LookupReflect(recv.Type()); found {
			if m, found := reflectx.MethodByName(typ, fn.Name()); found {
				ext, ok = m.Func, true
			}
		}
	}
	return
}

func makeInstr(interp *Interp, pfn *function, instr ssa.Instruction) func(fr *frame) {
	switch instr := instr.(type) {
	case *ssa.Alloc:
		if instr.Heap {
			typ := interp.preToType(instr.Type())
			ir := pfn.regIndex(instr)
			t := xtype.TypeOfType(typ.Elem())
			pt := xtype.TypeOfType(typ)
			return func(fr *frame) {
				fr.setReg(ir, xtype.New(t, pt))
			}
		}
		typ := interp.preToType(instr.Type())
		ir := pfn.regIndex(instr)
		t := xtype.TypeOfType(typ.Elem())
		pt := xtype.TypeOfType(typ)
		// ptr := xtype.NewPointer(t)
		elem := reflect.New(typ.Elem()).Elem()
		return func(fr *frame) {
			if v := fr.reg(ir); v != nil {
				reflect.ValueOf(v).Elem().Set(elem)
			} else {
				fr.setReg(ir, xtype.New(t, pt))
			}
		}
	case *ssa.Phi:
		ir := pfn.regIndex(instr)
		ie := make([]register, len(instr.Edges))
		for i, v := range instr.Edges {
			ie[i] = pfn.regIndex(v)
		}
		return func(fr *frame) {
			for i, pred := range instr.Block().Preds {
				if fr.pred == pred.Index {
					fr.setReg(ir, fr.reg(ie[i]))
					break
				}
			}
		}
	case *ssa.Call:
		return makeCallInstr(pfn, interp, instr, &instr.Call)
	case *ssa.BinOp:
		switch instr.Op {
		case token.ADD:
			return makeBinOpADD(pfn, instr)
		case token.SUB:
			return makeBinOpSUB(pfn, instr)
		case token.MUL:
			return makeBinOpMUL(pfn, instr)
		case token.QUO:
			return makeBinOpQUO(pfn, instr)
		case token.REM:
			return makeBinOpREM(pfn, instr)
		case token.AND:
			return makeBinOpAND(pfn, instr)
		case token.OR:
			return makeBinOpOR(pfn, instr)
		case token.XOR:
			return makeBinOpXOR(pfn, instr)
		case token.AND_NOT:
			return makeBinOpANDNOT(pfn, instr)
		case token.LSS:
			return makeBinOpLSS(pfn, instr)
		case token.LEQ:
			return makeBinOpLEQ(pfn, instr)
		case token.GTR:
			return makeBinOpGTR(pfn, instr)
		case token.GEQ:
			return makeBinOpGEQ(pfn, instr)
		case token.EQL:
			return makeBinOpEQL(pfn, instr)
		case token.NEQ:
			return makeBinOpNEQ(pfn, instr)
		case token.SHL:
			return makeBinOpSHL(pfn, instr)
		case token.SHR:
			return makeBinOpSHR(pfn, instr)
		default:
			panic("unreachable")
		}
	case *ssa.UnOp:
		switch instr.Op {
		case token.NOT:
			return makeUnOpNOT(pfn, instr)
		case token.SUB:
			return makeUnOpSUB(pfn, instr)
		case token.XOR:
			return makeUnOpXOR(pfn, instr)
		case token.ARROW:
			return makeUnOpARROW(pfn, instr)
		case token.MUL:
			return makeUnOpMUL(pfn, instr)
		default:
			panic("unreachable")
		}
	case *ssa.ChangeInterface:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame) {
			fr.setReg(ir, fr.reg(ix))
		}
	case *ssa.ChangeType:
		return makeTypeChangeInstr(pfn, instr)
	case *ssa.Convert:
		return makeConvertInstr(pfn, interp, instr)
	case *ssa.MakeInterface:
		typ := interp.preToType(instr.Type())
		ir := pfn.regIndex(instr)
		ix, kx, vx := pfn.regIndex3(instr.X)
		if kx.isStatic() {
			if typ == tyEmptyInterface {
				return func(fr *frame) {
					fr.setReg(ir, vx)
				}
			}
			v := reflect.New(typ).Elem()
			if vx != nil {
				SetValue(v, reflect.ValueOf(vx))
			}
			vx = v.Interface()
			return func(fr *frame) {
				fr.setReg(ir, vx)
			}
		}
		if typ == tyEmptyInterface {
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix))
			}
		}
		return func(fr *frame) {
			v := reflect.New(typ).Elem()
			if x := fr.reg(ix); x != nil {
				vx := reflect.ValueOf(x)
				SetValue(v, vx)
			}
			fr.setReg(ir, v.Interface())
		}
	case *ssa.MakeClosure:
		fn := instr.Fn.(*ssa.Function)
		typ := interp.preToType(fn.Type())
		ir := pfn.regIndex(instr)
		ib := make([]register, len(instr.Bindings))
		for i, v := range instr.Bindings {
			ib[i] = pfn.regIndex(v)
		}
		pfn := interp.loadFunction(fn)
		return func(fr *frame) {
			var bindings []value
			for i := range instr.Bindings {
				bindings = append(bindings, fr.reg(ib[i]))
			}
			v := pfn.makeFunction(typ, bindings)
			fr.setReg(ir, v.Interface())
		}
	case *ssa.MakeChan:
		typ := interp.preToType(instr.Type())
		ir := pfn.regIndex(instr)
		is := pfn.regIndex(instr.Size)
		if typ.ChanDir() == reflect.BothDir {
			return func(fr *frame) {
				size := fr.reg(is)
				buffer := asInt(size)
				if buffer < 0 {
					panic(fr.runtimeError(instr, "makechan: size out of range"))
				}
				fr.setReg(ir, reflect.MakeChan(typ, buffer).Interface())
			}
		}
		ctyp := reflect.ChanOf(reflect.BothDir, typ.Elem())
		return func(fr *frame) {
			size := fr.reg(is)
			buffer := asInt(size)
			if buffer < 0 {
				panic(fr.runtimeError(instr, "makechan: size out of range"))
			}
			fr.setReg(ir, reflect.MakeChan(ctyp, buffer).Convert(typ).Interface())
		}
	case *ssa.MakeMap:
		typ := instr.Type()
		rtyp := interp.preToType(typ)
		ir := pfn.regIndex(instr)
		if instr.Reserve == nil {
			return func(fr *frame) {
				fr.setReg(ir, reflect.MakeMap(rtyp).Interface())
			}
		}
		iv := pfn.regIndex(instr.Reserve)
		return func(fr *frame) {
			reserve := asInt(fr.reg(iv))
			fr.setReg(ir, reflect.MakeMapWithSize(rtyp, reserve).Interface())
		}
	case *ssa.MakeSlice:
		typ := interp.preToType(instr.Type())
		ir := pfn.regIndex(instr)
		il := pfn.regIndex(instr.Len)
		ic := pfn.regIndex(instr.Cap)
		return func(fr *frame) {
			Len := asInt(fr.reg(il))
			if Len < 0 || Len >= maxMemLen {
				panic(fr.runtimeError(instr, "makeslice: len out of range"))
			}
			Cap := asInt(fr.reg(ic))
			if Cap < 0 || Cap >= maxMemLen {
				panic(fr.runtimeError(instr, "makeslice: cap out of range"))
			}
			fr.setReg(ir, reflect.MakeSlice(typ, Len, Cap).Interface())
		}
	case *ssa.Slice:
		typ := interp.preToType(instr.Type())
		isNamed := typ.Kind() == reflect.Slice && typ != reflect.SliceOf(typ.Elem())
		_, makesliceCheck := instr.X.(*ssa.Alloc)
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		ih := pfn.regIndex(instr.High)
		il := pfn.regIndex(instr.Low)
		im := pfn.regIndex(instr.Max)
		if isNamed {
			return func(fr *frame) {
				fr.setReg(ir, slice(fr, instr, makesliceCheck, ix, ih, il, im).Convert(typ).Interface())
			}
		}
		return func(fr *frame) {
			fr.setReg(ir, slice(fr, instr, makesliceCheck, ix, ih, il, im).Interface())
		}
	case *ssa.FieldAddr:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame) {
			v, err := fieldAddrX(fr.reg(ix), instr.Field)
			if err != nil {
				panic(fr.runtimeError(instr, err.Error()))
			}
			fr.setReg(ir, v)
		}
	case *ssa.Field:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame) {
			v, err := fieldX(fr.reg(ix), instr.Field)
			if err != nil {
				panic(fr.runtimeError(instr, err.Error()))
			}
			fr.setReg(ir, v)
		}
	case *ssa.IndexAddr:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		ii := pfn.regIndex(instr.Index)
		return func(fr *frame) {
			x := fr.reg(ix)
			idx := fr.reg(ii)
			v := reflect.ValueOf(x)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			switch v.Kind() {
			case reflect.Slice:
			case reflect.Array:
			case reflect.Invalid:
				panic(fr.runtimeError(instr, "invalid memory address or nil pointer dereference"))
			default:
				panic(fmt.Sprintf("unexpected x type in IndexAddr: %T", x))
			}
			index := asInt(idx)
			if index < 0 {
				panic(fr.runtimeError(instr, fmt.Sprintf("index out of range [%v]", index)))
			} else if length := v.Len(); index >= length {
				panic(fr.runtimeError(instr, fmt.Sprintf("index out of range [%v] with length %v", index, length)))
			}
			fr.setReg(ir, v.Index(index).Addr().Interface())
		}
	case *ssa.Index:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		ii := pfn.regIndex(instr.Index)
		return func(fr *frame) {
			x := fr.reg(ix)
			idx := fr.reg(ii)
			index := asInt(idx)
			v := reflect.ValueOf(x)
			if index < 0 {
				panic(fr.runtimeError(instr, fmt.Sprintf("index out of range [%v]", index)))
			} else if length := v.Len(); index >= length {
				panic(fr.runtimeError(instr, fmt.Sprintf("index out of range [%v] with length %v", index, length)))
			}
			fr.setReg(ir, v.Index(index).Interface())
		}
	case *ssa.Lookup:
		typ := interp.preToType(instr.X.Type())
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		ii := pfn.regIndex(instr.Index)
		switch typ.Kind() {
		case reflect.String:
			return func(fr *frame) {
				v := fr.reg(ix)
				idx := fr.reg(ii)
				fr.setReg(ir, reflect.ValueOf(v).String()[asInt(idx)])
			}
		case reflect.Map:
			return func(fr *frame) {
				m := fr.reg(ix)
				idx := fr.reg(ii)
				vm := reflect.ValueOf(m)
				v := vm.MapIndex(reflect.ValueOf(idx))
				ok := v.IsValid()
				var rv value
				if ok {
					rv = v.Interface()
				} else {
					rv = reflect.New(typ.Elem()).Elem().Interface()
				}
				if instr.CommaOk {
					fr.setReg(ir, tuple{rv, ok})
				} else {
					fr.setReg(ir, rv)
				}
			}
		default:
			panic("unreachable")
		}
	case *ssa.Select:
		ir := pfn.regIndex(instr)
		ic := make([]register, len(instr.States))
		is := make([]register, len(instr.States))
		for i, state := range instr.States {
			ic[i] = pfn.regIndex(state.Chan)
			if state.Send != nil {
				is[i] = pfn.regIndex(state.Send)
			}
		}
		return func(fr *frame) {
			var cases []reflect.SelectCase
			if !instr.Blocking {
				cases = append(cases, reflect.SelectCase{
					Dir: reflect.SelectDefault,
				})
			}
			for i, state := range instr.States {
				var dir reflect.SelectDir
				if state.Dir == types.RecvOnly {
					dir = reflect.SelectRecv
				} else {
					dir = reflect.SelectSend
				}
				ch := reflect.ValueOf(fr.reg(ic[i]))
				var send reflect.Value
				if state.Send != nil {
					v := fr.reg(is[i])
					if v == nil {
						send = reflect.New(ch.Type().Elem()).Elem()
					} else {
						send = reflect.ValueOf(v)
					}
				}
				cases = append(cases, reflect.SelectCase{
					Dir:  dir,
					Chan: ch,
					Send: send,
				})
			}
			chosen, recv, recvOk := reflect.Select(cases)
			if !instr.Blocking {
				chosen-- // default case should have index -1.
			}
			r := tuple{chosen, recvOk}
			for n, st := range instr.States {
				if st.Dir == types.RecvOnly {
					var v value
					if n == chosen && recvOk {
						// No need to copy since send makes an unaliased copy.
						v = recv.Interface()
					} else {
						typ := interp.toType(st.Chan.Type())
						v = reflect.New(typ.Elem()).Elem().Interface()
					}
					r = append(r, v)
				}
			}
			fr.setReg(ir, r)
		}
	case *ssa.SliceToArrayPointer:
		typ := interp.preToType(instr.Type())
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame) {
			x := fr.reg(ix)
			v := reflect.ValueOf(x)
			vLen := v.Len()
			tLen := typ.Elem().Len()
			if tLen > vLen {
				panic(fr.runtimeError(instr, fmt.Sprintf(errSliceToArrayPointer, vLen, tLen)))
			}
			fr.setReg(ir, v.Convert(typ).Interface())
		}
	case *ssa.Range:
		typ := interp.preToType(instr.X.Type())
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		switch typ.Kind() {
		case reflect.String:
			return func(fr *frame) {
				v := fr.string(ix)
				fr.setReg(ir, &stringIter{Reader: strings.NewReader(v)})
			}
		case reflect.Map:
			return func(fr *frame) {
				v := fr.reg(ix)
				fr.setReg(ir, &mapIter{iter: reflect.ValueOf(v).MapRange()})
			}
		default:
			panic("unreachable")
		}
	case *ssa.Next:
		ir := pfn.regIndex(instr)
		ii := pfn.regIndex(instr.Iter)
		if instr.IsString {
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ii).(*stringIter).next())
			}
		}
		return func(fr *frame) {
			fr.setReg(ir, fr.reg(ii).(*mapIter).next())
		}
	case *ssa.TypeAssert:
		typ := interp.preToType(instr.AssertedType)
		xtyp := interp.preToType(instr.X.Type())
		ir := pfn.regIndex(instr)
		ix, kx, vx := pfn.regIndex3(instr.X)
		if kx.isStatic() {
			return func(fr *frame) {
				fr.setReg(ir, typeAssert(fr, instr, typ, xtyp, vx))
			}
		}
		return func(fr *frame) {
			v := fr.reg(ix)
			fr.setReg(ir, typeAssert(fr, instr, typ, xtyp, v))
		}
	case *ssa.Extract:
		if *instr.Referrers() == nil {
			return nil
		}
		ir := pfn.regIndex(instr)
		it := pfn.regIndex(instr.Tuple)
		return func(fr *frame) {
			fr.setReg(ir, fr.reg(it).(tuple)[instr.Index])
		}
	// Instructions executed for effect
	case *ssa.Jump:
		return func(fr *frame) {
			fr.pred, fr.block = fr.block.Index, fr.block.Succs[0]
			fr.ipc = fr.pfn.Blocks[fr.block.Index]
		}
	case *ssa.If:
		ic, kc, vc := pfn.regIndex3(instr.Cond)
		if kc == kindConst {
			if xtype.Bool(vc) {
				return func(fr *frame) {
					fr.pred = fr.block.Index
					fr.block = fr.block.Succs[0]
					fr.ipc = fr.pfn.Blocks[fr.block.Index]
				}
			}
			return func(fr *frame) {
				fr.pred = fr.block.Index
				fr.block = fr.block.Succs[1]
				fr.ipc = fr.pfn.Blocks[fr.block.Index]
			}
		}
		switch instr.Cond.Type().(type) {
		case *types.Basic:
			return func(fr *frame) {
				fr.pred = fr.block.Index
				if fr.reg(ic).(bool) {
					fr.block = fr.block.Succs[0]
				} else {
					fr.block = fr.block.Succs[1]
				}
				fr.ipc = fr.pfn.Blocks[fr.block.Index]
			}
		default:
			return func(fr *frame) {
				fr.pred = fr.block.Index
				if fr.bool(ic) {
					fr.block = fr.block.Succs[0]
				} else {
					fr.block = fr.block.Succs[1]
				}
				fr.ipc = fr.pfn.Blocks[fr.block.Index]
			}
		}
	case *ssa.Return:
		switch n := pfn.nres; n {
		case 0:
			return func(fr *frame) {
				fr.ipc = -1
			}
		case 1:
			ir, ik, iv := pfn.regIndex3(instr.Results[0])
			if ik.isStatic() {
				return func(fr *frame) {
					fr.stack[0] = iv
					fr.ipc = -1
				}
			}
			return func(fr *frame) {
				fr.stack[0] = fr.reg(ir)
				fr.ipc = -1
			}
		case 2:
			r1, k1, v1 := pfn.regIndex3(instr.Results[0])
			r2, k2, v2 := pfn.regIndex3(instr.Results[1])
			if k1.isStatic() && k2.isStatic() {
				return func(fr *frame) {
					fr.stack[0] = v1
					fr.stack[1] = v2
					fr.ipc = -1
				}
			} else if k1.isStatic() {
				return func(fr *frame) {
					fr.stack[0] = v1
					fr.stack[1] = fr.reg(r2)
					fr.ipc = -1
				}
			} else if k2.isStatic() {
				return func(fr *frame) {
					fr.stack[0] = fr.reg(r1)
					fr.stack[1] = v2
					fr.ipc = -1
				}
			}
			return func(fr *frame) {
				fr.stack[0] = fr.reg(r1)
				fr.stack[1] = fr.reg(r2)
				fr.ipc = -1
			}
		default:
			ir := make([]register, n)
			for i, v := range instr.Results {
				ir[i] = pfn.regIndex(v)
			}
			return func(fr *frame) {
				for i := 0; i < n; i++ {
					fr.stack[i] = fr.reg(ir[i])
				}
				fr.ipc = -1
			}
		}
	case *ssa.RunDefers:
		return func(fr *frame) {
			fr.runDefers()
		}
	case *ssa.Panic:
		ix := pfn.regIndex(instr.X)
		if interp.ctx.panicFunc != nil {
			return func(fr *frame) {
				var err error = PanicError{stack: debugStack(fr), Value: fr.reg(ix)}
				panic(interp.ctx.handlePanic(fr, instr, err))
			}
		}
		return func(fr *frame) {
			panic(PanicError{stack: debugStack(fr), Value: fr.reg(ix)})
		}
	case *ssa.Go:
		iv, ia, ib := getCallIndex(pfn, &instr.Call)
		return func(fr *frame) {
			fn, args := interp.prepareCall(fr, &instr.Call, iv, ia, ib)
			atomic.AddInt32(&interp.goroutines, 1)
			if interp.ctx.RunContext != nil {
				go func() {
					root := &frame{interp: interp}
					switch f := fn.(type) {
					case *ssa.Function:
						root.pfn = interp.funcs[f]
					case *closure:
						root.pfn = f.pfn
					}
					defer func() {
						e := recover()
						if e != nil {
							interp.cherror <- PanicError{stack: debugStack(root), Value: e}
						}
					}()
					interp.callDiscardsResult(root, fn, args, instr.Call.Args)
					atomic.AddInt32(&interp.goroutines, -1)
				}()
			} else {
				go func() {
					interp.callDiscardsResult(&frame{}, fn, args, instr.Call.Args)
					atomic.AddInt32(&interp.goroutines, -1)
				}()
			}
		}
	case *ssa.Defer:
		return makeDefer(interp, pfn, instr)
	case *ssa.Send:
		ic := pfn.regIndex(instr.Chan)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame) {
			c := fr.reg(ic)
			x := fr.reg(ix)
			ch := reflect.ValueOf(c)
			if x == nil {
				ch.Send(reflect.New(ch.Type().Elem()).Elem())
			} else {
				ch.Send(reflect.ValueOf(x))
			}
		}
	case *ssa.Store:
		// skip struct field _
		if addr, ok := instr.Addr.(*ssa.FieldAddr); ok {
			if s, ok := addr.X.Type().(*types.Pointer).Elem().(*types.Struct); ok {
				if s.Field(addr.Field).Name() == "_" {
					return nil
				}
			}
		}
		if pfn.Fn.Name() == "init" && pfn.Fn.Synthetic == "package initializer" {
			pfn.Interp.chkinit[instr.Addr.String()] = true
		}
		ia := pfn.regIndex(instr.Addr)
		iv, kv, vv := pfn.regIndex3(instr.Val)
		if kv.isStatic() {
			if vv == nil {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ia))
					SetValue(x.Elem(), reflect.New(x.Elem().Type()).Elem())
				}
			}
			return func(fr *frame) {
				x := reflect.ValueOf(fr.reg(ia))
				SetValue(x.Elem(), reflect.ValueOf(vv))
			}
		}
		return func(fr *frame) {
			x := reflect.ValueOf(fr.reg(ia))
			val := fr.reg(iv)
			v := reflect.ValueOf(val)
			if v.IsValid() {
				SetValue(x.Elem(), v)
			} else {
				SetValue(x.Elem(), reflect.New(x.Elem().Type()).Elem())
			}
		}
	case *ssa.MapUpdate:
		im := pfn.regIndex(instr.Map)
		ik := pfn.regIndex(instr.Key)
		iv, kv, vv := pfn.regIndex3(instr.Value)
		if kv.isStatic() {
			return func(fr *frame) {
				vm := reflect.ValueOf(fr.reg(im))
				vk := reflect.ValueOf(fr.reg(ik))
				vm.SetMapIndex(vk, reflect.ValueOf(vv))
			}
		}
		return func(fr *frame) {
			vm := reflect.ValueOf(fr.reg(im))
			vk := reflect.ValueOf(fr.reg(ik))
			v := fr.reg(iv)
			vm.SetMapIndex(vk, reflect.ValueOf(v))
		}
	case *ssa.DebugRef:
		if v, ok := instr.Object().(*types.Var); ok {
			ix := pfn.regIndex(instr.X)
			return func(fr *frame) {
				ref := &DebugInfo{DebugRef: instr, fset: interp.ctx.FileSet}
				ref.toValue = func() (*types.Var, interface{}, bool) {
					return v, fr.reg(ix), true
				}
				interp.ctx.debugFunc(ref)
			}
		}
		return func(fr *frame) {
			ref := &DebugInfo{DebugRef: instr, fset: interp.ctx.FileSet}
			ref.toValue = func() (*types.Var, interface{}, bool) {
				return nil, nil, false
			}
			interp.ctx.debugFunc(ref)
		}
	default:
		panic(fmt.Errorf("unreachable %T", instr))
	}
}

func getCallIndex(pfn *function, call *ssa.CallCommon) (iv register, ia []register, ib []register) {
	iv = pfn.regIndex(call.Value)
	ia = make([]register, len(call.Args))
	for i, v := range call.Args {
		ia[i] = pfn.regIndex(v)
	}
	if f, ok := call.Value.(*ssa.MakeClosure); ok {
		ib = make([]register, len(f.Bindings))
		for i, binding := range f.Bindings {
			ib[i] = pfn.regIndex(binding)
		}
	}
	return
}

var (
	typFramePtr = reflect.TypeOf((*frame)(nil))
)

func makeCallInstr(pfn *function, interp *Interp, instr ssa.Value, call *ssa.CallCommon) func(fr *frame) {
	ir := pfn.regIndex(instr)
	iv, ia, ib := getCallIndex(pfn, call)
	switch fn := call.Value.(type) {
	case *ssa.Builtin:
		return interp.makeBuiltinByStack(fn, call.Args, ir, ia)
	case *ssa.MakeClosure:
		ifn := interp.loadFunction(fn.Fn.(*ssa.Function))
		ia = append(ia, ib...)
		if ifn.Recover == nil {
			switch ifn.nres {
			case 0:
				return func(fr *frame) {
					interp.callFunctionByStackNoRecover0(fr, ifn, ir, ia)
				}
			case 1:
				return func(fr *frame) {
					interp.callFunctionByStackNoRecover1(fr, ifn, ir, ia)
				}
			default:
				return func(fr *frame) {
					interp.callFunctionByStackNoRecoverN(fr, ifn, ir, ia)
				}
			}
		}
		switch ifn.nres {
		case 0:
			return func(fr *frame) {
				interp.callFunctionByStack0(fr, ifn, ir, ia)
			}
		case 1:
			return func(fr *frame) {
				interp.callFunctionByStack1(fr, ifn, ir, ia)
			}
		default:
			return func(fr *frame) {
				interp.callFunctionByStackN(fr, ifn, ir, ia)
			}
		}
	case *ssa.Function:
		// "static func/method call"
		if fn.Blocks == nil {
			ext, ok := findExternFunc(interp, fn)
			if !ok {
				// skip pkg.init
				if fn.Pkg != nil && fn.Name() == "init" {
					return nil
				}
				panic(fmt.Errorf("no code for function: %v", fn))
			}
			if typ := ext.Type(); typ.NumIn() > 0 && typ.In(0) == typFramePtr {
				return func(fr *frame) {
					interp.callExternalWithFrameByStack(fr, ext, ir, ia)
				}
			}
			return func(fr *frame) {
				interp.callExternalByStack(fr, ext, ir, ia)
			}
		}
		ifn := interp.loadFunction(fn)
		if ifn.Recover == nil {
			switch ifn.nres {
			case 0:
				return func(fr *frame) {
					interp.callFunctionByStackNoRecover0(fr, ifn, ir, ia)
				}
			case 1:
				return func(fr *frame) {
					interp.callFunctionByStackNoRecover1(fr, ifn, ir, ia)
				}
			default:
				return func(fr *frame) {
					interp.callFunctionByStackNoRecoverN(fr, ifn, ir, ia)
				}
			}
		}
		switch ifn.nres {
		case 0:
			return func(fr *frame) {
				interp.callFunctionByStack0(fr, ifn, ir, ia)
			}
		case 1:
			return func(fr *frame) {
				interp.callFunctionByStack1(fr, ifn, ir, ia)
			}
		default:
			return func(fr *frame) {
				interp.callFunctionByStackN(fr, ifn, ir, ia)
			}
		}
	}
	// "dynamic method call" // ("invoke" mode)
	if call.IsInvoke() {
		return makeCallMethodInstr(interp, instr, call, ir, iv, ia)
	}
	// dynamic func call
	typ := interp.preToType(call.Value.Type())
	if typ.Kind() != reflect.Func {
		panic("unsupport")
	}
	if supportFuncVal && (interp.ctx.Mode&DisableDynamicFuncCallAnalysis == 0) {
		return dynamicFunCall(interp, iv, ir, ia)
	}
	return func(fr *frame) {
		fn := fr.reg(iv)
		v := reflect.ValueOf(fn)
		interp.callExternalByStack(fr, v, ir, ia)
	}
}

var (
	typeOfType      = reflect.TypeOf(reflect.TypeOf(0))
	vfnMethod       = reflect.ValueOf(reflectx.MethodByIndex)
	vfnMethodByName = reflect.ValueOf(reflectx.MethodByName)
)

func findUserMethod(typ reflect.Type, name string) (ext reflect.Value, ok bool) {
	if m, ok := reflectx.MethodByName(typ, name); ok {
		return m.Func, true
	}
	return
}

func findExternMethod(typ reflect.Type, name string) (ext reflect.Value, ok bool) {
	if typ == typeOfType {
		switch name {
		case "Method":
			return vfnMethod, true
		case "MethodByName":
			return vfnMethodByName, true
		}
	}
	if m, ok := typ.MethodByName(name); ok {
		return m.Func, true
	}
	return
}

func (i *Interp) findMethod(typ reflect.Type, mname string) (fn *ssa.Function, ok bool) {
	if mset, mok := i.msets[typ]; mok {
		fn, ok = mset[mname]
	}
	return
}

func makeCallMethodInstr(interp *Interp, instr ssa.Value, call *ssa.CallCommon, ir register, iv register, ia []register) func(fr *frame) {
	mname := call.Method.Name()
	ia = append([]register{iv}, ia...)
	var found bool
	var ext reflect.Value
	return func(fr *frame) {
		v := fr.reg(iv)
		if v == nil {
			panic(fr.runtimeError(instr, "runtime error: invalid memory address or nil pointer dereference"))
		}
		rtype := reflect.TypeOf(v)
		// find user type method *ssa.Function
		if mset, ok := interp.msets[rtype]; ok {
			if fn, ok := mset[mname]; ok {
				interp.callFunctionByStack(fr, interp.funcs[fn], ir, ia)
				return
			}
			ext, found = findUserMethod(rtype, mname)
		} else {
			ext, found = findExternMethod(rtype, mname)
		}
		if !found {
			panic(fr.plainError(instr, fmt.Sprintf("no code for method: %v.%v", rtype, mname)))
		}
		interp.callExternalByStack(fr, ext, ir, ia)
	}
}

type stringIter struct {
	*strings.Reader
	i int
}

func (it *stringIter) next() tuple {
	okv := make(tuple, 3)
	ch, n, err := it.ReadRune()
	ok := err != io.EOF
	okv[0] = ok
	if ok {
		okv[1] = it.i
		okv[2] = ch
	}
	it.i += n
	return okv
}

type mapIter struct {
	iter *reflect.MapIter
	ok   bool
}

func (it *mapIter) next() tuple {
	it.ok = it.iter.Next()
	if !it.ok {
		return []value{false, nil, nil}
	}
	k, v := it.iter.Key().Interface(), it.iter.Value().Interface()
	return []value{true, k, v}
}

func toKind(typ types.Type) reflect.Kind {
retry:
	switch t := typ.(type) {
	case *types.Basic:
		switch t.Kind() {
		case types.Bool:
			return reflect.Bool
		case types.Int:
			return reflect.Int
		case types.Int8:
			return reflect.Int8
		case types.Int16:
			return reflect.Int16
		case types.Int32:
			return reflect.Int32
		case types.Int64:
			return reflect.Int64
		case types.Uint:
			return reflect.Uint
		case types.Uint8:
			return reflect.Uint8
		case types.Uint16:
			return reflect.Uint16
		case types.Uint32:
			return reflect.Uint32
		case types.Uint64:
			return reflect.Uint64
		case types.Uintptr:
			return reflect.Uintptr
		case types.Float32:
			return reflect.Float32
		case types.Float64:
			return reflect.Float64
		case types.Complex64:
			return reflect.Complex64
		case types.Complex128:
			return reflect.Complex128
		case types.String:
			return reflect.String
		case types.UnsafePointer:
			return reflect.UnsafePointer
		// types for untyped values
		case types.UntypedBool:
			return reflect.Bool
		case types.UntypedInt:
			return reflect.Int
		case types.UntypedRune:
			return reflect.Int32
		case types.UntypedFloat:
			return reflect.Float64
		case types.UntypedComplex:
			return reflect.Complex128
		case types.UntypedString:
			return reflect.String
		case types.UntypedNil:
			return reflect.Invalid
		}
	case *types.Chan:
		return reflect.Chan
	case *types.Named:
		typ = t.Underlying()
		goto retry
	case *types.Signature:
		return reflect.Func
	case *types.Slice:
		return reflect.Slice
	case *types.Array:
		return reflect.Array
	case *types.Pointer:
		return reflect.Ptr
	case *types.Struct:
		return reflect.Struct
	case *types.Map:
		return reflect.Map
	case *types.Interface:
		return reflect.Interface
	case *types.Tuple:
		return reflect.Slice
	case *types.Alias:
		typ = types.Unalias(t)
		goto retry
	}
	return reflect.Invalid
}

func makeDefer(interp *Interp, pfn *function, instr *ssa.Defer) func(fr *frame) {
	iv, ia, ib := getCallIndex(pfn, &instr.Call)
	if instr.DeferStack == nil {
		return func(fr *frame) {
			fn, args := interp.prepareCall(fr, &instr.Call, iv, ia, ib)
			fr._defer = &_defer{
				fn:      fn,
				args:    args,
				ssaArgs: instr.Call.Args,
				tail:    fr._defer,
			}
		}
	}
	id := pfn.regIndex(instr.DeferStack)
	return func(fr *frame) {
		fn, args := interp.prepareCall(fr, &instr.Call, iv, ia, ib)
		defers := fr.reg(id).(**_defer)
		*defers = &_defer{
			fn:      fn,
			args:    args,
			ssaArgs: instr.Call.Args,
			tail:    *defers,
		}
	}
}
