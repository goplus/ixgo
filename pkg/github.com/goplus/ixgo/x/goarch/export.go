// export by github.com/goplus/ixgo/cmd/qexp

package goarch

import (
	q "github.com/goplus/ixgo/x/goarch"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name:        "goarch",
		Path:        "github.com/goplus/ixgo/x/goarch",
		Deps:        map[string]string{},
		Interfaces:  map[string]reflect.Type{},
		NamedTypes:  map[string]reflect.Type{},
		AliasTypes:  map[string]reflect.Type{},
		Vars:        map[string]reflect.Value{},
		Funcs:       map[string]reflect.Value{},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"PtrSize": {"untyped int", constant.MakeInt64(int64(q.PtrSize))},
		},
	})
}
