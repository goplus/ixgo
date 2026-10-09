//go:build !llgo
// +build !llgo

package ixgo

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/visualfc/funcval"
)

const (
	Compiler = runtime.Compiler
	IsLLGo   = false
)

func init() {
	if funcval.IsSupport {
		RegisterExternal("(reflect.Value).Pointer", func(fr *frame, v reflect.Value) uintptr {
			if v.Kind() == reflect.Func {
				if c := fr.interp.getMakeFuncVal(v.Interface()); c != nil && c.interp == fr.interp {
					return uintptr(c.pfn.base)
				}
			}
			return v.Pointer()
		})
		RegisterExternal("(reflect.Value).UnsafePointer", func(fr *frame, v reflect.Value) unsafe.Pointer {
			if v.Kind() == reflect.Func {
				if c := fr.interp.getMakeFuncVal(v.Interface()); c != nil && c.interp == fr.interp {
					return unsafe.Pointer(uintptr(c.pfn.base))
				}
			}
			return v.UnsafePointer()
		})
	}
}

//go:linkname runtimePanic runtime.gopanic
func runtimePanic(e interface{})

type funcinl struct {
	ones  uint32  // set to ^0 to distinguish from _func
	entry uintptr // entry of the real (the "outermost") frame
	name  string
	file  string
	line  int
}

func inlineFunc(entry uintptr) *funcinl {
	return &funcinl{ones: ^uint32(0), entry: entry}
}

func isInlineFunc(f *runtime.Func) bool {
	return (*funcinl)(unsafe.Pointer(f)).ones == ^uint32(0)
}

func runtimeFuncFileLine(fr *frame, f *runtime.Func, pc uintptr) (file string, line int) {
	entry := f.Entry()
	if isInlineFunc(f) && pc > entry {
		interp := fr.interp
		if pfn := findFuncByEntry(interp, int(entry)); pfn != nil {
			// pc-1 : fn.instr.pos
			pos := pfn.PosForPC(int(pc - entry - 1))
			if !pos.IsValid() {
				return "?", 0
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

const supportFuncVal = funcval.IsSupport

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

func (pfn *function) makeFunction(typ reflect.Type, env []value) reflect.Value {
	c := DirectFuncVal{interp: pfn.Interp, pfn: pfn, typ: typ, env: env}
	if supportFuncVal {
		if v, ok := makeDirectFunction(c, typ); ok {
			return v
		}
	}
	return reflect.MakeFunc(typ, c.callReflect)
}

var callbackReflectPC = reflect.ValueOf(DirectFuncVal{}.callReflect).Pointer()

type interpExt struct{}

// Callers must check interpreter ownership.
func (*interpExt) getMakeFuncVal(fn interface{}) *DirectFuncVal {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func || v.IsNil() {
		return nil
	}
	fv, n := funcval.Get(fn)
	switch {
	case n == 0 && (fv.Fn == directFuncVoidPC || isDirectFuncPC(fv.Fn)):
		// Direct method value.
	case n == 1 && fv.Fn == callbackReflectPC:
		// One reflect.MakeFunc bridge.
	default:
		return nil
	}
	return directFuncReceiver(fv)
}

func validateDirectFunc(typ reflect.Type, maker DirectFuncMaker) uintptr {
	sentinel := new(Interp)
	fn := new(function)
	env := []value{sentinel}
	v := maker(DirectFuncVal{interp: sentinel, pfn: fn, typ: typ, env: env})
	if !v.IsValid() || v.Kind() != reflect.Func || v.Type() != typ {
		panic(fmt.Sprintf("ixgo: DirectFunc maker returned %v, want %v", v.Type(), typ))
	}
	if v.IsNil() {
		panic("ixgo: DirectFunc maker returned a nil function")
	}
	fv, n := funcval.Get(v.Interface())
	if n != 0 {
		panic("ixgo: DirectFunc maker must return a method value")
	}
	got := directFuncReceiver(fv)
	if got == nil || got.interp != sentinel || got.pfn != fn || got.typ != typ || len(got.env) != 1 || got.env[0] != sentinel {
		panic("ixgo: DirectFunc receiver must embed ixgo.DirectFuncVal as its first field")
	}
	return fv.Fn
}

func directFuncReceiver(fv *funcval.FuncVal) *DirectFuncVal {
	if fv == nil {
		return nil
	}
	// gc ABI: DirectFuncVal is one word, followed by the method-value receiver.
	// The receiver must be at least as large as DirectFuncVal and must embed it
	// as its first field. Undersized receivers are rejected at registration
	// only after this read, so makers must return a real method value.
	return &(*struct {
		funcval.FuncVal
		receiver DirectFuncVal
	})(unsafe.Pointer(fv)).receiver
}
