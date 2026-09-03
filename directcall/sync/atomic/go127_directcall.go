// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package atomic

import (
	q "sync/atomic"

	"github.com/goplus/ixgo"
	unsafe "unsafe"
)

func init() {
	ixgo.RegisterDirectCalls("sync/atomic", map[string]ixgo.DirectCallAdapter{
		"(*Bool).CompareAndSwap":    method_ptr_Bool_CompareAndSwap,
		"(*Bool).Load":              method_ptr_Bool_Load,
		"(*Bool).Store":             method_ptr_Bool_Store,
		"(*Bool).Swap":              method_ptr_Bool_Swap,
		"(*Int32).Add":              method_ptr_Int32_Add,
		"(*Int32).And":              method_ptr_Int32_And,
		"(*Int32).CompareAndSwap":   method_ptr_Int32_CompareAndSwap,
		"(*Int32).Load":             method_ptr_Int32_Load,
		"(*Int32).Or":               method_ptr_Int32_Or,
		"(*Int32).Store":            method_ptr_Int32_Store,
		"(*Int32).Swap":             method_ptr_Int32_Swap,
		"(*Int64).Add":              method_ptr_Int64_Add,
		"(*Int64).And":              method_ptr_Int64_And,
		"(*Int64).CompareAndSwap":   method_ptr_Int64_CompareAndSwap,
		"(*Int64).Load":             method_ptr_Int64_Load,
		"(*Int64).Or":               method_ptr_Int64_Or,
		"(*Int64).Store":            method_ptr_Int64_Store,
		"(*Int64).Swap":             method_ptr_Int64_Swap,
		"(*Uint32).Add":             method_ptr_Uint32_Add,
		"(*Uint32).And":             method_ptr_Uint32_And,
		"(*Uint32).CompareAndSwap":  method_ptr_Uint32_CompareAndSwap,
		"(*Uint32).Load":            method_ptr_Uint32_Load,
		"(*Uint32).Or":              method_ptr_Uint32_Or,
		"(*Uint32).Store":           method_ptr_Uint32_Store,
		"(*Uint32).Swap":            method_ptr_Uint32_Swap,
		"(*Uint64).Add":             method_ptr_Uint64_Add,
		"(*Uint64).And":             method_ptr_Uint64_And,
		"(*Uint64).CompareAndSwap":  method_ptr_Uint64_CompareAndSwap,
		"(*Uint64).Load":            method_ptr_Uint64_Load,
		"(*Uint64).Or":              method_ptr_Uint64_Or,
		"(*Uint64).Store":           method_ptr_Uint64_Store,
		"(*Uint64).Swap":            method_ptr_Uint64_Swap,
		"(*Uintptr).Add":            method_ptr_Uintptr_Add,
		"(*Uintptr).And":            method_ptr_Uintptr_And,
		"(*Uintptr).CompareAndSwap": method_ptr_Uintptr_CompareAndSwap,
		"(*Uintptr).Load":           method_ptr_Uintptr_Load,
		"(*Uintptr).Or":             method_ptr_Uintptr_Or,
		"(*Uintptr).Store":          method_ptr_Uintptr_Store,
		"(*Uintptr).Swap":           method_ptr_Uintptr_Swap,
		"(*Value).CompareAndSwap":   method_ptr_Value_CompareAndSwap,
		"(*Value).Load":             method_ptr_Value_Load,
		"(*Value).Store":            method_ptr_Value_Store,
		"(*Value).Swap":             method_ptr_Value_Swap,
		"AddInt32":                  func_AddInt32,
		"AddInt64":                  func_AddInt64,
		"AddUint32":                 func_AddUint32,
		"AddUint64":                 func_AddUint64,
		"AddUintptr":                func_AddUintptr,
		"AndInt32":                  func_AndInt32,
		"AndInt64":                  func_AndInt64,
		"AndUint32":                 func_AndUint32,
		"AndUint64":                 func_AndUint64,
		"AndUintptr":                func_AndUintptr,
		"CompareAndSwapInt32":       func_CompareAndSwapInt32,
		"CompareAndSwapInt64":       func_CompareAndSwapInt64,
		"CompareAndSwapPointer":     func_CompareAndSwapPointer,
		"CompareAndSwapUint32":      func_CompareAndSwapUint32,
		"CompareAndSwapUint64":      func_CompareAndSwapUint64,
		"CompareAndSwapUintptr":     func_CompareAndSwapUintptr,
		"LoadInt32":                 func_LoadInt32,
		"LoadInt64":                 func_LoadInt64,
		"LoadPointer":               func_LoadPointer,
		"LoadUint32":                func_LoadUint32,
		"LoadUint64":                func_LoadUint64,
		"LoadUintptr":               func_LoadUintptr,
		"OrInt32":                   func_OrInt32,
		"OrInt64":                   func_OrInt64,
		"OrUint32":                  func_OrUint32,
		"OrUint64":                  func_OrUint64,
		"OrUintptr":                 func_OrUintptr,
		"StoreInt32":                func_StoreInt32,
		"StoreInt64":                func_StoreInt64,
		"StorePointer":              func_StorePointer,
		"StoreUint32":               func_StoreUint32,
		"StoreUint64":               func_StoreUint64,
		"StoreUintptr":              func_StoreUintptr,
		"SwapInt32":                 func_SwapInt32,
		"SwapInt64":                 func_SwapInt64,
		"SwapPointer":               func_SwapPointer,
		"SwapUint32":                func_SwapUint32,
		"SwapUint64":                func_SwapUint64,
		"SwapUintptr":               func_SwapUintptr,
	})
}

func func_AddInt32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddInt32(ixgo.DirectCallArg[*int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func func_AddInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddInt64(ixgo.DirectCallArg[*int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_AddUint32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddUint32(ixgo.DirectCallArg[*uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func func_AddUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddUint64(ixgo.DirectCallArg[*uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func func_AddUintptr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AddUintptr(ixgo.DirectCallArg[*uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1)))
}

func func_AndInt32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AndInt32(ixgo.DirectCallArg[*int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func func_AndInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AndInt64(ixgo.DirectCallArg[*int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_AndUint32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AndUint32(ixgo.DirectCallArg[*uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func func_AndUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AndUint64(ixgo.DirectCallArg[*uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func func_AndUintptr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AndUintptr(ixgo.DirectCallArg[*uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1)))
}

func method_ptr_Bool_CompareAndSwap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Bool).CompareAndSwap(ixgo.DirectCallArg[*q.Bool](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1), ixgo.DirectCallArg[bool](ctx, 2)))
}

func method_ptr_Bool_Load(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Bool).Load(ixgo.DirectCallArg[*q.Bool](ctx, 0)))
}

func method_ptr_Bool_Store(ctx ixgo.DirectCallContext) {
	(*q.Bool).Store(ixgo.DirectCallArg[*q.Bool](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_ptr_Bool_Swap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Bool).Swap(ixgo.DirectCallArg[*q.Bool](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1)))
}

func func_CompareAndSwapInt32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CompareAndSwapInt32(ixgo.DirectCallArg[*int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1), ixgo.DirectCallArg[int32](ctx, 2)))
}

func func_CompareAndSwapInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CompareAndSwapInt64(ixgo.DirectCallArg[*int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1), ixgo.DirectCallArg[int64](ctx, 2)))
}

func func_CompareAndSwapPointer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CompareAndSwapPointer(ixgo.DirectCallArg[*unsafe.Pointer](ctx, 0), ixgo.DirectCallArg[unsafe.Pointer](ctx, 1), ixgo.DirectCallArg[unsafe.Pointer](ctx, 2)))
}

func func_CompareAndSwapUint32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CompareAndSwapUint32(ixgo.DirectCallArg[*uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1), ixgo.DirectCallArg[uint32](ctx, 2)))
}

func func_CompareAndSwapUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CompareAndSwapUint64(ixgo.DirectCallArg[*uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1), ixgo.DirectCallArg[uint64](ctx, 2)))
}

func func_CompareAndSwapUintptr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CompareAndSwapUintptr(ixgo.DirectCallArg[*uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1), ixgo.DirectCallArg[uintptr](ctx, 2)))
}

func method_ptr_Int32_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int32).Add(ixgo.DirectCallArg[*q.Int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func method_ptr_Int32_And(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int32).And(ixgo.DirectCallArg[*q.Int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func method_ptr_Int32_CompareAndSwap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int32).CompareAndSwap(ixgo.DirectCallArg[*q.Int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1), ixgo.DirectCallArg[int32](ctx, 2)))
}

func method_ptr_Int32_Load(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int32).Load(ixgo.DirectCallArg[*q.Int32](ctx, 0)))
}

func method_ptr_Int32_Or(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int32).Or(ixgo.DirectCallArg[*q.Int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func method_ptr_Int32_Store(ctx ixgo.DirectCallContext) {
	(*q.Int32).Store(ixgo.DirectCallArg[*q.Int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1))
}

func method_ptr_Int32_Swap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int32).Swap(ixgo.DirectCallArg[*q.Int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func method_ptr_Int64_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int64).Add(ixgo.DirectCallArg[*q.Int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Int64_And(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int64).And(ixgo.DirectCallArg[*q.Int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Int64_CompareAndSwap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int64).CompareAndSwap(ixgo.DirectCallArg[*q.Int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1), ixgo.DirectCallArg[int64](ctx, 2)))
}

func method_ptr_Int64_Load(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int64).Load(ixgo.DirectCallArg[*q.Int64](ctx, 0)))
}

func method_ptr_Int64_Or(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int64).Or(ixgo.DirectCallArg[*q.Int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Int64_Store(ctx ixgo.DirectCallContext) {
	(*q.Int64).Store(ixgo.DirectCallArg[*q.Int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1))
}

func method_ptr_Int64_Swap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int64).Swap(ixgo.DirectCallArg[*q.Int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_LoadInt32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LoadInt32(ixgo.DirectCallArg[*int32](ctx, 0)))
}

func func_LoadInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LoadInt64(ixgo.DirectCallArg[*int64](ctx, 0)))
}

func func_LoadPointer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LoadPointer(ixgo.DirectCallArg[*unsafe.Pointer](ctx, 0)))
}

func func_LoadUint32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LoadUint32(ixgo.DirectCallArg[*uint32](ctx, 0)))
}

func func_LoadUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LoadUint64(ixgo.DirectCallArg[*uint64](ctx, 0)))
}

func func_LoadUintptr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LoadUintptr(ixgo.DirectCallArg[*uintptr](ctx, 0)))
}

func func_OrInt32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OrInt32(ixgo.DirectCallArg[*int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func func_OrInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OrInt64(ixgo.DirectCallArg[*int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_OrUint32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OrUint32(ixgo.DirectCallArg[*uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func func_OrUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OrUint64(ixgo.DirectCallArg[*uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func func_OrUintptr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OrUintptr(ixgo.DirectCallArg[*uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1)))
}

func func_StoreInt32(ctx ixgo.DirectCallContext) {
	q.StoreInt32(ixgo.DirectCallArg[*int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1))
}

func func_StoreInt64(ctx ixgo.DirectCallContext) {
	q.StoreInt64(ixgo.DirectCallArg[*int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1))
}

func func_StorePointer(ctx ixgo.DirectCallContext) {
	q.StorePointer(ixgo.DirectCallArg[*unsafe.Pointer](ctx, 0), ixgo.DirectCallArg[unsafe.Pointer](ctx, 1))
}

func func_StoreUint32(ctx ixgo.DirectCallContext) {
	q.StoreUint32(ixgo.DirectCallArg[*uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1))
}

func func_StoreUint64(ctx ixgo.DirectCallContext) {
	q.StoreUint64(ixgo.DirectCallArg[*uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1))
}

func func_StoreUintptr(ctx ixgo.DirectCallContext) {
	q.StoreUintptr(ixgo.DirectCallArg[*uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1))
}

func func_SwapInt32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SwapInt32(ixgo.DirectCallArg[*int32](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func func_SwapInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SwapInt64(ixgo.DirectCallArg[*int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_SwapPointer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SwapPointer(ixgo.DirectCallArg[*unsafe.Pointer](ctx, 0), ixgo.DirectCallArg[unsafe.Pointer](ctx, 1)))
}

func func_SwapUint32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SwapUint32(ixgo.DirectCallArg[*uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func func_SwapUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SwapUint64(ixgo.DirectCallArg[*uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func func_SwapUintptr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SwapUintptr(ixgo.DirectCallArg[*uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1)))
}

func method_ptr_Uint32_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint32).Add(ixgo.DirectCallArg[*q.Uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func method_ptr_Uint32_And(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint32).And(ixgo.DirectCallArg[*q.Uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func method_ptr_Uint32_CompareAndSwap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint32).CompareAndSwap(ixgo.DirectCallArg[*q.Uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1), ixgo.DirectCallArg[uint32](ctx, 2)))
}

func method_ptr_Uint32_Load(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint32).Load(ixgo.DirectCallArg[*q.Uint32](ctx, 0)))
}

func method_ptr_Uint32_Or(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint32).Or(ixgo.DirectCallArg[*q.Uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func method_ptr_Uint32_Store(ctx ixgo.DirectCallContext) {
	(*q.Uint32).Store(ixgo.DirectCallArg[*q.Uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1))
}

func method_ptr_Uint32_Swap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint32).Swap(ixgo.DirectCallArg[*q.Uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func method_ptr_Uint64_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint64).Add(ixgo.DirectCallArg[*q.Uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Uint64_And(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint64).And(ixgo.DirectCallArg[*q.Uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Uint64_CompareAndSwap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint64).CompareAndSwap(ixgo.DirectCallArg[*q.Uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1), ixgo.DirectCallArg[uint64](ctx, 2)))
}

func method_ptr_Uint64_Load(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint64).Load(ixgo.DirectCallArg[*q.Uint64](ctx, 0)))
}

func method_ptr_Uint64_Or(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint64).Or(ixgo.DirectCallArg[*q.Uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Uint64_Store(ctx ixgo.DirectCallContext) {
	(*q.Uint64).Store(ixgo.DirectCallArg[*q.Uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1))
}

func method_ptr_Uint64_Swap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uint64).Swap(ixgo.DirectCallArg[*q.Uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Uintptr_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uintptr).Add(ixgo.DirectCallArg[*q.Uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1)))
}

func method_ptr_Uintptr_And(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uintptr).And(ixgo.DirectCallArg[*q.Uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1)))
}

func method_ptr_Uintptr_CompareAndSwap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uintptr).CompareAndSwap(ixgo.DirectCallArg[*q.Uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1), ixgo.DirectCallArg[uintptr](ctx, 2)))
}

func method_ptr_Uintptr_Load(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uintptr).Load(ixgo.DirectCallArg[*q.Uintptr](ctx, 0)))
}

func method_ptr_Uintptr_Or(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uintptr).Or(ixgo.DirectCallArg[*q.Uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1)))
}

func method_ptr_Uintptr_Store(ctx ixgo.DirectCallContext) {
	(*q.Uintptr).Store(ixgo.DirectCallArg[*q.Uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1))
}

func method_ptr_Uintptr_Swap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uintptr).Swap(ixgo.DirectCallArg[*q.Uintptr](ctx, 0), ixgo.DirectCallArg[uintptr](ctx, 1)))
}

func method_ptr_Value_CompareAndSwap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).CompareAndSwap(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1), ixgo.DirectCallArg[interface{}](ctx, 2)))
}

func method_ptr_Value_Load(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Load(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_ptr_Value_Store(ctx ixgo.DirectCallContext) {
	(*q.Value).Store(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1))
}

func method_ptr_Value_Swap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Swap(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1)))
}
