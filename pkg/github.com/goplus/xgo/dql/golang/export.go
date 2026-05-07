// export by github.com/goplus/ixgo/cmd/qexp

package golang

import (
	q "github.com/goplus/xgo/dql/golang"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/dql/golang", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "golang",
			Path: "github.com/goplus/xgo/dql/golang",
			Deps: map[string]string{
				"bytes":                              "bytes",
				"github.com/goplus/xgo/dql":          "dql",
				"github.com/goplus/xgo/dql/reflects": "reflects",
				"github.com/qiniu/x/stream":          "stream",
				"go/ast":                             "ast",
				"go/parser":                          "parser",
				"go/token":                           "token",
				"io":                                 "io",
				"iter":                               "iter",
				"reflect":                            "reflect",
				"unsafe":                             "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Config":  reflect.TypeOf((*q.Config)(nil)).Elem(),
				"File":    reflect.TypeOf((*q.File)(nil)).Elem(),
				"NodeSet": reflect.TypeOf((*q.NodeSet)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"Node": reflect.TypeOf((*q.Node)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"From":         q.From,
				"New":          q.New,
				"NodeSet_Cast": q.NodeSet_Cast,
				"Nodes":        q.Nodes,
				"ParseFile":    q.ParseFile,
				"Root":         q.Root,
				"Source":       q.Source,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"XGoPackage": {Typ: "untyped string", Value: constant.MakeString(string(q.XGoPackage))},
			},
		}
	})
}
