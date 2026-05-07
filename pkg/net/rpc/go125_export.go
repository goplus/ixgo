// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package rpc

import (
	q "net/rpc"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("net/rpc", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "rpc",
			Path: "net/rpc",
			Deps: map[string]string{
				"bufio":         "bufio",
				"encoding/gob":  "gob",
				"errors":        "errors",
				"fmt":           "fmt",
				"go/token":      "token",
				"html/template": "template",
				"io":            "io",
				"log":           "log",
				"net":           "net",
				"net/http":      "http",
				"reflect":       "reflect",
				"slices":        "slices",
				"strings":       "strings",
				"sync":          "sync",
			},
			Interfaces: map[string]reflect.Type{
				"ClientCodec": reflect.TypeOf((*q.ClientCodec)(nil)).Elem(),
				"ServerCodec": reflect.TypeOf((*q.ServerCodec)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Call":        reflect.TypeOf((*q.Call)(nil)).Elem(),
				"Client":      reflect.TypeOf((*q.Client)(nil)).Elem(),
				"Request":     reflect.TypeOf((*q.Request)(nil)).Elem(),
				"Response":    reflect.TypeOf((*q.Response)(nil)).Elem(),
				"Server":      reflect.TypeOf((*q.Server)(nil)).Elem(),
				"ServerError": reflect.TypeOf((*q.ServerError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"DefaultServer": &q.DefaultServer,
				"ErrShutdown":   &q.ErrShutdown,
			},
			Funcs: map[string]interface{}{
				"Accept":             q.Accept,
				"Dial":               q.Dial,
				"DialHTTP":           q.DialHTTP,
				"DialHTTPPath":       q.DialHTTPPath,
				"HandleHTTP":         q.HandleHTTP,
				"NewClient":          q.NewClient,
				"NewClientWithCodec": q.NewClientWithCodec,
				"NewServer":          q.NewServer,
				"Register":           q.Register,
				"RegisterName":       q.RegisterName,
				"ServeCodec":         q.ServeCodec,
				"ServeConn":          q.ServeConn,
				"ServeRequest":       q.ServeRequest,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"DefaultDebugPath": {Typ: "untyped string", Value: constant.MakeString(string(q.DefaultDebugPath))},
				"DefaultRPCPath":   {Typ: "untyped string", Value: constant.MakeString(string(q.DefaultRPCPath))},
			},
		}
	})
}
