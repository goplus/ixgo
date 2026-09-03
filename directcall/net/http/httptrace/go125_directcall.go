// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package httptrace

import (
	q "net/http/httptrace"

	context "context"
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("net/http/httptrace", map[string]ixgo.DirectCallAdapter{
		"ContextClientTrace": func_ContextClientTrace,
		"WithClientTrace":    func_WithClientTrace,
	})
}

func func_ContextClientTrace(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ContextClientTrace(ixgo.DirectCallArg[context.Context](ctx, 0)))
}

func func_WithClientTrace(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WithClientTrace(ixgo.DirectCallArg[context.Context](ctx, 0), ixgo.DirectCallArg[*q.ClientTrace](ctx, 1)))
}
