// export by github.com/goplus/ixgo/cmd/qexp

package html

import (
	q "github.com/goplus/xgo/dql/html"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
			"strings":                    "strings",
			"unsafe":                     "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"File":    reflect.TypeOf((*q.File)(nil)).Elem(),
			"NodeSet": reflect.TypeOf((*q.NodeSet)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Node": reflect.TypeOf((*q.Node)(nil)).Elem(),
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
