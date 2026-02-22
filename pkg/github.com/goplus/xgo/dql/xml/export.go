// export by github.com/goplus/ixgo/cmd/qexp

package xml

import (
	q "github.com/goplus/xgo/dql/xml"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":          reflect.ValueOf(q.New),
			"NodeSet_Cast": reflect.ValueOf(q.NodeSet_Cast),
			"Nodes":        reflect.ValueOf(q.Nodes),
			"Parse":        reflect.ValueOf(q.Parse),
			"Root":         reflect.ValueOf(q.Root),
			"Source":       reflect.ValueOf(q.Source),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"XGoPackage": {"untyped bool", constant.MakeBool(bool(q.XGoPackage))},
		},
	})
}
