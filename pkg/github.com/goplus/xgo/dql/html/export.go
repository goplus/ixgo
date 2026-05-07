// export by github.com/goplus/ixgo/cmd/qexp

package html

import (
	q "github.com/goplus/xgo/dql/html"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/dql/html", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "html",
			Path: "github.com/goplus/xgo/dql/html",
			Deps: map[string]string{
				"bytes":                      "bytes",
				"fmt":                        "fmt",
				"github.com/goplus/xgo/dql":  "dql",
				"github.com/qiniu/x/stream":  "stream",
				"golang.org/x/net/html":      "html",
				"golang.org/x/net/html/atom": "atom",
				"io":                         "io",
				"iter":                       "iter",
				"os":                         "os",
				"unsafe":                     "unsafe",
			},
			Interfaces: map[string]reflect.Type{
				"NodeFilter": reflect.TypeOf((*q.NodeFilter)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Node":    reflect.TypeOf((*q.Node)(nil)).Elem(),
				"NodeSet": reflect.TypeOf((*q.NodeSet)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"ClassContains": q.ClassContains,
				"New":           q.New,
				"NodeSet_Cast":  q.NodeSet_Cast,
				"Nodes":         q.Nodes,
				"Parse":         q.Parse,
				"Root":          q.Root,
				"Source":        q.Source,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"XGoPackage": {Typ: "untyped bool", Value: constant.MakeBool(bool(q.XGoPackage))},
			},
		}
	})
}
