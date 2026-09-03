// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package pe

import (
	q "debug/pe"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("debug/pe", map[string]ixgo.DirectCallAdapter{
		"(*File).Close":        method_ptr_File_Close,
		"(*File).Section":      method_ptr_File_Section,
		"(*FormatError).Error": method_ptr_FormatError_Error,
		"(*Section).Open":      method_ptr_Section_Open,
	})
}

func method_ptr_File_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Close(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Section(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Section(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_FormatError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FormatError).Error(ixgo.DirectCallArg[*q.FormatError](ctx, 0)))
}

func method_ptr_Section_Open(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Section).Open(ixgo.DirectCallArg[*q.Section](ctx, 0)))
}
