// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package httputil

import (
	q "net/http/httputil"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("net/http/httputil", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "httputil",
			Path: "net/http/httputil",
			Deps: map[string]string{
				"bufio":                                 "bufio",
				"bytes":                                 "bytes",
				"context":                               "context",
				"errors":                                "errors",
				"fmt":                                   "fmt",
				"io":                                    "io",
				"log":                                   "log",
				"mime":                                  "mime",
				"net":                                   "net",
				"net/http":                              "http",
				"net/http/httptrace":                    "httptrace",
				"net/http/internal":                     "internal",
				"net/http/internal/ascii":               "ascii",
				"net/textproto":                         "textproto",
				"net/url":                               "url",
				"strings":                               "strings",
				"sync":                                  "sync",
				"sync/atomic":                           "atomic",
				"time":                                  "time",
				"vendor/golang.org/x/net/http/httpguts": "httpguts",
			},
			Interfaces: map[string]reflect.Type{
				"BufferPool": reflect.TypeOf((*q.BufferPool)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"ClientConn":   reflect.TypeOf((*q.ClientConn)(nil)).Elem(),
				"ProxyRequest": reflect.TypeOf((*q.ProxyRequest)(nil)).Elem(),
				"ReverseProxy": reflect.TypeOf((*q.ReverseProxy)(nil)).Elem(),
				"ServerConn":   reflect.TypeOf((*q.ServerConn)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrClosed":      &q.ErrClosed,
				"ErrLineTooLong": &q.ErrLineTooLong,
				"ErrPersistEOF":  &q.ErrPersistEOF,
				"ErrPipeline":    &q.ErrPipeline,
			},
			Funcs: map[string]interface{}{
				"DumpRequest":               q.DumpRequest,
				"DumpRequestOut":            q.DumpRequestOut,
				"DumpResponse":              q.DumpResponse,
				"NewChunkedReader":          q.NewChunkedReader,
				"NewChunkedWriter":          q.NewChunkedWriter,
				"NewClientConn":             q.NewClientConn,
				"NewProxyClientConn":        q.NewProxyClientConn,
				"NewServerConn":             q.NewServerConn,
				"NewSingleHostReverseProxy": q.NewSingleHostReverseProxy,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
