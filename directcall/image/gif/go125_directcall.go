// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package gif

import (
	q "image/gif"

	"github.com/goplus/ixgo"
	image "image"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("image/gif", map[string]ixgo.DirectCallAdapter{
		"Encode":    func_Encode,
		"EncodeAll": func_EncodeAll,
	})
}

func func_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Encode(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[image.Image](ctx, 1), ixgo.DirectCallArg[*q.Options](ctx, 2)))
}

func func_EncodeAll(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EncodeAll(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[*q.GIF](ctx, 1)))
}
