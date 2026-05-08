// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package tabwriter

import (
	q "text/tabwriter"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("text/tabwriter", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "tabwriter",
			Path: "text/tabwriter",
			Deps: map[string]string{
				"fmt":          "fmt",
				"io":           "io",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Writer": reflect.TypeOf((*q.Writer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"NewWriter": reflect.ValueOf(q.NewWriter),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"AlignRight":          {Typ: reflect.TypeOf(q.AlignRight), Value: constant.MakeInt64(int64(q.AlignRight))},
				"Debug":               {Typ: reflect.TypeOf(q.Debug), Value: constant.MakeInt64(int64(q.Debug))},
				"DiscardEmptyColumns": {Typ: reflect.TypeOf(q.DiscardEmptyColumns), Value: constant.MakeInt64(int64(q.DiscardEmptyColumns))},
				"FilterHTML":          {Typ: reflect.TypeOf(q.FilterHTML), Value: constant.MakeInt64(int64(q.FilterHTML))},
				"StripEscape":         {Typ: reflect.TypeOf(q.StripEscape), Value: constant.MakeInt64(int64(q.StripEscape))},
				"TabIndent":           {Typ: reflect.TypeOf(q.TabIndent), Value: constant.MakeInt64(int64(q.TabIndent))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"Escape": {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.Escape))},
			},
		}
	})
}
