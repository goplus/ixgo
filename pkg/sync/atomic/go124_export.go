// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package atomic

import (
	q "sync/atomic"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("sync/atomic", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "atomic",
			Path: "sync/atomic",
			Deps: map[string]string{
				"unsafe": "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Bool":    reflect.TypeOf((*q.Bool)(nil)).Elem(),
				"Int32":   reflect.TypeOf((*q.Int32)(nil)).Elem(),
				"Int64":   reflect.TypeOf((*q.Int64)(nil)).Elem(),
				"Uint32":  reflect.TypeOf((*q.Uint32)(nil)).Elem(),
				"Uint64":  reflect.TypeOf((*q.Uint64)(nil)).Elem(),
				"Uintptr": reflect.TypeOf((*q.Uintptr)(nil)).Elem(),
				"Value":   reflect.TypeOf((*q.Value)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"AddInt32":              q.AddInt32,
				"AddInt64":              q.AddInt64,
				"AddUint32":             q.AddUint32,
				"AddUint64":             q.AddUint64,
				"AddUintptr":            q.AddUintptr,
				"AndInt32":              q.AndInt32,
				"AndInt64":              q.AndInt64,
				"AndUint32":             q.AndUint32,
				"AndUint64":             q.AndUint64,
				"AndUintptr":            q.AndUintptr,
				"CompareAndSwapInt32":   q.CompareAndSwapInt32,
				"CompareAndSwapInt64":   q.CompareAndSwapInt64,
				"CompareAndSwapPointer": q.CompareAndSwapPointer,
				"CompareAndSwapUint32":  q.CompareAndSwapUint32,
				"CompareAndSwapUint64":  q.CompareAndSwapUint64,
				"CompareAndSwapUintptr": q.CompareAndSwapUintptr,
				"LoadInt32":             q.LoadInt32,
				"LoadInt64":             q.LoadInt64,
				"LoadPointer":           q.LoadPointer,
				"LoadUint32":            q.LoadUint32,
				"LoadUint64":            q.LoadUint64,
				"LoadUintptr":           q.LoadUintptr,
				"OrInt32":               q.OrInt32,
				"OrInt64":               q.OrInt64,
				"OrUint32":              q.OrUint32,
				"OrUint64":              q.OrUint64,
				"OrUintptr":             q.OrUintptr,
				"StoreInt32":            q.StoreInt32,
				"StoreInt64":            q.StoreInt64,
				"StorePointer":          q.StorePointer,
				"StoreUint32":           q.StoreUint32,
				"StoreUint64":           q.StoreUint64,
				"StoreUintptr":          q.StoreUintptr,
				"SwapInt32":             q.SwapInt32,
				"SwapInt64":             q.SwapInt64,
				"SwapPointer":           q.SwapPointer,
				"SwapUint32":            q.SwapUint32,
				"SwapUint64":            q.SwapUint64,
				"SwapUintptr":           q.SwapUintptr,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
