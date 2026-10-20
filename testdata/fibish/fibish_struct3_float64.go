//go:build !wasm

// run
package main

import "fmt"

type fibonacci struct {
	previous float64
	current  float64
}

func (p *fibonacci) fib(x int) float64 {
	if x < 2 {
		p.previous, p.current = 0, float64(x)
		return float64(x)
	}
	p.fib(x - 1)
	p.previous, p.current = p.current, p.previous+p.current
	return p.current
}

func main() {
	p := &fibonacci{}
	fmt.Printf("fib(%d)=%.0f\n", 40, p.fib(40))
}
