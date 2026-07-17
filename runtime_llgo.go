//go:build llgo
// +build llgo

package ixgo

import (
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sync/atomic"
)

func runtimeGC(fr *frame) {
	runtime.GC()
}

func debugStack(fr *frame) []byte {
	return debug.Stack()
}

func init() {
	RegisterExternal("os.Exit", func(fr *frame, code int) {
		interp := fr.interp
		if atomic.LoadInt32(&interp.goexited) == 1 {
			//os.Exit(code)
			interp.chexit <- code
		} else {
			panic(exitPanic(code))
		}
	})
	RegisterExternal("runtime.Goexit", func(fr *frame) {
		interp := fr.interp
		// main goroutine use panic
		if goroutineID() == interp.mainid {
			atomic.StoreInt32(&interp.goexited, 1)
			panic(goexitPanic(0))
		} else {
			runtime.Goexit()
		}
	})
	RegisterExternal("runtime.Caller", runtimeCaller)
	RegisterExternal("runtime.FuncForPC", runtimeFuncForPC)
}

func runtimeCaller(fr *frame, skip int) (pc uintptr, file string, line int, ok bool) {
	pc = fr.pc()
	if pfn := findFuncByPC(fr.interp, int(pc)); pfn != nil {
		pos := pfn.PosForPC(int(pc) - pfn.base - 1)
		if !pos.IsValid() {
			return
		}
		fpos := fr.interp.ctx.FileSet.Position(pos)
		if fpos.Filename == "" {
			return
		}
		file, line, ok = filepath.ToSlash(fpos.Filename), fpos.Line, true
	}
	return
}

func runtimeFuncForPC(fr *frame, pc uintptr) *runtime.Func {
	if pfn := findFuncByPC(fr.interp, int(pc)); pfn != nil {
		return runtimeFunc(pfn)
	}
	return runtime.FuncForPC(pc)
}

type funcinl struct {
	entry uintptr
	name  string
	pc    uintptr
	file  string
	line  int
}

func inlineFunc(entry uintptr) *funcinl {
	return &funcinl{entry: entry}
}
