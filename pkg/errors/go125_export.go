// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package errors

import (
	q "errors"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("errors", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "errors",
			Path: "errors",
			Deps: map[string]string{
				"internal/reflectlite": "reflectlite",
				"unsafe":               "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrUnsupported": &q.ErrUnsupported,
			},
			Funcs: map[string]interface{}{
				"As":     q.As,
				"Is":     q.Is,
				"Join":   q.Join,
				"New":    q.New,
				"Unwrap": q.Unwrap,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
