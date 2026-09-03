// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package ring

import (
	q "container/ring"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("container/ring", map[string]ixgo.DirectCallAdapter{
		"(*Ring).Do":     method_ptr_Ring_Do,
		"(*Ring).Len":    method_ptr_Ring_Len,
		"(*Ring).Link":   method_ptr_Ring_Link,
		"(*Ring).Move":   method_ptr_Ring_Move,
		"(*Ring).Next":   method_ptr_Ring_Next,
		"(*Ring).Prev":   method_ptr_Ring_Prev,
		"(*Ring).Unlink": method_ptr_Ring_Unlink,
		"New":            func_New,
	})
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[int](ctx, 0)))
}

func method_ptr_Ring_Do(ctx ixgo.DirectCallContext) {
	(*q.Ring).Do(ixgo.DirectCallArg[*q.Ring](ctx, 0), ixgo.DirectCallArg[func(any)](ctx, 1))
}

func method_ptr_Ring_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ring).Len(ixgo.DirectCallArg[*q.Ring](ctx, 0)))
}

func method_ptr_Ring_Link(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ring).Link(ixgo.DirectCallArg[*q.Ring](ctx, 0), ixgo.DirectCallArg[*q.Ring](ctx, 1)))
}

func method_ptr_Ring_Move(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ring).Move(ixgo.DirectCallArg[*q.Ring](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Ring_Next(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ring).Next(ixgo.DirectCallArg[*q.Ring](ctx, 0)))
}

func method_ptr_Ring_Prev(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ring).Prev(ixgo.DirectCallArg[*q.Ring](ctx, 0)))
}

func method_ptr_Ring_Unlink(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ring).Unlink(ixgo.DirectCallArg[*q.Ring](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}
