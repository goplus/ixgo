/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package memo

import (
	"math"
	"reflect"
)

// ScalarKey returns the scalar value's cache key, preserving floating bits.
func ScalarKey(v any) uint64 {
	switch x := v.(type) {
	case float32:
		return uint64(math.Float32bits(x))
	case float64:
		return math.Float64bits(x)
	}
	r := reflect.ValueOf(v)
	if r.Kind() >= reflect.Int && r.Kind() <= reflect.Int64 {
		return uint64(r.Int())
	}
	return r.Uint()
}

// IntegerBits returns an integer value modulo 2^64.
func IntegerBits(v reflect.Value) uint64 {
	if v.Kind() >= reflect.Int && v.Kind() <= reflect.Int64 {
		return uint64(v.Int())
	}
	return v.Uint()
}

// SetIntegerBits writes low bits using the destination integer width.
func SetIntegerBits(v reflect.Value, bits uint64) {
	if v.Kind() >= reflect.Int && v.Kind() <= reflect.Int64 {
		v.SetInt(int64(bits))
	} else {
		v.SetUint(bits)
	}
}

// SetScalarBits restores numerical state without changing floating-point bits.
func SetScalarBits(v reflect.Value, bits uint64) {
	switch v.Kind() {
	case reflect.Float32:
		v.SetFloat(float64(math.Float32frombits(uint32(bits))))
	case reflect.Float64:
		v.SetFloat(math.Float64frombits(bits))
	default:
		SetIntegerBits(v, bits)
	}
}

// ValueBits reads numerical state, including defined floating-point types.
func ValueBits(v reflect.Value) uint64 {
	switch v.Kind() {
	case reflect.Float32:
		return uint64(math.Float32bits(float32(v.Float())))
	case reflect.Float64:
		return math.Float64bits(v.Float())
	default:
		return IntegerBits(v)
	}
}
