// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package utf16

import (
	q "unicode/utf16"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("unicode/utf16", map[string]ixgo.DirectCallAdapter{
		"AppendRune":  func_AppendRune,
		"Decode":      func_Decode,
		"DecodeRune":  func_DecodeRune,
		"Encode":      func_Encode,
		"IsSurrogate": func_IsSurrogate,
		"RuneLen":     func_RuneLen,
	})
}

func func_AppendRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendRune(ixgo.DirectCallArg[[]uint16](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_Decode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Decode(ixgo.DirectCallArg[[]uint16](ctx, 0)))
}

func func_DecodeRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DecodeRune(ixgo.DirectCallArg[rune](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Encode(ixgo.DirectCallArg[[]rune](ctx, 0)))
}

func func_IsSurrogate(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsSurrogate(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_RuneLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RuneLen(ixgo.DirectCallArg[rune](ctx, 0)))
}
