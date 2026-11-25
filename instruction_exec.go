package ixgo

import (
	"github.com/visualfc/xtype"
	"go/token"
	"reflect"
)

// makeInstruction builds supported closures solely from execution metadata.
// Unsupported instructions retain their existing lowering.
func makeInstruction(d typedInstruction) (func(*frame), bool) {
	switch d.Op {
	case opJump:
		target := d.Aux.(jumpAux).Target
		return func(fr *frame) { fr.jump(target) }, true
	case opIf:
		aux := d.Aux.(ifAux)
		if aux.Constant {
			if aux.Condition {
				return func(fr *frame) { fr.jump(aux.Targets[0]) }, true
			}
			return func(fr *frame) { fr.jump(aux.Targets[1]) }, true
		}
		condition := d.Args[0]
		if aux.BasicBool {
			return func(fr *frame) {
				if fr.reg(condition).(bool) {
					fr.jump(aux.Targets[0])
				} else {
					fr.jump(aux.Targets[1])
				}
			}, true
		}
		return func(fr *frame) {
			if fr.bool(condition) {
				fr.jump(aux.Targets[0])
			} else {
				fr.jump(aux.Targets[1])
			}
		}, true
	case opPhi:
		aux := d.Aux.(phiAux)
		ir, ie, preds := aux.Results, aux.Edges, aux.Preds
		return func(fr *frame) {
			predIndex := -1
			for i, pred := range preds {
				if fr.pred == pred {
					predIndex = i
					break
				}
			}
			if predIndex < 0 {
				return
			}
			if cap(fr.phiVals) < len(ir) {
				fr.phiVals = make([]value, len(ir))
			}
			// Read all edges before writing results to preserve cyclic assignments.
			values := fr.phiVals[:len(ir)]
			for i := range ir {
				values[i] = fr.reg(ie[i][predIndex])
			}
			for i := range ir {
				fr.setReg(ir[i], values[i])
			}
		}, true
	case opBinOp:
		return makeBinaryInstruction(d)
	case opUnOp:
		return makeUnaryInstruction(d)
	}
	return nil, false
}

// makeBinaryInstruction builds binary operations from runtime descriptions.
func makeBinaryInstruction(d typedInstruction) (func(*frame), bool) {
	aux := d.Aux.(binaryOpAux)
	ir, ix, iy := d.Result, d.Args[0], d.Args[1]
	kx, ky := kindInvalid, kindInvalid
	if aux.Constant[0] {
		kx = kindConst
	}
	if aux.Constant[1] {
		ky = kindConst
	}
	vx, vy := aux.Values[0], aux.Values[1]
	switch aux.Op {
	case token.ADD:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) + vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)+fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) + vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)+fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) + vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)+fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) + vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)+fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) + vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)+fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) + vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)+fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) + vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)+fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) + vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)+fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) + vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)+fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) + vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)+fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) + vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)+fr.reg(iy).(uintptr)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := vx.(float32) + vy.(float32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float32)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(float32)) }, true
				} else if ky == kindConst {
					y := vy.(float32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)+fr.reg(iy).(float32)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := vx.(float64) + vy.(float64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float64)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(float64)) }, true
				} else if ky == kindConst {
					y := vy.(float64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)+fr.reg(iy).(float64)) }, true
				}
			case reflect.Complex64:
				if kx == kindConst && ky == kindConst {
					v := vx.(complex64) + vy.(complex64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(complex64)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(complex64)) }, true
				} else if ky == kindConst {
					y := vy.(complex64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)+fr.reg(iy).(complex64)) }, true
				}
			case reflect.Complex128:
				if kx == kindConst && ky == kindConst {
					v := vx.(complex128) + vy.(complex128)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(complex128)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(complex128)) }, true
				} else if ky == kindConst {
					y := vy.(complex128)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)+fr.reg(iy).(complex128)) }, true
				}
			case reflect.String:
				if kx == kindConst && ky == kindConst {
					v := vx.(string) + vy.(string)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(string)
					return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(string)) }, true
				} else if ky == kindConst {
					y := vy.(string)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string)+y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string)+fr.reg(iy).(string)) }, true
				}
			}
		} else {
			t := xtype.TypeOfType(typ)
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int(vx)+xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.int(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)+fr.int(iy))) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int8(vx)+xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.int8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)+fr.int8(iy))) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int16(vx)+xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.int16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)+fr.int16(iy))) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int32(vx)+xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.int32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)+fr.int32(iy))) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int64(vx)+xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.int64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)+fr.int64(iy))) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint(vx)+xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.uint(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)+fr.uint(iy))) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint8(vx)+xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.uint8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)+fr.uint8(iy))) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint16(vx)+xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.uint16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)+fr.uint16(iy))) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint32(vx)+xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.uint32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)+fr.uint32(iy))) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint64(vx)+xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.uint64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)+fr.uint64(iy))) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uintptr(vx)+xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.uintptr(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)+fr.uintptr(iy))) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Float32(vx)+xtype.Float32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.float32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float32(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float32(ix)+fr.float32(iy))) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Float64(vx)+xtype.Float64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.float64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float64(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float64(ix)+fr.float64(iy))) }, true
				}
			case reflect.Complex64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Complex64(vx)+xtype.Complex64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Complex64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.complex64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Complex64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex64(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex64(ix)+fr.complex64(iy))) }, true
				}
			case reflect.Complex128:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Complex128(vx)+xtype.Complex128(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Complex128(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.complex128(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Complex128(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex128(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex128(ix)+fr.complex128(iy))) }, true
				}
			case reflect.String:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.String(vx)+xtype.String(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.String(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x+fr.string(iy))) }, true
				} else if ky == kindConst {
					y := xtype.String(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.string(ix)+y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.string(ix)+fr.string(iy))) }, true
				}
			}
		}
		panic("unreachable")
	case token.SUB:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) - vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)-fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) - vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)-fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) - vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)-fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) - vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)-fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) - vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)-fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) - vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)-fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) - vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)-fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) - vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)-fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) - vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)-fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) - vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)-fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) - vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)-fr.reg(iy).(uintptr)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := vx.(float32) - vy.(float32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float32)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(float32)) }, true
				} else if ky == kindConst {
					y := vy.(float32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)-fr.reg(iy).(float32)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := vx.(float64) - vy.(float64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float64)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(float64)) }, true
				} else if ky == kindConst {
					y := vy.(float64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)-fr.reg(iy).(float64)) }, true
				}
			case reflect.Complex64:
				if kx == kindConst && ky == kindConst {
					v := vx.(complex64) - vy.(complex64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(complex64)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(complex64)) }, true
				} else if ky == kindConst {
					y := vy.(complex64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)-fr.reg(iy).(complex64)) }, true
				}
			case reflect.Complex128:
				if kx == kindConst && ky == kindConst {
					v := vx.(complex128) - vy.(complex128)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(complex128)
					return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(complex128)) }, true
				} else if ky == kindConst {
					y := vy.(complex128)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)-y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)-fr.reg(iy).(complex128)) }, true
				}
			}
		} else {
			t := xtype.TypeOfType(typ)
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int(vx)-xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.int(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)-fr.int(iy))) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int8(vx)-xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.int8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)-fr.int8(iy))) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int16(vx)-xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.int16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)-fr.int16(iy))) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int32(vx)-xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.int32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)-fr.int32(iy))) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int64(vx)-xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.int64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)-fr.int64(iy))) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint(vx)-xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.uint(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)-fr.uint(iy))) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint8(vx)-xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.uint8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)-fr.uint8(iy))) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint16(vx)-xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.uint16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)-fr.uint16(iy))) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint32(vx)-xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.uint32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)-fr.uint32(iy))) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint64(vx)-xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.uint64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)-fr.uint64(iy))) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uintptr(vx)-xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.uintptr(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)-fr.uintptr(iy))) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Float32(vx)-xtype.Float32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.float32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float32(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float32(ix)-fr.float32(iy))) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Float64(vx)-xtype.Float64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.float64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float64(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float64(ix)-fr.float64(iy))) }, true
				}
			case reflect.Complex64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Complex64(vx)-xtype.Complex64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Complex64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.complex64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Complex64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex64(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex64(ix)-fr.complex64(iy))) }, true
				}
			case reflect.Complex128:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Complex128(vx)-xtype.Complex128(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Complex128(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x-fr.complex128(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Complex128(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex128(ix)-y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex128(ix)-fr.complex128(iy))) }, true
				}
			}
		}
		panic("unreachable")
	case token.MUL:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) * vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)*fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) * vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)*fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) * vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)*fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) * vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)*fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) * vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)*fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) * vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)*fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) * vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)*fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) * vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)*fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) * vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)*fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) * vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)*fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) * vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)*fr.reg(iy).(uintptr)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := vx.(float32) * vy.(float32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float32)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(float32)) }, true
				} else if ky == kindConst {
					y := vy.(float32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)*fr.reg(iy).(float32)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := vx.(float64) * vy.(float64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float64)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(float64)) }, true
				} else if ky == kindConst {
					y := vy.(float64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)*fr.reg(iy).(float64)) }, true
				}
			case reflect.Complex64:
				if kx == kindConst && ky == kindConst {
					v := vx.(complex64) * vy.(complex64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(complex64)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(complex64)) }, true
				} else if ky == kindConst {
					y := vy.(complex64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)*fr.reg(iy).(complex64)) }, true
				}
			case reflect.Complex128:
				if kx == kindConst && ky == kindConst {
					v := vx.(complex128) * vy.(complex128)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(complex128)
					return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(complex128)) }, true
				} else if ky == kindConst {
					y := vy.(complex128)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)*y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)*fr.reg(iy).(complex128)) }, true
				}
			}
		} else {
			t := xtype.TypeOfType(typ)
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int(vx)*xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.int(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)*fr.int(iy))) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int8(vx)*xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.int8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)*fr.int8(iy))) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int16(vx)*xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.int16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)*fr.int16(iy))) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int32(vx)*xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.int32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)*fr.int32(iy))) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int64(vx)*xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.int64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)*fr.int64(iy))) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint(vx)*xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.uint(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)*fr.uint(iy))) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint8(vx)*xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.uint8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)*fr.uint8(iy))) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint16(vx)*xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.uint16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)*fr.uint16(iy))) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint32(vx)*xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.uint32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)*fr.uint32(iy))) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint64(vx)*xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.uint64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)*fr.uint64(iy))) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uintptr(vx)*xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.uintptr(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)*fr.uintptr(iy))) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Float32(vx)*xtype.Float32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.float32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float32(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float32(ix)*fr.float32(iy))) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Float64(vx)*xtype.Float64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.float64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float64(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float64(ix)*fr.float64(iy))) }, true
				}
			case reflect.Complex64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Complex64(vx)*xtype.Complex64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Complex64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.complex64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Complex64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex64(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex64(ix)*fr.complex64(iy))) }, true
				}
			case reflect.Complex128:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Complex128(vx)*xtype.Complex128(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Complex128(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x*fr.complex128(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Complex128(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex128(ix)*y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex128(ix)*fr.complex128(iy))) }, true
				}
			}
		}
		panic("unreachable")
	case token.QUO:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					x := vx.(int)
					y := vy.(int)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)/fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					x := vx.(int8)
					y := vy.(int8)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)/fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					x := vx.(int16)
					y := vy.(int16)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)/fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					x := vx.(int32)
					y := vy.(int32)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)/fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					x := vx.(int64)
					y := vy.(int64)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)/fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					x := vx.(uint)
					y := vy.(uint)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)/fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					x := vx.(uint8)
					y := vy.(uint8)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)/fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					x := vx.(uint16)
					y := vy.(uint16)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)/fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					x := vx.(uint32)
					y := vy.(uint32)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)/fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					x := vx.(uint64)
					y := vy.(uint64)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)/fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					x := vx.(uintptr)
					y := vy.(uintptr)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, x/y) }, true
					}
					v := x / y
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)/fr.reg(iy).(uintptr)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := vx.(float32) / vy.(float32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float32)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(float32)) }, true
				} else if ky == kindConst {
					y := vy.(float32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)/fr.reg(iy).(float32)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := vx.(float64) / vy.(float64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float64)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(float64)) }, true
				} else if ky == kindConst {
					y := vy.(float64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)/fr.reg(iy).(float64)) }, true
				}
			case reflect.Complex64:
				if kx == kindConst && ky == kindConst {
					v := vx.(complex64) / vy.(complex64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(complex64)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(complex64)) }, true
				} else if ky == kindConst {
					y := vy.(complex64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)/fr.reg(iy).(complex64)) }, true
				}
			case reflect.Complex128:
				if kx == kindConst && ky == kindConst {
					v := vx.(complex128) / vy.(complex128)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(complex128)
					return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(complex128)) }, true
				} else if ky == kindConst {
					y := vy.(complex128)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)/y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)/fr.reg(iy).(complex128)) }, true
				}
			}
		} else {
			t := xtype.TypeOfType(typ)
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					x := xtype.Int(vx)
					y := xtype.Int(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.int(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)/fr.int(iy))) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					x := xtype.Int8(vx)
					y := xtype.Int8(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.int8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)/fr.int8(iy))) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					x := xtype.Int16(vx)
					y := xtype.Int16(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.int16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)/fr.int16(iy))) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					x := xtype.Int32(vx)
					y := xtype.Int32(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.int32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)/fr.int32(iy))) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					x := xtype.Int64(vx)
					y := xtype.Int64(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.int64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)/fr.int64(iy))) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					x := xtype.Uint(vx)
					y := xtype.Uint(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.uint(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)/fr.uint(iy))) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					x := xtype.Uint8(vx)
					y := xtype.Uint8(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.uint8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)/fr.uint8(iy))) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					x := xtype.Uint16(vx)
					y := xtype.Uint16(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.uint16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)/fr.uint16(iy))) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					x := xtype.Uint32(vx)
					y := xtype.Uint32(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.uint32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)/fr.uint32(iy))) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					x := xtype.Uint64(vx)
					y := xtype.Uint64(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.uint64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)/fr.uint64(iy))) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					x := xtype.Uintptr(vx)
					y := xtype.Uintptr(vy)
					if y == 0 {
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/y)) }, true
					}
					v := xtype.Make(t, x/y)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.uintptr(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)/fr.uintptr(iy))) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Float32(vx)/xtype.Float32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.float32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float32(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float32(ix)/fr.float32(iy))) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Float64(vx)/xtype.Float64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.float64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float64(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.float64(ix)/fr.float64(iy))) }, true
				}
			case reflect.Complex64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Complex64(vx)/xtype.Complex64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Complex64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.complex64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Complex64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex64(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex64(ix)/fr.complex64(iy))) }, true
				}
			case reflect.Complex128:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Complex128(vx)/xtype.Complex128(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Complex128(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x/fr.complex128(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Complex128(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex128(ix)/y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.complex128(ix)/fr.complex128(iy))) }, true
				}
			}
		}
		panic("unreachable")
	case token.REM:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) % vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)%fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) % vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)%fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) % vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)%fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) % vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)%fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) % vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)%fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) % vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)%fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) % vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)%fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) % vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)%fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) % vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)%fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) % vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)%fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) % vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)%y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)%fr.reg(iy).(uintptr)) }, true
				}
			}
		} else {
			t := xtype.TypeOfType(typ)
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int(vx)%xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.int(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)%fr.int(iy))) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int8(vx)%xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.int8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)%fr.int8(iy))) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int16(vx)%xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.int16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)%fr.int16(iy))) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int32(vx)%xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.int32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)%fr.int32(iy))) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int64(vx)%xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.int64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)%fr.int64(iy))) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint(vx)%xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.uint(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)%fr.uint(iy))) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint8(vx)%xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.uint8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)%fr.uint8(iy))) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint16(vx)%xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.uint16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)%fr.uint16(iy))) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint32(vx)%xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.uint32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)%fr.uint32(iy))) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint64(vx)%xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.uint64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)%fr.uint64(iy))) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uintptr(vx)%xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x%fr.uintptr(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)%y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)%fr.uintptr(iy))) }, true
				}
			}
		}
		panic("unreachable")
	case token.AND:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) & vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)&fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) & vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)&fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) & vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)&fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) & vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)&fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) & vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)&fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) & vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)&fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) & vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)&fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) & vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)&fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) & vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)&fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) & vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)&fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) & vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)&y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)&fr.reg(iy).(uintptr)) }, true
				}
			}
		} else {
			t := xtype.TypeOfType(typ)
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int(vx)&xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.int(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)&fr.int(iy))) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int8(vx)&xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.int8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)&fr.int8(iy))) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int16(vx)&xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.int16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)&fr.int16(iy))) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int32(vx)&xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.int32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)&fr.int32(iy))) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int64(vx)&xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.int64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)&fr.int64(iy))) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint(vx)&xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.uint(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)&fr.uint(iy))) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint8(vx)&xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.uint8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)&fr.uint8(iy))) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint16(vx)&xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.uint16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)&fr.uint16(iy))) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint32(vx)&xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.uint32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)&fr.uint32(iy))) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint64(vx)&xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.uint64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)&fr.uint64(iy))) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uintptr(vx)&xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&fr.uintptr(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)&y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)&fr.uintptr(iy))) }, true
				}
			}
		}
		panic("unreachable")
	case token.OR:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) | vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)|fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) | vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)|fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) | vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)|fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) | vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)|fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) | vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)|fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) | vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)|fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) | vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)|fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) | vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)|fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) | vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)|fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) | vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)|fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) | vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)|y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)|fr.reg(iy).(uintptr)) }, true
				}
			}
		} else {
			t := xtype.TypeOfType(typ)
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int(vx)|xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.int(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)|fr.int(iy))) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int8(vx)|xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.int8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)|fr.int8(iy))) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int16(vx)|xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.int16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)|fr.int16(iy))) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int32(vx)|xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.int32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)|fr.int32(iy))) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int64(vx)|xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.int64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)|fr.int64(iy))) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint(vx)|xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.uint(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)|fr.uint(iy))) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint8(vx)|xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.uint8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)|fr.uint8(iy))) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint16(vx)|xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.uint16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)|fr.uint16(iy))) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint32(vx)|xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.uint32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)|fr.uint32(iy))) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint64(vx)|xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.uint64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)|fr.uint64(iy))) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uintptr(vx)|xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x|fr.uintptr(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)|y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)|fr.uintptr(iy))) }, true
				}
			}
		}
		panic("unreachable")
	case token.XOR:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) ^ vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)^fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) ^ vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)^fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) ^ vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)^fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) ^ vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)^fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) ^ vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)^fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) ^ vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)^fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) ^ vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)^fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) ^ vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)^fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) ^ vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)^fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) ^ vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)^fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) ^ vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)^fr.reg(iy).(uintptr)) }, true
				}
			}
		} else {
			t := xtype.TypeOfType(typ)
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int(vx)^xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.int(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)^fr.int(iy))) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int8(vx)^xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.int8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)^fr.int8(iy))) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int16(vx)^xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.int16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)^fr.int16(iy))) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int32(vx)^xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.int32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)^fr.int32(iy))) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int64(vx)^xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.int64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)^fr.int64(iy))) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint(vx)^xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.uint(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)^fr.uint(iy))) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint8(vx)^xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.uint8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)^fr.uint8(iy))) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint16(vx)^xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.uint16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)^fr.uint16(iy))) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint32(vx)^xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.uint32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)^fr.uint32(iy))) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint64(vx)^xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.uint64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)^fr.uint64(iy))) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uintptr(vx)^xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x^fr.uintptr(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)^fr.uintptr(iy))) }, true
				}
			}
		}
		panic("unreachable")
	case token.AND_NOT:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) &^ vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)&^fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) &^ vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)&^fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) &^ vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)&^fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) &^ vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)&^fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) &^ vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)&^fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) &^ vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)&^fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) &^ vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)&^fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) &^ vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)&^fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) &^ vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)&^fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) &^ vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)&^fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) &^ vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)&^y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)&^fr.reg(iy).(uintptr)) }, true
				}
			}
		} else {
			t := xtype.TypeOfType(typ)
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int(vx)&^xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.int(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)&^fr.int(iy))) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int8(vx)&^xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.int8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)&^fr.int8(iy))) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int16(vx)&^xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.int16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)&^fr.int16(iy))) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int32(vx)&^xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.int32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)&^fr.int32(iy))) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Int64(vx)&^xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.int64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)&^fr.int64(iy))) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint(vx)&^xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.uint(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)&^fr.uint(iy))) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint8(vx)&^xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.uint8(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)&^fr.uint8(iy))) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint16(vx)&^xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.uint16(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)&^fr.uint16(iy))) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint32(vx)&^xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.uint32(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)&^fr.uint32(iy))) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uint64(vx)&^xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.uint64(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)&^fr.uint64(iy))) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Make(t, xtype.Uintptr(vx)&^xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x&^fr.uintptr(iy))) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)&^y)) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)&^fr.uintptr(iy))) }, true
				}
			}
		}
		panic("unreachable")
	case token.LSS:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) < vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) < fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) < vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) < fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) < vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) < fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) < vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) < fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) < vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) < fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) < vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) < fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) < vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) < fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) < vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) < fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) < vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) < fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) < vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) < fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) < vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) < fr.reg(iy).(uintptr)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := vx.(float32) < vy.(float32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float32)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(float32)) }, true
				} else if ky == kindConst {
					y := vy.(float32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) < fr.reg(iy).(float32)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := vx.(float64) < vy.(float64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float64)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(float64)) }, true
				} else if ky == kindConst {
					y := vy.(float64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) < fr.reg(iy).(float64)) }, true
				}
			case reflect.String:
				if kx == kindConst && ky == kindConst {
					v := vx.(string) < vy.(string)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(string)
					return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(string)) }, true
				} else if ky == kindConst {
					y := vy.(string)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) < fr.reg(iy).(string)) }, true
				}
			}
		} else {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int(vx) < xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.int(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int(ix) < fr.int(iy)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int8(vx) < xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.int8(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int8(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int8(ix) < fr.int8(iy)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int16(vx) < xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.int16(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int16(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int16(ix) < fr.int16(iy)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int32(vx) < xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.int32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int32(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int32(ix) < fr.int32(iy)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int64(vx) < xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.int64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int64(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int64(ix) < fr.int64(iy)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint(vx) < xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.uint(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint(ix) < fr.uint(iy)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint8(vx) < xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.uint8(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) < fr.uint8(iy)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint16(vx) < xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.uint16(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) < fr.uint16(iy)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint32(vx) < xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.uint32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) < fr.uint32(iy)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint64(vx) < xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.uint64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) < fr.uint64(iy)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uintptr(vx) < xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.uintptr(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) < fr.uintptr(iy)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Float32(vx) < xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float32(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.float32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.float32(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.float32(ix) < fr.float32(iy)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Float64(vx) < xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float64(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.float64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.float64(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.float64(ix) < fr.float64(iy)) }, true
				}
			case reflect.String:
				if kx == kindConst && ky == kindConst {
					v := xtype.String(vx) < xtype.String(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.String(vx)
					return func(fr *frame) { fr.setReg(ir, x < fr.string(iy)) }, true
				} else if ky == kindConst {
					y := xtype.String(vy)
					return func(fr *frame) { fr.setReg(ir, fr.string(ix) < y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.string(ix) < fr.string(iy)) }, true
				}
			}
		}
		panic("unreachable")
	case token.LEQ:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) <= vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) <= fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) <= vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) <= fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) <= vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) <= fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) <= vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) <= fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) <= vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) <= fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) <= vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) <= fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) <= vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) <= fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) <= vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) <= fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) <= vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) <= fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) <= vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) <= fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) <= vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) <= fr.reg(iy).(uintptr)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := vx.(float32) <= vy.(float32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float32)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(float32)) }, true
				} else if ky == kindConst {
					y := vy.(float32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) <= fr.reg(iy).(float32)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := vx.(float64) <= vy.(float64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float64)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(float64)) }, true
				} else if ky == kindConst {
					y := vy.(float64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) <= fr.reg(iy).(float64)) }, true
				}
			case reflect.String:
				if kx == kindConst && ky == kindConst {
					v := vx.(string) <= vy.(string)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(string)
					return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(string)) }, true
				} else if ky == kindConst {
					y := vy.(string)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) <= fr.reg(iy).(string)) }, true
				}
			}
		} else {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int(vx) <= xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.int(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int(ix) <= fr.int(iy)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int8(vx) <= xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.int8(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int8(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int8(ix) <= fr.int8(iy)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int16(vx) <= xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.int16(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int16(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int16(ix) <= fr.int16(iy)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int32(vx) <= xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.int32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int32(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int32(ix) <= fr.int32(iy)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int64(vx) <= xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.int64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int64(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int64(ix) <= fr.int64(iy)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint(vx) <= xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.uint(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint(ix) <= fr.uint(iy)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint8(vx) <= xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.uint8(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) <= fr.uint8(iy)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint16(vx) <= xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.uint16(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) <= fr.uint16(iy)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint32(vx) <= xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.uint32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) <= fr.uint32(iy)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint64(vx) <= xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.uint64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) <= fr.uint64(iy)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uintptr(vx) <= xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.uintptr(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) <= fr.uintptr(iy)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Float32(vx) <= xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float32(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.float32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.float32(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.float32(ix) <= fr.float32(iy)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Float64(vx) <= xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float64(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.float64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.float64(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.float64(ix) <= fr.float64(iy)) }, true
				}
			case reflect.String:
				if kx == kindConst && ky == kindConst {
					v := xtype.String(vx) <= xtype.String(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.String(vx)
					return func(fr *frame) { fr.setReg(ir, x <= fr.string(iy)) }, true
				} else if ky == kindConst {
					y := xtype.String(vy)
					return func(fr *frame) { fr.setReg(ir, fr.string(ix) <= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.string(ix) <= fr.string(iy)) }, true
				}
			}
		}
		panic("unreachable")
	case token.GTR:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) > vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) > fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) > vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) > fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) > vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) > fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) > vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) > fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) > vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) > fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) > vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) > fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) > vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) > fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) > vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) > fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) > vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) > fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) > vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) > fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) > vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) > fr.reg(iy).(uintptr)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := vx.(float32) > vy.(float32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float32)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(float32)) }, true
				} else if ky == kindConst {
					y := vy.(float32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) > fr.reg(iy).(float32)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := vx.(float64) > vy.(float64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float64)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(float64)) }, true
				} else if ky == kindConst {
					y := vy.(float64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) > fr.reg(iy).(float64)) }, true
				}
			case reflect.String:
				if kx == kindConst && ky == kindConst {
					v := vx.(string) > vy.(string)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(string)
					return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(string)) }, true
				} else if ky == kindConst {
					y := vy.(string)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) > fr.reg(iy).(string)) }, true
				}
			}
		} else {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int(vx) > xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.int(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int(ix) > fr.int(iy)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int8(vx) > xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.int8(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int8(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int8(ix) > fr.int8(iy)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int16(vx) > xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.int16(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int16(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int16(ix) > fr.int16(iy)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int32(vx) > xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.int32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int32(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int32(ix) > fr.int32(iy)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int64(vx) > xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.int64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int64(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int64(ix) > fr.int64(iy)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint(vx) > xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.uint(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint(ix) > fr.uint(iy)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint8(vx) > xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.uint8(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) > fr.uint8(iy)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint16(vx) > xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.uint16(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) > fr.uint16(iy)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint32(vx) > xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.uint32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) > fr.uint32(iy)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint64(vx) > xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.uint64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) > fr.uint64(iy)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uintptr(vx) > xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.uintptr(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) > fr.uintptr(iy)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Float32(vx) > xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float32(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.float32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.float32(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.float32(ix) > fr.float32(iy)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Float64(vx) > xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float64(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.float64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.float64(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.float64(ix) > fr.float64(iy)) }, true
				}
			case reflect.String:
				if kx == kindConst && ky == kindConst {
					v := xtype.String(vx) > xtype.String(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.String(vx)
					return func(fr *frame) { fr.setReg(ir, x > fr.string(iy)) }, true
				} else if ky == kindConst {
					y := xtype.String(vy)
					return func(fr *frame) { fr.setReg(ir, fr.string(ix) > y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.string(ix) > fr.string(iy)) }, true
				}
			}
		}
		panic("unreachable")
	case token.GEQ:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := vx.(int) >= vy.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(int)) }, true
				} else if ky == kindConst {
					y := vy.(int)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) >= fr.reg(iy).(int)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := vx.(int8) >= vy.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int8)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(int8)) }, true
				} else if ky == kindConst {
					y := vy.(int8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) >= fr.reg(iy).(int8)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := vx.(int16) >= vy.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int16)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(int16)) }, true
				} else if ky == kindConst {
					y := vy.(int16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) >= fr.reg(iy).(int16)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := vx.(int32) >= vy.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int32)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(int32)) }, true
				} else if ky == kindConst {
					y := vy.(int32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) >= fr.reg(iy).(int32)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := vx.(int64) >= vy.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(int64)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(int64)) }, true
				} else if ky == kindConst {
					y := vy.(int64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) >= fr.reg(iy).(int64)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint) >= vy.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uint)) }, true
				} else if ky == kindConst {
					y := vy.(uint)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) >= fr.reg(iy).(uint)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint8) >= vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uint8)) }, true
				} else if ky == kindConst {
					y := vy.(uint8)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) >= fr.reg(iy).(uint8)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint16) >= vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uint16)) }, true
				} else if ky == kindConst {
					y := vy.(uint16)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) >= fr.reg(iy).(uint16)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint32) >= vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uint32)) }, true
				} else if ky == kindConst {
					y := vy.(uint32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) >= fr.reg(iy).(uint32)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := vx.(uint64) >= vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uint64)) }, true
				} else if ky == kindConst {
					y := vy.(uint64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) >= fr.reg(iy).(uint64)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := vx.(uintptr) >= vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uintptr)) }, true
				} else if ky == kindConst {
					y := vy.(uintptr)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) >= fr.reg(iy).(uintptr)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := vx.(float32) >= vy.(float32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float32)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(float32)) }, true
				} else if ky == kindConst {
					y := vy.(float32)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) >= fr.reg(iy).(float32)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := vx.(float64) >= vy.(float64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(float64)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(float64)) }, true
				} else if ky == kindConst {
					y := vy.(float64)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) >= fr.reg(iy).(float64)) }, true
				}
			case reflect.String:
				if kx == kindConst && ky == kindConst {
					v := vx.(string) >= vy.(string)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := vx.(string)
					return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(string)) }, true
				} else if ky == kindConst {
					y := vy.(string)
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) >= fr.reg(iy).(string)) }, true
				}
			}
		} else {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int(vx) >= xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.int(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int(ix) >= fr.int(iy)) }, true
				}
			case reflect.Int8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int8(vx) >= xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int8(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.int8(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int8(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int8(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int8(ix) >= fr.int8(iy)) }, true
				}
			case reflect.Int16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int16(vx) >= xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int16(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.int16(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int16(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int16(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int16(ix) >= fr.int16(iy)) }, true
				}
			case reflect.Int32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int32(vx) >= xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int32(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.int32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int32(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int32(ix) >= fr.int32(iy)) }, true
				}
			case reflect.Int64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Int64(vx) >= xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Int64(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.int64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Int64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.int64(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.int64(ix) >= fr.int64(iy)) }, true
				}
			case reflect.Uint:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint(vx) >= xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.uint(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint(ix) >= fr.uint(iy)) }, true
				}
			case reflect.Uint8:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint8(vx) >= xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint8(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.uint8(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint8(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) >= fr.uint8(iy)) }, true
				}
			case reflect.Uint16:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint16(vx) >= xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint16(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.uint16(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint16(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) >= fr.uint16(iy)) }, true
				}
			case reflect.Uint32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint32(vx) >= xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint32(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.uint32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) >= fr.uint32(iy)) }, true
				}
			case reflect.Uint64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uint64(vx) >= xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uint64(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.uint64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uint64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) >= fr.uint64(iy)) }, true
				}
			case reflect.Uintptr:
				if kx == kindConst && ky == kindConst {
					v := xtype.Uintptr(vx) >= xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Uintptr(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.uintptr(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Uintptr(vy)
					return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) >= fr.uintptr(iy)) }, true
				}
			case reflect.Float32:
				if kx == kindConst && ky == kindConst {
					v := xtype.Float32(vx) >= xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float32(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.float32(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Float32(vy)
					return func(fr *frame) { fr.setReg(ir, fr.float32(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.float32(ix) >= fr.float32(iy)) }, true
				}
			case reflect.Float64:
				if kx == kindConst && ky == kindConst {
					v := xtype.Float64(vx) >= xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.Float64(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.float64(iy)) }, true
				} else if ky == kindConst {
					y := xtype.Float64(vy)
					return func(fr *frame) { fr.setReg(ir, fr.float64(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.float64(ix) >= fr.float64(iy)) }, true
				}
			case reflect.String:
				if kx == kindConst && ky == kindConst {
					v := xtype.String(vx) >= xtype.String(vy)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else if kx == kindConst {
					x := xtype.String(vx)
					return func(fr *frame) { fr.setReg(ir, x >= fr.string(iy)) }, true
				} else if ky == kindConst {
					y := xtype.String(vy)
					return func(fr *frame) { fr.setReg(ir, fr.string(ix) >= y) }, true
				} else {
					return func(fr *frame) { fr.setReg(ir, fr.string(ix) >= fr.string(iy)) }, true
				}
			}
		}
		panic("unreachable")
	case token.EQL:
		xtyp := aux.Type
		ytyp := aux.RightType
		if xtyp == nil && ytyp == nil {
			return func(fr *frame) {
				fr.setReg(ir, true)
			}, true
		}
		switch xtyp.Kind() {
		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128,
			reflect.String:
			if kx == kindConst && ky == kindConst {
				v := vx == vy
				return func(fr *frame) {
					fr.setReg(ir, v)
				}, true
			} else if kx == kindConst {
				return func(fr *frame) {
					fr.setReg(ir, vx == fr.reg(iy))
				}, true
			} else if ky == kindConst {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix) == vy)
				}, true
			}
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix) == fr.reg(iy))
			}, true
		case reflect.Interface:
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix) == fr.reg(iy))
			}, true
		case reflect.Array:
			if xtyp == ytyp {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix) == fr.reg(iy))
				}, true
			}
			return func(fr *frame) {
				x := fr.reg(ix)
				y := fr.reg(iy)
				fr.setReg(ir, x == reflect.ValueOf(y).Convert(xtyp).Interface())
			}, true
		case reflect.Struct:
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix) == fr.reg(iy))
			}, true
		case reflect.UnsafePointer:
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix) == fr.reg(iy))
			}, true
		case reflect.Chan:
			xdir := xtyp.ChanDir()
			ydir := ytyp.ChanDir()
			if xdir != ydir {
				if xdir == reflect.BothDir {
					return func(fr *frame) {
						x := fr.reg(ix)
						x = reflect.ValueOf(x).Convert(ytyp).Interface()
						fr.setReg(ir, x == fr.reg(iy))
					}, true
				} else if ydir == reflect.BothDir {
					return func(fr *frame) {
						y := fr.reg(iy)
						y = reflect.ValueOf(y).Convert(xtyp).Interface()
						fr.setReg(ir, fr.reg(ix) == y)
					}, true
				}
			}
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix) == fr.reg(iy))
			}, true
		case reflect.Ptr:
			return func(fr *frame) {
				x := fr.reg(ix)
				y := fr.reg(iy)
				fr.setReg(ir, reflect.ValueOf(x).Pointer() == reflect.ValueOf(y).Pointer())
			}, true
		case reflect.Slice, reflect.Map, reflect.Func:
			return func(fr *frame) {
				x := fr.reg(ix)
				y := fr.reg(iy)
				b := reflect.ValueOf(x).IsNil() && reflect.ValueOf(y).IsNil()
				fr.setReg(ir, b)
			}, true
		default:
			panic("unreachable")
		}
	case token.NEQ:
		xtyp := aux.Type
		ytyp := aux.RightType
		if xtyp == nil && ytyp == nil {
			return func(fr *frame) {
				fr.setReg(ir, false)
			}, true
		}
		switch xtyp.Kind() {
		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128,
			reflect.String:
			if kx == kindConst && ky == kindConst {
				v := vx != vy
				return func(fr *frame) {
					fr.setReg(ir, v)
				}, true
			} else if kx == kindConst {
				return func(fr *frame) {
					fr.setReg(ir, vx != fr.reg(iy))
				}, true
			} else if ky == kindConst {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix) != vy)
				}, true
			}
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix) != fr.reg(iy))
			}, true
		case reflect.Interface:
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix) != fr.reg(iy))
			}, true
		case reflect.Array:
			if xtyp == ytyp {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix) != fr.reg(iy))
				}, true
			}
			return func(fr *frame) {
				x := fr.reg(ix)
				y := fr.reg(iy)
				fr.setReg(ir, x != reflect.ValueOf(y).Convert(xtyp).Interface())
			}, true
		case reflect.Struct:
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix) != fr.reg(iy))
			}, true
		case reflect.UnsafePointer:
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix) != fr.reg(iy))
			}, true
		case reflect.Chan:
			xdir := xtyp.ChanDir()
			ydir := ytyp.ChanDir()
			if xdir != ydir {
				if xdir == reflect.BothDir {
					return func(fr *frame) {
						x := fr.reg(ix)
						x = reflect.ValueOf(x).Convert(ytyp).Interface()
						fr.setReg(ir, x != fr.reg(iy))
					}, true
				} else if ydir == reflect.BothDir {
					return func(fr *frame) {
						y := fr.reg(iy)
						y = reflect.ValueOf(y).Convert(xtyp).Interface()
						fr.setReg(ir, fr.reg(ix) != y)
					}, true
				}
			}
			return func(fr *frame) {
				fr.setReg(ir, fr.reg(ix) != fr.reg(iy))
			}, true
		case reflect.Ptr:
			return func(fr *frame) {
				x := fr.reg(ix)
				y := fr.reg(iy)
				fr.setReg(ir, reflect.ValueOf(x).Pointer() != reflect.ValueOf(y).Pointer())
			}, true
		case reflect.Slice, reflect.Map, reflect.Func:
			return func(fr *frame) {
				x := fr.reg(ix)
				y := fr.reg(iy)
				b := reflect.ValueOf(x).IsNil() && reflect.ValueOf(y).IsNil()
				fr.setReg(ir, !b)
			}, true
		default:
			panic("unreachable")
		}
	case token.SHL:
		xtyp := aux.Type
		ytyp := aux.RightType
		xkind := xtyp.Kind()
		ykind := ytyp.Kind()
		if kx == kindConst && ky == kindConst {
			t := xtype.TypeOfType(xtyp)
			switch xkind {
			case reflect.Int:
				x := xtype.Int(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Int8:
				x := xtype.Int8(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Int16:
				x := xtype.Int16(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Int32:
				x := xtype.Int32(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Int64:
				x := xtype.Int64(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uint:
				x := xtype.Uint(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uint8:
				x := xtype.Uint8(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uint16:
				x := xtype.Uint16(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uint32:
				x := xtype.Uint32(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uint64:
				x := xtype.Uint64(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uintptr:
				x := xtype.Uintptr(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x<<xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x<<xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x<<xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x<<xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x<<xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x<<xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x<<xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x<<xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x<<xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x<<xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x<<xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			}
		}
		if xtyp.PkgPath() == "" {
			switch xkind {
			case reflect.Int:
				if kx == kindConst {
					x := vx.(int)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)<<fr.uintptr(iy)) }, true
					}
				}
			case reflect.Int8:
				if kx == kindConst {
					x := vx.(int8)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)<<fr.uintptr(iy)) }, true
					}
				}
			case reflect.Int16:
				if kx == kindConst {
					x := vx.(int16)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)<<fr.uintptr(iy)) }, true
					}
				}
			case reflect.Int32:
				if kx == kindConst {
					x := vx.(int32)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)<<fr.uintptr(iy)) }, true
					}
				}
			case reflect.Int64:
				if kx == kindConst {
					x := vx.(int64)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)<<fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uint:
				if kx == kindConst {
					x := vx.(uint)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)<<fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uint8:
				if kx == kindConst {
					x := vx.(uint8)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)<<fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uint16:
				if kx == kindConst {
					x := vx.(uint16)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)<<fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uint32:
				if kx == kindConst {
					x := vx.(uint32)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)<<fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uint64:
				if kx == kindConst {
					x := vx.(uint64)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)<<fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uintptr:
				if kx == kindConst {
					x := vx.(uintptr)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x<<fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)<<fr.uintptr(iy)) }, true
					}
				}
			}
		} else {
			t := xtype.TypeOfType(xtyp)
			switch xkind {
			case reflect.Int:
				if kx == kindConst {
					x := xtype.Int(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)<<fr.uintptr(iy))) }, true
					}
				}
			case reflect.Int8:
				if kx == kindConst {
					x := xtype.Int8(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)<<fr.uintptr(iy))) }, true
					}
				}
			case reflect.Int16:
				if kx == kindConst {
					x := xtype.Int16(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)<<fr.uintptr(iy))) }, true
					}
				}
			case reflect.Int32:
				if kx == kindConst {
					x := xtype.Int32(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)<<fr.uintptr(iy))) }, true
					}
				}
			case reflect.Int64:
				if kx == kindConst {
					x := xtype.Int64(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)<<fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uint:
				if kx == kindConst {
					x := xtype.Uint(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)<<fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uint8:
				if kx == kindConst {
					x := xtype.Uint8(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)<<fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uint16:
				if kx == kindConst {
					x := xtype.Uint16(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)<<fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uint32:
				if kx == kindConst {
					x := xtype.Uint32(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)<<fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uint64:
				if kx == kindConst {
					x := xtype.Uint64(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)<<fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uintptr:
				if kx == kindConst {
					x := xtype.Uintptr(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x<<fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)<<fr.uintptr(iy))) }, true
					}
				}
			}
		}
		panic("unreachable")
	case token.SHR:
		xtyp := aux.Type
		ytyp := aux.RightType
		xkind := xtyp.Kind()
		ykind := ytyp.Kind()
		if kx == kindConst && ky == kindConst {
			t := xtype.TypeOfType(xtyp)
			switch xkind {
			case reflect.Int:
				x := xtype.Int(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Int8:
				x := xtype.Int8(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Int16:
				x := xtype.Int16(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Int32:
				x := xtype.Int32(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Int64:
				x := xtype.Int64(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uint:
				x := xtype.Uint(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uint8:
				x := xtype.Uint8(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uint16:
				x := xtype.Uint16(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uint32:
				x := xtype.Uint32(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uint64:
				x := xtype.Uint64(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			case reflect.Uintptr:
				x := xtype.Uintptr(vx)
				switch ykind {
				case reflect.Int:
					v := xtype.Make(t, x>>xtype.Int(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int8:
					v := xtype.Make(t, x>>xtype.Int8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int16:
					v := xtype.Make(t, x>>xtype.Int16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int32:
					v := xtype.Make(t, x>>xtype.Int32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Int64:
					v := xtype.Make(t, x>>xtype.Int64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint:
					v := xtype.Make(t, x>>xtype.Uint(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint8:
					v := xtype.Make(t, x>>xtype.Uint8(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint16:
					v := xtype.Make(t, x>>xtype.Uint16(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint32:
					v := xtype.Make(t, x>>xtype.Uint32(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uint64:
					v := xtype.Make(t, x>>xtype.Uint64(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				case reflect.Uintptr:
					v := xtype.Make(t, x>>xtype.Uintptr(vy))
					return func(fr *frame) { fr.setReg(ir, v) }, true
				}
			}
		}
		if xtyp.PkgPath() == "" {
			switch xkind {
			case reflect.Int:
				if kx == kindConst {
					x := vx.(int)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)>>fr.uintptr(iy)) }, true
					}
				}
			case reflect.Int8:
				if kx == kindConst {
					x := vx.(int8)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)>>fr.uintptr(iy)) }, true
					}
				}
			case reflect.Int16:
				if kx == kindConst {
					x := vx.(int16)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)>>fr.uintptr(iy)) }, true
					}
				}
			case reflect.Int32:
				if kx == kindConst {
					x := vx.(int32)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)>>fr.uintptr(iy)) }, true
					}
				}
			case reflect.Int64:
				if kx == kindConst {
					x := vx.(int64)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)>>fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uint:
				if kx == kindConst {
					x := vx.(uint)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)>>fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uint8:
				if kx == kindConst {
					x := vx.(uint8)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)>>fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uint16:
				if kx == kindConst {
					x := vx.(uint16)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)>>fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uint32:
				if kx == kindConst {
					x := vx.(uint32)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)>>fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uint64:
				if kx == kindConst {
					x := vx.(uint64)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)>>fr.uintptr(iy)) }, true
					}
				}
			case reflect.Uintptr:
				if kx == kindConst {
					x := vx.(uintptr)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, x>>fr.uintptr(iy)) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>y) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.int(iy)) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.int8(iy)) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.int16(iy)) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.int32(iy)) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.int64(iy)) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.uint(iy)) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.uint8(iy)) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.uint16(iy)) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.uint32(iy)) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.uint64(iy)) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)>>fr.uintptr(iy)) }, true
					}
				}
			}
		} else {
			t := xtype.TypeOfType(xtyp)
			switch xkind {
			case reflect.Int:
				if kx == kindConst {
					x := xtype.Int(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int(ix)>>fr.uintptr(iy))) }, true
					}
				}
			case reflect.Int8:
				if kx == kindConst {
					x := xtype.Int8(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int8(ix)>>fr.uintptr(iy))) }, true
					}
				}
			case reflect.Int16:
				if kx == kindConst {
					x := xtype.Int16(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int16(ix)>>fr.uintptr(iy))) }, true
					}
				}
			case reflect.Int32:
				if kx == kindConst {
					x := xtype.Int32(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int32(ix)>>fr.uintptr(iy))) }, true
					}
				}
			case reflect.Int64:
				if kx == kindConst {
					x := xtype.Int64(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.int64(ix)>>fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uint:
				if kx == kindConst {
					x := xtype.Uint(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint(ix)>>fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uint8:
				if kx == kindConst {
					x := xtype.Uint8(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint8(ix)>>fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uint16:
				if kx == kindConst {
					x := xtype.Uint16(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint16(ix)>>fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uint32:
				if kx == kindConst {
					x := xtype.Uint32(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint32(ix)>>fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uint64:
				if kx == kindConst {
					x := xtype.Uint64(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uint64(ix)>>fr.uintptr(iy))) }, true
					}
				}
			case reflect.Uintptr:
				if kx == kindConst {
					x := xtype.Uintptr(vx)
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, x>>fr.uintptr(iy))) }, true
					}
				} else if ky == kindConst {
					switch ykind {
					case reflect.Int:
						y := xtype.Int(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					case reflect.Int8:
						y := xtype.Int8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					case reflect.Int16:
						y := xtype.Int16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					case reflect.Int32:
						y := xtype.Int32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					case reflect.Int64:
						y := xtype.Int64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					case reflect.Uint:
						y := xtype.Uint(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					case reflect.Uint8:
						y := xtype.Uint8(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					case reflect.Uint16:
						y := xtype.Uint16(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					case reflect.Uint32:
						y := xtype.Uint32(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					case reflect.Uint64:
						y := xtype.Uint64(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					case reflect.Uintptr:
						y := xtype.Uintptr(vy)
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>y)) }, true
					}
				} else {
					switch ykind {
					case reflect.Int:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.int(iy))) }, true
					case reflect.Int8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.int8(iy))) }, true
					case reflect.Int16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.int16(iy))) }, true
					case reflect.Int32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.int32(iy))) }, true
					case reflect.Int64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.int64(iy))) }, true
					case reflect.Uint:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.uint(iy))) }, true
					case reflect.Uint8:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.uint8(iy))) }, true
					case reflect.Uint16:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.uint16(iy))) }, true
					case reflect.Uint32:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.uint32(iy))) }, true
					case reflect.Uint64:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.uint64(iy))) }, true
					case reflect.Uintptr:
						return func(fr *frame) { fr.setReg(ir, xtype.Make(t, fr.uintptr(ix)>>fr.uintptr(iy))) }, true
					}
				}
			}
		}
		panic("unreachable")
	}
	return nil, false
}

// makeUnaryInstruction builds unary operations from runtime descriptions.
func makeUnaryInstruction(d typedInstruction) (func(*frame), bool) {
	aux := d.Aux.(unaryOpAux)
	ir, ix := d.Result, d.Args[0]
	kx := kindInvalid
	if aux.Constant {
		kx = kindConst
	}
	vx := aux.Value
	switch aux.Op {
	case token.SUB:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst {
					v := -vx.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(int)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int8:
				if kx == kindConst {
					v := -vx.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(int8)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int16:
				if kx == kindConst {
					v := -vx.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(int16)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int32:
				if kx == kindConst {
					v := -vx.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(int32)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int64:
				if kx == kindConst {
					v := -vx.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(int64)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint:
				if kx == kindConst {
					v := -vx.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(uint)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint8:
				if kx == kindConst {
					v := -vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(uint8)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint16:
				if kx == kindConst {
					v := -vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(uint16)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint32:
				if kx == kindConst {
					v := -vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(uint32)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint64:
				if kx == kindConst {
					v := -vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(uint64)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uintptr:
				if kx == kindConst {
					v := -vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(uintptr)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Float32:
				if kx == kindConst {
					v := -vx.(float32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(float32)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Float64:
				if kx == kindConst {
					v := -vx.(float64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(float64)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Complex64:
				if kx == kindConst {
					v := -vx.(complex64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(complex64)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Complex128:
				if kx == kindConst {
					v := -vx.(complex128)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := -fr.reg(ix).(complex128)
						fr.setReg(ir, v)
					}, true
				}
			}
		} else {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst {
					v := xtype.NegInt(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegInt(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int8:
				if kx == kindConst {
					v := xtype.NegInt8(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegInt8(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int16:
				if kx == kindConst {
					v := xtype.NegInt16(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegInt16(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int32:
				if kx == kindConst {
					v := xtype.NegInt32(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegInt32(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int64:
				if kx == kindConst {
					v := xtype.NegInt64(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegInt64(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint:
				if kx == kindConst {
					v := xtype.NegUint(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegUint(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint8:
				if kx == kindConst {
					v := xtype.NegUint8(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegUint8(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint16:
				if kx == kindConst {
					v := xtype.NegUint16(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegUint16(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint32:
				if kx == kindConst {
					v := xtype.NegUint32(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegUint32(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint64:
				if kx == kindConst {
					v := xtype.NegUint64(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegUint64(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uintptr:
				if kx == kindConst {
					v := xtype.NegUintptr(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegUintptr(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Float32:
				if kx == kindConst {
					v := xtype.NegFloat32(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegFloat32(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Float64:
				if kx == kindConst {
					v := xtype.NegFloat64(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegFloat64(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Complex64:
				if kx == kindConst {
					v := xtype.NegComplex64(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegComplex64(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Complex128:
				if kx == kindConst {
					v := xtype.NegComplex128(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.NegComplex128(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			}
		}
		panic("unreachable")
	case token.XOR:
		typ := aux.Type
		if typ.PkgPath() == "" {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst {
					v := ^vx.(int)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(int)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int8:
				if kx == kindConst {
					v := ^vx.(int8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(int8)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int16:
				if kx == kindConst {
					v := ^vx.(int16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(int16)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int32:
				if kx == kindConst {
					v := ^vx.(int32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(int32)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int64:
				if kx == kindConst {
					v := ^vx.(int64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(int64)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint:
				if kx == kindConst {
					v := ^vx.(uint)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(uint)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint8:
				if kx == kindConst {
					v := ^vx.(uint8)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(uint8)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint16:
				if kx == kindConst {
					v := ^vx.(uint16)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(uint16)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint32:
				if kx == kindConst {
					v := ^vx.(uint32)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(uint32)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint64:
				if kx == kindConst {
					v := ^vx.(uint64)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(uint64)
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uintptr:
				if kx == kindConst {
					v := ^vx.(uintptr)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := ^fr.reg(ix).(uintptr)
						fr.setReg(ir, v)
					}, true
				}
			}
		} else {
			switch typ.Kind() {
			case reflect.Int:
				if kx == kindConst {
					v := xtype.XorInt(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorInt(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int8:
				if kx == kindConst {
					v := xtype.XorInt8(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorInt8(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int16:
				if kx == kindConst {
					v := xtype.XorInt16(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorInt16(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int32:
				if kx == kindConst {
					v := xtype.XorInt32(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorInt32(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Int64:
				if kx == kindConst {
					v := xtype.XorInt64(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorInt64(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint:
				if kx == kindConst {
					v := xtype.XorUint(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorUint(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint8:
				if kx == kindConst {
					v := xtype.XorUint8(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorUint8(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint16:
				if kx == kindConst {
					v := xtype.XorUint16(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorUint16(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint32:
				if kx == kindConst {
					v := xtype.XorUint32(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorUint32(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uint64:
				if kx == kindConst {
					v := xtype.XorUint64(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorUint64(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			case reflect.Uintptr:
				if kx == kindConst {
					v := xtype.XorUintptr(vx)
					return func(fr *frame) { fr.setReg(ir, v) }, true
				} else {
					return func(fr *frame) {
						v := xtype.XorUintptr(fr.reg(ix))
						fr.setReg(ir, v)
					}, true
				}
			}
		}
		panic("unreachable")
	case token.NOT:
		typ := aux.Type
		if typ.PkgPath() == "" {
			if kx == kindConst {
				v := !vx.(bool)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}, true
			}
			return func(fr *frame) {
				v := !fr.reg(ix).(bool)
				fr.setReg(ir, v)
			}, true
		} else {
			if kx == kindConst {
				v := xtype.Not(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}, true
			}
			return func(fr *frame) {
				v := xtype.Not(fr.reg(ix))
				fr.setReg(ir, v)
			}, true
		}

	case token.MUL:
		if aux.Global {
			x := reflect.ValueOf(aux.Value)
			return func(fr *frame) {
				elem := x.Elem()
				if !elem.IsValid() {
					panic(fr.runtimeError(fr.pfn.InstrForPC(int(fr.ipc)-1), "invalid memory address or nil pointer dereference"))
				}
				fr.setReg(ir, elem.Interface())
			}, true
		}
		return func(fr *frame) {
			elem := reflect.ValueOf(fr.reg(ix)).Elem()
			if !elem.IsValid() {
				panic(fr.runtimeError(fr.pfn.InstrForPC(int(fr.ipc)-1), "invalid memory address or nil pointer dereference"))
			}
			fr.setReg(ir, elem.Interface())
		}, true
	case token.ARROW:
		typ := aux.Type.Elem()
		if aux.Global {
			x := reflect.ValueOf(aux.Value)
			if aux.CommaOk {
				return func(fr *frame) {
					if fr.pfn.Interp.gcEnabled {
						fr.gcFrames()
					}
					v, ok := x.Recv()
					if !ok {
						v = reflect.New(typ).Elem()
					}
					fr.setReg(ir, tuple{v.Interface(), ok})
				}, true
			}
			return func(fr *frame) {
				if fr.pfn.Interp.gcEnabled {
					fr.gcFrames()
				}
				v, ok := x.Recv()
				if !ok {
					v = reflect.New(typ).Elem()
				}
				fr.setReg(ir, v.Interface())
			}, true
		}
		if aux.CommaOk {
			return func(fr *frame) {
				x := reflect.ValueOf(fr.reg(ix))
				if fr.pfn.Interp.gcEnabled {
					fr.gcFrames()
				}
				v, ok := x.Recv()
				if !ok {
					v = reflect.New(typ).Elem()
				}
				fr.setReg(ir, tuple{v.Interface(), ok})
			}, true
		}
		return func(fr *frame) {
			x := reflect.ValueOf(fr.reg(ix))
			if fr.pfn.Interp.gcEnabled {
				fr.gcFrames()
			}
			v, ok := x.Recv()
			if !ok {
				v = reflect.New(typ).Elem()
			}
			fr.setReg(ir, v.Interface())
		}, true
	}
	return nil, false
}
