// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package filepath

import (
	q "path/filepath"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("path/filepath", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "filepath",
			Path: "path/filepath",
			Deps: map[string]string{
				"errors":                "errors",
				"internal/bytealg":      "bytealg",
				"internal/filepathlite": "filepathlite",
				"io/fs":                 "fs",
				"os":                    "os",
				"runtime":               "runtime",
				"slices":                "slices",
				"strings":               "strings",
				"syscall":               "syscall",
				"unicode/utf8":          "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"WalkFunc": reflect.TypeOf((*q.WalkFunc)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrBadPattern": &q.ErrBadPattern,
				"SkipAll":       &q.SkipAll,
				"SkipDir":       &q.SkipDir,
			},
			Funcs: map[string]interface{}{
				"Abs":          q.Abs,
				"Base":         q.Base,
				"Clean":        q.Clean,
				"Dir":          q.Dir,
				"EvalSymlinks": q.EvalSymlinks,
				"Ext":          q.Ext,
				"FromSlash":    q.FromSlash,
				"Glob":         q.Glob,
				"HasPrefix":    q.HasPrefix,
				"IsAbs":        q.IsAbs,
				"IsLocal":      q.IsLocal,
				"Join":         q.Join,
				"Localize":     q.Localize,
				"Match":        q.Match,
				"Rel":          q.Rel,
				"Split":        q.Split,
				"SplitList":    q.SplitList,
				"ToSlash":      q.ToSlash,
				"VolumeName":   q.VolumeName,
				"Walk":         q.Walk,
				"WalkDir":      q.WalkDir,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"ListSeparator": {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.ListSeparator))},
				"Separator":     {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.Separator))},
			},
		}
	})
}
