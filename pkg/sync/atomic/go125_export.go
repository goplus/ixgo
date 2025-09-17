// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package atomic

import (
	q "sync/atomic"

	"reflect"
	_ "unsafe"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"AddInt32":              reflect.ValueOf(q.AddInt32),
			"AddInt64":              reflect.ValueOf(q.AddInt64),
			"AddUint32":             reflect.ValueOf(q.AddUint32),
			"AddUint64":             reflect.ValueOf(q.AddUint64),
			"AddUintptr":            reflect.ValueOf(q.AddUintptr),
			"AndInt32":              reflect.ValueOf(q.AndInt32),
			"AndInt64":              reflect.ValueOf(q.AndInt64),
			"AndUint32":             reflect.ValueOf(q.AndUint32),
			"AndUint64":             reflect.ValueOf(q.AndUint64),
			"AndUintptr":            reflect.ValueOf(q.AndUintptr),
			"CompareAndSwapInt32":   reflect.ValueOf(q.CompareAndSwapInt32),
			"CompareAndSwapInt64":   reflect.ValueOf(q.CompareAndSwapInt64),
			"CompareAndSwapPointer": reflect.ValueOf(q.CompareAndSwapPointer),
			"CompareAndSwapUint32":  reflect.ValueOf(q.CompareAndSwapUint32),
			"CompareAndSwapUint64":  reflect.ValueOf(q.CompareAndSwapUint64),
			"CompareAndSwapUintptr": reflect.ValueOf(q.CompareAndSwapUintptr),
			"LoadInt32":             reflect.ValueOf(q.LoadInt32),
			"LoadInt64":             reflect.ValueOf(q.LoadInt64),
			"LoadPointer":           reflect.ValueOf(q.LoadPointer),
			"LoadUint32":            reflect.ValueOf(q.LoadUint32),
			"LoadUint64":            reflect.ValueOf(q.LoadUint64),
			"LoadUintptr":           reflect.ValueOf(q.LoadUintptr),
			"OrInt32":               reflect.ValueOf(q.OrInt32),
			"OrInt64":               reflect.ValueOf(q.OrInt64),
			"OrUint32":              reflect.ValueOf(q.OrUint32),
			"OrUint64":              reflect.ValueOf(q.OrUint64),
			"OrUintptr":             reflect.ValueOf(q.OrUintptr),
			"StoreInt32":            reflect.ValueOf(q.StoreInt32),
			"StoreInt64":            reflect.ValueOf(q.StoreInt64),
			"StorePointer":          reflect.ValueOf(q.StorePointer),
			"StoreUint32":           reflect.ValueOf(q.StoreUint32),
			"StoreUint64":           reflect.ValueOf(q.StoreUint64),
			"StoreUintptr":          reflect.ValueOf(q.StoreUintptr),
			"SwapInt32":             reflect.ValueOf(q.SwapInt32),
			"SwapInt64":             reflect.ValueOf(q.SwapInt64),
			"SwapPointer":           reflect.ValueOf(q.SwapPointer),
			"SwapUint32":            reflect.ValueOf(q.SwapUint32),
			"SwapUint64":            reflect.ValueOf(q.SwapUint64),
			"SwapUintptr":           reflect.ValueOf(q.SwapUintptr),
			"runtime_procPin":       reflect.ValueOf(_runtime_procPin),
			"runtime_procUnpin":     reflect.ValueOf(_runtime_procUnpin),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
		Source:        source,
	})
}

//go:linkname _runtime_procPin sync/atomic.runtime_procPin
func _runtime_procPin() int

//go:linkname _runtime_procUnpin sync/atomic.runtime_procUnpin
func _runtime_procUnpin()

var source = "package atomic\n\nimport \"unsafe\"\n\nfunc SwapInt32(addr *int32, new int32) (old int32)\n\nfunc SwapUint32(addr *uint32, new uint32) (old uint32)\n\nfunc SwapUintptr(addr *uintptr, new uintptr) (old uintptr)\n\nfunc SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)\n\nfunc CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)\n\nfunc CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)\n\nfunc CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)\n\nfunc CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)\n\nfunc AddInt32(addr *int32, delta int32) (new int32)\n\nfunc AddUint32(addr *uint32, delta uint32) (new uint32)\n\nfunc AddUintptr(addr *uintptr, delta uintptr) (new uintptr)\n\nfunc AndInt32(addr *int32, mask int32) (old int32)\n\nfunc AndUint32(addr *uint32, mask uint32) (old uint32)\n\nfunc AndUintptr(addr *uintptr, mask uintptr) (old uintptr)\n\nfunc OrInt32(addr *int32, mask int32) (old int32)\n\nfunc OrUint32(addr *uint32, mask uint32) (old uint32)\n\nfunc OrUintptr(addr *uintptr, mask uintptr) (old uintptr)\n\nfunc LoadInt32(addr *int32) (val int32)\n\nfunc LoadUint32(addr *uint32) (val uint32)\n\nfunc LoadUintptr(addr *uintptr) (val uintptr)\n\nfunc LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)\n\nfunc StoreInt32(addr *int32, val int32)\n\nfunc StoreUint32(addr *uint32, val uint32)\n\nfunc StoreUintptr(addr *uintptr, val uintptr)\n\nfunc StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)\nfunc SwapInt64(addr *int64, new int64) (old int64)\n\nfunc SwapUint64(addr *uint64, new uint64) (old uint64)\n\nfunc CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)\n\nfunc CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)\n\nfunc AddInt64(addr *int64, delta int64) (new int64)\n\nfunc AddUint64(addr *uint64, delta uint64) (new uint64)\n\nfunc AndInt64(addr *int64, mask int64) (old int64)\n\nfunc AndUint64(addr *uint64, mask uint64) (old uint64)\n\nfunc OrInt64(addr *int64, mask int64) (old int64)\n\nfunc OrUint64(addr *uint64, mask uint64) (old uint64)\n\nfunc LoadInt64(addr *int64) (val int64)\n\nfunc LoadUint64(addr *uint64) (val uint64)\n\nfunc StoreInt64(addr *int64, val int64)\n\nfunc StoreUint64(addr *uint64, val uint64)\n\ntype Bool struct {\n\t_\tnoCopy\n\tv\tuint32\n}\n\nfunc (x *Bool) Load() bool\n\nfunc (x *Bool) Store(val bool)\n\nfunc (x *Bool) Swap(new bool) (old bool)\n\nfunc (x *Bool) CompareAndSwap(old, new bool) (swapped bool)\n\nfunc b32(b bool) uint32 {\n\tif b {\n\t\treturn 1\n\t}\n\treturn 0\n}\n\ntype Pointer[T any] struct {\n\t_\t[0]*T\n\n\t_\tnoCopy\n\tv\tunsafe.Pointer\n}\n\nfunc (x *Pointer[T]) Load() *T\t{ return (*T)(LoadPointer(&x.v)) }\n\nfunc (x *Pointer[T]) Store(val *T)\t{ StorePointer(&x.v, unsafe.Pointer(val)) }\n\nfunc (x *Pointer[T]) Swap(new *T) (old *T)\t{ return (*T)(SwapPointer(&x.v, unsafe.Pointer(new))) }\n\nfunc (x *Pointer[T]) CompareAndSwap(old, new *T) (swapped bool) {\n\treturn CompareAndSwapPointer(&x.v, unsafe.Pointer(old), unsafe.Pointer(new))\n}\n\ntype Int32 struct {\n\t_\tnoCopy\n\tv\tint32\n}\n\nfunc (x *Int32) Load() int32\n\nfunc (x *Int32) Store(val int32)\n\nfunc (x *Int32) Swap(new int32) (old int32)\n\nfunc (x *Int32) CompareAndSwap(old, new int32) (swapped bool)\n\nfunc (x *Int32) Add(delta int32) (new int32)\n\nfunc (x *Int32) And(mask int32) (old int32)\n\nfunc (x *Int32) Or(mask int32) (old int32)\n\ntype Int64 struct {\n\t_\tnoCopy\n\t_\talign64\n\tv\tint64\n}\n\nfunc (x *Int64) Load() int64\n\nfunc (x *Int64) Store(val int64)\n\nfunc (x *Int64) Swap(new int64) (old int64)\n\nfunc (x *Int64) CompareAndSwap(old, new int64) (swapped bool)\n\nfunc (x *Int64) Add(delta int64) (new int64)\n\nfunc (x *Int64) And(mask int64) (old int64)\n\nfunc (x *Int64) Or(mask int64) (old int64)\n\ntype Uint32 struct {\n\t_\tnoCopy\n\tv\tuint32\n}\n\nfunc (x *Uint32) Load() uint32\n\nfunc (x *Uint32) Store(val uint32)\n\nfunc (x *Uint32) Swap(new uint32) (old uint32)\n\nfunc (x *Uint32) CompareAndSwap(old, new uint32) (swapped bool)\n\nfunc (x *Uint32) Add(delta uint32) (new uint32)\n\nfunc (x *Uint32) And(mask uint32) (old uint32)\n\nfunc (x *Uint32) Or(mask uint32) (old uint32)\n\ntype Uint64 struct {\n\t_\tnoCopy\n\t_\talign64\n\tv\tuint64\n}\n\nfunc (x *Uint64) Load() uint64\n\nfunc (x *Uint64) Store(val uint64)\n\nfunc (x *Uint64) Swap(new uint64) (old uint64)\n\nfunc (x *Uint64) CompareAndSwap(old, new uint64) (swapped bool)\n\nfunc (x *Uint64) Add(delta uint64) (new uint64)\n\nfunc (x *Uint64) And(mask uint64) (old uint64)\n\nfunc (x *Uint64) Or(mask uint64) (old uint64)\n\ntype Uintptr struct {\n\t_\tnoCopy\n\tv\tuintptr\n}\n\nfunc (x *Uintptr) Load() uintptr\n\nfunc (x *Uintptr) Store(val uintptr)\n\nfunc (x *Uintptr) Swap(new uintptr) (old uintptr)\n\nfunc (x *Uintptr) CompareAndSwap(old, new uintptr) (swapped bool)\n\nfunc (x *Uintptr) Add(delta uintptr) (new uintptr)\n\nfunc (x *Uintptr) And(mask uintptr) (old uintptr)\n\nfunc (x *Uintptr) Or(mask uintptr) (old uintptr)\n\ntype noCopy struct{}\n\nfunc (*noCopy) Lock()\t{}\nfunc (*noCopy) Unlock()\t{}\n\ntype align64 struct{}\ntype Value struct {\n\tv any\n}\n\ntype efaceWords struct {\n\ttyp\tunsafe.Pointer\n\tdata\tunsafe.Pointer\n}\n\nfunc (v *Value) Load() (val any)\n\nvar firstStoreInProgress byte\n\nfunc (v *Value) Store(val any)\n\nfunc (v *Value) Swap(new any) (old any)\n\nfunc (v *Value) CompareAndSwap(old, new any) (swapped bool)\n\nfunc runtime_procPin() int\nfunc runtime_procUnpin()\n"
