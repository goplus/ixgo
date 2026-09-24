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
