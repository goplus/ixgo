// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package pem

import (
	q "encoding/pem"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/pem", map[string]ixgo.DirectCallAdapter{
		"Encode":         func_Encode,
		"EncodeToMemory": func_EncodeToMemory,
	})
}

func func_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Encode(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[*q.Block](ctx, 1)))
}

func func_EncodeToMemory(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EncodeToMemory(ixgo.DirectCallArg[*q.Block](ctx, 0)))
}
