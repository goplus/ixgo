// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package lzw

import (
	q "compress/lzw"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("compress/lzw", map[string]ixgo.DirectCallAdapter{
		"(*Reader).Close": method_ptr_Reader_Close,
		"(*Reader).Reset": method_ptr_Reader_Reset,
		"(*Writer).Close": method_ptr_Writer_Close,
		"(*Writer).Reset": method_ptr_Writer_Reset,
		"NewReader":       func_NewReader,
		"NewWriter":       func_NewWriter,
	})
}

func func_NewReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReader(ixgo.DirectCallArg[io.Reader](ctx, 0), ixgo.DirectCallArg[q.Order](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_NewWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriter(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[q.Order](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Reader_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).Close(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Reader_Reset(ctx ixgo.DirectCallContext) {
	(*q.Reader).Reset(ixgo.DirectCallArg[*q.Reader](ctx, 0), ixgo.DirectCallArg[io.Reader](ctx, 1), ixgo.DirectCallArg[q.Order](ctx, 2), ixgo.DirectCallArg[int](ctx, 3))
}

func method_ptr_Writer_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Close(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Reset(ctx ixgo.DirectCallContext) {
	(*q.Writer).Reset(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[q.Order](ctx, 2), ixgo.DirectCallArg[int](ctx, 3))
}
