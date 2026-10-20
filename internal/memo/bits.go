// Package memo contains shared runtime helpers for memoization.
package memo

import (
	"math"
	"reflect"
)

// ScalarKey returns a stable key while preserving floating-point bit patterns.
func ScalarKey(v any) uint64 {
	switch x := v.(type) {
	case float32:
		return uint64(math.Float32bits(x))
	case float64:
		return math.Float64bits(x)
	case int:
		return uint64(x)
	case int8:
		return uint64(x)
	case int16:
		return uint64(x)
	case int32:
		return uint64(x)
	case int64:
		return uint64(x)
	case uint:
		return uint64(x)
	case uint8:
		return uint64(x)
	case uint16:
		return uint64(x)
	case uint32:
		return uint64(x)
	case uint64:
		return x
	case uintptr:
		return uint64(x)
	}
	return ValueBits(reflect.ValueOf(v))
}

func IntegerBits(v reflect.Value) uint64 {
	if v.Kind() >= reflect.Int && v.Kind() <= reflect.Int64 {
		return uint64(v.Int())
	}
	return v.Uint()
}

func SetIntegerBits(v reflect.Value, bits uint64) {
	if v.Kind() >= reflect.Int && v.Kind() <= reflect.Int64 {
		v.SetInt(int64(bits))
	} else {
		v.SetUint(bits)
	}
}

func SetScalarBits(v reflect.Value, bits uint64) {
	switch v.Kind() {
	case reflect.Float32:
		v.Set(reflect.ValueOf(math.Float32frombits(uint32(bits))).Convert(v.Type()))
	case reflect.Float64:
		v.SetFloat(math.Float64frombits(bits))
	default:
		SetIntegerBits(v, bits)
	}
}

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
