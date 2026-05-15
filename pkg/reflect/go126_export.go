// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package reflect

import (
	q "reflect"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "reflect",
		Path: "reflect",
		Deps: map[string]string{
			"errors":                "errors",
			"internal/abi":          "abi",
			"internal/bytealg":      "bytealg",
			"internal/goarch":       "goarch",
			"internal/race":         "race",
			"internal/runtime/maps": "maps",
			"internal/runtime/sys":  "sys",
			"internal/strconv":      "strconv",
			"internal/unsafeheader": "unsafeheader",
			"iter":                  "iter",
			"math":                  "math",
			"runtime":               "runtime",
			"strconv":               "strconv",
			"sync":                  "sync",
			"unicode":               "unicode",
			"unicode/utf8":          "utf8",
			"unsafe":                "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Type": reflect.TypeOf((*q.Type)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"ChanDir":      reflect.TypeOf((*q.ChanDir)(nil)).Elem(),
			"Kind":         reflect.TypeOf((*q.Kind)(nil)).Elem(),
			"MapIter":      reflect.TypeOf((*q.MapIter)(nil)).Elem(),
			"Method":       reflect.TypeOf((*q.Method)(nil)).Elem(),
			"SelectCase":   reflect.TypeOf((*q.SelectCase)(nil)).Elem(),
			"SelectDir":    reflect.TypeOf((*q.SelectDir)(nil)).Elem(),
			"SliceHeader":  reflect.TypeOf((*q.SliceHeader)(nil)).Elem(),
			"StringHeader": reflect.TypeOf((*q.StringHeader)(nil)).Elem(),
			"StructField":  reflect.TypeOf((*q.StructField)(nil)).Elem(),
			"StructTag":    reflect.TypeOf((*q.StructTag)(nil)).Elem(),
			"Value":        reflect.TypeOf((*q.Value)(nil)).Elem(),
			"ValueError":   reflect.TypeOf((*q.ValueError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Append":          reflect.ValueOf(q.Append),
			"AppendSlice":     reflect.ValueOf(q.AppendSlice),
			"ArrayOf":         reflect.ValueOf(q.ArrayOf),
			"ChanOf":          reflect.ValueOf(q.ChanOf),
			"Copy":            reflect.ValueOf(q.Copy),
			"DeepEqual":       reflect.ValueOf(q.DeepEqual),
			"FuncOf":          reflect.ValueOf(q.FuncOf),
			"Indirect":        reflect.ValueOf(q.Indirect),
			"MakeChan":        reflect.ValueOf(q.MakeChan),
			"MakeFunc":        reflect.ValueOf(q.MakeFunc),
			"MakeMap":         reflect.ValueOf(q.MakeMap),
			"MakeMapWithSize": reflect.ValueOf(q.MakeMapWithSize),
			"MakeSlice":       reflect.ValueOf(q.MakeSlice),
			"MapOf":           reflect.ValueOf(q.MapOf),
			"New":             reflect.ValueOf(q.New),
			"NewAt":           reflect.ValueOf(q.NewAt),
			"PointerTo":       reflect.ValueOf(q.PointerTo),
			"PtrTo":           reflect.ValueOf(q.PtrTo),
			"Select":          reflect.ValueOf(q.Select),
			"SliceAt":         reflect.ValueOf(q.SliceAt),
			"SliceOf":         reflect.ValueOf(q.SliceOf),
			"StructOf":        reflect.ValueOf(q.StructOf),
			"Swapper":         reflect.ValueOf(q.Swapper),
			"TypeOf":          reflect.ValueOf(q.TypeOf),
			"ValueOf":         reflect.ValueOf(q.ValueOf),
			"VisibleFields":   reflect.ValueOf(q.VisibleFields),
			"Zero":            reflect.ValueOf(q.Zero),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"Array":         {Typ: reflect.TypeOf(q.Array), Value: constant.MakeInt64(int64(q.Array))},
			"Bool":          {Typ: reflect.TypeOf(q.Bool), Value: constant.MakeInt64(int64(q.Bool))},
			"BothDir":       {Typ: reflect.TypeOf(q.BothDir), Value: constant.MakeInt64(int64(q.BothDir))},
			"Chan":          {Typ: reflect.TypeOf(q.Chan), Value: constant.MakeInt64(int64(q.Chan))},
			"Complex128":    {Typ: reflect.TypeOf(q.Complex128), Value: constant.MakeInt64(int64(q.Complex128))},
			"Complex64":     {Typ: reflect.TypeOf(q.Complex64), Value: constant.MakeInt64(int64(q.Complex64))},
			"Float32":       {Typ: reflect.TypeOf(q.Float32), Value: constant.MakeInt64(int64(q.Float32))},
			"Float64":       {Typ: reflect.TypeOf(q.Float64), Value: constant.MakeInt64(int64(q.Float64))},
			"Func":          {Typ: reflect.TypeOf(q.Func), Value: constant.MakeInt64(int64(q.Func))},
			"Int":           {Typ: reflect.TypeOf(q.Int), Value: constant.MakeInt64(int64(q.Int))},
			"Int16":         {Typ: reflect.TypeOf(q.Int16), Value: constant.MakeInt64(int64(q.Int16))},
			"Int32":         {Typ: reflect.TypeOf(q.Int32), Value: constant.MakeInt64(int64(q.Int32))},
			"Int64":         {Typ: reflect.TypeOf(q.Int64), Value: constant.MakeInt64(int64(q.Int64))},
			"Int8":          {Typ: reflect.TypeOf(q.Int8), Value: constant.MakeInt64(int64(q.Int8))},
			"Interface":     {Typ: reflect.TypeOf(q.Interface), Value: constant.MakeInt64(int64(q.Interface))},
			"Invalid":       {Typ: reflect.TypeOf(q.Invalid), Value: constant.MakeInt64(int64(q.Invalid))},
			"Map":           {Typ: reflect.TypeOf(q.Map), Value: constant.MakeInt64(int64(q.Map))},
			"Pointer":       {Typ: reflect.TypeOf(q.Pointer), Value: constant.MakeInt64(int64(q.Pointer))},
			"Ptr":           {Typ: reflect.TypeOf(q.Ptr), Value: constant.MakeInt64(int64(q.Ptr))},
			"RecvDir":       {Typ: reflect.TypeOf(q.RecvDir), Value: constant.MakeInt64(int64(q.RecvDir))},
			"SelectDefault": {Typ: reflect.TypeOf(q.SelectDefault), Value: constant.MakeInt64(int64(q.SelectDefault))},
			"SelectRecv":    {Typ: reflect.TypeOf(q.SelectRecv), Value: constant.MakeInt64(int64(q.SelectRecv))},
			"SelectSend":    {Typ: reflect.TypeOf(q.SelectSend), Value: constant.MakeInt64(int64(q.SelectSend))},
			"SendDir":       {Typ: reflect.TypeOf(q.SendDir), Value: constant.MakeInt64(int64(q.SendDir))},
			"Slice":         {Typ: reflect.TypeOf(q.Slice), Value: constant.MakeInt64(int64(q.Slice))},
			"String":        {Typ: reflect.TypeOf(q.String), Value: constant.MakeInt64(int64(q.String))},
			"Struct":        {Typ: reflect.TypeOf(q.Struct), Value: constant.MakeInt64(int64(q.Struct))},
			"Uint":          {Typ: reflect.TypeOf(q.Uint), Value: constant.MakeInt64(int64(q.Uint))},
			"Uint16":        {Typ: reflect.TypeOf(q.Uint16), Value: constant.MakeInt64(int64(q.Uint16))},
			"Uint32":        {Typ: reflect.TypeOf(q.Uint32), Value: constant.MakeInt64(int64(q.Uint32))},
			"Uint64":        {Typ: reflect.TypeOf(q.Uint64), Value: constant.MakeInt64(int64(q.Uint64))},
			"Uint8":         {Typ: reflect.TypeOf(q.Uint8), Value: constant.MakeInt64(int64(q.Uint8))},
			"Uintptr":       {Typ: reflect.TypeOf(q.Uintptr), Value: constant.MakeInt64(int64(q.Uintptr))},
			"UnsafePointer": {Typ: reflect.TypeOf(q.UnsafePointer), Value: constant.MakeInt64(int64(q.UnsafePointer))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
