// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package aes

import (
	q "crypto/aes"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/aes", map[string]ixgo.DirectCallAdapter{
		"(*KeySizeError).Error": method_ptr_KeySizeError_Error,
		"(KeySizeError).Error":  method_KeySizeError_Error,
	})
}

func method_KeySizeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.KeySizeError.Error(ixgo.DirectCallArg[q.KeySizeError](ctx, 0)))
}

func method_ptr_KeySizeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.KeySizeError).Error(ixgo.DirectCallArg[*q.KeySizeError](ctx, 0)))
}
