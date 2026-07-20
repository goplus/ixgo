//go:build llgo
// +build llgo

package ixgo

import (
	"go/token"
	"path/filepath"
	"reflect"
	"runtime"
	"unsafe"
)

func getMakeFuncClosure(v reflect.Value) *makeFuncClosure {
	pv := (*reflectValue)(unsafe.Pointer(&v))
	fn := (*llgoClosure)(pv.ptr)
	return (*makeFuncClosure)(fn.env)
}

type llgoClosure struct {
	fn  unsafe.Pointer
	env unsafe.Pointer
}

type makeFuncClosure struct {
	fn  unsafe.Pointer
	env *struct {
		interp **Interp
		pfn    **function
		typ    *reflect.Type
		env    *[]interface{}
	}
}

func init() {
	RegisterExternal("(reflect.Value).Pointer", func(fr *frame, v reflect.Value) uintptr {
		if v.Kind() == reflect.Func {
			if c := getMakeFuncClosure(v); c != nil && c.env != nil && *c.env.interp == fr.interp {
				return uintptr((*c.env.pfn).base)
			}
		}
		return v.Pointer()
	})
	RegisterExternal("(reflect.Value).UnsafePointer", func(fr *frame, v reflect.Value) unsafe.Pointer {
		if v.Kind() == reflect.Func {
			if c := getMakeFuncClosure(v); c != nil && c.env != nil && *c.env.interp == fr.interp {
				return unsafe.Pointer(uintptr((*c.env.pfn).base))
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
