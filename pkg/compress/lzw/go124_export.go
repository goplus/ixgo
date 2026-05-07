// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package lzw

import (
	q "compress/lzw"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("compress/lzw", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "lzw",
			Path: "compress/lzw",
			Deps: map[string]string{
				"bufio":  "bufio",
				"errors": "errors",
				"fmt":    "fmt",
				"io":     "io",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Order":  reflect.TypeOf((*q.Order)(nil)).Elem(),
				"Reader": reflect.TypeOf((*q.Reader)(nil)).Elem(),
				"Writer": reflect.TypeOf((*q.Writer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewReader": q.NewReader,
				"NewWriter": q.NewWriter,
			},
			TypedConsts: map[string]interface{}{
				"LSB": q.LSB,
				"MSB": q.MSB,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
