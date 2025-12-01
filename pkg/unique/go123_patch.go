//go:build go1.23 && !go1.25
// +build go1.23,!go1.25

package unique

import (
	_ "embed"
	_ "unique"
	_ "unsafe"

	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/abi"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/goarch"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/isync"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/stringslite"

	"github.com/goplus/ixgo"
)

//go:embed _patch/clone.go
var patch_clone []byte

//go:embed _patch/handle.go
var patch_handle []byte

func init() {
	ixgo.RegisterExternal("unique.runtime_registerUniqueMapCleanup", runtime_registerUniqueMapCleanup)
	ixgo.RegisterPatch("unique", patch_clone, patch_handle)
}

//go:linkname runtime_registerUniqueMapCleanup unique.runtime_registerUniqueMapCleanup
func runtime_registerUniqueMapCleanup(cleanup func())
