// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package suffixarray

import (
	q "index/suffixarray"

	"github.com/goplus/ixgo"
	io "io"
	regexp "regexp"
)

func init() {
	ixgo.RegisterDirectCalls("index/suffixarray", map[string]ixgo.DirectCallAdapter{
		"(*Index).Bytes":        method_ptr_Index_Bytes,
		"(*Index).FindAllIndex": method_ptr_Index_FindAllIndex,
		"(*Index).Lookup":       method_ptr_Index_Lookup,
		"(*Index).Read":         method_ptr_Index_Read,
		"(*Index).Write":        method_ptr_Index_Write,
		"New":                   func_New,
	})
}

func method_ptr_Index_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Index).Bytes(ixgo.DirectCallArg[*q.Index](ctx, 0)))
}

func method_ptr_Index_FindAllIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Index).FindAllIndex(ixgo.DirectCallArg[*q.Index](ctx, 0), ixgo.DirectCallArg[*regexp.Regexp](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Index_Lookup(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Index).Lookup(ixgo.DirectCallArg[*q.Index](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Index_Read(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Index).Read(ixgo.DirectCallArg[*q.Index](ctx, 0), ixgo.DirectCallArg[io.Reader](ctx, 1)))
}

func method_ptr_Index_Write(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Index).Write(ixgo.DirectCallArg[*q.Index](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[[]byte](ctx, 0)))
}
