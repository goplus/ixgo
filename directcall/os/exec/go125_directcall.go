// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package exec

import (
	q "os/exec"

	context "context"
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("os/exec", map[string]ixgo.DirectCallAdapter{
		"(*Cmd).Environ":     method_ptr_Cmd_Environ,
		"(*Cmd).Run":         method_ptr_Cmd_Run,
		"(*Cmd).Start":       method_ptr_Cmd_Start,
		"(*Cmd).String":      method_ptr_Cmd_String,
		"(*Cmd).Wait":        method_ptr_Cmd_Wait,
		"(*Error).Error":     method_ptr_Error_Error,
		"(*Error).Unwrap":    method_ptr_Error_Unwrap,
		"(*ExitError).Error": method_ptr_ExitError_Error,
		"Command":            func_Command,
		"CommandContext":     func_CommandContext,
	})
}

func method_ptr_Cmd_Environ(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Cmd).Environ(ixgo.DirectCallArg[*q.Cmd](ctx, 0)))
}

func method_ptr_Cmd_Run(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Cmd).Run(ixgo.DirectCallArg[*q.Cmd](ctx, 0)))
}

func method_ptr_Cmd_Start(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Cmd).Start(ixgo.DirectCallArg[*q.Cmd](ctx, 0)))
}

func method_ptr_Cmd_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Cmd).String(ixgo.DirectCallArg[*q.Cmd](ctx, 0)))
}

func method_ptr_Cmd_Wait(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Cmd).Wait(ixgo.DirectCallArg[*q.Cmd](ctx, 0)))
}

func func_Command(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Command(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[[]string](ctx, 1)...))
}

func func_CommandContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CommandContext(ixgo.DirectCallArg[context.Context](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[[]string](ctx, 2)...))
}

func method_ptr_Error_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Error).Error(ixgo.DirectCallArg[*q.Error](ctx, 0)))
}

func method_ptr_Error_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Error).Unwrap(ixgo.DirectCallArg[*q.Error](ctx, 0)))
}

func method_ptr_ExitError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ExitError).Error(ixgo.DirectCallArg[*q.ExitError](ctx, 0)))
}
