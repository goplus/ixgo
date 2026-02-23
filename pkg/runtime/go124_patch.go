//go:build go1.24
// +build go1.24

package runtime

import (
	_ "embed"
	"runtime"
	"unsafe"

	"github.com/goplus/ixgo"
)

//go:embed _patch/mcleanup.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("runtime.addCleanup", addCleanup)
	ixgo.RegisterPatch("runtime", patch_data)
}

func addCleanup(p unsafe.Pointer, fn func()) runtime.Cleanup {
	return runtime.AddCleanup[uintptr, int]((*uintptr)(p), func(int) { fn() }, 0)
}
