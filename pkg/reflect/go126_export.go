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
			"Array":         {reflect.TypeOf(q.Array), constant.MakeInt64(int64(q.Array))},
			"Bool":          {reflect.TypeOf(q.Bool), constant.MakeInt64(int64(q.Bool))},
			"BothDir":       {reflect.TypeOf(q.BothDir), constant.MakeInt64(int64(q.BothDir))},
			"Chan":          {reflect.TypeOf(q.Chan), constant.MakeInt64(int64(q.Chan))},
			"Complex128":    {reflect.TypeOf(q.Complex128), constant.MakeInt64(int64(q.Complex128))},
			"Complex64":     {reflect.TypeOf(q.Complex64), constant.MakeInt64(int64(q.Complex64))},
			"Float32":       {reflect.TypeOf(q.Float32), constant.MakeInt64(int64(q.Float32))},
			"Float64":       {reflect.TypeOf(q.Float64), constant.MakeInt64(int64(q.Float64))},
			"Func":          {reflect.TypeOf(q.Func), constant.MakeInt64(int64(q.Func))},
			"Int":           {reflect.TypeOf(q.Int), constant.MakeInt64(int64(q.Int))},
			"Int16":         {reflect.TypeOf(q.Int16), constant.MakeInt64(int64(q.Int16))},
			"Int32":         {reflect.TypeOf(q.Int32), constant.MakeInt64(int64(q.Int32))},
			"Int64":         {reflect.TypeOf(q.Int64), constant.MakeInt64(int64(q.Int64))},
			"Int8":          {reflect.TypeOf(q.Int8), constant.MakeInt64(int64(q.Int8))},
			"Interface":     {reflect.TypeOf(q.Interface), constant.MakeInt64(int64(q.Interface))},
			"Invalid":       {reflect.TypeOf(q.Invalid), constant.MakeInt64(int64(q.Invalid))},
			"Map":           {reflect.TypeOf(q.Map), constant.MakeInt64(int64(q.Map))},
			"Pointer":       {reflect.TypeOf(q.Pointer), constant.MakeInt64(int64(q.Pointer))},
			"Ptr":           {reflect.TypeOf(q.Ptr), constant.MakeInt64(int64(q.Ptr))},
			"RecvDir":       {reflect.TypeOf(q.RecvDir), constant.MakeInt64(int64(q.RecvDir))},
			"SelectDefault": {reflect.TypeOf(q.SelectDefault), constant.MakeInt64(int64(q.SelectDefault))},
			"SelectRecv":    {reflect.TypeOf(q.SelectRecv), constant.MakeInt64(int64(q.SelectRecv))},
			"SelectSend":    {reflect.TypeOf(q.SelectSend), constant.MakeInt64(int64(q.SelectSend))},
			"SendDir":       {reflect.TypeOf(q.SendDir), constant.MakeInt64(int64(q.SendDir))},
			"Slice":         {reflect.TypeOf(q.Slice), constant.MakeInt64(int64(q.Slice))},
			"String":        {reflect.TypeOf(q.String), constant.MakeInt64(int64(q.String))},
			"Struct":        {reflect.TypeOf(q.Struct), constant.MakeInt64(int64(q.Struct))},
			"Uint":          {reflect.TypeOf(q.Uint), constant.MakeInt64(int64(q.Uint))},
			"Uint16":        {reflect.TypeOf(q.Uint16), constant.MakeInt64(int64(q.Uint16))},
			"Uint32":        {reflect.TypeOf(q.Uint32), constant.MakeInt64(int64(q.Uint32))},
			"Uint64":        {reflect.TypeOf(q.Uint64), constant.MakeInt64(int64(q.Uint64))},
			"Uint8":         {reflect.TypeOf(q.Uint8), constant.MakeInt64(int64(q.Uint8))},
			"Uintptr":       {reflect.TypeOf(q.Uintptr), constant.MakeInt64(int64(q.Uintptr))},
			"UnsafePointer": {reflect.TypeOf(q.UnsafePointer), constant.MakeInt64(int64(q.UnsafePointer))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
