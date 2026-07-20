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

func init() {
	if funcval.IsSupport {
		RegisterExternal("(reflect.Value).Pointer", func(fr *frame, v reflect.Value) uintptr {
			if v.Kind() == reflect.Func {
				if fv, n := funcval.Get(v.Interface()); n == 1 {
					if c := (*makeFuncVal)(unsafe.Pointer(fv)); c.interp == fr.interp {
						return uintptr(c.pfn.base)
					}
				}
			}
			return v.Pointer()
		})
		RegisterExternal("(reflect.Value).UnsafePointer", func(fr *frame, v reflect.Value) unsafe.Pointer {
			if v.Kind() == reflect.Func {
				if fv, n := funcval.Get(v.Interface()); n == 1 {
					if c := (*makeFuncVal)(unsafe.Pointer(fv)); c.interp == fr.interp {
						return unsafe.Pointer(uintptr(c.pfn.base))
					}
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
