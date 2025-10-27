//go:build go1.19
// +build go1.19

package atomic

import (
	_ "embed"

	"github.com/goplus/ixgo"
)

//go:embed _patch/type.go
var patch_data []byte

func init() {
	ixgo.RegisterPatch("sync/atomic", patch_data)
}
