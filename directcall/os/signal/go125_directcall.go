// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package signal

import (
	q "os/signal"

	"github.com/goplus/ixgo"
	os "os"
)

func init() {
	ixgo.RegisterDirectCalls("os/signal", map[string]ixgo.DirectCallAdapter{
		"Ignore":  func_Ignore,
		"Ignored": func_Ignored,
		"Notify":  func_Notify,
		"Reset":   func_Reset,
		"Stop":    func_Stop,
	})
}

func func_Ignore(ctx ixgo.DirectCallContext) {
	q.Ignore(ixgo.DirectCallArg[[]os.Signal](ctx, 0)...)
}

func func_Ignored(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Ignored(ixgo.DirectCallArg[os.Signal](ctx, 0)))
}

func func_Notify(ctx ixgo.DirectCallContext) {
	q.Notify(ixgo.DirectCallArg[chan<- os.Signal](ctx, 0), ixgo.DirectCallArg[[]os.Signal](ctx, 1)...)
}

func func_Reset(ctx ixgo.DirectCallContext) {
	q.Reset(ixgo.DirectCallArg[[]os.Signal](ctx, 0)...)
}

func func_Stop(ctx ixgo.DirectCallContext) {
	q.Stop(ixgo.DirectCallArg[chan<- os.Signal](ctx, 0))
}
