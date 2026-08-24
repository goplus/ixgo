//go:build go1.23

package main

import "unique"

func main() {
	a, b := unique.Make("x"), unique.Make("x")
	if a.Value() != "x" || a != b {
		panic("Make")
	}
}
