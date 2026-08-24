//go:build go1.27
// +build go1.27

package rand

import (
	_ "embed"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/math/bits"
)

//go:embed _patch/rand_go127.go
var patch_data_go127 []byte

func init() {
	ixgo.RegisterPatch("math/rand/v2", patch_data_go127)
}
