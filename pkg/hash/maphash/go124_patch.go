//go:build go1.24
// +build go1.24

package maphash

import (
	_ "embed"

	"github.com/goplus/ixgo"

	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/abi"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/goarch"
)

//go:embed _patch/maphash.go
var patch_data []byte

func init() {
	ixgo.RegisterPatch("hash/maphash", patch_data)
}
