// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package io

import (
	q "io"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("io", map[string]ixgo.DirectCallAdapter{
		"(*PipeReader).Close":          method_ptr_PipeReader_Close,
		"(*PipeReader).CloseWithError": method_ptr_PipeReader_CloseWithError,
		"(*PipeWriter).Close":          method_ptr_PipeWriter_Close,
		"(*PipeWriter).CloseWithError": method_ptr_PipeWriter_CloseWithError,
		"(*SectionReader).Size":        method_ptr_SectionReader_Size,
		"LimitReader":                  func_LimitReader,
		"MultiReader":                  func_MultiReader,
		"MultiWriter":                  func_MultiWriter,
		"NewOffsetWriter":              func_NewOffsetWriter,
		"NewSectionReader":             func_NewSectionReader,
		"NopCloser":                    func_NopCloser,
		"TeeReader":                    func_TeeReader,
	})
}

func func_LimitReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LimitReader(ixgo.DirectCallArg[q.Reader](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_MultiReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MultiReader(ixgo.DirectCallArg[[]q.Reader](ctx, 0)...))
}

func func_MultiWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MultiWriter(ixgo.DirectCallArg[[]q.Writer](ctx, 0)...))
}

func func_NewOffsetWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewOffsetWriter(ixgo.DirectCallArg[q.WriterAt](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_NewSectionReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewSectionReader(ixgo.DirectCallArg[q.ReaderAt](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1), ixgo.DirectCallArg[int64](ctx, 2)))
}

func func_NopCloser(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NopCloser(ixgo.DirectCallArg[q.Reader](ctx, 0)))
}

func method_ptr_PipeReader_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PipeReader).Close(ixgo.DirectCallArg[*q.PipeReader](ctx, 0)))
}

func method_ptr_PipeReader_CloseWithError(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PipeReader).CloseWithError(ixgo.DirectCallArg[*q.PipeReader](ctx, 0), ixgo.DirectCallArg[error](ctx, 1)))
}

func method_ptr_PipeWriter_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PipeWriter).Close(ixgo.DirectCallArg[*q.PipeWriter](ctx, 0)))
}

func method_ptr_PipeWriter_CloseWithError(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PipeWriter).CloseWithError(ixgo.DirectCallArg[*q.PipeWriter](ctx, 0), ixgo.DirectCallArg[error](ctx, 1)))
}

func method_ptr_SectionReader_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SectionReader).Size(ixgo.DirectCallArg[*q.SectionReader](ctx, 0)))
}

func func_TeeReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TeeReader(ixgo.DirectCallArg[q.Reader](ctx, 0), ixgo.DirectCallArg[q.Writer](ctx, 1)))
}
