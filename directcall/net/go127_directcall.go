// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package net

import (
	q "net"

	"github.com/goplus/ixgo"
	netip "net/netip"
	time "time"
)

func init() {
	ixgo.RegisterDirectCalls("net", map[string]ixgo.DirectCallAdapter{
		"(*AddrError).Error":               method_ptr_AddrError_Error,
		"(*AddrError).Temporary":           method_ptr_AddrError_Temporary,
		"(*AddrError).Timeout":             method_ptr_AddrError_Timeout,
		"(*DNSConfigError).Error":          method_ptr_DNSConfigError_Error,
		"(*DNSConfigError).Temporary":      method_ptr_DNSConfigError_Temporary,
		"(*DNSConfigError).Timeout":        method_ptr_DNSConfigError_Timeout,
		"(*DNSConfigError).Unwrap":         method_ptr_DNSConfigError_Unwrap,
		"(*DNSError).Error":                method_ptr_DNSError_Error,
		"(*DNSError).Temporary":            method_ptr_DNSError_Temporary,
		"(*DNSError).Timeout":              method_ptr_DNSError_Timeout,
		"(*DNSError).Unwrap":               method_ptr_DNSError_Unwrap,
		"(*Dialer).MultipathTCP":           method_ptr_Dialer_MultipathTCP,
		"(*Dialer).SetMultipathTCP":        method_ptr_Dialer_SetMultipathTCP,
		"(*Flags).String":                  method_ptr_Flags_String,
		"(*HardwareAddr).String":           method_ptr_HardwareAddr_String,
		"(*IP).DefaultMask":                method_ptr_IP_DefaultMask,
		"(*IP).Equal":                      method_ptr_IP_Equal,
		"(*IP).IsGlobalUnicast":            method_ptr_IP_IsGlobalUnicast,
		"(*IP).IsInterfaceLocalMulticast":  method_ptr_IP_IsInterfaceLocalMulticast,
		"(*IP).IsLinkLocalMulticast":       method_ptr_IP_IsLinkLocalMulticast,
		"(*IP).IsLinkLocalUnicast":         method_ptr_IP_IsLinkLocalUnicast,
		"(*IP).IsLoopback":                 method_ptr_IP_IsLoopback,
		"(*IP).IsMulticast":                method_ptr_IP_IsMulticast,
		"(*IP).IsPrivate":                  method_ptr_IP_IsPrivate,
		"(*IP).IsUnspecified":              method_ptr_IP_IsUnspecified,
		"(*IP).Mask":                       method_ptr_IP_Mask,
		"(*IP).String":                     method_ptr_IP_String,
		"(*IP).To16":                       method_ptr_IP_To16,
		"(*IP).To4":                        method_ptr_IP_To4,
		"(*IP).UnmarshalText":              method_ptr_IP_UnmarshalText,
		"(*IPAddr).Network":                method_ptr_IPAddr_Network,
		"(*IPAddr).String":                 method_ptr_IPAddr_String,
		"(*IPMask).String":                 method_ptr_IPMask_String,
		"(*IPNet).Contains":                method_ptr_IPNet_Contains,
		"(*IPNet).Network":                 method_ptr_IPNet_Network,
		"(*IPNet).String":                  method_ptr_IPNet_String,
		"(*InvalidAddrError).Error":        method_ptr_InvalidAddrError_Error,
		"(*InvalidAddrError).Temporary":    method_ptr_InvalidAddrError_Temporary,
		"(*InvalidAddrError).Timeout":      method_ptr_InvalidAddrError_Timeout,
		"(*ListenConfig).MultipathTCP":     method_ptr_ListenConfig_MultipathTCP,
		"(*ListenConfig).SetMultipathTCP":  method_ptr_ListenConfig_SetMultipathTCP,
		"(*OpError).Error":                 method_ptr_OpError_Error,
		"(*OpError).Temporary":             method_ptr_OpError_Temporary,
		"(*OpError).Timeout":               method_ptr_OpError_Timeout,
		"(*OpError).Unwrap":                method_ptr_OpError_Unwrap,
		"(*ParseError).Error":              method_ptr_ParseError_Error,
		"(*ParseError).Temporary":          method_ptr_ParseError_Temporary,
		"(*ParseError).Timeout":            method_ptr_ParseError_Timeout,
		"(*TCPAddr).AddrPort":              method_ptr_TCPAddr_AddrPort,
		"(*TCPAddr).Network":               method_ptr_TCPAddr_Network,
		"(*TCPAddr).String":                method_ptr_TCPAddr_String,
		"(*TCPConn).CloseRead":             method_ptr_TCPConn_CloseRead,
		"(*TCPConn).CloseWrite":            method_ptr_TCPConn_CloseWrite,
		"(*TCPConn).SetKeepAlive":          method_ptr_TCPConn_SetKeepAlive,
		"(*TCPConn).SetKeepAliveConfig":    method_ptr_TCPConn_SetKeepAliveConfig,
		"(*TCPConn).SetKeepAlivePeriod":    method_ptr_TCPConn_SetKeepAlivePeriod,
		"(*TCPConn).SetLinger":             method_ptr_TCPConn_SetLinger,
		"(*TCPConn).SetNoDelay":            method_ptr_TCPConn_SetNoDelay,
		"(*TCPListener).Addr":              method_ptr_TCPListener_Addr,
		"(*TCPListener).Close":             method_ptr_TCPListener_Close,
		"(*TCPListener).SetDeadline":       method_ptr_TCPListener_SetDeadline,
		"(*UDPAddr).AddrPort":              method_ptr_UDPAddr_AddrPort,
		"(*UDPAddr).Network":               method_ptr_UDPAddr_Network,
		"(*UDPAddr).String":                method_ptr_UDPAddr_String,
		"(*UnixAddr).Network":              method_ptr_UnixAddr_Network,
		"(*UnixAddr).String":               method_ptr_UnixAddr_String,
		"(*UnixConn).CloseRead":            method_ptr_UnixConn_CloseRead,
		"(*UnixConn).CloseWrite":           method_ptr_UnixConn_CloseWrite,
		"(*UnixListener).Addr":             method_ptr_UnixListener_Addr,
		"(*UnixListener).Close":            method_ptr_UnixListener_Close,
		"(*UnixListener).SetDeadline":      method_ptr_UnixListener_SetDeadline,
		"(*UnixListener).SetUnlinkOnClose": method_ptr_UnixListener_SetUnlinkOnClose,
		"(*UnknownNetworkError).Error":     method_ptr_UnknownNetworkError_Error,
		"(*UnknownNetworkError).Temporary": method_ptr_UnknownNetworkError_Temporary,
		"(*UnknownNetworkError).Timeout":   method_ptr_UnknownNetworkError_Timeout,
		"(Flags).String":                   method_Flags_String,
		"(HardwareAddr).String":            method_HardwareAddr_String,
		"(IP).DefaultMask":                 method_IP_DefaultMask,
		"(IP).Equal":                       method_IP_Equal,
		"(IP).IsGlobalUnicast":             method_IP_IsGlobalUnicast,
		"(IP).IsInterfaceLocalMulticast":   method_IP_IsInterfaceLocalMulticast,
		"(IP).IsLinkLocalMulticast":        method_IP_IsLinkLocalMulticast,
		"(IP).IsLinkLocalUnicast":          method_IP_IsLinkLocalUnicast,
		"(IP).IsLoopback":                  method_IP_IsLoopback,
		"(IP).IsMulticast":                 method_IP_IsMulticast,
		"(IP).IsPrivate":                   method_IP_IsPrivate,
		"(IP).IsUnspecified":               method_IP_IsUnspecified,
		"(IP).Mask":                        method_IP_Mask,
		"(IP).String":                      method_IP_String,
		"(IP).To16":                        method_IP_To16,
		"(IP).To4":                         method_IP_To4,
		"(IPMask).String":                  method_IPMask_String,
		"(InvalidAddrError).Error":         method_InvalidAddrError_Error,
		"(InvalidAddrError).Temporary":     method_InvalidAddrError_Temporary,
		"(InvalidAddrError).Timeout":       method_InvalidAddrError_Timeout,
		"(UnknownNetworkError).Error":      method_UnknownNetworkError_Error,
		"(UnknownNetworkError).Temporary":  method_UnknownNetworkError_Temporary,
		"(UnknownNetworkError).Timeout":    method_UnknownNetworkError_Timeout,
		"CIDRMask":                         func_CIDRMask,
		"IPv4":                             func_IPv4,
		"IPv4Mask":                         func_IPv4Mask,
		"JoinHostPort":                     func_JoinHostPort,
		"ParseIP":                          func_ParseIP,
		"TCPAddrFromAddrPort":              func_TCPAddrFromAddrPort,
		"UDPAddrFromAddrPort":              func_UDPAddrFromAddrPort,
	})
}

func method_ptr_AddrError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrError).Error(ixgo.DirectCallArg[*q.AddrError](ctx, 0)))
}

func method_ptr_AddrError_Temporary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrError).Temporary(ixgo.DirectCallArg[*q.AddrError](ctx, 0)))
}

func method_ptr_AddrError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrError).Timeout(ixgo.DirectCallArg[*q.AddrError](ctx, 0)))
}

func func_CIDRMask(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CIDRMask(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_DNSConfigError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DNSConfigError).Error(ixgo.DirectCallArg[*q.DNSConfigError](ctx, 0)))
}

func method_ptr_DNSConfigError_Temporary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DNSConfigError).Temporary(ixgo.DirectCallArg[*q.DNSConfigError](ctx, 0)))
}

func method_ptr_DNSConfigError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DNSConfigError).Timeout(ixgo.DirectCallArg[*q.DNSConfigError](ctx, 0)))
}

func method_ptr_DNSConfigError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DNSConfigError).Unwrap(ixgo.DirectCallArg[*q.DNSConfigError](ctx, 0)))
}

func method_ptr_DNSError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DNSError).Error(ixgo.DirectCallArg[*q.DNSError](ctx, 0)))
}

func method_ptr_DNSError_Temporary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DNSError).Temporary(ixgo.DirectCallArg[*q.DNSError](ctx, 0)))
}

func method_ptr_DNSError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DNSError).Timeout(ixgo.DirectCallArg[*q.DNSError](ctx, 0)))
}

func method_ptr_DNSError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DNSError).Unwrap(ixgo.DirectCallArg[*q.DNSError](ctx, 0)))
}

func method_ptr_Dialer_MultipathTCP(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Dialer).MultipathTCP(ixgo.DirectCallArg[*q.Dialer](ctx, 0)))
}

func method_ptr_Dialer_SetMultipathTCP(ctx ixgo.DirectCallContext) {
	(*q.Dialer).SetMultipathTCP(ixgo.DirectCallArg[*q.Dialer](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_Flags_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Flags.String(ixgo.DirectCallArg[q.Flags](ctx, 0)))
}

func method_ptr_Flags_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Flags).String(ixgo.DirectCallArg[*q.Flags](ctx, 0)))
}

func method_HardwareAddr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HardwareAddr.String(ixgo.DirectCallArg[q.HardwareAddr](ctx, 0)))
}

func method_ptr_HardwareAddr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.HardwareAddr).String(ixgo.DirectCallArg[*q.HardwareAddr](ctx, 0)))
}

func method_IP_DefaultMask(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.DefaultMask(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_DefaultMask(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).DefaultMask(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.Equal(ixgo.DirectCallArg[q.IP](ctx, 0), ixgo.DirectCallArg[q.IP](ctx, 1)))
}

func method_ptr_IP_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).Equal(ixgo.DirectCallArg[*q.IP](ctx, 0), ixgo.DirectCallArg[q.IP](ctx, 1)))
}

func method_IP_IsGlobalUnicast(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.IsGlobalUnicast(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_IsGlobalUnicast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).IsGlobalUnicast(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_IsInterfaceLocalMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.IsInterfaceLocalMulticast(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_IsInterfaceLocalMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).IsInterfaceLocalMulticast(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_IsLinkLocalMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.IsLinkLocalMulticast(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_IsLinkLocalMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).IsLinkLocalMulticast(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_IsLinkLocalUnicast(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.IsLinkLocalUnicast(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_IsLinkLocalUnicast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).IsLinkLocalUnicast(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_IsLoopback(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.IsLoopback(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_IsLoopback(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).IsLoopback(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_IsMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.IsMulticast(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_IsMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).IsMulticast(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_IsPrivate(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.IsPrivate(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_IsPrivate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).IsPrivate(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_IsUnspecified(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.IsUnspecified(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_IsUnspecified(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).IsUnspecified(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_Mask(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.Mask(ixgo.DirectCallArg[q.IP](ctx, 0), ixgo.DirectCallArg[q.IPMask](ctx, 1)))
}

func method_ptr_IP_Mask(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).Mask(ixgo.DirectCallArg[*q.IP](ctx, 0), ixgo.DirectCallArg[q.IPMask](ctx, 1)))
}

func method_IP_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.String(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).String(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_To16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.To16(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_To16(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).To16(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_IP_To4(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IP.To4(ixgo.DirectCallArg[q.IP](ctx, 0)))
}

func method_ptr_IP_To4(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).To4(ixgo.DirectCallArg[*q.IP](ctx, 0)))
}

func method_ptr_IP_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IP).UnmarshalText(ixgo.DirectCallArg[*q.IP](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_IPAddr_Network(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IPAddr).Network(ixgo.DirectCallArg[*q.IPAddr](ctx, 0)))
}

func method_ptr_IPAddr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IPAddr).String(ixgo.DirectCallArg[*q.IPAddr](ctx, 0)))
}

func method_IPMask_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IPMask.String(ixgo.DirectCallArg[q.IPMask](ctx, 0)))
}

func method_ptr_IPMask_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IPMask).String(ixgo.DirectCallArg[*q.IPMask](ctx, 0)))
}

func method_ptr_IPNet_Contains(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IPNet).Contains(ixgo.DirectCallArg[*q.IPNet](ctx, 0), ixgo.DirectCallArg[q.IP](ctx, 1)))
}

func method_ptr_IPNet_Network(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IPNet).Network(ixgo.DirectCallArg[*q.IPNet](ctx, 0)))
}

func method_ptr_IPNet_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IPNet).String(ixgo.DirectCallArg[*q.IPNet](ctx, 0)))
}

func func_IPv4(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IPv4(ixgo.DirectCallArg[byte](ctx, 0), ixgo.DirectCallArg[byte](ctx, 1), ixgo.DirectCallArg[byte](ctx, 2), ixgo.DirectCallArg[byte](ctx, 3)))
}

func func_IPv4Mask(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IPv4Mask(ixgo.DirectCallArg[byte](ctx, 0), ixgo.DirectCallArg[byte](ctx, 1), ixgo.DirectCallArg[byte](ctx, 2), ixgo.DirectCallArg[byte](ctx, 3)))
}

func method_InvalidAddrError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.InvalidAddrError.Error(ixgo.DirectCallArg[q.InvalidAddrError](ctx, 0)))
}

func method_ptr_InvalidAddrError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InvalidAddrError).Error(ixgo.DirectCallArg[*q.InvalidAddrError](ctx, 0)))
}

func method_InvalidAddrError_Temporary(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.InvalidAddrError.Temporary(ixgo.DirectCallArg[q.InvalidAddrError](ctx, 0)))
}

func method_ptr_InvalidAddrError_Temporary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InvalidAddrError).Temporary(ixgo.DirectCallArg[*q.InvalidAddrError](ctx, 0)))
}

func method_InvalidAddrError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.InvalidAddrError.Timeout(ixgo.DirectCallArg[q.InvalidAddrError](ctx, 0)))
}

func method_ptr_InvalidAddrError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InvalidAddrError).Timeout(ixgo.DirectCallArg[*q.InvalidAddrError](ctx, 0)))
}

func func_JoinHostPort(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.JoinHostPort(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_ListenConfig_MultipathTCP(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ListenConfig).MultipathTCP(ixgo.DirectCallArg[*q.ListenConfig](ctx, 0)))
}

func method_ptr_ListenConfig_SetMultipathTCP(ctx ixgo.DirectCallContext) {
	(*q.ListenConfig).SetMultipathTCP(ixgo.DirectCallArg[*q.ListenConfig](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_ptr_OpError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OpError).Error(ixgo.DirectCallArg[*q.OpError](ctx, 0)))
}

func method_ptr_OpError_Temporary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OpError).Temporary(ixgo.DirectCallArg[*q.OpError](ctx, 0)))
}

func method_ptr_OpError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OpError).Timeout(ixgo.DirectCallArg[*q.OpError](ctx, 0)))
}

func method_ptr_OpError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OpError).Unwrap(ixgo.DirectCallArg[*q.OpError](ctx, 0)))
}

func method_ptr_ParseError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ParseError).Error(ixgo.DirectCallArg[*q.ParseError](ctx, 0)))
}

func method_ptr_ParseError_Temporary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ParseError).Temporary(ixgo.DirectCallArg[*q.ParseError](ctx, 0)))
}

func method_ptr_ParseError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ParseError).Timeout(ixgo.DirectCallArg[*q.ParseError](ctx, 0)))
}

func func_ParseIP(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ParseIP(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_TCPAddr_AddrPort(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPAddr).AddrPort(ixgo.DirectCallArg[*q.TCPAddr](ctx, 0)))
}

func method_ptr_TCPAddr_Network(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPAddr).Network(ixgo.DirectCallArg[*q.TCPAddr](ctx, 0)))
}

func method_ptr_TCPAddr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPAddr).String(ixgo.DirectCallArg[*q.TCPAddr](ctx, 0)))
}

func func_TCPAddrFromAddrPort(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TCPAddrFromAddrPort(ixgo.DirectCallArg[netip.AddrPort](ctx, 0)))
}

func method_ptr_TCPConn_CloseRead(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPConn).CloseRead(ixgo.DirectCallArg[*q.TCPConn](ctx, 0)))
}

func method_ptr_TCPConn_CloseWrite(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPConn).CloseWrite(ixgo.DirectCallArg[*q.TCPConn](ctx, 0)))
}

func method_ptr_TCPConn_SetKeepAlive(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPConn).SetKeepAlive(ixgo.DirectCallArg[*q.TCPConn](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1)))
}

func method_ptr_TCPConn_SetKeepAliveConfig(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPConn).SetKeepAliveConfig(ixgo.DirectCallArg[*q.TCPConn](ctx, 0), ixgo.DirectCallArg[q.KeepAliveConfig](ctx, 1)))
}

func method_ptr_TCPConn_SetKeepAlivePeriod(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPConn).SetKeepAlivePeriod(ixgo.DirectCallArg[*q.TCPConn](ctx, 0), ixgo.DirectCallArg[time.Duration](ctx, 1)))
}

func method_ptr_TCPConn_SetLinger(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPConn).SetLinger(ixgo.DirectCallArg[*q.TCPConn](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_TCPConn_SetNoDelay(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPConn).SetNoDelay(ixgo.DirectCallArg[*q.TCPConn](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1)))
}

func method_ptr_TCPListener_Addr(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPListener).Addr(ixgo.DirectCallArg[*q.TCPListener](ctx, 0)))
}

func method_ptr_TCPListener_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPListener).Close(ixgo.DirectCallArg[*q.TCPListener](ctx, 0)))
}

func method_ptr_TCPListener_SetDeadline(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TCPListener).SetDeadline(ixgo.DirectCallArg[*q.TCPListener](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func method_ptr_UDPAddr_AddrPort(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UDPAddr).AddrPort(ixgo.DirectCallArg[*q.UDPAddr](ctx, 0)))
}

func method_ptr_UDPAddr_Network(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UDPAddr).Network(ixgo.DirectCallArg[*q.UDPAddr](ctx, 0)))
}

func method_ptr_UDPAddr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UDPAddr).String(ixgo.DirectCallArg[*q.UDPAddr](ctx, 0)))
}

func func_UDPAddrFromAddrPort(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UDPAddrFromAddrPort(ixgo.DirectCallArg[netip.AddrPort](ctx, 0)))
}

func method_ptr_UnixAddr_Network(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnixAddr).Network(ixgo.DirectCallArg[*q.UnixAddr](ctx, 0)))
}

func method_ptr_UnixAddr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnixAddr).String(ixgo.DirectCallArg[*q.UnixAddr](ctx, 0)))
}

func method_ptr_UnixConn_CloseRead(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnixConn).CloseRead(ixgo.DirectCallArg[*q.UnixConn](ctx, 0)))
}

func method_ptr_UnixConn_CloseWrite(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnixConn).CloseWrite(ixgo.DirectCallArg[*q.UnixConn](ctx, 0)))
}

func method_ptr_UnixListener_Addr(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnixListener).Addr(ixgo.DirectCallArg[*q.UnixListener](ctx, 0)))
}

func method_ptr_UnixListener_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnixListener).Close(ixgo.DirectCallArg[*q.UnixListener](ctx, 0)))
}

func method_ptr_UnixListener_SetDeadline(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnixListener).SetDeadline(ixgo.DirectCallArg[*q.UnixListener](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func method_ptr_UnixListener_SetUnlinkOnClose(ctx ixgo.DirectCallContext) {
	(*q.UnixListener).SetUnlinkOnClose(ixgo.DirectCallArg[*q.UnixListener](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_UnknownNetworkError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnknownNetworkError.Error(ixgo.DirectCallArg[q.UnknownNetworkError](ctx, 0)))
}

func method_ptr_UnknownNetworkError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnknownNetworkError).Error(ixgo.DirectCallArg[*q.UnknownNetworkError](ctx, 0)))
}

func method_UnknownNetworkError_Temporary(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnknownNetworkError.Temporary(ixgo.DirectCallArg[q.UnknownNetworkError](ctx, 0)))
}

func method_ptr_UnknownNetworkError_Temporary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnknownNetworkError).Temporary(ixgo.DirectCallArg[*q.UnknownNetworkError](ctx, 0)))
}

func method_UnknownNetworkError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnknownNetworkError.Timeout(ixgo.DirectCallArg[q.UnknownNetworkError](ctx, 0)))
}

func method_ptr_UnknownNetworkError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnknownNetworkError).Timeout(ixgo.DirectCallArg[*q.UnknownNetworkError](ctx, 0)))
}
