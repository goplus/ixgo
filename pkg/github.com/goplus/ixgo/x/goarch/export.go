// export by github.com/goplus/ixgo/cmd/qexp

package goarch

import (
	q "github.com/goplus/ixgo/x/goarch"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/ixgo/x/goarch", func() *ixgo.Package {
		return &ixgo.Package{
			Name:        "goarch",
			Path:        "github.com/goplus/ixgo/x/goarch",
			Deps:        map[string]string{},
			Interfaces:  map[string]reflect.Type{},
			NamedTypes:  map[string]reflect.Type{},
			AliasTypes:  map[string]reflect.Type{},
			Vars:        map[string]interface{}{},
			Funcs:       map[string]interface{}{},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"PtrSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.PtrSize))},
			},
		}
	})
}
