// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package tabwriter

import (
	q "text/tabwriter"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("text/tabwriter", map[string]ixgo.DirectCallAdapter{
		"(*Writer).Flush": method_ptr_Writer_Flush,
		"(*Writer).Init":  method_ptr_Writer_Init,
		"NewWriter":       func_NewWriter,
	})
}

func func_NewWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriter(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3), ixgo.DirectCallArg[byte](ctx, 4), ixgo.DirectCallArg[uint](ctx, 5)))
}

func method_ptr_Writer_Flush(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Flush(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Init(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Init(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3), ixgo.DirectCallArg[int](ctx, 4), ixgo.DirectCallArg[byte](ctx, 5), ixgo.DirectCallArg[uint](ctx, 6)))
}
