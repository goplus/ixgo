// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package draw

import (
	q "image/draw"

	"github.com/goplus/ixgo"
	image "image"
)

func init() {
	ixgo.RegisterDirectCalls("image/draw", map[string]ixgo.DirectCallAdapter{
		"(*Op).Draw": method_ptr_Op_Draw,
		"(Op).Draw":  method_Op_Draw,
		"Draw":       func_Draw,
		"DrawMask":   func_DrawMask,
	})
}

func func_Draw(ctx ixgo.DirectCallContext) {
	q.Draw(ixgo.DirectCallArg[q.Image](ctx, 0), ixgo.DirectCallArg[image.Rectangle](ctx, 1), ixgo.DirectCallArg[image.Image](ctx, 2), ixgo.DirectCallArg[image.Point](ctx, 3), ixgo.DirectCallArg[q.Op](ctx, 4))
}

func func_DrawMask(ctx ixgo.DirectCallContext) {
	q.DrawMask(ixgo.DirectCallArg[q.Image](ctx, 0), ixgo.DirectCallArg[image.Rectangle](ctx, 1), ixgo.DirectCallArg[image.Image](ctx, 2), ixgo.DirectCallArg[image.Point](ctx, 3), ixgo.DirectCallArg[image.Image](ctx, 4), ixgo.DirectCallArg[image.Point](ctx, 5), ixgo.DirectCallArg[q.Op](ctx, 6))
}

func method_Op_Draw(ctx ixgo.DirectCallContext) {
	q.Op.Draw(ixgo.DirectCallArg[q.Op](ctx, 0), ixgo.DirectCallArg[q.Image](ctx, 1), ixgo.DirectCallArg[image.Rectangle](ctx, 2), ixgo.DirectCallArg[image.Image](ctx, 3), ixgo.DirectCallArg[image.Point](ctx, 4))
}

func method_ptr_Op_Draw(ctx ixgo.DirectCallContext) {
	(*q.Op).Draw(ixgo.DirectCallArg[*q.Op](ctx, 0), ixgo.DirectCallArg[q.Image](ctx, 1), ixgo.DirectCallArg[image.Rectangle](ctx, 2), ixgo.DirectCallArg[image.Image](ctx, 3), ixgo.DirectCallArg[image.Point](ctx, 4))
}
