// export by github.com/goplus/ixgo/cmd/qexp

package parsertest

import (
	q "github.com/goplus/xgo/tpl/parser/parsertest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/parser/parsertest", func() *ixgo.Package {
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Expect":     q.Expect,
				"Fprint":     q.Fprint,
				"FprintNode": q.FprintNode,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
