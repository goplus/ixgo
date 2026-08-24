package maphash

import "hash/maphash"

// Hasher defines hashing and equivalence operations for values of type T.
type Hasher[T any] interface {
	Hash(*maphash.Hash, T)
	Equal(x, y T) bool
}

// ComparableHasher implements Hasher using the built-in comparable relation.
type ComparableHasher[T comparable] struct {
	_ [0]func(T)
}

func (ComparableHasher[T]) Hash(h *maphash.Hash, v T) { WriteComparable(h, v) }
func (ComparableHasher[T]) Equal(x, y T) bool         { return x == y }
