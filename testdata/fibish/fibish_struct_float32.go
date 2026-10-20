//go:build !wasm

// run
package main

import "fmt"

type fibonacci struct{}

func (fibonacci) fib(x int) (float32, float32) {
	if x < 3 {
		return 0, float32(x)
	}
	a, b := fibonacci{}.fib(x - 2)
	c, d := fibonacci{}.fib(x - 1)
	return a + d, b + c
}

func main() {
	a, b := fibonacci{}.fib(40)
	fmt.Printf("fib(%d)=%.0f,%.0f\n", 40, a, b)
}
