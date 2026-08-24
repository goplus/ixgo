//go:build go1.27

package main

import "hash/maphash"

func main() {
	s := maphash.MakeSeed()
	if maphash.Comparable(s, 1) != maphash.Comparable(s, 1) {
		panic("Comparable")
	}
	var h maphash.Hash
	h.SetSeed(s)
	maphash.WriteComparable(&h, "x")
	var x maphash.ComparableHasher[string]
	if !x.Equal("x", "x") || x.Equal("x", "y") {
		panic("ComparableHasher")
	}
}
