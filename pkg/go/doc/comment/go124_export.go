// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package comment

import (
	q "go/doc/comment"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "comment",
		Path: "go/doc/comment",
		Deps: map[string]string{
			"bytes":        "bytes",
			"fmt":          "fmt",
			"slices":       "slices",
			"sort":         "sort",
			"strconv":      "strconv",
			"strings":      "strings",
			"unicode":      "unicode",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Block": reflect.TypeOf((*q.Block)(nil)).Elem(),
			"Text":  reflect.TypeOf((*q.Text)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Code":      reflect.TypeOf((*q.Code)(nil)).Elem(),
			"Doc":       reflect.TypeOf((*q.Doc)(nil)).Elem(),
			"DocLink":   reflect.TypeOf((*q.DocLink)(nil)).Elem(),
			"Heading":   reflect.TypeOf((*q.Heading)(nil)).Elem(),
			"Italic":    reflect.TypeOf((*q.Italic)(nil)).Elem(),
			"Link":      reflect.TypeOf((*q.Link)(nil)).Elem(),
			"LinkDef":   reflect.TypeOf((*q.LinkDef)(nil)).Elem(),
			"List":      reflect.TypeOf((*q.List)(nil)).Elem(),
			"ListItem":  reflect.TypeOf((*q.ListItem)(nil)).Elem(),
			"Paragraph": reflect.TypeOf((*q.Paragraph)(nil)).Elem(),
			"Parser":    reflect.TypeOf((*q.Parser)(nil)).Elem(),
			"Plain":     reflect.TypeOf((*q.Plain)(nil)).Elem(),
			"Printer":   reflect.TypeOf((*q.Printer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"DefaultLookupPackage": reflect.ValueOf(q.DefaultLookupPackage),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
