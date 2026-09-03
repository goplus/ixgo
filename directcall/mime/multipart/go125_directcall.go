// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package multipart

import (
	q "mime/multipart"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("mime/multipart", map[string]ixgo.DirectCallAdapter{
		"(*Form).RemoveAll":             method_ptr_Form_RemoveAll,
		"(*Part).Close":                 method_ptr_Part_Close,
		"(*Part).FileName":              method_ptr_Part_FileName,
		"(*Part).FormName":              method_ptr_Part_FormName,
		"(*Writer).Boundary":            method_ptr_Writer_Boundary,
		"(*Writer).Close":               method_ptr_Writer_Close,
		"(*Writer).FormDataContentType": method_ptr_Writer_FormDataContentType,
		"(*Writer).SetBoundary":         method_ptr_Writer_SetBoundary,
		"(*Writer).WriteField":          method_ptr_Writer_WriteField,
		"FileContentDisposition":        func_FileContentDisposition,
		"NewReader":                     func_NewReader,
		"NewWriter":                     func_NewWriter,
	})
}

func func_FileContentDisposition(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FileContentDisposition(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Form_RemoveAll(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Form).RemoveAll(ixgo.DirectCallArg[*q.Form](ctx, 0)))
}

func func_NewReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReader(ixgo.DirectCallArg[io.Reader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_NewWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriter(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func method_ptr_Part_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Part).Close(ixgo.DirectCallArg[*q.Part](ctx, 0)))
}

func method_ptr_Part_FileName(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Part).FileName(ixgo.DirectCallArg[*q.Part](ctx, 0)))
}

func method_ptr_Part_FormName(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Part).FormName(ixgo.DirectCallArg[*q.Part](ctx, 0)))
}

func method_ptr_Writer_Boundary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Boundary(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Close(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_FormDataContentType(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).FormDataContentType(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_SetBoundary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).SetBoundary(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Writer_WriteField(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).WriteField(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}
