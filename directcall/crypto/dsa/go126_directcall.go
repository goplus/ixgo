// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package dsa

import (
	q "crypto/dsa"

	"github.com/goplus/ixgo"
	io "io"
	big "math/big"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/dsa", map[string]ixgo.DirectCallAdapter{
		"GenerateKey":        func_GenerateKey,
		"GenerateParameters": func_GenerateParameters,
		"Verify":             func_Verify,
	})
}

func func_GenerateKey(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.GenerateKey(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0), ixgo.DirectCallArg[io.Reader](ctx, 1)))
}

func func_GenerateParameters(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.GenerateParameters(ixgo.DirectCallArg[*q.Parameters](ctx, 0), ixgo.DirectCallArg[io.Reader](ctx, 1), ixgo.DirectCallArg[q.ParameterSizes](ctx, 2)))
}

func func_Verify(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Verify(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[*big.Int](ctx, 2), ixgo.DirectCallArg[*big.Int](ctx, 3)))
}
