//go:build llgo
// +build llgo

package ixgo

import (
	"runtime"
	_ "unsafe"
)

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
