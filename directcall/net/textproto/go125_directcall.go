// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package textproto

import (
	q "net/textproto"

	bufio "bufio"
	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("net/textproto", map[string]ixgo.DirectCallAdapter{
		"(*Conn).Close":             method_ptr_Conn_Close,
		"(*Error).Error":            method_ptr_Error_Error,
		"(*MIMEHeader).Add":         method_ptr_MIMEHeader_Add,
		"(*MIMEHeader).Del":         method_ptr_MIMEHeader_Del,
		"(*MIMEHeader).Get":         method_ptr_MIMEHeader_Get,
		"(*MIMEHeader).Set":         method_ptr_MIMEHeader_Set,
		"(*MIMEHeader).Values":      method_ptr_MIMEHeader_Values,
		"(*Pipeline).EndRequest":    method_ptr_Pipeline_EndRequest,
		"(*Pipeline).EndResponse":   method_ptr_Pipeline_EndResponse,
		"(*Pipeline).Next":          method_ptr_Pipeline_Next,
		"(*Pipeline).StartRequest":  method_ptr_Pipeline_StartRequest,
		"(*Pipeline).StartResponse": method_ptr_Pipeline_StartResponse,
		"(*ProtocolError).Error":    method_ptr_ProtocolError_Error,
		"(*Reader).DotReader":       method_ptr_Reader_DotReader,
		"(*Writer).DotWriter":       method_ptr_Writer_DotWriter,
		"(*Writer).PrintfLine":      method_ptr_Writer_PrintfLine,
		"(MIMEHeader).Add":          method_MIMEHeader_Add,
		"(MIMEHeader).Del":          method_MIMEHeader_Del,
		"(MIMEHeader).Get":          method_MIMEHeader_Get,
		"(MIMEHeader).Set":          method_MIMEHeader_Set,
		"(MIMEHeader).Values":       method_MIMEHeader_Values,
		"(ProtocolError).Error":     method_ProtocolError_Error,
		"CanonicalMIMEHeaderKey":    func_CanonicalMIMEHeaderKey,
		"NewConn":                   func_NewConn,
		"NewReader":                 func_NewReader,
		"NewWriter":                 func_NewWriter,
		"TrimBytes":                 func_TrimBytes,
		"TrimString":                func_TrimString,
	})
}

func func_CanonicalMIMEHeaderKey(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CanonicalMIMEHeaderKey(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_Conn_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).Close(ixgo.DirectCallArg[*q.Conn](ctx, 0)))
}

func method_ptr_Error_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Error).Error(ixgo.DirectCallArg[*q.Error](ctx, 0)))
}

func method_MIMEHeader_Add(ctx ixgo.DirectCallContext) {
	q.MIMEHeader.Add(ixgo.DirectCallArg[q.MIMEHeader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ptr_MIMEHeader_Add(ctx ixgo.DirectCallContext) {
	(*q.MIMEHeader).Add(ixgo.DirectCallArg[*q.MIMEHeader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_MIMEHeader_Del(ctx ixgo.DirectCallContext) {
	q.MIMEHeader.Del(ixgo.DirectCallArg[q.MIMEHeader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_MIMEHeader_Del(ctx ixgo.DirectCallContext) {
	(*q.MIMEHeader).Del(ixgo.DirectCallArg[*q.MIMEHeader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_MIMEHeader_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MIMEHeader.Get(ixgo.DirectCallArg[q.MIMEHeader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_MIMEHeader_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MIMEHeader).Get(ixgo.DirectCallArg[*q.MIMEHeader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_MIMEHeader_Set(ctx ixgo.DirectCallContext) {
	q.MIMEHeader.Set(ixgo.DirectCallArg[q.MIMEHeader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ptr_MIMEHeader_Set(ctx ixgo.DirectCallContext) {
	(*q.MIMEHeader).Set(ixgo.DirectCallArg[*q.MIMEHeader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_MIMEHeader_Values(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MIMEHeader.Values(ixgo.DirectCallArg[q.MIMEHeader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_MIMEHeader_Values(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MIMEHeader).Values(ixgo.DirectCallArg[*q.MIMEHeader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_NewConn(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewConn(ixgo.DirectCallArg[io.ReadWriteCloser](ctx, 0)))
}

func func_NewReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReader(ixgo.DirectCallArg[*bufio.Reader](ctx, 0)))
}

func func_NewWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriter(ixgo.DirectCallArg[*bufio.Writer](ctx, 0)))
}

func method_ptr_Pipeline_EndRequest(ctx ixgo.DirectCallContext) {
	(*q.Pipeline).EndRequest(ixgo.DirectCallArg[*q.Pipeline](ctx, 0), ixgo.DirectCallArg[uint](ctx, 1))
}

func method_ptr_Pipeline_EndResponse(ctx ixgo.DirectCallContext) {
	(*q.Pipeline).EndResponse(ixgo.DirectCallArg[*q.Pipeline](ctx, 0), ixgo.DirectCallArg[uint](ctx, 1))
}

func method_ptr_Pipeline_Next(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pipeline).Next(ixgo.DirectCallArg[*q.Pipeline](ctx, 0)))
}

func method_ptr_Pipeline_StartRequest(ctx ixgo.DirectCallContext) {
	(*q.Pipeline).StartRequest(ixgo.DirectCallArg[*q.Pipeline](ctx, 0), ixgo.DirectCallArg[uint](ctx, 1))
}

func method_ptr_Pipeline_StartResponse(ctx ixgo.DirectCallContext) {
	(*q.Pipeline).StartResponse(ixgo.DirectCallArg[*q.Pipeline](ctx, 0), ixgo.DirectCallArg[uint](ctx, 1))
}

func method_ProtocolError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ProtocolError.Error(ixgo.DirectCallArg[q.ProtocolError](ctx, 0)))
}

func method_ptr_ProtocolError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProtocolError).Error(ixgo.DirectCallArg[*q.ProtocolError](ctx, 0)))
}

func method_ptr_Reader_DotReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).DotReader(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func func_TrimBytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrimBytes(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_TrimString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrimString(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_Writer_DotWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).DotWriter(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_PrintfLine(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).PrintfLine(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[[]any](ctx, 2)...))
}
