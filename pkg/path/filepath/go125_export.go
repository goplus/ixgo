// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package filepath

import (
	q "path/filepath"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"ErrBadPattern": reflect.ValueOf(&q.ErrBadPattern),
			"SkipAll":       reflect.ValueOf(&q.SkipAll),
			"SkipDir":       reflect.ValueOf(&q.SkipDir),
		},
		Funcs: map[string]reflect.Value{
			"Abs":          reflect.ValueOf(q.Abs),
			"Base":         reflect.ValueOf(q.Base),
			"Clean":        reflect.ValueOf(q.Clean),
			"Dir":          reflect.ValueOf(q.Dir),
			"EvalSymlinks": reflect.ValueOf(q.EvalSymlinks),
			"Ext":          reflect.ValueOf(q.Ext),
			"FromSlash":    reflect.ValueOf(q.FromSlash),
			"Glob":         reflect.ValueOf(q.Glob),
			"HasPrefix":    reflect.ValueOf(q.HasPrefix),
			"IsAbs":        reflect.ValueOf(q.IsAbs),
			"IsLocal":      reflect.ValueOf(q.IsLocal),
			"Join":         reflect.ValueOf(q.Join),
			"Localize":     reflect.ValueOf(q.Localize),
			"Match":        reflect.ValueOf(q.Match),
			"Rel":          reflect.ValueOf(q.Rel),
			"Split":        reflect.ValueOf(q.Split),
			"SplitList":    reflect.ValueOf(q.SplitList),
			"ToSlash":      reflect.ValueOf(q.ToSlash),
			"VolumeName":   reflect.ValueOf(q.VolumeName),
			"Walk":         reflect.ValueOf(q.Walk),
			"WalkDir":      reflect.ValueOf(q.WalkDir),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"ListSeparator": {"untyped rune", constant.MakeInt64(int64(q.ListSeparator))},
			"Separator":     {"untyped rune", constant.MakeInt64(int64(q.Separator))},
		},
	})
}
