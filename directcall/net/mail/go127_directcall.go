// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package mail

import (
	q "net/mail"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("net/mail", map[string]ixgo.DirectCallAdapter{
		"(*Address).String": method_ptr_Address_String,
		"(*Header).Get":     method_ptr_Header_Get,
		"(Header).Get":      method_Header_Get,
	})
}

func method_ptr_Address_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Address).String(ixgo.DirectCallArg[*q.Address](ctx, 0)))
}

func method_Header_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Header.Get(ixgo.DirectCallArg[q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Header_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Header).Get(ixgo.DirectCallArg[*q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}
