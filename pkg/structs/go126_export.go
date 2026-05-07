// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package structs

import (
	q "structs"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("structs", func() *ixgo.Package {
		return &ixgo.Package{
			Name:       "structs",
			Path:       "structs",
			Deps:       map[string]string{},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"HostLayout": reflect.TypeOf((*q.HostLayout)(nil)).Elem(),
			},
			AliasTypes:    map[string]reflect.Type{},
			Vars:          map[string]interface{}{},
			Funcs:         map[string]interface{}{},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
