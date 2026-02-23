//go:build go1.26
// +build go1.26

package ixgo

import (
	"bytes"
	"strconv"
)

func writefloat64(out *bytes.Buffer, v float64) {
	var buf [20]byte
	out.Write(strconv.AppendFloat(buf[:0], v, 'g', -1, 64))
}

func writefloat32(out *bytes.Buffer, v float32) {
	var buf [20]byte
	out.Write(strconv.AppendFloat(buf[:0], float64(v), 'g', -1, 32))
}

func writecomplex128(out *bytes.Buffer, c complex128) {
	var buf [44]byte
	out.Write(AppendComplex(buf[:0], c, 'g', -1, 128))
}

func writecomplex64(out *bytes.Buffer, c complex64) {
	var buf [44]byte
	out.Write(AppendComplex(buf[:0], complex128(c), 'g', -1, 64))
}

/*
func writeuint(out *bytes.Buffer, v uint64) {
	// Note: Avoiding strconv.AppendUint so that it's clearer
	// that there are no allocations in this routine.
	// cmd/link/internal/ld.TestAbstractOriginSanity
	// sees the append and doesn't realize it doesn't allocate.
	var buf [20]byte
	i := strconv.RuntimeFormatBase10(buf[:], v)
	out.Write(buf[i:])
}

func writeint(out *bytes.Buffer, v int64) {
	// Note: Avoiding strconv.AppendUint so that it's clearer
	// that there are no allocations in this routine.
	// cmd/link/internal/ld.TestAbstractOriginSanity
	// sees the append and doesn't realize it doesn't allocate.
	neg := v < 0
	u := uint64(v)
	if neg {
		u = -u
	}
	var buf [20]byte
	i := strconv.RuntimeFormatBase10(buf[:], u)
	if neg {
		i--
		buf[i] = '-'
	}
	out.Write(buf[i:])
}
*/

// AppendComplex appends the result of FormatComplex to dst.
// It is here for the runtime.
// There is no public strconv.AppendComplex.
func AppendComplex(dst []byte, c complex128, fmt byte, prec, bitSize int) []byte {
	if bitSize != 64 && bitSize != 128 {
		panic("invalid bitSize")
	}
	bitSize >>= 1 // complex64 uses float32 internally

	dst = append(dst, '(')
	dst = strconv.AppendFloat(dst, real(c), fmt, prec, bitSize)
	i := len(dst)
	dst = strconv.AppendFloat(dst, imag(c), fmt, prec, bitSize)
	// Check if imaginary part has a sign. If not, add one.
	if dst[i] != '+' && dst[i] != '-' {
		dst = append(dst, 0)
		copy(dst[i+1:], dst[i:])
		dst[i] = '+'
	}
	dst = append(dst, "i)"...)
	return dst
}
