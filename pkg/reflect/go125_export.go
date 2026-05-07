// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package reflect

import (
	q "reflect"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("reflect", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "reflect",
			Path: "reflect",
			Deps: map[string]string{
				"errors":                "errors",
				"internal/abi":          "abi",
				"internal/bytealg":      "bytealg",
				"internal/goarch":       "goarch",
				"internal/itoa":         "itoa",
				"internal/race":         "race",
				"internal/runtime/maps": "maps",
				"internal/runtime/sys":  "sys",
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Append":          q.Append,
				"AppendSlice":     q.AppendSlice,
				"ArrayOf":         q.ArrayOf,
				"ChanOf":          q.ChanOf,
				"Copy":            q.Copy,
				"DeepEqual":       q.DeepEqual,
				"FuncOf":          q.FuncOf,
				"Indirect":        q.Indirect,
				"MakeChan":        q.MakeChan,
				"MakeFunc":        q.MakeFunc,
				"MakeMap":         q.MakeMap,
				"MakeMapWithSize": q.MakeMapWithSize,
				"MakeSlice":       q.MakeSlice,
				"MapOf":           q.MapOf,
				"New":             q.New,
				"NewAt":           q.NewAt,
				"PointerTo":       q.PointerTo,
				"PtrTo":           q.PtrTo,
				"Select":          q.Select,
				"SliceAt":         q.SliceAt,
				"SliceOf":         q.SliceOf,
				"StructOf":        q.StructOf,
				"Swapper":         q.Swapper,
				"TypeOf":          q.TypeOf,
				"ValueOf":         q.ValueOf,
				"VisibleFields":   q.VisibleFields,
				"Zero":            q.Zero,
			},
			TypedConsts: map[string]interface{}{
				"Array":         q.Array,
				"Bool":          q.Bool,
				"BothDir":       q.BothDir,
				"Chan":          q.Chan,
				"Complex128":    q.Complex128,
				"Complex64":     q.Complex64,
				"Float32":       q.Float32,
				"Float64":       q.Float64,
				"Func":          q.Func,
				"Int":           q.Int,
				"Int16":         q.Int16,
				"Int32":         q.Int32,
				"Int64":         q.Int64,
				"Int8":          q.Int8,
				"Interface":     q.Interface,
				"Invalid":       q.Invalid,
				"Map":           q.Map,
				"Pointer":       q.Pointer,
				"Ptr":           q.Ptr,
				"RecvDir":       q.RecvDir,
				"SelectDefault": q.SelectDefault,
				"SelectRecv":    q.SelectRecv,
				"SelectSend":    q.SelectSend,
				"SendDir":       q.SendDir,
				"Slice":         q.Slice,
				"String":        q.String,
				"Struct":        q.Struct,
				"Uint":          q.Uint,
				"Uint16":        q.Uint16,
				"Uint32":        q.Uint32,
				"Uint64":        q.Uint64,
				"Uint8":         q.Uint8,
				"Uintptr":       q.Uintptr,
				"UnsafePointer": q.UnsafePointer,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
