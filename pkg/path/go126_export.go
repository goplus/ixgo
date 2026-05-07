// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package path

import (
	q "path"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("path", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "path",
			Path: "path",
			Deps: map[string]string{
				"errors":           "errors",
				"internal/bytealg": "bytealg",
				"unicode/utf8":     "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrBadPattern": &q.ErrBadPattern,
			},
			Funcs: map[string]interface{}{
				"Base":  q.Base,
				"Clean": q.Clean,
				"Dir":   q.Dir,
				"Ext":   q.Ext,
				"IsAbs": q.IsAbs,
				"Join":  q.Join,
				"Match": q.Match,
				"Split": q.Split,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
