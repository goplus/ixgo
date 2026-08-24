package hkdf

import (
	"errors"
	"hash"
)

// Extract generates a pseudorandom key from secret and salt.
func Extract[H hash.Hash](h func() H, secret, salt []byte) ([]byte, error) {
	if salt == nil {
		salt = make([]byte, h().Size())
	}
	return hmacSum(h, salt, secret), nil
}

// Expand derives a key from a pseudorandom key and context information.
func Expand[H hash.Hash](h func() H, pseudorandomKey []byte, info string, keyLength int) ([]byte, error) {
	limit := h().Size() * 255
	if keyLength > limit {
		return nil, errors.New("hkdf: requested key length too large")
	}

	out := make([]byte, 0, keyLength)
	var counter byte
	var previous []byte
	for len(out) < keyLength {
		counter++
		if counter == 0 {
			panic("hkdf: counter overflow")
		}
		previous = hmacSum(h, pseudorandomKey, previous, []byte(info), []byte{counter})
		remain := keyLength - len(out)
		if remain > len(previous) {
			remain = len(previous)
		}
		out = append(out, previous[:remain]...)
	}
	return out, nil
}

func hmacSum[H hash.Hash](h func() H, key []byte, data ...[]byte) []byte {
	ch := h()
	blockSize := ch.BlockSize()
	if len(key) > blockSize {
		ch.Write(key)
		key = ch.Sum(nil)
	}
	padded := make([]byte, blockSize)
	copy(padded, key)
	ipad := make([]byte, blockSize)
	opad := make([]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		ipad[i] = padded[i] ^ 0x36
		opad[i] = padded[i] ^ 0x5c
	}
	inner := h()
	inner.Write(ipad)
	for _, part := range data {
		inner.Write(part)
	}
	digest := inner.Sum(nil)
	outer := h()
	outer.Write(opad)
	outer.Write(digest)
	return outer.Sum(nil)
}

// Key derives a key from secret, salt, and context information.
func Key[H hash.Hash](h func() H, secret, salt []byte, info string, keyLength int) ([]byte, error) {
	prk, err := Extract(h, secret, salt)
	if err != nil {
		return nil, err
	}
	return Expand(h, prk, info, keyLength)
}
