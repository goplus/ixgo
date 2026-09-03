// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package hex

import (
	q "encoding/hex"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/hex", map[string]ixgo.DirectCallAdapter{
		"(*InvalidByteError).Error": method_ptr_InvalidByteError_Error,
		"(InvalidByteError).Error":  method_InvalidByteError_Error,
		"AppendEncode":              func_AppendEncode,
		"DecodedLen":                func_DecodedLen,
		"Dump":                      func_Dump,
		"Dumper":                    func_Dumper,
		"Encode":                    func_Encode,
		"EncodeToString":            func_EncodeToString,
		"EncodedLen":                func_EncodedLen,
		"NewDecoder":                func_NewDecoder,
		"NewEncoder":                func_NewEncoder,
	})
}

func func_AppendEncode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendEncode(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_DecodedLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DecodedLen(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_Dump(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Dump(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_Dumper(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Dumper(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func func_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Encode(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_EncodeToString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EncodeToString(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_EncodedLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EncodedLen(ixgo.DirectCallArg[int](ctx, 0)))
}

func method_InvalidByteError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.InvalidByteError.Error(ixgo.DirectCallArg[q.InvalidByteError](ctx, 0)))
}

func method_ptr_InvalidByteError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InvalidByteError).Error(ixgo.DirectCallArg[*q.InvalidByteError](ctx, 0)))
}

func func_NewDecoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewDecoder(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewEncoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewEncoder(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}
