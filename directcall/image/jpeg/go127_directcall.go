// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package jpeg

import (
	q "image/jpeg"

	"github.com/goplus/ixgo"
	image "image"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("image/jpeg", map[string]ixgo.DirectCallAdapter{
		"(*FormatError).Error":      method_ptr_FormatError_Error,
		"(*UnsupportedError).Error": method_ptr_UnsupportedError_Error,
		"(FormatError).Error":       method_FormatError_Error,
		"(UnsupportedError).Error":  method_UnsupportedError_Error,
		"Encode":                    func_Encode,
	})
}

func func_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Encode(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[image.Image](ctx, 1), ixgo.DirectCallArg[*q.Options](ctx, 2)))
}

func method_FormatError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatError.Error(ixgo.DirectCallArg[q.FormatError](ctx, 0)))
}

func method_ptr_FormatError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FormatError).Error(ixgo.DirectCallArg[*q.FormatError](ctx, 0)))
}

func method_UnsupportedError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnsupportedError.Error(ixgo.DirectCallArg[q.UnsupportedError](ctx, 0)))
}

func method_ptr_UnsupportedError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnsupportedError).Error(ixgo.DirectCallArg[*q.UnsupportedError](ctx, 0)))
}
