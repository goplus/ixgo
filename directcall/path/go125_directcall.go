// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package path

import (
	q "path"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("path", map[string]ixgo.DirectCallAdapter{
		"Base":  func_Base,
		"Clean": func_Clean,
		"Dir":   func_Dir,
		"Ext":   func_Ext,
		"IsAbs": func_IsAbs,
		"Join":  func_Join,
	})
}

func func_Base(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Base(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Clean(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Clean(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Dir(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Dir(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Ext(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Ext(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_IsAbs(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsAbs(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Join(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Join(ixgo.DirectCallArg[[]string](ctx, 0)...))
}
