//go:build go1.27

package main

import "math/rand/v2"

type source struct{ n uint64 }

func (s *source) Uint64() uint64 { s.n++; return s.n }

func main() {
	if n := rand.N(10); n < 0 || n >= 10 {
		panic("rand.N")
	}
	r := rand.New(&source{})
	if n := r.N(10); n < 0 || n >= 10 {
		panic("N")
	}
	var n8 int8 = r.N(int8(8))
	if n8 < 0 || n8 >= 8 {
		panic("N[int8]")
	}
	var nu uint32 = r.N(uint32(32))
	if nu >= 32 {
		panic("N[uint32]")
	}
	panicked := false
	func() {
		defer func() { panicked = recover() != nil }()
		r.N(0)
	}()
	if !panicked {
		panic("N zero did not panic")
	}
}
