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
	ixgo.RegisterPackage("text/tabwriter", func() *ixgo.Package {
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewWriter": q.NewWriter,
			},
			TypedConsts: map[string]interface{}{
				"AlignRight":          q.AlignRight,
				"Debug":               q.Debug,
				"DiscardEmptyColumns": q.DiscardEmptyColumns,
				"FilterHTML":          q.FilterHTML,
				"StripEscape":         q.StripEscape,
				"TabIndent":           q.TabIndent,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"Escape": {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.Escape))},
			},
		}
	})
}
