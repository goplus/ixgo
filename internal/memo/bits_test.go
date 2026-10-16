/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package memo

import (
	"math"
	"reflect"
	"testing"
)

func TestScalarKey(t *testing.T) {
	cases := []struct {
		v    any
		want uint64
	}{
		{int(-1), math.MaxUint64}, {int8(-128), math.MaxUint64 - 127},
		{uint64(math.MaxUint64), math.MaxUint64}, {uintptr(42), 42},
		{float32(0), 0}, {float32(math.Copysign(0, -1)), 1 << 31},
		{float64(0), 0}, {math.Copysign(0, -1), 1 << 63},
		{math.Float32frombits(0x7fc00001), 0x7fc00001},
		{math.Float64frombits(0x7ff8000000000001), 0x7ff8000000000001},
		{math.Inf(1), 0x7ff0000000000000},
	}
	for _, tc := range cases {
		if got := ScalarKey(tc.v); got != tc.want {
			t.Fatalf("%T: %x want %x", tc.v, got, tc.want)
		}
	}
}

func TestIntegerBits(t *testing.T) {
	for _, v := range []any{int(0), int8(0), int16(0), int32(0), int64(0), uint(0), uint8(0), uint16(0), uint32(0), uint64(0), uintptr(0)} {
		r := reflect.New(reflect.TypeOf(v)).Elem()
		SetIntegerBits(r, math.MaxUint64)
		bits := IntegerBits(r)
		want := uint64(math.MaxUint64)
		if r.Kind() >= reflect.Uint && r.Kind() <= reflect.Uintptr && r.Type().Bits() < 64 {
			want = (uint64(1) << r.Type().Bits()) - 1
		}
		if bits != want {
			t.Fatalf("%T: %x want %x", v, bits, want)
		}
		SetIntegerBits(r, bits+1)
		if IntegerBits(r) != 0 {
			t.Fatalf("%T: overflow did not wrap", v)
		}
	}
}

func TestScalarBitsRoundTrip(t *testing.T) {
	for _, typ := range []reflect.Type{reflect.TypeOf(float32(0)), reflect.TypeOf(float64(0))} {
		cases := []uint64{0, 1, 0x80000000, 0x7f800000, 0x7fc00001, 0xffc00002}
		if typ.Kind() == reflect.Float64 {
			cases = []uint64{0, 1, 1 << 63, 0x7ff0000000000000, 0x7ff8000000000001, 0xfff8000000000002}
		}
		for _, bits := range cases {
			v := reflect.New(typ).Elem()
			SetScalarBits(v, bits)
			if got := ValueBits(v); got != bits {
				t.Fatalf("%v: %x != %x", typ, got, bits)
			}
		}
	}
}
