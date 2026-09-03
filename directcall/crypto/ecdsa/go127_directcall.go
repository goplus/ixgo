// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package ecdsa

import (
	q "crypto/ecdsa"

	crypto "crypto"
	"github.com/goplus/ixgo"
	big "math/big"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/ecdsa", map[string]ixgo.DirectCallAdapter{
		"(*PrivateKey).Equal":  method_ptr_PrivateKey_Equal,
		"(*PrivateKey).Public": method_ptr_PrivateKey_Public,
		"(*PublicKey).Equal":   method_ptr_PublicKey_Equal,
		"Verify":               func_Verify,
		"VerifyASN1":           func_VerifyASN1,
	})
}

func method_ptr_PrivateKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Equal(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0), ixgo.DirectCallArg[crypto.PrivateKey](ctx, 1)))
}

func method_ptr_PrivateKey_Public(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Public(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0)))
}

func method_ptr_PublicKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PublicKey).Equal(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[crypto.PublicKey](ctx, 1)))
}

func func_Verify(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Verify(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[*big.Int](ctx, 2), ixgo.DirectCallArg[*big.Int](ctx, 3)))
}

func func_VerifyASN1(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.VerifyASN1(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}
