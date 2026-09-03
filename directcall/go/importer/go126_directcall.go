// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package importer

import (
	q "go/importer"

	"github.com/goplus/ixgo"
	token "go/token"
)

func init() {
	ixgo.RegisterDirectCalls("go/importer", map[string]ixgo.DirectCallAdapter{
		"Default":     func_Default,
		"For":         func_For,
		"ForCompiler": func_ForCompiler,
	})
}

func func_Default(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Default())
}

func func_For(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.For(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[q.Lookup](ctx, 1)))
}

func func_ForCompiler(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ForCompiler(ixgo.DirectCallArg[*token.FileSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[q.Lookup](ctx, 2)))
}
