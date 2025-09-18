// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package template

import (
	q "html/template"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"HTMLEscape":       reflect.ValueOf(q.HTMLEscape),
			"HTMLEscapeString": reflect.ValueOf(q.HTMLEscapeString),
			"HTMLEscaper":      reflect.ValueOf(q.HTMLEscaper),
			"IsTrue":           reflect.ValueOf(q.IsTrue),
			"JSEscape":         reflect.ValueOf(q.JSEscape),
			"JSEscapeString":   reflect.ValueOf(q.JSEscapeString),
			"JSEscaper":        reflect.ValueOf(q.JSEscaper),
			"Must":             reflect.ValueOf(q.Must),
			"New":              reflect.ValueOf(q.New),
			"ParseFS":          reflect.ValueOf(q.ParseFS),
			"ParseFiles":       reflect.ValueOf(q.ParseFiles),
			"ParseGlob":        reflect.ValueOf(q.ParseGlob),
			"URLQueryEscaper":  reflect.ValueOf(q.URLQueryEscaper),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"ErrAmbigContext":      {reflect.TypeOf(q.ErrAmbigContext), constant.MakeInt64(int64(q.ErrAmbigContext))},
			"ErrBadHTML":           {reflect.TypeOf(q.ErrBadHTML), constant.MakeInt64(int64(q.ErrBadHTML))},
			"ErrBranchEnd":         {reflect.TypeOf(q.ErrBranchEnd), constant.MakeInt64(int64(q.ErrBranchEnd))},
			"ErrEndContext":        {reflect.TypeOf(q.ErrEndContext), constant.MakeInt64(int64(q.ErrEndContext))},
			"ErrJSTemplate":        {reflect.TypeOf(q.ErrJSTemplate), constant.MakeInt64(int64(q.ErrJSTemplate))},
			"ErrNoSuchTemplate":    {reflect.TypeOf(q.ErrNoSuchTemplate), constant.MakeInt64(int64(q.ErrNoSuchTemplate))},
			"ErrOutputContext":     {reflect.TypeOf(q.ErrOutputContext), constant.MakeInt64(int64(q.ErrOutputContext))},
			"ErrPartialCharset":    {reflect.TypeOf(q.ErrPartialCharset), constant.MakeInt64(int64(q.ErrPartialCharset))},
			"ErrPartialEscape":     {reflect.TypeOf(q.ErrPartialEscape), constant.MakeInt64(int64(q.ErrPartialEscape))},
			"ErrPredefinedEscaper": {reflect.TypeOf(q.ErrPredefinedEscaper), constant.MakeInt64(int64(q.ErrPredefinedEscaper))},
			"ErrRangeLoopReentry":  {reflect.TypeOf(q.ErrRangeLoopReentry), constant.MakeInt64(int64(q.ErrRangeLoopReentry))},
			"ErrSlashAmbig":        {reflect.TypeOf(q.ErrSlashAmbig), constant.MakeInt64(int64(q.ErrSlashAmbig))},
			"OK":                   {reflect.TypeOf(q.OK), constant.MakeInt64(int64(q.OK))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
