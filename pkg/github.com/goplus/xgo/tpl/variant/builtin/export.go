// export by github.com/goplus/ixgo/cmd/qexp

package buitin

import (
	q "github.com/goplus/xgo/tpl/variant/builtin"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/variant/builtin", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "buitin",
			Path: "github.com/goplus/xgo/tpl/variant/builtin",
			Deps: map[string]string{
				"github.com/goplus/xgo/tpl/variant": "variant",
				"reflect":                           "reflect",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"CastInt": q.CastInt,
				"Type":    q.Type,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
