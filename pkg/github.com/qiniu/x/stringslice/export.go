// export by github.com/goplus/ixgo/cmd/qexp

package stringslice

import (
	q "github.com/qiniu/x/stringslice"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/qiniu/x/stringslice", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "stringslice",
			Path: "github.com/qiniu/x/stringslice",
			Deps: map[string]string{
				"github.com/qiniu/x/stringutil": "stringutil",
				"strings":                       "strings",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Capitalize": q.Capitalize,
				"Repeat":     q.Repeat,
				"Replace":    q.Replace,
				"ReplaceAll": q.ReplaceAll,
				"ToLower":    q.ToLower,
				"ToTitle":    q.ToTitle,
				"ToUpper":    q.ToUpper,
				"Trim":       q.Trim,
				"TrimLeft":   q.TrimLeft,
				"TrimPrefix": q.TrimPrefix,
				"TrimRight":  q.TrimRight,
				"TrimSpace":  q.TrimSpace,
				"TrimSuffix": q.TrimSuffix,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
