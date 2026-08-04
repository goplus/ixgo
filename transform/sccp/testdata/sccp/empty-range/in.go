package test

var index int
var ints []*int
var next func() int

func target() int {
	for index, *(ints[next()]) = range []int{} {
	}
	return index
}
