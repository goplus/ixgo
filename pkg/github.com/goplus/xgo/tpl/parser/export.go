// export by github.com/goplus/ixgo/cmd/qexp

package parser

import (
	q "github.com/goplus/xgo/tpl/parser"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ParseEx":   reflect.ValueOf(q.ParseEx),
			"ParseFile": reflect.ValueOf(q.ParseFile),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
