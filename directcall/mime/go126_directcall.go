// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package mime

import (
	q "mime"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("mime", map[string]ixgo.DirectCallAdapter{
		"(*WordEncoder).Encode": method_ptr_WordEncoder_Encode,
		"(WordEncoder).Encode":  method_WordEncoder_Encode,
		"AddExtensionType":      func_AddExtensionType,
		"FormatMediaType":       func_FormatMediaType,
		"TypeByExtension":       func_TypeByExtension,
	})
}

func func_AddExtensionType(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddExtensionType(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_FormatMediaType(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatMediaType(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[map[string]string](ctx, 1)))
}

func func_TypeByExtension(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeByExtension(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_WordEncoder_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WordEncoder.Encode(ixgo.DirectCallArg[q.WordEncoder](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_WordEncoder_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.WordEncoder).Encode(ixgo.DirectCallArg[*q.WordEncoder](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}
