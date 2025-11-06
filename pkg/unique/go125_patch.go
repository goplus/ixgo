//go:build go1.25
// +build go1.25

package unique

import (
	_ "embed"
	_ "unique"
	_ "unsafe"

	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/abi"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/isync"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/stringslite"

	"github.com/goplus/ixgo"
)

//go:embed _patch125/clone.go
var patch_clone []byte

//go:embed _patch125/handle.go
var patch_handle []byte

//go:embed _patch125/canonmap.go
var patch_canonmap []byte

func init() {
	ixgo.RegisterExternal("unique@patch.runtime_rand", rand)
	ixgo.RegisterPatch("unique", patch_clone, patch_handle, patch_canonmap)
}

//go:linkname rand runtime.rand
func rand()
