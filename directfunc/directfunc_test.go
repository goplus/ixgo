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

package directfunc_test

import (
	"testing"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/directfunc"
	"github.com/visualfc/funcval"
)

func TestImportRegistersDirectFuncs(t *testing.T) {
	ctx := ixgo.NewContext(0)
	interp, err := ctx.LoadInterp("main.go", `package main
type boom struct{}
func (boom) Error() string { return "boom" }
func Make() []interface{} {
	n := 0
	return []interface{}{
		func() { n++ },
		func() bool { return n%2 == 0 },
		func() int { return n },
		func() int64 { return int64(n) },
		func() string { return "ok" },
		func() error { return boom{} },
		func(x int) { n += x },
	}
}
func main() {}
`)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(interp.UnsafeRelease)
	v, err := interp.RunFunc("Make")
	if err != nil {
		t.Fatal(err)
	}
	values := v.([]interface{})
	if funcval.IsSupport {
		wantBridges := []int{0, 0, 0, 0, 0, 0, 1}
		for i, fn := range values {
			_, n := funcval.Get(fn)
			if n != wantBridges[i] {
				t.Errorf("function %d (%T) uses %d reflection bridges; want %d", i, fn, n, wantBridges[i])
			}
		}
	}
	void := values[0].(func())
	boolean := values[1].(func() bool)
	count := values[2].(func() int)
	i64 := values[3].(func() int64)
	str := values[4].(func() string)
	errcb := values[5].(func() error)
	add := values[6].(func(int))
	if !boolean() {
		t.Fatal("even(0) = false")
	}
	void()
	if boolean() || count() != 1 || i64() != 1 {
		t.Fatalf("after increment: even = %v, count = %d, int64 = %d", boolean(), count(), i64())
	}
	if str() != "ok" {
		t.Fatalf("string function = %q; want ok", str())
	}
	if got := errcb(); got == nil || got.Error() != "boom" {
		t.Fatalf("error function = %v; want boom", got)
	}
	add(4)
	if count() != 5 {
		t.Fatalf("after add(4): count = %d; want 5", count())
	}
}
