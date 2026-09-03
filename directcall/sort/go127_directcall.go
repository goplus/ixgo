// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package sort

import (
	q "sort"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("sort", map[string]ixgo.DirectCallAdapter{
		"(*Float64Slice).Len":    method_ptr_Float64Slice_Len,
		"(*Float64Slice).Less":   method_ptr_Float64Slice_Less,
		"(*Float64Slice).Search": method_ptr_Float64Slice_Search,
		"(*Float64Slice).Sort":   method_ptr_Float64Slice_Sort,
		"(*Float64Slice).Swap":   method_ptr_Float64Slice_Swap,
		"(*IntSlice).Len":        method_ptr_IntSlice_Len,
		"(*IntSlice).Less":       method_ptr_IntSlice_Less,
		"(*IntSlice).Search":     method_ptr_IntSlice_Search,
		"(*IntSlice).Sort":       method_ptr_IntSlice_Sort,
		"(*IntSlice).Swap":       method_ptr_IntSlice_Swap,
		"(*StringSlice).Len":     method_ptr_StringSlice_Len,
		"(*StringSlice).Less":    method_ptr_StringSlice_Less,
		"(*StringSlice).Search":  method_ptr_StringSlice_Search,
		"(*StringSlice).Sort":    method_ptr_StringSlice_Sort,
		"(*StringSlice).Swap":    method_ptr_StringSlice_Swap,
		"(Float64Slice).Len":     method_Float64Slice_Len,
		"(Float64Slice).Less":    method_Float64Slice_Less,
		"(Float64Slice).Search":  method_Float64Slice_Search,
		"(Float64Slice).Sort":    method_Float64Slice_Sort,
		"(Float64Slice).Swap":    method_Float64Slice_Swap,
		"(IntSlice).Len":         method_IntSlice_Len,
		"(IntSlice).Less":        method_IntSlice_Less,
		"(IntSlice).Search":      method_IntSlice_Search,
		"(IntSlice).Sort":        method_IntSlice_Sort,
		"(IntSlice).Swap":        method_IntSlice_Swap,
		"(StringSlice).Len":      method_StringSlice_Len,
		"(StringSlice).Less":     method_StringSlice_Less,
		"(StringSlice).Search":   method_StringSlice_Search,
		"(StringSlice).Sort":     method_StringSlice_Sort,
		"(StringSlice).Swap":     method_StringSlice_Swap,
		"Float64s":               func_Float64s,
		"Float64sAreSorted":      func_Float64sAreSorted,
		"Ints":                   func_Ints,
		"IntsAreSorted":          func_IntsAreSorted,
		"IsSorted":               func_IsSorted,
		"Reverse":                func_Reverse,
		"Search":                 func_Search,
		"SearchFloat64s":         func_SearchFloat64s,
		"SearchInts":             func_SearchInts,
		"SearchStrings":          func_SearchStrings,
		"Slice":                  func_Slice,
		"SliceIsSorted":          func_SliceIsSorted,
		"SliceStable":            func_SliceStable,
		"Sort":                   func_Sort,
		"Stable":                 func_Stable,
		"Strings":                func_Strings,
		"StringsAreSorted":       func_StringsAreSorted,
	})
}

func method_Float64Slice_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float64Slice.Len(ixgo.DirectCallArg[q.Float64Slice](ctx, 0)))
}

func method_ptr_Float64Slice_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float64Slice).Len(ixgo.DirectCallArg[*q.Float64Slice](ctx, 0)))
}

func method_Float64Slice_Less(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float64Slice.Less(ixgo.DirectCallArg[q.Float64Slice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Float64Slice_Less(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float64Slice).Less(ixgo.DirectCallArg[*q.Float64Slice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_Float64Slice_Search(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float64Slice.Search(ixgo.DirectCallArg[q.Float64Slice](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func method_ptr_Float64Slice_Search(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float64Slice).Search(ixgo.DirectCallArg[*q.Float64Slice](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func method_Float64Slice_Sort(ctx ixgo.DirectCallContext) {
	q.Float64Slice.Sort(ixgo.DirectCallArg[q.Float64Slice](ctx, 0))
}

func method_ptr_Float64Slice_Sort(ctx ixgo.DirectCallContext) {
	(*q.Float64Slice).Sort(ixgo.DirectCallArg[*q.Float64Slice](ctx, 0))
}

func method_Float64Slice_Swap(ctx ixgo.DirectCallContext) {
	q.Float64Slice.Swap(ixgo.DirectCallArg[q.Float64Slice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func method_ptr_Float64Slice_Swap(ctx ixgo.DirectCallContext) {
	(*q.Float64Slice).Swap(ixgo.DirectCallArg[*q.Float64Slice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func func_Float64s(ctx ixgo.DirectCallContext) {
	q.Float64s(ixgo.DirectCallArg[[]float64](ctx, 0))
}

func func_Float64sAreSorted(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float64sAreSorted(ixgo.DirectCallArg[[]float64](ctx, 0)))
}

func method_IntSlice_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IntSlice.Len(ixgo.DirectCallArg[q.IntSlice](ctx, 0)))
}

func method_ptr_IntSlice_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IntSlice).Len(ixgo.DirectCallArg[*q.IntSlice](ctx, 0)))
}

func method_IntSlice_Less(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IntSlice.Less(ixgo.DirectCallArg[q.IntSlice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_IntSlice_Less(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IntSlice).Less(ixgo.DirectCallArg[*q.IntSlice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_IntSlice_Search(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IntSlice.Search(ixgo.DirectCallArg[q.IntSlice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_IntSlice_Search(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IntSlice).Search(ixgo.DirectCallArg[*q.IntSlice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_IntSlice_Sort(ctx ixgo.DirectCallContext) {
	q.IntSlice.Sort(ixgo.DirectCallArg[q.IntSlice](ctx, 0))
}

func method_ptr_IntSlice_Sort(ctx ixgo.DirectCallContext) {
	(*q.IntSlice).Sort(ixgo.DirectCallArg[*q.IntSlice](ctx, 0))
}

func method_IntSlice_Swap(ctx ixgo.DirectCallContext) {
	q.IntSlice.Swap(ixgo.DirectCallArg[q.IntSlice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func method_ptr_IntSlice_Swap(ctx ixgo.DirectCallContext) {
	(*q.IntSlice).Swap(ixgo.DirectCallArg[*q.IntSlice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func func_Ints(ctx ixgo.DirectCallContext) {
	q.Ints(ixgo.DirectCallArg[[]int](ctx, 0))
}

func func_IntsAreSorted(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IntsAreSorted(ixgo.DirectCallArg[[]int](ctx, 0)))
}

func func_IsSorted(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsSorted(ixgo.DirectCallArg[q.Interface](ctx, 0)))
}

func func_Reverse(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Reverse(ixgo.DirectCallArg[q.Interface](ctx, 0)))
}

func func_Search(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Search(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[func(int) bool](ctx, 1)))
}

func func_SearchFloat64s(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SearchFloat64s(ixgo.DirectCallArg[[]float64](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func func_SearchInts(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SearchInts(ixgo.DirectCallArg[[]int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_SearchStrings(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SearchStrings(ixgo.DirectCallArg[[]string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Slice(ctx ixgo.DirectCallContext) {
	q.Slice(ixgo.DirectCallArg[interface{}](ctx, 0), ixgo.DirectCallArg[func(i int, j int) bool](ctx, 1))
}

func func_SliceIsSorted(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SliceIsSorted(ixgo.DirectCallArg[interface{}](ctx, 0), ixgo.DirectCallArg[func(i int, j int) bool](ctx, 1)))
}

func func_SliceStable(ctx ixgo.DirectCallContext) {
	q.SliceStable(ixgo.DirectCallArg[interface{}](ctx, 0), ixgo.DirectCallArg[func(i int, j int) bool](ctx, 1))
}

func func_Sort(ctx ixgo.DirectCallContext) {
	q.Sort(ixgo.DirectCallArg[q.Interface](ctx, 0))
}

func func_Stable(ctx ixgo.DirectCallContext) {
	q.Stable(ixgo.DirectCallArg[q.Interface](ctx, 0))
}

func method_StringSlice_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StringSlice.Len(ixgo.DirectCallArg[q.StringSlice](ctx, 0)))
}

func method_ptr_StringSlice_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StringSlice).Len(ixgo.DirectCallArg[*q.StringSlice](ctx, 0)))
}

func method_StringSlice_Less(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StringSlice.Less(ixgo.DirectCallArg[q.StringSlice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_StringSlice_Less(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StringSlice).Less(ixgo.DirectCallArg[*q.StringSlice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_StringSlice_Search(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StringSlice.Search(ixgo.DirectCallArg[q.StringSlice](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_StringSlice_Search(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StringSlice).Search(ixgo.DirectCallArg[*q.StringSlice](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_StringSlice_Sort(ctx ixgo.DirectCallContext) {
	q.StringSlice.Sort(ixgo.DirectCallArg[q.StringSlice](ctx, 0))
}

func method_ptr_StringSlice_Sort(ctx ixgo.DirectCallContext) {
	(*q.StringSlice).Sort(ixgo.DirectCallArg[*q.StringSlice](ctx, 0))
}

func method_StringSlice_Swap(ctx ixgo.DirectCallContext) {
	q.StringSlice.Swap(ixgo.DirectCallArg[q.StringSlice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func method_ptr_StringSlice_Swap(ctx ixgo.DirectCallContext) {
	(*q.StringSlice).Swap(ixgo.DirectCallArg[*q.StringSlice](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func func_Strings(ctx ixgo.DirectCallContext) {
	q.Strings(ixgo.DirectCallArg[[]string](ctx, 0))
}

func func_StringsAreSorted(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StringsAreSorted(ixgo.DirectCallArg[[]string](ctx, 0)))
}
