// export by github.com/goplus/ixgo/cmd/qexp

package reflects

import (
	q "github.com/goplus/xgo/dql/reflects"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/dql/reflects", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "reflects",
			Path: "github.com/goplus/xgo/dql/reflects",
			Deps: map[string]string{
				"github.com/goplus/xgo/dql": "dql",
				"iter":                      "iter",
				"reflect":                   "reflect",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Node":    reflect.TypeOf((*q.Node)(nil)).Elem(),
				"NodeSet": reflect.TypeOf((*q.NodeSet)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New":          q.New,
				"NodeSet_Cast": q.NodeSet_Cast,
				"Nodes":        q.Nodes,
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
