// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package httputil

import (
	q "net/http/httputil"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"ErrClosed":      reflect.ValueOf(&q.ErrClosed),
			"ErrLineTooLong": reflect.ValueOf(&q.ErrLineTooLong),
			"ErrPersistEOF":  reflect.ValueOf(&q.ErrPersistEOF),
			"ErrPipeline":    reflect.ValueOf(&q.ErrPipeline),
		},
		Funcs: map[string]reflect.Value{
			"DumpRequest":               reflect.ValueOf(q.DumpRequest),
			"DumpRequestOut":            reflect.ValueOf(q.DumpRequestOut),
			"DumpResponse":              reflect.ValueOf(q.DumpResponse),
			"NewChunkedReader":          reflect.ValueOf(q.NewChunkedReader),
			"NewChunkedWriter":          reflect.ValueOf(q.NewChunkedWriter),
			"NewClientConn":             reflect.ValueOf(q.NewClientConn),
			"NewProxyClientConn":        reflect.ValueOf(q.NewProxyClientConn),
			"NewServerConn":             reflect.ValueOf(q.NewServerConn),
			"NewSingleHostReverseProxy": reflect.ValueOf(q.NewSingleHostReverseProxy),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
