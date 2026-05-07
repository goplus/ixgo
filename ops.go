// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ixgo

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"unsafe"

	"golang.org/x/tools/go/ssa"
)

func xtypeValue(c *ssa.Const, kind types.BasicKind) value {
	switch kind {
	case types.Bool, types.UntypedBool:
		return constant.BoolVal(c.Value)
	case types.Int, types.UntypedInt:
		// Assume sizeof(int) is same on host and target.
		return int(c.Int64())
	case types.Int8:
		return int8(c.Int64())
	case types.Int16:
		return int16(c.Int64())
	case types.Int32, types.UntypedRune:
		return int32(c.Int64())
	case types.Int64:
		return c.Int64()
	case types.Uint:
		// Assume sizeof(uint) is same on host and target.
		return uint(c.Uint64())
	case types.Uint8:
		return uint8(c.Uint64())
	case types.Uint16:
		return uint16(c.Uint64())
	case types.Uint32:
		return uint32(c.Uint64())
	case types.Uint64:
		return c.Uint64()
	case types.Uintptr:
		// Assume sizeof(uintptr) is same on host and target.
		return uintptr(c.Uint64())
	case types.Float32:
		return float32(c.Float64())
	case types.Float64, types.UntypedFloat:
		return c.Float64()
	case types.Complex64:
		return complex64(c.Complex128())
	case types.Complex128, types.UntypedComplex:
		return c.Complex128()
	case types.String, types.UntypedString:
		if c.Value.Kind() == constant.String {
			return constant.StringVal(c.Value)
		}
		return string(rune(c.Int64()))
	case types.UnsafePointer:
		return unsafe.Pointer(uintptr(c.Uint64()))
	}
	panic("unreachable")
}

// constValue returns the value of the constant with the
// dynamic type tag appropriate for c.Type().
func constToValue(i *Interp, c *ssa.Const) value {
	typ := c.Type()
	if c.Value == nil {
		if xtype, ok := typ.(*types.Basic); ok && xtype.Kind() == types.UntypedNil {
			return nil
		}
		return reflect.Zero(i.preToType(typ)).Interface()
	}
	if xtype, ok := typ.(*types.Basic); ok {
		return xtypeValue(c, xtype.Kind())
	} else if xtype, ok := typ.Underlying().(*types.Basic); ok {
		v := xtypeValue(c, xtype.Kind())
		nv := reflect.New(i.preToType(typ)).Elem()
		SetValue(nv, reflect.ValueOf(v))
		return nv.Interface()
	}
	panic(fmt.Sprintf("unparser constValue: %s", c))
}

func globalToValue(i *Interp, key *ssa.Global) (interface{}, bool) {
	if key.Pkg != nil {
		pkgpath := key.Pkg.Pkg.Path()
		if pkg, ok := i.installed(pkgpath); ok {
			if ext, ok := pkg.Vars[key.Name()]; ok {
				return ext, true
			}
		}
	}
	if v, ok := i.globals[key.String()]; ok {
		return v, true
	}
	return nil, false
}

func staticToValue(i *Interp, value ssa.Value) (interface{}, bool) {
	switch v := value.(type) {
	case *ssa.Global:
		return globalToValue(i, v)
	case *ssa.Const:
		return constToValue(i, v), true
	}
	return nil, false
}

// asInt converts x, which must be an integer, to an int suitable for
// use as a slice or array index or operand to make().
func asInt(x value) int {
	switch x := x.(type) {
	case int:
		return x
	case int8:
		return int(x)
	case int16:
		return int(x)
	case int32:
		return int(x)
	case int64:
		return int(x)
	case uint:
		return int(x)
	case uint8:
		return int(x)
	case uint16:
		return int(x)
	case uint32:
		return int(x)
	case uint64:
		return int(x)
	case uintptr:
		return int(x)
	default:
		v := reflect.ValueOf(x)
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int(v.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return int(v.Uint())
		}
	}
	panic(fmt.Sprintf("cannot convert %T to int", x))
}

// asUint64 converts x, which must be an unsigned integer, to a uint64
// suitable for use as a bitwise shift count.
func asUint64(x value) uint64 {
	switch x := x.(type) {
	case int:
		if x >= 0 {
			return uint64(x)
		}
	case int8:
		if x >= 0 {
			return uint64(x)
		}
	case int16:
		if x >= 0 {
			return uint64(x)
		}
	case int32:
		if x >= 0 {
			return uint64(x)
		}
	case int64:
		if x >= 0 {
			return uint64(x)
		}
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
	default:
		v := reflect.ValueOf(x)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return v.Uint()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x := v.Int()
			if x >= 0 {
				return uint64(x)
			}
		default:
			panic(fmt.Sprintf("cannot convert %T to uint64", x))
		}
	}
	panic(RuntimeError("negative shift amount"))
}

// slice returns x[lo:hi:max].  Any of lo, hi and max may be nil.
func slice(fr *frame, instr *ssa.Slice, makesliceCheck bool, ix, ih, il, im register) reflect.Value {
	// x := fr.get(instr.X)
	x := fr.reg(ix)
	var Len, Cap int
	v := reflect.ValueOf(x)
	// *array
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	kind := v.Kind()
	switch kind {
	case reflect.String:
		Len = v.Len()
		Cap = Len
	case reflect.Slice, reflect.Array:
		Len = v.Len()
		Cap = v.Cap()
	}

	lo := 0
	hi := Len
	max := Cap
	var slice3 bool
	if instr.Low != nil {
		// lo = asInt(fr.get(instr.Low))
		lo = asInt(fr.reg(il))
	}
	if instr.High != nil {
		// hi = asInt(fr.get(instr.High))
		hi = asInt(fr.reg(ih))
	}
	if instr.Max != nil {
		// max = asInt(fr.get(instr.Max))
		max = asInt(fr.reg(im))
		slice3 = true
	}

	if makesliceCheck {
		if hi < 0 {
			panic(fr.runtimeError(instr, "makeslice: len out of range"))
		} else if hi > max {
			panic(fr.runtimeError(instr, "makeslice: cap out of range"))
		}
	} else {
		if slice3 {
			if max < 0 {
				panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [::%v]", max)))
			} else if max > Cap {
				if kind == reflect.Slice {
					panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [::%v] with capacity %v", max, Cap)))
				} else {
					panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [::%v] with length %v", max, Cap)))
				}
			} else if hi < 0 {
				panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [:%v:]", hi)))
			} else if hi > max {
				panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [:%v:%v]", hi, max)))
			} else if lo < 0 {
				panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [%v::]", lo)))
			} else if lo > hi {
				panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [%v:%v:]", lo, hi)))
			}
		} else {
			if hi < 0 {
				panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [:%v]", hi)))
			} else if hi > Cap {
				if kind == reflect.Slice {
					panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [:%v] with capacity %v", hi, Cap)))
				} else {
					panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [:%v] with length %v", hi, Cap)))
				}
			} else if lo < 0 {
				panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [%v:]", lo)))
			} else if lo > hi {
				panic(fr.runtimeError(instr, fmt.Sprintf("slice bounds out of range [%v:%v]", lo, hi)))
			}
		}
	}
	switch kind {
	case reflect.String:
		// optimization x[len(x):], see $GOROOT/test/slicecap.go
		if lo == hi {
			return v.Slice(0, 0)
		}
		return v.Slice(lo, hi)
	case reflect.Slice, reflect.Array:
		return v.Slice3(lo, hi, max)
	}
	panic(fmt.Sprintf("slice: unexpected X type: %T", x))
}

func opADD(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) + y.(int)
	case int8:
		return x.(int8) + y.(int8)
	case int16:
		return x.(int16) + y.(int16)
	case int32:
		return x.(int32) + y.(int32)
	case int64:
		return x.(int64) + y.(int64)
	case uint:
		return x.(uint) + y.(uint)
	case uint8:
		return x.(uint8) + y.(uint8)
	case uint16:
		return x.(uint16) + y.(uint16)
	case uint32:
		return x.(uint32) + y.(uint32)
	case uint64:
		return x.(uint64) + y.(uint64)
	case uintptr:
		return x.(uintptr) + y.(uintptr)
	case float32:
		return x.(float32) + y.(float32)
	case float64:
		return x.(float64) + y.(float64)
	case complex64:
		return x.(complex64) + y.(complex64)
	case complex128:
		return x.(complex128) + y.(complex128)
	case string:
		return x.(string) + y.(string)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			r := reflect.New(vx.Type()).Elem()
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(vx.Int() + vy.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(vx.Uint() + vy.Uint())
			case reflect.Float32, reflect.Float64:
				r.SetFloat(vx.Float() + vy.Float())
			case reflect.Complex64, reflect.Complex128:
				r.SetComplex(vx.Complex() + vy.Complex())
			case reflect.String:
				r.SetString(vx.String() + vy.String())
			default:
				goto failed
			}
			return r.Interface()
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T + %T", x, y))
}

func opSUB(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) - y.(int)
	case int8:
		return x.(int8) - y.(int8)
	case int16:
		return x.(int16) - y.(int16)
	case int32:
		return x.(int32) - y.(int32)
	case int64:
		return x.(int64) - y.(int64)
	case uint:
		return x.(uint) - y.(uint)
	case uint8:
		return x.(uint8) - y.(uint8)
	case uint16:
		return x.(uint16) - y.(uint16)
	case uint32:
		return x.(uint32) - y.(uint32)
	case uint64:
		return x.(uint64) - y.(uint64)
	case uintptr:
		return x.(uintptr) - y.(uintptr)
	case float32:
		return x.(float32) - y.(float32)
	case float64:
		return x.(float64) - y.(float64)
	case complex64:
		return x.(complex64) - y.(complex64)
	case complex128:
		return x.(complex128) - y.(complex128)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			r := reflect.New(vx.Type()).Elem()
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(vx.Int() - vy.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(vx.Uint() - vy.Uint())
			case reflect.Float32, reflect.Float64:
				r.SetFloat(vx.Float() - vy.Float())
			case reflect.Complex64, reflect.Complex128:
				r.SetComplex(vx.Complex() - vy.Complex())
			default:
				goto failed
			}
			return r.Interface()
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T - %T", x, y))
}

func opMUL(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) * y.(int)
	case int8:
		return x.(int8) * y.(int8)
	case int16:
		return x.(int16) * y.(int16)
	case int32:
		return x.(int32) * y.(int32)
	case int64:
		return x.(int64) * y.(int64)
	case uint:
		return x.(uint) * y.(uint)
	case uint8:
		return x.(uint8) * y.(uint8)
	case uint16:
		return x.(uint16) * y.(uint16)
	case uint32:
		return x.(uint32) * y.(uint32)
	case uint64:
		return x.(uint64) * y.(uint64)
	case uintptr:
		return x.(uintptr) * y.(uintptr)
	case float32:
		return x.(float32) * y.(float32)
	case float64:
		return x.(float64) * y.(float64)
	case complex64:
		return x.(complex64) * y.(complex64)
	case complex128:
		return x.(complex128) * y.(complex128)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			r := reflect.New(vx.Type()).Elem()
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(vx.Int() * vy.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(vx.Uint() * vy.Uint())
			case reflect.Float32, reflect.Float64:
				r.SetFloat(vx.Float() * vy.Float())
			case reflect.Complex64, reflect.Complex128:
				r.SetComplex(vx.Complex() * vy.Complex())
			default:
				goto failed
			}
			return r.Interface()
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T * %T", x, y))
}

func opQuo(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) / y.(int)
	case int8:
		return x.(int8) / y.(int8)
	case int16:
		return x.(int16) / y.(int16)
	case int32:
		return x.(int32) / y.(int32)
	case int64:
		return x.(int64) / y.(int64)
	case uint:
		return x.(uint) / y.(uint)
	case uint8:
		return x.(uint8) / y.(uint8)
	case uint16:
		return x.(uint16) / y.(uint16)
	case uint32:
		return x.(uint32) / y.(uint32)
	case uint64:
		return x.(uint64) / y.(uint64)
	case uintptr:
		return x.(uintptr) / y.(uintptr)
	case float32:
		return x.(float32) / y.(float32)
	case float64:
		return x.(float64) / y.(float64)
	case complex64:
		return x.(complex64) / y.(complex64)
	case complex128:
		return x.(complex128) / y.(complex128)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			r := reflect.New(vx.Type()).Elem()
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(vx.Int() / vy.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(vx.Uint() / vy.Uint())
			case reflect.Float32, reflect.Float64:
				r.SetFloat(vx.Float() / vy.Float())
			case reflect.Complex64, reflect.Complex128:
				r.SetComplex(vx.Complex() / vy.Complex())
			default:
				goto failed
			}
			return r.Interface()
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T / %T", x, y))
}

func opREM(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) % y.(int)
	case int8:
		return x.(int8) % y.(int8)
	case int16:
		return x.(int16) % y.(int16)
	case int32:
		return x.(int32) % y.(int32)
	case int64:
		return x.(int64) % y.(int64)
	case uint:
		return x.(uint) % y.(uint)
	case uint8:
		return x.(uint8) % y.(uint8)
	case uint16:
		return x.(uint16) % y.(uint16)
	case uint32:
		return x.(uint32) % y.(uint32)
	case uint64:
		return x.(uint64) % y.(uint64)
	case uintptr:
		return x.(uintptr) % y.(uintptr)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			r := reflect.New(vx.Type()).Elem()
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(vx.Int() % vy.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(vx.Uint() % vy.Uint())
			default:
				goto failed
			}
			return r.Interface()
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T %% %T", x, y))
}

func opAND(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) & y.(int)
	case int8:
		return x.(int8) & y.(int8)
	case int16:
		return x.(int16) & y.(int16)
	case int32:
		return x.(int32) & y.(int32)
	case int64:
		return x.(int64) & y.(int64)
	case uint:
		return x.(uint) & y.(uint)
	case uint8:
		return x.(uint8) & y.(uint8)
	case uint16:
		return x.(uint16) & y.(uint16)
	case uint32:
		return x.(uint32) & y.(uint32)
	case uint64:
		return x.(uint64) & y.(uint64)
	case uintptr:
		return x.(uintptr) & y.(uintptr)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			r := reflect.New(vx.Type()).Elem()
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(vx.Int() & vy.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(vx.Uint() & vy.Uint())
			default:
				goto failed
			}
			return r.Interface()
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T && %T", x, y))
}

func opOR(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) | y.(int)
	case int8:
		return x.(int8) | y.(int8)
	case int16:
		return x.(int16) | y.(int16)
	case int32:
		return x.(int32) | y.(int32)
	case int64:
		return x.(int64) | y.(int64)
	case uint:
		return x.(uint) | y.(uint)
	case uint8:
		return x.(uint8) | y.(uint8)
	case uint16:
		return x.(uint16) | y.(uint16)
	case uint32:
		return x.(uint32) | y.(uint32)
	case uint64:
		return x.(uint64) | y.(uint64)
	case uintptr:
		return x.(uintptr) | y.(uintptr)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			r := reflect.New(vx.Type()).Elem()
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(vx.Int() | vy.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(vx.Uint() | vy.Uint())
			default:
				goto failed
			}
			return r.Interface()
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T | %T", x, y))
}

func opXOR(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) ^ y.(int)
	case int8:
		return x.(int8) ^ y.(int8)
	case int16:
		return x.(int16) ^ y.(int16)
	case int32:
		return x.(int32) ^ y.(int32)
	case int64:
		return x.(int64) ^ y.(int64)
	case uint:
		return x.(uint) ^ y.(uint)
	case uint8:
		return x.(uint8) ^ y.(uint8)
	case uint16:
		return x.(uint16) ^ y.(uint16)
	case uint32:
		return x.(uint32) ^ y.(uint32)
	case uint64:
		return x.(uint64) ^ y.(uint64)
	case uintptr:
		return x.(uintptr) ^ y.(uintptr)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			r := reflect.New(vx.Type()).Elem()
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(vx.Int() ^ vy.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(vx.Uint() ^ vy.Uint())
			default:
				goto failed
			}
			return r.Interface()
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T ^ %T", x, y))
}

func opANDNOT(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) &^ y.(int)
	case int8:
		return x.(int8) &^ y.(int8)
	case int16:
		return x.(int16) &^ y.(int16)
	case int32:
		return x.(int32) &^ y.(int32)
	case int64:
		return x.(int64) &^ y.(int64)
	case uint:
		return x.(uint) &^ y.(uint)
	case uint8:
		return x.(uint8) &^ y.(uint8)
	case uint16:
		return x.(uint16) &^ y.(uint16)
	case uint32:
		return x.(uint32) &^ y.(uint32)
	case uint64:
		return x.(uint64) &^ y.(uint64)
	case uintptr:
		return x.(uintptr) &^ y.(uintptr)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			r := reflect.New(vx.Type()).Elem()
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(vx.Int() &^ vy.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(vx.Uint() &^ vy.Uint())
			default:
				goto failed
			}
			return r.Interface()
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T &^ %T", x, y))
}

func opSHL(x, _y value) value {
	y := asUint64(_y)
	switch x.(type) {
	case int:
		return x.(int) << y
	case int8:
		return x.(int8) << y
	case int16:
		return x.(int16) << y
	case int32:
		return x.(int32) << y
	case int64:
		return x.(int64) << y
	case uint:
		return x.(uint) << y
	case uint8:
		return x.(uint8) << y
	case uint16:
		return x.(uint16) << y
	case uint32:
		return x.(uint32) << y
	case uint64:
		return x.(uint64) << y
	case uintptr:
		return x.(uintptr) << y
	default:
		vx := reflect.ValueOf(x)
		r := reflect.New(vx.Type()).Elem()
		switch vx.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			r.SetInt(vx.Int() << y)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			r.SetUint(vx.Uint() << y)
		default:
			goto failed
		}
		return r.Interface()
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T << %T", x, y))
}

func opSHR(x, _y value) value {
	y := asUint64(_y)
	switch x.(type) {
	case int:
		return x.(int) >> y
	case int8:
		return x.(int8) >> y
	case int16:
		return x.(int16) >> y
	case int32:
		return x.(int32) >> y
	case int64:
		return x.(int64) >> y
	case uint:
		return x.(uint) >> y
	case uint8:
		return x.(uint8) >> y
	case uint16:
		return x.(uint16) >> y
	case uint32:
		return x.(uint32) >> y
	case uint64:
		return x.(uint64) >> y
	case uintptr:
		return x.(uintptr) >> y
	default:
		vx := reflect.ValueOf(x)
		r := reflect.New(vx.Type()).Elem()
		switch vx.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			r.SetInt(vx.Int() >> y)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			r.SetUint(vx.Uint() >> y)
		default:
			goto failed
		}
		return r.Interface()
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T >> %T", x, y))
}

func opLSS(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) < y.(int)
	case int8:
		return x.(int8) < y.(int8)
	case int16:
		return x.(int16) < y.(int16)
	case int32:
		return x.(int32) < y.(int32)
	case int64:
		return x.(int64) < y.(int64)
	case uint:
		return x.(uint) < y.(uint)
	case uint8:
		return x.(uint8) < y.(uint8)
	case uint16:
		return x.(uint16) < y.(uint16)
	case uint32:
		return x.(uint32) < y.(uint32)
	case uint64:
		return x.(uint64) < y.(uint64)
	case uintptr:
		return x.(uintptr) < y.(uintptr)
	case float32:
		return x.(float32) < y.(float32)
	case float64:
		return x.(float64) < y.(float64)
	case string:
		return x.(string) < y.(string)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return vx.Int() < vy.Int()
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				return vx.Uint() < vy.Uint()
			case reflect.Float32, reflect.Float64:
				return vx.Float() < vy.Float()
			case reflect.String:
				return vx.String() < vy.String()
			default:
				goto failed
			}
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T < %T", x, y))
}

func opLEQ(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) <= y.(int)
	case int8:
		return x.(int8) <= y.(int8)
	case int16:
		return x.(int16) <= y.(int16)
	case int32:
		return x.(int32) <= y.(int32)
	case int64:
		return x.(int64) <= y.(int64)
	case uint:
		return x.(uint) <= y.(uint)
	case uint8:
		return x.(uint8) <= y.(uint8)
	case uint16:
		return x.(uint16) <= y.(uint16)
	case uint32:
		return x.(uint32) <= y.(uint32)
	case uint64:
		return x.(uint64) <= y.(uint64)
	case uintptr:
		return x.(uintptr) <= y.(uintptr)
	case float32:
		return x.(float32) <= y.(float32)
	case float64:
		return x.(float64) <= y.(float64)
	case string:
		return x.(string) <= y.(string)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return vx.Int() <= vy.Int()
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				return vx.Uint() <= vy.Uint()
			case reflect.Float32, reflect.Float64:
				return vx.Float() <= vy.Float()
			case reflect.String:
				return vx.String() <= vy.String()
			default:
				goto failed
			}
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T <= %T", x, y))
}

func opGTR(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) > y.(int)
	case int8:
		return x.(int8) > y.(int8)
	case int16:
		return x.(int16) > y.(int16)
	case int32:
		return x.(int32) > y.(int32)
	case int64:
		return x.(int64) > y.(int64)
	case uint:
		return x.(uint) > y.(uint)
	case uint8:
		return x.(uint8) > y.(uint8)
	case uint16:
		return x.(uint16) > y.(uint16)
	case uint32:
		return x.(uint32) > y.(uint32)
	case uint64:
		return x.(uint64) > y.(uint64)
	case uintptr:
		return x.(uintptr) > y.(uintptr)
	case float32:
		return x.(float32) > y.(float32)
	case float64:
		return x.(float64) > y.(float64)
	case string:
		return x.(string) > y.(string)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return vx.Int() > vy.Int()
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				return vx.Uint() > vy.Uint()
			case reflect.Float32, reflect.Float64:
				return vx.Float() > vy.Float()
			case reflect.String:
				return vx.String() > vy.String()
			default:
				goto failed
			}
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T > %T", x, y))
}

func opGEQ(x, y value) value {
	switch x.(type) {
	case int:
		return x.(int) >= y.(int)
	case int8:
		return x.(int8) >= y.(int8)
	case int16:
		return x.(int16) >= y.(int16)
	case int32:
		return x.(int32) >= y.(int32)
	case int64:
		return x.(int64) >= y.(int64)
	case uint:
		return x.(uint) >= y.(uint)
	case uint8:
		return x.(uint8) >= y.(uint8)
	case uint16:
		return x.(uint16) >= y.(uint16)
	case uint32:
		return x.(uint32) >= y.(uint32)
	case uint64:
		return x.(uint64) >= y.(uint64)
	case uintptr:
		return x.(uintptr) >= y.(uintptr)
	case float32:
		return x.(float32) >= y.(float32)
	case float64:
		return x.(float64) >= y.(float64)
	case string:
		return x.(string) >= y.(string)
	default:
		vx := reflect.ValueOf(x)
		vy := reflect.ValueOf(y)
		if kind := vx.Kind(); kind == vy.Kind() {
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return vx.Int() >= vy.Int()
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				return vx.Uint() >= vy.Uint()
			case reflect.Float32, reflect.Float64:
				return vx.Float() >= vy.Float()
			case reflect.String:
				return vx.String() >= vy.String()
			default:
				goto failed
			}
		}
	}
failed:
	panic(fmt.Sprintf("invalid binary op: %T >= %T", x, y))
}

// binop implements all arithmetic and logical binary operators for
// numeric datatypes and strings.  Both operands must have identical
// dynamic type.
func binop(instr *ssa.BinOp, t types.Type, x, y value) value {
	switch instr.Op {
	case token.ADD:
		return opADD(x, y)
	case token.SUB:
		return opSUB(x, y)
	case token.MUL:
		return opMUL(x, y)
	case token.QUO:
		return opQuo(x, y)
	case token.REM:
		return opREM(x, y)
	case token.AND:
		return opAND(x, y)
	case token.OR:
		return opOR(x, y)
	case token.XOR:
		return opXOR(x, y)
	case token.AND_NOT:
		return opANDNOT(x, y)
	case token.SHL:
		return opSHL(x, y)
	case token.SHR:
		return opSHR(x, y)
	case token.LSS:
		return opLSS(x, y)
	case token.LEQ:
		return opLEQ(x, y)
	case token.EQL:
		return opEQL(instr, x, y)
	case token.NEQ:
		return !opEQL(instr, x, y)
	case token.GTR:
		return opGTR(x, y)
	case token.GEQ:
		return opGEQ(x, y)
	}
	panic(fmt.Sprintf("invalid binary op: %T %s %T", x, instr.Op, y))
}

func IsConstNil(v ssa.Value) bool {
	switch c := v.(type) {
	case *ssa.Const:
		return c.IsNil()
	}
	return false
}

func IsNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Slice, reflect.Map, reflect.Func:
		return v.IsNil()
	case reflect.Chan, reflect.Ptr, reflect.UnsafePointer, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

func opEQL(instr *ssa.BinOp, x, y interface{}) bool {
	vx := reflect.ValueOf(x)
	vy := reflect.ValueOf(y)
	if vx.Kind() != vy.Kind() {
		return false
	}
	if IsConstNil(instr.X) {
		return IsNil(vy)
	} else if IsConstNil(instr.Y) {
		return IsNil(vx)
	}
	return equalValue(vx, vy)
}

func equalNil(vx, vy reflect.Value) bool {
	if IsNil(vx) {
		return IsNil(vy)
	} else if IsNil(vy) {
		return IsNil(vx)
	}
	return equalValue(vx, vy)
}

func equalValue(vx, vy reflect.Value) bool {
	if kind := vx.Kind(); kind == vy.Kind() {
		switch kind {
		case reflect.Invalid:
			return true
		case reflect.Chan:
			dirx := vx.Type().ChanDir()
			diry := vy.Type().ChanDir()
			if dirx != diry {
				if dirx == reflect.BothDir {
					return vy.Interface() == vx.Convert(vy.Type()).Interface()
				} else if diry == reflect.BothDir {
					return vx.Interface() == vy.Convert(vx.Type()).Interface()
				}
			} else {
				return vx.Interface() == vy.Interface()
			}
		case reflect.Ptr:
			return vx.Pointer() == vy.Pointer()
		case reflect.Struct:
			return equalStruct(vx, vy)
		case reflect.Array:
			return equalArray(vx, vy)
		default:
			return vx.Interface() == vy.Interface()
		}
	}
	return false
}

func equalArray(vx, vy reflect.Value) bool {
	xlen := vx.Len()
	if xlen != vy.Len() {
		return false
	}
	if vx.Type().Elem() != vy.Type().Elem() {
		return false
	}
	for i := 0; i < xlen; i++ {
		fx := vx.Index(i)
		fy := vy.Index(i)
		if !equalNil(fx, fy) {
			return false
		}
	}
	return true
}

func equalStruct(vx, vy reflect.Value) bool {
	typ := vx.Type()
	if typ != vy.Type() {
		return false
	}
	return vx.Interface() == vy.Interface()
}

func unop(instr *ssa.UnOp, x value) value {
	switch instr.Op {
	case token.ARROW: // receive
		vx := reflect.ValueOf(x)
		v, ok := vx.Recv()
		if !ok {
			v = reflect.New(vx.Type().Elem()).Elem()
		}
		if instr.CommaOk {
			return tuple{v.Interface(), ok}
		}
		return v.Interface()
		// if !ok {
		// 	v = zero(instr.X.Type().Underlying().(*types.Chan).Elem())
		// }
		// if instr.CommaOk {
		// 	v = tuple{v, ok}
		// }
		// return v
	case token.SUB:
		switch x := x.(type) {
		case int:
			return -x
		case int8:
			return -x
		case int16:
			return -x
		case int32:
			return -x
		case int64:
			return -x
		case uint:
			return -x
		case uint8:
			return -x
		case uint16:
			return -x
		case uint32:
			return -x
		case uint64:
			return -x
		case uintptr:
			return -x
		case float32:
			return -x
		case float64:
			return -x
		case complex64:
			return -x
		case complex128:
			return -x
		default:
			v := reflect.ValueOf(x)
			r := reflect.New(v.Type()).Elem()
			switch v.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(-v.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(-v.Uint())
			case reflect.Float32, reflect.Float64:
				r.SetFloat(-v.Float())
			case reflect.Complex64, reflect.Complex128:
				r.SetComplex(-v.Complex())
			}
			return r.Interface()
		}
	case token.MUL:
		v := reflect.ValueOf(x).Elem()
		if !v.IsValid() {
			panic(RuntimeError("invalid memory address or nil pointer dereference"))
		}
		return v.Interface()
		//return load(deref(instr.X.Type()), x.(*value))
	case token.NOT:
		switch x := x.(type) {
		case bool:
			return !x
		default:
			v := reflect.ValueOf(x)
			if v.Kind() == reflect.Bool {
				r := reflect.New(v.Type()).Elem()
				if v.Bool() {
					return v.Interface()
				}
				r.SetBool(true)
				return r.Interface()
			}
		}
		// return !x.(bool)
	case token.XOR:
		switch x := x.(type) {
		case int:
			return ^x
		case int8:
			return ^x
		case int16:
			return ^x
		case int32:
			return ^x
		case int64:
			return ^x
		case uint:
			return ^x
		case uint8:
			return ^x
		case uint16:
			return ^x
		case uint32:
			return ^x
		case uint64:
			return ^x
		case uintptr:
			return ^x
		default:
			vx := reflect.ValueOf(x)
			r := reflect.New(vx.Type()).Elem()
			switch vx.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				r.SetInt(^r.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				r.SetUint(^r.Uint())
			default:
				goto failed
			}
			return r.Interface()
		}
	}
failed:
	panic(fmt.Sprintf("invalid unary op %s %T", instr.Op, x))
}

// typeAssert checks whether dynamic type of itf is instr.AssertedType.
// It returns the extracted value on success, and panics on failure,
// unless instr.CommaOk, in which case it always returns a "value,ok" tuple.
func typeAssert(fr *frame, instr *ssa.TypeAssert, typ reflect.Type, xtyp reflect.Type, iv interface{}) value {
	var v value
	var err error
	if iv == nil {
		if xtyp == tyErrorInterface {
			err = fr.runtimeError(instr, "invalid memory address or nil pointer dereference")
		} else if xtyp == typ {
			err = fr.plainError(instr, fmt.Sprintf("interface conversion: interface is nil, not %v", typ))
		} else {
			err = fr.plainError(instr, fmt.Sprintf("interface conversion: %v is nil, not %v", xtyp, typ))
		}
		// }
	} else {
		rv := reflect.ValueOf(iv)
		rt := rv.Type()
		if typ == rt {
			v = iv
		} else {
			if !rt.AssignableTo(typ) {
				err = fr.runtimeError(instr, fmt.Sprintf("interface conversion: %v is %v, not %v", xtyp, rt, typ))
				if itype, ok := instr.AssertedType.Underlying().(*types.Interface); ok {
					if it, ok := fr.interp.findType(rt, false); ok {
						if meth, _ := types.MissingMethod(it, itype, true); meth != nil {
							err = fr.runtimeError(instr, fmt.Sprintf("interface conversion: %v is not %v: missing method %s",
								rt, instr.AssertedType, meth.Name()))
						}
					}
				} else if typ.PkgPath() == rt.PkgPath() && typ.Name() == rt.Name() {
					t1, ok1 := fr.interp.findType(typ, false)
					t2, ok2 := fr.interp.findType(rt, false)
					if ok1 && ok2 {
						n1, ok1 := t1.(*types.Named)
						n2, ok2 := t2.(*types.Named)
						if ok1 && ok2 && n1.Obj().Parent() != n2.Obj().Parent() {
							err = fr.runtimeError(instr, fmt.Sprintf("interface conversion: %v is %v, not %v (types from different scopes)", xtyp, rt, typ))
						}
					}
				}
			} else {
				v = rv.Convert(typ).Interface()
			}
		}
	}
	if err != nil {
		if !instr.CommaOk {
			panic(err)
		}
		return tuple{reflect.New(typ).Elem().Interface(), false}
	}
	if instr.CommaOk {
		return tuple{v, true}
	}
	return v
	// err := ""
	// if itf.t == nil {
	// 	err = fmt.Sprintf("interface conversion: interface is nil, not %s", instr.AssertedType)

	// } else if idst, ok := instr.AssertedType.Underlying().(*types.Interface); ok {
	// 	v = itf
	// 	err = checkInterface(i, idst, itf)

	// } else if types.Identical(itf.t, instr.AssertedType) {
	// 	v = itf.v // extract value

	// } else {
	// 	err = fmt.Sprintf("interface conversion: interface is %s, not %s", itf.t, instr.AssertedType)
	// }

	// if err != "" {
	// 	if !instr.CommaOk {
	// 		panic(err)
	// 	}
	// 	return tuple{zero(instr.AssertedType), false}
	// }
	// if instr.CommaOk {
	// 	return tuple{v, true}
	// }
	// return v
}

// widen widens a xtype typed value x to the widest type of its
// category, one of:
//
//	bool, int64, uint64, float64, complex128, string.
//
// This is inefficient but reduces the size of the cross-product of
// cases we have to consider.
func widen(x value) value {
	switch y := x.(type) {
	case bool, int64, uint64, float64, complex128, string, unsafe.Pointer:
		return x
	case int:
		return int64(y)
	case int8:
		return int64(y)
	case int16:
		return int64(y)
	case int32:
		return int64(y)
	case uint:
		return uint64(y)
	case uint8:
		return uint64(y)
	case uint16:
		return uint64(y)
	case uint32:
		return uint64(y)
	case uintptr:
		return uint64(y)
	case float32:
		return float64(y)
	case complex64:
		return complex128(y)
	}
	panic(fmt.Sprintf("cannot widen %T", x))
}

//go:nocheckptr
func toUnsafePointer(v uintptr) unsafe.Pointer {
	return unsafe.Pointer(v)
}

type reflectValue struct {
	typ  unsafe.Pointer
	ptr  unsafe.Pointer
	flag uintptr
}

func convert(x interface{}, typ reflect.Type) interface{} {
	v := reflect.ValueOf(x)
	vk := v.Kind()
	switch typ.Kind() {
	case reflect.UnsafePointer:
		if vk == reflect.Uintptr {
			return toUnsafePointer(uintptr(v.Uint()))
		} else if vk == reflect.Ptr {
			return unsafe.Pointer(v.Pointer())
		}
	case reflect.Uintptr:
		if vk == reflect.UnsafePointer {
			return v.Pointer()
		}
	case reflect.Ptr:
		if vk == reflect.UnsafePointer {
			return reflect.NewAt(typ.Elem(), unsafe.Pointer(v.Pointer())).Interface()
		}
	case reflect.Slice:
		if v.Kind() == reflect.String {
			elem := typ.Elem()
			switch elem.Kind() {
			case reflect.Uint8:
				if elem.PkgPath() != "" {
					dst := reflect.New(typ).Elem()
					dst.SetBytes([]byte(v.String()))
					return dst.Interface()
				}
			case reflect.Int32:
				if elem.PkgPath() != "" {
					dst := reflect.New(typ).Elem()
					*(*[]rune)((*reflectValue)(unsafe.Pointer(&dst)).ptr) = []rune(v.String())
					return dst.Interface()
				}
			}
		}
	case reflect.String:
		if v.Kind() == reflect.Slice {
			elem := v.Type().Elem()
			switch elem.Kind() {
			case reflect.Uint8:
				if elem.PkgPath() != "" {
					v = reflect.ValueOf(string(v.Bytes()))
				}
			case reflect.Int32:
				if elem.PkgPath() != "" {
					v = reflect.ValueOf(*(*[]rune)(((*reflectValue)(unsafe.Pointer(&v))).ptr))
				}
			}
		}
	}
	return v.Convert(typ).Interface()
}
