// export by github.com/goplus/ixgo/cmd/qexp

package abi

import (
	q "github.com/goplus/ixgo/x/abi"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "abi",
		Path: "github.com/goplus/ixgo/x/abi",
		Deps: map[string]string{
			"unsafe": "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"EmptyInterface": reflect.TypeOf((*q.EmptyInterface)(nil)).Elem(),
			"ITab":           reflect.TypeOf((*q.ITab)(nil)).Elem(),
			"Kind":           reflect.TypeOf((*q.Kind)(nil)).Elem(),
			"NameOff":        reflect.TypeOf((*q.NameOff)(nil)).Elem(),
			"TFlag":          reflect.TypeOf((*q.TFlag)(nil)).Elem(),
			"TextOff":        reflect.TypeOf((*q.TextOff)(nil)).Elem(),
			"Type":           reflect.TypeOf((*q.Type)(nil)).Elem(),
			"TypeOff":        reflect.TypeOf((*q.TypeOff)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NoEscape": reflect.ValueOf(q.NoEscape),
			"TypeOf":   reflect.ValueOf(q.TypeOf),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"Array":               {reflect.TypeOf(q.Array), constant.MakeInt64(int64(q.Array))},
			"Bool":                {reflect.TypeOf(q.Bool), constant.MakeInt64(int64(q.Bool))},
			"Chan":                {reflect.TypeOf(q.Chan), constant.MakeInt64(int64(q.Chan))},
			"Complex128":          {reflect.TypeOf(q.Complex128), constant.MakeInt64(int64(q.Complex128))},
			"Complex64":           {reflect.TypeOf(q.Complex64), constant.MakeInt64(int64(q.Complex64))},
			"Float32":             {reflect.TypeOf(q.Float32), constant.MakeInt64(int64(q.Float32))},
			"Float64":             {reflect.TypeOf(q.Float64), constant.MakeInt64(int64(q.Float64))},
			"Func":                {reflect.TypeOf(q.Func), constant.MakeInt64(int64(q.Func))},
			"Int":                 {reflect.TypeOf(q.Int), constant.MakeInt64(int64(q.Int))},
			"Int16":               {reflect.TypeOf(q.Int16), constant.MakeInt64(int64(q.Int16))},
			"Int32":               {reflect.TypeOf(q.Int32), constant.MakeInt64(int64(q.Int32))},
			"Int64":               {reflect.TypeOf(q.Int64), constant.MakeInt64(int64(q.Int64))},
			"Int8":                {reflect.TypeOf(q.Int8), constant.MakeInt64(int64(q.Int8))},
			"Interface":           {reflect.TypeOf(q.Interface), constant.MakeInt64(int64(q.Interface))},
			"Invalid":             {reflect.TypeOf(q.Invalid), constant.MakeInt64(int64(q.Invalid))},
			"KindDirectIface":     {reflect.TypeOf(q.KindDirectIface), constant.MakeInt64(int64(q.KindDirectIface))},
			"KindMask":            {reflect.TypeOf(q.KindMask), constant.MakeInt64(int64(q.KindMask))},
			"Map":                 {reflect.TypeOf(q.Map), constant.MakeInt64(int64(q.Map))},
			"Pointer":             {reflect.TypeOf(q.Pointer), constant.MakeInt64(int64(q.Pointer))},
			"Slice":               {reflect.TypeOf(q.Slice), constant.MakeInt64(int64(q.Slice))},
			"String":              {reflect.TypeOf(q.String), constant.MakeInt64(int64(q.String))},
			"Struct":              {reflect.TypeOf(q.Struct), constant.MakeInt64(int64(q.Struct))},
			"TFlagExtraStar":      {reflect.TypeOf(q.TFlagExtraStar), constant.MakeInt64(int64(q.TFlagExtraStar))},
			"TFlagGCMaskOnDemand": {reflect.TypeOf(q.TFlagGCMaskOnDemand), constant.MakeInt64(int64(q.TFlagGCMaskOnDemand))},
			"TFlagNamed":          {reflect.TypeOf(q.TFlagNamed), constant.MakeInt64(int64(q.TFlagNamed))},
			"TFlagRegularMemory":  {reflect.TypeOf(q.TFlagRegularMemory), constant.MakeInt64(int64(q.TFlagRegularMemory))},
			"TFlagUncommon":       {reflect.TypeOf(q.TFlagUncommon), constant.MakeInt64(int64(q.TFlagUncommon))},
			"Uint":                {reflect.TypeOf(q.Uint), constant.MakeInt64(int64(q.Uint))},
			"Uint16":              {reflect.TypeOf(q.Uint16), constant.MakeInt64(int64(q.Uint16))},
			"Uint32":              {reflect.TypeOf(q.Uint32), constant.MakeInt64(int64(q.Uint32))},
			"Uint64":              {reflect.TypeOf(q.Uint64), constant.MakeInt64(int64(q.Uint64))},
			"Uint8":               {reflect.TypeOf(q.Uint8), constant.MakeInt64(int64(q.Uint8))},
			"Uintptr":             {reflect.TypeOf(q.Uintptr), constant.MakeInt64(int64(q.Uintptr))},
			"UnsafePointer":       {reflect.TypeOf(q.UnsafePointer), constant.MakeInt64(int64(q.UnsafePointer))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
		Source:        source,
	})
}

var source = "package abi\n\nimport \"unsafe\"\n\nfunc NoEscape(p unsafe.Pointer) unsafe.Pointer\n\nvar alwaysFalse bool\nvar escapeSink any\n\nfunc Escape[T any](x T) T {\n\tif alwaysFalse {\n\t\tescapeSink = x\n\t}\n\treturn x\n}\n\ntype ITab struct {\n\tInter\t*Type\n\tType\t*Type\n\tHash\tuint32\n\tFun\t[1]uintptr\n}\n\ntype EmptyInterface struct {\n\tType\t*Type\n\tData\tunsafe.Pointer\n}\ntype Type struct {\n\tSize_\t\tuintptr\n\tPtrBytes\tuintptr\n\tHash\t\tuint32\n\tTFlag\t\tTFlag\n\tAlign_\t\tuint8\n\tFieldAlign_\tuint8\n\tKind_\t\tKind\n\n\tEqual\tfunc(unsafe.Pointer, unsafe.Pointer) bool\n\n\tGCData\t\t*byte\n\tStr\t\tNameOff\n\tPtrToThis\tTypeOff\n}\n\ntype Kind uint8\n\nconst (\n\tInvalid\tKind\t= iota\n\tBool\n\tInt\n\tInt8\n\tInt16\n\tInt32\n\tInt64\n\tUint\n\tUint8\n\tUint16\n\tUint32\n\tUint64\n\tUintptr\n\tFloat32\n\tFloat64\n\tComplex64\n\tComplex128\n\tArray\n\tChan\n\tFunc\n\tInterface\n\tMap\n\tPointer\n\tSlice\n\tString\n\tStruct\n\tUnsafePointer\n)\n\nconst (\n\tKindDirectIface\tKind\t= 1 << 5\n\tKindMask\tKind\t= (1 << 5) - 1\n)\n\ntype TFlag uint8\n\nconst (\n\tTFlagUncommon\tTFlag\t= 1 << 0\n\n\tTFlagExtraStar\tTFlag\t= 1 << 1\n\n\tTFlagNamed\tTFlag\t= 1 << 2\n\n\tTFlagRegularMemory\tTFlag\t= 1 << 3\n\n\tTFlagGCMaskOnDemand\tTFlag\t= 1 << 4\n)\n\ntype NameOff int32\n\ntype TypeOff int32\n\ntype TextOff int32\n\nfunc (k Kind) String() string\n\nvar kindNames = []string{\n\tInvalid:\t\"invalid\",\n\tBool:\t\t\"bool\",\n\tInt:\t\t\"int\",\n\tInt8:\t\t\"int8\",\n\tInt16:\t\t\"int16\",\n\tInt32:\t\t\"int32\",\n\tInt64:\t\t\"int64\",\n\tUint:\t\t\"uint\",\n\tUint8:\t\t\"uint8\",\n\tUint16:\t\t\"uint16\",\n\tUint32:\t\t\"uint32\",\n\tUint64:\t\t\"uint64\",\n\tUintptr:\t\"uintptr\",\n\tFloat32:\t\"float32\",\n\tFloat64:\t\"float64\",\n\tComplex64:\t\"complex64\",\n\tComplex128:\t\"complex128\",\n\tArray:\t\t\"array\",\n\tChan:\t\t\"chan\",\n\tFunc:\t\t\"func\",\n\tInterface:\t\"interface\",\n\tMap:\t\t\"map\",\n\tPointer:\t\"ptr\",\n\tSlice:\t\t\"slice\",\n\tString:\t\t\"string\",\n\tStruct:\t\t\"struct\",\n\tUnsafePointer:\t\"unsafe.Pointer\",\n}\n\nfunc TypeOf(a any) *Type\n\nfunc (t *Type) Kind() Kind\n\nfunc (t *Type) HasName() bool\n\nfunc (t *Type) Pointers() bool\n\nfunc (t *Type) IfaceIndir() bool\n\nfunc (t *Type) IsDirectIface() bool\n\nfunc (t *Type) GcSlice(begin, end uintptr) []byte\n"
