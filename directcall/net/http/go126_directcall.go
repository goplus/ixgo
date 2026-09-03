// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package http

import (
	q "net/http"

	context "context"
	"github.com/goplus/ixgo"
	io "io"
	fs "io/fs"
	net "net"
	url "net/url"
	time "time"
)

func init() {
	ixgo.RegisterDirectCalls("net/http", map[string]ixgo.DirectCallAdapter{
		"(*Client).CloseIdleConnections":                    method_ptr_Client_CloseIdleConnections,
		"(*ClientConn).Available":                           method_ptr_ClientConn_Available,
		"(*ClientConn).Close":                               method_ptr_ClientConn_Close,
		"(*ClientConn).Err":                                 method_ptr_ClientConn_Err,
		"(*ClientConn).InFlight":                            method_ptr_ClientConn_InFlight,
		"(*ClientConn).Release":                             method_ptr_ClientConn_Release,
		"(*ClientConn).Reserve":                             method_ptr_ClientConn_Reserve,
		"(*ClientConn).SetStateHook":                        method_ptr_ClientConn_SetStateHook,
		"(*ConnState).String":                               method_ptr_ConnState_String,
		"(*Cookie).String":                                  method_ptr_Cookie_String,
		"(*Cookie).Valid":                                   method_ptr_Cookie_Valid,
		"(*CrossOriginProtection).AddInsecureBypassPattern": method_ptr_CrossOriginProtection_AddInsecureBypassPattern,
		"(*CrossOriginProtection).AddTrustedOrigin":         method_ptr_CrossOriginProtection_AddTrustedOrigin,
		"(*CrossOriginProtection).Check":                    method_ptr_CrossOriginProtection_Check,
		"(*CrossOriginProtection).Handler":                  method_ptr_CrossOriginProtection_Handler,
		"(*CrossOriginProtection).SetDenyHandler":           method_ptr_CrossOriginProtection_SetDenyHandler,
		"(*HandlerFunc).ServeHTTP":                          method_ptr_HandlerFunc_ServeHTTP,
		"(*Header).Add":                                     method_ptr_Header_Add,
		"(*Header).Clone":                                   method_ptr_Header_Clone,
		"(*Header).Del":                                     method_ptr_Header_Del,
		"(*Header).Get":                                     method_ptr_Header_Get,
		"(*Header).Set":                                     method_ptr_Header_Set,
		"(*Header).Values":                                  method_ptr_Header_Values,
		"(*Header).Write":                                   method_ptr_Header_Write,
		"(*Header).WriteSubset":                             method_ptr_Header_WriteSubset,
		"(*MaxBytesError).Error":                            method_ptr_MaxBytesError_Error,
		"(*ProtocolError).Error":                            method_ptr_ProtocolError_Error,
		"(*ProtocolError).Is":                               method_ptr_ProtocolError_Is,
		"(*Protocols).HTTP1":                                method_ptr_Protocols_HTTP1,
		"(*Protocols).HTTP2":                                method_ptr_Protocols_HTTP2,
		"(*Protocols).SetHTTP1":                             method_ptr_Protocols_SetHTTP1,
		"(*Protocols).SetHTTP2":                             method_ptr_Protocols_SetHTTP2,
		"(*Protocols).SetUnencryptedHTTP2":                  method_ptr_Protocols_SetUnencryptedHTTP2,
		"(*Protocols).String":                               method_ptr_Protocols_String,
		"(*Protocols).UnencryptedHTTP2":                     method_ptr_Protocols_UnencryptedHTTP2,
		"(*Request).AddCookie":                              method_ptr_Request_AddCookie,
		"(*Request).Clone":                                  method_ptr_Request_Clone,
		"(*Request).Context":                                method_ptr_Request_Context,
		"(*Request).Cookies":                                method_ptr_Request_Cookies,
		"(*Request).CookiesNamed":                           method_ptr_Request_CookiesNamed,
		"(*Request).FormValue":                              method_ptr_Request_FormValue,
		"(*Request).ParseForm":                              method_ptr_Request_ParseForm,
		"(*Request).ParseMultipartForm":                     method_ptr_Request_ParseMultipartForm,
		"(*Request).PathValue":                              method_ptr_Request_PathValue,
		"(*Request).PostFormValue":                          method_ptr_Request_PostFormValue,
		"(*Request).ProtoAtLeast":                           method_ptr_Request_ProtoAtLeast,
		"(*Request).Referer":                                method_ptr_Request_Referer,
		"(*Request).SetBasicAuth":                           method_ptr_Request_SetBasicAuth,
		"(*Request).SetPathValue":                           method_ptr_Request_SetPathValue,
		"(*Request).UserAgent":                              method_ptr_Request_UserAgent,
		"(*Request).WithContext":                            method_ptr_Request_WithContext,
		"(*Request).Write":                                  method_ptr_Request_Write,
		"(*Request).WriteProxy":                             method_ptr_Request_WriteProxy,
		"(*Response).Cookies":                               method_ptr_Response_Cookies,
		"(*Response).ProtoAtLeast":                          method_ptr_Response_ProtoAtLeast,
		"(*Response).Write":                                 method_ptr_Response_Write,
		"(*ResponseController).EnableFullDuplex":            method_ptr_ResponseController_EnableFullDuplex,
		"(*ResponseController).Flush":                       method_ptr_ResponseController_Flush,
		"(*ResponseController).SetReadDeadline":             method_ptr_ResponseController_SetReadDeadline,
		"(*ResponseController).SetWriteDeadline":            method_ptr_ResponseController_SetWriteDeadline,
		"(*ServeMux).Handle":                                method_ptr_ServeMux_Handle,
		"(*ServeMux).HandleFunc":                            method_ptr_ServeMux_HandleFunc,
		"(*ServeMux).ServeHTTP":                             method_ptr_ServeMux_ServeHTTP,
		"(*Server).Close":                                   method_ptr_Server_Close,
		"(*Server).ListenAndServe":                          method_ptr_Server_ListenAndServe,
		"(*Server).ListenAndServeTLS":                       method_ptr_Server_ListenAndServeTLS,
		"(*Server).RegisterOnShutdown":                      method_ptr_Server_RegisterOnShutdown,
		"(*Server).Serve":                                   method_ptr_Server_Serve,
		"(*Server).ServeTLS":                                method_ptr_Server_ServeTLS,
		"(*Server).SetKeepAlivesEnabled":                    method_ptr_Server_SetKeepAlivesEnabled,
		"(*Server).Shutdown":                                method_ptr_Server_Shutdown,
		"(*Transport).CancelRequest":                        method_ptr_Transport_CancelRequest,
		"(*Transport).Clone":                                method_ptr_Transport_Clone,
		"(*Transport).CloseIdleConnections":                 method_ptr_Transport_CloseIdleConnections,
		"(*Transport).RegisterProtocol":                     method_ptr_Transport_RegisterProtocol,
		"(ConnState).String":                                method_ConnState_String,
		"(HandlerFunc).ServeHTTP":                           method_HandlerFunc_ServeHTTP,
		"(Header).Add":                                      method_Header_Add,
		"(Header).Clone":                                    method_Header_Clone,
		"(Header).Del":                                      method_Header_Del,
		"(Header).Get":                                      method_Header_Get,
		"(Header).Set":                                      method_Header_Set,
		"(Header).Values":                                   method_Header_Values,
		"(Header).Write":                                    method_Header_Write,
		"(Header).WriteSubset":                              method_Header_WriteSubset,
		"(Protocols).HTTP1":                                 method_Protocols_HTTP1,
		"(Protocols).HTTP2":                                 method_Protocols_HTTP2,
		"(Protocols).String":                                method_Protocols_String,
		"(Protocols).UnencryptedHTTP2":                      method_Protocols_UnencryptedHTTP2,
		"AllowQuerySemicolons":                              func_AllowQuerySemicolons,
		"CanonicalHeaderKey":                                func_CanonicalHeaderKey,
		"DetectContentType":                                 func_DetectContentType,
		"Error":                                             func_Error,
		"FS":                                                func_FS,
		"FileServer":                                        func_FileServer,
		"FileServerFS":                                      func_FileServerFS,
		"Handle":                                            func_Handle,
		"HandleFunc":                                        func_HandleFunc,
		"ListenAndServe":                                    func_ListenAndServe,
		"ListenAndServeTLS":                                 func_ListenAndServeTLS,
		"MaxBytesHandler":                                   func_MaxBytesHandler,
		"MaxBytesReader":                                    func_MaxBytesReader,
		"NewCrossOriginProtection":                          func_NewCrossOriginProtection,
		"NewFileTransport":                                  func_NewFileTransport,
		"NewFileTransportFS":                                func_NewFileTransportFS,
		"NewResponseController":                             func_NewResponseController,
		"NewServeMux":                                       func_NewServeMux,
		"NotFound":                                          func_NotFound,
		"NotFoundHandler":                                   func_NotFoundHandler,
		"ProxyURL":                                          func_ProxyURL,
		"Redirect":                                          func_Redirect,
		"RedirectHandler":                                   func_RedirectHandler,
		"Serve":                                             func_Serve,
		"ServeContent":                                      func_ServeContent,
		"ServeFile":                                         func_ServeFile,
		"ServeFileFS":                                       func_ServeFileFS,
		"ServeTLS":                                          func_ServeTLS,
		"SetCookie":                                         func_SetCookie,
		"StatusText":                                        func_StatusText,
		"StripPrefix":                                       func_StripPrefix,
		"TimeoutHandler":                                    func_TimeoutHandler,
	})
}

func func_AllowQuerySemicolons(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AllowQuerySemicolons(ixgo.DirectCallArg[q.Handler](ctx, 0)))
}

func func_CanonicalHeaderKey(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CanonicalHeaderKey(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_Client_CloseIdleConnections(ctx ixgo.DirectCallContext) {
	(*q.Client).CloseIdleConnections(ixgo.DirectCallArg[*q.Client](ctx, 0))
}

func method_ptr_ClientConn_Available(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientConn).Available(ixgo.DirectCallArg[*q.ClientConn](ctx, 0)))
}

func method_ptr_ClientConn_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientConn).Close(ixgo.DirectCallArg[*q.ClientConn](ctx, 0)))
}

func method_ptr_ClientConn_Err(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientConn).Err(ixgo.DirectCallArg[*q.ClientConn](ctx, 0)))
}

func method_ptr_ClientConn_InFlight(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientConn).InFlight(ixgo.DirectCallArg[*q.ClientConn](ctx, 0)))
}

func method_ptr_ClientConn_Release(ctx ixgo.DirectCallContext) {
	(*q.ClientConn).Release(ixgo.DirectCallArg[*q.ClientConn](ctx, 0))
}

func method_ptr_ClientConn_Reserve(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientConn).Reserve(ixgo.DirectCallArg[*q.ClientConn](ctx, 0)))
}

func method_ptr_ClientConn_SetStateHook(ctx ixgo.DirectCallContext) {
	(*q.ClientConn).SetStateHook(ixgo.DirectCallArg[*q.ClientConn](ctx, 0), ixgo.DirectCallArg[func(*q.ClientConn)](ctx, 1))
}

func method_ConnState_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ConnState.String(ixgo.DirectCallArg[q.ConnState](ctx, 0)))
}

func method_ptr_ConnState_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ConnState).String(ixgo.DirectCallArg[*q.ConnState](ctx, 0)))
}

func method_ptr_Cookie_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Cookie).String(ixgo.DirectCallArg[*q.Cookie](ctx, 0)))
}

func method_ptr_Cookie_Valid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Cookie).Valid(ixgo.DirectCallArg[*q.Cookie](ctx, 0)))
}

func method_ptr_CrossOriginProtection_AddInsecureBypassPattern(ctx ixgo.DirectCallContext) {
	(*q.CrossOriginProtection).AddInsecureBypassPattern(ixgo.DirectCallArg[*q.CrossOriginProtection](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_CrossOriginProtection_AddTrustedOrigin(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CrossOriginProtection).AddTrustedOrigin(ixgo.DirectCallArg[*q.CrossOriginProtection](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_CrossOriginProtection_Check(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CrossOriginProtection).Check(ixgo.DirectCallArg[*q.CrossOriginProtection](ctx, 0), ixgo.DirectCallArg[*q.Request](ctx, 1)))
}

func method_ptr_CrossOriginProtection_Handler(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CrossOriginProtection).Handler(ixgo.DirectCallArg[*q.CrossOriginProtection](ctx, 0), ixgo.DirectCallArg[q.Handler](ctx, 1)))
}

func method_ptr_CrossOriginProtection_SetDenyHandler(ctx ixgo.DirectCallContext) {
	(*q.CrossOriginProtection).SetDenyHandler(ixgo.DirectCallArg[*q.CrossOriginProtection](ctx, 0), ixgo.DirectCallArg[q.Handler](ctx, 1))
}

func func_DetectContentType(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DetectContentType(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_Error(ctx ixgo.DirectCallContext) {
	q.Error(ixgo.DirectCallArg[q.ResponseWriter](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func func_FS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FS(ixgo.DirectCallArg[fs.FS](ctx, 0)))
}

func func_FileServer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FileServer(ixgo.DirectCallArg[q.FileSystem](ctx, 0)))
}

func func_FileServerFS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FileServerFS(ixgo.DirectCallArg[fs.FS](ctx, 0)))
}

func func_Handle(ctx ixgo.DirectCallContext) {
	q.Handle(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[q.Handler](ctx, 1))
}

func func_HandleFunc(ctx ixgo.DirectCallContext) {
	q.HandleFunc(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[func(q.ResponseWriter, *q.Request)](ctx, 1))
}

func method_HandlerFunc_ServeHTTP(ctx ixgo.DirectCallContext) {
	q.HandlerFunc.ServeHTTP(ixgo.DirectCallArg[q.HandlerFunc](ctx, 0), ixgo.DirectCallArg[q.ResponseWriter](ctx, 1), ixgo.DirectCallArg[*q.Request](ctx, 2))
}

func method_ptr_HandlerFunc_ServeHTTP(ctx ixgo.DirectCallContext) {
	(*q.HandlerFunc).ServeHTTP(ixgo.DirectCallArg[*q.HandlerFunc](ctx, 0), ixgo.DirectCallArg[q.ResponseWriter](ctx, 1), ixgo.DirectCallArg[*q.Request](ctx, 2))
}

func method_Header_Add(ctx ixgo.DirectCallContext) {
	q.Header.Add(ixgo.DirectCallArg[q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ptr_Header_Add(ctx ixgo.DirectCallContext) {
	(*q.Header).Add(ixgo.DirectCallArg[*q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_Header_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Header.Clone(ixgo.DirectCallArg[q.Header](ctx, 0)))
}

func method_ptr_Header_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Header).Clone(ixgo.DirectCallArg[*q.Header](ctx, 0)))
}

func method_Header_Del(ctx ixgo.DirectCallContext) {
	q.Header.Del(ixgo.DirectCallArg[q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_Header_Del(ctx ixgo.DirectCallContext) {
	(*q.Header).Del(ixgo.DirectCallArg[*q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_Header_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Header.Get(ixgo.DirectCallArg[q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Header_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Header).Get(ixgo.DirectCallArg[*q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_Header_Set(ctx ixgo.DirectCallContext) {
	q.Header.Set(ixgo.DirectCallArg[q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ptr_Header_Set(ctx ixgo.DirectCallContext) {
	(*q.Header).Set(ixgo.DirectCallArg[*q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_Header_Values(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Header.Values(ixgo.DirectCallArg[q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Header_Values(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Header).Values(ixgo.DirectCallArg[*q.Header](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_Header_Write(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Header.Write(ixgo.DirectCallArg[q.Header](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1)))
}

func method_ptr_Header_Write(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Header).Write(ixgo.DirectCallArg[*q.Header](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1)))
}

func method_Header_WriteSubset(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Header.WriteSubset(ixgo.DirectCallArg[q.Header](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[map[string]bool](ctx, 2)))
}

func method_ptr_Header_WriteSubset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Header).WriteSubset(ixgo.DirectCallArg[*q.Header](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[map[string]bool](ctx, 2)))
}

func func_ListenAndServe(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ListenAndServe(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[q.Handler](ctx, 1)))
}

func func_ListenAndServeTLS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ListenAndServeTLS(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[q.Handler](ctx, 3)))
}

func method_ptr_MaxBytesError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MaxBytesError).Error(ixgo.DirectCallArg[*q.MaxBytesError](ctx, 0)))
}

func func_MaxBytesHandler(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MaxBytesHandler(ixgo.DirectCallArg[q.Handler](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_MaxBytesReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MaxBytesReader(ixgo.DirectCallArg[q.ResponseWriter](ctx, 0), ixgo.DirectCallArg[io.ReadCloser](ctx, 1), ixgo.DirectCallArg[int64](ctx, 2)))
}

func func_NewCrossOriginProtection(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCrossOriginProtection())
}

func func_NewFileTransport(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewFileTransport(ixgo.DirectCallArg[q.FileSystem](ctx, 0)))
}

func func_NewFileTransportFS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewFileTransportFS(ixgo.DirectCallArg[fs.FS](ctx, 0)))
}

func func_NewResponseController(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewResponseController(ixgo.DirectCallArg[q.ResponseWriter](ctx, 0)))
}

func func_NewServeMux(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewServeMux())
}

func func_NotFound(ctx ixgo.DirectCallContext) {
	q.NotFound(ixgo.DirectCallArg[q.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*q.Request](ctx, 1))
}

func func_NotFoundHandler(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NotFoundHandler())
}

func method_ptr_ProtocolError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProtocolError).Error(ixgo.DirectCallArg[*q.ProtocolError](ctx, 0)))
}

func method_ptr_ProtocolError_Is(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProtocolError).Is(ixgo.DirectCallArg[*q.ProtocolError](ctx, 0), ixgo.DirectCallArg[error](ctx, 1)))
}

func method_Protocols_HTTP1(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Protocols.HTTP1(ixgo.DirectCallArg[q.Protocols](ctx, 0)))
}

func method_ptr_Protocols_HTTP1(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Protocols).HTTP1(ixgo.DirectCallArg[*q.Protocols](ctx, 0)))
}

func method_Protocols_HTTP2(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Protocols.HTTP2(ixgo.DirectCallArg[q.Protocols](ctx, 0)))
}

func method_ptr_Protocols_HTTP2(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Protocols).HTTP2(ixgo.DirectCallArg[*q.Protocols](ctx, 0)))
}

func method_ptr_Protocols_SetHTTP1(ctx ixgo.DirectCallContext) {
	(*q.Protocols).SetHTTP1(ixgo.DirectCallArg[*q.Protocols](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_ptr_Protocols_SetHTTP2(ctx ixgo.DirectCallContext) {
	(*q.Protocols).SetHTTP2(ixgo.DirectCallArg[*q.Protocols](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_ptr_Protocols_SetUnencryptedHTTP2(ctx ixgo.DirectCallContext) {
	(*q.Protocols).SetUnencryptedHTTP2(ixgo.DirectCallArg[*q.Protocols](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_Protocols_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Protocols.String(ixgo.DirectCallArg[q.Protocols](ctx, 0)))
}

func method_ptr_Protocols_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Protocols).String(ixgo.DirectCallArg[*q.Protocols](ctx, 0)))
}

func method_Protocols_UnencryptedHTTP2(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Protocols.UnencryptedHTTP2(ixgo.DirectCallArg[q.Protocols](ctx, 0)))
}

func method_ptr_Protocols_UnencryptedHTTP2(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Protocols).UnencryptedHTTP2(ixgo.DirectCallArg[*q.Protocols](ctx, 0)))
}

func func_ProxyURL(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ProxyURL(ixgo.DirectCallArg[*url.URL](ctx, 0)))
}

func func_Redirect(ctx ixgo.DirectCallContext) {
	q.Redirect(ixgo.DirectCallArg[q.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*q.Request](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[int](ctx, 3))
}

func func_RedirectHandler(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RedirectHandler(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Request_AddCookie(ctx ixgo.DirectCallContext) {
	(*q.Request).AddCookie(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[*q.Cookie](ctx, 1))
}

func method_ptr_Request_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).Clone(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1)))
}

func method_ptr_Request_Context(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).Context(ixgo.DirectCallArg[*q.Request](ctx, 0)))
}

func method_ptr_Request_Cookies(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).Cookies(ixgo.DirectCallArg[*q.Request](ctx, 0)))
}

func method_ptr_Request_CookiesNamed(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).CookiesNamed(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Request_FormValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).FormValue(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Request_ParseForm(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).ParseForm(ixgo.DirectCallArg[*q.Request](ctx, 0)))
}

func method_ptr_Request_ParseMultipartForm(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).ParseMultipartForm(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Request_PathValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).PathValue(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Request_PostFormValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).PostFormValue(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Request_ProtoAtLeast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).ProtoAtLeast(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Request_Referer(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).Referer(ixgo.DirectCallArg[*q.Request](ctx, 0)))
}

func method_ptr_Request_SetBasicAuth(ctx ixgo.DirectCallContext) {
	(*q.Request).SetBasicAuth(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ptr_Request_SetPathValue(ctx ixgo.DirectCallContext) {
	(*q.Request).SetPathValue(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ptr_Request_UserAgent(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).UserAgent(ixgo.DirectCallArg[*q.Request](ctx, 0)))
}

func method_ptr_Request_WithContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).WithContext(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1)))
}

func method_ptr_Request_Write(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).Write(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1)))
}

func method_ptr_Request_WriteProxy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Request).WriteProxy(ixgo.DirectCallArg[*q.Request](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1)))
}

func method_ptr_Response_Cookies(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Response).Cookies(ixgo.DirectCallArg[*q.Response](ctx, 0)))
}

func method_ptr_Response_ProtoAtLeast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Response).ProtoAtLeast(ixgo.DirectCallArg[*q.Response](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Response_Write(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Response).Write(ixgo.DirectCallArg[*q.Response](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1)))
}

func method_ptr_ResponseController_EnableFullDuplex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ResponseController).EnableFullDuplex(ixgo.DirectCallArg[*q.ResponseController](ctx, 0)))
}

func method_ptr_ResponseController_Flush(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ResponseController).Flush(ixgo.DirectCallArg[*q.ResponseController](ctx, 0)))
}

func method_ptr_ResponseController_SetReadDeadline(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ResponseController).SetReadDeadline(ixgo.DirectCallArg[*q.ResponseController](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func method_ptr_ResponseController_SetWriteDeadline(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ResponseController).SetWriteDeadline(ixgo.DirectCallArg[*q.ResponseController](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func func_Serve(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Serve(ixgo.DirectCallArg[net.Listener](ctx, 0), ixgo.DirectCallArg[q.Handler](ctx, 1)))
}

func func_ServeContent(ctx ixgo.DirectCallContext) {
	q.ServeContent(ixgo.DirectCallArg[q.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*q.Request](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[time.Time](ctx, 3), ixgo.DirectCallArg[io.ReadSeeker](ctx, 4))
}

func func_ServeFile(ctx ixgo.DirectCallContext) {
	q.ServeFile(ixgo.DirectCallArg[q.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*q.Request](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func func_ServeFileFS(ctx ixgo.DirectCallContext) {
	q.ServeFileFS(ixgo.DirectCallArg[q.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*q.Request](ctx, 1), ixgo.DirectCallArg[fs.FS](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func method_ptr_ServeMux_Handle(ctx ixgo.DirectCallContext) {
	(*q.ServeMux).Handle(ixgo.DirectCallArg[*q.ServeMux](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[q.Handler](ctx, 2))
}

func method_ptr_ServeMux_HandleFunc(ctx ixgo.DirectCallContext) {
	(*q.ServeMux).HandleFunc(ixgo.DirectCallArg[*q.ServeMux](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[func(q.ResponseWriter, *q.Request)](ctx, 2))
}

func method_ptr_ServeMux_ServeHTTP(ctx ixgo.DirectCallContext) {
	(*q.ServeMux).ServeHTTP(ixgo.DirectCallArg[*q.ServeMux](ctx, 0), ixgo.DirectCallArg[q.ResponseWriter](ctx, 1), ixgo.DirectCallArg[*q.Request](ctx, 2))
}

func func_ServeTLS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ServeTLS(ixgo.DirectCallArg[net.Listener](ctx, 0), ixgo.DirectCallArg[q.Handler](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_Server_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).Close(ixgo.DirectCallArg[*q.Server](ctx, 0)))
}

func method_ptr_Server_ListenAndServe(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).ListenAndServe(ixgo.DirectCallArg[*q.Server](ctx, 0)))
}

func method_ptr_Server_ListenAndServeTLS(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).ListenAndServeTLS(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_Server_RegisterOnShutdown(ctx ixgo.DirectCallContext) {
	(*q.Server).RegisterOnShutdown(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[func()](ctx, 1))
}

func method_ptr_Server_Serve(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).Serve(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[net.Listener](ctx, 1)))
}

func method_ptr_Server_ServeTLS(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).ServeTLS(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[net.Listener](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_Server_SetKeepAlivesEnabled(ctx ixgo.DirectCallContext) {
	(*q.Server).SetKeepAlivesEnabled(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_ptr_Server_Shutdown(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Server).Shutdown(ixgo.DirectCallArg[*q.Server](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1)))
}

func func_SetCookie(ctx ixgo.DirectCallContext) {
	q.SetCookie(ixgo.DirectCallArg[q.ResponseWriter](ctx, 0), ixgo.DirectCallArg[*q.Cookie](ctx, 1))
}

func func_StatusText(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StatusText(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_StripPrefix(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StripPrefix(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[q.Handler](ctx, 1)))
}

func func_TimeoutHandler(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TimeoutHandler(ixgo.DirectCallArg[q.Handler](ctx, 0), ixgo.DirectCallArg[time.Duration](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_Transport_CancelRequest(ctx ixgo.DirectCallContext) {
	(*q.Transport).CancelRequest(ixgo.DirectCallArg[*q.Transport](ctx, 0), ixgo.DirectCallArg[*q.Request](ctx, 1))
}

func method_ptr_Transport_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Transport).Clone(ixgo.DirectCallArg[*q.Transport](ctx, 0)))
}

func method_ptr_Transport_CloseIdleConnections(ctx ixgo.DirectCallContext) {
	(*q.Transport).CloseIdleConnections(ixgo.DirectCallArg[*q.Transport](ctx, 0))
}

func method_ptr_Transport_RegisterProtocol(ctx ixgo.DirectCallContext) {
	(*q.Transport).RegisterProtocol(ixgo.DirectCallArg[*q.Transport](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[q.RoundTripper](ctx, 2))
}
