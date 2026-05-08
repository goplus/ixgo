// export by github.com/goplus/ixgo/cmd/qexp

package cl

import (
	q "github.com/goplus/xgo/tpl/cl"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("github.com/goplus/xgo/tpl/cl", func() *ixgo.Package {
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
			Vars: map[string]reflect.Value{
				"ErrNoDocFound": reflect.ValueOf(&q.ErrNoDocFound),
			},
			Funcs: map[string]reflect.Value{
				"LogConflict": reflect.ValueOf(q.LogConflict),
				"New":         reflect.ValueOf(q.New),
				"NewEx":       reflect.ValueOf(q.NewEx),
			},
			TypedConsts:   map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
