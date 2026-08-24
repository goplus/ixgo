//go:build go1.23

package main

import "iter"

func main() {
	next, stop := iter.Pull(iter.Seq[int](func(y func(int) bool) { y(1) }))
	defer stop()
	if v, ok := next(); !ok || v != 1 {
		panic("Pull")
	}
	next2, stop2 := iter.Pull2(iter.Seq2[string, int](func(y func(string, int) bool) { y("x", 1) }))
	defer stop2()
	if k, v, ok := next2(); !ok || k != "x" || v != 1 {
		panic("Pull2")
	}
}
