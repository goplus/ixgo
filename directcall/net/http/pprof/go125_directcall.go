// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package pprof

import (
	q "net/http/pprof"

	"github.com/goplus/ixgo"
	http "net/http"
)

func init() {
	ixgo.RegisterDirectCalls("net/http/pprof", map[string]ixgo.DirectCallAdapter{
		"Cmdline": func_Cmdline,
		"Handler": func_Handler,
		"Index":   func_Index,
		"Profile": func_Profile,
		"Symbol":  func_Symbol,
		"Trace":   func_Trace,
	})
}

func func_Cmdline(ctx ixgo.DirectCallContext) {
	q.Cmdline(ixgo.DirectCallArg[http.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*http.Request](ctx, 1))
}

func func_Handler(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Handler(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Index(ctx ixgo.DirectCallContext) {
	q.Index(ixgo.DirectCallArg[http.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*http.Request](ctx, 1))
}

func func_Profile(ctx ixgo.DirectCallContext) {
	q.Profile(ixgo.DirectCallArg[http.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*http.Request](ctx, 1))
}

func func_Symbol(ctx ixgo.DirectCallContext) {
	q.Symbol(ixgo.DirectCallArg[http.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*http.Request](ctx, 1))
}

func func_Trace(ctx ixgo.DirectCallContext) {
	q.Trace(ixgo.DirectCallArg[http.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*http.Request](ctx, 1))
}
