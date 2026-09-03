// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package bzip2

import (
	q "compress/bzip2"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("compress/bzip2", map[string]ixgo.DirectCallAdapter{
		"(*StructuralError).Error": method_ptr_StructuralError_Error,
		"(StructuralError).Error":  method_StructuralError_Error,
		"NewReader":                func_NewReader,
	})
}

func func_NewReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func method_StructuralError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StructuralError.Error(ixgo.DirectCallArg[q.StructuralError](ctx, 0)))
}

func method_ptr_StructuralError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StructuralError).Error(ixgo.DirectCallArg[*q.StructuralError](ctx, 0)))
}
