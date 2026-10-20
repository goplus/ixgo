// run

//go:build !wasm

package main

import "fmt"

type fibonacci struct {
	result int
}

// Each method keeps uglyfib's mutually recursive shape. The accumulator that
// was previously *int is stored in the receiver field.
func (p *fibonacci) f(x int, xm1, xm2 *int) {
	y := [2]int{x - 4, 0}
	if x < 2 {
		p.result += x
		return
	}
	x -= 3
	p.g(*xm1, xm2, &x)
	p.h(*xm2, &x, &y[0])
}

func (p *fibonacci) g(x int, xm1, xm2 *int) {
	y := [3]int{x - 4, 0, 0}
	if x < 2 {
		p.result += x
		return
	}
	x -= 3
	p.k(*xm2, &x, &y[0])
	p.h(*xm1, xm2, &x)
}

func (p *fibonacci) h(x int, xm1, xm2 *int) {
	y := [4]int{x - 4, 0, 0, 0}
	if x < 2 {
		p.result += x
		return
	}
	x -= 3
	p.k(*xm1, xm2, &x)
	p.f(*xm2, &x, &y[0])
}

func (p *fibonacci) k(x int, xm1, xm2 *int) {
	y := [5]int{x - 4, 0, 0, 0, 0}
	if x < 2 {
		p.result += x
		return
	}
	x -= 3
	p.f(*xm2, &x, &y[0])
	p.g(*xm1, xm2, &x)
}

func main() {
	x := 40
	xm1, xm2 := x-1, x-2
	p := &fibonacci{}
	p.f(x, &xm1, &xm2)
	fmt.Printf("Fib(%d)=%d\n", x, p.result)
}
