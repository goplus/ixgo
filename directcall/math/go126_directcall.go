// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package math

import (
	q "math"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("math", map[string]ixgo.DirectCallAdapter{
		"Abs":             func_Abs,
		"Acos":            func_Acos,
		"Acosh":           func_Acosh,
		"Asin":            func_Asin,
		"Asinh":           func_Asinh,
		"Atan":            func_Atan,
		"Atan2":           func_Atan2,
		"Atanh":           func_Atanh,
		"Cbrt":            func_Cbrt,
		"Ceil":            func_Ceil,
		"Copysign":        func_Copysign,
		"Cos":             func_Cos,
		"Cosh":            func_Cosh,
		"Dim":             func_Dim,
		"Erf":             func_Erf,
		"Erfc":            func_Erfc,
		"Erfcinv":         func_Erfcinv,
		"Erfinv":          func_Erfinv,
		"Exp":             func_Exp,
		"Exp2":            func_Exp2,
		"Expm1":           func_Expm1,
		"FMA":             func_FMA,
		"Float32bits":     func_Float32bits,
		"Float32frombits": func_Float32frombits,
		"Float64bits":     func_Float64bits,
		"Float64frombits": func_Float64frombits,
		"Floor":           func_Floor,
		"Gamma":           func_Gamma,
		"Hypot":           func_Hypot,
		"Ilogb":           func_Ilogb,
		"Inf":             func_Inf,
		"IsInf":           func_IsInf,
		"IsNaN":           func_IsNaN,
		"J0":              func_J0,
		"J1":              func_J1,
		"Jn":              func_Jn,
		"Ldexp":           func_Ldexp,
		"Log":             func_Log,
		"Log10":           func_Log10,
		"Log1p":           func_Log1p,
		"Log2":            func_Log2,
		"Logb":            func_Logb,
		"Max":             func_Max,
		"Min":             func_Min,
		"Mod":             func_Mod,
		"NaN":             func_NaN,
		"Nextafter":       func_Nextafter,
		"Nextafter32":     func_Nextafter32,
		"Pow":             func_Pow,
		"Pow10":           func_Pow10,
		"Remainder":       func_Remainder,
		"Round":           func_Round,
		"RoundToEven":     func_RoundToEven,
		"Signbit":         func_Signbit,
		"Sin":             func_Sin,
		"Sinh":            func_Sinh,
		"Sqrt":            func_Sqrt,
		"Tan":             func_Tan,
		"Tanh":            func_Tanh,
		"Trunc":           func_Trunc,
		"Y0":              func_Y0,
		"Y1":              func_Y1,
		"Yn":              func_Yn,
	})
}

func func_Abs(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Abs(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Acos(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Acos(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Acosh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Acosh(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Asin(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Asin(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Asinh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Asinh(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Atan(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Atan(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Atan2(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Atan2(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Atanh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Atanh(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Cbrt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Cbrt(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Ceil(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Ceil(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Copysign(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Copysign(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Cos(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Cos(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Cosh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Cosh(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Dim(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Dim(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Erf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Erf(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Erfc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Erfc(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Erfcinv(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Erfcinv(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Erfinv(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Erfinv(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Exp(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Exp(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Exp2(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Exp2(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Expm1(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Expm1(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_FMA(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FMA(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1), ixgo.DirectCallArg[float64](ctx, 2)))
}

func func_Float32bits(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float32bits(ixgo.DirectCallArg[float32](ctx, 0)))
}

func func_Float32frombits(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float32frombits(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_Float64bits(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float64bits(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Float64frombits(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float64frombits(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_Floor(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Floor(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Gamma(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Gamma(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Hypot(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Hypot(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Ilogb(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Ilogb(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Inf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Inf(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_IsInf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsInf(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_IsNaN(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsNaN(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_J0(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.J0(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_J1(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.J1(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Jn(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Jn(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Ldexp(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Ldexp(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_Log(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Log(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Log10(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Log10(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Log1p(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Log1p(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Log2(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Log2(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Logb(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Logb(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Max(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Max(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Min(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Min(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Mod(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Mod(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_NaN(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NaN())
}

func func_Nextafter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Nextafter(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Nextafter32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Nextafter32(ixgo.DirectCallArg[float32](ctx, 0), ixgo.DirectCallArg[float32](ctx, 1)))
}

func func_Pow(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pow(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Pow10(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pow10(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_Remainder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Remainder(ixgo.DirectCallArg[float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_Round(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Round(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_RoundToEven(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RoundToEven(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Signbit(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Signbit(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Sin(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sin(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Sinh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sinh(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Sqrt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sqrt(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Tan(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Tan(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Tanh(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Tanh(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Trunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Trunc(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Y0(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Y0(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Y1(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Y1(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Yn(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Yn(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}
