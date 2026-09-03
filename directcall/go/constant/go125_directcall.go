// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package constant

import (
	q "go/constant"

	"github.com/goplus/ixgo"
	token "go/token"
)

func init() {
	ixgo.RegisterDirectCalls("go/constant", map[string]ixgo.DirectCallAdapter{
		"(*Kind).String":  method_ptr_Kind_String,
		"(Kind).String":   method_Kind_String,
		"BinaryOp":        func_BinaryOp,
		"BitLen":          func_BitLen,
		"BoolVal":         func_BoolVal,
		"Bytes":           func_Bytes,
		"Compare":         func_Compare,
		"Denom":           func_Denom,
		"Imag":            func_Imag,
		"Make":            func_Make,
		"MakeBool":        func_MakeBool,
		"MakeFloat64":     func_MakeFloat64,
		"MakeFromBytes":   func_MakeFromBytes,
		"MakeFromLiteral": func_MakeFromLiteral,
		"MakeImag":        func_MakeImag,
		"MakeInt64":       func_MakeInt64,
		"MakeString":      func_MakeString,
		"MakeUint64":      func_MakeUint64,
		"MakeUnknown":     func_MakeUnknown,
		"Num":             func_Num,
		"Real":            func_Real,
		"Shift":           func_Shift,
		"Sign":            func_Sign,
		"StringVal":       func_StringVal,
		"ToComplex":       func_ToComplex,
		"ToFloat":         func_ToFloat,
		"ToInt":           func_ToInt,
		"UnaryOp":         func_UnaryOp,
		"Val":             func_Val,
	})
}

func func_BinaryOp(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.BinaryOp(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[token.Token](ctx, 1), ixgo.DirectCallArg[q.Value](ctx, 2)))
}

func func_BitLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.BitLen(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_BoolVal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.BoolVal(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Bytes(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Compare(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[token.Token](ctx, 1), ixgo.DirectCallArg[q.Value](ctx, 2)))
}

func func_Denom(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Denom(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_Imag(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Imag(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_Kind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Kind.String(ixgo.DirectCallArg[q.Kind](ctx, 0)))
}

func method_ptr_Kind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Kind).String(ixgo.DirectCallArg[*q.Kind](ctx, 0)))
}

func func_Make(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Make(ixgo.DirectCallArg[any](ctx, 0)))
}

func func_MakeBool(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeBool(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_MakeFloat64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeFloat64(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_MakeFromBytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeFromBytes(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_MakeFromLiteral(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeFromLiteral(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[token.Token](ctx, 1), ixgo.DirectCallArg[uint](ctx, 2)))
}

func func_MakeImag(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeImag(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_MakeInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeInt64(ixgo.DirectCallArg[int64](ctx, 0)))
}

func func_MakeString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeString(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_MakeUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeUint64(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_MakeUnknown(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeUnknown())
}

func func_Num(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Num(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_Real(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Real(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_Shift(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Shift(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[token.Token](ctx, 1), ixgo.DirectCallArg[uint](ctx, 2)))
}

func func_Sign(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sign(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_StringVal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StringVal(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_ToComplex(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToComplex(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_ToFloat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToFloat(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_ToInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToInt(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func func_UnaryOp(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnaryOp(ixgo.DirectCallArg[token.Token](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1), ixgo.DirectCallArg[uint](ctx, 2)))
}

func func_Val(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Val(ixgo.DirectCallArg[q.Value](ctx, 0)))
}
