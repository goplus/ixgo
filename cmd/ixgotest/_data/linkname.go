package main

import (
	"iter"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

func main() {
	cleanup()
	iterPull()
}

func cleanup() {
	type Blob struct {
		data []byte
	}

	b := &Blob{data: make([]byte, 1000)}
	now := time.Now()
	var since time.Duration
	var done atomic.Bool
	runtime.AddCleanup(b, func(created time.Time) {
		since = time.Since(created) / time.Millisecond
		done.Store(true)
	}, now)

	time.Sleep(10 * time.Millisecond)
	b = nil

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
