// export by github.com/goplus/ixgo/cmd/qexp

package format

import (
	q "github.com/goplus/xgo/format"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "format",
		Path: "github.com/goplus/xgo/format",
		Deps: map[string]string{
			"bytes":                         "bytes",
			"github.com/goplus/xgo/ast":     "ast",
			"github.com/goplus/xgo/parser":  "parser",
			"github.com/goplus/xgo/printer": "printer",
			"github.com/goplus/xgo/token":   "token",
			"io":                            "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Node":   reflect.ValueOf(q.Node),
			"Source": reflect.ValueOf(q.Source),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
