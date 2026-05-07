// export by github.com/goplus/ixgo/cmd/qexp

package json

import (
	q "github.com/goplus/xgo/dql/json"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/dql/json", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "json",
			Path: "github.com/goplus/xgo/dql/json",
			Deps: map[string]string{
				"bytes":                          "bytes",
				"encoding/json":                  "json",
				"github.com/goplus/xgo/dql/maps": "maps",
				"github.com/qiniu/x/stream":      "stream",
				"io":                             "io",
				"iter":                           "iter",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{
				"Node":    reflect.TypeOf((*q.Node)(nil)).Elem(),
				"NodeSet": reflect.TypeOf((*q.NodeSet)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New":    q.New,
				"Source": q.Source,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"XGoPackage": {Typ: "untyped string", Value: constant.MakeString(string(q.XGoPackage))},
			},
		}
	})
}
