//go:build !linknamefix
// +build !linknamefix

package iter

import (
	_ "embed"
	_ "unsafe"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/race"
)

//go:embed _patch/iter.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("runtime.newcoro", newcoro)
	ixgo.RegisterExternal("runtime.coroswitch", coroswitch)
	ixgo.RegisterPatch("iter", patch_data)
}

type coro = struct{}

func newcoro(fn func(*coro)) *coro {
	return &coro{}
}

func coroswitch(*coro) {
	println(`ixgo warning: iter.coroswitch not impl. need build -tags=linknamefix -ldflags="-checklinkname=0"`)
}
