// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package ascii85

import (
	q "encoding/ascii85"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/ascii85", map[string]ixgo.DirectCallAdapter{
		"(*CorruptInputError).Error": method_ptr_CorruptInputError_Error,
		"(CorruptInputError).Error":  method_CorruptInputError_Error,
		"Encode":                     func_Encode,
		"MaxEncodedLen":              func_MaxEncodedLen,
		"NewDecoder":                 func_NewDecoder,
		"NewEncoder":                 func_NewEncoder,
	})
}

func method_CorruptInputError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CorruptInputError.Error(ixgo.DirectCallArg[q.CorruptInputError](ctx, 0)))
}

func method_ptr_CorruptInputError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CorruptInputError).Error(ixgo.DirectCallArg[*q.CorruptInputError](ctx, 0)))
}

func func_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Encode(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_MaxEncodedLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MaxEncodedLen(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_NewDecoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewDecoder(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewEncoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewEncoder(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}
