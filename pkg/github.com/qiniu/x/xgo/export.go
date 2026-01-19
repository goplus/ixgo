// export by github.com/goplus/ixgo/cmd/qexp

package xgo

import (
	q "github.com/qiniu/x/xgo"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name:       "xgo",
		Path:       "github.com/qiniu/x/xgo",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"IntRange": reflect.TypeOf((*q.IntRange)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewRange__0": reflect.ValueOf(q.NewRange__0),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"XGoPackage": {"untyped bool", constant.MakeBool(bool(q.XGoPackage))},
		},
	})
}
