// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package html

import (
	q "html"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("html", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "html",
			Path: "html",
			Deps: map[string]string{
				"strings":      "strings",
				"sync":         "sync",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"EscapeString":   q.EscapeString,
				"UnescapeString": q.UnescapeString,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
