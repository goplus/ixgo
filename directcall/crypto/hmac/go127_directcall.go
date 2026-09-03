// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package hmac

import (
	q "crypto/hmac"

	"github.com/goplus/ixgo"
	hash "hash"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/hmac", map[string]ixgo.DirectCallAdapter{
		"Equal": func_Equal,
		"New":   func_New,
	})
}

func func_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Equal(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[func() hash.Hash](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}
