// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package coverage

import (
	q "runtime/coverage"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("runtime/coverage", map[string]ixgo.DirectCallAdapter{
		"ClearCounters":    func_ClearCounters,
		"WriteCounters":    func_WriteCounters,
		"WriteCountersDir": func_WriteCountersDir,
		"WriteMeta":        func_WriteMeta,
		"WriteMetaDir":     func_WriteMetaDir,
	})
}

func func_ClearCounters(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ClearCounters())
}

func func_WriteCounters(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WriteCounters(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func func_WriteCountersDir(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WriteCountersDir(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_WriteMeta(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WriteMeta(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func func_WriteMetaDir(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WriteMetaDir(ixgo.DirectCallArg[string](ctx, 0)))
}
