// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package gzip

import (
	q "compress/gzip"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("compress/gzip", map[string]ixgo.DirectCallAdapter{
		"(*Reader).Close":       method_ptr_Reader_Close,
		"(*Reader).Multistream": method_ptr_Reader_Multistream,
		"(*Reader).Reset":       method_ptr_Reader_Reset,
		"(*Writer).Close":       method_ptr_Writer_Close,
		"(*Writer).Flush":       method_ptr_Writer_Flush,
		"(*Writer).Reset":       method_ptr_Writer_Reset,
		"NewWriter":             func_NewWriter,
	})
}

func func_NewWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriter(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func method_ptr_Reader_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).Close(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Reader_Multistream(ctx ixgo.DirectCallContext) {
	(*q.Reader).Multistream(ixgo.DirectCallArg[*q.Reader](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_ptr_Reader_Reset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).Reset(ixgo.DirectCallArg[*q.Reader](ctx, 0), ixgo.DirectCallArg[io.Reader](ctx, 1)))
}

func method_ptr_Writer_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Close(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Flush(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Flush(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Reset(ctx ixgo.DirectCallContext) {
	(*q.Writer).Reset(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1))
}
