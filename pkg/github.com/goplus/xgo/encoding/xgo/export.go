// export by github.com/goplus/ixgo/cmd/qexp

package xgo

import (
	q "github.com/goplus/xgo/encoding/xgo"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/encoding/xgo", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "xgo",
			Path: "github.com/goplus/xgo/encoding/xgo",
			Deps: map[string]string{
				"github.com/goplus/xgo/dql/xgo": "xgo",
				"github.com/goplus/xgo/parser":  "parser",
				"strings":                       "strings",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{
				"Object": reflect.TypeOf((*q.Object)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New": q.New,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"XGoPackage": {Typ: "untyped string", Value: constant.MakeString(string(q.XGoPackage))},
			},
		}
	})
}
