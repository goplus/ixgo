package abi

import (
	_ "embed"

	"github.com/goplus/ixgo"
)

//go:embed _patch/abi.go
var patch_data []byte

func init() {
	ixgo.RegisterPatch("github.com/goplus/ixgo/x/abi", patch_data)
}
