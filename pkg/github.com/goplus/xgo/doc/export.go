// export by github.com/goplus/ixgo/cmd/qexp

package doc

import (
	q "github.com/goplus/xgo/doc"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "doc",
		Path: "github.com/goplus/xgo/doc",
		Deps: map[string]string{
			"go/ast":   "ast",
			"go/doc":   "doc",
			"go/token": "token",
			"sort":     "sort",
			"strconv":  "strconv",
			"strings":  "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Transform": reflect.ValueOf(q.Transform),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
