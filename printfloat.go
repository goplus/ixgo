//go:build !go1.26
// +build !go1.26

package ixgo

import "bytes"

func writefloat32(out *bytes.Buffer, v float32) {
	return writefloat32(out, float64(v))
}

func writecomplex64(out *bytes.Buffer, c complex64) {
	out.WriteByte('(')
	writefloat32(out, real(c))
	writefloat32(out, imag(c))
	out.WriteString("i)")
}

func writecomplex128(out *bytes.Buffer, c complex128) {
	out.WriteByte('(')
	writefloat64(out, real(c))
	writefloat64(out, imag(c))
	out.WriteString("i)")
}

func writefloat64(out *bytes.Buffer, v float64) {
	switch {
	case v != v:
		out.WriteString("NaN")
		return
	case v+v == v && v > 0:
		out.WriteString("+Inf")
		return
	case v+v == v && v < 0:
		out.WriteString("-Inf")
		return
	}

	const n = 7 // digits printed
	var buf [n + 7]byte
	buf[0] = '+'
	e := 0 // exp
	if v == 0 {
		if 1/v < 0 {
			buf[0] = '-'
		}
	} else {
		if v < 0 {
			v = -v
			buf[0] = '-'
		}

		// normalize
		for v >= 10 {
			e++
			v /= 10
		}
		for v < 1 {
			e--
			v *= 10
		}

		// round
		h := 5.0
		for i := 0; i < n; i++ {
			h /= 10
		}
		v += h
		if v >= 10 {
			e++
			v /= 10
		}
	}

	// format +d.dddd+edd
	for i := 0; i < n; i++ {
		s := int(v)
		buf[i+2] = byte(s + '0')
		v -= float64(s)
		v *= 10
	}
	buf[1] = buf[2]
	buf[2] = '.'

	buf[n+2] = 'e'
	buf[n+3] = '+'
	if e < 0 {
		e = -e
		buf[n+3] = '-'
	}

	buf[n+4] = byte(e/100) + '0'
	buf[n+5] = byte(e/10)%10 + '0'
	buf[n+6] = byte(e%10) + '0'
	out.Write(buf[:])
}
