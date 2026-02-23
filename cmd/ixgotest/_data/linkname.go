package main

import (
	"fmt"
	"iter"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

func main() {
	cleanup()
	iterPull()
	iterPull2()
}

func cleanup() {
	type Blob struct {
		data []byte
	}

	var since time.Duration
	var done atomic.Bool
	b := &Blob{data: make([]byte, 1000)}
	now := time.Now()
	runtime.AddCleanup(b, func(created time.Time) {
		since = time.Since(created) / time.Millisecond
		done.Store(true)
	}, now)

	b = nil
	time.Sleep(10 * time.Millisecond)

	deadline := time.Now().Add(2 * time.Second)
	for !done.Load() && time.Now().Before(deadline) {
		runtime.GC()
		time.Sleep(5 * time.Millisecond)
	}
	if !done.Load() {
		panic("cleanup timeout")
	}
	println(since)
}

func Reversed[V any](s []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				return
			}
		}
	}
}

func iterPull() {
	s := []int{1, 2, 3, 4, 5}

	next, stop := iter.Pull(Reversed(s))
	defer stop()

	var r string
	for {
		v, ok := next()
		if !ok {
			break
		}
		r += strconv.Itoa(v)
	}
	if r != "54321" {
		panic("error: " + r)
	}
}

func rangeMap(m map[string]int) iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func iterPull2() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	next, stop := iter.Pull2(rangeMap(m))
	defer stop()

	for {
		k, v, ok := next()
		if !ok {
			break
		}
		fmt.Printf("%s: %d\n", k, v)
	}
}
