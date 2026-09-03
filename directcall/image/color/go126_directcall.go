// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package color

import (
	q "image/color"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("image/color", map[string]ixgo.DirectCallAdapter{
		"(*Palette).Convert": method_ptr_Palette_Convert,
		"(*Palette).Index":   method_ptr_Palette_Index,
		"(Palette).Convert":  method_Palette_Convert,
		"(Palette).Index":    method_Palette_Index,
		"ModelFunc":          func_ModelFunc,
	})
}

func func_ModelFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ModelFunc(ixgo.DirectCallArg[func(q.Color) q.Color](ctx, 0)))
}

func method_Palette_Convert(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Palette.Convert(ixgo.DirectCallArg[q.Palette](ctx, 0), ixgo.DirectCallArg[q.Color](ctx, 1)))
}

func method_ptr_Palette_Convert(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Palette).Convert(ixgo.DirectCallArg[*q.Palette](ctx, 0), ixgo.DirectCallArg[q.Color](ctx, 1)))
}

func method_Palette_Index(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Palette.Index(ixgo.DirectCallArg[q.Palette](ctx, 0), ixgo.DirectCallArg[q.Color](ctx, 1)))
}

func method_ptr_Palette_Index(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Palette).Index(ixgo.DirectCallArg[*q.Palette](ctx, 0), ixgo.DirectCallArg[q.Color](ctx, 1)))
}
