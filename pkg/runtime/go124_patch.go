//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package runtime

import (
	_ "embed"
	_ "runtime"
	"unsafe"

	"github.com/goplus/ixgo"
)

//go:embed _patch/mcleanup_go124.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("runtime.inUserArenaChunk", inUserArenaChunk)
	ixgo.RegisterExternal("runtime.findObject", findObject)
	ixgo.RegisterExternal("runtime.createfing", createfing)
	ixgo.RegisterExternal("runtime.addCleanup", addCleanup)
	ixgo.RegisterExternal("runtime.isGoPointerWithoutSpan", isGoPointerWithoutSpan)
	ixgo.RegisterPatch("runtime", patch_data)
}

// inUserArenaChunk returns true if p points to a user arena chunk.
//
//go:linkname inUserArenaChunk runtime.inUserArenaChunk
func inUserArenaChunk(p uintptr) bool

//go:linkname findObject runtime.findObject
//go:nosplit
func findObject(p, refBase, refOff uintptr) (base uintptr, s unsafe.Pointer, objIndex uintptr)

//go:linkname createfing runtime.createfing
func createfing()

// addCleanup attaches a cleanup function to the object. Multiple
// cleanups are allowed on an object, and even the same pointer.
// A cleanup id is returned which can be used to uniquely identify
// the cleanup.
//
//go:linkname addCleanup runtime.addCleanup
func addCleanup(p unsafe.Pointer, f *funcval) uint64

//go:linkname isGoPointerWithoutSpan runtime.isGoPointerWithoutSpan
func isGoPointerWithoutSpan(p unsafe.Pointer) bool

type funcval struct {
	fn uintptr
	// variable-size, fn-specific data here
}
