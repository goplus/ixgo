package memo

import (
	"math"
	"reflect"
	"testing"
)

func TestScalarKeyUintptr(t *testing.T) {
	const want = uintptr(0x1234)
	if got := ScalarKey(want); got != uint64(want) {
		t.Fatalf("ScalarKey(%#x) = %#x", want, got)
	}
}

func TestValueBits(t *testing.T) {
	if got := ValueBits(reflect.ValueOf(float32(math.Copysign(0, -1)))); got != 0x80000000 {
		t.Fatal(got)
	}
	if got := ValueBits(reflect.ValueOf(float64(math.Copysign(0, -1)))); got != 0x8000000000000000 {
		t.Fatal(got)
	}
}
