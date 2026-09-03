// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package tls

import (
	q "crypto/tls"

	context "context"
	"github.com/goplus/ixgo"
	net "net"
	time "time"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/tls", map[string]ixgo.DirectCallAdapter{
		"(*AlertError).Error":                           method_ptr_AlertError_Error,
		"(*CertificateRequestInfo).Context":             method_ptr_CertificateRequestInfo_Context,
		"(*CertificateRequestInfo).SupportsCertificate": method_ptr_CertificateRequestInfo_SupportsCertificate,
		"(*CertificateVerificationError).Error":         method_ptr_CertificateVerificationError_Error,
		"(*CertificateVerificationError).Unwrap":        method_ptr_CertificateVerificationError_Unwrap,
		"(*ClientAuthType).String":                      method_ptr_ClientAuthType_String,
		"(*ClientHelloInfo).Context":                    method_ptr_ClientHelloInfo_Context,
		"(*ClientHelloInfo).SupportsCertificate":        method_ptr_ClientHelloInfo_SupportsCertificate,
		"(*Config).BuildNameToCertificate":              method_ptr_Config_BuildNameToCertificate,
		"(*Config).Clone":                               method_ptr_Config_Clone,
		"(*Config).SetSessionTicketKeys":                method_ptr_Config_SetSessionTicketKeys,
		"(*Conn).Close":                                 method_ptr_Conn_Close,
		"(*Conn).CloseWrite":                            method_ptr_Conn_CloseWrite,
		"(*Conn).ConnectionState":                       method_ptr_Conn_ConnectionState,
		"(*Conn).Handshake":                             method_ptr_Conn_Handshake,
		"(*Conn).HandshakeContext":                      method_ptr_Conn_HandshakeContext,
		"(*Conn).LocalAddr":                             method_ptr_Conn_LocalAddr,
		"(*Conn).NetConn":                               method_ptr_Conn_NetConn,
		"(*Conn).OCSPResponse":                          method_ptr_Conn_OCSPResponse,
		"(*Conn).RemoteAddr":                            method_ptr_Conn_RemoteAddr,
		"(*Conn).SetDeadline":                           method_ptr_Conn_SetDeadline,
		"(*Conn).SetReadDeadline":                       method_ptr_Conn_SetReadDeadline,
		"(*Conn).SetWriteDeadline":                      method_ptr_Conn_SetWriteDeadline,
		"(*Conn).VerifyHostname":                        method_ptr_Conn_VerifyHostname,
		"(*CurveID).String":                             method_ptr_CurveID_String,
		"(*ECHRejectionError).Error":                    method_ptr_ECHRejectionError_Error,
		"(*QUICConn).Close":                             method_ptr_QUICConn_Close,
		"(*QUICConn).ConnectionState":                   method_ptr_QUICConn_ConnectionState,
		"(*QUICConn).HandleData":                        method_ptr_QUICConn_HandleData,
		"(*QUICConn).NextEvent":                         method_ptr_QUICConn_NextEvent,
		"(*QUICConn).SendSessionTicket":                 method_ptr_QUICConn_SendSessionTicket,
		"(*QUICConn).SetTransportParameters":            method_ptr_QUICConn_SetTransportParameters,
		"(*QUICConn).Start":                             method_ptr_QUICConn_Start,
		"(*QUICConn).StoreSession":                      method_ptr_QUICConn_StoreSession,
		"(*QUICEncryptionLevel).String":                 method_ptr_QUICEncryptionLevel_String,
		"(*RecordHeaderError).Error":                    method_ptr_RecordHeaderError_Error,
		"(*SignatureScheme).String":                     method_ptr_SignatureScheme_String,
		"(AlertError).Error":                            method_AlertError_Error,
		"(ClientAuthType).String":                       method_ClientAuthType_String,
		"(CurveID).String":                              method_CurveID_String,
		"(QUICEncryptionLevel).String":                  method_QUICEncryptionLevel_String,
		"(RecordHeaderError).Error":                     method_RecordHeaderError_Error,
		"(SignatureScheme).String":                      method_SignatureScheme_String,
		"CipherSuiteName":                               func_CipherSuiteName,
		"CipherSuites":                                  func_CipherSuites,
		"Client":                                        func_Client,
		"InsecureCipherSuites":                          func_InsecureCipherSuites,
		"NewLRUClientSessionCache":                      func_NewLRUClientSessionCache,
		"NewListener":                                   func_NewListener,
		"QUICClient":                                    func_QUICClient,
		"QUICServer":                                    func_QUICServer,
		"Server":                                        func_Server,
		"VersionName":                                   func_VersionName,
	})
}

func method_AlertError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AlertError.Error(ixgo.DirectCallArg[q.AlertError](ctx, 0)))
}

func method_ptr_AlertError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AlertError).Error(ixgo.DirectCallArg[*q.AlertError](ctx, 0)))
}

func method_ptr_CertificateRequestInfo_Context(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertificateRequestInfo).Context(ixgo.DirectCallArg[*q.CertificateRequestInfo](ctx, 0)))
}

func method_ptr_CertificateRequestInfo_SupportsCertificate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertificateRequestInfo).SupportsCertificate(ixgo.DirectCallArg[*q.CertificateRequestInfo](ctx, 0), ixgo.DirectCallArg[*q.Certificate](ctx, 1)))
}

func method_ptr_CertificateVerificationError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertificateVerificationError).Error(ixgo.DirectCallArg[*q.CertificateVerificationError](ctx, 0)))
}

func method_ptr_CertificateVerificationError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CertificateVerificationError).Unwrap(ixgo.DirectCallArg[*q.CertificateVerificationError](ctx, 0)))
}

func func_CipherSuiteName(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CipherSuiteName(ixgo.DirectCallArg[uint16](ctx, 0)))
}

func func_CipherSuites(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CipherSuites())
}

func func_Client(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Client(ixgo.DirectCallArg[net.Conn](ctx, 0), ixgo.DirectCallArg[*q.Config](ctx, 1)))
}

func method_ClientAuthType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ClientAuthType.String(ixgo.DirectCallArg[q.ClientAuthType](ctx, 0)))
}

func method_ptr_ClientAuthType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientAuthType).String(ixgo.DirectCallArg[*q.ClientAuthType](ctx, 0)))
}

func method_ptr_ClientHelloInfo_Context(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientHelloInfo).Context(ixgo.DirectCallArg[*q.ClientHelloInfo](ctx, 0)))
}

func method_ptr_ClientHelloInfo_SupportsCertificate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ClientHelloInfo).SupportsCertificate(ixgo.DirectCallArg[*q.ClientHelloInfo](ctx, 0), ixgo.DirectCallArg[*q.Certificate](ctx, 1)))
}

func method_ptr_Config_BuildNameToCertificate(ctx ixgo.DirectCallContext) {
	(*q.Config).BuildNameToCertificate(ixgo.DirectCallArg[*q.Config](ctx, 0))
}

func method_ptr_Config_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Config).Clone(ixgo.DirectCallArg[*q.Config](ctx, 0)))
}

func method_ptr_Config_SetSessionTicketKeys(ctx ixgo.DirectCallContext) {
	(*q.Config).SetSessionTicketKeys(ixgo.DirectCallArg[*q.Config](ctx, 0), ixgo.DirectCallArg[[][32]byte](ctx, 1))
}

func method_ptr_Conn_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).Close(ixgo.DirectCallArg[*q.Conn](ctx, 0)))
}

func method_ptr_Conn_CloseWrite(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).CloseWrite(ixgo.DirectCallArg[*q.Conn](ctx, 0)))
}

func method_ptr_Conn_ConnectionState(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).ConnectionState(ixgo.DirectCallArg[*q.Conn](ctx, 0)))
}

func method_ptr_Conn_Handshake(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).Handshake(ixgo.DirectCallArg[*q.Conn](ctx, 0)))
}

func method_ptr_Conn_HandshakeContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).HandshakeContext(ixgo.DirectCallArg[*q.Conn](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1)))
}

func method_ptr_Conn_LocalAddr(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).LocalAddr(ixgo.DirectCallArg[*q.Conn](ctx, 0)))
}

func method_ptr_Conn_NetConn(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).NetConn(ixgo.DirectCallArg[*q.Conn](ctx, 0)))
}

func method_ptr_Conn_OCSPResponse(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).OCSPResponse(ixgo.DirectCallArg[*q.Conn](ctx, 0)))
}

func method_ptr_Conn_RemoteAddr(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).RemoteAddr(ixgo.DirectCallArg[*q.Conn](ctx, 0)))
}

func method_ptr_Conn_SetDeadline(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).SetDeadline(ixgo.DirectCallArg[*q.Conn](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func method_ptr_Conn_SetReadDeadline(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).SetReadDeadline(ixgo.DirectCallArg[*q.Conn](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func method_ptr_Conn_SetWriteDeadline(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).SetWriteDeadline(ixgo.DirectCallArg[*q.Conn](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func method_ptr_Conn_VerifyHostname(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).VerifyHostname(ixgo.DirectCallArg[*q.Conn](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_CurveID_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CurveID.String(ixgo.DirectCallArg[q.CurveID](ctx, 0)))
}

func method_ptr_CurveID_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CurveID).String(ixgo.DirectCallArg[*q.CurveID](ctx, 0)))
}

func method_ptr_ECHRejectionError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ECHRejectionError).Error(ixgo.DirectCallArg[*q.ECHRejectionError](ctx, 0)))
}

func func_InsecureCipherSuites(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.InsecureCipherSuites())
}

func func_NewLRUClientSessionCache(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewLRUClientSessionCache(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_NewListener(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewListener(ixgo.DirectCallArg[net.Listener](ctx, 0), ixgo.DirectCallArg[*q.Config](ctx, 1)))
}

func func_QUICClient(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.QUICClient(ixgo.DirectCallArg[*q.QUICConfig](ctx, 0)))
}

func method_ptr_QUICConn_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.QUICConn).Close(ixgo.DirectCallArg[*q.QUICConn](ctx, 0)))
}

func method_ptr_QUICConn_ConnectionState(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.QUICConn).ConnectionState(ixgo.DirectCallArg[*q.QUICConn](ctx, 0)))
}

func method_ptr_QUICConn_HandleData(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.QUICConn).HandleData(ixgo.DirectCallArg[*q.QUICConn](ctx, 0), ixgo.DirectCallArg[q.QUICEncryptionLevel](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}

func method_ptr_QUICConn_NextEvent(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.QUICConn).NextEvent(ixgo.DirectCallArg[*q.QUICConn](ctx, 0)))
}

func method_ptr_QUICConn_SendSessionTicket(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.QUICConn).SendSessionTicket(ixgo.DirectCallArg[*q.QUICConn](ctx, 0), ixgo.DirectCallArg[q.QUICSessionTicketOptions](ctx, 1)))
}

func method_ptr_QUICConn_SetTransportParameters(ctx ixgo.DirectCallContext) {
	(*q.QUICConn).SetTransportParameters(ixgo.DirectCallArg[*q.QUICConn](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1))
}

func method_ptr_QUICConn_Start(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.QUICConn).Start(ixgo.DirectCallArg[*q.QUICConn](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1)))
}

func method_ptr_QUICConn_StoreSession(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.QUICConn).StoreSession(ixgo.DirectCallArg[*q.QUICConn](ctx, 0), ixgo.DirectCallArg[*q.SessionState](ctx, 1)))
}

func method_QUICEncryptionLevel_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.QUICEncryptionLevel.String(ixgo.DirectCallArg[q.QUICEncryptionLevel](ctx, 0)))
}

func method_ptr_QUICEncryptionLevel_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.QUICEncryptionLevel).String(ixgo.DirectCallArg[*q.QUICEncryptionLevel](ctx, 0)))
}

func func_QUICServer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.QUICServer(ixgo.DirectCallArg[*q.QUICConfig](ctx, 0)))
}

func method_RecordHeaderError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RecordHeaderError.Error(ixgo.DirectCallArg[q.RecordHeaderError](ctx, 0)))
}

func method_ptr_RecordHeaderError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RecordHeaderError).Error(ixgo.DirectCallArg[*q.RecordHeaderError](ctx, 0)))
}

func func_Server(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Server(ixgo.DirectCallArg[net.Conn](ctx, 0), ixgo.DirectCallArg[*q.Config](ctx, 1)))
}

func method_SignatureScheme_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SignatureScheme.String(ixgo.DirectCallArg[q.SignatureScheme](ctx, 0)))
}

func method_ptr_SignatureScheme_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SignatureScheme).String(ixgo.DirectCallArg[*q.SignatureScheme](ctx, 0)))
}

func func_VersionName(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.VersionName(ixgo.DirectCallArg[uint16](ctx, 0)))
}
