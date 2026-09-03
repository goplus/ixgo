// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package gosym

import (
	q "debug/gosym"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("debug/gosym", map[string]ixgo.DirectCallAdapter{
		"(*DecodingError).Error":    method_ptr_DecodingError_Error,
		"(*LineTable).LineToPC":     method_ptr_LineTable_LineToPC,
		"(*LineTable).PCToLine":     method_ptr_LineTable_PCToLine,
		"(*Sym).BaseName":           method_ptr_Sym_BaseName,
		"(*Sym).PackageName":        method_ptr_Sym_PackageName,
		"(*Sym).ReceiverName":       method_ptr_Sym_ReceiverName,
		"(*Sym).Static":             method_ptr_Sym_Static,
		"(*Table).LookupFunc":       method_ptr_Table_LookupFunc,
		"(*Table).LookupSym":        method_ptr_Table_LookupSym,
		"(*Table).PCToFunc":         method_ptr_Table_PCToFunc,
		"(*Table).SymByAddr":        method_ptr_Table_SymByAddr,
		"(*UnknownFileError).Error": method_ptr_UnknownFileError_Error,
		"(*UnknownLineError).Error": method_ptr_UnknownLineError_Error,
		"(UnknownFileError).Error":  method_UnknownFileError_Error,
		"NewLineTable":              func_NewLineTable,
	})
}

func method_ptr_DecodingError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DecodingError).Error(ixgo.DirectCallArg[*q.DecodingError](ctx, 0)))
}

func method_ptr_LineTable_LineToPC(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LineTable).LineToPC(ixgo.DirectCallArg[*q.LineTable](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[uint64](ctx, 2)))
}

func method_ptr_LineTable_PCToLine(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LineTable).PCToLine(ixgo.DirectCallArg[*q.LineTable](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func func_NewLineTable(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewLineTable(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Sym_BaseName(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Sym).BaseName(ixgo.DirectCallArg[*q.Sym](ctx, 0)))
}

func method_ptr_Sym_PackageName(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Sym).PackageName(ixgo.DirectCallArg[*q.Sym](ctx, 0)))
}

func method_ptr_Sym_ReceiverName(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Sym).ReceiverName(ixgo.DirectCallArg[*q.Sym](ctx, 0)))
}

func method_ptr_Sym_Static(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Sym).Static(ixgo.DirectCallArg[*q.Sym](ctx, 0)))
}

func method_ptr_Table_LookupFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Table).LookupFunc(ixgo.DirectCallArg[*q.Table](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Table_LookupSym(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Table).LookupSym(ixgo.DirectCallArg[*q.Table](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Table_PCToFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Table).PCToFunc(ixgo.DirectCallArg[*q.Table](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Table_SymByAddr(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Table).SymByAddr(ixgo.DirectCallArg[*q.Table](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_UnknownFileError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnknownFileError.Error(ixgo.DirectCallArg[q.UnknownFileError](ctx, 0)))
}

func method_ptr_UnknownFileError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnknownFileError).Error(ixgo.DirectCallArg[*q.UnknownFileError](ctx, 0)))
}

func method_ptr_UnknownLineError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnknownLineError).Error(ixgo.DirectCallArg[*q.UnknownLineError](ctx, 0)))
}
