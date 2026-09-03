// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package unicode

import (
	q "unicode"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("unicode", map[string]ixgo.DirectCallAdapter{
		"(*SpecialCase).ToLower": method_ptr_SpecialCase_ToLower,
		"(*SpecialCase).ToTitle": method_ptr_SpecialCase_ToTitle,
		"(*SpecialCase).ToUpper": method_ptr_SpecialCase_ToUpper,
		"(SpecialCase).ToLower":  method_SpecialCase_ToLower,
		"(SpecialCase).ToTitle":  method_SpecialCase_ToTitle,
		"(SpecialCase).ToUpper":  method_SpecialCase_ToUpper,
		"In":                     func_In,
		"Is":                     func_Is,
		"IsControl":              func_IsControl,
		"IsDigit":                func_IsDigit,
		"IsGraphic":              func_IsGraphic,
		"IsLetter":               func_IsLetter,
		"IsLower":                func_IsLower,
		"IsMark":                 func_IsMark,
		"IsNumber":               func_IsNumber,
		"IsOneOf":                func_IsOneOf,
		"IsPrint":                func_IsPrint,
		"IsPunct":                func_IsPunct,
		"IsSpace":                func_IsSpace,
		"IsSymbol":               func_IsSymbol,
		"IsTitle":                func_IsTitle,
		"IsUpper":                func_IsUpper,
		"SimpleFold":             func_SimpleFold,
		"To":                     func_To,
		"ToLower":                func_ToLower,
		"ToTitle":                func_ToTitle,
		"ToUpper":                func_ToUpper,
	})
}

func func_In(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.In(ixgo.DirectCallArg[rune](ctx, 0), ixgo.DirectCallArg[[]*q.RangeTable](ctx, 1)...))
}

func func_Is(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Is(ixgo.DirectCallArg[*q.RangeTable](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_IsControl(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsControl(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsDigit(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsDigit(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsGraphic(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsGraphic(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsLetter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsLetter(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsLower(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsLower(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsMark(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsMark(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsNumber(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsNumber(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsOneOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsOneOf(ixgo.DirectCallArg[[]*q.RangeTable](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_IsPrint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsPrint(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsPunct(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsPunct(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsSpace(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsSpace(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsSymbol(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsSymbol(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsTitle(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsTitle(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_IsUpper(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsUpper(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_SimpleFold(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SimpleFold(ixgo.DirectCallArg[rune](ctx, 0)))
}

func method_SpecialCase_ToLower(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SpecialCase.ToLower(ixgo.DirectCallArg[q.SpecialCase](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func method_ptr_SpecialCase_ToLower(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SpecialCase).ToLower(ixgo.DirectCallArg[*q.SpecialCase](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func method_SpecialCase_ToTitle(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SpecialCase.ToTitle(ixgo.DirectCallArg[q.SpecialCase](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func method_ptr_SpecialCase_ToTitle(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SpecialCase).ToTitle(ixgo.DirectCallArg[*q.SpecialCase](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func method_SpecialCase_ToUpper(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SpecialCase.ToUpper(ixgo.DirectCallArg[q.SpecialCase](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func method_ptr_SpecialCase_ToUpper(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SpecialCase).ToUpper(ixgo.DirectCallArg[*q.SpecialCase](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_To(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.To(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_ToLower(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToLower(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_ToTitle(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToTitle(ixgo.DirectCallArg[rune](ctx, 0)))
}

func func_ToUpper(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToUpper(ixgo.DirectCallArg[rune](ctx, 0)))
}
