//go:build !wasm

// run
package main

import "fmt"

type fibonacci struct {
	previous float32
	current  float32
}

func (p *fibonacci) fib(x int) float32 {
	if x < 2 {
		p.previous, p.current = 0, float32(x)
		return float32(x)
	}
	p.fib(x - 1)
	p.previous, p.current = p.current, p.previous+p.current
	return p.current
}

func main() {
	p := &fibonacci{}
	fmt.Printf("fib(%d)=%.0f\n", 40, p.fib(40))
}
