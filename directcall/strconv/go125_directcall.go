// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package strconv

import (
	q "strconv"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("strconv", map[string]ixgo.DirectCallAdapter{
		"(*NumError).Error":        method_ptr_NumError_Error,
		"(*NumError).Unwrap":       method_ptr_NumError_Unwrap,
		"AppendBool":               func_AppendBool,
		"AppendFloat":              func_AppendFloat,
		"AppendInt":                func_AppendInt,
		"AppendQuote":              func_AppendQuote,
		"AppendQuoteRune":          func_AppendQuoteRune,
		"AppendQuoteRuneToASCII":   func_AppendQuoteRuneToASCII,
		"AppendQuoteRuneToGraphic": func_AppendQuoteRuneToGraphic,
		"AppendQuoteToASCII":       func_AppendQuoteToASCII,
		"AppendQuoteToGraphic":     func_AppendQuoteToGraphic,
		"AppendUint":               func_AppendUint,
		"CanBackquote":             func_CanBackquote,
		"FormatBool":               func_FormatBool,
		"FormatComplex":            func_FormatComplex,
		"FormatFloat":              func_FormatFloat,
		"FormatInt":                func_FormatInt,
		"FormatUint":               func_FormatUint,
		"IsGraphic":                func_IsGraphic,
		"IsPrint":                  func_IsPrint,
		"Itoa":                     func_Itoa,
		"Quote":                    func_Quote,
		"QuoteRune":                func_QuoteRune,
		"QuoteRuneToASCII":         func_QuoteRuneToASCII,
		"QuoteRuneToGraphic":       func_QuoteRuneToGraphic,
		"QuoteToASCII":             func_QuoteToASCII,
		"QuoteToGraphic":           func_QuoteToGraphic,
	})
}

func func_AppendBool(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendBool(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1)))
}

func func_AppendFloat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendFloat(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1), ixgo.DirectCallArg[byte](ctx, 2), ixgo.DirectCallArg[int](ctx, 3), ixgo.DirectCallArg[int](ctx, 4)))
}

func func_AppendInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendInt(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_AppendQuote(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendQuote(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_AppendQuoteRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendQuoteRune(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_AppendQuoteRuneToASCII(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendQuoteRuneToASCII(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_AppendQuoteRuneToGraphic(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendQuoteRuneToGraphic(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_AppendQuoteToASCII(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendQuoteToASCII(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_AppendQuoteToGraphic(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendQuoteToGraphic(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_AppendUint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendUint(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_CanBackquote(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CanBackquote(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_FormatBool(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatBool(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_FormatComplex(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatComplex(ixgo.DirectCallArg[complex128](ctx, 0), ixgo.DirectCallArg[byte](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func func_FormatFloat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatFloat(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[byte](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func func_FormatInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatInt(ixgo.DirectCallArg[int64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_FormatUint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatUint(ixgo.DirectCallArg[uint64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_IsGraphic(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsGraphic(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsPrint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsPrint(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_Itoa(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Itoa(ixgo.DirectCallArg[int](ctx, 0)))
}

func method_ptr_NumError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NumError).Error(ixgo.DirectCallArg[*q.NumError](ctx, 0)))
}

func method_ptr_NumError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NumError).Unwrap(ixgo.DirectCallArg[*q.NumError](ctx, 0)))
}

func func_Quote(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Quote(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_QuoteRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.QuoteRune(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_QuoteRuneToASCII(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.QuoteRuneToASCII(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_QuoteRuneToGraphic(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.QuoteRuneToGraphic(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_QuoteToASCII(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.QuoteToASCII(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_QuoteToGraphic(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.QuoteToGraphic(ixgo.DirectCallArg[string](ctx, 0)))
}
