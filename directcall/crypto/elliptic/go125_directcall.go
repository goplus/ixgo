// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package elliptic

import (
	q "crypto/elliptic"

	"github.com/goplus/ixgo"
	big "math/big"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/elliptic", map[string]ixgo.DirectCallAdapter{
		"(*CurveParams).IsOnCurve": method_ptr_CurveParams_IsOnCurve,
		"(*CurveParams).Params":    method_ptr_CurveParams_Params,
		"Marshal":                  func_Marshal,
		"MarshalCompressed":        func_MarshalCompressed,
		"P224":                     func_P224,
		"P256":                     func_P256,
		"P384":                     func_P384,
		"P521":                     func_P521,
	})
}

func method_ptr_CurveParams_IsOnCurve(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CurveParams).IsOnCurve(ixgo.DirectCallArg[*q.CurveParams](ctx, 0), ixgo.DirectCallArg[*big.Int](ctx, 1), ixgo.DirectCallArg[*big.Int](ctx, 2)))
}

func method_ptr_CurveParams_Params(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CurveParams).Params(ixgo.DirectCallArg[*q.CurveParams](ctx, 0)))
}

func func_Marshal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Marshal(ixgo.DirectCallArg[q.Curve](ctx, 0), ixgo.DirectCallArg[*big.Int](ctx, 1), ixgo.DirectCallArg[*big.Int](ctx, 2)))
}

func func_MarshalCompressed(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MarshalCompressed(ixgo.DirectCallArg[q.Curve](ctx, 0), ixgo.DirectCallArg[*big.Int](ctx, 1), ixgo.DirectCallArg[*big.Int](ctx, 2)))
}

func func_P224(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.P224())
}

func func_P256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.P256())
}

func func_P384(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.P384())
}

func func_P521(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.P521())
}
