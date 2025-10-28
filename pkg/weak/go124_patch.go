//go:build go1.24
// +build go1.24

package weak

import (
	_ "embed"
	_ "weak"

	"unsafe"

	"github.com/goplus/ixgo"
)

//go:embed _patch/weak.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("weak.runtime_registerWeakPointer", runtime_registerWeakPointer)
	ixgo.RegisterExternal("weak.runtime_makeStrongFromWeak", runtime_makeStrongFromWeak)
	ixgo.RegisterPatch("weak", patch_data)
}

//go:linkname runtime_registerWeakPointer weak.runtime_registerWeakPointer
func runtime_registerWeakPointer(unsafe.Pointer) unsafe.Pointer

//go:linkname runtime_makeStrongFromWeak weak.runtime_makeStrongFromWeak
func runtime_makeStrongFromWeak(unsafe.Pointer) unsafe.Pointer
