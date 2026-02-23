package runtime

import (
	"runtime"
	"unsafe"
)

func AddCleanup[T, S any](ptr *T, cleanup func(S), arg S) (r runtime.Cleanup) {
	return addCleanup(unsafe.Pointer(ptr), func() { cleanup(arg) })
}

//go:linkname addCleanup runtime.addCleanup
func addCleanup(p unsafe.Pointer, fn func()) runtime.Cleanup
