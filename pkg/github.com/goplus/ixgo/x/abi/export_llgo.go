// export by github.com/goplus/ixgo/cmd/qexp

//go:build llgo
// +build llgo

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
			"ArrayType":      reflect.TypeOf((*q.ArrayType)(nil)).Elem(),
			"ChanDir":        reflect.TypeOf((*q.ChanDir)(nil)).Elem(),
			"ChanType":       reflect.TypeOf((*q.ChanType)(nil)).Elem(),
			"EmptyInterface": reflect.TypeOf((*q.EmptyInterface)(nil)).Elem(),
			"FuncType":       reflect.TypeOf((*q.FuncType)(nil)).Elem(),
			"ITab":           reflect.TypeOf((*q.ITab)(nil)).Elem(),
			"Imethod":        reflect.TypeOf((*q.Imethod)(nil)).Elem(),
			"InterfaceType":  reflect.TypeOf((*q.InterfaceType)(nil)).Elem(),
			"Kind":           reflect.TypeOf((*q.Kind)(nil)).Elem(),
			"MapType":        reflect.TypeOf((*q.MapType)(nil)).Elem(),
			"Method":         reflect.TypeOf((*q.Method)(nil)).Elem(),
			"OldMapType":     reflect.TypeOf((*q.OldMapType)(nil)).Elem(),
			"PtrType":        reflect.TypeOf((*q.PtrType)(nil)).Elem(),
			"SliceType":      reflect.TypeOf((*q.SliceType)(nil)).Elem(),
			"StructField":    reflect.TypeOf((*q.StructField)(nil)).Elem(),
			"StructType":     reflect.TypeOf((*q.StructType)(nil)).Elem(),
			"SwissMapType":   reflect.TypeOf((*q.SwissMapType)(nil)).Elem(),
			"TFlag":          reflect.TypeOf((*q.TFlag)(nil)).Elem(),
			"Type":           reflect.TypeOf((*q.Type)(nil)).Elem(),
			"UncommonType":   reflect.TypeOf((*q.UncommonType)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Text": reflect.TypeOf((*q.Text)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"IsExported": reflect.ValueOf(q.IsExported),
			"NoEscape":   reflect.ValueOf(q.NoEscape),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"Array":               {Typ: reflect.TypeOf(q.Array), Value: constant.MakeInt64(int64(q.Array))},
			"Bool":                {Typ: reflect.TypeOf(q.Bool), Value: constant.MakeInt64(int64(q.Bool))},
			"BothDir":             {Typ: reflect.TypeOf(q.BothDir), Value: constant.MakeInt64(int64(q.BothDir))},
			"Chan":                {Typ: reflect.TypeOf(q.Chan), Value: constant.MakeInt64(int64(q.Chan))},
			"Complex128":          {Typ: reflect.TypeOf(q.Complex128), Value: constant.MakeInt64(int64(q.Complex128))},
			"Complex64":           {Typ: reflect.TypeOf(q.Complex64), Value: constant.MakeInt64(int64(q.Complex64))},
			"Float32":             {Typ: reflect.TypeOf(q.Float32), Value: constant.MakeInt64(int64(q.Float32))},
			"Float64":             {Typ: reflect.TypeOf(q.Float64), Value: constant.MakeInt64(int64(q.Float64))},
			"Func":                {Typ: reflect.TypeOf(q.Func), Value: constant.MakeInt64(int64(q.Func))},
			"Int":                 {Typ: reflect.TypeOf(q.Int), Value: constant.MakeInt64(int64(q.Int))},
			"Int16":               {Typ: reflect.TypeOf(q.Int16), Value: constant.MakeInt64(int64(q.Int16))},
			"Int32":               {Typ: reflect.TypeOf(q.Int32), Value: constant.MakeInt64(int64(q.Int32))},
			"Int64":               {Typ: reflect.TypeOf(q.Int64), Value: constant.MakeInt64(int64(q.Int64))},
			"Int8":                {Typ: reflect.TypeOf(q.Int8), Value: constant.MakeInt64(int64(q.Int8))},
			"Interface":           {Typ: reflect.TypeOf(q.Interface), Value: constant.MakeInt64(int64(q.Interface))},
			"Invalid":             {Typ: reflect.TypeOf(q.Invalid), Value: constant.MakeInt64(int64(q.Invalid))},
			"InvalidDir":          {Typ: reflect.TypeOf(q.InvalidDir), Value: constant.MakeInt64(int64(q.InvalidDir))},
			"Map":                 {Typ: reflect.TypeOf(q.Map), Value: constant.MakeInt64(int64(q.Map))},
			"Pointer":             {Typ: reflect.TypeOf(q.Pointer), Value: constant.MakeInt64(int64(q.Pointer))},
			"RecvDir":             {Typ: reflect.TypeOf(q.RecvDir), Value: constant.MakeInt64(int64(q.RecvDir))},
			"SendDir":             {Typ: reflect.TypeOf(q.SendDir), Value: constant.MakeInt64(int64(q.SendDir))},
			"Slice":               {Typ: reflect.TypeOf(q.Slice), Value: constant.MakeInt64(int64(q.Slice))},
			"String":              {Typ: reflect.TypeOf(q.String), Value: constant.MakeInt64(int64(q.String))},
			"Struct":              {Typ: reflect.TypeOf(q.Struct), Value: constant.MakeInt64(int64(q.Struct))},
			"SwissMapCtrlEmpty":   {Typ: reflect.TypeOf(q.SwissMapCtrlEmpty), Value: constant.MakeUint64(uint64(q.SwissMapCtrlEmpty))},
			"TFlagClosure":        {Typ: reflect.TypeOf(q.TFlagClosure), Value: constant.MakeInt64(int64(q.TFlagClosure))},
			"TFlagExtraStar":      {Typ: reflect.TypeOf(q.TFlagExtraStar), Value: constant.MakeInt64(int64(q.TFlagExtraStar))},
			"TFlagGCMaskOnDemand": {Typ: reflect.TypeOf(q.TFlagGCMaskOnDemand), Value: constant.MakeInt64(int64(q.TFlagGCMaskOnDemand))},
			"TFlagNamed":          {Typ: reflect.TypeOf(q.TFlagNamed), Value: constant.MakeInt64(int64(q.TFlagNamed))},
			"TFlagRegularMemory":  {Typ: reflect.TypeOf(q.TFlagRegularMemory), Value: constant.MakeInt64(int64(q.TFlagRegularMemory))},
			"TFlagUncommon":       {Typ: reflect.TypeOf(q.TFlagUncommon), Value: constant.MakeInt64(int64(q.TFlagUncommon))},
			"TFlagVariadic":       {Typ: reflect.TypeOf(q.TFlagVariadic), Value: constant.MakeInt64(int64(q.TFlagVariadic))},
			"Uint":                {Typ: reflect.TypeOf(q.Uint), Value: constant.MakeInt64(int64(q.Uint))},
			"Uint16":              {Typ: reflect.TypeOf(q.Uint16), Value: constant.MakeInt64(int64(q.Uint16))},
			"Uint32":              {Typ: reflect.TypeOf(q.Uint32), Value: constant.MakeInt64(int64(q.Uint32))},
			"Uint64":              {Typ: reflect.TypeOf(q.Uint64), Value: constant.MakeInt64(int64(q.Uint64))},
			"Uint8":               {Typ: reflect.TypeOf(q.Uint8), Value: constant.MakeInt64(int64(q.Uint8))},
			"Uintptr":             {Typ: reflect.TypeOf(q.Uintptr), Value: constant.MakeInt64(int64(q.Uintptr))},
			"UnsafePointer":       {Typ: reflect.TypeOf(q.UnsafePointer), Value: constant.MakeInt64(int64(q.UnsafePointer))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"KindDirectIface":        {Typ: "untyped int", Value: constant.MakeInt64(int64(q.KindDirectIface))},
			"KindGCProg":             {Typ: "untyped int", Value: constant.MakeInt64(int64(q.KindGCProg))},
			"KindMask":               {Typ: "untyped int", Value: constant.MakeInt64(int64(q.KindMask))},
			"MaxPtrmaskBytes":        {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MaxPtrmaskBytes))},
			"OldMapBucketCount":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.OldMapBucketCount))},
			"OldMapBucketCountBits":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.OldMapBucketCountBits))},
			"OldMapMaxElemBytes":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.OldMapMaxElemBytes))},
			"OldMapMaxKeyBytes":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.OldMapMaxKeyBytes))},
			"SwissMapGroupSlots":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapGroupSlots))},
			"SwissMapGroupSlotsBits": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapGroupSlotsBits))},
			"SwissMapHashMightPanic": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapHashMightPanic))},
			"SwissMapIndirectElem":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapIndirectElem))},
			"SwissMapIndirectKey":    {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapIndirectKey))},
			"SwissMapMaxElemBytes":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapMaxElemBytes))},
			"SwissMapMaxKeyBytes":    {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapMaxKeyBytes))},
			"SwissMapNeedKeyUpdate":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapNeedKeyUpdate))},
		},
	})
}
