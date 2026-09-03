// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package html

import (
	q "html"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("html", map[string]ixgo.DirectCallAdapter{
		"EscapeString":   func_EscapeString,
		"UnescapeString": func_UnescapeString,
	})
}

func func_EscapeString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EscapeString(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_UnescapeString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnescapeString(ixgo.DirectCallArg[string](ctx, 0)))
}
