// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package metrics

import (
	q "runtime/metrics"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("runtime/metrics", map[string]ixgo.DirectCallAdapter{
		"(*Value).Float64":          method_ptr_Value_Float64,
		"(*Value).Float64Histogram": method_ptr_Value_Float64Histogram,
		"(*Value).Kind":             method_ptr_Value_Kind,
		"(*Value).Uint64":           method_ptr_Value_Uint64,
		"(Value).Float64":           method_Value_Float64,
		"(Value).Float64Histogram":  method_Value_Float64Histogram,
		"(Value).Kind":              method_Value_Kind,
		"(Value).Uint64":            method_Value_Uint64,
		"All":                       func_All,
		"Read":                      func_Read,
	})
}

func func_All(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.All())
}

func func_Read(ctx ixgo.DirectCallContext) {
	q.Read(ixgo.DirectCallArg[[]q.Sample](ctx, 0))
}

func method_Value_Float64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Float64(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Float64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Float64(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Float64Histogram(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Float64Histogram(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Float64Histogram(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Float64Histogram(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Kind(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Kind(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Uint64(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Uint64(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}
