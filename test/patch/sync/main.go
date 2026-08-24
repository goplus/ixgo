//go:build go1.21

package main

import "sync"

func main() {
	n := 0
	f := sync.OnceFunc(func() { n++ })
	f()
	f()
	if n != 1 {
		panic("OnceFunc")
	}
	if sync.OnceValue(func() int { return 1 })() != 1 {
		panic("OnceValue")
	}
	a, b := sync.OnceValues(func() (int, string) { return 1, "x" })()
	if a != 1 || b != "x" {
		panic("OnceValues")
	}
}
