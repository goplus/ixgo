// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package template

import (
	q "html/template"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("html/template", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "template",
			Path: "html/template",
			Deps: map[string]string{
				"bytes":               "bytes",
				"encoding/json":       "json",
				"fmt":                 "fmt",
				"html":                "html",
				"internal/godebug":    "godebug",
				"io":                  "io",
				"io/fs":               "fs",
				"maps":                "maps",
				"os":                  "os",
				"path":                "path",
				"path/filepath":       "filepath",
				"reflect":             "reflect",
				"regexp":              "regexp",
				"strconv":             "strconv",
				"strings":             "strings",
				"sync":                "sync",
				"text/template":       "template",
				"text/template/parse": "parse",
				"unicode":             "unicode",
				"unicode/utf8":        "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"CSS":       reflect.TypeOf((*q.CSS)(nil)).Elem(),
				"Error":     reflect.TypeOf((*q.Error)(nil)).Elem(),
				"ErrorCode": reflect.TypeOf((*q.ErrorCode)(nil)).Elem(),
				"HTML":      reflect.TypeOf((*q.HTML)(nil)).Elem(),
				"HTMLAttr":  reflect.TypeOf((*q.HTMLAttr)(nil)).Elem(),
				"JS":        reflect.TypeOf((*q.JS)(nil)).Elem(),
				"JSStr":     reflect.TypeOf((*q.JSStr)(nil)).Elem(),
				"Srcset":    reflect.TypeOf((*q.Srcset)(nil)).Elem(),
				"Template":  reflect.TypeOf((*q.Template)(nil)).Elem(),
				"URL":       reflect.TypeOf((*q.URL)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"FuncMap": reflect.TypeOf((*q.FuncMap)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
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
			TypedConsts: map[string]interface{}{
				"ErrAmbigContext":      q.ErrAmbigContext,
				"ErrBadHTML":           q.ErrBadHTML,
				"ErrBranchEnd":         q.ErrBranchEnd,
				"ErrEndContext":        q.ErrEndContext,
				"ErrJSTemplate":        q.ErrJSTemplate,
				"ErrNoSuchTemplate":    q.ErrNoSuchTemplate,
				"ErrOutputContext":     q.ErrOutputContext,
				"ErrPartialCharset":    q.ErrPartialCharset,
				"ErrPartialEscape":     q.ErrPartialEscape,
				"ErrPredefinedEscaper": q.ErrPredefinedEscaper,
				"ErrRangeLoopReentry":  q.ErrRangeLoopReentry,
				"ErrSlashAmbig":        q.ErrSlashAmbig,
				"OK":                   q.OK,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
