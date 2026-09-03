// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package httputil

import (
	q "net/http/httputil"

	bufio "bufio"
	"github.com/goplus/ixgo"
	io "io"
	net "net"
	http "net/http"
	url "net/url"
)

func init() {
	ixgo.RegisterDirectCalls("net/http/httputil", map[string]ixgo.DirectCallAdapter{
		"(*ClientConn).Close":           method_ptr_ClientConn_Close,
		"(*ClientConn).Pending":         method_ptr_ClientConn_Pending,
		"(*ClientConn).Write":           method_ptr_ClientConn_Write,
		"(*ProxyRequest).SetURL":        method_ptr_ProxyRequest_SetURL,
		"(*ProxyRequest).SetXForwarded": method_ptr_ProxyRequest_SetXForwarded,
		"(*ReverseProxy).ServeHTTP":     method_ptr_ReverseProxy_ServeHTTP,
		"(*ServerConn).Close":           method_ptr_ServerConn_Close,
		"(*ServerConn).Pending":         method_ptr_ServerConn_Pending,
		"(*ServerConn).Write":           method_ptr_ServerConn_Write,
		"NewChunkedReader":              func_NewChunkedReader,
		"NewChunkedWriter":              func_NewChunkedWriter,
		"NewClientConn":                 func_NewClientConn,
		"NewProxyClientConn":            func_NewProxyClientConn,
		"NewServerConn":                 func_NewServerConn,
		"NewSingleHostReverseProxy":     func_NewSingleHostReverseProxy,
	})
}

func method_ptr_ClientConn_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientConn).Close(ixgo.DirectCallArg[*q.ClientConn](ctx, 0)))
}

func method_ptr_ClientConn_Pending(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientConn).Pending(ixgo.DirectCallArg[*q.ClientConn](ctx, 0)))
}

func method_ptr_ClientConn_Write(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientConn).Write(ixgo.DirectCallArg[*q.ClientConn](ctx, 0), ixgo.DirectCallArg[*http.Request](ctx, 1)))
}

func func_NewChunkedReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewChunkedReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewChunkedWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewChunkedWriter(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func func_NewClientConn(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewClientConn(ixgo.DirectCallArg[net.Conn](ctx, 0), ixgo.DirectCallArg[*bufio.Reader](ctx, 1)))
}

func func_NewProxyClientConn(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewProxyClientConn(ixgo.DirectCallArg[net.Conn](ctx, 0), ixgo.DirectCallArg[*bufio.Reader](ctx, 1)))
}

func func_NewServerConn(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewServerConn(ixgo.DirectCallArg[net.Conn](ctx, 0), ixgo.DirectCallArg[*bufio.Reader](ctx, 1)))
}

func func_NewSingleHostReverseProxy(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewSingleHostReverseProxy(ixgo.DirectCallArg[*url.URL](ctx, 0)))
}

func method_ptr_ProxyRequest_SetURL(ctx ixgo.DirectCallContext) {
	(*q.ProxyRequest).SetURL(ixgo.DirectCallArg[*q.ProxyRequest](ctx, 0), ixgo.DirectCallArg[*url.URL](ctx, 1))
}

func method_ptr_ProxyRequest_SetXForwarded(ctx ixgo.DirectCallContext) {
	(*q.ProxyRequest).SetXForwarded(ixgo.DirectCallArg[*q.ProxyRequest](ctx, 0))
}

func method_ptr_ReverseProxy_ServeHTTP(ctx ixgo.DirectCallContext) {
	(*q.ReverseProxy).ServeHTTP(ixgo.DirectCallArg[*q.ReverseProxy](ctx, 0), ixgo.DirectCallArg[http.ResponseWriter](ctx, 1), ixgo.DirectCallArg[*http.Request](ctx, 2))
}

func method_ptr_ServerConn_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ServerConn).Close(ixgo.DirectCallArg[*q.ServerConn](ctx, 0)))
}

func method_ptr_ServerConn_Pending(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ServerConn).Pending(ixgo.DirectCallArg[*q.ServerConn](ctx, 0)))
}

func method_ptr_ServerConn_Write(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ServerConn).Write(ixgo.DirectCallArg[*q.ServerConn](ctx, 0), ixgo.DirectCallArg[*http.Request](ctx, 1), ixgo.DirectCallArg[*http.Response](ctx, 2)))
}
