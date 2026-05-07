// export by github.com/goplus/ixgo/cmd/qexp

package stringslite

import (
	q "github.com/goplus/ixgo/x/stringslite"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/ixgo/x/stringslite", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "stringslite",
			Path: "github.com/goplus/ixgo/x/stringslite",
			Deps: map[string]string{
				"unsafe": "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Clone": q.Clone,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
