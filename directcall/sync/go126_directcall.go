// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package sync

import (
	q "sync"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("sync", map[string]ixgo.DirectCallAdapter{
		"(*Cond).Broadcast":       method_ptr_Cond_Broadcast,
		"(*Cond).Signal":          method_ptr_Cond_Signal,
		"(*Cond).Wait":            method_ptr_Cond_Wait,
		"(*Map).Clear":            method_ptr_Map_Clear,
		"(*Map).CompareAndDelete": method_ptr_Map_CompareAndDelete,
		"(*Map).CompareAndSwap":   method_ptr_Map_CompareAndSwap,
		"(*Map).Delete":           method_ptr_Map_Delete,
		"(*Map).Range":            method_ptr_Map_Range,
		"(*Map).Store":            method_ptr_Map_Store,
		"(*Mutex).Lock":           method_ptr_Mutex_Lock,
		"(*Mutex).TryLock":        method_ptr_Mutex_TryLock,
		"(*Mutex).Unlock":         method_ptr_Mutex_Unlock,
		"(*Once).Do":              method_ptr_Once_Do,
		"(*Pool).Get":             method_ptr_Pool_Get,
		"(*Pool).Put":             method_ptr_Pool_Put,
		"(*RWMutex).Lock":         method_ptr_RWMutex_Lock,
		"(*RWMutex).RLock":        method_ptr_RWMutex_RLock,
		"(*RWMutex).RLocker":      method_ptr_RWMutex_RLocker,
		"(*RWMutex).RUnlock":      method_ptr_RWMutex_RUnlock,
		"(*RWMutex).TryLock":      method_ptr_RWMutex_TryLock,
		"(*RWMutex).TryRLock":     method_ptr_RWMutex_TryRLock,
		"(*RWMutex).Unlock":       method_ptr_RWMutex_Unlock,
		"(*WaitGroup).Add":        method_ptr_WaitGroup_Add,
		"(*WaitGroup).Done":       method_ptr_WaitGroup_Done,
		"(*WaitGroup).Go":         method_ptr_WaitGroup_Go,
		"(*WaitGroup).Wait":       method_ptr_WaitGroup_Wait,
		"NewCond":                 func_NewCond,
		"OnceFunc":                func_OnceFunc,
	})
}

func method_ptr_Cond_Broadcast(ctx ixgo.DirectCallContext) {
	(*q.Cond).Broadcast(ixgo.DirectCallArg[*q.Cond](ctx, 0))
}

func method_ptr_Cond_Signal(ctx ixgo.DirectCallContext) {
	(*q.Cond).Signal(ixgo.DirectCallArg[*q.Cond](ctx, 0))
}

func method_ptr_Cond_Wait(ctx ixgo.DirectCallContext) {
	(*q.Cond).Wait(ixgo.DirectCallArg[*q.Cond](ctx, 0))
}

func method_ptr_Map_Clear(ctx ixgo.DirectCallContext) {
	(*q.Map).Clear(ixgo.DirectCallArg[*q.Map](ctx, 0))
}

func method_ptr_Map_CompareAndDelete(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Map).CompareAndDelete(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[any](ctx, 1), ixgo.DirectCallArg[any](ctx, 2)))
}

func method_ptr_Map_CompareAndSwap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Map).CompareAndSwap(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[any](ctx, 1), ixgo.DirectCallArg[any](ctx, 2), ixgo.DirectCallArg[any](ctx, 3)))
}

func method_ptr_Map_Delete(ctx ixgo.DirectCallContext) {
	(*q.Map).Delete(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[any](ctx, 1))
}

func method_ptr_Map_Range(ctx ixgo.DirectCallContext) {
	(*q.Map).Range(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[func(key any, value any) bool](ctx, 1))
}

func method_ptr_Map_Store(ctx ixgo.DirectCallContext) {
	(*q.Map).Store(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[any](ctx, 1), ixgo.DirectCallArg[any](ctx, 2))
}

func method_ptr_Mutex_Lock(ctx ixgo.DirectCallContext) {
	(*q.Mutex).Lock(ixgo.DirectCallArg[*q.Mutex](ctx, 0))
}

func method_ptr_Mutex_TryLock(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Mutex).TryLock(ixgo.DirectCallArg[*q.Mutex](ctx, 0)))
}

func method_ptr_Mutex_Unlock(ctx ixgo.DirectCallContext) {
	(*q.Mutex).Unlock(ixgo.DirectCallArg[*q.Mutex](ctx, 0))
}

func func_NewCond(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCond(ixgo.DirectCallArg[q.Locker](ctx, 0)))
}

func method_ptr_Once_Do(ctx ixgo.DirectCallContext) {
	(*q.Once).Do(ixgo.DirectCallArg[*q.Once](ctx, 0), ixgo.DirectCallArg[func()](ctx, 1))
}

func func_OnceFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OnceFunc(ixgo.DirectCallArg[func()](ctx, 0)))
}

func method_ptr_Pool_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pool).Get(ixgo.DirectCallArg[*q.Pool](ctx, 0)))
}

func method_ptr_Pool_Put(ctx ixgo.DirectCallContext) {
	(*q.Pool).Put(ixgo.DirectCallArg[*q.Pool](ctx, 0), ixgo.DirectCallArg[any](ctx, 1))
}

func method_ptr_RWMutex_Lock(ctx ixgo.DirectCallContext) {
	(*q.RWMutex).Lock(ixgo.DirectCallArg[*q.RWMutex](ctx, 0))
}

func method_ptr_RWMutex_RLock(ctx ixgo.DirectCallContext) {
	(*q.RWMutex).RLock(ixgo.DirectCallArg[*q.RWMutex](ctx, 0))
}

func method_ptr_RWMutex_RLocker(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RWMutex).RLocker(ixgo.DirectCallArg[*q.RWMutex](ctx, 0)))
}

func method_ptr_RWMutex_RUnlock(ctx ixgo.DirectCallContext) {
	(*q.RWMutex).RUnlock(ixgo.DirectCallArg[*q.RWMutex](ctx, 0))
}

func method_ptr_RWMutex_TryLock(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RWMutex).TryLock(ixgo.DirectCallArg[*q.RWMutex](ctx, 0)))
}

func method_ptr_RWMutex_TryRLock(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RWMutex).TryRLock(ixgo.DirectCallArg[*q.RWMutex](ctx, 0)))
}

func method_ptr_RWMutex_Unlock(ctx ixgo.DirectCallContext) {
	(*q.RWMutex).Unlock(ixgo.DirectCallArg[*q.RWMutex](ctx, 0))
}

func method_ptr_WaitGroup_Add(ctx ixgo.DirectCallContext) {
	(*q.WaitGroup).Add(ixgo.DirectCallArg[*q.WaitGroup](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_WaitGroup_Done(ctx ixgo.DirectCallContext) {
	(*q.WaitGroup).Done(ixgo.DirectCallArg[*q.WaitGroup](ctx, 0))
}

func method_ptr_WaitGroup_Go(ctx ixgo.DirectCallContext) {
	(*q.WaitGroup).Go(ixgo.DirectCallArg[*q.WaitGroup](ctx, 0), ixgo.DirectCallArg[func()](ctx, 1))
}

func method_ptr_WaitGroup_Wait(ctx ixgo.DirectCallContext) {
	(*q.WaitGroup).Wait(ixgo.DirectCallArg[*q.WaitGroup](ctx, 0))
}
