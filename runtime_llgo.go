//go:build llgo
// +build llgo

package ixgo

import "runtime"

func runtimeGC(fr *frame) {
	runtime.GC()
}

func debugStack(fr *frame) []byte {
	return nil
}
