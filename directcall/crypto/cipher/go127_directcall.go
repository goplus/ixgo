// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package cipher

import (
	q "crypto/cipher"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/cipher", map[string]ixgo.DirectCallAdapter{
		"(*StreamWriter).Close": method_ptr_StreamWriter_Close,
		"(StreamWriter).Close":  method_StreamWriter_Close,
		"NewCBCDecrypter":       func_NewCBCDecrypter,
		"NewCBCEncrypter":       func_NewCBCEncrypter,
		"NewCFBDecrypter":       func_NewCFBDecrypter,
		"NewCFBEncrypter":       func_NewCFBEncrypter,
		"NewCTR":                func_NewCTR,
		"NewOFB":                func_NewOFB,
	})
}

func func_NewCBCDecrypter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCBCDecrypter(ixgo.DirectCallArg[q.Block](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_NewCBCEncrypter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCBCEncrypter(ixgo.DirectCallArg[q.Block](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_NewCFBDecrypter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCFBDecrypter(ixgo.DirectCallArg[q.Block](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_NewCFBEncrypter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCFBEncrypter(ixgo.DirectCallArg[q.Block](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_NewCTR(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCTR(ixgo.DirectCallArg[q.Block](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_NewOFB(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewOFB(ixgo.DirectCallArg[q.Block](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_StreamWriter_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StreamWriter.Close(ixgo.DirectCallArg[q.StreamWriter](ctx, 0)))
}

func method_ptr_StreamWriter_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StreamWriter).Close(ixgo.DirectCallArg[*q.StreamWriter](ctx, 0)))
}
