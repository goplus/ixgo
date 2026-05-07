// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package strings

import (
	q "strings"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("strings", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "strings",
			Path: "strings",
			Deps: map[string]string{
				"errors":               "errors",
				"internal/abi":         "abi",
				"internal/bytealg":     "bytealg",
				"internal/stringslite": "stringslite",
				"io":                   "io",
				"iter":                 "iter",
				"math/bits":            "bits",
				"sync":                 "sync",
				"unicode":              "unicode",
				"unicode/utf8":         "utf8",
				"unsafe":               "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Builder":  reflect.TypeOf((*q.Builder)(nil)).Elem(),
				"Reader":   reflect.TypeOf((*q.Reader)(nil)).Elem(),
				"Replacer": reflect.TypeOf((*q.Replacer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Clone":          q.Clone,
				"Compare":        q.Compare,
				"Contains":       q.Contains,
				"ContainsAny":    q.ContainsAny,
				"ContainsFunc":   q.ContainsFunc,
				"ContainsRune":   q.ContainsRune,
				"Count":          q.Count,
				"Cut":            q.Cut,
				"CutPrefix":      q.CutPrefix,
				"CutSuffix":      q.CutSuffix,
				"EqualFold":      q.EqualFold,
				"Fields":         q.Fields,
				"FieldsFunc":     q.FieldsFunc,
				"FieldsFuncSeq":  q.FieldsFuncSeq,
				"FieldsSeq":      q.FieldsSeq,
				"HasPrefix":      q.HasPrefix,
				"HasSuffix":      q.HasSuffix,
				"Index":          q.Index,
				"IndexAny":       q.IndexAny,
				"IndexByte":      q.IndexByte,
				"IndexFunc":      q.IndexFunc,
				"IndexRune":      q.IndexRune,
				"Join":           q.Join,
				"LastIndex":      q.LastIndex,
				"LastIndexAny":   q.LastIndexAny,
				"LastIndexByte":  q.LastIndexByte,
				"LastIndexFunc":  q.LastIndexFunc,
				"Lines":          q.Lines,
				"Map":            q.Map,
				"NewReader":      q.NewReader,
				"NewReplacer":    q.NewReplacer,
				"Repeat":         q.Repeat,
				"Replace":        q.Replace,
				"ReplaceAll":     q.ReplaceAll,
				"Split":          q.Split,
				"SplitAfter":     q.SplitAfter,
				"SplitAfterN":    q.SplitAfterN,
				"SplitAfterSeq":  q.SplitAfterSeq,
				"SplitN":         q.SplitN,
				"SplitSeq":       q.SplitSeq,
				"Title":          q.Title,
				"ToLower":        q.ToLower,
				"ToLowerSpecial": q.ToLowerSpecial,
				"ToTitle":        q.ToTitle,
				"ToTitleSpecial": q.ToTitleSpecial,
				"ToUpper":        q.ToUpper,
				"ToUpperSpecial": q.ToUpperSpecial,
				"ToValidUTF8":    q.ToValidUTF8,
				"Trim":           q.Trim,
				"TrimFunc":       q.TrimFunc,
				"TrimLeft":       q.TrimLeft,
				"TrimLeftFunc":   q.TrimLeftFunc,
				"TrimPrefix":     q.TrimPrefix,
				"TrimRight":      q.TrimRight,
				"TrimRightFunc":  q.TrimRightFunc,
				"TrimSpace":      q.TrimSpace,
				"TrimSuffix":     q.TrimSuffix,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
