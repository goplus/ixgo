// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package parser

import (
	q "go/parser"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("go/parser", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "parser",
			Path: "go/parser",
			Deps: map[string]string{
				"bytes":               "bytes",
				"errors":              "errors",
				"fmt":                 "fmt",
				"go/ast":              "ast",
				"go/build/constraint": "constraint",
				"go/scanner":          "scanner",
				"go/token":            "token",
				"io":                  "io",
				"io/fs":               "fs",
				"os":                  "os",
				"path/filepath":       "filepath",
				"strings":             "strings",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Mode": reflect.TypeOf((*q.Mode)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"ParseDir":      reflect.ValueOf(q.ParseDir),
				"ParseExpr":     reflect.ValueOf(q.ParseExpr),
				"ParseExprFrom": reflect.ValueOf(q.ParseExprFrom),
				"ParseFile":     reflect.ValueOf(q.ParseFile),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"AllErrors":            {Typ: reflect.TypeOf(q.AllErrors), Value: constant.MakeInt64(int64(q.AllErrors))},
				"DeclarationErrors":    {Typ: reflect.TypeOf(q.DeclarationErrors), Value: constant.MakeInt64(int64(q.DeclarationErrors))},
				"ImportsOnly":          {Typ: reflect.TypeOf(q.ImportsOnly), Value: constant.MakeInt64(int64(q.ImportsOnly))},
				"PackageClauseOnly":    {Typ: reflect.TypeOf(q.PackageClauseOnly), Value: constant.MakeInt64(int64(q.PackageClauseOnly))},
				"ParseComments":        {Typ: reflect.TypeOf(q.ParseComments), Value: constant.MakeInt64(int64(q.ParseComments))},
				"SkipObjectResolution": {Typ: reflect.TypeOf(q.SkipObjectResolution), Value: constant.MakeInt64(int64(q.SkipObjectResolution))},
				"SpuriousErrors":       {Typ: reflect.TypeOf(q.SpuriousErrors), Value: constant.MakeInt64(int64(q.SpuriousErrors))},
				"Trace":                {Typ: reflect.TypeOf(q.Trace), Value: constant.MakeInt64(int64(q.Trace))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
