package ixgo

import (
	"bytes"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"testing"
	"unsafe"

	"golang.org/x/tools/go/ssa"

	foo "github.com/goplus/ixgo/testdata/foo/v2"
	"github.com/goplus/ixgo/testdata/getpkg"
	unknownpkg "github.com/goplus/ixgo/testdata/unknownpkg"
)

type getpkgFieldHolder struct {
	T getpkg.T
}

type getpkgCurpkgHolder struct {
	A foo.T
	B unknownpkg.U
}

func TestGetPackageLoadsRegistered(t *testing.T) {
	const path = "ixgo.test/getpkgreg"
	RegisterPackage(&Package{
		Name: "getpkgreg",
		Path: path,
		UntypedConsts: map[string]UntypedConst{
			"Answer": {Typ: "untyped int", Value: constant.MakeInt64(42)},
		},
	})
	loader := NewContext(0).Loader.(*TypesLoader)
	pkg := loader.GetPackage(path)
	if !pkg.Complete() {
		t.Fatalf("GetPackage(%q) returned an incomplete placeholder", path)
	}
	if pkg.Scope().Lookup("Answer") == nil {
		t.Fatal("registered const Answer was not installed")
	}
}

func TestGetPackageUnknownIsPlaceholder(t *testing.T) {
	loader := NewContext(0).Loader.(*TypesLoader)
	pkg := loader.GetPackage("ixgo.test/does-not-exist")
	if pkg.Complete() {
		t.Fatal("unknown package should remain a placeholder")
	}
}

func TestGetPackageLoadsRegisteredFromReflectField(t *testing.T) {
	const depPath = "github.com/goplus/ixgo/testdata/getpkg"
	RegisterPackage(&Package{
		Name: "getpkg",
		Path: depPath,
		NamedTypes: map[string]reflect.Type{
			"T": reflect.TypeOf((*getpkg.T)(nil)).Elem(),
		},
	})
	RegisterPackage(&Package{
		Name: "holder",
		Path: "ixgo.test/getpkgholder",
		NamedTypes: map[string]reflect.Type{
			"Holder": reflect.TypeOf((*getpkgFieldHolder)(nil)).Elem(),
		},
	})
	ctx := NewContext(0)
	if _, err := ctx.Loader.Import("ixgo.test/getpkgholder"); err != nil {
		t.Fatal(err)
	}
	pkg, err := ctx.Loader.Import(depPath)
	if err != nil {
		t.Fatal(err)
	}
	if !pkg.Complete() {
		t.Fatal("getpkg should be loaded from the registry, not left as a placeholder")
	}
	obj := pkg.Scope().Lookup("T")
	if obj == nil {
		t.Fatal("missing type T")
	}
	if _, ok := obj.Type().Underlying().(*types.Struct); !ok {
		t.Fatalf("T type = %T, want struct", obj.Type().Underlying())
	}
}

func TestGetPackagePreservesCurpkgAfterNestedImport(t *testing.T) {
	const (
		fooPath     = "github.com/goplus/ixgo/testdata/foo/v2"
		unknownPath = "github.com/goplus/ixgo/testdata/unknownpkg"
		holderPath  = "ixgo.test/curpkgholder"
	)
	RegisterPackage(&Package{
		Name: "foo",
		Path: fooPath,
		NamedTypes: map[string]reflect.Type{
			"T": reflect.TypeOf((*foo.T)(nil)).Elem(),
		},
	})
	RegisterPackage(&Package{
		Name: "holder",
		Path: holderPath,
		Deps: map[string]string{
			fooPath:     "foo",
			unknownPath: "othername",
		},
		NamedTypes: map[string]reflect.Type{
			"Holder": reflect.TypeOf((*getpkgCurpkgHolder)(nil)).Elem(),
		},
	})
	ctx := NewContext(0)
	if _, err := ctx.Loader.Import(holderPath); err != nil {
		t.Fatal(err)
	}
	loader := ctx.Loader.(*TypesLoader)
	pkg := loader.GetPackage(unknownPath)
	if got, want := pkg.Name(), "othername"; got != want {
		t.Fatalf("unknown package name = %q, want %q (curpkg.Deps lost after nested Import)", got, want)
	}
}

func TestGetPackageRegisteredImportErrorFallsBack(t *testing.T) {
	const path = "ixgo.test/badsrc"
	RegisterPackage(&Package{
		Name:   "badsrc",
		Path:   path,
		Source: "package badsrc\nfunc (",
	})
	loader := NewContext(0).Loader.(*TypesLoader)
	pkg := loader.GetPackage(path)
	if pkg.Complete() {
		t.Fatal("failed registered source load should fall back to a placeholder")
	}
	if got, want := pkg.Name(), "badsrc"; got != want {
		t.Fatalf("placeholder name = %q, want %q", got, want)
	}
}

type namedInt int
type namedUint uint
type namedFloat float64
type namedComplex complex128
type namedString string
type namedInt8 int8

func TestAsIntAndUint64(t *testing.T) {
	if asInt(int(7)) != 7 || asInt(int8(7)) != 7 || asInt(int16(7)) != 7 ||
		asInt(int32(7)) != 7 || asInt(int64(7)) != 7 || asInt(uint(7)) != 7 ||
		asInt(uint8(7)) != 7 || asInt(uint16(7)) != 7 || asInt(uint32(7)) != 7 ||
		asInt(uint64(7)) != 7 || asInt(uintptr(7)) != 7 {
		t.Fatal("asInt basic")
	}
	if asInt(namedInt(9)) != 9 || asInt(namedUint(9)) != 9 {
		t.Fatal("asInt named")
	}
	mustPanic(t, func() { asInt("x") })

	if asUint64(int(3)) != 3 || asUint64(int8(3)) != 3 || asUint64(int16(3)) != 3 ||
		asUint64(int32(3)) != 3 || asUint64(int64(3)) != 3 || asUint64(uint(3)) != 3 ||
		asUint64(uint8(3)) != 3 || asUint64(uint16(3)) != 3 || asUint64(uint32(3)) != 3 ||
		asUint64(uint64(3)) != 3 || asUint64(uintptr(3)) != 3 {
		t.Fatal("asUint64 basic")
	}
	if asUint64(namedUint(4)) != 4 || asUint64(namedInt(4)) != 4 {
		t.Fatal("asUint64 named")
	}
	mustPanic(t, func() { asUint64(int(-1)) })
	mustPanic(t, func() { asUint64(namedInt(-1)) })
	mustPanic(t, func() { asUint64("x") })
}

func TestLegacyBinops(t *testing.T) {
	if opADD(1, 2).(int) != 3 || opADD(int8(1), int8(2)).(int8) != 3 ||
		opADD(int16(1), int16(2)).(int16) != 3 || opADD(int32(1), int32(2)).(int32) != 3 ||
		opADD(int64(1), int64(2)).(int64) != 3 || opADD(uint(1), uint(2)).(uint) != 3 ||
		opADD(uint8(1), uint8(2)).(uint8) != 3 || opADD(uint16(1), uint16(2)).(uint16) != 3 ||
		opADD(uint32(1), uint32(2)).(uint32) != 3 || opADD(uint64(1), uint64(2)).(uint64) != 3 ||
		opADD(uintptr(1), uintptr(2)).(uintptr) != 3 || opADD(float32(1), float32(2)).(float32) != 3 ||
		opADD(float64(1), float64(2)).(float64) != 3 || opADD(complex64(1), complex64(2)).(complex64) != 3 ||
		opADD(complex128(1), complex128(2)).(complex128) != 3 || opADD("a", "b").(string) != "ab" {
		t.Fatal("opADD basic")
	}
	if opADD(namedInt(1), namedInt(2)).(namedInt) != 3 ||
		opADD(namedUint(1), namedUint(2)).(namedUint) != 3 ||
		opADD(namedFloat(1), namedFloat(2)).(namedFloat) != 3 ||
		opADD(namedComplex(1), namedComplex(2)).(namedComplex) != 3 ||
		opADD(namedString("a"), namedString("b")).(namedString) != "ab" {
		t.Fatal("opADD named")
	}
	mustPanic(t, func() { opADD(true, true) })

	if opSUB(5, 2).(int) != 3 || opSUB(int8(5), int8(2)).(int8) != 3 ||
		opSUB(int16(5), int16(2)).(int16) != 3 || opSUB(int32(5), int32(2)).(int32) != 3 ||
		opSUB(int64(5), int64(2)).(int64) != 3 || opSUB(uint(5), uint(2)).(uint) != 3 ||
		opSUB(uint8(5), uint8(2)).(uint8) != 3 || opSUB(uint16(5), uint16(2)).(uint16) != 3 ||
		opSUB(uint32(5), uint32(2)).(uint32) != 3 || opSUB(uint64(5), uint64(2)).(uint64) != 3 ||
		opSUB(uintptr(5), uintptr(2)).(uintptr) != 3 || opSUB(float32(5), float32(2)).(float32) != 3 ||
		opSUB(float64(5), float64(2)).(float64) != 3 || opSUB(complex64(5), complex64(2)).(complex64) != 3 ||
		opSUB(complex128(5), complex128(2)).(complex128) != 3 {
		t.Fatal("opSUB")
	}
	if opSUB(namedInt(5), namedInt(2)).(namedInt) != 3 {
		t.Fatal("opSUB named")
	}

	if opMUL(3, 4).(int) != 12 || opMUL(int8(3), int8(4)).(int8) != 12 ||
		opMUL(int16(3), int16(4)).(int16) != 12 || opMUL(int32(3), int32(4)).(int32) != 12 ||
		opMUL(int64(3), int64(4)).(int64) != 12 || opMUL(uint(3), uint(4)).(uint) != 12 ||
		opMUL(uint8(3), uint8(4)).(uint8) != 12 || opMUL(uint16(3), uint16(4)).(uint16) != 12 ||
		opMUL(uint32(3), uint32(4)).(uint32) != 12 || opMUL(uint64(3), uint64(4)).(uint64) != 12 ||
		opMUL(uintptr(3), uintptr(4)).(uintptr) != 12 || opMUL(float32(3), float32(4)).(float32) != 12 ||
		opMUL(float64(3), float64(4)).(float64) != 12 || opMUL(complex64(3), complex64(4)).(complex64) != 12 ||
		opMUL(complex128(3), complex128(4)).(complex128) != 12 {
		t.Fatal("opMUL")
	}
	if opMUL(namedInt(3), namedInt(4)).(namedInt) != 12 {
		t.Fatal("opMUL named")
	}

	if opQuo(8, 2).(int) != 4 || opQuo(int8(8), int8(2)).(int8) != 4 ||
		opQuo(int16(8), int16(2)).(int16) != 4 || opQuo(int32(8), int32(2)).(int32) != 4 ||
		opQuo(int64(8), int64(2)).(int64) != 4 || opQuo(uint(8), uint(2)).(uint) != 4 ||
		opQuo(uint8(8), uint8(2)).(uint8) != 4 || opQuo(uint16(8), uint16(2)).(uint16) != 4 ||
		opQuo(uint32(8), uint32(2)).(uint32) != 4 || opQuo(uint64(8), uint64(2)).(uint64) != 4 ||
		opQuo(uintptr(8), uintptr(2)).(uintptr) != 4 || opQuo(float32(8), float32(2)).(float32) != 4 ||
		opQuo(float64(8), float64(2)).(float64) != 4 || opQuo(complex64(8), complex64(2)).(complex64) != 4 ||
		opQuo(complex128(8), complex128(2)).(complex128) != 4 {
		t.Fatal("opQuo")
	}
	if opQuo(namedInt(8), namedInt(2)).(namedInt) != 4 {
		t.Fatal("opQuo named")
	}

	if opREM(7, 3).(int) != 1 || opREM(int8(7), int8(3)).(int8) != 1 ||
		opREM(int16(7), int16(3)).(int16) != 1 || opREM(int32(7), int32(3)).(int32) != 1 ||
		opREM(int64(7), int64(3)).(int64) != 1 || opREM(uint(7), uint(3)).(uint) != 1 ||
		opREM(uint8(7), uint8(3)).(uint8) != 1 || opREM(uint16(7), uint16(3)).(uint16) != 1 ||
		opREM(uint32(7), uint32(3)).(uint32) != 1 || opREM(uint64(7), uint64(3)).(uint64) != 1 ||
		opREM(uintptr(7), uintptr(3)).(uintptr) != 1 {
		t.Fatal("opREM")
	}
	if opREM(namedInt(7), namedInt(3)).(namedInt) != 1 {
		t.Fatal("opREM named")
	}

	if opAND(6, 3).(int) != 2 || opAND(int8(6), int8(3)).(int8) != 2 ||
		opAND(int16(6), int16(3)).(int16) != 2 || opAND(int32(6), int32(3)).(int32) != 2 ||
		opAND(int64(6), int64(3)).(int64) != 2 || opAND(uint(6), uint(3)).(uint) != 2 ||
		opAND(uint8(6), uint8(3)).(uint8) != 2 || opAND(uint16(6), uint16(3)).(uint16) != 2 ||
		opAND(uint32(6), uint32(3)).(uint32) != 2 || opAND(uint64(6), uint64(3)).(uint64) != 2 ||
		opAND(uintptr(6), uintptr(3)).(uintptr) != 2 {
		t.Fatal("opAND")
	}
	if opAND(namedInt(6), namedInt(3)).(namedInt) != 2 {
		t.Fatal("opAND named")
	}

	if opOR(4, 1).(int) != 5 || opOR(int8(4), int8(1)).(int8) != 5 ||
		opOR(int16(4), int16(1)).(int16) != 5 || opOR(int32(4), int32(1)).(int32) != 5 ||
		opOR(int64(4), int64(1)).(int64) != 5 || opOR(uint(4), uint(1)).(uint) != 5 ||
		opOR(uint8(4), uint8(1)).(uint8) != 5 || opOR(uint16(4), uint16(1)).(uint16) != 5 ||
		opOR(uint32(4), uint32(1)).(uint32) != 5 || opOR(uint64(4), uint64(1)).(uint64) != 5 ||
		opOR(uintptr(4), uintptr(1)).(uintptr) != 5 {
		t.Fatal("opOR")
	}
	if opOR(namedInt(4), namedInt(1)).(namedInt) != 5 {
		t.Fatal("opOR named")
	}

	if opXOR(6, 3).(int) != 5 || opXOR(int8(6), int8(3)).(int8) != 5 ||
		opXOR(int16(6), int16(3)).(int16) != 5 || opXOR(int32(6), int32(3)).(int32) != 5 ||
		opXOR(int64(6), int64(3)).(int64) != 5 || opXOR(uint(6), uint(3)).(uint) != 5 ||
		opXOR(uint8(6), uint8(3)).(uint8) != 5 || opXOR(uint16(6), uint16(3)).(uint16) != 5 ||
		opXOR(uint32(6), uint32(3)).(uint32) != 5 || opXOR(uint64(6), uint64(3)).(uint64) != 5 ||
		opXOR(uintptr(6), uintptr(3)).(uintptr) != 5 {
		t.Fatal("opXOR")
	}
	if opXOR(namedInt(6), namedInt(3)).(namedInt) != 5 {
		t.Fatal("opXOR named")
	}

	if opANDNOT(7, 1).(int) != 6 || opANDNOT(int8(7), int8(1)).(int8) != 6 ||
		opANDNOT(int16(7), int16(1)).(int16) != 6 || opANDNOT(int32(7), int32(1)).(int32) != 6 ||
		opANDNOT(int64(7), int64(1)).(int64) != 6 || opANDNOT(uint(7), uint(1)).(uint) != 6 ||
		opANDNOT(uint8(7), uint8(1)).(uint8) != 6 || opANDNOT(uint16(7), uint16(1)).(uint16) != 6 ||
		opANDNOT(uint32(7), uint32(1)).(uint32) != 6 || opANDNOT(uint64(7), uint64(1)).(uint64) != 6 ||
		opANDNOT(uintptr(7), uintptr(1)).(uintptr) != 6 {
		t.Fatal("opANDNOT")
	}
	if opANDNOT(namedInt(7), namedInt(1)).(namedInt) != 6 {
		t.Fatal("opANDNOT named")
	}

	if opSHL(1, uint(2)).(int) != 4 || opSHL(int8(1), uint8(2)).(int8) != 4 ||
		opSHL(int16(1), uint16(2)).(int16) != 4 || opSHL(int32(1), uint32(2)).(int32) != 4 ||
		opSHL(int64(1), uint64(2)).(int64) != 4 || opSHL(uint(1), uint(2)).(uint) != 4 ||
		opSHL(uint8(1), uint8(2)).(uint8) != 4 || opSHL(uint16(1), uint16(2)).(uint16) != 4 ||
		opSHL(uint32(1), uint32(2)).(uint32) != 4 || opSHL(uint64(1), uint64(2)).(uint64) != 4 ||
		opSHL(uintptr(1), uintptr(2)).(uintptr) != 4 {
		t.Fatal("opSHL")
	}
	if opSHL(namedInt(1), namedUint(2)).(namedInt) != 4 {
		t.Fatal("opSHL named")
	}

	if opSHR(8, uint(2)).(int) != 2 || opSHR(int8(8), uint8(2)).(int8) != 2 ||
		opSHR(int16(8), uint16(2)).(int16) != 2 || opSHR(int32(8), uint32(2)).(int32) != 2 ||
		opSHR(int64(8), uint64(2)).(int64) != 2 || opSHR(uint(8), uint(2)).(uint) != 2 ||
		opSHR(uint8(8), uint8(2)).(uint8) != 2 || opSHR(uint16(8), uint16(2)).(uint16) != 2 ||
		opSHR(uint32(8), uint32(2)).(uint32) != 2 || opSHR(uint64(8), uint64(2)).(uint64) != 2 ||
		opSHR(uintptr(8), uintptr(2)).(uintptr) != 2 {
		t.Fatal("opSHR")
	}
	if opSHR(namedInt(8), namedUint(2)).(namedInt) != 2 {
		t.Fatal("opSHR named")
	}

	if !opLSS(1, 2).(bool) || !opLSS(int8(1), int8(2)).(bool) ||
		!opLSS(int16(1), int16(2)).(bool) || !opLSS(int32(1), int32(2)).(bool) ||
		!opLSS(int64(1), int64(2)).(bool) || !opLSS(uint(1), uint(2)).(bool) ||
		!opLSS(uint8(1), uint8(2)).(bool) || !opLSS(uint16(1), uint16(2)).(bool) ||
		!opLSS(uint32(1), uint32(2)).(bool) || !opLSS(uint64(1), uint64(2)).(bool) ||
		!opLSS(uintptr(1), uintptr(2)).(bool) || !opLSS(float32(1), float32(2)).(bool) ||
		!opLSS(float64(1), float64(2)).(bool) || !opLSS("a", "b").(bool) {
		t.Fatal("opLSS")
	}
	if !opLSS(namedInt(1), namedInt(2)).(bool) {
		t.Fatal("opLSS named")
	}

	if !opLEQ(2, 2).(bool) || !opLEQ(int8(2), int8(2)).(bool) ||
		!opLEQ(int16(2), int16(2)).(bool) || !opLEQ(int32(2), int32(2)).(bool) ||
		!opLEQ(int64(2), int64(2)).(bool) || !opLEQ(uint(2), uint(2)).(bool) ||
		!opLEQ(uint8(2), uint8(2)).(bool) || !opLEQ(uint16(2), uint16(2)).(bool) ||
		!opLEQ(uint32(2), uint32(2)).(bool) || !opLEQ(uint64(2), uint64(2)).(bool) ||
		!opLEQ(uintptr(2), uintptr(2)).(bool) || !opLEQ(float32(2), float32(2)).(bool) ||
		!opLEQ(float64(2), float64(2)).(bool) || !opLEQ("b", "b").(bool) {
		t.Fatal("opLEQ")
	}
	if !opLEQ(namedInt(2), namedInt(2)).(bool) {
		t.Fatal("opLEQ named")
	}

	if !opGTR(3, 2).(bool) || !opGTR(int8(3), int8(2)).(bool) ||
		!opGTR(int16(3), int16(2)).(bool) || !opGTR(int32(3), int32(2)).(bool) ||
		!opGTR(int64(3), int64(2)).(bool) || !opGTR(uint(3), uint(2)).(bool) ||
		!opGTR(uint8(3), uint8(2)).(bool) || !opGTR(uint16(3), uint16(2)).(bool) ||
		!opGTR(uint32(3), uint32(2)).(bool) || !opGTR(uint64(3), uint64(2)).(bool) ||
		!opGTR(uintptr(3), uintptr(2)).(bool) || !opGTR(float32(3), float32(2)).(bool) ||
		!opGTR(float64(3), float64(2)).(bool) || !opGTR("c", "b").(bool) {
		t.Fatal("opGTR")
	}
	if !opGTR(namedInt(3), namedInt(2)).(bool) {
		t.Fatal("opGTR named")
	}

	if !opGEQ(2, 2).(bool) || !opGEQ(int8(2), int8(2)).(bool) ||
		!opGEQ(int16(2), int16(2)).(bool) || !opGEQ(int32(2), int32(2)).(bool) ||
		!opGEQ(int64(2), int64(2)).(bool) || !opGEQ(uint(2), uint(2)).(bool) ||
		!opGEQ(uint8(2), uint8(2)).(bool) || !opGEQ(uint16(2), uint16(2)).(bool) ||
		!opGEQ(uint32(2), uint32(2)).(bool) || !opGEQ(uint64(2), uint64(2)).(bool) ||
		!opGEQ(uintptr(2), uintptr(2)).(bool) || !opGEQ(float32(2), float32(2)).(bool) ||
		!opGEQ(float64(2), float64(2)).(bool) || !opGEQ("b", "b").(bool) {
		t.Fatal("opGEQ")
	}
	if !opGEQ(namedInt(2), namedInt(2)).(bool) {
		t.Fatal("opGEQ named")
	}
}

func TestBinopUnopConvert(t *testing.T) {
	add := &ssa.BinOp{Op: token.ADD}
	if binop(add, types.Typ[types.Int], 1, 2).(int) != 3 {
		t.Fatal("binop ADD")
	}
	if binop(&ssa.BinOp{Op: token.SUB}, nil, 5, 2).(int) != 3 {
		t.Fatal("binop SUB")
	}
	if binop(&ssa.BinOp{Op: token.MUL}, nil, 3, 4).(int) != 12 {
		t.Fatal("binop MUL")
	}
	if binop(&ssa.BinOp{Op: token.QUO}, nil, 8, 2).(int) != 4 {
		t.Fatal("binop QUO")
	}
	if binop(&ssa.BinOp{Op: token.REM}, nil, 7, 3).(int) != 1 {
		t.Fatal("binop REM")
	}
	if binop(&ssa.BinOp{Op: token.AND}, nil, 6, 3).(int) != 2 {
		t.Fatal("binop AND")
	}
	if binop(&ssa.BinOp{Op: token.OR}, nil, 4, 1).(int) != 5 {
		t.Fatal("binop OR")
	}
	if binop(&ssa.BinOp{Op: token.XOR}, nil, 6, 3).(int) != 5 {
		t.Fatal("binop XOR")
	}
	if binop(&ssa.BinOp{Op: token.AND_NOT}, nil, 7, 1).(int) != 6 {
		t.Fatal("binop AND_NOT")
	}
	if binop(&ssa.BinOp{Op: token.SHL}, nil, 1, uint(2)).(int) != 4 {
		t.Fatal("binop SHL")
	}
	if binop(&ssa.BinOp{Op: token.SHR}, nil, 8, uint(2)).(int) != 2 {
		t.Fatal("binop SHR")
	}
	if !binop(&ssa.BinOp{Op: token.LSS}, nil, 1, 2).(bool) {
		t.Fatal("binop LSS")
	}
	if !binop(&ssa.BinOp{Op: token.LEQ}, nil, 2, 2).(bool) {
		t.Fatal("binop LEQ")
	}
	if !binop(&ssa.BinOp{Op: token.GTR}, nil, 3, 2).(bool) {
		t.Fatal("binop GTR")
	}
	if !binop(&ssa.BinOp{Op: token.GEQ}, nil, 2, 2).(bool) {
		t.Fatal("binop GEQ")
	}
	eql := &ssa.BinOp{Op: token.EQL}
	if !binop(eql, nil, 2, 2).(bool) {
		t.Fatal("binop EQL")
	}
	if binop(&ssa.BinOp{Op: token.NEQ}, nil, 2, 3).(bool) != true {
		t.Fatal("binop NEQ")
	}
	mustPanic(t, func() { binop(&ssa.BinOp{Op: token.ILLEGAL}, nil, 1, 2) })

	if unop(&ssa.UnOp{Op: token.SUB}, 3).(int) != -3 ||
		unop(&ssa.UnOp{Op: token.SUB}, int8(3)).(int8) != -3 ||
		unop(&ssa.UnOp{Op: token.SUB}, int16(3)).(int16) != -3 ||
		unop(&ssa.UnOp{Op: token.SUB}, int32(3)).(int32) != -3 ||
		unop(&ssa.UnOp{Op: token.SUB}, int64(3)).(int64) != -3 ||
		unop(&ssa.UnOp{Op: token.SUB}, float32(3)).(float32) != -3 ||
		unop(&ssa.UnOp{Op: token.SUB}, float64(3)).(float64) != -3 ||
		unop(&ssa.UnOp{Op: token.SUB}, complex64(3)).(complex64) != -3 ||
		unop(&ssa.UnOp{Op: token.SUB}, complex128(3)).(complex128) != -3 {
		t.Fatal("unop SUB")
	}
	_ = unop(&ssa.UnOp{Op: token.SUB}, uint(3))
	_ = unop(&ssa.UnOp{Op: token.SUB}, uint8(3))
	_ = unop(&ssa.UnOp{Op: token.SUB}, uint16(3))
	_ = unop(&ssa.UnOp{Op: token.SUB}, uint32(3))
	_ = unop(&ssa.UnOp{Op: token.SUB}, uint64(3))
	_ = unop(&ssa.UnOp{Op: token.SUB}, uintptr(3))
	if unop(&ssa.UnOp{Op: token.SUB}, namedInt(3)).(namedInt) != -3 {
		t.Fatal("unop SUB named")
	}
	if unop(&ssa.UnOp{Op: token.NOT}, true).(bool) != false {
		t.Fatal("unop NOT")
	}
	type namedBool bool
	if unop(&ssa.UnOp{Op: token.NOT}, namedBool(false)).(namedBool) != true {
		t.Fatal("unop NOT named")
	}
	if unop(&ssa.UnOp{Op: token.XOR}, 0).(int) != -1 ||
		unop(&ssa.UnOp{Op: token.XOR}, int8(0)).(int8) != -1 ||
		unop(&ssa.UnOp{Op: token.XOR}, int16(0)).(int16) != -1 ||
		unop(&ssa.UnOp{Op: token.XOR}, int32(0)).(int32) != -1 ||
		unop(&ssa.UnOp{Op: token.XOR}, int64(0)).(int64) != -1 ||
		unop(&ssa.UnOp{Op: token.XOR}, uint8(0)).(uint8) != 255 {
		t.Fatal("unop XOR")
	}
	if unop(&ssa.UnOp{Op: token.XOR}, namedInt(0)).(namedInt) != -1 {
		t.Fatal("unop XOR named")
	}
	n := 9
	if unop(&ssa.UnOp{Op: token.MUL}, &n).(int) != 9 {
		t.Fatal("unop MUL")
	}
	ch := make(chan int, 1)
	ch <- 11
	if unop(&ssa.UnOp{Op: token.ARROW}, ch).(int) != 11 {
		t.Fatal("unop ARROW")
	}
	close(ch)
	if unop(&ssa.UnOp{Op: token.ARROW, CommaOk: true}, ch).(tuple)[1] != false {
		t.Fatal("unop ARROW commaok")
	}
	mustPanic(t, func() { unop(&ssa.UnOp{Op: token.ILLEGAL}, 1) })

	if convert(1, reflect.TypeOf(int64(0))).(int64) != 1 {
		t.Fatal("convert int")
	}
	up := uintptr(0)
	if convert(up, reflect.TypeOf(unsafe.Pointer(nil))) == nil {
		t.Fatal("convert uintptr to unsafe")
	}
	p := unsafe.Pointer(&n)
	if convert(p, reflect.TypeOf(uintptr(0))).(uintptr) == 0 {
		t.Fatal("convert unsafe to uintptr")
	}
	pt := reflect.TypeOf((*int)(nil))
	if convert(p, pt) == nil {
		t.Fatal("convert unsafe to ptr")
	}
	type myByte byte
	type myRune rune
	if convert("ab", reflect.TypeOf([]myByte{})).([]myByte)[0] != 'a' {
		t.Fatal("convert string to named []byte")
	}
	if convert("中", reflect.TypeOf([]myRune{})).([]myRune)[0] != '中' {
		t.Fatal("convert string to named []rune")
	}
	if convert([]myByte{'a', 'b'}, reflect.TypeOf("")).(string) != "ab" {
		t.Fatal("convert named []byte to string")
	}
	if convert([]myRune{'中'}, reflect.TypeOf("")).(string) != "中" {
		t.Fatal("convert named []rune to string")
	}
}

func TestEqualWidenIsNil(t *testing.T) {
	if !IsNil(reflect.ValueOf([]int(nil))) || IsNil(reflect.ValueOf(1)) {
		t.Fatal("IsNil")
	}
	if IsConstNil(ssa.NewConst(constant.MakeInt64(1), types.Typ[types.Int])) {
		t.Fatal("IsConstNil int")
	}
	nilc := ssa.NewConst(nil, types.NewPointer(types.Typ[types.Int]))
	if !IsConstNil(nilc) {
		t.Fatal("IsConstNil pointer")
	}
	if IsConstNil(nil) {
		t.Fatal("IsConstNil other")
	}

	a1 := reflect.ValueOf([2]int{1, 2})
	a2 := reflect.ValueOf([2]int{1, 2})
	a3 := reflect.ValueOf([2]int{1, 3})
	if !equalArray(a1, a2) || equalArray(a1, a3) {
		t.Fatal("equalArray")
	}
	if equalArray(a1, reflect.ValueOf([1]int{1})) {
		t.Fatal("equalArray len")
	}
	st := reflect.ValueOf(struct{ A int }{1})
	if !equalStruct(st, st) {
		t.Fatal("equalStruct")
	}
	if !equalValue(a1, a2) || equalValue(a1, a3) {
		t.Fatal("equalValue array")
	}
	p := 1
	if !equalValue(reflect.ValueOf(&p), reflect.ValueOf(&p)) {
		t.Fatal("equalValue ptr")
	}
	ch := make(chan int)
	if !equalValue(reflect.ValueOf(ch), reflect.ValueOf(ch)) {
		t.Fatal("equalValue chan")
	}
	if !equalNil(reflect.ValueOf([]int(nil)), reflect.ValueOf([]int(nil))) {
		t.Fatal("equalNil")
	}

	if widen(true) != true || widen(int64(1)) != int64(1) ||
		widen(int(2)) != int64(2) || widen(int8(2)) != int64(2) ||
		widen(int16(2)) != int64(2) || widen(int32(2)) != int64(2) ||
		widen(uint(2)) != uint64(2) || widen(uint8(2)) != uint64(2) ||
		widen(uint16(2)) != uint64(2) || widen(uint32(2)) != uint64(2) ||
		widen(uintptr(2)) != uint64(2) || widen(float32(2)) != float64(2) ||
		widen(complex64(2)) != complex128(2) || widen("s") != "s" {
		t.Fatal("widen")
	}
	mustPanic(t, func() { widen([]int{}) })

	instr := &ssa.BinOp{Op: token.EQL, X: nilc}
	if !opEQL(instr, (*int)(nil), (*int)(nil)) {
		t.Fatal("opEQL nil X")
	}
	instr.X, instr.Y = ssa.NewConst(constant.MakeInt64(1), types.Typ[types.Int]), nilc
	if !opEQL(instr, (*int)(nil), (*int)(nil)) {
		t.Fatal("opEQL nil Y")
	}
}

func TestFieldXAndErrors(t *testing.T) {
	type T struct{ A, B int }
	v := &T{1, 2}
	p, err := fieldAddrX(v, 0)
	if err != nil || *p.(*int) != 1 {
		t.Fatal("fieldAddrX", p, err)
	}
	got, err := fieldX(v, 1)
	if err != nil || got.(int) != 2 {
		t.Fatal("fieldX", got, err)
	}
	if _, err := fieldAddrX((*T)(nil), 0); err == nil {
		t.Fatal("fieldAddrX nil")
	}
	if _, err := fieldX((*T)(nil), 0); err == nil {
		t.Fatal("fieldX nil")
	}
	if _, err := fieldAddrAt((*T)(nil), 0, reflect.TypeOf(0)); err == nil {
		t.Fatal("fieldAddrAt nil")
	}
	ms := allMethodX(reflect.TypeOf(v))
	_ = ms

	if (ExitError(2)).Error() == "" {
		t.Fatal("ExitError")
	}
	if (PlainError("p")).Error() != "p" {
		t.Fatal("PlainError")
	}
	(PlainError("p")).RuntimeError()
	if (RuntimeError("r")).Error() != "runtime error: r" {
		t.Fatal("RuntimeError")
	}
	(RuntimeError("r")).RuntimeError()
	pe := PanicError{Value: 1}
	if pe.Error() == "" || pe.Stack() != nil {
		t.Fatal("PanicError")
	}
	fe := FatalError{Value: "x"}
	if fe.Error() == "" {
		t.Fatal("FatalError")
	}
}

func TestWritevalueNamed(t *testing.T) {
	var buf bytes.Buffer
	writevalue(&buf, namedFloat(1.5), true)
	writevalue(&buf, namedComplex(1+2i), true)
	writeany(&buf, namedInt(3))
	writeany(&buf, namedFloat(1.25))
	writeany(&buf, namedComplex(3+4i))
	writeany(&buf, namedString("s"))
	writevalue(&buf, complex128(1+2i), true)
	writevalue(&buf, complex64(1+2i), true)
	writevalue(&buf, float32(1.5), true)
	writevalue(&buf, float64(1.5), true)
	var inner interface{ M() }
	var outer interface{} = inner
	writevalue(&buf, outer, true)
	type namedUintptr uintptr
	writevalue(&buf, namedUintptr(1), true)
	writeany(&buf, error(PlainError("e")))
	writeinterface(&buf, 1)
	writevalue(&buf, []int{1}, false)
	writevalue(&buf, map[int]int{1: 1}, false)
	writevalue(&buf, (func())(nil), true)
	writevalue(&buf, (chan int)(nil), true)
	writevalue(&buf, unsafe.Pointer(nil), true)
	writevalue(&buf, [2]int{1, 2}, true)
	writevalue(&buf, namedString("s"), true)
	mustPanic(t, func() {
		var b bytes.Buffer
		writevalue(&b, struct{ A int }{1}, false)
	})
	writevalue(&buf, struct{ A int }{1}, true)
}

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	fn()
}

func TestEmptyTypeAndOffsetof(t *testing.T) {
	for _, kind := range []reflect.Kind{
		reflect.Array, reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Map, reflect.Ptr, reflect.Slice, reflect.Struct, reflect.String,
	} {
		if emptyType(kind) == nil {
			t.Fatalf("emptyType %v", kind)
		}
	}
	_ = emptyType(reflect.Int)

	pkg := types.NewPackage("p", "p")
	x := types.NewField(token.NoPos, pkg, "X", types.Typ[types.Int], false)
	inner := types.NewNamed(types.NewTypeName(token.NoPos, pkg, "Inner", nil), types.NewStruct([]*types.Var{x}, nil), nil)
	emb := types.NewField(token.NoPos, pkg, "Inner", inner, true)
	outer := types.NewStruct([]*types.Var{emb}, nil)
	sizes := types.SizesFor("gc", "amd64")
	off, err := selOffsetof(sizes, outer, []int{0, 0}, "X")
	if err != nil || off != 0 {
		t.Fatalf("selOffsetof %v %v", off, err)
	}
	ptrEmb := types.NewField(token.NoPos, pkg, "Inner", types.NewPointer(inner), true)
	ptrOuter := types.NewStruct([]*types.Var{ptrEmb}, nil)
	if _, err := selOffsetof(sizes, ptrOuter, []int{0, 0}, "X"); err == nil {
		t.Fatal("expected pointer embed error")
	}

	recv := types.NewVar(token.NoPos, pkg, "r", types.Typ[types.Int])
	methSig := types.NewSignatureType(recv, nil, nil, nil, nil, false)
	meth := types.NewFunc(token.NoPos, pkg, "M", methSig)
	iface := types.NewInterfaceType([]*types.Func{meth}, nil)
	iface.Complete()
	param := types.NewVar(token.NoPos, pkg, overloadArgs, iface)
	sig := types.NewSignatureType(nil, nil, nil, types.NewTuple(param), nil, false)
	if !isXGoOverloadFunc(sig) {
		t.Fatal("isXGoOverloadFunc")
	}
	if isXGoOverloadFunc(types.NewSignatureType(nil, nil, nil, nil, nil, false)) {
		t.Fatal("empty sig")
	}
	_ = loc(token.NewFileSet(), token.NoPos)
	fset := token.NewFileSet()
	fset.AddFile("x.go", 1, 32)
	_ = loc(fset, token.Pos(2))
}
