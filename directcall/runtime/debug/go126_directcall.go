// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package debug

import (
	q "runtime/debug"

	"github.com/goplus/ixgo"
	os "os"
)

func init() {
	ixgo.RegisterDirectCalls("runtime/debug", map[string]ixgo.DirectCallAdapter{
		"(*BuildInfo).String": method_ptr_BuildInfo_String,
		"FreeOSMemory":        func_FreeOSMemory,
		"PrintStack":          func_PrintStack,
		"ReadGCStats":         func_ReadGCStats,
		"SetCrashOutput":      func_SetCrashOutput,
		"SetGCPercent":        func_SetGCPercent,
		"SetMaxStack":         func_SetMaxStack,
		"SetMaxThreads":       func_SetMaxThreads,
		"SetMemoryLimit":      func_SetMemoryLimit,
		"SetPanicOnFault":     func_SetPanicOnFault,
		"SetTraceback":        func_SetTraceback,
		"Stack":               func_Stack,
		"WriteHeapDump":       func_WriteHeapDump,
	})
}

func method_ptr_BuildInfo_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BuildInfo).String(ixgo.DirectCallArg[*q.BuildInfo](ctx, 0)))
}

func func_FreeOSMemory(ctx ixgo.DirectCallContext) {
	q.FreeOSMemory()
}

func func_PrintStack(ctx ixgo.DirectCallContext) {
	q.PrintStack()
}

func func_ReadGCStats(ctx ixgo.DirectCallContext) {
	q.ReadGCStats(ixgo.DirectCallArg[*q.GCStats](ctx, 0))
}

func func_SetCrashOutput(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SetCrashOutput(ixgo.DirectCallArg[*os.File](ctx, 0), ixgo.DirectCallArg[q.CrashOptions](ctx, 1)))
}

func func_SetGCPercent(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SetGCPercent(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_SetMaxStack(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SetMaxStack(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_SetMaxThreads(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SetMaxThreads(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_SetMemoryLimit(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SetMemoryLimit(ixgo.DirectCallArg[int64](ctx, 0)))
}

func func_SetPanicOnFault(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SetPanicOnFault(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_SetTraceback(ctx ixgo.DirectCallContext) {
	q.SetTraceback(ixgo.DirectCallArg[string](ctx, 0))
}

func func_Stack(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Stack())
}

func func_WriteHeapDump(ctx ixgo.DirectCallContext) {
	q.WriteHeapDump(ixgo.DirectCallArg[uintptr](ctx, 0))
}
