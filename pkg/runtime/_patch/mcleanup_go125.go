// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.25
// +build go1.25

package runtime

import (
	"runtime"
	"unsafe"

	"github.com/goplus/ixgo/x/abi"
)

// AddCleanup attaches a cleanup function to ptr. Some time after ptr is no longer
// reachable, the runtime will call cleanup(arg) in a separate goroutine.
//
// A typical use is that ptr is an object wrapping an underlying resource (e.g.,
// a File object wrapping an OS file descriptor), arg is the underlying resource
// (e.g., the OS file descriptor), and the cleanup function releases the underlying
// resource (e.g., by calling the close system call).
//
// There are few constraints on ptr. In particular, multiple cleanups may be
// attached to the same pointer, or to different pointers within the same
// allocation.
//
// If ptr is reachable from cleanup or arg, ptr will never be collected
// and the cleanup will never run. As a protection against simple cases of this,
// AddCleanup panics if arg is equal to ptr.
//
// There is no specified order in which cleanups will run.
// In particular, if several objects point to each other and all become
// unreachable at the same time, their cleanups all become eligible to run
// and can run in any order. This is true even if the objects form a cycle.
//
// Cleanups run concurrently with any user-created goroutines.
// Cleanups may also run concurrently with one another (unlike finalizers).
// If a cleanup function must run for a long time, it should create a new goroutine
// to avoid blocking the execution of other cleanups.
//
// If ptr has both a cleanup and a finalizer, the cleanup will only run once
// it has been finalized and becomes unreachable without an associated finalizer.
//
// The cleanup(arg) call is not always guaranteed to run; in particular it is not
// guaranteed to run before program exit.
//
// Cleanups are not guaranteed to run if the size of T is zero bytes, because
// it may share same address with other zero-size objects in memory. See
// https://go.dev/ref/spec#Size_and_alignment_guarantees.
//
// It is not guaranteed that a cleanup will run for objects allocated
// in initializers for package-level variables. Such objects may be
// linker-allocated, not heap-allocated.
//
// Note that because cleanups may execute arbitrarily far into the future
// after an object is no longer referenced, the runtime is allowed to perform
// a space-saving optimization that batches objects together in a single
// allocation slot. The cleanup for an unreferenced object in such an
// allocation may never run if it always exists in the same batch as a
// referenced object. Typically, this batching only happens for tiny
// (on the order of 16 bytes or less) and pointer-free objects.
//
// A cleanup may run as soon as an object becomes unreachable.
// In order to use cleanups correctly, the program must ensure that
// the object is reachable until it is safe to run its cleanup.
// Objects stored in global variables, or that can be found by tracing
// pointers from a global variable, are reachable. A function argument or
// receiver may become unreachable at the last point where the function
// mentions it. To ensure a cleanup does not get called prematurely,
// pass the object to the [KeepAlive] function after the last point
// where the object must remain reachable.
func AddCleanup[T, S any](ptr *T, cleanup func(S), arg S) (r runtime.Cleanup) {
	// Explicitly force ptr to escape to the heap.
	ptr = abi.Escape(ptr)

	// The pointer to the object must be valid.
	if ptr == nil {
		panic("runtime.AddCleanup: ptr is nil")
	}
	usptr := uintptr(unsafe.Pointer(ptr))

	// Check that arg is not equal to ptr.
	if kind := abi.TypeOf(arg).Kind(); kind == abi.Pointer || kind == abi.UnsafePointer {
		if unsafe.Pointer(ptr) == *((*unsafe.Pointer)(unsafe.Pointer(&arg))) {
			panic("runtime.AddCleanup: ptr is equal to arg, cleanup will never run")
		}
	}
	if inUserArenaChunk(usptr) {
		// Arena-allocated objects are not eligible for cleanup.
		panic("runtime.AddCleanup: ptr is arena-allocated")
	}
	// if debug.sbrk != 0 {
	// 	// debug.sbrk never frees memory, so no cleanup will ever run
	// 	// (and we don't have the data structures to record them).
	// 	// Return a noop cleanup.
	// 	return Cleanup{}
	// }

	fn := func() {
		cleanup(arg)
	}
	// Closure must escape.
	fv := *(**funcval)(unsafe.Pointer(&fn))
	fv = abi.Escape(fv)

	// Find the containing object.
	base, _, _ := findObject(usptr, 0, 0)
	if base == 0 {
		if isGoPointerWithoutSpan(unsafe.Pointer(ptr)) {
			// Cleanup is a noop.
			return runtime.Cleanup{}
		}
		panic("runtime.AddCleanup: ptr not in allocated block")
	}

	// Create another G if necessary.
	// if gcCleanups.needG() {
	// 	gcCleanups.createGs()
	// }

	createGIfNecessary()

	id := addCleanup(unsafe.Pointer(ptr), fv)
	// if debug.checkfinalizers != 0 {
	// 	cleanupFn := *(**funcval)(unsafe.Pointer(&cleanup))
	// 	setCleanupContext(unsafe.Pointer(ptr), abi.TypeFor[T](), sys.GetCallerPC(), cleanupFn.fn, id)
	// }
	return *(*runtime.Cleanup)(unsafe.Pointer(&Cleanup{
		id:  id,
		ptr: usptr,
	}))
}

type funcval struct {
	fn uintptr
	// variable-size, fn-specific data here
}

// Cleanup is a handle to a cleanup call for a specific object.
type Cleanup struct {
	// id is the unique identifier for the cleanup within the arena.
	id uint64
	// ptr contains the pointer to the object.
	ptr uintptr
}

// inUserArenaChunk returns true if p points to a user arena chunk.
//
//go:linkname inUserArenaChunk runtime.inUserArenaChunk
func inUserArenaChunk(p uintptr) bool

// findObject returns the base address for the heap object containing
// the address p, the object's span, and the index of the object in s.
// If p does not point into a heap object, it returns base == 0.
//
// If p points is an invalid heap pointer and debug.invalidptr != 0,
// findObject panics.
//
// refBase and refOff optionally give the base address of the object
// in which the pointer p was found and the byte offset at which it
// was found. These are used for error reporting.
//
// It is nosplit so it is safe for p to be a pointer to the current goroutine's stack.
// Since p is a uintptr, it would not be adjusted if the stack were to move.
//
// findObject should be an internal detail,
// but widely used packages access it using linkname.
// Notable members of the hall of shame include:
//   - github.com/bytedance/sonic
//
// Do not remove or change the type signature.
// See go.dev/issue/67401.
//
//go:linkname findObject runtime.findObject
//go:nosplit
func findObject(p, refBase, refOff uintptr) (base uintptr, s unsafe.Pointer, objIndex uintptr)

//go:linkname createfing runtime.createfing
func createfing()

// addCleanup attaches a cleanup function to the object. Multiple
// cleanups are allowed on an object, and even the same pointer.
// A cleanup id is returned which can be used to uniquely identify
// the cleanup.
//
//go:linkname addCleanup runtime.addCleanup
func addCleanup(p unsafe.Pointer, f *funcval) uint64

//go:linkname isGoPointerWithoutSpan runtime.isGoPointerWithoutSpan
func isGoPointerWithoutSpan(p unsafe.Pointer) bool

//go:linkname createGIfNecessary runtime.createGIfNecessary
func createGIfNecessary()
