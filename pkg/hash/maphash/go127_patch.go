//go:build go1.27
// +build go1.27

package maphash

import (
	_ "embed"

	"github.com/goplus/ixgo"
)

//go:embed _patch/maphash_go127.go
var patch_data_go127 []byte

func init() {
	ixgo.RegisterPatch("hash/maphash", patch_data_go127)
}
