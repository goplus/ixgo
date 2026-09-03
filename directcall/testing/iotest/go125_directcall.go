// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package iotest

import (
	q "testing/iotest"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("testing/iotest", map[string]ixgo.DirectCallAdapter{
		"DataErrReader":  func_DataErrReader,
		"ErrReader":      func_ErrReader,
		"HalfReader":     func_HalfReader,
		"NewReadLogger":  func_NewReadLogger,
		"NewWriteLogger": func_NewWriteLogger,
		"OneByteReader":  func_OneByteReader,
		"TestReader":     func_TestReader,
		"TimeoutReader":  func_TimeoutReader,
		"TruncateWriter": func_TruncateWriter,
	})
}

func func_DataErrReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DataErrReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_ErrReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ErrReader(ixgo.DirectCallArg[error](ctx, 0)))
}

func func_HalfReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HalfReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewReadLogger(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReadLogger(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[io.Reader](ctx, 1)))
}

func func_NewWriteLogger(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriteLogger(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1)))
}

func func_OneByteReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OneByteReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_TestReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TestReader(ixgo.DirectCallArg[io.Reader](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_TimeoutReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TimeoutReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_TruncateWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TruncateWriter(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}
