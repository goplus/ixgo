// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package netip

import (
	q "net/netip"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("net/netip", map[string]ixgo.DirectCallAdapter{
		"(*Addr).AppendTo":                  method_ptr_Addr_AppendTo,
		"(*Addr).As16":                      method_ptr_Addr_As16,
		"(*Addr).As4":                       method_ptr_Addr_As4,
		"(*Addr).AsSlice":                   method_ptr_Addr_AsSlice,
		"(*Addr).BitLen":                    method_ptr_Addr_BitLen,
		"(*Addr).Compare":                   method_ptr_Addr_Compare,
		"(*Addr).Is4":                       method_ptr_Addr_Is4,
		"(*Addr).Is4In6":                    method_ptr_Addr_Is4In6,
		"(*Addr).Is6":                       method_ptr_Addr_Is6,
		"(*Addr).IsGlobalUnicast":           method_ptr_Addr_IsGlobalUnicast,
		"(*Addr).IsInterfaceLocalMulticast": method_ptr_Addr_IsInterfaceLocalMulticast,
		"(*Addr).IsLinkLocalMulticast":      method_ptr_Addr_IsLinkLocalMulticast,
		"(*Addr).IsLinkLocalUnicast":        method_ptr_Addr_IsLinkLocalUnicast,
		"(*Addr).IsLoopback":                method_ptr_Addr_IsLoopback,
		"(*Addr).IsMulticast":               method_ptr_Addr_IsMulticast,
		"(*Addr).IsPrivate":                 method_ptr_Addr_IsPrivate,
		"(*Addr).IsUnspecified":             method_ptr_Addr_IsUnspecified,
		"(*Addr).IsValid":                   method_ptr_Addr_IsValid,
		"(*Addr).Less":                      method_ptr_Addr_Less,
		"(*Addr).Next":                      method_ptr_Addr_Next,
		"(*Addr).Prev":                      method_ptr_Addr_Prev,
		"(*Addr).String":                    method_ptr_Addr_String,
		"(*Addr).StringExpanded":            method_ptr_Addr_StringExpanded,
		"(*Addr).Unmap":                     method_ptr_Addr_Unmap,
		"(*Addr).UnmarshalBinary":           method_ptr_Addr_UnmarshalBinary,
		"(*Addr).UnmarshalText":             method_ptr_Addr_UnmarshalText,
		"(*Addr).WithZone":                  method_ptr_Addr_WithZone,
		"(*Addr).Zone":                      method_ptr_Addr_Zone,
		"(*AddrPort).Addr":                  method_ptr_AddrPort_Addr,
		"(*AddrPort).AppendTo":              method_ptr_AddrPort_AppendTo,
		"(*AddrPort).Compare":               method_ptr_AddrPort_Compare,
		"(*AddrPort).IsValid":               method_ptr_AddrPort_IsValid,
		"(*AddrPort).Port":                  method_ptr_AddrPort_Port,
		"(*AddrPort).String":                method_ptr_AddrPort_String,
		"(*AddrPort).UnmarshalBinary":       method_ptr_AddrPort_UnmarshalBinary,
		"(*AddrPort).UnmarshalText":         method_ptr_AddrPort_UnmarshalText,
		"(*Prefix).Addr":                    method_ptr_Prefix_Addr,
		"(*Prefix).AppendTo":                method_ptr_Prefix_AppendTo,
		"(*Prefix).Bits":                    method_ptr_Prefix_Bits,
		"(*Prefix).Compare":                 method_ptr_Prefix_Compare,
		"(*Prefix).Contains":                method_ptr_Prefix_Contains,
		"(*Prefix).IsSingleIP":              method_ptr_Prefix_IsSingleIP,
		"(*Prefix).IsValid":                 method_ptr_Prefix_IsValid,
		"(*Prefix).Masked":                  method_ptr_Prefix_Masked,
		"(*Prefix).Overlaps":                method_ptr_Prefix_Overlaps,
		"(*Prefix).String":                  method_ptr_Prefix_String,
		"(*Prefix).UnmarshalBinary":         method_ptr_Prefix_UnmarshalBinary,
		"(*Prefix).UnmarshalText":           method_ptr_Prefix_UnmarshalText,
		"(Addr).AppendTo":                   method_Addr_AppendTo,
		"(Addr).As16":                       method_Addr_As16,
		"(Addr).As4":                        method_Addr_As4,
		"(Addr).AsSlice":                    method_Addr_AsSlice,
		"(Addr).BitLen":                     method_Addr_BitLen,
		"(Addr).Compare":                    method_Addr_Compare,
		"(Addr).Is4":                        method_Addr_Is4,
		"(Addr).Is4In6":                     method_Addr_Is4In6,
		"(Addr).Is6":                        method_Addr_Is6,
		"(Addr).IsGlobalUnicast":            method_Addr_IsGlobalUnicast,
		"(Addr).IsInterfaceLocalMulticast":  method_Addr_IsInterfaceLocalMulticast,
		"(Addr).IsLinkLocalMulticast":       method_Addr_IsLinkLocalMulticast,
		"(Addr).IsLinkLocalUnicast":         method_Addr_IsLinkLocalUnicast,
		"(Addr).IsLoopback":                 method_Addr_IsLoopback,
		"(Addr).IsMulticast":                method_Addr_IsMulticast,
		"(Addr).IsPrivate":                  method_Addr_IsPrivate,
		"(Addr).IsUnspecified":              method_Addr_IsUnspecified,
		"(Addr).IsValid":                    method_Addr_IsValid,
		"(Addr).Less":                       method_Addr_Less,
		"(Addr).Next":                       method_Addr_Next,
		"(Addr).Prev":                       method_Addr_Prev,
		"(Addr).String":                     method_Addr_String,
		"(Addr).StringExpanded":             method_Addr_StringExpanded,
		"(Addr).Unmap":                      method_Addr_Unmap,
		"(Addr).WithZone":                   method_Addr_WithZone,
		"(Addr).Zone":                       method_Addr_Zone,
		"(AddrPort).Addr":                   method_AddrPort_Addr,
		"(AddrPort).AppendTo":               method_AddrPort_AppendTo,
		"(AddrPort).Compare":                method_AddrPort_Compare,
		"(AddrPort).IsValid":                method_AddrPort_IsValid,
		"(AddrPort).Port":                   method_AddrPort_Port,
		"(AddrPort).String":                 method_AddrPort_String,
		"(Prefix).Addr":                     method_Prefix_Addr,
		"(Prefix).AppendTo":                 method_Prefix_AppendTo,
		"(Prefix).Bits":                     method_Prefix_Bits,
		"(Prefix).Compare":                  method_Prefix_Compare,
		"(Prefix).Contains":                 method_Prefix_Contains,
		"(Prefix).IsSingleIP":               method_Prefix_IsSingleIP,
		"(Prefix).IsValid":                  method_Prefix_IsValid,
		"(Prefix).Masked":                   method_Prefix_Masked,
		"(Prefix).Overlaps":                 method_Prefix_Overlaps,
		"(Prefix).String":                   method_Prefix_String,
		"AddrFrom16":                        func_AddrFrom16,
		"AddrFrom4":                         func_AddrFrom4,
		"AddrPortFrom":                      func_AddrPortFrom,
		"IPv4Unspecified":                   func_IPv4Unspecified,
		"IPv6LinkLocalAllNodes":             func_IPv6LinkLocalAllNodes,
		"IPv6LinkLocalAllRouters":           func_IPv6LinkLocalAllRouters,
		"IPv6Loopback":                      func_IPv6Loopback,
		"IPv6Unspecified":                   func_IPv6Unspecified,
		"MustParseAddr":                     func_MustParseAddr,
		"MustParseAddrPort":                 func_MustParseAddrPort,
		"MustParsePrefix":                   func_MustParsePrefix,
		"PrefixFrom":                        func_PrefixFrom,
	})
}

func method_Addr_AppendTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.AppendTo(ixgo.DirectCallArg[q.Addr](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Addr_AppendTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).AppendTo(ixgo.DirectCallArg[*q.Addr](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_Addr_As16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.As16(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_As16(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).As16(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_As4(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.As4(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_As4(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).As4(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_AsSlice(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.AsSlice(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_AsSlice(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).AsSlice(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_BitLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.BitLen(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_BitLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).BitLen(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.Compare(ixgo.DirectCallArg[q.Addr](ctx, 0), ixgo.DirectCallArg[q.Addr](ctx, 1)))
}

func method_ptr_Addr_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).Compare(ixgo.DirectCallArg[*q.Addr](ctx, 0), ixgo.DirectCallArg[q.Addr](ctx, 1)))
}

func method_Addr_Is4(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.Is4(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_Is4(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).Is4(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_Is4In6(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.Is4In6(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_Is4In6(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).Is4In6(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_Is6(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.Is6(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_Is6(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).Is6(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_IsGlobalUnicast(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.IsGlobalUnicast(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_IsGlobalUnicast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).IsGlobalUnicast(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_IsInterfaceLocalMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.IsInterfaceLocalMulticast(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_IsInterfaceLocalMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).IsInterfaceLocalMulticast(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_IsLinkLocalMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.IsLinkLocalMulticast(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_IsLinkLocalMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).IsLinkLocalMulticast(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_IsLinkLocalUnicast(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.IsLinkLocalUnicast(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_IsLinkLocalUnicast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).IsLinkLocalUnicast(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_IsLoopback(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.IsLoopback(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_IsLoopback(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).IsLoopback(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_IsMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.IsMulticast(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_IsMulticast(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).IsMulticast(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_IsPrivate(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.IsPrivate(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_IsPrivate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).IsPrivate(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_IsUnspecified(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.IsUnspecified(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_IsUnspecified(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).IsUnspecified(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.IsValid(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).IsValid(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_Less(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.Less(ixgo.DirectCallArg[q.Addr](ctx, 0), ixgo.DirectCallArg[q.Addr](ctx, 1)))
}

func method_ptr_Addr_Less(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).Less(ixgo.DirectCallArg[*q.Addr](ctx, 0), ixgo.DirectCallArg[q.Addr](ctx, 1)))
}

func method_Addr_Next(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.Next(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_Next(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).Next(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_Prev(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.Prev(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_Prev(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).Prev(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.String(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).String(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_StringExpanded(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.StringExpanded(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_StringExpanded(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).StringExpanded(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_Addr_Unmap(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.Unmap(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_Unmap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).Unmap(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func method_ptr_Addr_UnmarshalBinary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).UnmarshalBinary(ixgo.DirectCallArg[*q.Addr](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Addr_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).UnmarshalText(ixgo.DirectCallArg[*q.Addr](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_Addr_WithZone(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.WithZone(ixgo.DirectCallArg[q.Addr](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Addr_WithZone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).WithZone(ixgo.DirectCallArg[*q.Addr](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_Addr_Zone(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Addr.Zone(ixgo.DirectCallArg[q.Addr](ctx, 0)))
}

func method_ptr_Addr_Zone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Addr).Zone(ixgo.DirectCallArg[*q.Addr](ctx, 0)))
}

func func_AddrFrom16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddrFrom16(ixgo.DirectCallArg[[16]byte](ctx, 0)))
}

func func_AddrFrom4(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddrFrom4(ixgo.DirectCallArg[[4]byte](ctx, 0)))
}

func method_AddrPort_Addr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddrPort.Addr(ixgo.DirectCallArg[q.AddrPort](ctx, 0)))
}

func method_ptr_AddrPort_Addr(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrPort).Addr(ixgo.DirectCallArg[*q.AddrPort](ctx, 0)))
}

func method_AddrPort_AppendTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddrPort.AppendTo(ixgo.DirectCallArg[q.AddrPort](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_AddrPort_AppendTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrPort).AppendTo(ixgo.DirectCallArg[*q.AddrPort](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_AddrPort_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddrPort.Compare(ixgo.DirectCallArg[q.AddrPort](ctx, 0), ixgo.DirectCallArg[q.AddrPort](ctx, 1)))
}

func method_ptr_AddrPort_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrPort).Compare(ixgo.DirectCallArg[*q.AddrPort](ctx, 0), ixgo.DirectCallArg[q.AddrPort](ctx, 1)))
}

func method_AddrPort_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddrPort.IsValid(ixgo.DirectCallArg[q.AddrPort](ctx, 0)))
}

func method_ptr_AddrPort_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrPort).IsValid(ixgo.DirectCallArg[*q.AddrPort](ctx, 0)))
}

func method_AddrPort_Port(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddrPort.Port(ixgo.DirectCallArg[q.AddrPort](ctx, 0)))
}

func method_ptr_AddrPort_Port(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrPort).Port(ixgo.DirectCallArg[*q.AddrPort](ctx, 0)))
}

func method_AddrPort_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddrPort.String(ixgo.DirectCallArg[q.AddrPort](ctx, 0)))
}

func method_ptr_AddrPort_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrPort).String(ixgo.DirectCallArg[*q.AddrPort](ctx, 0)))
}

func method_ptr_AddrPort_UnmarshalBinary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrPort).UnmarshalBinary(ixgo.DirectCallArg[*q.AddrPort](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_AddrPort_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AddrPort).UnmarshalText(ixgo.DirectCallArg[*q.AddrPort](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_AddrPortFrom(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddrPortFrom(ixgo.DirectCallArg[q.Addr](ctx, 0), ixgo.DirectCallArg[uint16](ctx, 1)))
}

func func_IPv4Unspecified(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IPv4Unspecified())
}

func func_IPv6LinkLocalAllNodes(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IPv6LinkLocalAllNodes())
}

func func_IPv6LinkLocalAllRouters(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IPv6LinkLocalAllRouters())
}

func func_IPv6Loopback(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IPv6Loopback())
}

func func_IPv6Unspecified(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IPv6Unspecified())
}

func func_MustParseAddr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MustParseAddr(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_MustParseAddrPort(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MustParseAddrPort(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_MustParsePrefix(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MustParsePrefix(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_Prefix_Addr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Prefix.Addr(ixgo.DirectCallArg[q.Prefix](ctx, 0)))
}

func method_ptr_Prefix_Addr(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).Addr(ixgo.DirectCallArg[*q.Prefix](ctx, 0)))
}

func method_Prefix_AppendTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Prefix.AppendTo(ixgo.DirectCallArg[q.Prefix](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Prefix_AppendTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).AppendTo(ixgo.DirectCallArg[*q.Prefix](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_Prefix_Bits(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Prefix.Bits(ixgo.DirectCallArg[q.Prefix](ctx, 0)))
}

func method_ptr_Prefix_Bits(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).Bits(ixgo.DirectCallArg[*q.Prefix](ctx, 0)))
}

func method_Prefix_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Prefix.Compare(ixgo.DirectCallArg[q.Prefix](ctx, 0), ixgo.DirectCallArg[q.Prefix](ctx, 1)))
}

func method_ptr_Prefix_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).Compare(ixgo.DirectCallArg[*q.Prefix](ctx, 0), ixgo.DirectCallArg[q.Prefix](ctx, 1)))
}

func method_Prefix_Contains(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Prefix.Contains(ixgo.DirectCallArg[q.Prefix](ctx, 0), ixgo.DirectCallArg[q.Addr](ctx, 1)))
}

func method_ptr_Prefix_Contains(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).Contains(ixgo.DirectCallArg[*q.Prefix](ctx, 0), ixgo.DirectCallArg[q.Addr](ctx, 1)))
}

func method_Prefix_IsSingleIP(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Prefix.IsSingleIP(ixgo.DirectCallArg[q.Prefix](ctx, 0)))
}

func method_ptr_Prefix_IsSingleIP(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).IsSingleIP(ixgo.DirectCallArg[*q.Prefix](ctx, 0)))
}

func method_Prefix_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Prefix.IsValid(ixgo.DirectCallArg[q.Prefix](ctx, 0)))
}

func method_ptr_Prefix_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).IsValid(ixgo.DirectCallArg[*q.Prefix](ctx, 0)))
}

func method_Prefix_Masked(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Prefix.Masked(ixgo.DirectCallArg[q.Prefix](ctx, 0)))
}

func method_ptr_Prefix_Masked(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).Masked(ixgo.DirectCallArg[*q.Prefix](ctx, 0)))
}

func method_Prefix_Overlaps(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Prefix.Overlaps(ixgo.DirectCallArg[q.Prefix](ctx, 0), ixgo.DirectCallArg[q.Prefix](ctx, 1)))
}

func method_ptr_Prefix_Overlaps(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).Overlaps(ixgo.DirectCallArg[*q.Prefix](ctx, 0), ixgo.DirectCallArg[q.Prefix](ctx, 1)))
}

func method_Prefix_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Prefix.String(ixgo.DirectCallArg[q.Prefix](ctx, 0)))
}

func method_ptr_Prefix_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).String(ixgo.DirectCallArg[*q.Prefix](ctx, 0)))
}

func method_ptr_Prefix_UnmarshalBinary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).UnmarshalBinary(ixgo.DirectCallArg[*q.Prefix](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Prefix_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prefix).UnmarshalText(ixgo.DirectCallArg[*q.Prefix](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_PrefixFrom(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PrefixFrom(ixgo.DirectCallArg[q.Addr](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}
