// export by github.com/goplus/ixgo/cmd/qexp

package xml

import (
	q "github.com/goplus/xgo/dql/xml"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/dql/xml", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "xml",
			Path: "github.com/goplus/xgo/dql/xml",
			Deps: map[string]string{
				"bytes":                     "bytes",
				"encoding/xml":              "xml",
				"fmt":                       "fmt",
				"github.com/goplus/xgo/dql": "dql",
				"github.com/qiniu/x/stream": "stream",
				"io":                        "io",
				"iter":                      "iter",
				"os":                        "os",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Node":    reflect.TypeOf((*q.Node)(nil)).Elem(),
				"NodeSet": reflect.TypeOf((*q.NodeSet)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"Attr":     reflect.TypeOf((*q.Attr)(nil)).Elem(),
				"CharData": reflect.TypeOf((*q.CharData)(nil)).Elem(),
				"Name":     reflect.TypeOf((*q.Name)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New":          q.New,
				"NodeSet_Cast": q.NodeSet_Cast,
				"Nodes":        q.Nodes,
				"Parse":        q.Parse,
				"Root":         q.Root,
				"Source":       q.Source,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"XGoPackage": {Typ: "untyped bool", Value: constant.MakeBool(bool(q.XGoPackage))},
			},
		}
	})
}
