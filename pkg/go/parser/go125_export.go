// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package parser

import (
	q "go/parser"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
			"AllErrors":            {reflect.TypeOf(q.AllErrors), constant.MakeInt64(int64(q.AllErrors))},
			"DeclarationErrors":    {reflect.TypeOf(q.DeclarationErrors), constant.MakeInt64(int64(q.DeclarationErrors))},
			"ImportsOnly":          {reflect.TypeOf(q.ImportsOnly), constant.MakeInt64(int64(q.ImportsOnly))},
			"PackageClauseOnly":    {reflect.TypeOf(q.PackageClauseOnly), constant.MakeInt64(int64(q.PackageClauseOnly))},
			"ParseComments":        {reflect.TypeOf(q.ParseComments), constant.MakeInt64(int64(q.ParseComments))},
			"SkipObjectResolution": {reflect.TypeOf(q.SkipObjectResolution), constant.MakeInt64(int64(q.SkipObjectResolution))},
			"SpuriousErrors":       {reflect.TypeOf(q.SpuriousErrors), constant.MakeInt64(int64(q.SpuriousErrors))},
			"Trace":                {reflect.TypeOf(q.Trace), constant.MakeInt64(int64(q.Trace))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
