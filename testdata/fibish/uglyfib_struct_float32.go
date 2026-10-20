//go:build !wasm

// run
package main

import "fmt"

type fibonacci struct{ result float32 }

func (p *fibonacci) f(x float32, xm1, xm2 *float32) {
	y := [2]float32{x - 4, 0}
	if x < 2 {
		p.result += x
		return
	}
	x -= 3
	p.g(*xm1, xm2, &x)
	p.h(*xm2, &x, &y[0])
}
func (p *fibonacci) g(x float32, xm1, xm2 *float32) {
	y := [3]float32{x - 4, 0, 0}
	if x < 2 {
		p.result += x
		return
	}
	x -= 3
	p.k(*xm2, &x, &y[0])
	p.h(*xm1, xm2, &x)
}
func (p *fibonacci) h(x float32, xm1, xm2 *float32) {
	y := [4]float32{x - 4, 0, 0, 0}
	if x < 2 {
		p.result += x
		return
	}
	x -= 3
	p.k(*xm1, xm2, &x)
	p.f(*xm2, &x, &y[0])
}
func (p *fibonacci) k(x float32, xm1, xm2 *float32) {
	y := [5]float32{x - 4, 0, 0, 0, 0}
	if x < 2 {
		p.result += x
		return
	}
	x -= 3
	p.f(*xm2, &x, &y[0])
	p.g(*xm1, xm2, &x)
}
func main() {
	x := float32(40)
	a, b := x-1, x-2
	p := &fibonacci{}
	p.f(x, &a, &b)
	fmt.Printf("Fib(%.0f)=%.0f\n", x, p.result)
}
