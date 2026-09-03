// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package cmplx

import (
	q "math/cmplx"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("math/cmplx", map[string]ixgo.DirectCallAdapter{
		"Abs":   func_Abs,
		"Acos":  func_Acos,
		"Acosh": func_Acosh,
		"Asin":  func_Asin,
		"Asinh": func_Asinh,
		"Atan":  func_Atan,
		"Atanh": func_Atanh,
		"Conj":  func_Conj,
		"Cos":   func_Cos,
		"Cosh":  func_Cosh,
		"Cot":   func_Cot,
		"Exp":   func_Exp,
		"Inf":   func_Inf,
		"IsInf": func_IsInf,
		"IsNaN": func_IsNaN,
		"Log":   func_Log,
		"Log10": func_Log10,
		"NaN":   func_NaN,
		"Phase": func_Phase,
		"Pow":   func_Pow,
		"Rect":  func_Rect,
		"Sin":   func_Sin,
		"Sinh":  func_Sinh,
		"Sqrt":  func_Sqrt,
		"Tan":   func_Tan,
		"Tanh":  func_Tanh,
	})
}

func func_Abs(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Abs(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Acos(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Acos(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Acosh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Acosh(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Asin(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Asin(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Asinh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Asinh(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Atan(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Atan(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Atanh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Atanh(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Conj(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Conj(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Cos(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Cos(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Cosh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Cosh(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Cot(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Cot(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Exp(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Exp(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Inf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Inf())
}

func func_IsInf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsInf(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_IsNaN(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsNaN(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Log(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Log(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Log10(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Log10(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_NaN(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NaN())
}

func func_Phase(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Phase(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Pow(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pow(ixgo.DirectCallArg[complex128](ctx, 0), ixgo.DirectCallArg[complex128](ctx, 1)))
}

func func_Rect(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rect(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Sin(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sin(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Sinh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sinh(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Sqrt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sqrt(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Tan(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Tan(ixgo.DirectCallArg[complex128](ctx, 0)))
}

func func_Tanh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Tanh(ixgo.DirectCallArg[complex128](ctx, 0)))
}
