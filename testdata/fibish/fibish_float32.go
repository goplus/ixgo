// run

//go:build !wasm

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at https://go.dev/LICENSE.

package main

import "fmt"

// Test that register results are correctly returned (and passed)

//go:registerparams
//go:noinline
func f(x float32) (float32, float32) {

	if x < 3 {
		return 0, x
	}

	a, b := f(x - 2)
	c, d := f(x - 1)
	return a + d, b + c
}

func main() {
	x := float32(40)
	a, b := f(x)
	fmt.Printf("f(%.0f)=%.0f,%.0f\n", x, a, b)
}
