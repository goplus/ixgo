// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package fcgi

import (
	q "net/http/fcgi"

	"github.com/goplus/ixgo"
	net "net"
	http "net/http"
)

func init() {
	ixgo.RegisterDirectCalls("net/http/fcgi", map[string]ixgo.DirectCallAdapter{
		"ProcessEnv": func_ProcessEnv,
		"Serve":      func_Serve,
	})
}

func func_ProcessEnv(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ProcessEnv(ixgo.DirectCallArg[*http.Request](ctx, 0)))
}

func func_Serve(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Serve(ixgo.DirectCallArg[net.Listener](ctx, 0), ixgo.DirectCallArg[http.Handler](ctx, 1)))
}
