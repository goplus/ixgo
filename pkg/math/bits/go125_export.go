// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package bits

import (
	q "math/bits"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "bits",
		Path: "math/bits",
		Deps: map[string]string{
			"unsafe": "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Add":             reflect.ValueOf(q.Add),
			"Add32":           reflect.ValueOf(q.Add32),
			"Add64":           reflect.ValueOf(q.Add64),
			"Div":             reflect.ValueOf(q.Div),
			"Div32":           reflect.ValueOf(q.Div32),
			"Div64":           reflect.ValueOf(q.Div64),
			"LeadingZeros":    reflect.ValueOf(q.LeadingZeros),
			"LeadingZeros16":  reflect.ValueOf(q.LeadingZeros16),
			"LeadingZeros32":  reflect.ValueOf(q.LeadingZeros32),
			"LeadingZeros64":  reflect.ValueOf(q.LeadingZeros64),
			"LeadingZeros8":   reflect.ValueOf(q.LeadingZeros8),
			"Len":             reflect.ValueOf(q.Len),
			"Len16":           reflect.ValueOf(q.Len16),
			"Len32":           reflect.ValueOf(q.Len32),
			"Len64":           reflect.ValueOf(q.Len64),
			"Len8":            reflect.ValueOf(q.Len8),
			"Mul":             reflect.ValueOf(q.Mul),
			"Mul32":           reflect.ValueOf(q.Mul32),
			"Mul64":           reflect.ValueOf(q.Mul64),
			"OnesCount":       reflect.ValueOf(q.OnesCount),
			"OnesCount16":     reflect.ValueOf(q.OnesCount16),
			"OnesCount32":     reflect.ValueOf(q.OnesCount32),
			"OnesCount64":     reflect.ValueOf(q.OnesCount64),
			"OnesCount8":      reflect.ValueOf(q.OnesCount8),
			"Rem":             reflect.ValueOf(q.Rem),
			"Rem32":           reflect.ValueOf(q.Rem32),
			"Rem64":           reflect.ValueOf(q.Rem64),
			"Reverse":         reflect.ValueOf(q.Reverse),
			"Reverse16":       reflect.ValueOf(q.Reverse16),
			"Reverse32":       reflect.ValueOf(q.Reverse32),
			"Reverse64":       reflect.ValueOf(q.Reverse64),
			"Reverse8":        reflect.ValueOf(q.Reverse8),
			"ReverseBytes":    reflect.ValueOf(q.ReverseBytes),
			"ReverseBytes16":  reflect.ValueOf(q.ReverseBytes16),
			"ReverseBytes32":  reflect.ValueOf(q.ReverseBytes32),
			"ReverseBytes64":  reflect.ValueOf(q.ReverseBytes64),
			"RotateLeft":      reflect.ValueOf(q.RotateLeft),
			"RotateLeft16":    reflect.ValueOf(q.RotateLeft16),
			"RotateLeft32":    reflect.ValueOf(q.RotateLeft32),
			"RotateLeft64":    reflect.ValueOf(q.RotateLeft64),
			"RotateLeft8":     reflect.ValueOf(q.RotateLeft8),
			"Sub":             reflect.ValueOf(q.Sub),
			"Sub32":           reflect.ValueOf(q.Sub32),
			"Sub64":           reflect.ValueOf(q.Sub64),
			"TrailingZeros":   reflect.ValueOf(q.TrailingZeros),
			"TrailingZeros16": reflect.ValueOf(q.TrailingZeros16),
			"TrailingZeros32": reflect.ValueOf(q.TrailingZeros32),
			"TrailingZeros64": reflect.ValueOf(q.TrailingZeros64),
			"TrailingZeros8":  reflect.ValueOf(q.TrailingZeros8),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"UintSize": {"untyped int", constant.MakeInt64(int64(q.UintSize))},
		},
	})
}
