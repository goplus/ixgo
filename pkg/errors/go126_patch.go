//go:build go1.26
// +build go1.26

package errors

import (
	_ "embed"

	"github.com/goplus/ixgo"
)

//go:embed _patch/errors.go
var patch_data []byte

func init() {
	ixgo.RegisterPatch("errors", patch_data)
}
