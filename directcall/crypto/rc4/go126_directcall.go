// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package rc4

import (
	q "crypto/rc4"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/rc4", map[string]ixgo.DirectCallAdapter{
		"(*Cipher).Reset":        method_ptr_Cipher_Reset,
		"(*Cipher).XORKeyStream": method_ptr_Cipher_XORKeyStream,
		"(*KeySizeError).Error":  method_ptr_KeySizeError_Error,
		"(KeySizeError).Error":   method_KeySizeError_Error,
	})
}

func method_ptr_Cipher_Reset(ctx ixgo.DirectCallContext) {
	(*q.Cipher).Reset(ixgo.DirectCallArg[*q.Cipher](ctx, 0))
}

func method_ptr_Cipher_XORKeyStream(ctx ixgo.DirectCallContext) {
	(*q.Cipher).XORKeyStream(ixgo.DirectCallArg[*q.Cipher](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2))
}

func method_KeySizeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.KeySizeError.Error(ixgo.DirectCallArg[q.KeySizeError](ctx, 0)))
}

func method_ptr_KeySizeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.KeySizeError).Error(ixgo.DirectCallArg[*q.KeySizeError](ctx, 0)))
}
