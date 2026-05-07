// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package bzip2

import (
	q "compress/bzip2"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("compress/bzip2", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "bzip2",
			Path: "compress/bzip2",
			Deps: map[string]string{
				"bufio":  "bufio",
				"cmp":    "cmp",
				"io":     "io",
				"slices": "slices",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"StructuralError": reflect.TypeOf((*q.StructuralError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewReader": q.NewReader,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
