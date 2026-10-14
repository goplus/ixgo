// run

// Copyright 2026 The GoPlus Authors (goplus.org). All rights reserved.
// Licensed under the Apache License, Version 2.0.

package main

import "fmt"

// f computes the same pair as fibish.go, retaining the previous two pairs.
func f(x int) (int, int) {
	if x < 3 {
		return 0, x
	}

	// (a, b) = f(i-2), (c, d) = f(i-1).
	a, b := 0, 1
	c, d := 0, 2
	for i := 3; i <= x; i++ {
		a, b, c, d = c, d, a+d, b+c
	}
	return c, d
}

func main() {
	x := 40
	a, b := f(x)
	fmt.Printf("f(%d)=%d,%d\n", x, a, b)
}
