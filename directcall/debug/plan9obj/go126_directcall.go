// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package plan9obj

import (
	q "debug/plan9obj"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("debug/plan9obj", map[string]ixgo.DirectCallAdapter{
		"(*File).Close":   method_ptr_File_Close,
		"(*File).Section": method_ptr_File_Section,
		"(*Section).Open": method_ptr_Section_Open,
	})
}

func method_ptr_File_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Close(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Section(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Section(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Section_Open(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Section).Open(ixgo.DirectCallArg[*q.Section](ctx, 0)))
}
