// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package testing

import (
	q "testing"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("testing", map[string]ixgo.DirectCallAdapter{
		"(*B).Elapsed":                         method_ptr_B_Elapsed,
		"(*B).Loop":                            method_ptr_B_Loop,
		"(*B).ReportAllocs":                    method_ptr_B_ReportAllocs,
		"(*B).ReportMetric":                    method_ptr_B_ReportMetric,
		"(*B).ResetTimer":                      method_ptr_B_ResetTimer,
		"(*B).Run":                             method_ptr_B_Run,
		"(*B).RunParallel":                     method_ptr_B_RunParallel,
		"(*B).SetBytes":                        method_ptr_B_SetBytes,
		"(*B).SetParallelism":                  method_ptr_B_SetParallelism,
		"(*B).StartTimer":                      method_ptr_B_StartTimer,
		"(*B).StopTimer":                       method_ptr_B_StopTimer,
		"(*BenchmarkResult).AllocedBytesPerOp": method_ptr_BenchmarkResult_AllocedBytesPerOp,
		"(*BenchmarkResult).AllocsPerOp":       method_ptr_BenchmarkResult_AllocsPerOp,
		"(*BenchmarkResult).MemString":         method_ptr_BenchmarkResult_MemString,
		"(*BenchmarkResult).NsPerOp":           method_ptr_BenchmarkResult_NsPerOp,
		"(*BenchmarkResult).String":            method_ptr_BenchmarkResult_String,
		"(*F).Add":                             method_ptr_F_Add,
		"(*F).Fail":                            method_ptr_F_Fail,
		"(*F).Fuzz":                            method_ptr_F_Fuzz,
		"(*F).Helper":                          method_ptr_F_Helper,
		"(*F).Skipped":                         method_ptr_F_Skipped,
		"(*M).Run":                             method_ptr_M_Run,
		"(*PB).Next":                           method_ptr_PB_Next,
		"(*T).Chdir":                           method_ptr_T_Chdir,
		"(*T).Parallel":                        method_ptr_T_Parallel,
		"(*T).Run":                             method_ptr_T_Run,
		"(*T).Setenv":                          method_ptr_T_Setenv,
		"(BenchmarkResult).AllocedBytesPerOp":  method_BenchmarkResult_AllocedBytesPerOp,
		"(BenchmarkResult).AllocsPerOp":        method_BenchmarkResult_AllocsPerOp,
		"(BenchmarkResult).MemString":          method_BenchmarkResult_MemString,
		"(BenchmarkResult).NsPerOp":            method_BenchmarkResult_NsPerOp,
		"(BenchmarkResult).String":             method_BenchmarkResult_String,
		"AllocsPerRun":                         func_AllocsPerRun,
		"Benchmark":                            func_Benchmark,
		"CoverMode":                            func_CoverMode,
		"Coverage":                             func_Coverage,
		"Init":                                 func_Init,
		"Main":                                 func_Main,
		"RegisterCover":                        func_RegisterCover,
		"RunBenchmarks":                        func_RunBenchmarks,
		"RunExamples":                          func_RunExamples,
		"RunTests":                             func_RunTests,
		"Short":                                func_Short,
		"Testing":                              func_Testing,
		"Verbose":                              func_Verbose,
	})
}

func func_AllocsPerRun(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AllocsPerRun(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[func()](ctx, 1)))
}

func method_ptr_B_Elapsed(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.B).Elapsed(ixgo.DirectCallArg[*q.B](ctx, 0)))
}

func method_ptr_B_Loop(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.B).Loop(ixgo.DirectCallArg[*q.B](ctx, 0)))
}

func method_ptr_B_ReportAllocs(ctx ixgo.DirectCallContext) {
	(*q.B).ReportAllocs(ixgo.DirectCallArg[*q.B](ctx, 0))
}

func method_ptr_B_ReportMetric(ctx ixgo.DirectCallContext) {
	(*q.B).ReportMetric(ixgo.DirectCallArg[*q.B](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ptr_B_ResetTimer(ctx ixgo.DirectCallContext) {
	(*q.B).ResetTimer(ixgo.DirectCallArg[*q.B](ctx, 0))
}

func method_ptr_B_Run(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.B).Run(ixgo.DirectCallArg[*q.B](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[func(b *q.B)](ctx, 2)))
}

func method_ptr_B_RunParallel(ctx ixgo.DirectCallContext) {
	(*q.B).RunParallel(ixgo.DirectCallArg[*q.B](ctx, 0), ixgo.DirectCallArg[func(*q.PB)](ctx, 1))
}

func method_ptr_B_SetBytes(ctx ixgo.DirectCallContext) {
	(*q.B).SetBytes(ixgo.DirectCallArg[*q.B](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1))
}

func method_ptr_B_SetParallelism(ctx ixgo.DirectCallContext) {
	(*q.B).SetParallelism(ixgo.DirectCallArg[*q.B](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_B_StartTimer(ctx ixgo.DirectCallContext) {
	(*q.B).StartTimer(ixgo.DirectCallArg[*q.B](ctx, 0))
}

func method_ptr_B_StopTimer(ctx ixgo.DirectCallContext) {
	(*q.B).StopTimer(ixgo.DirectCallArg[*q.B](ctx, 0))
}

func func_Benchmark(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Benchmark(ixgo.DirectCallArg[func(b *q.B)](ctx, 0)))
}

func method_BenchmarkResult_AllocedBytesPerOp(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.BenchmarkResult.AllocedBytesPerOp(ixgo.DirectCallArg[q.BenchmarkResult](ctx, 0)))
}

func method_ptr_BenchmarkResult_AllocedBytesPerOp(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BenchmarkResult).AllocedBytesPerOp(ixgo.DirectCallArg[*q.BenchmarkResult](ctx, 0)))
}

func method_BenchmarkResult_AllocsPerOp(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.BenchmarkResult.AllocsPerOp(ixgo.DirectCallArg[q.BenchmarkResult](ctx, 0)))
}

func method_ptr_BenchmarkResult_AllocsPerOp(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BenchmarkResult).AllocsPerOp(ixgo.DirectCallArg[*q.BenchmarkResult](ctx, 0)))
}

func method_BenchmarkResult_MemString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.BenchmarkResult.MemString(ixgo.DirectCallArg[q.BenchmarkResult](ctx, 0)))
}

func method_ptr_BenchmarkResult_MemString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BenchmarkResult).MemString(ixgo.DirectCallArg[*q.BenchmarkResult](ctx, 0)))
}

func method_BenchmarkResult_NsPerOp(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.BenchmarkResult.NsPerOp(ixgo.DirectCallArg[q.BenchmarkResult](ctx, 0)))
}

func method_ptr_BenchmarkResult_NsPerOp(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BenchmarkResult).NsPerOp(ixgo.DirectCallArg[*q.BenchmarkResult](ctx, 0)))
}

func method_BenchmarkResult_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.BenchmarkResult.String(ixgo.DirectCallArg[q.BenchmarkResult](ctx, 0)))
}

func method_ptr_BenchmarkResult_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BenchmarkResult).String(ixgo.DirectCallArg[*q.BenchmarkResult](ctx, 0)))
}

func func_CoverMode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CoverMode())
}

func func_Coverage(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Coverage())
}

func method_ptr_F_Add(ctx ixgo.DirectCallContext) {
	(*q.F).Add(ixgo.DirectCallArg[*q.F](ctx, 0), ixgo.DirectCallArg[[]any](ctx, 1)...)
}

func method_ptr_F_Fail(ctx ixgo.DirectCallContext) {
	(*q.F).Fail(ixgo.DirectCallArg[*q.F](ctx, 0))
}

func method_ptr_F_Fuzz(ctx ixgo.DirectCallContext) {
	(*q.F).Fuzz(ixgo.DirectCallArg[*q.F](ctx, 0), ixgo.DirectCallArg[any](ctx, 1))
}

func method_ptr_F_Helper(ctx ixgo.DirectCallContext) {
	(*q.F).Helper(ixgo.DirectCallArg[*q.F](ctx, 0))
}

func method_ptr_F_Skipped(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.F).Skipped(ixgo.DirectCallArg[*q.F](ctx, 0)))
}

func func_Init(ctx ixgo.DirectCallContext) {
	q.Init()
}

func method_ptr_M_Run(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.M).Run(ixgo.DirectCallArg[*q.M](ctx, 0)))
}

func func_Main(ctx ixgo.DirectCallContext) {
	q.Main(ixgo.DirectCallArg[func(pat string, str string) (bool, error)](ctx, 0), ixgo.DirectCallArg[[]q.InternalTest](ctx, 1), ixgo.DirectCallArg[[]q.InternalBenchmark](ctx, 2), ixgo.DirectCallArg[[]q.InternalExample](ctx, 3))
}

func method_ptr_PB_Next(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PB).Next(ixgo.DirectCallArg[*q.PB](ctx, 0)))
}

func func_RegisterCover(ctx ixgo.DirectCallContext) {
	q.RegisterCover(ixgo.DirectCallArg[q.Cover](ctx, 0))
}

func func_RunBenchmarks(ctx ixgo.DirectCallContext) {
	q.RunBenchmarks(ixgo.DirectCallArg[func(pat string, str string) (bool, error)](ctx, 0), ixgo.DirectCallArg[[]q.InternalBenchmark](ctx, 1))
}

func func_RunExamples(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RunExamples(ixgo.DirectCallArg[func(pat string, str string) (bool, error)](ctx, 0), ixgo.DirectCallArg[[]q.InternalExample](ctx, 1)))
}

func func_RunTests(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RunTests(ixgo.DirectCallArg[func(pat string, str string) (bool, error)](ctx, 0), ixgo.DirectCallArg[[]q.InternalTest](ctx, 1)))
}

func func_Short(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Short())
}

func method_ptr_T_Chdir(ctx ixgo.DirectCallContext) {
	(*q.T).Chdir(ixgo.DirectCallArg[*q.T](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_T_Parallel(ctx ixgo.DirectCallContext) {
	(*q.T).Parallel(ixgo.DirectCallArg[*q.T](ctx, 0))
}

func method_ptr_T_Run(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.T).Run(ixgo.DirectCallArg[*q.T](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[func(t *q.T)](ctx, 2)))
}

func method_ptr_T_Setenv(ctx ixgo.DirectCallContext) {
	(*q.T).Setenv(ixgo.DirectCallArg[*q.T](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func func_Testing(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Testing())
}

func func_Verbose(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Verbose())
}
