// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package netip

import (
	q "net/netip"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("net/netip", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "netip",
			Path: "net/netip",
			Deps: map[string]string{
				"cmp":                "cmp",
				"errors":             "errors",
				"internal/bytealg":   "bytealg",
				"internal/byteorder": "byteorder",
				"internal/itoa":      "itoa",
				"math":               "math",
				"math/bits":          "bits",
				"strconv":            "strconv",
				"unique":             "unique",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Addr":     reflect.TypeOf((*q.Addr)(nil)).Elem(),
				"AddrPort": reflect.TypeOf((*q.AddrPort)(nil)).Elem(),
				"Prefix":   reflect.TypeOf((*q.Prefix)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"AddrFrom16":              q.AddrFrom16,
				"AddrFrom4":               q.AddrFrom4,
				"AddrFromSlice":           q.AddrFromSlice,
				"AddrPortFrom":            q.AddrPortFrom,
				"IPv4Unspecified":         q.IPv4Unspecified,
				"IPv6LinkLocalAllNodes":   q.IPv6LinkLocalAllNodes,
				"IPv6LinkLocalAllRouters": q.IPv6LinkLocalAllRouters,
				"IPv6Loopback":            q.IPv6Loopback,
				"IPv6Unspecified":         q.IPv6Unspecified,
				"MustParseAddr":           q.MustParseAddr,
				"MustParseAddrPort":       q.MustParseAddrPort,
				"MustParsePrefix":         q.MustParsePrefix,
				"ParseAddr":               q.ParseAddr,
				"ParseAddrPort":           q.ParseAddrPort,
				"ParsePrefix":             q.ParsePrefix,
				"PrefixFrom":              q.PrefixFrom,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
