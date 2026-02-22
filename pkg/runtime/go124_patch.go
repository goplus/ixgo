//go:build go1.24 && !linknamefix
// +build go1.24,!linknamefix

package runtime

import (
	_ "embed"
	_ "unsafe"

	"github.com/goplus/ixgo"
)

//go:embed _patch/mcleanup_empty.go
var patch_data []byte

func init() {
	ixgo.RegisterPatch("runtime", patch_data)
}
