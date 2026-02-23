// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package format

import (
	q "go/format"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Node":   reflect.ValueOf(q.Node),
			"Source": reflect.ValueOf(q.Source),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
