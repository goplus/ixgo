//go:build !llgo
// +build !llgo

package ixgo

import (
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
	c := makeFuncVal{interp: pfn.Interp, pfn: pfn, typ: typ, env: env}
	if supportFuncVal {
		switch typ {
		case callbackVoidType:
			return reflect.ValueOf(c.callVoid)
		case callbackBoolType:
			return reflect.ValueOf(c.callBool)
		}
	}
	return reflect.MakeFunc(typ, c.callReflect)
}

type makeFuncVal struct {
	interp *Interp
	pfn    *function
	typ    reflect.Type
	env    []value
}

var (
	callbackVoidType  = reflect.TypeFor[func()]()
	callbackBoolType  = reflect.TypeFor[func() bool]()
	callbackVoidPC    = reflect.ValueOf(makeFuncVal{}.callVoid).Pointer()
	callbackBoolPC    = reflect.ValueOf(makeFuncVal{}.callBool).Pointer()
	callbackReflectPC = reflect.ValueOf(makeFuncVal{}.callReflect).Pointer()
)

func (c makeFuncVal) callVoid() {
	c.interp.callFunctionDiscardsResult(c.interp.tryDeferFrame(), c.pfn, nil, c.env)
}

func (c makeFuncVal) callBool() bool {
	result := c.interp.callFunction(c.interp.tryDeferFrame(), c.pfn, nil, c.env)
	return result != nil && result.(bool)
}

func (c makeFuncVal) callReflect(args []reflect.Value) []reflect.Value {
	return c.interp.callFunctionByReflect(c.interp.tryDeferFrame(), c.pfn, c.typ, args, c.env)
}

type interpExt struct{}

// Callers must check interpreter ownership.
func (*interpExt) getMakeFuncVal(fn interface{}) *makeFuncVal {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func || v.IsNil() {
		return nil
	}
	fv, n := funcval.Get(fn)
	switch {
	case n == 0 && (fv.Fn == callbackVoidPC || fv.Fn == callbackBoolPC):
		// Direct method value.
	case n == 1 && fv.Fn == callbackReflectPC:
		// One reflect.MakeFunc bridge.
	default:
		return nil
	}
	// gc ABI: FuncVal is one word, followed by the receiver.
	// The call methods must retain value receivers.
	return &(*struct {
		funcval.FuncVal
		receiver makeFuncVal
	})(unsafe.Pointer(fv)).receiver
}
