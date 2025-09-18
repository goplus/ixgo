// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package jsonrpc

import (
	q "net/rpc/jsonrpc"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Dial":           reflect.ValueOf(q.Dial),
			"NewClient":      reflect.ValueOf(q.NewClient),
			"NewClientCodec": reflect.ValueOf(q.NewClientCodec),
			"NewServerCodec": reflect.ValueOf(q.NewServerCodec),
			"ServeConn":      reflect.ValueOf(q.ServeConn),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
