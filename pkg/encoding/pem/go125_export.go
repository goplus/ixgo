// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package pem

import (
	q "encoding/pem"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "pem",
		Path: "encoding/pem",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding/base64": "base64",
			"errors":          "errors",
			"io":              "io",
			"slices":          "slices",
			"strings":         "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Block": reflect.TypeOf((*q.Block)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":         reflect.ValueOf(q.Decode),
			"Encode":         reflect.ValueOf(q.Encode),
			"EncodeToMemory": reflect.ValueOf(q.EncodeToMemory),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
