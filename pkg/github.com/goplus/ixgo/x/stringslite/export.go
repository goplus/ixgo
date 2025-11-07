// export by github.com/goplus/ixgo/cmd/qexp

package stringslite

import (
	q "github.com/goplus/ixgo/x/stringslite"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "stringslite",
		Path: "github.com/goplus/ixgo/x/stringslite",
		Deps: map[string]string{
			"unsafe": "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Clone": reflect.ValueOf(q.Clone),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
