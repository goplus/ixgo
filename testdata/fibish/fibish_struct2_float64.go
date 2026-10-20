//go:build !wasm

// run
package main

import "fmt"

type fibonacci struct{ offset float64 }

func (p *fibonacci) fib(x int) (float64, float64) {
	if x < 3 {
		return p.offset, float64(x) + p.offset
	}
	a, b := p.fib(x - 2)
	c, d := p.fib(x - 1)
	return a + d, b + c
}

func main() {
	p := &fibonacci{offset: 0}
	a, b := p.fib(40)
	fmt.Printf("fib(%d)=%.0f,%.0f\n", 40, a, b)
}
