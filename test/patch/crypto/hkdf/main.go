//go:build go1.24

package main

import (
	"bytes"
	"crypto/hkdf"
	"crypto/sha256"
)

func main() {
	p, e := hkdf.Extract(sha256.New, []byte("secret"), nil)
	if e != nil {
		panic(e)
	}
	x, e := hkdf.Expand(sha256.New, p, "info", 32)
	if e != nil || len(x) != 32 {
		panic("hkdf.Expand")
	}
	y, e := hkdf.Key(sha256.New, []byte("secret"), nil, "info", 32)
	if e != nil || !bytes.Equal(x, y) {
		panic("hkdf.Key")
	}
}
