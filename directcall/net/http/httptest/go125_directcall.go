// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package httptest

import (
	q "net/http/httptest"

	context "context"
	"github.com/goplus/ixgo"
	io "io"
	http "net/http"
)

func init() {
	ixgo.RegisterDirectCalls("net/http/httptest", map[string]ixgo.DirectCallAdapter{
		"(*ResponseRecorder).Flush":        method_ptr_ResponseRecorder_Flush,
		"(*ResponseRecorder).Header":       method_ptr_ResponseRecorder_Header,
		"(*ResponseRecorder).Result":       method_ptr_ResponseRecorder_Result,
		"(*ResponseRecorder).WriteHeader":  method_ptr_ResponseRecorder_WriteHeader,
		"(*Server).Certificate":            method_ptr_Server_Certificate,
		"(*Server).Client":                 method_ptr_Server_Client,
		"(*Server).Close":                  method_ptr_Server_Close,
		"(*Server).CloseClientConnections": method_ptr_Server_CloseClientConnections,
		"(*Server).Start":                  method_ptr_Server_Start,
		"(*Server).StartTLS":               method_ptr_Server_StartTLS,
		"NewRecorder":                      func_NewRecorder,
		"NewRequest":                       func_NewRequest,
		"NewRequestWithContext":            func_NewRequestWithContext,
		"NewServer":                        func_NewServer,
		"NewTLSServer":                     func_NewTLSServer,
		"NewUnstartedServer":               func_NewUnstartedServer,
	})
}

func func_NewRecorder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewRecorder())
}

func func_NewRequest(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewRequest(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[io.Reader](ctx, 2)))
}

func func_NewRequestWithContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewRequestWithContext(ixgo.DirectCallArg[context.Context](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[io.Reader](ctx, 3)))
}

func func_NewServer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewServer(ixgo.DirectCallArg[http.Handler](ctx, 0)))
}

func func_NewTLSServer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewTLSServer(ixgo.DirectCallArg[http.Handler](ctx, 0)))
}

func func_NewUnstartedServer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewUnstartedServer(ixgo.DirectCallArg[http.Handler](ctx, 0)))
}

func method_ptr_ResponseRecorder_Flush(ctx ixgo.DirectCallContext) {
	(*q.ResponseRecorder).Flush(ixgo.DirectCallArg[*q.ResponseRecorder](ctx, 0))
}

func method_ptr_ResponseRecorder_Header(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ResponseRecorder).Header(ixgo.DirectCallArg[*q.ResponseRecorder](ctx, 0)))
}

func method_ptr_ResponseRecorder_Result(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ResponseRecorder).Result(ixgo.DirectCallArg[*q.ResponseRecorder](ctx, 0)))
}

func method_ptr_ResponseRecorder_WriteHeader(ctx ixgo.DirectCallContext) {
	(*q.ResponseRecorder).WriteHeader(ixgo.DirectCallArg[*q.ResponseRecorder](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_Server_Certificate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).Certificate(ixgo.DirectCallArg[*q.Server](ctx, 0)))
}

func method_ptr_Server_Client(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).Client(ixgo.DirectCallArg[*q.Server](ctx, 0)))
}

func method_ptr_Server_Close(ctx ixgo.DirectCallContext) {
	(*q.Server).Close(ixgo.DirectCallArg[*q.Server](ctx, 0))
}

func method_ptr_Server_CloseClientConnections(ctx ixgo.DirectCallContext) {
	(*q.Server).CloseClientConnections(ixgo.DirectCallArg[*q.Server](ctx, 0))
}

func method_ptr_Server_Start(ctx ixgo.DirectCallContext) {
	(*q.Server).Start(ixgo.DirectCallArg[*q.Server](ctx, 0))
}

func method_ptr_Server_StartTLS(ctx ixgo.DirectCallContext) {
	(*q.Server).StartTLS(ixgo.DirectCallArg[*q.Server](ctx, 0))
}
