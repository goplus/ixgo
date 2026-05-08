// export by github.com/goplus/ixgo/cmd/qexp

package parsertest

import (
	q "github.com/goplus/xgo/tpl/parser/parsertest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("github.com/goplus/xgo/tpl/parser/parsertest", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "parsertest",
			Path: "github.com/goplus/xgo/tpl/parser/parsertest",
			Deps: map[string]string{
				"bytes":                           "bytes",
				"fmt":                             "fmt",
				"github.com/goplus/xgo/tpl/ast":   "ast",
				"github.com/goplus/xgo/tpl/token": "token",
				"github.com/qiniu/x/test":         "test",
				"io":                              "io",
				"log":                             "log",
				"reflect":                         "reflect",
				"testing":                         "testing",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"Expect":     reflect.ValueOf(q.Expect),
				"Fprint":     reflect.ValueOf(q.Fprint),
				"FprintNode": reflect.ValueOf(q.FprintNode),
			},
			TypedConsts:   map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
