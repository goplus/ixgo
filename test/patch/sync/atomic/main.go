//go:build go1.19

package main

import "sync/atomic"

func main() {
	var p atomic.Pointer[int]
	x, y := 1, 2
	p.Store(&x)
	if p.Load() != &x || p.Swap(&y) != &x || !p.CompareAndSwap(&y, &x) {
		panic("Pointer")
	}
}
