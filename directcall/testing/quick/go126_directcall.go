// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package quick

import (
	q "testing/quick"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("testing/quick", map[string]ixgo.DirectCallAdapter{
		"(*CheckEqualError).Error": method_ptr_CheckEqualError_Error,
		"(*CheckError).Error":      method_ptr_CheckError_Error,
		"(*SetupError).Error":      method_ptr_SetupError_Error,
		"(SetupError).Error":       method_SetupError_Error,
		"Check":                    func_Check,
		"CheckEqual":               func_CheckEqual,
	})
}

func func_Check(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Check(ixgo.DirectCallArg[any](ctx, 0), ixgo.DirectCallArg[*q.Config](ctx, 1)))
}

func func_CheckEqual(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CheckEqual(ixgo.DirectCallArg[any](ctx, 0), ixgo.DirectCallArg[any](ctx, 1), ixgo.DirectCallArg[*q.Config](ctx, 2)))
}

func method_ptr_CheckEqualError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CheckEqualError).Error(ixgo.DirectCallArg[*q.CheckEqualError](ctx, 0)))
}

func method_ptr_CheckError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CheckError).Error(ixgo.DirectCallArg[*q.CheckError](ctx, 0)))
}

func method_SetupError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SetupError.Error(ixgo.DirectCallArg[q.SetupError](ctx, 0)))
}

func method_ptr_SetupError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SetupError).Error(ixgo.DirectCallArg[*q.SetupError](ctx, 0)))
}
