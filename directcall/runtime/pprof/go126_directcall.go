// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package pprof

import (
	q "runtime/pprof"

	context "context"
	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("runtime/pprof", map[string]ixgo.DirectCallAdapter{
		"(*Profile).Add":     method_ptr_Profile_Add,
		"(*Profile).Count":   method_ptr_Profile_Count,
		"(*Profile).Name":    method_ptr_Profile_Name,
		"(*Profile).Remove":  method_ptr_Profile_Remove,
		"(*Profile).WriteTo": method_ptr_Profile_WriteTo,
		"Do":                 func_Do,
		"ForLabels":          func_ForLabels,
		"Labels":             func_Labels,
		"Lookup":             func_Lookup,
		"NewProfile":         func_NewProfile,
		"Profiles":           func_Profiles,
		"SetGoroutineLabels": func_SetGoroutineLabels,
		"StartCPUProfile":    func_StartCPUProfile,
		"StopCPUProfile":     func_StopCPUProfile,
		"WithLabels":         func_WithLabels,
		"WriteHeapProfile":   func_WriteHeapProfile,
	})
}

func func_Do(ctx ixgo.DirectCallContext) {
	q.Do(ixgo.DirectCallArg[context.Context](ctx, 0), ixgo.DirectCallArg[q.LabelSet](ctx, 1), ixgo.DirectCallArg[func(context.Context)](ctx, 2))
}

func func_ForLabels(ctx ixgo.DirectCallContext) {
	q.ForLabels(ixgo.DirectCallArg[context.Context](ctx, 0), ixgo.DirectCallArg[func(key string, value string) bool](ctx, 1))
}

func func_Labels(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Labels(ixgo.DirectCallArg[[]string](ctx, 0)...))
}

func func_Lookup(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Lookup(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_NewProfile(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewProfile(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_Profile_Add(ctx ixgo.DirectCallContext) {
	(*q.Profile).Add(ixgo.DirectCallArg[*q.Profile](ctx, 0), ixgo.DirectCallArg[any](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func method_ptr_Profile_Count(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Profile).Count(ixgo.DirectCallArg[*q.Profile](ctx, 0)))
}

func method_ptr_Profile_Name(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Profile).Name(ixgo.DirectCallArg[*q.Profile](ctx, 0)))
}

func method_ptr_Profile_Remove(ctx ixgo.DirectCallContext) {
	(*q.Profile).Remove(ixgo.DirectCallArg[*q.Profile](ctx, 0), ixgo.DirectCallArg[any](ctx, 1))
}

func method_ptr_Profile_WriteTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Profile).WriteTo(ixgo.DirectCallArg[*q.Profile](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_Profiles(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Profiles())
}

func func_SetGoroutineLabels(ctx ixgo.DirectCallContext) {
	q.SetGoroutineLabels(ixgo.DirectCallArg[context.Context](ctx, 0))
}

func func_StartCPUProfile(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StartCPUProfile(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func func_StopCPUProfile(ctx ixgo.DirectCallContext) {
	q.StopCPUProfile()
}

func func_WithLabels(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WithLabels(ixgo.DirectCallArg[context.Context](ctx, 0), ixgo.DirectCallArg[q.LabelSet](ctx, 1)))
}

func func_WriteHeapProfile(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WriteHeapProfile(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}
