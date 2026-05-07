// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package suffixarray

import (
	q "index/suffixarray"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("index/suffixarray", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "suffixarray",
			Path: "index/suffixarray",
			Deps: map[string]string{
				"bytes":           "bytes",
				"encoding/binary": "binary",
				"errors":          "errors",
				"io":              "io",
				"math":            "math",
				"regexp":          "regexp",
				"slices":          "slices",
				"sort":            "sort",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Index": reflect.TypeOf((*q.Index)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New": q.New,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
