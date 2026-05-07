// export by github.com/goplus/ixgo/cmd/qexp

package types

import (
	q "github.com/goplus/xgo/tpl/types"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/types", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "types",
			Path: "github.com/goplus/xgo/tpl/types",
			Deps: map[string]string{
				"github.com/goplus/xgo/tpl/token": "token",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Token": reflect.TypeOf((*q.Token)(nil)).Elem(),
			},
			AliasTypes:    map[string]reflect.Type{},
			Vars:          map[string]interface{}{},
			Funcs:         map[string]interface{}{},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
