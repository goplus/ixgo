// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package rpc

import (
	q "net/rpc"

	"github.com/goplus/ixgo"
	io "io"
	net "net"
	http "net/http"
)

func init() {
	ixgo.RegisterDirectCalls("net/rpc", map[string]ixgo.DirectCallAdapter{
		"(*Client).Call":         method_ptr_Client_Call,
		"(*Client).Close":        method_ptr_Client_Close,
		"(*Client).Go":           method_ptr_Client_Go,
		"(*Server).Accept":       method_ptr_Server_Accept,
		"(*Server).HandleHTTP":   method_ptr_Server_HandleHTTP,
		"(*Server).Register":     method_ptr_Server_Register,
		"(*Server).RegisterName": method_ptr_Server_RegisterName,
		"(*Server).ServeCodec":   method_ptr_Server_ServeCodec,
		"(*Server).ServeConn":    method_ptr_Server_ServeConn,
		"(*Server).ServeHTTP":    method_ptr_Server_ServeHTTP,
		"(*Server).ServeRequest": method_ptr_Server_ServeRequest,
		"(*ServerError).Error":   method_ptr_ServerError_Error,
		"(ServerError).Error":    method_ServerError_Error,
		"Accept":                 func_Accept,
		"HandleHTTP":             func_HandleHTTP,
		"NewClient":              func_NewClient,
		"NewClientWithCodec":     func_NewClientWithCodec,
		"NewServer":              func_NewServer,
		"Register":               func_Register,
		"RegisterName":           func_RegisterName,
		"ServeCodec":             func_ServeCodec,
		"ServeConn":              func_ServeConn,
		"ServeRequest":           func_ServeRequest,
	})
}

func func_Accept(ctx ixgo.DirectCallContext) {
	q.Accept(ixgo.DirectCallArg[net.Listener](ctx, 0))
}

func method_ptr_Client_Call(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Call(ixgo.DirectCallArg[*q.Client](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[any](ctx, 2), ixgo.DirectCallArg[any](ctx, 3)))
}

func method_ptr_Client_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Close(ixgo.DirectCallArg[*q.Client](ctx, 0)))
}

func method_ptr_Client_Go(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Go(ixgo.DirectCallArg[*q.Client](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[any](ctx, 2), ixgo.DirectCallArg[any](ctx, 3), ixgo.DirectCallArg[chan *q.Call](ctx, 4)))
}

func func_HandleHTTP(ctx ixgo.DirectCallContext) {
	q.HandleHTTP()
}

func func_NewClient(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewClient(ixgo.DirectCallArg[io.ReadWriteCloser](ctx, 0)))
}

func func_NewClientWithCodec(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewClientWithCodec(ixgo.DirectCallArg[q.ClientCodec](ctx, 0)))
}

func func_NewServer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewServer())
}

func func_Register(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Register(ixgo.DirectCallArg[any](ctx, 0)))
}

func func_RegisterName(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RegisterName(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func func_ServeCodec(ctx ixgo.DirectCallContext) {
	q.ServeCodec(ixgo.DirectCallArg[q.ServerCodec](ctx, 0))
}

func func_ServeConn(ctx ixgo.DirectCallContext) {
	q.ServeConn(ixgo.DirectCallArg[io.ReadWriteCloser](ctx, 0))
}

func func_ServeRequest(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ServeRequest(ixgo.DirectCallArg[q.ServerCodec](ctx, 0)))
}

func method_ptr_Server_Accept(ctx ixgo.DirectCallContext) {
	(*q.Server).Accept(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[net.Listener](ctx, 1))
}

func method_ptr_Server_HandleHTTP(ctx ixgo.DirectCallContext) {
	(*q.Server).HandleHTTP(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ptr_Server_Register(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).Register(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_Server_RegisterName(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).RegisterName(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[any](ctx, 2)))
}

func method_ptr_Server_ServeCodec(ctx ixgo.DirectCallContext) {
	(*q.Server).ServeCodec(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[q.ServerCodec](ctx, 1))
}

func method_ptr_Server_ServeConn(ctx ixgo.DirectCallContext) {
	(*q.Server).ServeConn(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[io.ReadWriteCloser](ctx, 1))
}

func method_ptr_Server_ServeHTTP(ctx ixgo.DirectCallContext) {
	(*q.Server).ServeHTTP(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[http.ResponseWriter](ctx, 1), ixgo.DirectCallArg[*http.Request](ctx, 2))
}

func method_ptr_Server_ServeRequest(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).ServeRequest(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[q.ServerCodec](ctx, 1)))
}

func method_ServerError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ServerError.Error(ixgo.DirectCallArg[q.ServerError](ctx, 0)))
}

func method_ptr_ServerError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ServerError).Error(ixgo.DirectCallArg[*q.ServerError](ctx, 0)))
}
