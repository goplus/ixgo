//go:build llgo
// +build llgo

package ixgo

import (
	"runtime"
	"sync/atomic"
)

func runtimeGC(fr *frame) {
	runtime.GC()
}

func debugStack(fr *frame) []byte {
	return nil
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
}
