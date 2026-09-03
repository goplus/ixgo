// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package fmt

import (
	q "fmt"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("fmt", map[string]ixgo.DirectCallAdapter{
		"Append":       func_Append,
		"Appendf":      func_Appendf,
		"Appendln":     func_Appendln,
		"Errorf":       func_Errorf,
		"FormatString": func_FormatString,
		"Sprint":       func_Sprint,
		"Sprintf":      func_Sprintf,
		"Sprintln":     func_Sprintln,
	})
}

func func_Append(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Append(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[[]any](ctx, 1)...))
}

func func_Appendf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Appendf(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[[]any](ctx, 2)...))
}

func func_Appendln(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Appendln(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[[]any](ctx, 1)...))
}

func func_Errorf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Errorf(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[[]any](ctx, 1)...))
}

func func_FormatString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatString(ixgo.DirectCallArg[q.State](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_Sprint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sprint(ixgo.DirectCallArg[[]any](ctx, 0)...))
}

func func_Sprintf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sprintf(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[[]any](ctx, 1)...))
}

func func_Sprintln(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sprintln(ixgo.DirectCallArg[[]any](ctx, 0)...))
}
