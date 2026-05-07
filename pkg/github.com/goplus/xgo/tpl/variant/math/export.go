// export by github.com/goplus/ixgo/cmd/qexp

package math

import (
	q "github.com/goplus/xgo/tpl/variant/math"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/variant/math", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "math",
			Path: "github.com/goplus/xgo/tpl/variant/math",
			Deps: map[string]string{
				"github.com/goplus/xgo/tpl/token":   "token",
				"github.com/goplus/xgo/tpl/variant": "variant",
				"math":                              "math",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Max": q.Max,
				"Min": q.Min,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
