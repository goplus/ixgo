//go:build !wasm

// run
package main

import "fmt"

func f(x float32, xm1, xm2, p *float32) {
	y := [2]float32{x - 4, 0}
	if x < 2 {
		*p += x
		return
	}
	x -= 3
	g(*xm1, xm2, &x, p)
	h(*xm2, &x, &y[0], p)
}

func g(x float32, xm1, xm2, p *float32) {
	y := [3]float32{x - 4, 0, 0}
	if x < 2 {
		*p += x
		return
	}
	x -= 3
	k(*xm2, &x, &y[0], p)
	h(*xm1, xm2, &x, p)
}

func h(x float32, xm1, xm2, p *float32) {
	y := [4]float32{x - 4, 0, 0, 0}
	if x < 2 {
		*p += x
		return
	}
	x -= 3
	k(*xm1, xm2, &x, p)
	f(*xm2, &x, &y[0], p)
}

func k(x float32, xm1, xm2, p *float32) {
	y := [5]float32{x - 4, 0, 0, 0, 0}
	if x < 2 {
		*p += x
		return
	}
	x -= 3
	f(*xm2, &x, &y[0], p)
	g(*xm1, xm2, &x, p)
}

func main() {
	x := float32(40)
	xm1, xm2 := x-1, x-2
	var result float32
	f(x, &xm1, &xm2, &result)
	fmt.Printf("Fib(%.0f)=%.0f\n", x, result)
}
