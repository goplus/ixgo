//go:build go1.24

package main

import "weak"

func main() {
	x := 1
	p := weak.Make(&x)
	if p.Value() != &x {
		panic("Make")
	}
}
