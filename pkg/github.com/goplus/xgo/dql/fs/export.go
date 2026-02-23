// export by github.com/goplus/ixgo/cmd/qexp

package fs

import (
	q "github.com/goplus/xgo/dql/fs"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "fs",
		Path: "github.com/goplus/xgo/dql/fs",
		Deps: map[string]string{
			"errors":                    "errors",
			"github.com/goplus/xgo/dql": "dql",
			"io/fs":                     "fs",
			"iter":                      "iter",
			"os":                        "os",
			"path":                      "path",
			"time":                      "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Node":    reflect.TypeOf((*q.Node)(nil)).Elem(),
			"NodeSet": reflect.TypeOf((*q.NodeSet)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Dir":   reflect.ValueOf(q.Dir),
			"New":   reflect.ValueOf(q.New),
			"Nodes": reflect.ValueOf(q.Nodes),
			"Root":  reflect.ValueOf(q.Root),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
