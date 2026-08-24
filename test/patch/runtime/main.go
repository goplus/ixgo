//go:build go1.24

package main

import "runtime"

func main() {
	x := 1
	h := runtime.AddCleanup(&x, func(string) {}, "cleanup")
	h.Stop()
}
