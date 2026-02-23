// export by github.com/goplus/ixgo/cmd/qexp

package json

import (
	q "github.com/goplus/xgo/dql/json"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":    reflect.ValueOf(q.New),
			"Source": reflect.ValueOf(q.Source),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"XGoPackage": {"untyped string", constant.MakeString(string(q.XGoPackage))},
		},
	})
}
