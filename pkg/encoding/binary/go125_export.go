// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package binary

import (
	q "encoding/binary"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("encoding/binary", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars: map[string]interface{}{
				"BigEndian":    &q.BigEndian,
				"LittleEndian": &q.LittleEndian,
				"NativeEndian": &q.NativeEndian,
			},
			Funcs: map[string]interface{}{
				"Append":        q.Append,
				"AppendUvarint": q.AppendUvarint,
				"AppendVarint":  q.AppendVarint,
				"Decode":        q.Decode,
				"Encode":        q.Encode,
				"PutUvarint":    q.PutUvarint,
				"PutVarint":     q.PutVarint,
				"Read":          q.Read,
				"ReadUvarint":   q.ReadUvarint,
				"ReadVarint":    q.ReadVarint,
				"Size":          q.Size,
				"Uvarint":       q.Uvarint,
				"Varint":        q.Varint,
				"Write":         q.Write,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"MaxVarintLen16": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MaxVarintLen16))},
				"MaxVarintLen32": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MaxVarintLen32))},
				"MaxVarintLen64": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MaxVarintLen64))},
			},
		}
	})
}
