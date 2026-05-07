// export by github.com/goplus/ixgo/cmd/qexp

package parser

import (
	q "github.com/goplus/xgo/tpl/parser"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/parser", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "parser",
			Path: "github.com/goplus/xgo/tpl/parser",
			Deps: map[string]string{
				"github.com/goplus/xgo/tpl/ast":     "ast",
				"github.com/goplus/xgo/tpl/scanner": "scanner",
				"github.com/goplus/xgo/tpl/token":   "token",
				"github.com/qiniu/x/stream":         "stream",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Config": reflect.TypeOf((*q.Config)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"RetProcParser": reflect.TypeOf((*q.RetProcParser)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"ParseEx":   q.ParseEx,
				"ParseFile": q.ParseFile,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
