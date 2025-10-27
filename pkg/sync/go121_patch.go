//go:build go1.21
// +build go1.21

package sync

import (
	_ "embed"

	"github.com/goplus/ixgo"
)

//go:embed _patch/oncefunc.go
var patch_data []byte

func init() {
	ixgo.RegisterPatch("sync", patch_data)
}
