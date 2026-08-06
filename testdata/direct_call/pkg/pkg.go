package pkg

import (
	"go/constant"
	"reflect"
)

const Marker = 1

// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

// Inspect returns value when typ describes an int.
func Inspect(typ reflect.Type, value constant.Value) int {
	if typ.Kind() != reflect.Int {
		return 0
	}
	got, _ := constant.Int64Val(value)
	return int(got)
}

// Number is used by the generated direct-call fixture.
type Number int

// Value returns n as an int.
func (n Number) Value() int {
	return int(n)
}
