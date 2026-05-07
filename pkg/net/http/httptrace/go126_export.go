// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package httptrace

import (
	q "net/http/httptrace"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("net/http/httptrace", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "httptrace",
			Path: "net/http/httptrace",
			Deps: map[string]string{
				"context":           "context",
				"crypto/tls":        "tls",
				"internal/nettrace": "nettrace",
				"net":               "net",
				"net/textproto":     "textproto",
				"reflect":           "reflect",
				"time":              "time",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"ClientTrace":      reflect.TypeOf((*q.ClientTrace)(nil)).Elem(),
				"DNSDoneInfo":      reflect.TypeOf((*q.DNSDoneInfo)(nil)).Elem(),
				"DNSStartInfo":     reflect.TypeOf((*q.DNSStartInfo)(nil)).Elem(),
				"GotConnInfo":      reflect.TypeOf((*q.GotConnInfo)(nil)).Elem(),
				"WroteRequestInfo": reflect.TypeOf((*q.WroteRequestInfo)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"ContextClientTrace": q.ContextClientTrace,
				"WithClientTrace":    q.WithClientTrace,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
