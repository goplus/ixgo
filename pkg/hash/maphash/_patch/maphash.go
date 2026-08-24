package maphash

import (
	"hash/maphash"
	"unsafe"

	"github.com/goplus/ixgo/x/abi"
	"github.com/goplus/ixgo/x/goarch"
)

func Comparable[T comparable](seed maphash.Seed, v T) uint64 {
	return comparableHash(v, seed)
}

type _Seed struct {
	s uint64
}

const use64BitHash = goarch.PtrSize == 8 && goarch.IsWasm == 0

func comparableHash[T comparable](v T, seed maphash.Seed) uint64 {
	s := (*_Seed)(unsafe.Pointer(&seed)).s
	var m map[T]struct{}
	mTyp := abi.TypeOf(m)
	hasher := (*abi.MapType)(unsafe.Pointer(mTyp)).Hasher
	if use64BitHash {
		return uint64(hasher(abi.NoEscape(unsafe.Pointer(&v)), uintptr(s)))
	}
	lo := hasher(abi.NoEscape(unsafe.Pointer(&v)), uintptr(uint32(s)))
	hi := hasher(abi.NoEscape(unsafe.Pointer(&v)), uintptr(s>>32))
	return uint64(hi)<<32 | uint64(lo)
}

// WriteComparable adds x to the data hashed by h.
func WriteComparable[T comparable](h *maphash.Hash, x T) {
	//abi.EscapeNonString(x)
	_h := (*_Hash)(unsafe.Pointer(h))
	// writeComparable directly operates on h.state
	// without using h.buf. Mix in the buffer length so it won't
	// commute with a buffered write, which either changes h.n or changes
	// h.state.
	if _h.n != 0 {
		writeComparable(h, _h.n)
	}
	writeComparable(h, x)
}

type _Hash struct {
	_     [0]func()     // not comparable
	seed  _Seed         // initial seed used for this hash
	state _Seed         // current hash of all flushed bytes
	buf   [bufSize]byte // unflushed byte buffer
	n     int           // number of unflushed bytes
}

// bufSize is the size of the Hash write buffer.
// The buffer ensures that writes depend only on the sequence of bytes,
// not the sequence of WriteByte/Write/WriteString calls,
// by always calling rthash with a full buffer (except for the tail).
const bufSize = 128

func writeComparable[T comparable](h *maphash.Hash, v T) {
	_h := (*_Hash)(unsafe.Pointer(h))
	_h.state.s = comparableHash(v, *(*maphash.Seed)(unsafe.Pointer(&_h.state)))
}
