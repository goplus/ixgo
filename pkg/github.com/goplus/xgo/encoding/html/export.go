// export by github.com/goplus/ixgo/cmd/qexp

package html

import (
	q "github.com/goplus/xgo/encoding/html"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "html",
		Path: "github.com/goplus/xgo/encoding/html",
		Deps: map[string]string{
			"github.com/goplus/xgo/dql/html": "html",
			"strings":                        "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"Object": reflect.TypeOf((*q.Object)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"XGoPackage": {"untyped string", constant.MakeString(string(q.XGoPackage))},
		},
	})
}
