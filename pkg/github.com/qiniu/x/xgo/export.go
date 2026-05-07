// export by github.com/goplus/ixgo/cmd/qexp

package xgo

import (
	q "github.com/qiniu/x/xgo"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/qiniu/x/xgo", func() *ixgo.Package {
		return &ixgo.Package{
			Name:       "xgo",
			Path:       "github.com/qiniu/x/xgo",
			Deps:       map[string]string{},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"IntRange": reflect.TypeOf((*q.IntRange)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewRange__0": q.NewRange__0,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"XGoPackage": {Typ: "untyped bool", Value: constant.MakeBool(bool(q.XGoPackage))},
			},
		}
	})
}
