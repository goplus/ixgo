// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package scanner

import (
	q "text/scanner"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("text/scanner", map[string]ixgo.DirectCallAdapter{
		"(*Position).IsValid":  method_ptr_Position_IsValid,
		"(*Position).String":   method_ptr_Position_String,
		"(*Scanner).Init":      method_ptr_Scanner_Init,
		"(*Scanner).Next":      method_ptr_Scanner_Next,
		"(*Scanner).Peek":      method_ptr_Scanner_Peek,
		"(*Scanner).Pos":       method_ptr_Scanner_Pos,
		"(*Scanner).Scan":      method_ptr_Scanner_Scan,
		"(*Scanner).TokenText": method_ptr_Scanner_TokenText,
		"(Position).String":    method_Position_String,
		"TokenString":          func_TokenString,
	})
}

func method_ptr_Position_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Position).IsValid(ixgo.DirectCallArg[*q.Position](ctx, 0)))
}

func method_Position_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Position.String(ixgo.DirectCallArg[q.Position](ctx, 0)))
}

func method_ptr_Position_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Position).String(ixgo.DirectCallArg[*q.Position](ctx, 0)))
}

func method_ptr_Scanner_Init(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).Init(ixgo.DirectCallArg[*q.Scanner](ctx, 0), ixgo.DirectCallArg[io.Reader](ctx, 1)))
}

func method_ptr_Scanner_Next(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).Next(ixgo.DirectCallArg[*q.Scanner](ctx, 0)))
}

func method_ptr_Scanner_Peek(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).Peek(ixgo.DirectCallArg[*q.Scanner](ctx, 0)))
}

func method_ptr_Scanner_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).Pos(ixgo.DirectCallArg[*q.Scanner](ctx, 0)))
}

func method_ptr_Scanner_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).Scan(ixgo.DirectCallArg[*q.Scanner](ctx, 0)))
}

func method_ptr_Scanner_TokenText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).TokenText(ixgo.DirectCallArg[*q.Scanner](ctx, 0)))
}

func func_TokenString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TokenString(ixgo.DirectCallArg[rune](ctx, 0)))
}
