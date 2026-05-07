// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package format

import (
	q "go/format"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("go/format", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "format",
			Path: "go/format",
			Deps: map[string]string{
				"bytes":      "bytes",
				"fmt":        "fmt",
				"go/ast":     "ast",
				"go/parser":  "parser",
				"go/printer": "printer",
				"go/token":   "token",
				"io":         "io",
				"strings":    "strings",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Node":   q.Node,
				"Source": q.Source,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
