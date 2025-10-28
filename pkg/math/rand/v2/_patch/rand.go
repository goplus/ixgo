package rand

import (
	_ "math/rand/v2"
	_ "unsafe"
)

func N[Int intType](n Int) Int {
	if n <= 0 {
		panic("invalid argument to N")
	}
	return Int(globalRandUint64n(uint64(n)))
}

type intType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

//go:linkname globalRandUint64n math/rand/v2.globalRandUint64n
func globalRandUint64n(uint64) uint64
