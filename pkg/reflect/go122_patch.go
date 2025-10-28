//go:build go1.22
// +build go1.22

package reflect

import (
	_ "embed"

	"github.com/goplus/ixgo"
)

//go:embed _patch/type.go
var patch_data []byte

func init() {
	ixgo.RegisterPatch("reflect", patch_data)
}
