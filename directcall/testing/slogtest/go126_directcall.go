// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package slogtest

import (
	q "testing/slogtest"

	"github.com/goplus/ixgo"
	slog "log/slog"
	testing "testing"
)

func init() {
	ixgo.RegisterDirectCalls("testing/slogtest", map[string]ixgo.DirectCallAdapter{
		"Run":         func_Run,
		"TestHandler": func_TestHandler,
	})
}

func func_Run(ctx ixgo.DirectCallContext) {
	q.Run(ixgo.DirectCallArg[*testing.T](ctx, 0), ixgo.DirectCallArg[func(*testing.T) slog.Handler](ctx, 1), ixgo.DirectCallArg[func(*testing.T) map[string]any](ctx, 2))
}

func func_TestHandler(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TestHandler(ixgo.DirectCallArg[slog.Handler](ctx, 0), ixgo.DirectCallArg[func() []map[string]any](ctx, 1)))
}
