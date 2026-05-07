// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package regexp

import (
	q "regexp"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("regexp", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "regexp",
			Path: "regexp",
			Deps: map[string]string{
				"bytes":         "bytes",
				"io":            "io",
				"regexp/syntax": "syntax",
				"slices":        "slices",
				"strconv":       "strconv",
				"strings":       "strings",
				"sync":          "sync",
				"unicode":       "unicode",
				"unicode/utf8":  "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Regexp": reflect.TypeOf((*q.Regexp)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Compile":          q.Compile,
				"CompilePOSIX":     q.CompilePOSIX,
				"Match":            q.Match,
				"MatchReader":      q.MatchReader,
				"MatchString":      q.MatchString,
				"MustCompile":      q.MustCompile,
				"MustCompilePOSIX": q.MustCompilePOSIX,
				"QuoteMeta":        q.QuoteMeta,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
