// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package bits

import (
	q "math/bits"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("math/bits", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "bits",
			Path: "math/bits",
			Deps: map[string]string{
				"unsafe": "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Add":             q.Add,
				"Add32":           q.Add32,
				"Add64":           q.Add64,
				"Div":             q.Div,
				"Div32":           q.Div32,
				"Div64":           q.Div64,
				"LeadingZeros":    q.LeadingZeros,
				"LeadingZeros16":  q.LeadingZeros16,
				"LeadingZeros32":  q.LeadingZeros32,
				"LeadingZeros64":  q.LeadingZeros64,
				"LeadingZeros8":   q.LeadingZeros8,
				"Len":             q.Len,
				"Len16":           q.Len16,
				"Len32":           q.Len32,
				"Len64":           q.Len64,
				"Len8":            q.Len8,
				"Mul":             q.Mul,
				"Mul32":           q.Mul32,
				"Mul64":           q.Mul64,
				"OnesCount":       q.OnesCount,
				"OnesCount16":     q.OnesCount16,
				"OnesCount32":     q.OnesCount32,
				"OnesCount64":     q.OnesCount64,
				"OnesCount8":      q.OnesCount8,
				"Rem":             q.Rem,
				"Rem32":           q.Rem32,
				"Rem64":           q.Rem64,
				"Reverse":         q.Reverse,
				"Reverse16":       q.Reverse16,
				"Reverse32":       q.Reverse32,
				"Reverse64":       q.Reverse64,
				"Reverse8":        q.Reverse8,
				"ReverseBytes":    q.ReverseBytes,
				"ReverseBytes16":  q.ReverseBytes16,
				"ReverseBytes32":  q.ReverseBytes32,
				"ReverseBytes64":  q.ReverseBytes64,
				"RotateLeft":      q.RotateLeft,
				"RotateLeft16":    q.RotateLeft16,
				"RotateLeft32":    q.RotateLeft32,
				"RotateLeft64":    q.RotateLeft64,
				"RotateLeft8":     q.RotateLeft8,
				"Sub":             q.Sub,
				"Sub32":           q.Sub32,
				"Sub64":           q.Sub64,
				"TrailingZeros":   q.TrailingZeros,
				"TrailingZeros16": q.TrailingZeros16,
				"TrailingZeros32": q.TrailingZeros32,
				"TrailingZeros64": q.TrailingZeros64,
				"TrailingZeros8":  q.TrailingZeros8,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"UintSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.UintSize))},
			},
		}
	})
}
