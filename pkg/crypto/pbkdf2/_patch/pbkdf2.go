package pbkdf2

import (
	"errors"
	"hash"
)

// Key derives a key from password, salt, and iteration count using PBKDF2.
func Key[Hash hash.Hash](h func() Hash, password string, salt []byte, iter, keyLength int) ([]byte, error) {
	if keyLength <= 0 {
		return nil, errors.New("pkbdf2: keyLength must be larger than 0")
	}

	hashLen := h().Size()
	numBlocks := int((int64(keyLength) + int64(hashLen) - 1) / int64(hashLen))
	const maxBlocks = int64(1<<32 - 1)
	if keyLength+hashLen < keyLength || int64(numBlocks) > maxBlocks {
		return nil, errors.New("pbkdf2: keyLength too long")
	}

	prfKey := []byte(password)
	dk := make([]byte, 0, numBlocks*hashLen)
	u := make([]byte, hashLen)
	var block [4]byte
	for i := 1; i <= numBlocks; i++ {
		block[0] = byte(i >> 24)
		block[1] = byte(i >> 16)
		block[2] = byte(i >> 8)
		block[3] = byte(i)
		t := hmacSum(h, prfKey, salt, block[:])
		dk = append(dk, t...)
		copy(u, t)
		for n := 2; n <= iter; n++ {
			u = hmacSum(h, prfKey, u)
			for j := range t {
				t[j] ^= u[j]
			}
		}
		copy(dk[len(dk)-hashLen:], t)
	}
	return dk[:keyLength], nil
}

func hmacSum[Hash hash.Hash](h func() Hash, key []byte, data ...[]byte) []byte {
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
