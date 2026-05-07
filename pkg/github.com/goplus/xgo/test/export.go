// export by github.com/goplus/ixgo/cmd/qexp

package test

import (
	q "github.com/goplus/xgo/test"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/test", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "test",
			Path: "github.com/goplus/xgo/test",
			Deps: map[string]string{
				"os":      "os",
				"testing": "testing",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"App":  reflect.TypeOf((*q.App)(nil)).Elem(),
				"Case": reflect.TypeOf((*q.Case)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Gopt_App_TestMain":  q.Gopt_App_TestMain,
				"Gopt_Case_TestMain": q.Gopt_Case_TestMain,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"GopPackage": {Typ: "untyped bool", Value: constant.MakeBool(bool(q.GopPackage))},
			},
		}
	})
}
