// export by github.com/goplus/ixgo/cmd/qexp

package fs

import (
	q "github.com/goplus/xgo/dql/fs"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/dql/fs", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Dir":   q.Dir,
				"New":   q.New,
				"Nodes": q.Nodes,
				"Root":  q.Root,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
