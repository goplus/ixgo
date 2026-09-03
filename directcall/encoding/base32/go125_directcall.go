// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package base32

import (
	q "encoding/base32"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/base32", map[string]ixgo.DirectCallAdapter{
		"(*CorruptInputError).Error": method_ptr_CorruptInputError_Error,
		"(*Encoding).AppendEncode":   method_ptr_Encoding_AppendEncode,
		"(*Encoding).DecodedLen":     method_ptr_Encoding_DecodedLen,
		"(*Encoding).Encode":         method_ptr_Encoding_Encode,
		"(*Encoding).EncodeToString": method_ptr_Encoding_EncodeToString,
		"(*Encoding).EncodedLen":     method_ptr_Encoding_EncodedLen,
		"(*Encoding).WithPadding":    method_ptr_Encoding_WithPadding,
		"(CorruptInputError).Error":  method_CorruptInputError_Error,
		"(Encoding).WithPadding":     method_Encoding_WithPadding,
		"NewDecoder":                 func_NewDecoder,
		"NewEncoder":                 func_NewEncoder,
		"NewEncoding":                func_NewEncoding,
	})
}

func method_CorruptInputError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CorruptInputError.Error(ixgo.DirectCallArg[q.CorruptInputError](ctx, 0)))
}

func method_ptr_CorruptInputError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CorruptInputError).Error(ixgo.DirectCallArg[*q.CorruptInputError](ctx, 0)))
}

func method_ptr_Encoding_AppendEncode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoding).AppendEncode(ixgo.DirectCallArg[*q.Encoding](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}

func method_ptr_Encoding_DecodedLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoding).DecodedLen(ixgo.DirectCallArg[*q.Encoding](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Encoding_Encode(ctx ixgo.DirectCallContext) {
	(*q.Encoding).Encode(ixgo.DirectCallArg[*q.Encoding](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2))
}

func method_ptr_Encoding_EncodeToString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoding).EncodeToString(ixgo.DirectCallArg[*q.Encoding](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Encoding_EncodedLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoding).EncodedLen(ixgo.DirectCallArg[*q.Encoding](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_Encoding_WithPadding(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Encoding.WithPadding(ixgo.DirectCallArg[q.Encoding](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func method_ptr_Encoding_WithPadding(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoding).WithPadding(ixgo.DirectCallArg[*q.Encoding](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_NewDecoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewDecoder(ixgo.DirectCallArg[*q.Encoding](ctx, 0), ixgo.DirectCallArg[io.Reader](ctx, 1)))
}

func func_NewEncoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewEncoder(ixgo.DirectCallArg[*q.Encoding](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1)))
}

func func_NewEncoding(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewEncoding(ixgo.DirectCallArg[string](ctx, 0)))
}
