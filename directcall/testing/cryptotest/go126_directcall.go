// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package cryptotest

import (
	q "testing/cryptotest"

	"github.com/goplus/ixgo"
	testing "testing"
)

func init() {
	ixgo.RegisterDirectCalls("testing/cryptotest", map[string]ixgo.DirectCallAdapter{
		"SetGlobalRandom": func_SetGlobalRandom,
	})
}

func func_SetGlobalRandom(ctx ixgo.DirectCallContext) {
	q.SetGlobalRandom(ixgo.DirectCallArg[*testing.T](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1))
}
