// run

//go:build !wasm

package main

import "fmt"

type fibonacci struct {
	previous int
	current  int
}

// fib records the previous pair in fields before descending recursively.
func (p *fibonacci) fib(x int) int {
	if x < 2 {
		p.previous, p.current = 0, x
		return x
	}

	p.fib(x - 1)
	p.previous, p.current = p.current, p.previous+p.current
	return p.current
}

func main() {
	p := &fibonacci{}
	fmt.Printf("fib(%d)=%d\n", 40, p.fib(40))
}
