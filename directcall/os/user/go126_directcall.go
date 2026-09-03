// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package user

import (
	q "os/user"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("os/user", map[string]ixgo.DirectCallAdapter{
		"(*UnknownGroupError).Error":   method_ptr_UnknownGroupError_Error,
		"(*UnknownGroupIdError).Error": method_ptr_UnknownGroupIdError_Error,
		"(*UnknownUserError).Error":    method_ptr_UnknownUserError_Error,
		"(*UnknownUserIdError).Error":  method_ptr_UnknownUserIdError_Error,
		"(UnknownGroupError).Error":    method_UnknownGroupError_Error,
		"(UnknownGroupIdError).Error":  method_UnknownGroupIdError_Error,
		"(UnknownUserError).Error":     method_UnknownUserError_Error,
		"(UnknownUserIdError).Error":   method_UnknownUserIdError_Error,
	})
}

func method_UnknownGroupError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnknownGroupError.Error(ixgo.DirectCallArg[q.UnknownGroupError](ctx, 0)))
}

func method_ptr_UnknownGroupError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnknownGroupError).Error(ixgo.DirectCallArg[*q.UnknownGroupError](ctx, 0)))
}

func method_UnknownGroupIdError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnknownGroupIdError.Error(ixgo.DirectCallArg[q.UnknownGroupIdError](ctx, 0)))
}

func method_ptr_UnknownGroupIdError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnknownGroupIdError).Error(ixgo.DirectCallArg[*q.UnknownGroupIdError](ctx, 0)))
}

func method_UnknownUserError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnknownUserError.Error(ixgo.DirectCallArg[q.UnknownUserError](ctx, 0)))
}

func method_ptr_UnknownUserError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnknownUserError).Error(ixgo.DirectCallArg[*q.UnknownUserError](ctx, 0)))
}

func method_UnknownUserIdError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnknownUserIdError.Error(ixgo.DirectCallArg[q.UnknownUserIdError](ctx, 0)))
}

func method_ptr_UnknownUserIdError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnknownUserIdError).Error(ixgo.DirectCallArg[*q.UnknownUserIdError](ctx, 0)))
}
