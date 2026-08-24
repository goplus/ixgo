//go:build go1.18

package main

import "github.com/goplus/ixgo/x/abi"

func main() {
	if abi.TypeFor[int]() == nil || abi.Escape(1) != 1 {
		panic("abi patch")
	}
}
