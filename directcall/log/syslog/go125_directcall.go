// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package syslog

import (
	q "log/syslog"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("log/syslog", map[string]ixgo.DirectCallAdapter{
		"(*Writer).Alert":   method_ptr_Writer_Alert,
		"(*Writer).Close":   method_ptr_Writer_Close,
		"(*Writer).Crit":    method_ptr_Writer_Crit,
		"(*Writer).Debug":   method_ptr_Writer_Debug,
		"(*Writer).Emerg":   method_ptr_Writer_Emerg,
		"(*Writer).Err":     method_ptr_Writer_Err,
		"(*Writer).Info":    method_ptr_Writer_Info,
		"(*Writer).Notice":  method_ptr_Writer_Notice,
		"(*Writer).Warning": method_ptr_Writer_Warning,
	})
}

func method_ptr_Writer_Alert(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Alert(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Writer_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Close(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Crit(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Crit(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Writer_Debug(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Debug(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Writer_Emerg(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Emerg(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Writer_Err(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Err(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Writer_Info(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Info(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Writer_Notice(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Notice(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Writer_Warning(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Warning(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}
