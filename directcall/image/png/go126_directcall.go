// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package png

import (
	q "image/png"

	"github.com/goplus/ixgo"
	image "image"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("image/png", map[string]ixgo.DirectCallAdapter{
		"(*Encoder).Encode":         method_ptr_Encoder_Encode,
		"(*FormatError).Error":      method_ptr_FormatError_Error,
		"(*UnsupportedError).Error": method_ptr_UnsupportedError_Error,
		"(FormatError).Error":       method_FormatError_Error,
		"(UnsupportedError).Error":  method_UnsupportedError_Error,
		"Encode":                    func_Encode,
	})
}

func func_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Encode(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[image.Image](ctx, 1)))
}

func method_ptr_Encoder_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).Encode(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[image.Image](ctx, 2)))
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
