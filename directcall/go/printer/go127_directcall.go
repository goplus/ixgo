// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package printer

import (
	q "go/printer"

	"github.com/goplus/ixgo"
	token "go/token"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("go/printer", map[string]ixgo.DirectCallAdapter{
		"(*Config).Fprint": method_ptr_Config_Fprint,
		"Fprint":           func_Fprint,
	})
}

func method_ptr_Config_Fprint(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Config).Fprint(ixgo.DirectCallArg[*q.Config](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[*token.FileSet](ctx, 2), ixgo.DirectCallArg[interface{}](ctx, 3)))
}

func func_Fprint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Fprint(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[*token.FileSet](ctx, 1), ixgo.DirectCallArg[interface{}](ctx, 2)))
}
