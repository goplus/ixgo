// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package scanner

import (
	q "go/scanner"

	"github.com/goplus/ixgo"
	token "go/token"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("go/scanner", map[string]ixgo.DirectCallAdapter{
		"(*Error).Error":               method_ptr_Error_Error,
		"(*ErrorList).Add":             method_ptr_ErrorList_Add,
		"(*ErrorList).Err":             method_ptr_ErrorList_Err,
		"(*ErrorList).Error":           method_ptr_ErrorList_Error,
		"(*ErrorList).Len":             method_ptr_ErrorList_Len,
		"(*ErrorList).Less":            method_ptr_ErrorList_Less,
		"(*ErrorList).RemoveMultiples": method_ptr_ErrorList_RemoveMultiples,
		"(*ErrorList).Reset":           method_ptr_ErrorList_Reset,
		"(*ErrorList).Sort":            method_ptr_ErrorList_Sort,
		"(*ErrorList).Swap":            method_ptr_ErrorList_Swap,
		"(*Scanner).End":               method_ptr_Scanner_End,
		"(*Scanner).Init":              method_ptr_Scanner_Init,
		"(Error).Error":                method_Error_Error,
		"(ErrorList).Err":              method_ErrorList_Err,
		"(ErrorList).Error":            method_ErrorList_Error,
		"(ErrorList).Len":              method_ErrorList_Len,
		"(ErrorList).Less":             method_ErrorList_Less,
		"(ErrorList).Sort":             method_ErrorList_Sort,
		"(ErrorList).Swap":             method_ErrorList_Swap,
		"PrintError":                   func_PrintError,
	})
}

func method_Error_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Error.Error(ixgo.DirectCallArg[q.Error](ctx, 0)))
}

func method_ptr_Error_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Error).Error(ixgo.DirectCallArg[*q.Error](ctx, 0)))
}

func method_ptr_ErrorList_Add(ctx ixgo.DirectCallContext) {
	(*q.ErrorList).Add(ixgo.DirectCallArg[*q.ErrorList](ctx, 0), ixgo.DirectCallArg[token.Position](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ErrorList_Err(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ErrorList.Err(ixgo.DirectCallArg[q.ErrorList](ctx, 0)))
}

func method_ptr_ErrorList_Err(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ErrorList).Err(ixgo.DirectCallArg[*q.ErrorList](ctx, 0)))
}

func method_ErrorList_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ErrorList.Error(ixgo.DirectCallArg[q.ErrorList](ctx, 0)))
}

func method_ptr_ErrorList_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ErrorList).Error(ixgo.DirectCallArg[*q.ErrorList](ctx, 0)))
}

func method_ErrorList_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ErrorList.Len(ixgo.DirectCallArg[q.ErrorList](ctx, 0)))
}

func method_ptr_ErrorList_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ErrorList).Len(ixgo.DirectCallArg[*q.ErrorList](ctx, 0)))
}

func method_ErrorList_Less(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ErrorList.Less(ixgo.DirectCallArg[q.ErrorList](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_ErrorList_Less(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ErrorList).Less(ixgo.DirectCallArg[*q.ErrorList](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_ErrorList_RemoveMultiples(ctx ixgo.DirectCallContext) {
	(*q.ErrorList).RemoveMultiples(ixgo.DirectCallArg[*q.ErrorList](ctx, 0))
}

func method_ptr_ErrorList_Reset(ctx ixgo.DirectCallContext) {
	(*q.ErrorList).Reset(ixgo.DirectCallArg[*q.ErrorList](ctx, 0))
}

func method_ErrorList_Sort(ctx ixgo.DirectCallContext) {
	q.ErrorList.Sort(ixgo.DirectCallArg[q.ErrorList](ctx, 0))
}

func method_ptr_ErrorList_Sort(ctx ixgo.DirectCallContext) {
	(*q.ErrorList).Sort(ixgo.DirectCallArg[*q.ErrorList](ctx, 0))
}

func method_ErrorList_Swap(ctx ixgo.DirectCallContext) {
	q.ErrorList.Swap(ixgo.DirectCallArg[q.ErrorList](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func method_ptr_ErrorList_Swap(ctx ixgo.DirectCallContext) {
	(*q.ErrorList).Swap(ixgo.DirectCallArg[*q.ErrorList](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func func_PrintError(ctx ixgo.DirectCallContext) {
	q.PrintError(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[error](ctx, 1))
}

func method_ptr_Scanner_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).End(ixgo.DirectCallArg[*q.Scanner](ctx, 0)))
}

func method_ptr_Scanner_Init(ctx ixgo.DirectCallContext) {
	(*q.Scanner).Init(ixgo.DirectCallArg[*q.Scanner](ctx, 0), ixgo.DirectCallArg[*token.File](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2), ixgo.DirectCallArg[q.ErrorHandler](ctx, 3), ixgo.DirectCallArg[q.Mode](ctx, 4))
}
