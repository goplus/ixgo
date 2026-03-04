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
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"reflect"
	"strings"

	"github.com/goplus/ixgo/load"
	"github.com/visualfc/xtype"
	"golang.org/x/tools/go/ssa"
)

const (
	fnBase = 100
)

func checkPackages(intp *Interp, pkgs []*ssa.Package) (err error) {
	if intp.ctx.Mode&DisableRecover == 0 {
		defer func() {
			if v := recover(); v != nil {
				if e, ok := v.(error); ok {
					err = e
				} else {
					err = fmt.Errorf("%v", v)
				}
			}
		}()
	}
	visit := visitor{
		intp: intp,
		prog: intp.mainpkg.Prog,
		pkgs: make(map[*ssa.Package]bool),
		seen: make(map[*ssa.Function]bool),
		base: fnBase,
	}
	for _, pkg := range pkgs {
		visit.pkgs[pkg] = true
	}
	visit.program()
	return
}

type visitor struct {
	intp *Interp
	prog *ssa.Program
	pkgs map[*ssa.Package]bool
	seen map[*ssa.Function]bool
	base int
}

func (visit *visitor) program() {
	chks := make(map[string]bool)
	chks[""] = true // anonymous struct embed named type
	for pkg := range visit.pkgs {
		chks[pkg.Pkg.Path()] = true
	}

	isExtern := func(typ reflect.Type) bool {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		return !chks[typ.PkgPath()]
	}

	methodsOf := func(T types.Type) {
		if types.IsInterface(T) {
			return
		}
		typ := visit.intp.preToType(T)
		// skip extern type
		if isExtern(typ) {
			return
		}
		mmap := make(map[string]*ssa.Function)
		mset := visit.prog.MethodSets.MethodSet(T)
		checkXGo := visit.intp.ctx.Mode&CheckGopOverloadFunc != 0
		for i, n := 0, mset.Len(); i < n; i++ {
			sel := mset.At(i)
			obj := sel.Obj()
			// skip embed extern type method
			if pkg := obj.Pkg(); pkg != nil {
				if !chks[pkg.Path()] {
					continue
				}
				if checkXGo && isXGoOverloadFunc(obj.Type().(*types.Signature)) {
					continue
				}
			}
			fn := visit.prog.MethodValue(sel)
			mmap[obj.Name()] = fn
			visit.function(fn)
		}
		visit.intp.msets[typ] = mmap
	}

	exportedTypeHack := func(t *ssa.Type) {
		if ast.IsExported(t.Name()) && !types.IsInterface(t.Type()) {
			if named, ok := t.Type().(*types.Named); ok && !hasTypeParam(named) {
				methodsOf(named)                   //  T
				methodsOf(types.NewPointer(named)) // *T
			}
		}
	}

	for pkg := range visit.pkgs {
		for _, mem := range pkg.Members {
			switch mem := mem.(type) {
			case *ssa.Function:
				visit.function(mem)
			case *ssa.Type:
				exportedTypeHack(mem)
			}
		}
	}

	for _, T := range visit.prog.RuntimeTypes() {
		methodsOf(T)
	}
}

func (visit *visitor) findLinkSym(fn *ssa.Function) (*load.LinkSym, bool) {
	if sp, ok := visit.intp.ctx.pkgs[fn.Pkg.Pkg.Path()]; ok {
		for _, link := range sp.Links {
			if link.Name == fn.Name() {
				return link, true
			}
		}
	}
	return nil, false
}

func (visit *visitor) findFunction(sym *load.LinkSym) *ssa.Function {
	for pkg := range visit.pkgs {
		if pkg.Pkg.Path() == sym.Linkname.PkgPath {
			if typ := sym.Linkname.Recv; typ != "" {
				var star bool
				if typ[0] == '*' {
					star = true
					typ = typ[1:]
				}
				if obj := pkg.Pkg.Scope().Lookup(typ); obj != nil {
					t := obj.Type()
					if star {
						t = types.NewPointer(t)
					}
					return visit.prog.LookupMethod(t, pkg.Pkg, sym.Linkname.Method)
				}
			}
			return pkg.Func(sym.Linkname.Name)
		}
	}
	return nil
}

func wrapMethodType(sig *types.Signature) *types.Signature {
	params := sig.Params()
	n := params.Len()
	list := make([]*types.Var, n+1)
	list[0] = sig.Recv()
	for i := 0; i < n; i++ {
		list[i+1] = params.At(i)
	}
	return types.NewSignature(nil, types.NewTuple(list...), sig.Results(), sig.Variadic())
}

func (visit *visitor) findLinkFunc(sym *load.LinkSym) (ext reflect.Value, ok bool) {
	ext, ok = findExternLinkFunc(visit.intp, &sym.Linkname)
	if ok {
		return
	}
	if link := visit.findFunction(sym); link != nil {
		visit.function(link)
		sig := link.Signature
		if sig.Recv() != nil {
			sig = wrapMethodType(sig)
		}
		typ := visit.intp.preToType(sig)
		pfn := visit.intp.loadFunction(link)
		ext = pfn.makeFunction(typ, nil)
		ok = true
	}
	return
}

func (visit *visitor) function(fn *ssa.Function) {
	if visit.seen[fn] {
		return
	}
	if hasTypeParam(fn.Type()) {
		return
	}
	visit.seen[fn] = true
	fnPath := fn.String()
	if f, ok := visit.intp.ctx.override[fnPath]; ok &&
		visit.intp.preToType(fn.Type()) == f.Type() {
		fn.Blocks = nil
		return
	}
	if fn.Blocks == nil {
		if _, ok := visit.pkgs[fn.Pkg]; ok {
			if _, ok = findExternFunc(visit.intp, fn); !ok {
				if sym, ok := visit.findLinkSym(fn); ok {
					if ext, ok := visit.findLinkFunc(sym); ok {
						typ := visit.intp.preToType(fn.Type())
						ftyp := ext.Type()
						if typ != ftyp {
							ext = xtype.ConvertFunc(ext, xtype.TypeOfType(typ))
						}
						visit.intp.ctx.override[fnPath] = ext
						return
					}
				}
				if visit.intp.ctx.Mode&EnableNoStrict != 0 {
					typ := visit.intp.preToType(fn.Type())
					numOut := typ.NumOut()
					if numOut == 0 {
						visit.intp.ctx.override[fnPath] = reflect.MakeFunc(typ, func(args []reflect.Value) (results []reflect.Value) {
							return
						})
					} else {
						visit.intp.ctx.override[fnPath] = reflect.MakeFunc(typ, func(args []reflect.Value) (results []reflect.Value) {
							results = make([]reflect.Value, numOut)
							for i := 0; i < numOut; i++ {
								results[i] = reflect.New(typ.Out(i)).Elem()
							}
							return
						})
					}
					println(fmt.Sprintf("ixgo warning: %v: %v missing function body", visit.intp.ctx.FileSet.Position(fn.Pos()), fnPath))
					return
				}
				panic(fmt.Errorf("%v: %v missing function body", visit.intp.ctx.FileSet.Position(fn.Pos()), fnPath))
			}
		}
		return
	}
	if len(fn.TypeArgs()) != 0 {
		visit.intp.record.EnterInstance(fn)
		defer visit.intp.record.LeaveInstance(fn)
	}
	visit.intp.loadType(fn.Type())
	for _, alloc := range fn.Locals {
		visit.intp.loadType(alloc.Type())
		visit.intp.loadType(deref(alloc.Type()))
	}
	pfn := visit.intp.loadFunction(fn)
	for _, p := range fn.Params {
		pfn.regIndex(p)
	}
	for _, p := range fn.FreeVars {
		pfn.regIndex(p)
	}
	var buf [32]*ssa.Value // avoid alloc in common case
	for _, b := range fn.Blocks {
		Instrs := make([]func(*frame), len(b.Instrs))
		ssaInstrs := make([]ssa.Instruction, len(b.Instrs))
		var index int
		n := len(b.Instrs)
		for i := 0; i < n; i++ {
			instr := b.Instrs[i]
			ops := instr.Operands(buf[:0])
			switch instr := instr.(type) {
			case *ssa.Alloc:
				visit.intp.loadType(instr.Type())
				visit.intp.loadType(deref(instr.Type()))
			case *ssa.Next:
				// skip *ssa.opaqueType: iter
				ops = nil
			case *ssa.Extract:
				// skip
				ops = nil
			case *ssa.TypeAssert:
				visit.intp.loadType(instr.AssertedType)
			case *ssa.MakeChan:
				visit.intp.loadType(instr.Type())
			case *ssa.MakeMap:
				visit.intp.loadType(instr.Type())
			case *ssa.MakeSlice:
				visit.intp.loadType(instr.Type())
			case *ssa.SliceToArrayPointer:
				visit.intp.loadType(instr.Type())
			case *ssa.Convert:
				visit.intp.loadType(instr.Type())
			case *ssa.ChangeType:
				visit.intp.loadType(instr.Type())
			case *ssa.MakeInterface:
				visit.intp.loadType(instr.Type())
			}
			for _, op := range ops {
				switch v := (*op).(type) {
				case *ssa.Function:
					visit.function(v)
				case nil:
					// skip
				default:
					visit.intp.loadType(v.Type())
				}
			}
			pfn.makeInstr = instr
			ifn := makeInstr(visit.intp, pfn, instr)
			if ifn == nil {
				continue
			}
			if visit.intp.ctx.evalMode && fn.String() == "main.init" {
				if visit.intp.ctx.evalInit == nil {
					visit.intp.ctx.evalInit = make(map[string]bool)
				}
				if call, ok := instr.(*ssa.Call); ok {
					key := call.String()
					if strings.HasPrefix(key, "init#") {
						if visit.intp.ctx.evalInit[key] {
							ifn = func(fr *frame) {}
						} else {
							visit.intp.ctx.evalInit[key] = true
						}
					}
				}
			}
			if visit.intp.ctx.evalCallFn != nil {
				if call, ok := instr.(*ssa.Call); ok {
					ir := pfn.regIndex(call)
					results := call.Call.Signature().Results()
					ofn := ifn
					switch results.Len() {
					case 0:
						ifn = func(fr *frame) {
							ofn(fr)
							visit.intp.ctx.evalCallFn(visit.intp, call)
						}
					case 1:
						ifn = func(fr *frame) {
							ofn(fr)
							visit.intp.ctx.evalCallFn(visit.intp, call, fr.reg(ir))
						}
					default:
						ifn = func(fr *frame) {
							ofn(fr)
							r := fr.reg(ir).(tuple)
							visit.intp.ctx.evalCallFn(visit.intp, call, r...)
						}
					}
				}
			}
			if visit.intp.ctx.Mode&EnableTracing != 0 {
				ofn := ifn
				ifn = func(fr *frame) {
					if v, ok := instr.(ssa.Value); ok {
						log.Printf("\t%-20T %v = %-40v\t%v\n", instr, v.Name(), instr, v.Type())
					} else {
						log.Printf("\t%-20T %v\n", instr, instr)
					}
					ofn(fr)
				}
				if index == 0 {
					ofn := ifn
					bi := b.Index
					common := b.Comment
					ifn = func(fr *frame) {
						log.Printf(".%v %v\n", bi, common)
						ofn(fr)
					}
				}
				if index == 0 && b.Index == 0 {
					ofn := ifn
					ifn = func(fr *frame) {
						log.Printf("Entering %v%v.", fr.pfn.Fn, loc(fr.interp.ctx.FileSet, fr.pfn.Fn.Pos()))
						ofn(fr)
					}
				}
				if _, ok := instr.(*ssa.Return); ok {
					ofn := ifn
					ifn = func(fr *frame) {
						ofn(fr)
						var caller ssa.Instruction
						if fr.caller != nil {
							caller = fr.caller.pfn.InstrForPC(fr.caller.ipc - 1)
						}
						if caller == nil {
							log.Printf("Leaving %v.\n", fr.pfn.Fn)
						} else {
							log.Printf("Leaving %v, resuming %v call %v%v.\n",
								fr.pfn.Fn, fr.caller.pfn.Fn, caller, loc(fr.interp.ctx.FileSet, caller.Pos()))
						}
					}
				}
			}
			Instrs[index] = ifn
			ssaInstrs[index] = instr
			index++
		}
		offset := len(pfn.Instrs)
		pfn.Blocks = append(pfn.Blocks, offset)
		pfn.Instrs = append(pfn.Instrs, Instrs[:index]...)
		pfn.ssaInstrs = append(pfn.ssaInstrs, ssaInstrs[:index]...)
		if b == fn.Recover && visit.intp.ctx.Mode&DisableRecover == 0 {
			pfn.Recover = pfn.Instrs[offset:]
		}
	}
	pfn.makeInstr = nil
	pfn.base = visit.base
	visit.base += len(pfn.ssaInstrs) + 2
	pfn.initPool()
}

func loc(fset *token.FileSet, pos token.Pos) string {
	if pos == token.NoPos {
		return ""
	}
	return " at " + fset.Position(pos).String()
}
