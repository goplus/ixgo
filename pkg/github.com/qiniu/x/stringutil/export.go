// export by github.com/goplus/ixgo/cmd/qexp

package stringutil

import (
	q "github.com/qiniu/x/stringutil"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/qiniu/x/stringutil", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "stringutil",
			Path: "github.com/qiniu/x/stringutil",
			Deps: map[string]string{
				"unicode":      "unicode",
				"unicode/utf8": "utf8",
				"unsafe":       "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Capitalize": q.Capitalize,
				"Concat":     q.Concat,
				"Diff":       q.Diff,
				"String":     q.String,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
