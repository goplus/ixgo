// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package flate

import (
	q "compress/flate"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("compress/flate", map[string]ixgo.DirectCallAdapter{
		"(*CorruptInputError).Error": method_ptr_CorruptInputError_Error,
		"(*InternalError).Error":     method_ptr_InternalError_Error,
		"(*ReadError).Error":         method_ptr_ReadError_Error,
		"(*WriteError).Error":        method_ptr_WriteError_Error,
		"(*Writer).Close":            method_ptr_Writer_Close,
		"(*Writer).Flush":            method_ptr_Writer_Flush,
		"(*Writer).Reset":            method_ptr_Writer_Reset,
		"(CorruptInputError).Error":  method_CorruptInputError_Error,
		"(InternalError).Error":      method_InternalError_Error,
		"NewReader":                  func_NewReader,
		"NewReaderDict":              func_NewReaderDict,
	})
}

func method_CorruptInputError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CorruptInputError.Error(ixgo.DirectCallArg[q.CorruptInputError](ctx, 0)))
}

func method_ptr_CorruptInputError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CorruptInputError).Error(ixgo.DirectCallArg[*q.CorruptInputError](ctx, 0)))
}

func method_InternalError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.InternalError.Error(ixgo.DirectCallArg[q.InternalError](ctx, 0)))
}

func method_ptr_InternalError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InternalError).Error(ixgo.DirectCallArg[*q.InternalError](ctx, 0)))
}

func func_NewReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewReaderDict(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReaderDict(ixgo.DirectCallArg[io.Reader](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_ReadError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ReadError).Error(ixgo.DirectCallArg[*q.ReadError](ctx, 0)))
}

func method_ptr_WriteError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.WriteError).Error(ixgo.DirectCallArg[*q.WriteError](ctx, 0)))
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
