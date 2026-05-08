// export by github.com/goplus/ixgo/cmd/qexp

package buitin

import (
	q "github.com/goplus/xgo/tpl/variant/builtin"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("github.com/goplus/xgo/tpl/variant/builtin", func() *ixgo.Package {
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
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"CastInt": reflect.ValueOf(q.CastInt),
				"Type":    reflect.ValueOf(q.Type),
			},
			TypedConsts:   map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
