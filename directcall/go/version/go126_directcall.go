// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package version

import (
	q "go/version"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("go/version", map[string]ixgo.DirectCallAdapter{
		"Compare": func_Compare,
		"IsValid": func_IsValid,
		"Lang":    func_Lang,
	})
}

func func_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Compare(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsValid(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Lang(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Lang(ixgo.DirectCallArg[string](ctx, 0)))
}
