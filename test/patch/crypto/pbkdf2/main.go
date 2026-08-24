//go:build go1.24

package main

import (
	"crypto/pbkdf2"
	"crypto/sha256"
)

func main() {
	k, e := pbkdf2.Key(sha256.New, "password", []byte("salt"), 2, 32)
	if e != nil || len(k) != 32 || k[0] == 0 {
		panic("pbkdf2.Key")
	}
}
