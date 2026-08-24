package rand

import (
	q "math/rand/v2"
	"math/bits"
)

// Rand is the Go 1.27 random number generator.
// It is defined in the patch so that its generic N method is type-checkable.
type Rand struct { src q.Source }

func New(src q.Source) *Rand { return &Rand{src: src} }
func (r *Rand) Int64() int64 { return int64(r.src.Uint64() &^ (1 << 63)) }
func (r *Rand) Uint32() uint32 { return uint32(r.src.Uint64() >> 32) }
func (r *Rand) Uint64() uint64 { return r.src.Uint64() }
func (r *Rand) Int32() int32 { return int32(r.src.Uint64() >> 33) }
func (r *Rand) Int() int { return int(uint(r.src.Uint64()) << 1 >> 1) }
func (r *Rand) Uint() uint { return uint(r.src.Uint64()) }

func (r *Rand) Int64N(n int64) int64 { if n <= 0 { panic("invalid argument to Int64N") }; return int64(r.uint64n(uint64(n))) }
func (r *Rand) Uint64N(n uint64) uint64 { if n == 0 { panic("invalid argument to Uint64N") }; return r.uint64n(n) }
func (r *Rand) Int32N(n int32) int32 { if n <= 0 { panic("invalid argument to Int32N") }; return int32(r.uint64n(uint64(n))) }
func (r *Rand) Uint32N(n uint32) uint32 { if n == 0 { panic("invalid argument to Uint32N") }; return uint32(r.uint64n(uint64(n))) }
func (r *Rand) IntN(n int) int { if n <= 0 { panic("invalid argument to IntN") }; return int(r.uint64n(uint64(n))) }
func (r *Rand) UintN(n uint) uint { if n == 0 { panic("invalid argument to UintN") }; return uint(r.uint64n(uint64(n))) }

type intType127 interface { ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr }

func (r *Rand) N[Int intType127](n Int) Int { if n <= 0 { panic("invalid argument to N") }; return Int(r.uint64n(uint64(n))) }

func (r *Rand) uint64n(n uint64) uint64 {
	if n&(n-1) == 0 { return r.Uint64() & (n - 1) }
	hi, lo := bits.Mul64(r.Uint64(), n)
	if lo < n { for threshold := -n % n; lo < threshold; { hi, lo = bits.Mul64(r.Uint64(), n) } }
	return hi
}

func (r *Rand) Float64() float64 { return float64(r.Uint64()<<11>>11) / (1 << 53) }
func (r *Rand) Float32() float32 { return float32(r.Uint32()>>8) / (1 << 24) }
func (r *Rand) Perm(n int) []int { p := make([]int, n); for i := range p { p[i] = i }; r.Shuffle(n, func(i, j int) { p[i], p[j] = p[j], p[i] }); return p }
func (r *Rand) Shuffle(n int, swap func(i, j int)) { if n < 0 { panic("invalid argument to Shuffle") }; for i := n - 1; i > 0; i-- { swap(i, int(r.uint64n(uint64(i+1)))) } }
