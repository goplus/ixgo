// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package utf8

import (
	q "unicode/utf8"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("unicode/utf8", map[string]ixgo.DirectCallAdapter{
		"AppendRune":        func_AppendRune,
		"EncodeRune":        func_EncodeRune,
		"FullRune":          func_FullRune,
		"FullRuneInString":  func_FullRuneInString,
		"RuneCount":         func_RuneCount,
		"RuneCountInString": func_RuneCountInString,
		"RuneLen":           func_RuneLen,
		"RuneStart":         func_RuneStart,
		"Valid":             func_Valid,
		"ValidRune":         func_ValidRune,
		"ValidString":       func_ValidString,
	})
}

func func_AppendRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendRune(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_EncodeRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EncodeRune(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_FullRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FullRune(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_FullRuneInString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FullRuneInString(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_RuneCount(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RuneCount(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_RuneCountInString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RuneCountInString(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_RuneLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RuneLen(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_RuneStart(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RuneStart(ixgo.DirectCallArg[byte](ctx, 0)))
}

func func_Valid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Valid(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_ValidRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ValidRune(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_ValidString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ValidString(ixgo.DirectCallArg[string](ctx, 0)))
}
