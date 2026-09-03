// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package trace

import (
	q "runtime/trace"

	context "context"
	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("runtime/trace", map[string]ixgo.DirectCallAdapter{
		"(*FlightRecorder).Enabled": method_ptr_FlightRecorder_Enabled,
		"(*FlightRecorder).Start":   method_ptr_FlightRecorder_Start,
		"(*FlightRecorder).Stop":    method_ptr_FlightRecorder_Stop,
		"(*Region).End":             method_ptr_Region_End,
		"(*Task).End":               method_ptr_Task_End,
		"IsEnabled":                 func_IsEnabled,
		"Log":                       func_Log,
		"Logf":                      func_Logf,
		"NewFlightRecorder":         func_NewFlightRecorder,
		"Start":                     func_Start,
		"StartRegion":               func_StartRegion,
		"Stop":                      func_Stop,
		"WithRegion":                func_WithRegion,
	})
}

func method_ptr_FlightRecorder_Enabled(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlightRecorder).Enabled(ixgo.DirectCallArg[*q.FlightRecorder](ctx, 0)))
}

func method_ptr_FlightRecorder_Start(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlightRecorder).Start(ixgo.DirectCallArg[*q.FlightRecorder](ctx, 0)))
}

func method_ptr_FlightRecorder_Stop(ctx ixgo.DirectCallContext) {
	(*q.FlightRecorder).Stop(ixgo.DirectCallArg[*q.FlightRecorder](ctx, 0))
}

func func_IsEnabled(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsEnabled())
}

func func_Log(ctx ixgo.DirectCallContext) {
	q.Log(ixgo.DirectCallArg[context.Context](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func func_Logf(ctx ixgo.DirectCallContext) {
	q.Logf(ixgo.DirectCallArg[context.Context](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[[]any](ctx, 3)...)
}

func func_NewFlightRecorder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewFlightRecorder(ixgo.DirectCallArg[q.FlightRecorderConfig](ctx, 0)))
}

func method_ptr_Region_End(ctx ixgo.DirectCallContext) {
	(*q.Region).End(ixgo.DirectCallArg[*q.Region](ctx, 0))
}

func func_Start(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Start(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func func_StartRegion(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StartRegion(ixgo.DirectCallArg[context.Context](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Stop(ctx ixgo.DirectCallContext) {
	q.Stop()
}

func method_ptr_Task_End(ctx ixgo.DirectCallContext) {
	(*q.Task).End(ixgo.DirectCallArg[*q.Task](ctx, 0))
}

func func_WithRegion(ctx ixgo.DirectCallContext) {
	q.WithRegion(ixgo.DirectCallArg[context.Context](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[func()](ctx, 2))
}
