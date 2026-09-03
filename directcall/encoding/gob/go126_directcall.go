// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package gob

import (
	q "encoding/gob"

	"github.com/goplus/ixgo"
	io "io"
	reflect "reflect"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/gob", map[string]ixgo.DirectCallAdapter{
		"(*Decoder).Decode":      method_ptr_Decoder_Decode,
		"(*Decoder).DecodeValue": method_ptr_Decoder_DecodeValue,
		"(*Encoder).Encode":      method_ptr_Encoder_Encode,
		"(*Encoder).EncodeValue": method_ptr_Encoder_EncodeValue,
		"NewDecoder":             func_NewDecoder,
		"NewEncoder":             func_NewEncoder,
		"Register":               func_Register,
		"RegisterName":           func_RegisterName,
	})
}

func method_ptr_Decoder_Decode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).Decode(ixgo.DirectCallArg[*q.Decoder](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_Decoder_DecodeValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).DecodeValue(ixgo.DirectCallArg[*q.Decoder](ctx, 0), ixgo.DirectCallArg[reflect.Value](ctx, 1)))
}

func method_ptr_Encoder_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).Encode(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_Encoder_EncodeValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).EncodeValue(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[reflect.Value](ctx, 1)))
}

func func_NewDecoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewDecoder(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewEncoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewEncoder(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func func_Register(ctx ixgo.DirectCallContext) {
	q.Register(ixgo.DirectCallArg[any](ctx, 0))
}

func func_RegisterName(ctx ixgo.DirectCallContext) {
	q.RegisterName(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[any](ctx, 1))
}
