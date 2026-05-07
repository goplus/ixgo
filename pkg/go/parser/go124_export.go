// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package parser

import (
	q "go/parser"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("go/parser", func() *ixgo.Package {
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"ParseDir":      q.ParseDir,
				"ParseExpr":     q.ParseExpr,
				"ParseExprFrom": q.ParseExprFrom,
				"ParseFile":     q.ParseFile,
			},
			TypedConsts: map[string]interface{}{
				"AllErrors":            q.AllErrors,
				"DeclarationErrors":    q.DeclarationErrors,
				"ImportsOnly":          q.ImportsOnly,
				"PackageClauseOnly":    q.PackageClauseOnly,
				"ParseComments":        q.ParseComments,
				"SkipObjectResolution": q.SkipObjectResolution,
				"SpuriousErrors":       q.SpuriousErrors,
				"Trace":                q.Trace,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
