//go:build go1.25
// +build go1.25

package reflect

import (
	_ "embed"

	"github.com/goplus/ixgo"
)

//go:embed _patch/type_go125.go
var patch_data_125 []byte

func init() {
	ixgo.RegisterPatch("reflect", patch_data_125)
}
