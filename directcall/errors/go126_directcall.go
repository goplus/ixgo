// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package errors

import (
	q "errors"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("errors", map[string]ixgo.DirectCallAdapter{
		"As":     func_As,
		"Is":     func_Is,
		"Join":   func_Join,
		"New":    func_New,
		"Unwrap": func_Unwrap,
	})
}

func func_As(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.As(ixgo.DirectCallArg[error](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func func_Is(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Is(ixgo.DirectCallArg[error](ctx, 0), ixgo.DirectCallArg[error](ctx, 1)))
}

func func_Join(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Join(ixgo.DirectCallArg[[]error](ctx, 0)...))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Unwrap(ixgo.DirectCallArg[error](ctx, 0)))
}
