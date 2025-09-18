// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package encoding

import (
	q "encoding"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "encoding",
		Path: "encoding",
		Deps: map[string]string{},
		Interfaces: map[string]reflect.Type{
			"BinaryAppender":    reflect.TypeOf((*q.BinaryAppender)(nil)).Elem(),
			"BinaryMarshaler":   reflect.TypeOf((*q.BinaryMarshaler)(nil)).Elem(),
			"BinaryUnmarshaler": reflect.TypeOf((*q.BinaryUnmarshaler)(nil)).Elem(),
			"TextAppender":      reflect.TypeOf((*q.TextAppender)(nil)).Elem(),
			"TextMarshaler":     reflect.TypeOf((*q.TextMarshaler)(nil)).Elem(),
			"TextUnmarshaler":   reflect.TypeOf((*q.TextUnmarshaler)(nil)).Elem(),
		},
		NamedTypes:    map[string]reflect.Type{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
