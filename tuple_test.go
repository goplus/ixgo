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
	"strings"
	"sync"
	"testing"
)

func TestTupleAfterReuse(t *testing.T) {
	interp := loadTuple(t)
	for _, test := range []struct {
		name      string
		first     int
		later     int
		laterWant Tuple
	}{
		{name: "Pair", first: 11, later: 23, laterWant: Tuple{23, 24}},
		{name: "Recursive", first: 31, later: 42, laterWant: Tuple{42, 43}},
		{name: "Recovering", first: 11, later: -1, laterWant: Tuple{-1, -2}},
		{name: "EnvPair", first: 11, later: 23, laterWant: Tuple{23, 24}},
	} {
		t.Run(test.name, func(t *testing.T) {
			runTuple(t, interp, test.name, test.first)
			result := runTuple(t, interp, test.name, test.first)
			want := Tuple{test.first, test.first + 1}
			if !reflect.DeepEqual(result, want) {
				t.Fatalf("result before reuse = %v, want %v", result, want)
			}
			if got := runTuple(t, interp, test.name, test.later); !reflect.DeepEqual(got, test.laterWant) {
				t.Errorf("later result = %v, want %v", got, test.laterWant)
			}
			if !reflect.DeepEqual(result, want) {
				t.Errorf("retained result = %v, want %v", result, want)
			}
		})
	}
}

func TestTupleConcurrent(t *testing.T) {
	interp := loadTuple(t)
	runTuple(t, interp, "Pair", 0)
	const workers, calls = 4, 40
	results := make([]Tuple, workers*calls)
	var wg sync.WaitGroup
	for worker := 0; worker < workers; worker++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			for n := 0; n < calls; n++ {
				want := worker*calls + n
				value, err := interp.RunFunc("Pair", want)
				if err != nil {
					t.Error(err)
					return
				}
				results[want] = value.(Tuple)
			}
		}(worker)
	}
	wg.Wait()
	for want, result := range results {
		if len(result) != 2 || result[0] != want || result[1] != want+1 {
			t.Fatalf("retained tuple = %v, want (%d, %d)", result, want, want+1)
		}
	}
}

func TestTupleClosures(t *testing.T) {
	interp := loadTuple(t)
	const calls = 64
	results := make([]Tuple, calls)
	for n := range results {
		results[n] = runTuple(t, interp, "Capture", n)
	}
	for n, result := range results {
		for slot := 0; slot < 2; slot++ {
			if got := result[slot].(func() int)(); got != n+slot {
				t.Fatalf("retained closure %d/%d returned %d, want %d", n, slot, got, n+slot)
			}
		}
	}
}

func TestFrameCleanup(t *testing.T) {
	interp := loadFrame(t, EnableCachedReg)
	p := interp.funcs[interp.mainpkg.Func("Work")]
	if len(p.Fn.Locals) != 1 || p.Fn.Locals[0].Heap || len(p.cacheRegs) == 0 {
		t.Fatal("fixture needs one reusable local and a hidden cache")
	}
	argReg := p.regIndex(p.Fn.Params[0])
	localReg := p.regIndex(p.Fn.Locals[0])
	var active *frame
	first := p.Instrs[0]
	p.Instrs[0] = func(fr *frame) {
		active = fr
		fr.phiVals = []value{new(int)}
		first(fr)
	}
	for call := 0; call < 3; call++ {
		var local value
		result := runFrame(t, interp, "Work", func() { local = active.stack[localReg] })
		if result != 7 || local == nil {
			t.Fatalf("call %d: result = %v, local = %v", call, result, local)
		}
		if active.caller != nil || active.callee != nil || active._defer != nil || active._panic != nil || active.deferid != 0 {
			t.Fatalf("call %d: returned frame retains its call chain", call)
		}
		if active.stack[argReg] != nil {
			t.Errorf("call %d: callback argument retained", call)
		}
		if active.stack[localReg] != local || !reflect.ValueOf(local).Elem().IsZero() {
			t.Errorf("call %d: reusable local was replaced or retained a value", call)
		}
		for _, reg := range p.cacheRegs {
			if active.stack[reg] != nil {
				t.Errorf("call %d: hidden cache %d retained %T", call, reg, active.stack[reg])
			}
		}
		for _, v := range active.phiVals {
			if v != nil {
				t.Errorf("call %d: phi scratch retained %T", call, v)
			}
		}
	}
}

func TestFrameResults(t *testing.T) {
	interp := loadFrame(t, 0)
	retained := runFrame(t, interp, "Values", 7).(Tuple)
	pointer := runFrame(t, interp, "LocalPointer", 7)
	for n := 20; n < 24; n++ {
		runFrame(t, interp, "Values", n)
		runFrame(t, interp, "LocalPointer", n)
	}
	check := func(record reflect.Value, want int) {
		t.Helper()
		if got := record.FieldByName("Numbers").Interface().([]int); !reflect.DeepEqual(got, []int{want, want + 1}) {
			t.Errorf("numbers = %v; want [%d %d]", got, want, want+1)
		}
		if got := record.FieldByName("Lookup").Interface().(map[string]int)["n"]; got != want {
			t.Errorf("map value = %d; want %d", got, want)
		}
		if got := record.FieldByName("Any").Interface(); got != want {
			t.Errorf("interface value = %v; want %d", got, want)
		}
	}
	check(reflect.ValueOf(retained[0]), 7)
	check(reflect.ValueOf(retained[1]).Elem(), 9)
	if got := retained[2].(func() int)(); got != 7 {
		t.Errorf("captured value = %d; want 7", got)
	}
	if got := *pointer.(*int); got != 7 {
		t.Errorf("returned local pointer = %d; want 7", got)
	}
}

func TestFramePanicStack(t *testing.T) {
	interp := loadFrame(t, 0)
	runFrame(t, interp, "Crash", 1)
	_, err := interp.RunFunc("Crash", -1)
	panicked, ok := err.(FatalError)
	if !ok {
		t.Fatalf("panic = %v; want FatalError", err)
	}
	stack := string(panicked.Stack())
	for _, name := range []string{"main.Crash", "main.crashLeaf", "frame.go:"} {
		if !strings.Contains(stack, name) {
			t.Errorf("panic stack missing %q:\n%s", name, stack)
		}
	}
	runFrame(t, interp, "Crash", 2)
	if string(panicked.Stack()) != stack {
		t.Error("retained panic stack changed after another call")
	}
}

func loadTuple(t *testing.T) *Interp {
	t.Helper()
	ctx := NewContext(SupportMultipleInterp)
	ctx.SetLeastCallForEnablePool(0)
	interp, err := ctx.LoadInterp("tuple.go", tupleSource)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(interp.UnsafeRelease)
	return interp
}

func runTuple(t *testing.T, interp *Interp, name string, n int) Tuple {
	t.Helper()
	v, err := interp.RunFunc(name, n)
	if err != nil {
		t.Fatal(err)
	}
	return v.(Tuple)
}

func loadFrame(t testing.TB, mode Mode) *Interp {
	t.Helper()
	ctx := NewContext(SupportMultipleInterp | mode)
	ctx.SetLeastCallForEnablePool(0)
	interp, err := ctx.LoadInterp("frame.go", frameSource)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(interp.UnsafeRelease)
	return interp
}

func runFrame(t testing.TB, interp *Interp, name string, args ...interface{}) interface{} {
	t.Helper()
	v, err := interp.RunFunc(name, args...)
	if err != nil {
		t.Fatal(err)
	}
	return v
}

const tupleSource = `package main
func Pair(n int) (int, int) { return n, n + 1 }
func Recursive(n int) (int, int) {
	if n <= 10 { return n, n + 1 }
	a, b := Recursive(n - 10)
	return a + 10, b + 10
}
func Recovering(n int) (a, b int) {
	defer func() {
		if recover() != nil { a, b = -1, -2 }
	}()
	if n < 0 { panic("boom") }
	return n, n + 1
}
func EnvPair(n int) (int, int) { return pairFunc(n)() }
func pairFunc(n int) func() (int, int) {
	return func() (int, int) { return n, n + 1 }
}
func Capture(n int) (func() int, func() int) {
	data := new([256]byte)
	data[0], data[1] = byte(n), byte(n + 1)
	return func() int { return int(data[0]) }, func() int { return int(data[1]) }
}
`

const frameSource = `package main
var bias int
func Add(n int) int { return n + bias }
func Work(cb func()) int {
	bias = 3
	var local struct { callbacks [1]func(); value int }
	local.callbacks[0] = cb
	local.value = 4
	if local.callbacks[0] != nil { local.callbacks[0]() }
	return Add(local.value)
}

type Data struct {
	Numbers []int
	Lookup map[string]int
	Any interface{}
}
func Values(n int) (Data, *Data, func() int) {
	var local Data
	local.Numbers = []int{n, n + 1}
	local.Lookup = map[string]int{"n": n}
	local.Any = n
	ptr := &Data{[]int{n + 2, n + 3}, map[string]int{"n": n + 2}, n + 2}
	return local, ptr, func() int { return n }
}
func LocalPointer(n int) *int {
	var local struct{ N int }
	local.N = n
	return &local.N
}

func Crash(n int) int { return crashLeaf(n) }
func crashLeaf(n int) int {
	var p *int
	if n < 0 { return *p }
	return n
}
`
