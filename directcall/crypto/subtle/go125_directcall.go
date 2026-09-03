// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package subtle

import (
	q "crypto/subtle"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/subtle", map[string]ixgo.DirectCallAdapter{
		"ConstantTimeByteEq":        func_ConstantTimeByteEq,
		"ConstantTimeCompare":       func_ConstantTimeCompare,
		"ConstantTimeCopy":          func_ConstantTimeCopy,
		"ConstantTimeEq":            func_ConstantTimeEq,
		"ConstantTimeLessOrEq":      func_ConstantTimeLessOrEq,
		"ConstantTimeSelect":        func_ConstantTimeSelect,
		"WithDataIndependentTiming": func_WithDataIndependentTiming,
		"XORBytes":                  func_XORBytes,
	})
}

func func_ConstantTimeByteEq(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ConstantTimeByteEq(ixgo.DirectCallArg[uint8](ctx, 0), ixgo.DirectCallArg[uint8](ctx, 1)))
}

func func_ConstantTimeCompare(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ConstantTimeCompare(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_ConstantTimeCopy(ctx ixgo.DirectCallContext) {
	q.ConstantTimeCopy(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2))
}

func func_ConstantTimeEq(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ConstantTimeEq(ixgo.DirectCallArg[int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func func_ConstantTimeLessOrEq(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ConstantTimeLessOrEq(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_ConstantTimeSelect(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ConstantTimeSelect(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_WithDataIndependentTiming(ctx ixgo.DirectCallContext) {
	q.WithDataIndependentTiming(ixgo.DirectCallArg[func()](ctx, 0))
}

func func_XORBytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.XORBytes(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}
