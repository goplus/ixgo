// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sccp

import (
	"go/constant"
	"go/token"
	"go/types"
	"math"
	"math/bits"

	"golang.org/x/tools/go/ssa"
)

// This file plays the role that rewriteValuegeneric plays for the compiler's
// sccp: given an operation whose arguments are all known constants, evaluate
// it at compile time with exact machine semantics (fixed-width two's
// complement integers, IEEE floats). Each fold function returns nil when the
// operation cannot or must not be folded — e.g. division by zero or negative
// shift count, whose run-time panics must be preserved — which computeLattice
// maps to the Bottom lattice.

func basicOf(t types.Type) (*types.Basic, bool) {
	b, ok := t.Underlying().(*types.Basic)
	return b, ok
}

// intInfo returns the width in bits and the signedness of an integer type.
// Untyped integer kinds are rejected: their machine width is unknown.
func intInfo(b *types.Basic) (w uint, signed bool, ok bool) {
	switch b.Kind() {
	case types.Int:
		return bits.UintSize, true, true
	case types.Int8:
		return 8, true, true
	case types.Int16:
		return 16, true, true
	case types.Int32:
		return 32, true, true
	case types.Int64:
		return 64, true, true
	case types.Uint:
		return bits.UintSize, false, true
	case types.Uint8:
		return 8, false, true
	case types.Uint16:
		return 16, false, true
	case types.Uint32:
		return 32, false, true
	case types.Uint64:
		return 64, false, true
	case types.Uintptr:
		return bits.UintSize, false, true
	}
	return 0, false, false
}

// foldableType reports whether constants of type t can be folded by this
// pass: booleans, fixed-width integers and floats. This mirrors the set of
// constant ops handled by the compiler's sccp (OpConst8..64, OpConstBool,
// OpConst32F/64F); strings, complex and nil constants are not supported.
func foldableType(t types.Type) bool {
	b, ok := basicOf(t)
	if !ok {
		return false
	}
	if b.Info()&types.IsBoolean != 0 {
		return true
	}
	if _, _, ok := intInfo(b); ok {
		return true
	}
	switch b.Kind() {
	case types.Float32, types.Float64:
		return true
	}
	return false
}

// foldableConvert reports whether a conversion from src to dst can be folded:
// both sides must be fixed-width numeric (integer or float) types. This is
// the analog of the compiler's Trunc*/ZeroExt*/SignExt*/Cvt* op set.
func foldableConvert(src, dst types.Type) bool {
	sb, ok := basicOf(src)
	if !ok {
		return false
	}
	db, ok := basicOf(dst)
	if !ok {
		return false
	}
	return numericKind(sb) && numericKind(db)
}

func numericKind(b *types.Basic) bool {
	if _, _, ok := intInfo(b); ok {
		return true
	}
	switch b.Kind() {
	case types.Float32, types.Float64:
		return true
	}
	return false
}

// truncS truncates v to a signed integer of w bits, sign-extending back
// into an int64.
func truncS(v int64, w uint) int64 {
	sh := 64 - w
	return v << sh >> sh
}

// truncU truncates v to an unsigned integer of w bits.
func truncU(v uint64, w uint) uint64 {
	if w == 64 {
		return v
	}
	return v & (1<<w - 1)
}

func float64Bits(f float64) uint64 {
	return math.Float64bits(f)
}

func boolConst(v bool, typ types.Type) *ssa.Const {
	return ssa.NewConst(constant.MakeBool(v), typ)
}

func intConst(v int64, w uint, typ types.Type) *ssa.Const {
	return ssa.NewConst(constant.MakeInt64(truncS(v, w)), typ)
}

func uintConst(v uint64, w uint, typ types.Type) *ssa.Const {
	return ssa.NewConst(constant.MakeUint64(truncU(v, w)), typ)
}

// floatConst returns a float constant for r, or nil if r is not exactly
// representable as a go/constant value: infinities and NaNs cannot be
// represented at all, and go/constant canonicalizes negative zero to
// positive zero which would change the sign observable via 1/x.
func floatConst(r float64, typ types.Type) *ssa.Const {
	if math.IsInf(r, 0) || math.IsNaN(r) || (r == 0 && math.Signbit(r)) {
		return nil
	}
	return ssa.NewConst(constant.MakeFloat64(r), typ)
}

// shiftCount extracts a shift count from a constant of any integer type.
// A negative count panics at run time, so it is not folded.
func shiftCount(y *ssa.Const) (uint64, bool) {
	b, ok := basicOf(y.Type())
	if !ok || b.Info()&types.IsInteger == 0 {
		return 0, false
	}
	if b.Info()&types.IsUnsigned != 0 {
		return y.Uint64(), true
	}
	v := y.Int64()
	if v < 0 {
		return 0, false
	}
	return uint64(v), true
}

// foldBinOp evaluates x op y, where typ is the static type of the result
// (the operand type for arithmetic, boolean for comparisons).
func foldBinOp(op token.Token, x, y *ssa.Const, typ types.Type) *ssa.Const {
	xb, ok := basicOf(x.Type())
	if !ok {
		return nil
	}

	// bool: only == and != exist
	if xb.Info()&types.IsBoolean != 0 {
		if x.Value.Kind() != constant.Bool || y.Value.Kind() != constant.Bool {
			return nil
		}
		a, b := constant.BoolVal(x.Value), constant.BoolVal(y.Value)
		switch op {
		case token.EQL:
			return boolConst(a == b, typ)
		case token.NEQ:
			return boolConst(a != b, typ)
		}
		return nil
	}

	// integer
	if w, signed, ok := intInfo(xb); ok {
		// shifts: the count may have any integer type and is handled apart
		if op == token.SHL || op == token.SHR {
			c, ok := shiftCount(y)
			if !ok {
				return nil
			}
			// Go shift semantics: counts >= width are well-defined (the
			// native int64/uint64 shift below already behaves that way)
			if signed {
				a := x.Int64()
				if op == token.SHL {
					return intConst(a<<c, w, typ)
				}
				return intConst(a>>c, w, typ)
			}
			u := x.Uint64()
			if op == token.SHL {
				return uintConst(u<<c, w, typ)
			}
			return uintConst(u>>c, w, typ)
		}
		if signed {
			a, b := x.Int64(), y.Int64()
			switch op {
			case token.ADD:
				return intConst(a+b, w, typ)
			case token.SUB:
				return intConst(a-b, w, typ)
			case token.MUL:
				return intConst(a*b, w, typ)
			case token.QUO:
				if b == 0 {
					return nil // division by zero panics at run time
				}
				return intConst(a/b, w, typ)
			case token.REM:
				if b == 0 {
					return nil
				}
				return intConst(a%b, w, typ)
			case token.AND:
				return intConst(a&b, w, typ)
			case token.OR:
				return intConst(a|b, w, typ)
			case token.XOR:
				return intConst(a^b, w, typ)
			case token.AND_NOT:
				return intConst(a&^b, w, typ)
			case token.EQL:
				return boolConst(a == b, typ)
			case token.NEQ:
				return boolConst(a != b, typ)
			case token.LSS:
				return boolConst(a < b, typ)
			case token.LEQ:
				return boolConst(a <= b, typ)
			case token.GTR:
				return boolConst(a > b, typ)
			case token.GEQ:
				return boolConst(a >= b, typ)
			}
			return nil
		}
		a, b := x.Uint64(), y.Uint64()
		switch op {
		case token.ADD:
			return uintConst(a+b, w, typ)
		case token.SUB:
			return uintConst(a-b, w, typ)
		case token.MUL:
			return uintConst(a*b, w, typ)
		case token.QUO:
			if b == 0 {
				return nil
			}
			return uintConst(a/b, w, typ)
		case token.REM:
			if b == 0 {
				return nil
			}
			return uintConst(a%b, w, typ)
		case token.AND:
			return uintConst(a&b, w, typ)
		case token.OR:
			return uintConst(a|b, w, typ)
		case token.XOR:
			return uintConst(a^b, w, typ)
		case token.AND_NOT:
			return uintConst(a&^b, w, typ)
		case token.EQL:
			return boolConst(a == b, typ)
		case token.NEQ:
			return boolConst(a != b, typ)
		case token.LSS:
			return boolConst(a < b, typ)
		case token.LEQ:
			return boolConst(a <= b, typ)
		case token.GTR:
			return boolConst(a > b, typ)
		case token.GEQ:
			return boolConst(a >= b, typ)
		}
		return nil
	}

	// float
	switch xb.Kind() {
	case types.Float32:
		a, b := float32(x.Float64()), float32(y.Float64())
		switch op {
		case token.ADD:
			return floatConst(float64(a+b), typ)
		case token.SUB:
			return floatConst(float64(a-b), typ)
		case token.MUL:
			return floatConst(float64(a*b), typ)
		case token.QUO:
			return floatConst(float64(a/b), typ) // x/0 yields Inf, rejected by floatConst
		case token.EQL:
			return boolConst(a == b, typ)
		case token.NEQ:
			return boolConst(a != b, typ)
		case token.LSS:
			return boolConst(a < b, typ)
		case token.LEQ:
			return boolConst(a <= b, typ)
		case token.GTR:
			return boolConst(a > b, typ)
		case token.GEQ:
			return boolConst(a >= b, typ)
		}
	case types.Float64:
		a, b := x.Float64(), y.Float64()
		switch op {
		case token.ADD:
			return floatConst(a+b, typ)
		case token.SUB:
			return floatConst(a-b, typ)
		case token.MUL:
			return floatConst(a*b, typ)
		case token.QUO:
			return floatConst(a/b, typ)
		case token.EQL:
			return boolConst(a == b, typ)
		case token.NEQ:
			return boolConst(a != b, typ)
		case token.LSS:
			return boolConst(a < b, typ)
		case token.LEQ:
			return boolConst(a <= b, typ)
		case token.GTR:
			return boolConst(a > b, typ)
		case token.GEQ:
			return boolConst(a >= b, typ)
		}
	}
	return nil
}

// foldUnOp evaluates op x, where typ is the static type of the result.
func foldUnOp(op token.Token, x *ssa.Const, typ types.Type) *ssa.Const {
	xb, ok := basicOf(x.Type())
	if !ok {
		return nil
	}
	switch op {
	case token.SUB: // negate
		if w, signed, ok := intInfo(xb); ok {
			if signed {
				return intConst(-x.Int64(), w, typ)
			}
			return uintConst(-x.Uint64(), w, typ)
		}
		switch xb.Kind() {
		case types.Float32:
			return floatConst(float64(-float32(x.Float64())), typ)
		case types.Float64:
			return floatConst(-x.Float64(), typ)
		}
	case token.XOR: // bitwise complement
		if w, signed, ok := intInfo(xb); ok {
			if signed {
				return intConst(^x.Int64(), w, typ)
			}
			return uintConst(^x.Uint64(), w, typ)
		}
	case token.NOT: // not
		if xb.Info()&types.IsBoolean != 0 && x.Value.Kind() == constant.Bool {
			return boolConst(!constant.BoolVal(x.Value), typ)
		}
	}
	return nil
}

// foldConvert evaluates a numeric conversion of x to typ.
func foldConvert(x *ssa.Const, typ types.Type) *ssa.Const {
	sb, ok := basicOf(x.Type())
	if !ok {
		return nil
	}
	db, ok := basicOf(typ)
	if !ok {
		return nil
	}

	// integer source
	if _, ssigned, sok := intInfo(sb); sok {
		if dw, dsigned, dok := intInfo(db); dok {
			// int -> int: sign- or zero-extend, then truncate to the
			// destination width (Trunc*/ZeroExt*/SignExt*)
			var raw uint64
			if ssigned {
				raw = uint64(x.Int64())
			} else {
				raw = x.Uint64()
			}
			if dsigned {
				return intConst(int64(raw), dw, typ)
			}
			return uintConst(raw, dw, typ)
		}
		// int -> float (Cvt*to*F)
		var f float64
		if ssigned {
			f = float64(x.Int64())
		} else {
			f = float64(x.Uint64())
		}
		switch db.Kind() {
		case types.Float32:
			return floatConst(float64(float32(f)), typ)
		case types.Float64:
			return floatConst(f, typ)
		}
		return nil
	}

	// float source
	switch sb.Kind() {
	case types.Float32, types.Float64:
	default:
		return nil
	}
	f := x.Float64()
	switch db.Kind() {
	case types.Float32:
		return floatConst(float64(float32(f)), typ)
	case types.Float64:
		return floatConst(f, typ)
	}
	if dw, dsigned, dok := intInfo(db); dok {
		// float -> int (Cvt*Fto*): the result of an out-of-range conversion
		// is implementation-defined at run time, so only fold in-range values
		v := math.Trunc(f)
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return nil
		}
		if dsigned {
			lim := math.Ldexp(1, int(dw)-1) // 2^(w-1), exact in float64
			if !(v >= -lim && v < lim) {
				return nil
			}
			return intConst(int64(v), dw, typ)
		}
		lim := math.Ldexp(1, int(dw)) // 2^w
		if !(v >= 0 && v < lim) {
			return nil
		}
		return uintConst(uint64(v), dw, typ)
	}
	return nil
}
