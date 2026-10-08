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

const (
	Compiler = "llgo"
	IsLLGo   = true
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

//go:linkname runtimePanic github.com/xgo-dev/llgo/runtime/internal/runtime.Panic
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

// llgoClosure is the FFI-backed data block used by LLGo's reflect.MakeFunc.
type llgoClosure struct {
	fn  unsafe.Pointer
	env unsafe.Pointer
}

func (i *interpExt) getMakeFuncValue(v reflect.Value) *FuncVal {
	pv := (*reflectValue)(unsafe.Pointer(&v))
	return i.loadMakeFunc(pv.ptr)
}

func (i *interpExt) getMakeFuncVal(v interface{}) *FuncVal {
	e := (*emptyInterface)(unsafe.Pointer(&v))
	return i.loadMakeFunc(e.word)
}

func (i *interpExt) loadMakeFunc(ptr unsafe.Pointer) *FuncVal {
	if ptr == nil {
		return nil
	}
	if r, ok := i.makeFuncs.Load(ptr); ok {
		return r.(*FuncVal)
	}
	r, ok := i.makeFuncs.Load((*llgoClosure)(ptr).fn)
	if !ok {
		return nil
	}
	return r.(*FuncVal)
}

func (pfn *function) makeFunction(typ reflect.Type, env []value) reflect.Value {
	interp := pfn.Interp
	c := &FuncVal{interp: interp, pfn: pfn, typ: typ, env: env}
	if v, ok := makeTypedFunction(*c, typ); ok {
		interp.makeFuncs.Store((*reflectValue)(unsafe.Pointer(&v)).ptr, c)
		return v
	}
	v := reflect.MakeFunc(typ, func(args []reflect.Value) []reflect.Value {
		return interp.callFunctionByReflect(interp.tryDeferFrame(), pfn, typ, args, env)
	})
	fn := (*reflectValue)(unsafe.Pointer(&v)).ptr
	interp.makeFuncs.Store((*llgoClosure)(fn).fn, c)
	return v
}
