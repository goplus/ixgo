// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package bytes

import (
	q "bytes"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("bytes", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "bytes",
			Path: "bytes",
			Deps: map[string]string{
				"errors":           "errors",
				"internal/bytealg": "bytealg",
				"io":               "io",
				"iter":             "iter",
				"math/bits":        "bits",
				"unicode":          "unicode",
				"unicode/utf8":     "utf8",
				"unsafe":           "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Buffer": reflect.TypeOf((*q.Buffer)(nil)).Elem(),
				"Reader": reflect.TypeOf((*q.Reader)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrTooLarge": &q.ErrTooLarge,
			},
			Funcs: map[string]interface{}{
				"Clone":           q.Clone,
				"Compare":         q.Compare,
				"Contains":        q.Contains,
				"ContainsAny":     q.ContainsAny,
				"ContainsFunc":    q.ContainsFunc,
				"ContainsRune":    q.ContainsRune,
				"Count":           q.Count,
				"Cut":             q.Cut,
				"CutPrefix":       q.CutPrefix,
				"CutSuffix":       q.CutSuffix,
				"Equal":           q.Equal,
				"EqualFold":       q.EqualFold,
				"Fields":          q.Fields,
				"FieldsFunc":      q.FieldsFunc,
				"FieldsFuncSeq":   q.FieldsFuncSeq,
				"FieldsSeq":       q.FieldsSeq,
				"HasPrefix":       q.HasPrefix,
				"HasSuffix":       q.HasSuffix,
				"Index":           q.Index,
				"IndexAny":        q.IndexAny,
				"IndexByte":       q.IndexByte,
				"IndexFunc":       q.IndexFunc,
				"IndexRune":       q.IndexRune,
				"Join":            q.Join,
				"LastIndex":       q.LastIndex,
				"LastIndexAny":    q.LastIndexAny,
				"LastIndexByte":   q.LastIndexByte,
				"LastIndexFunc":   q.LastIndexFunc,
				"Lines":           q.Lines,
				"Map":             q.Map,
				"NewBuffer":       q.NewBuffer,
				"NewBufferString": q.NewBufferString,
				"NewReader":       q.NewReader,
				"Repeat":          q.Repeat,
				"Replace":         q.Replace,
				"ReplaceAll":      q.ReplaceAll,
				"Runes":           q.Runes,
				"Split":           q.Split,
				"SplitAfter":      q.SplitAfter,
				"SplitAfterN":     q.SplitAfterN,
				"SplitAfterSeq":   q.SplitAfterSeq,
				"SplitN":          q.SplitN,
				"SplitSeq":        q.SplitSeq,
				"Title":           q.Title,
				"ToLower":         q.ToLower,
				"ToLowerSpecial":  q.ToLowerSpecial,
				"ToTitle":         q.ToTitle,
				"ToTitleSpecial":  q.ToTitleSpecial,
				"ToUpper":         q.ToUpper,
				"ToUpperSpecial":  q.ToUpperSpecial,
				"ToValidUTF8":     q.ToValidUTF8,
				"Trim":            q.Trim,
				"TrimFunc":        q.TrimFunc,
				"TrimLeft":        q.TrimLeft,
				"TrimLeftFunc":    q.TrimLeftFunc,
				"TrimPrefix":      q.TrimPrefix,
				"TrimRight":       q.TrimRight,
				"TrimRightFunc":   q.TrimRightFunc,
				"TrimSpace":       q.TrimSpace,
				"TrimSuffix":      q.TrimSuffix,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"MinRead": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MinRead))},
			},
		}
	})
}
