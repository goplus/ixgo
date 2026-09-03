// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package csv

import (
	q "encoding/csv"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/csv", map[string]ixgo.DirectCallAdapter{
		"(*ParseError).Error":   method_ptr_ParseError_Error,
		"(*ParseError).Unwrap":  method_ptr_ParseError_Unwrap,
		"(*Reader).InputOffset": method_ptr_Reader_InputOffset,
		"(*Writer).Error":       method_ptr_Writer_Error,
		"(*Writer).Flush":       method_ptr_Writer_Flush,
		"(*Writer).Write":       method_ptr_Writer_Write,
		"(*Writer).WriteAll":    method_ptr_Writer_WriteAll,
		"NewReader":             func_NewReader,
		"NewWriter":             func_NewWriter,
	})
}

func func_NewReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriter(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func method_ptr_ParseError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ParseError).Error(ixgo.DirectCallArg[*q.ParseError](ctx, 0)))
}

func method_ptr_ParseError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ParseError).Unwrap(ixgo.DirectCallArg[*q.ParseError](ctx, 0)))
}

func method_ptr_Reader_InputOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).InputOffset(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Writer_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Error(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Flush(ctx ixgo.DirectCallContext) {
	(*q.Writer).Flush(ixgo.DirectCallArg[*q.Writer](ctx, 0))
}

func method_ptr_Writer_Write(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Write(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[[]string](ctx, 1)))
}

func method_ptr_Writer_WriteAll(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).WriteAll(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[[][]string](ctx, 1)))
}
