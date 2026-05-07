// export by github.com/goplus/ixgo/cmd/qexp

package xml

import (
	q "github.com/goplus/xgo/encoding/xml"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/encoding/xml", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "xml",
			Path: "github.com/goplus/xgo/encoding/xml",
			Deps: map[string]string{
				"github.com/goplus/xgo/dql/xml": "xml",
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
