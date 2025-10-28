//go:build go1.22
// +build go1.22

package rand

import (
	_ "embed"
	_ "unsafe"

	"math/rand/v2"

	"github.com/goplus/ixgo"
)

//go:embed _patch/rand.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("math/rand/v2.globalRandUint64n", func(n uint64) uint64 {
		return randUint64n(globalRand, n)
	})
	ixgo.RegisterPatch("math/rand/v2", patch_data)
}

//go:linkname globalRand math/rand/v2.globalRand
var globalRand *rand.Rand

//go:linkname randUint64n math/rand/v2.(*Rand).uint64n
func randUint64n(*rand.Rand, uint64) uint64
