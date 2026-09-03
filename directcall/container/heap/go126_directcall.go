// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package heap

import (
	q "container/heap"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("container/heap", map[string]ixgo.DirectCallAdapter{
		"Fix":    func_Fix,
		"Init":   func_Init,
		"Pop":    func_Pop,
		"Push":   func_Push,
		"Remove": func_Remove,
	})
}

func func_Fix(ctx ixgo.DirectCallContext) {
	q.Fix(ixgo.DirectCallArg[q.Interface](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func func_Init(ctx ixgo.DirectCallContext) {
	q.Init(ixgo.DirectCallArg[q.Interface](ctx, 0))
}

func func_Pop(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pop(ixgo.DirectCallArg[q.Interface](ctx, 0)))
}

func func_Push(ctx ixgo.DirectCallContext) {
	q.Push(ixgo.DirectCallArg[q.Interface](ctx, 0), ixgo.DirectCallArg[any](ctx, 1))
}

func func_Remove(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Remove(ixgo.DirectCallArg[q.Interface](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}
