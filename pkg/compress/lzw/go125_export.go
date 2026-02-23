// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package lzw

import (
	q "compress/lzw"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(q.NewReader),
			"NewWriter": reflect.ValueOf(q.NewWriter),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"LSB": {reflect.TypeOf(q.LSB), constant.MakeInt64(int64(q.LSB))},
			"MSB": {reflect.TypeOf(q.MSB), constant.MakeInt64(int64(q.MSB))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
