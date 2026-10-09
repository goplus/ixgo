package main

// Tests of defer.  (Deferred recover() belongs is recover.go.)

import "fmt"

func deferMutatesResults(noArgReturn bool) (a, b int) {
	defer func() {
		if a != 1 || b != 2 {
			panic(fmt.Sprint(a, b))
		}
		a, b = 3, 4
	}()
	if noArgReturn {
		a, b = 1, 2
		return
	}
	return 1, 2
}

func init() {
	a, b := deferMutatesResults(true)
	if a != 3 || b != 4 {
		panic(fmt.Sprint(a, b))
	}
	a, b = deferMutatesResults(false)
	if a != 3 || b != 4 {
		panic(fmt.Sprint(a, b))
	}
}

// We concatenate init blocks to make a single function, but we must
// run defers at the end of each block, not the combined function.
var deferCount = 0

func init() {
	deferCount = 1
	defer func() {
		deferCount++
	}()
	// defer runs HERE
}

func init() {
	// Strictly speaking the spec says deferCount may be 0 or 2
	// since the relative order of init blocks is unspecified.
	if deferCount != 2 {
		panic(deferCount) // defer call has not run!
	}
}

func yield2(yield func(int) bool) {
	_ = yield(1) && yield(2)
}

func yield3(yield func(int) bool) {
	_ = yield(1) && yield(2) && yield(3)
}

var rangeDeferCount int

func rangeFuncOnlyDefers() {
	n := 0
	for _ = range yield2 {
		for _ = range yield3 {
			n++
			defer func() { rangeDeferCount++ }()
		}
	}
}

func rangeFuncNamedReturn() (r int) {
	for range yield2 {
		defer func() { r++ }()
	}
	return
}

func main() {
	rangeFuncOnlyDefers()
	if rangeDeferCount != 6 {
		panic(rangeDeferCount)
	}
	if rangeFuncNamedReturn() != 2 {
		panic("named return")
	}
}
