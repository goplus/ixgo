// export by github.com/goplus/ixgo/cmd/qexp

package test

import (
	q "github.com/goplus/xgo/test"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Gopt_App_TestMain":  reflect.ValueOf(q.Gopt_App_TestMain),
			"Gopt_Case_TestMain": reflect.ValueOf(q.Gopt_Case_TestMain),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"GopPackage": {Typ: "untyped bool", Value: constant.MakeBool(bool(q.GopPackage))},
		},
	})
}
