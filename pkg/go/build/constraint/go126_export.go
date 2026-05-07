// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package constraint

import (
	q "go/build/constraint"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("go/build/constraint", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "constraint",
			Path: "go/build/constraint",
			Deps: map[string]string{
				"errors":       "errors",
				"strconv":      "strconv",
				"strings":      "strings",
				"unicode":      "unicode",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{
				"Expr": reflect.TypeOf((*q.Expr)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"AndExpr":     reflect.TypeOf((*q.AndExpr)(nil)).Elem(),
				"NotExpr":     reflect.TypeOf((*q.NotExpr)(nil)).Elem(),
				"OrExpr":      reflect.TypeOf((*q.OrExpr)(nil)).Elem(),
				"SyntaxError": reflect.TypeOf((*q.SyntaxError)(nil)).Elem(),
				"TagExpr":     reflect.TypeOf((*q.TagExpr)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"GoVersion":      q.GoVersion,
				"IsGoBuild":      q.IsGoBuild,
				"IsPlusBuild":    q.IsPlusBuild,
				"Parse":          q.Parse,
				"PlusBuildLines": q.PlusBuildLines,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
