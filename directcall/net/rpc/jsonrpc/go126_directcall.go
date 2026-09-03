// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package jsonrpc

import (
	q "net/rpc/jsonrpc"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("net/rpc/jsonrpc", map[string]ixgo.DirectCallAdapter{
		"NewClient":      func_NewClient,
		"NewClientCodec": func_NewClientCodec,
		"NewServerCodec": func_NewServerCodec,
		"ServeConn":      func_ServeConn,
	})
}

func func_NewClient(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewClient(ixgo.DirectCallArg[io.ReadWriteCloser](ctx, 0)))
}

func func_NewClientCodec(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewClientCodec(ixgo.DirectCallArg[io.ReadWriteCloser](ctx, 0)))
}

func func_NewServerCodec(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewServerCodec(ixgo.DirectCallArg[io.ReadWriteCloser](ctx, 0)))
}

func func_ServeConn(ctx ixgo.DirectCallContext) {
	q.ServeConn(ixgo.DirectCallArg[io.ReadWriteCloser](ctx, 0))
}
