//go:build linknamefix
// +build linknamefix

package iter

import (
	_ "embed"
	_ "runtime"
	_ "unsafe"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/race"
)

//go:embed _patch/iter_linkname.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("runtime.newcoro", newcoro)
	ixgo.RegisterExternal("runtime.coroswitch", coroswitch)
	ixgo.RegisterPatch("iter", patch_data)
}

type coro = struct{}

//go:linkname newcoro runtime.newcoro
func newcoro(fn func(*coro)) *coro

//go:linkname coroswitch runtime.coroswitch
func coroswitch(*coro)
