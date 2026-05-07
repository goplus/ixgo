// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package template

import (
	q "text/template"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("text/template", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "template",
			Path: "text/template",
			Deps: map[string]string{
				"errors":              "errors",
				"fmt":                 "fmt",
				"internal/fmtsort":    "fmtsort",
				"io":                  "io",
				"io/fs":               "fs",
				"maps":                "maps",
				"net/url":             "url",
				"os":                  "os",
				"path":                "path",
				"path/filepath":       "filepath",
				"reflect":             "reflect",
				"runtime":             "runtime",
				"strings":             "strings",
				"sync":                "sync",
				"text/template/parse": "parse",
				"unicode":             "unicode",
				"unicode/utf8":        "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"ExecError": reflect.TypeOf((*q.ExecError)(nil)).Elem(),
				"FuncMap":   reflect.TypeOf((*q.FuncMap)(nil)).Elem(),
				"Template":  reflect.TypeOf((*q.Template)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"HTMLEscape":       q.HTMLEscape,
				"HTMLEscapeString": q.HTMLEscapeString,
				"HTMLEscaper":      q.HTMLEscaper,
				"IsTrue":           q.IsTrue,
				"JSEscape":         q.JSEscape,
				"JSEscapeString":   q.JSEscapeString,
				"JSEscaper":        q.JSEscaper,
				"Must":             q.Must,
				"New":              q.New,
				"ParseFS":          q.ParseFS,
				"ParseFiles":       q.ParseFiles,
				"ParseGlob":        q.ParseGlob,
				"URLQueryEscaper":  q.URLQueryEscaper,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
