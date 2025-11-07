package stringslite

import "unsafe"

func Clone(s string) string {
	if len(s) == 0 {
		return ""
	}
	b := make([]byte, len(s))
	copy(b, s)
	return string(unsafe.Slice(&b[0], len(b)))
}
