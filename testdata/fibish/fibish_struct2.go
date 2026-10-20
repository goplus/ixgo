// run

//go:build !wasm

package main

import "fmt"

type fibonacci struct {
	offset int
}

func (p *fibonacci) fib(x int) (int, int) {
	if x < 3 {
		return p.offset, x + p.offset
	}
	a, b := p.fib(x - 2)
	c, d := p.fib(x - 1)
	return a + d, b + c
}

func main() {
	p := &fibonacci{offset: 0}
	a, b := p.fib(40)
	fmt.Printf("fib(%d)=%d,%d\n", 40, a, b)
}
