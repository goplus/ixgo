// export by github.com/goplus/ixgo/cmd/qexp

package cl

import (
	q "github.com/goplus/xgo/tpl/cl"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/cl", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "cl",
			Path: "github.com/goplus/xgo/tpl/cl",
			Deps: map[string]string{
				"fmt":                               "fmt",
				"github.com/goplus/xgo/tpl/ast":     "ast",
				"github.com/goplus/xgo/tpl/matcher": "matcher",
				"github.com/goplus/xgo/tpl/token":   "token",
				"github.com/qiniu/x/errors":         "errors",
				"os":                                "os",
				"strconv":                           "strconv",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Config": reflect.TypeOf((*q.Config)(nil)).Elem(),
				"Result": reflect.TypeOf((*q.Result)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrNoDocFound": &q.ErrNoDocFound,
			},
			Funcs: map[string]interface{}{
				"LogConflict": q.LogConflict,
				"New":         q.New,
				"NewEx":       q.NewEx,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
