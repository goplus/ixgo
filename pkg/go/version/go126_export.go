// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package version

import (
	q "go/version"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("go/version", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "version",
			Path: "go/version",
			Deps: map[string]string{
				"internal/gover": "gover",
				"strings":        "strings",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Compare": q.Compare,
				"IsValid": q.IsValid,
				"Lang":    q.Lang,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
