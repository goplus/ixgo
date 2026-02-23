//go:build !linknamefix
// +build !linknamefix

package iter

import (
	_ "embed"
	"iter"
	_ "unsafe"

	"github.com/goplus/ixgo"
)

//go:embed _patch/iter.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("iter.pullany", pullany)
	ixgo.RegisterExternal("iter.pullany2", pullany2)
	ixgo.RegisterPatch("iter", patch_data)
}

func pullany(seq func(yield func(v any) bool)) (next func() (any, bool), stop func()) {
	return iter.Pull(seq)
}

func pullany2(seq func(yield func(k, v any) bool)) (next func() (any, any, bool), stop func()) {
	return iter.Pull2(seq)
}
