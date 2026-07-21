//go:build llgo
// +build llgo

package ixgo

import (
	"go/token"
	"path/filepath"
	"reflect"
	"runtime"
	"sync"
	"unsafe"
)

func init() {
	RegisterExternal("(reflect.Value).Pointer", func(fr *frame, v reflect.Value) uintptr {
		if v.Kind() == reflect.Func {
			if c := fr.interp.getMakeFuncValue(v); c != nil && c.interp == fr.interp {
				return uintptr(c.pfn.base)
			}
		}
		return v.Pointer()
	})
	RegisterExternal("(reflect.Value).UnsafePointer", func(fr *frame, v reflect.Value) unsafe.Pointer {
		if v.Kind() == reflect.Func {
			if c := fr.interp.getMakeFuncValue(v); c != nil && c.interp == fr.interp {
				return unsafe.Pointer(uintptr(c.pfn.base))
			}
		}
		return v.UnsafePointer()
	})
}

type funcinl struct {
	entry uintptr
	name  string
	pc    uintptr
	file  string
	line  int
}

func inlineFunc(entry uintptr) *funcinl {
	fn := &funcinl{entry: entry}
	return fn
}

func isInlineFunc(f *runtime.Func) bool {
	return true
}

//go:linkname runtimePanic github.com/goplus/llgo/runtime/internal/runtime.Panic
func runtimePanic(e interface{})

func runtimeFuncFileLine(fr *frame, f *runtime.Func, pc uintptr) (file string, line int) {
	entry := f.Entry()
	if isInlineFunc(f) && pc >= entry {
		interp := fr.interp
		if pfn := findFuncByEntry(interp, int(entry)); pfn != nil {
			var pos token.Pos
			if pc == entry {
				pos = pfn.Fn.Pos()
			} else {
				// pc-1 : fn.instr.pos
				pos = pfn.PosForPC(int(pc - entry - 1))
				if !pos.IsValid() {
					return "?", 0
				}
			}
			fpos := interp.ctx.FileSet.Position(pos)
			if fpos.Filename == "" {
				return "??", fpos.Line
			}
			file, line = filepath.ToSlash(fpos.Filename), fpos.Line
			return
		}
	}
	return f.FileLine(pc)
}

const supportFuncVal = true

func dynamicFunCall(interp *Interp, iv register, ir register, ia []register) func(fr *frame) {
	return func(fr *frame) {
		fn := fr.reg(iv)
		if c := interp.getMakeFuncVal(fn); c != nil && c.interp == interp {
			if c.pfn.Recover == nil {
				interp.callFunctionByStackNoRecoverWithEnv(fr, c.pfn, ir, ia, c.env)
			} else {
				interp.callFunctionByStackWithEnv(fr, c.pfn, ir, ia, c.env)
			}
			return
		}
		v := reflect.ValueOf(fn)
		interp.callExternalByStack(fr, v, ir, ia)
	}
}

type interpExt struct {
	makeFuncs sync.Map
}

func (i *interpExt) getMakeFuncValue(v reflect.Value) *makeFuncVal {
	pv := (*reflectValue)(unsafe.Pointer(&v))
	if r, ok := i.makeFuncs.Load(pv.ptr); ok {
		return r.(*makeFuncVal)
	}
	return nil
}

func (i *interpExt) getMakeFuncVal(v interface{}) *makeFuncVal {
	e := (*emptyInterface)(unsafe.Pointer(&v))
	if r, ok := i.makeFuncs.Load(e.word); ok {
		return r.(*makeFuncVal)
	}
	return nil
}

func (pfn *function) makeFunction(typ reflect.Type, env []value) reflect.Value {
	interp := pfn.Interp
	v := reflect.MakeFunc(typ, func(args []reflect.Value) []reflect.Value {
		return interp.callFunctionByReflect(interp.tryDeferFrame(), pfn, typ, args, env)
	})
	fn := (*reflectValue)(unsafe.Pointer(&v)).ptr
	interp.makeFuncs.Store(fn, &makeFuncVal{
		Fn:     uintptr(fn),
		interp: interp,
		pfn:    pfn,
		typ:    typ,
		env:    env,
	})
	return v
}

type makeFuncVal struct {
	Fn     uintptr
	interp *Interp
	pfn    *function
	typ    reflect.Type
	env    []interface{}
}
