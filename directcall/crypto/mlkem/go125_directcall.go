// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package mlkem

import (
	q "crypto/mlkem"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/mlkem", map[string]ixgo.DirectCallAdapter{
		"(*DecapsulationKey1024).Bytes":            method_ptr_DecapsulationKey1024_Bytes,
		"(*DecapsulationKey1024).EncapsulationKey": method_ptr_DecapsulationKey1024_EncapsulationKey,
		"(*DecapsulationKey768).Bytes":             method_ptr_DecapsulationKey768_Bytes,
		"(*DecapsulationKey768).EncapsulationKey":  method_ptr_DecapsulationKey768_EncapsulationKey,
		"(*EncapsulationKey1024).Bytes":            method_ptr_EncapsulationKey1024_Bytes,
		"(*EncapsulationKey768).Bytes":             method_ptr_EncapsulationKey768_Bytes,
	})
}

func method_ptr_DecapsulationKey1024_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DecapsulationKey1024).Bytes(ixgo.DirectCallArg[*q.DecapsulationKey1024](ctx, 0)))
}

func method_ptr_DecapsulationKey1024_EncapsulationKey(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DecapsulationKey1024).EncapsulationKey(ixgo.DirectCallArg[*q.DecapsulationKey1024](ctx, 0)))
}

func method_ptr_DecapsulationKey768_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DecapsulationKey768).Bytes(ixgo.DirectCallArg[*q.DecapsulationKey768](ctx, 0)))
}

func method_ptr_DecapsulationKey768_EncapsulationKey(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DecapsulationKey768).EncapsulationKey(ixgo.DirectCallArg[*q.DecapsulationKey768](ctx, 0)))
}

func method_ptr_EncapsulationKey1024_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.EncapsulationKey1024).Bytes(ixgo.DirectCallArg[*q.EncapsulationKey1024](ctx, 0)))
}

func method_ptr_EncapsulationKey768_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.EncapsulationKey768).Bytes(ixgo.DirectCallArg[*q.EncapsulationKey768](ctx, 0)))
}
