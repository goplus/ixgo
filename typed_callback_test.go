/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ixgo

import (
	"reflect"
	"sync"
	"testing"

	"github.com/visualfc/funcval"
)

func TestTypedCallbacks(t *testing.T) {
	for _, mode := range []struct {
		name string
		mode Mode
	}{
		{name: "uncached"},
		{name: "cached", mode: EnableCachedReg},
		{name: "dynamic_fallback", mode: EnableCachedReg | DisableDynamicFuncCallAnalysis},
	} {
		t.Run(mode.name, func(t *testing.T) {
			interp := loadCallbacks(t, mode.mode, 0)
			left := newCallbacks(t, interp, 0)
			right := newCallbacks(t, interp, 1)
			increment, even, other := left.increment, left.even, right.even
			if !even() || other() {
				t.Fatal("independent closures lost their initial environment")
			}
			// Retained closures survive frame reuse.
			for i := 0; i < 4; i++ {
				increment()
				if even() || other() {
					t.Fatal("native callback changed the wrong closure environment")
				}
				runFunc(t, interp, "InvokeVoid", increment)
				if !runFunc(t, interp, "InvokeBool", even).(bool) {
					t.Fatal("interpreted dynamic callback lost its environment")
				}
			}
			nativeCalls := 0
			runFunc(t, interp, "InvokeVoid", func() { nativeCalls++ })
			if nativeCalls != 1 || !runFunc(t, interp, "InvokeBool", func() bool { return true }).(bool) {
				t.Fatal("native function fallback failed")
			}
			otherInterp := loadCallbacks(t, mode.mode, 0)
			runFunc(t, otherInterp, "InvokeVoid", increment)
			if runFunc(t, otherInterp, "InvokeBool", even).(bool) || other() {
				t.Fatal("callback invoked from another interpreter lost its environment")
			}
			invoke := func(f func()) { f() }
			reentrant := runFunc(t, interp, "Reentrant", invoke, 0).(func() bool)
			if reentrant() || !reentrant() {
				t.Fatal("nested host callbacks lost their environment")
			}
			if !runFunc(t, interp, "Signatures", invoke, func(f func() bool) bool { return f() }).(bool) {
				t.Fatal("named or unsupported callback signature failed")
			}
			fallbacks := runFunc(t, interp, "Fallbacks").([]interface{})
			if !reflect.ValueOf(fallbacks[0]).Call(nil)[0].Bool() {
				t.Fatal("named bool result changed")
			}
			if got := fallbacks[1].(func(int) int)(41); got != 42 {
				t.Fatalf("callback with argument = %d; want 42", got)
			}
			if got := fallbacks[2].(func(...int) int)(20, 22); got != 42 {
				t.Fatalf("variadic callback = %d; want 42", got)
			}
			if value, ok := fallbacks[3].(func() (int, bool))(); value != 42 || !ok {
				t.Fatalf("multiple callback results = (%d, %v); want (42, true)", value, ok)
			}
			if !IsLLGo {
				names := runFunc(t, interp, "PointerNames", increment, even).([]string)
				for i, name := range names {
					if name != []string{"main.MakeCallbacks.func1", "main.MakeCallbacks.func2"}[i] {
						t.Errorf("callback %d runtime name = %q; want interpreted closure", i, name)
					}
				}
			}
		})
	}
}

func TestTypedCallbackFunctionValues(t *testing.T) {
	if !funcval.IsSupport {
		t.Skip("requires gc function values")
	}
	i := loadCallbacks(t, 0, 0)
	values := runFunc(t, i, "MakeCallbacks", 1).([]interface{})
	for index, fn := range values {
		_, reflectBridges := funcval.Get(fn)
		want := 0
		if index == 2 {
			want = 1 // Reflection fallback.
		}
		if reflectBridges != want {
			t.Errorf("callback %T uses %d reflection bridges; want %d", fn, reflectBridges, want)
		}
		c := i.getMakeFuncVal(fn)
		if c == nil || c.interp != i || c.pfn == nil || len(c.env) != 1 {
			t.Fatalf("callback %T lost its interpreter, function or closure environment", fn)
		}
	}
	external := reflect.MakeFunc(reflect.TypeFor[func()](), func([]reflect.Value) []reflect.Value {
		return nil
	}).Interface()
	for _, fn := range []interface{}{
		nil, (func())(nil), (func() bool)(nil), func() {}, func() bool { return true }, external,
		42, "not a function", new(int), []int{1}, map[string]int{"x": 1}, make(chan int),
	} {
		if i.getMakeFuncVal(fn) != nil {
			t.Errorf("value %T identified as an interpreted function", fn)
		}
	}
}

func TestTypedCallbackExternalMakeFunc(t *testing.T) {
	interp := loadCallbacks(t, EnableCachedReg, 0)
	calls := 0
	void := reflect.MakeFunc(reflect.TypeFor[func()](), func([]reflect.Value) []reflect.Value {
		calls++
		return nil
	}).Interface()
	boolean := reflect.MakeFunc(reflect.TypeFor[func() bool](), func([]reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf(calls%2 == 0)}
	}).Interface()
	for i := 1; i <= 4; i++ {
		runFunc(t, interp, "InvokeVoid", void)
		if got := runFunc(t, interp, "InvokeBool", boolean).(bool); got != (i%2 == 0) {
			t.Fatalf("reflected callback result = %v after %d calls", got, i)
		}
	}
	if calls != 4 {
		t.Fatalf("reflected callback called %d times; want 4", calls)
	}
}

func TestTypedCallbacksConcurrent(t *testing.T) {
	const workers, calls = 8, 100
	interp := loadCallbacks(t, EnableCachedReg, 0)
	callbacks := make([]callbackSet, workers)
	for worker := range callbacks {
		callbacks[worker] = newCallbacks(t, interp, worker)
	}
	var wg sync.WaitGroup
	var failed [workers]bool
	for worker, functions := range callbacks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment, even := functions.increment, functions.even
			for call := 0; call < calls; call++ {
				increment()
				if even() != ((worker+call+1)%2 == 0) {
					failed[worker] = true
				}
			}
		}()
	}
	wg.Wait()
	for worker, functions := range callbacks {
		if got := functions.count(); got != worker+calls || failed[worker] {
			t.Errorf("worker %d: count = %d, incorrect predicate = %v", worker, got, failed[worker])
		}
	}
}

func TestTypedCallbackRecover(t *testing.T) {
	interp := loadCallbacks(t, EnableCachedReg, 0)
	for _, value := range []interface{}{"panic", nil} {
		marks := 0
		void := runFunc(t, interp, "RecoverVoid", value, func() { marks++ }).(func())
		boolean := runFunc(t, interp, "RecoverBool", value).(func() bool)
		for i := 0; i < 3; i++ {
			void()
			if !boolean() {
				t.Fatalf("bool callback failed to recover panic(%v)", value)
			}
		}
		panicFunc := runFunc(t, interp, "Panic", value).(func())
		if got, ok := catchPanic(panicFunc).(PanicError); !ok || got.Value != value {
			t.Fatalf("script panic = %v; want PanicError carrying %v", got, value)
		}
		if marks != 3 {
			t.Fatalf("void callback recovered %d times; want 3", marks)
		}
	}
	reentrant := runFunc(t, interp, "DeferredReentry", func(f func()) { f() }).(func() bool)
	if !reentrant() {
		t.Fatal("nested host callback changed deferred recover semantics")
	}
	// The host consumes its own stop sentinel.
	sentinel := new(int)
	stop := runFunc(t, interp, "Stop", func() { panic(sentinel) }).(func())
	for i := 0; i < 3; i++ {
		if got := catchPanic(stop); got != sentinel {
			t.Fatalf("host recovered %v; want original sentinel %p", got, sentinel)
		}
	}
	if !runFunc(t, interp, "InvokeBool", func() bool { return true }).(bool) {
		t.Fatal("interpreter unusable after host recovered sentinel")
	}
}

func TestCallbackAbort(t *testing.T) {
	for _, function := range []struct {
		name       string
		want, zero []interface{}
	}{
		{name: "Void"},
		{name: "Bool", want: []interface{}{true}, zero: []interface{}{false}},
		{name: "Int", want: []interface{}{42}, zero: []interface{}{0}},
		{name: "Multi", want: []interface{}{42, true}, zero: []interface{}{0, false}},
	} {
		for _, pool := range []struct {
			name      string
			threshold int
		}{
			{name: "pooled", threshold: 0},
			{name: "unpooled", threshold: 1 << 30},
		} {
			for _, state := range []struct {
				name   string
				warmup int
				during bool
			}{
				{name: "cold_before"},
				{name: "hot_before", warmup: 2},
				{name: "cold_during", during: true},
				{name: "hot_during", warmup: 2, during: true},
			} {
				t.Run(function.name+"/"+pool.name+"/"+state.name, func(t *testing.T) {
					interp := loadCallbacks(t, EnableCachedReg, pool.threshold)
					hooks, bodies := 0, 0
					abortInside := false
					runFunc(t, interp, "SetHooks", func() {
						hooks++
						if abortInside {
							interp.Abort()
						}
					}, func() { bodies++ })
					fn, ok := interp.GetFunc(function.name)
					if !ok {
						t.Fatalf("missing callback %s", function.name)
					}
					for i := 0; i < state.warmup; i++ {
						checkAbortCall(t, fn, function.want)
					}
					if hooks != state.warmup || bodies != state.warmup {
						t.Fatalf("warmup: hooks = %d, bodies = %d; want %d", hooks, bodies, state.warmup)
					}
					wantHooks := state.warmup
					if state.during {
						abortInside = true
						wantHooks++
					} else {
						interp.Abort()
					}
					for i := 0; i < 3; i++ {
						checkAbortCall(t, fn, function.zero)
						if hooks != wantHooks || bodies != state.warmup {
							t.Fatalf("after abort: hooks = %d, bodies = %d; want %d, %d", hooks, bodies, wantHooks, state.warmup)
						}
					}
				})
			}
		}
	}
}

func TestCallbackResults(t *testing.T) {
	interp := loadCallbacks(t, EnableCachedReg, 0)
	var retained Tuple
	want := Tuple{42, true}
	abort := false
	runFunc(t, interp, "SetHooks", func() {
		if retained != nil && !reflect.DeepEqual(retained, want) {
			t.Errorf("retained result before Return = %v, want %v", retained, want)
		}
		if abort {
			interp.Abort()
		}
	}, func() {})
	runFunc(t, interp, "Multi")
	retained = runFunc(t, interp, "Multi").(Tuple)
	fn, ok := interp.GetFunc("Multi")
	if !ok {
		t.Fatal("missing Multi callback")
	}
	checkAbortCall(t, fn, []interface{}{42, true})
	abort = true
	checkAbortCall(t, fn, []interface{}{0, false})
	if got := runFunc(t, interp, "Multi"); !reflect.DeepEqual(got, Tuple{nil, nil}) {
		t.Errorf("aborted tuple = %v, want two empty slots", got)
	}
	if !reflect.DeepEqual(retained, want) {
		t.Errorf("retained result after Abort = %v, want %v", retained, want)
	}
}

func BenchmarkTypedCallback(b *testing.B) {
	for _, name := range []string{"void", "bool", "create_void", "create_bool", "create_int", "dynamic_void_128", "dynamic_bool_128", "stop"} {
		b.Run(name, func(b *testing.B) {
			interp := loadCallbacks(b, EnableCachedReg, 0)
			callbacks := newCallbacks(b, interp, 0)
			var callback interface{} = callbacks.increment
			switch name {
			case "bool":
				callback = callbacks.even
			case "create_void":
				callback = runFunc(b, interp, "CreateAndCallVoid", func(f func()) { f() })
			case "create_bool":
				callback = runFunc(b, interp, "CreateAndCallBool", func(f func() bool) bool { return f() })
			case "create_int":
				callback = runFunc(b, interp, "CreateAndCallInt", func(f func() int) int { return f() })
			case "dynamic_void_128":
				callback = runFunc(b, interp, "DynamicVoid", callbacks.increment)
			case "dynamic_bool_128":
				callback = runFunc(b, interp, "DynamicBool", callbacks.even)
			case "stop":
				sentinel := new(int)
				callback = runFunc(b, interp, "Stop", func() { panic(sentinel) })
			}
			b.ReportAllocs()
			switch callback := callback.(type) {
			case func() int:
				for b.Loop() {
					callback()
				}
			case func() bool:
				for b.Loop() {
					callback()
				}
			case func():
				if name == "stop" {
					for b.Loop() {
						catchPanic(callback)
					}
				} else {
					for b.Loop() {
						callback()
					}
				}
			}
		})
	}
}

type callbackSet struct {
	increment func()
	even      func() bool
	count     func() int
}

func loadCallbacks(t testing.TB, mode Mode, threshold int) *Interp {
	t.Helper()
	ctx := NewContext(mode | SupportMultipleInterp)
	ctx.SetLeastCallForEnablePool(threshold)
	interp, err := ctx.LoadInterp("callbacks.go", typedCallbackSource)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(interp.UnsafeRelease)
	return interp
}

func runFunc(t testing.TB, interp *Interp, name string, args ...interface{}) interface{} {
	t.Helper()
	v, err := interp.RunFunc(name, args...)
	if err != nil {
		t.Fatal(err)
	}
	return v
}

func newCallbacks(t testing.TB, interp *Interp, start int) callbackSet {
	t.Helper()
	v := runFunc(t, interp, "MakeCallbacks", start).([]interface{})
	return callbackSet{v[0].(func()), v[1].(func() bool), v[2].(func() int)}
}

func catchPanic(f func()) (value interface{}) {
	defer func() { value = recover() }()
	f()
	return nil
}

func checkAbortCall(t *testing.T, fn interface{}, want []interface{}) {
	t.Helper()
	var got []interface{}
	switch fn := fn.(type) {
	case func():
		fn()
	case func() bool:
		got = []interface{}{fn()}
	case func() int:
		got = []interface{}{fn()}
	case func() (int, bool):
		n, ok := fn()
		got = []interface{}{n, ok}
	default:
		t.Fatalf("unexpected callback type %T", fn)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("callback result = %v; want %v", got, want)
	}
}

const typedCallbackSource = `package main

import (
	"reflect"
	"runtime"
)

func MakeCallbacks(start int) []interface{} {
	n := start
	return []interface{}{
		func() { n++ },
		func() bool { return n % 2 == 0 },
		func() int { return n },
	}
}

func InvokeVoid(f func()) { f() }
func InvokeBool(f func() bool) bool { return f() }

func Reentrant(invoke func(func()), start int) func() bool {
	n := start
	return func() bool {
		invoke(func() { n++ })
		return n % 2 == 0
	}
}

func Signatures(call func(func()), predicate func(func() bool) bool) bool {
	type Action func()
	type Predicate func() bool
	n := 0
	a := Action(func() { n++ })
	p := Predicate(func() bool { return n == 1 })
	call(a)
	return predicate(p)
}

func Fallbacks() []interface{} {
	type Flag bool
	return []interface{}{
		func() Flag { return true },
		func(x int) int { return x + 1 },
		func(xs ...int) int { return xs[0] + xs[1] },
		func() (int, bool) { return 42, true },
	}
}

func RecoverVoid(value interface{}, mark func()) func() {
	return func() {
		defer func() { recover(); mark() }()
		panic(value)
	}
}

func RecoverBool(value interface{}) func() bool {
	return func() (ok bool) {
		defer func() { recover(); ok = true }()
		panic(value)
	}
}

func Stop(stop func()) func() { return func() { stop() } }
func Panic(value interface{}) func() { return func() { panic(value) } }

func DeferredReentry(invoke func(func())) func() bool {
	return func() (ok bool) {
		defer func() {
			nestedRecovered := false
			invoke(func() { nestedRecovered = recover() != nil })
			ok = !nestedRecovered && recover() == "panic"
		}()
		panic("panic")
	}
}

func PointerNames(a func(), b func() bool) []string {
	for _, f := range []interface{}{(func())(nil), (func() bool)(nil)} {
		v := reflect.ValueOf(f)
		if v.Pointer() != 0 || v.UnsafePointer() != nil { panic("nil function pointer") }
	}
	return []string{
		runtime.FuncForPC(reflect.ValueOf(a).Pointer()).Name(),
		runtime.FuncForPC(reflect.ValueOf(b).Pointer()).Name(),
	}
}

func CreateAndCallVoid(call func(func())) func() {
	n := 0
	return func() { call(func() { n++ }) }
}

func CreateAndCallBool(call func(func() bool) bool) func() bool {
	n := 0
	return func() bool { return call(func() bool { n++; return n % 2 == 0 }) }
}

func CreateAndCallInt(call func(func() int) int) func() int {
	n := 0
	return func() int { return call(func() int { n++; return n }) }
}

func DynamicVoid(f func()) func() {
	return func() { for i := 0; i < 128; i++ { f() } }
}

func DynamicBool(f func() bool) func() bool {
	return func() bool {
		var result bool
		for i := 0; i < 128; i++ { result = f() }
		return result
	}
}

var hook, body func()

func SetHooks(h, b func()) { hook, body = h, b }
func Void() { hook(); body() }
func Bool() bool { hook(); body(); return true }
func Int() int { hook(); body(); return 42 }
func Multi() (int, bool) { hook(); body(); return 42, true }

func main() {}
`
