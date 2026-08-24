//go:build go1.22

package main

import "math/rand/v2"

func main() {
	if n := rand.N(10); n < 0 || n >= 10 {
		panic("N")
	}
}
