// export by github.com/goplus/ixgo/cmd/qexp

package maps

import (
	q "github.com/goplus/xgo/dql/maps"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "maps",
		Path: "github.com/goplus/xgo/dql/maps",
		Deps: map[string]string{
			"github.com/goplus/xgo/dql": "dql",
			"iter":                      "iter",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Node":    reflect.TypeOf((*q.Node)(nil)).Elem(),
			"NodeSet": reflect.TypeOf((*q.NodeSet)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":          reflect.ValueOf(q.New),
			"NodeSet_Cast": reflect.ValueOf(q.NodeSet_Cast),
			"Nodes":        reflect.ValueOf(q.Nodes),
			"Root":         reflect.ValueOf(q.Root),
			"Source":       reflect.ValueOf(q.Source),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"XGoPackage": {"untyped bool", constant.MakeBool(bool(q.XGoPackage))},
		},
	})
}
