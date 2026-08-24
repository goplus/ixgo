//go:build go1.24
// +build go1.24

package hkdf

import (
	_ "embed"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/errors"
	_ "github.com/goplus/ixgo/pkg/hash"
)

//go:embed _patch/hkdf.go
var patch_data []byte

func init() {
	ixgo.RegisterPatch("crypto/hkdf", patch_data)
}
