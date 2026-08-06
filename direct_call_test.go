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
	"go/token"
	"go/types"
	"reflect"
	"sync/atomic"
	"testing"

	directcalltest "github.com/goplus/ixgo/testdata/direct_call/host"
	"golang.org/x/tools/go/ssa"
)

const (
	testDirectCallPkgPath     = "ixgo.test/direct_call"
	testDirectCallSelector    = "Add"
	testDirectCallExternalKey = testDirectCallPkgPath + ".Add"
	testDirectCallHostPkgPath = "github.com/goplus/ixgo/testdata/direct_call/host"
	testCounterSelector       = "(*Counter).Value"
	testCounterExternalKey    = "(*" + testDirectCallHostPkgPath + ".Counter).Value"
	testValueCounterSelector  = "ValueCounter.Value"
	testValuePointerSelector  = "(*ValueCounter).Value"
	testRecorderSelector      = "(*Recorder).Record"
)

func directCallAdd(a, b int) int {
	return a + b
}

func directCallWrongAdd(a, b int) int {
	return a + b + 1
}

func directCallWrongCounterValue(*directcalltest.Counter) int {
	return -1
}

func registerStaticDirectCallPackage(binding DirectCallAdapter) {
	RegisterPackage(&Package{
		Name:          "directcall",
		Path:          testDirectCallPkgPath,
		Interfaces:    map[string]reflect.Type{},
		NamedTypes:    map[string]reflect.Type{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{"Add": reflect.ValueOf(directCallAdd)},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
	RegisterDirectCalls(testDirectCallPkgPath, map[string]DirectCallAdapter{
		testDirectCallSelector: binding,
	})
}

func registerHostDirectCalls(bindings map[string]DirectCallAdapter) {
	RegisterPackage(&Package{
		Name:       "host",
		Path:       testDirectCallHostPkgPath,
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Counter":         reflect.TypeOf(directcalltest.Counter(0)),
			"FallbackCounter": reflect.TypeOf(directcalltest.FallbackCounter(0)),
			"Recorder":        reflect.TypeOf(directcalltest.Recorder{}),
			"ValueCounter":    reflect.TypeOf(directcalltest.ValueCounter(0)),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
	RegisterDirectCalls(testDirectCallHostPkgPath, bindings)
}

func clearExternalCallOverride(t testing.TB, key string) {
	t.Helper()
	RegisterExternal(key, nil)
	t.Cleanup(func() { RegisterExternal(key, nil) })
}

func TestDirectCallContext(t *testing.T) {
	fr := &frame{stack: []value{nil, nil}}
	ctx := DirectCallContext{frame: fr, result: 1, args: []register{0}}
	if got := DirectCallArg[*int](ctx, 0); got != nil {
		t.Fatalf("nil argument = %v; want nil", got)
	}
	ctx.SetResult(42)
	if got := fr.reg(1); got != 42 {
		t.Fatalf("result = %v; want 42", got)
	}
}

func directCallSignature(fn *ssa.Function) *types.Signature {
	sig := fn.Signature
	if sig.Recv() != nil {
		sig = wrapMethodType(sig)
	}
	return sig
}

func TestDirectCallSignatureIncludesMethodReceiver(t *testing.T) {
	recv := types.NewVar(token.NoPos, nil, "recv", types.Typ[types.Int])
	param := types.NewVar(token.NoPos, nil, "v", types.Typ[types.String])
	result := types.NewVar(token.NoPos, nil, "", types.Typ[types.Bool])
	sig := types.NewSignature(recv, types.NewTuple(param), types.NewTuple(result), false)

	got := directCallSignature(&ssa.Function{Signature: sig})
	if got.Recv() != nil {
		t.Fatalf("direct-call signature still has receiver %v", got.Recv())
	}
	if got.Params().Len() != 2 || got.Params().At(0) != recv || got.Params().At(1) != param {
		t.Fatalf("direct-call params = %v; want (%v, %v)", got.Params(), recv, param)
	}
	if got.Results() != sig.Results() {
		t.Fatalf("direct-call results = %v; want %v", got.Results(), sig.Results())
	}
}

func TestDirectCallMethodKeyUsesGoSymbol(t *testing.T) {
	tests := []struct {
		name string
		typ  reflect.Type
		want string
	}{
		{name: "value", typ: reflect.TypeOf(directcalltest.ValueCounter(0)), want: "(" + testDirectCallHostPkgPath + ".ValueCounter).Value"},
		{name: "pointer", typ: reflect.TypeOf((*directcalltest.Counter)(nil)), want: testCounterExternalKey},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key, ok := directCallMethodKey(test.typ, "Value")
			if !ok {
				t.Fatal("directCallMethodKey did not recognize a named receiver")
			}
			if key != test.want {
				t.Fatalf("method key = %q; want %q", key, test.want)
			}
		})
	}
}

func TestDirectCallSymbolQualifiesMethods(t *testing.T) {
	const pkgPath = "example.com/p"
	tests := map[string]string{
		"T.Value":    "(example.com/p.T).Value",
		"(*T).Value": "(*example.com/p.T).Value",
		"Add":        "example.com/p.Add",
	}
	for selector, want := range tests {
		if got := directCallSymbol(pkgPath, selector); got != want {
			t.Errorf("directCallSymbol(%q) = %q; want %q", selector, got, want)
		}
	}
}

func TestDirectCallRegistrationMergesAndReplaces(t *testing.T) {
	const pkgPath = "ixgo.test/direct_call_registration"
	first := func(DirectCallContext) {}
	replacement := func(DirectCallContext) {}
	valueMethod := func(DirectCallContext) {}
	pointerMethod := func(DirectCallContext) {}

	RegisterDirectCalls(pkgPath, map[string]DirectCallAdapter{
		"Add":   first,
		"(T).M": valueMethod,
	})
	RegisterDirectCalls(pkgPath, map[string]DirectCallAdapter{
		"Add":    replacement,
		"(*T).M": pointerMethod,
	})

	tests := []struct {
		key  string
		want DirectCallAdapter
	}{
		{key: pkgPath + ".Add", want: replacement},
		{key: "(" + pkgPath + ".T).M", want: valueMethod},
		{key: "(*" + pkgPath + ".T).M", want: pointerMethod},
	}
	interp := &Interp{ctx: &Context{}}
	for _, test := range tests {
		got, ok := lookupDirectCallBinding(interp, test.key)
		if !ok {
			t.Fatalf("binding %q was not registered", test.key)
		}
		if got == nil {
			t.Fatalf("binding %q has nil adapter", test.key)
		}
	}
}

func TestStaticPackageCallUsesDirectCall(t *testing.T) {
	clearExternalCallOverride(t, testDirectCallExternalKey)
	var calls int
	registerStaticDirectCallPackage(func(ctx DirectCallContext) {
		calls++
		ctx.SetResult(directCallAdd(DirectCallArg[int](ctx, 0), DirectCallArg[int](ctx, 1)))
	})

	const source = `package main

import directcall "ixgo.test/direct_call"

func main() {
	if got := directcall.Add(20, 22); got != 42 {
		panic(got)
	}
}
`
	if _, err := NewContext(0).RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
	if calls != 1 {
		t.Fatalf("direct adapter calls = %d; want 1", calls)
	}
}

func TestStaticMethodCallUsesDirectCall(t *testing.T) {
	clearExternalCallOverride(t, testCounterExternalKey)
	var calls int
	registerHostDirectCalls(map[string]DirectCallAdapter{
		testCounterSelector: func(ctx DirectCallContext) {
			calls++
			ctx.SetResult((*directcalltest.Counter).Value(DirectCallArg[*directcalltest.Counter](ctx, 0)))
		},
	})

	const source = `package main

import host "github.com/goplus/ixgo/testdata/direct_call/host"

func main() {
	counter := host.Counter(7)
	if got := counter.Value(); got != 7 {
		panic(got)
	}
}
`
	if _, err := NewContext(0).RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
	if calls != 1 {
		t.Fatalf("direct adapter calls = %d; want 1", calls)
	}
}

func TestInvokeUsesDirectCallsAndReflectFallback(t *testing.T) {
	var pointerCalls, valueCalls, valuePointerCalls int
	registerHostDirectCalls(map[string]DirectCallAdapter{
		testCounterSelector: func(ctx DirectCallContext) {
			pointerCalls++
			ctx.SetResult((*directcalltest.Counter).Value(DirectCallArg[*directcalltest.Counter](ctx, 0)))
		},
		testValueCounterSelector: func(ctx DirectCallContext) {
			valueCalls++
			ctx.SetResult(directcalltest.ValueCounter.Value(DirectCallArg[directcalltest.ValueCounter](ctx, 0)))
		},
		testValuePointerSelector: func(ctx DirectCallContext) {
			valuePointerCalls++
			ctx.SetResult((*directcalltest.ValueCounter).Value(DirectCallArg[*directcalltest.ValueCounter](ctx, 0)))
		},
	})

	const source = `package main

import host "github.com/goplus/ixgo/testdata/direct_call/host"

type counter interface {
	Value() int
}

func value(v counter) int {
	return v.Value()
}

func main() {
	pointer := host.Counter(7)
	valueCounter := host.ValueCounter(8)
	valuePointer := host.ValueCounter(10)
	fallback := host.FallbackCounter(9)
	if value(&pointer) != 7 || value(valueCounter) != 8 || value(&valuePointer) != 10 || value(&fallback) != 9 {
		panic("unexpected counter value")
	}
}
`
	if _, err := NewContext(0).RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
	if pointerCalls != 1 || valueCalls != 1 || valuePointerCalls != 1 {
		t.Fatalf("direct adapter calls = (%d, %d, %d); want (1, 1, 1)", pointerCalls, valueCalls, valuePointerCalls)
	}
}

func TestInvokeUsesRegisteredSelector(t *testing.T) {
	var directCalls int
	registerHostDirectCalls(map[string]DirectCallAdapter{
		testCounterSelector: func(ctx DirectCallContext) {
			directCalls++
			ctx.SetResult(-1)
		},
	})

	const source = `package main

import host "github.com/goplus/ixgo/testdata/direct_call/host"

type counter interface {
	Value() int
}

func value(v counter) int {
	return v.Value()
}

func main() {
	counter := host.Counter(7)
	if got := value(&counter); got != -1 {
		panic(got)
	}
}
`
	if _, err := NewContext(0).RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
	if directCalls != 1 {
		t.Fatalf("direct adapter calls = %d; want 1", directCalls)
	}
}

func TestDirectCallGoAndInvokeDeferUseRegularPath(t *testing.T) {
	var directCalls int32
	registerHostDirectCalls(map[string]DirectCallAdapter{
		testRecorderSelector: func(ctx DirectCallContext) {
			atomic.AddInt32(&directCalls, 1)
			(*directcalltest.Recorder).Record(
				DirectCallArg[*directcalltest.Recorder](ctx, 0),
				DirectCallArg[int](ctx, 1),
			)
		},
	})

	const source = `package main

import host "github.com/goplus/ixgo/testdata/direct_call/host"

type recorder interface {
	Record(int)
}

func main() {
	target := host.Recorder{Values: make(chan int, 2)}
	go target.Record(1)
	if got := <-target.Values; got != 1 {
		panic(got)
	}

	var sink recorder = &target
	defer func() {
		if got := <-target.Values; got != 2 {
			panic(got)
		}
	}()
	defer sink.Record(2)
}
`
	if _, err := NewContext(0).RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
	if got := atomic.LoadInt32(&directCalls); got != 0 {
		t.Fatalf("go/defer direct adapter calls = %d; want 0", got)
	}
}

func TestContextOverrideSkipsStaticDirectCall(t *testing.T) {
	clearExternalCallOverride(t, testDirectCallExternalKey)
	var directCalls, overrideCalls int
	registerStaticDirectCallPackage(func(ctx DirectCallContext) {
		directCalls++
		ctx.SetResult(-1)
	})
	ctx := NewContext(0)
	ctx.RegisterExternal(testDirectCallExternalKey, func(a, b int) int {
		overrideCalls++
		return a + b + 1
	})

	const source = `package main

import directcall "ixgo.test/direct_call"

func main() {
	if got := directcall.Add(20, 22); got != 43 {
		panic(got)
	}
}
`
	if _, err := ctx.RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
	if directCalls != 0 || overrideCalls != 1 {
		t.Fatalf("calls: direct=%d override=%d; want 0, 1", directCalls, overrideCalls)
	}
}

func TestRegisteredExternalSkipsStaticDirectCall(t *testing.T) {
	clearExternalCallOverride(t, testDirectCallExternalKey)
	var directCalls, overrideCalls int
	registerStaticDirectCallPackage(func(ctx DirectCallContext) {
		directCalls++
		ctx.SetResult(-1)
	})
	RegisterExternal(testDirectCallExternalKey, func(a, b int) int {
		overrideCalls++
		return a + b + 1
	})

	const source = `package main

import directcall "ixgo.test/direct_call"

func main() {
	if got := directcall.Add(20, 22); got != 43 {
		panic(got)
	}
}
`
	if _, err := NewContext(0).RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
	if directCalls != 0 || overrideCalls != 1 {
		t.Fatalf("calls: direct=%d override=%d; want 0, 1", directCalls, overrideCalls)
	}
}

func TestStaticDirectCallUsesRegisteredSelector(t *testing.T) {
	clearExternalCallOverride(t, testDirectCallExternalKey)
	var directCalls int
	registerStaticDirectCallPackage(func(ctx DirectCallContext) {
		directCalls++
		ctx.SetResult(-1)
	})

	const source = `package main

import directcall "ixgo.test/direct_call"

func main() {
	if got := directcall.Add(20, 22); got != -1 {
		panic(got)
	}
}
`
	if _, err := NewContext(0).RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
	if directCalls != 1 {
		t.Fatalf("direct adapter calls = %d; want 1", directCalls)
	}
}

func TestStaticDirectCallIgnoresRemovedTargetSignature(t *testing.T) {
	clearExternalCallOverride(t, testDirectCallExternalKey)
	var directCalls int
	registerStaticDirectCallPackage(func(ctx DirectCallContext) {
		directCalls++
		ctx.SetResult(-1)
	})

	const source = `package main

import directcall "ixgo.test/direct_call"

func main() {
	if got := directcall.Add(20, 22); got != -1 {
		panic(got)
	}
}
`
	if _, err := NewContext(0).RunFile("main.go", source, nil); err != nil {
		t.Fatal(err)
	}
	if directCalls != 1 {
		t.Fatalf("direct adapter calls = %d; want 1", directCalls)
	}
}

func TestDirectCallPreservesDeferFrame(t *testing.T) {
	adapter := DirectCallAdapter(func(ctx DirectCallContext) {
		ctx.SetResult(directCallAdd(DirectCallArg[int](ctx, 0), DirectCallArg[int](ctx, 1)))
	})
	interp := &Interp{}
	fr := &frame{stack: []value{20, 22, nil}, deferid: 7}
	interp.invokeDirectCall(fr, adapter, 2, []register{0, 1})
	if got := fr.reg(2); got != 42 {
		t.Fatalf("result = %v; want 42", got)
	}
	got, ok := interp.deferMap.Load(uint64(7))
	if !ok || got != fr {
		t.Fatalf("defer frame = (%v, %v); want (%p, true)", got, ok, fr)
	}
}

func BenchmarkDirectCall(b *testing.B) {
	fnValue := reflect.ValueOf(directCallAdd)
	registerStaticDirectCallPackage(func(ctx DirectCallContext) {
		ctx.SetResult(directCallAdd(DirectCallArg[int](ctx, 0), DirectCallArg[int](ctx, 1)))
	})
	RegisterExternal(testDirectCallExternalKey, nil)
	interp := &Interp{preloadTypes: map[types.Type]reflect.Type{
		types.Typ[types.Int]: reflect.TypeOf(int(0)),
	}}
	fr := &frame{stack: []value{20, 22, nil}}
	ia := []register{0, 1}

	b.Run("direct", func(b *testing.B) {
		adapter := DirectCallAdapter(func(ctx DirectCallContext) {
			ctx.SetResult(directCallAdd(DirectCallArg[int](ctx, 0), DirectCallArg[int](ctx, 1)))
		})
		call := func(fr *frame) {
			interp.invokeDirectCall(fr, adapter, 2, ia)
		}
		b.ReportAllocs()
		for b.Loop() {
			call(fr)
		}
	})
	b.Run("reflect", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			interp.callExternalByStack(fr, fnValue, 2, ia)
		}
	})

	registerHostDirectCalls(map[string]DirectCallAdapter{
		testCounterSelector: func(ctx DirectCallContext) {
			ctx.SetResult((*directcalltest.Counter).Value(DirectCallArg[*directcalltest.Counter](ctx, 0)))
		},
	})
	result := types.NewVar(token.NoPos, nil, "", types.Typ[types.Int])
	method := types.NewFunc(token.NoPos, nil, "Value", types.NewSignature(nil, types.NewTuple(), types.NewTuple(result), false))
	methodCall := &ssa.CallCommon{Method: method}
	var counter directcalltest.Counter
	methodFrame := &frame{stack: []value{&counter, nil}}
	methodArgs := []register{0}

	b.Run("dynamic_direct", func(b *testing.B) {
		call := makeCallMethodInstr(interp, nil, methodCall, 1, 0, nil)
		b.ReportAllocs()
		for b.Loop() {
			call(methodFrame)
		}
	})
	var fallback directcalltest.FallbackCounter
	fallbackFrame := &frame{stack: []value{&fallback, nil}}
	b.Run("dynamic_fallback", func(b *testing.B) {
		call := makeCallMethodInstr(interp, nil, methodCall, 1, 0, nil)
		b.ReportAllocs()
		for b.Loop() {
			call(fallbackFrame)
		}
	})
	b.Run("dynamic_reflect", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			fn, ok := findExternMethod(reflect.TypeOf(methodFrame.reg(0)), "Value")
			if !ok {
				b.Fatal("method not found")
			}
			interp.callExternalByStack(methodFrame, fn, 1, methodArgs)
		}
	})
}
