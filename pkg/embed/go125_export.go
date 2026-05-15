// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package embed

import (
	q "embed"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "embed",
		Path: "embed",
		Deps: map[string]string{
			"errors":               "errors",
			"internal/bytealg":     "bytealg",
			"internal/stringslite": "stringslite",
			"io":                   "io",
			"io/fs":                "fs",
			"time":                 "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"FS": reflect.TypeOf((*q.FS)(nil)).Elem(),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
