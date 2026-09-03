// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package smtp

import (
	q "net/smtp"

	tls "crypto/tls"
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("net/smtp", map[string]ixgo.DirectCallAdapter{
		"(*Client).Auth":     method_ptr_Client_Auth,
		"(*Client).Close":    method_ptr_Client_Close,
		"(*Client).Hello":    method_ptr_Client_Hello,
		"(*Client).Mail":     method_ptr_Client_Mail,
		"(*Client).Noop":     method_ptr_Client_Noop,
		"(*Client).Quit":     method_ptr_Client_Quit,
		"(*Client).Rcpt":     method_ptr_Client_Rcpt,
		"(*Client).Reset":    method_ptr_Client_Reset,
		"(*Client).StartTLS": method_ptr_Client_StartTLS,
		"(*Client).Verify":   method_ptr_Client_Verify,
		"CRAMMD5Auth":        func_CRAMMD5Auth,
		"PlainAuth":          func_PlainAuth,
		"SendMail":           func_SendMail,
	})
}

func func_CRAMMD5Auth(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CRAMMD5Auth(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Client_Auth(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Auth(ixgo.DirectCallArg[*q.Client](ctx, 0), ixgo.DirectCallArg[q.Auth](ctx, 1)))
}

func method_ptr_Client_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Close(ixgo.DirectCallArg[*q.Client](ctx, 0)))
}

func method_ptr_Client_Hello(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Hello(ixgo.DirectCallArg[*q.Client](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Client_Mail(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Mail(ixgo.DirectCallArg[*q.Client](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Client_Noop(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Noop(ixgo.DirectCallArg[*q.Client](ctx, 0)))
}

func method_ptr_Client_Quit(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Quit(ixgo.DirectCallArg[*q.Client](ctx, 0)))
}

func method_ptr_Client_Rcpt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Rcpt(ixgo.DirectCallArg[*q.Client](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Client_Reset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Reset(ixgo.DirectCallArg[*q.Client](ctx, 0)))
}

func method_ptr_Client_StartTLS(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).StartTLS(ixgo.DirectCallArg[*q.Client](ctx, 0), ixgo.DirectCallArg[*tls.Config](ctx, 1)))
}

func method_ptr_Client_Verify(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Client).Verify(ixgo.DirectCallArg[*q.Client](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_PlainAuth(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PlainAuth(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func func_SendMail(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SendMail(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[q.Auth](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[[]string](ctx, 3), ixgo.DirectCallArg[[]byte](ctx, 4)))
}
