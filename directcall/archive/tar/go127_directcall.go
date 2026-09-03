// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package tar

import (
	q "archive/tar"

	"github.com/goplus/ixgo"
	io "io"
	fs "io/fs"
)

func init() {
	ixgo.RegisterDirectCalls("archive/tar", map[string]ixgo.DirectCallAdapter{
		"(*Format).String":      method_ptr_Format_String,
		"(*Header).FileInfo":    method_ptr_Header_FileInfo,
		"(*Writer).AddFS":       method_ptr_Writer_AddFS,
		"(*Writer).Close":       method_ptr_Writer_Close,
		"(*Writer).Flush":       method_ptr_Writer_Flush,
		"(*Writer).WriteHeader": method_ptr_Writer_WriteHeader,
		"(Format).String":       method_Format_String,
		"NewReader":             func_NewReader,
		"NewWriter":             func_NewWriter,
	})
}

func method_Format_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Format.String(ixgo.DirectCallArg[q.Format](ctx, 0)))
}

func method_ptr_Format_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Format).String(ixgo.DirectCallArg[*q.Format](ctx, 0)))
}

func method_ptr_Header_FileInfo(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Header).FileInfo(ixgo.DirectCallArg[*q.Header](ctx, 0)))
}

func func_NewReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriter(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func method_ptr_Writer_AddFS(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).AddFS(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[fs.FS](ctx, 1)))
}

func method_ptr_Writer_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Close(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Flush(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Flush(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_WriteHeader(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).WriteHeader(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[*q.Header](ctx, 1)))
}
