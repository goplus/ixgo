// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package runtime

import (
	q "runtime"

	"github.com/goplus/ixgo"
	unsafe "unsafe"
)

func init() {
	ixgo.RegisterDirectCalls("runtime", map[string]ixgo.DirectCallAdapter{
		"(*Cleanup).Stop":                    method_ptr_Cleanup_Stop,
		"(*Func).Entry":                      method_ptr_Func_Entry,
		"(*Func).Name":                       method_ptr_Func_Name,
		"(*MemProfileRecord).InUseBytes":     method_ptr_MemProfileRecord_InUseBytes,
		"(*MemProfileRecord).InUseObjects":   method_ptr_MemProfileRecord_InUseObjects,
		"(*MemProfileRecord).Stack":          method_ptr_MemProfileRecord_Stack,
		"(*PanicNilError).Error":             method_ptr_PanicNilError_Error,
		"(*PanicNilError).RuntimeError":      method_ptr_PanicNilError_RuntimeError,
		"(*Pinner).Pin":                      method_ptr_Pinner_Pin,
		"(*Pinner).Unpin":                    method_ptr_Pinner_Unpin,
		"(*StackRecord).Stack":               method_ptr_StackRecord_Stack,
		"(*TypeAssertionError).Error":        method_ptr_TypeAssertionError_Error,
		"(*TypeAssertionError).RuntimeError": method_ptr_TypeAssertionError_RuntimeError,
		"(Cleanup).Stop":                     method_Cleanup_Stop,
		"Breakpoint":                         func_Breakpoint,
		"CPUProfile":                         func_CPUProfile,
		"Callers":                            func_Callers,
		"CallersFrames":                      func_CallersFrames,
		"FuncForPC":                          func_FuncForPC,
		"GC":                                 func_GC,
		"GOMAXPROCS":                         func_GOMAXPROCS,
		"GOROOT":                             func_GOROOT,
		"Goexit":                             func_Goexit,
		"Gosched":                            func_Gosched,
		"KeepAlive":                          func_KeepAlive,
		"LockOSThread":                       func_LockOSThread,
		"NumCPU":                             func_NumCPU,
		"NumCgoCall":                         func_NumCgoCall,
		"NumGoroutine":                       func_NumGoroutine,
		"ReadMemStats":                       func_ReadMemStats,
		"ReadTrace":                          func_ReadTrace,
		"SetBlockProfileRate":                func_SetBlockProfileRate,
		"SetCPUProfileRate":                  func_SetCPUProfileRate,
		"SetCgoTraceback":                    func_SetCgoTraceback,
		"SetDefaultGOMAXPROCS":               func_SetDefaultGOMAXPROCS,
		"SetFinalizer":                       func_SetFinalizer,
		"SetMutexProfileFraction":            func_SetMutexProfileFraction,
		"Stack":                              func_Stack,
		"StartTrace":                         func_StartTrace,
		"StopTrace":                          func_StopTrace,
		"UnlockOSThread":                     func_UnlockOSThread,
		"Version":                            func_Version,
	})
}

func func_Breakpoint(ctx ixgo.DirectCallContext) {
	q.Breakpoint()
}

func func_CPUProfile(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CPUProfile())
}

func func_Callers(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Callers(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[[]uintptr](ctx, 1)))
}

func func_CallersFrames(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CallersFrames(ixgo.DirectCallArg[[]uintptr](ctx, 0)))
}

func method_Cleanup_Stop(ctx ixgo.DirectCallContext) {
	q.Cleanup.Stop(ixgo.DirectCallArg[q.Cleanup](ctx, 0))
}

func method_ptr_Cleanup_Stop(ctx ixgo.DirectCallContext) {
	(*q.Cleanup).Stop(ixgo.DirectCallArg[*q.Cleanup](ctx, 0))
}

func method_ptr_Func_Entry(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Func).Entry(ixgo.DirectCallArg[*q.Func](ctx, 0)))
}

func method_ptr_Func_Name(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Func).Name(ixgo.DirectCallArg[*q.Func](ctx, 0)))
}

func func_FuncForPC(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FuncForPC(ixgo.DirectCallArg[uintptr](ctx, 0)))
}

func func_GC(ctx ixgo.DirectCallContext) {
	q.GC()
}

func func_GOMAXPROCS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.GOMAXPROCS(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_GOROOT(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.GOROOT())
}

func func_Goexit(ctx ixgo.DirectCallContext) {
	q.Goexit()
}

func func_Gosched(ctx ixgo.DirectCallContext) {
	q.Gosched()
}

func func_KeepAlive(ctx ixgo.DirectCallContext) {
	q.KeepAlive(ixgo.DirectCallArg[interface{}](ctx, 0))
}

func func_LockOSThread(ctx ixgo.DirectCallContext) {
	q.LockOSThread()
}

func method_ptr_MemProfileRecord_InUseBytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MemProfileRecord).InUseBytes(ixgo.DirectCallArg[*q.MemProfileRecord](ctx, 0)))
}

func method_ptr_MemProfileRecord_InUseObjects(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MemProfileRecord).InUseObjects(ixgo.DirectCallArg[*q.MemProfileRecord](ctx, 0)))
}

func method_ptr_MemProfileRecord_Stack(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MemProfileRecord).Stack(ixgo.DirectCallArg[*q.MemProfileRecord](ctx, 0)))
}

func func_NumCPU(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NumCPU())
}

func func_NumCgoCall(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NumCgoCall())
}

func func_NumGoroutine(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NumGoroutine())
}

func method_ptr_PanicNilError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PanicNilError).Error(ixgo.DirectCallArg[*q.PanicNilError](ctx, 0)))
}

func method_ptr_PanicNilError_RuntimeError(ctx ixgo.DirectCallContext) {
	(*q.PanicNilError).RuntimeError(ixgo.DirectCallArg[*q.PanicNilError](ctx, 0))
}

func method_ptr_Pinner_Pin(ctx ixgo.DirectCallContext) {
	(*q.Pinner).Pin(ixgo.DirectCallArg[*q.Pinner](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1))
}

func method_ptr_Pinner_Unpin(ctx ixgo.DirectCallContext) {
	(*q.Pinner).Unpin(ixgo.DirectCallArg[*q.Pinner](ctx, 0))
}

func func_ReadMemStats(ctx ixgo.DirectCallContext) {
	q.ReadMemStats(ixgo.DirectCallArg[*q.MemStats](ctx, 0))
}

func func_ReadTrace(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ReadTrace())
}

func func_SetBlockProfileRate(ctx ixgo.DirectCallContext) {
	q.SetBlockProfileRate(ixgo.DirectCallArg[int](ctx, 0))
}

func func_SetCPUProfileRate(ctx ixgo.DirectCallContext) {
	q.SetCPUProfileRate(ixgo.DirectCallArg[int](ctx, 0))
}

func func_SetCgoTraceback(ctx ixgo.DirectCallContext) {
	q.SetCgoTraceback(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[unsafe.Pointer](ctx, 1), ixgo.DirectCallArg[unsafe.Pointer](ctx, 2), ixgo.DirectCallArg[unsafe.Pointer](ctx, 3))
}

func func_SetDefaultGOMAXPROCS(ctx ixgo.DirectCallContext) {
	q.SetDefaultGOMAXPROCS()
}

func func_SetFinalizer(ctx ixgo.DirectCallContext) {
	q.SetFinalizer(ixgo.DirectCallArg[interface{}](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1))
}

func func_SetMutexProfileFraction(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SetMutexProfileFraction(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_Stack(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Stack(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1)))
}

func method_ptr_StackRecord_Stack(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StackRecord).Stack(ixgo.DirectCallArg[*q.StackRecord](ctx, 0)))
}

func func_StartTrace(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StartTrace())
}

func func_StopTrace(ctx ixgo.DirectCallContext) {
	q.StopTrace()
}

func method_ptr_TypeAssertionError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAssertionError).Error(ixgo.DirectCallArg[*q.TypeAssertionError](ctx, 0)))
}

func method_ptr_TypeAssertionError_RuntimeError(ctx ixgo.DirectCallContext) {
	(*q.TypeAssertionError).RuntimeError(ixgo.DirectCallArg[*q.TypeAssertionError](ctx, 0))
}

func func_UnlockOSThread(ctx ixgo.DirectCallContext) {
	q.UnlockOSThread()
}

func func_Version(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Version())
}
