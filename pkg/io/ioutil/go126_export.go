// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package ioutil

import (
	q "io/ioutil"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("io/ioutil", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "ioutil",
			Path: "io/ioutil",
			Deps: map[string]string{
				"io":      "io",
				"io/fs":   "fs",
				"os":      "os",
				"slices":  "slices",
				"strings": "strings",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"Discard": &q.Discard,
			},
			Funcs: map[string]interface{}{
				"NopCloser": q.NopCloser,
				"ReadAll":   q.ReadAll,
				"ReadDir":   q.ReadDir,
				"ReadFile":  q.ReadFile,
				"TempDir":   q.TempDir,
				"TempFile":  q.TempFile,
				"WriteFile": q.WriteFile,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
