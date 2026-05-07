// export by github.com/goplus/ixgo/cmd/qexp

package abi

import (
	q "github.com/goplus/ixgo/x/abi"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/ixgo/x/abi", func() *ixgo.Package {
		return &ixgo.Package{
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
				"Method":         reflect.TypeOf((*q.Method)(nil)).Elem(),
				"Name":           reflect.TypeOf((*q.Name)(nil)).Elem(),
				"NameOff":        reflect.TypeOf((*q.NameOff)(nil)).Elem(),
				"OldMapType":     reflect.TypeOf((*q.OldMapType)(nil)).Elem(),
				"PtrType":        reflect.TypeOf((*q.PtrType)(nil)).Elem(),
				"SliceType":      reflect.TypeOf((*q.SliceType)(nil)).Elem(),
				"StructField":    reflect.TypeOf((*q.StructField)(nil)).Elem(),
				"StructType":     reflect.TypeOf((*q.StructType)(nil)).Elem(),
				"SwissMapType":   reflect.TypeOf((*q.SwissMapType)(nil)).Elem(),
				"TFlag":          reflect.TypeOf((*q.TFlag)(nil)).Elem(),
				"TextOff":        reflect.TypeOf((*q.TextOff)(nil)).Elem(),
				"Type":           reflect.TypeOf((*q.Type)(nil)).Elem(),
				"TypeOff":        reflect.TypeOf((*q.TypeOff)(nil)).Elem(),
				"UncommonType":   reflect.TypeOf((*q.UncommonType)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewName":  q.NewName,
				"NoEscape": q.NoEscape,
				"TypeOf":   q.TypeOf,
			},
			TypedConsts: map[string]interface{}{
				"Array":               q.Array,
				"Bool":                q.Bool,
				"BothDir":             q.BothDir,
				"Chan":                q.Chan,
				"Complex128":          q.Complex128,
				"Complex64":           q.Complex64,
				"Float32":             q.Float32,
				"Float64":             q.Float64,
				"Func":                q.Func,
				"Int":                 q.Int,
				"Int16":               q.Int16,
				"Int32":               q.Int32,
				"Int64":               q.Int64,
				"Int8":                q.Int8,
				"Interface":           q.Interface,
				"Invalid":             q.Invalid,
				"InvalidDir":          q.InvalidDir,
				"KindDirectIface":     q.KindDirectIface,
				"KindMask":            q.KindMask,
				"Map":                 q.Map,
				"Pointer":             q.Pointer,
				"RecvDir":             q.RecvDir,
				"SendDir":             q.SendDir,
				"Slice":               q.Slice,
				"String":              q.String,
				"Struct":              q.Struct,
				"SwissMapCtrlEmpty":   q.SwissMapCtrlEmpty,
				"TFlagExtraStar":      q.TFlagExtraStar,
				"TFlagGCMaskOnDemand": q.TFlagGCMaskOnDemand,
				"TFlagNamed":          q.TFlagNamed,
				"TFlagRegularMemory":  q.TFlagRegularMemory,
				"TFlagUncommon":       q.TFlagUncommon,
				"Uint":                q.Uint,
				"Uint16":              q.Uint16,
				"Uint32":              q.Uint32,
				"Uint64":              q.Uint64,
				"Uint8":               q.Uint8,
				"Uintptr":             q.Uintptr,
				"UnsafePointer":       q.UnsafePointer,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"MaxPtrmaskBytes":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MaxPtrmaskBytes))},
				"OldMapBucketCount":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.OldMapBucketCount))},
				"OldMapBucketCountBits":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.OldMapBucketCountBits))},
				"OldMapMaxElemBytes":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.OldMapMaxElemBytes))},
				"OldMapMaxKeyBytes":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.OldMapMaxKeyBytes))},
				"SwissMapGroupSlots":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapGroupSlots))},
				"SwissMapGroupSlotsBits":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapGroupSlotsBits))},
				"SwissMapHashMightPanic":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapHashMightPanic))},
				"SwissMapIndirectElem":    {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapIndirectElem))},
				"SwissMapIndirectKey":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapIndirectKey))},
				"SwissMapMaxElemBytes":    {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapMaxElemBytes))},
				"SwissMapMaxKeyBytes":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapMaxKeyBytes))},
				"SwissMapNeedKeyUpdate":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SwissMapNeedKeyUpdate))},
				"TraceArgsDotdotdot":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TraceArgsDotdotdot))},
				"TraceArgsEndAgg":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TraceArgsEndAgg))},
				"TraceArgsEndSeq":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TraceArgsEndSeq))},
				"TraceArgsLimit":          {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TraceArgsLimit))},
				"TraceArgsMaxDepth":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TraceArgsMaxDepth))},
				"TraceArgsMaxLen":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TraceArgsMaxLen))},
				"TraceArgsOffsetTooLarge": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TraceArgsOffsetTooLarge))},
				"TraceArgsSpecial":        {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TraceArgsSpecial))},
				"TraceArgsStartAgg":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TraceArgsStartAgg))},
			},
		}
	})
}
