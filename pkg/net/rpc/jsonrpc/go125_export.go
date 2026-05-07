// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package jsonrpc

import (
	q "net/rpc/jsonrpc"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("net/rpc/jsonrpc", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "jsonrpc",
			Path: "net/rpc/jsonrpc",
			Deps: map[string]string{
				"encoding/json": "json",
				"errors":        "errors",
				"fmt":           "fmt",
				"io":            "io",
				"net":           "net",
				"net/rpc":       "rpc",
				"sync":          "sync",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Dial":           q.Dial,
				"NewClient":      q.NewClient,
				"NewClientCodec": q.NewClientCodec,
				"NewServerCodec": q.NewServerCodec,
				"ServeConn":      q.ServeConn,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
