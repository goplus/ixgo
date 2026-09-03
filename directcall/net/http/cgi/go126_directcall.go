// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package cgi

import (
	q "net/http/cgi"

	"github.com/goplus/ixgo"
	http "net/http"
)

func init() {
	ixgo.RegisterDirectCalls("net/http/cgi", map[string]ixgo.DirectCallAdapter{
		"(*Handler).ServeHTTP": method_ptr_Handler_ServeHTTP,
		"Serve":                func_Serve,
	})
}

func method_ptr_Handler_ServeHTTP(ctx ixgo.DirectCallContext) {
	(*q.Handler).ServeHTTP(ixgo.DirectCallArg[*q.Handler](ctx, 0), ixgo.DirectCallArg[http.ResponseWriter](ctx, 1), ixgo.DirectCallArg[*http.Request](ctx, 2))
}

func func_Serve(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Serve(ixgo.DirectCallArg[http.Handler](ctx, 0)))
}
