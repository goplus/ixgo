//go:build go1.22
// +build go1.22

package rand

import (
	_ "embed"
	_ "runtime"
	_ "unsafe"

	"math/rand/v2"

	"github.com/goplus/ixgo"
)

//go:embed _patch/rand.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("math/rand/v2.globalRandUint64n", func(n uint64) uint64 {
		return globalRand.Uint64N(n)
	})
	ixgo.RegisterPatch("math/rand/v2", patch_data)
}

// globalRand is the source of random numbers for the top-level
// convenience functions.
var globalRand = rand.New(runtimeSource{})

//go:linkname runtime_rand runtime.rand
func runtime_rand() uint64

// runtimeSource is a Source that uses the runtime fastrand functions.
type runtimeSource struct{}

func (runtimeSource) Uint64() uint64 {
	return runtime_rand()
}
