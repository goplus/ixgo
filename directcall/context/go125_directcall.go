// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package context

import (
	q "context"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("context", map[string]ixgo.DirectCallAdapter{
		"AfterFunc":     func_AfterFunc,
		"Background":    func_Background,
		"Cause":         func_Cause,
		"TODO":          func_TODO,
		"WithValue":     func_WithValue,
		"WithoutCancel": func_WithoutCancel,
	})
}

func func_AfterFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AfterFunc(ixgo.DirectCallArg[q.Context](ctx, 0), ixgo.DirectCallArg[func()](ctx, 1)))
}

func func_Background(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Background())
}

func func_Cause(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Cause(ixgo.DirectCallArg[q.Context](ctx, 0)))
}

func func_TODO(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TODO())
}

func func_WithValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WithValue(ixgo.DirectCallArg[q.Context](ctx, 0), ixgo.DirectCallArg[any](ctx, 1), ixgo.DirectCallArg[any](ctx, 2)))
}

func func_WithoutCancel(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WithoutCancel(ixgo.DirectCallArg[q.Context](ctx, 0)))
}
