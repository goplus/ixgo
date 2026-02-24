// export by github.com/goplus/ixgo/cmd/qexp

package golang

import (
	q "github.com/goplus/xgo/dql/golang"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"From":         reflect.ValueOf(q.From),
			"New":          reflect.ValueOf(q.New),
			"NodeSet_Cast": reflect.ValueOf(q.NodeSet_Cast),
			"Nodes":        reflect.ValueOf(q.Nodes),
			"ParseFile":    reflect.ValueOf(q.ParseFile),
			"Root":         reflect.ValueOf(q.Root),
			"Source":       reflect.ValueOf(q.Source),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"XGoPackage": {"untyped string", constant.MakeString(string(q.XGoPackage))},
		},
	})
}
