// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package list

import (
	q "container/list"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("container/list", func() *ixgo.Package {
		return &ixgo.Package{
			Name:       "list",
			Path:       "container/list",
			Deps:       map[string]string{},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Element": reflect.TypeOf((*q.Element)(nil)).Elem(),
				"List":    reflect.TypeOf((*q.List)(nil)).Elem(),
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
