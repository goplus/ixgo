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
	ixgo.RegisterPackage("encoding/pem", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Decode":         q.Decode,
				"Encode":         q.Encode,
				"EncodeToMemory": q.EncodeToMemory,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
