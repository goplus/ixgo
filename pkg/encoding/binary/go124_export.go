// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package binary

import (
	q "encoding/binary"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "binary",
		Path: "encoding/binary",
		Deps: map[string]string{
			"errors":  "errors",
			"io":      "io",
			"math":    "math",
			"reflect": "reflect",
			"slices":  "slices",
			"sync":    "sync",
		},
		Interfaces: map[string]reflect.Type{
			"AppendByteOrder": reflect.TypeOf((*q.AppendByteOrder)(nil)).Elem(),
			"ByteOrder":       reflect.TypeOf((*q.ByteOrder)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"BigEndian":    reflect.ValueOf(&q.BigEndian),
			"LittleEndian": reflect.ValueOf(&q.LittleEndian),
			"NativeEndian": reflect.ValueOf(&q.NativeEndian),
		},
		Funcs: map[string]reflect.Value{
			"Append":        reflect.ValueOf(q.Append),
			"AppendUvarint": reflect.ValueOf(q.AppendUvarint),
			"AppendVarint":  reflect.ValueOf(q.AppendVarint),
			"Decode":        reflect.ValueOf(q.Decode),
			"Encode":        reflect.ValueOf(q.Encode),
			"PutUvarint":    reflect.ValueOf(q.PutUvarint),
			"PutVarint":     reflect.ValueOf(q.PutVarint),
			"Read":          reflect.ValueOf(q.Read),
			"ReadUvarint":   reflect.ValueOf(q.ReadUvarint),
			"ReadVarint":    reflect.ValueOf(q.ReadVarint),
			"Size":          reflect.ValueOf(q.Size),
			"Uvarint":       reflect.ValueOf(q.Uvarint),
			"Varint":        reflect.ValueOf(q.Varint),
			"Write":         reflect.ValueOf(q.Write),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"MaxVarintLen16": {"untyped int", constant.MakeInt64(int64(q.MaxVarintLen16))},
			"MaxVarintLen32": {"untyped int", constant.MakeInt64(int64(q.MaxVarintLen32))},
			"MaxVarintLen64": {"untyped int", constant.MakeInt64(int64(q.MaxVarintLen64))},
		},
	})
}
