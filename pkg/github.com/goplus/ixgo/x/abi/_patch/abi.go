package abi

import (
	"github.com/goplus/ixgo/x/abi"
)

var alwaysFalse bool
var escapeSink any

// Escape forces any pointers in x to escape to the heap.
func Escape[T any](x T) T {
	if alwaysFalse {
		escapeSink = x
	}
	return x
}

// TypeFor returns the abi.Type for a type parameter.
func TypeFor[T any]() *abi.Type {
	var v T
	if t := abi.TypeOf(v); t != nil {
		return t // optimize for T being a non-interface kind
	}
	return abi.TypeOf((*T)(nil)).Elem() // only for an interface kind
}
