//go:build go1.24
// +build go1.24

package pbkdf2

import (
	_ "embed"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/errors"
	_ "github.com/goplus/ixgo/pkg/hash"
)

//go:embed _patch/pbkdf2.go
var patch_data []byte

func init() {
	ixgo.RegisterPatch("crypto/pbkdf2", patch_data)
}
