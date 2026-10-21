/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ixgo

import (
	"go/token"
	"sync"
	"sync/atomic"
)

// ExecutionPosition reports a goroutine's source location.
type ExecutionPosition struct {
	GoroutineID uint64
	Position    token.Position
	// Clear removes the location. Position is zero for clear events.
	Clear bool
	// Sequence increases within a tracker. Keep the latest per goroutine,
	// including clears, to ignore delayed events.
	Sequence uint64
}

// ExecutionOptions controls source-position reporting.
type ExecutionOptions struct {
	// Filter runs during compilation; nil accepts all valid positions.
	// It must return consistent results and be safe for concurrent compilation.
	Filter func(token.Position) bool
	// ReportRepeated reports every accepted DebugRef position, even on the same
	// line. Caller restores and clears remain deduplicated.
	ReportRepeated bool
}

// ExecutionTracker tracks each goroutine's source position, merging consecutive
// positions on the same line unless ReportRepeated is set.
// Calls inherit their caller's position and restore it after user defers;
// the outermost call clears it on exit. Positions come from DebugRef metadata,
// not a callback before every statement. Share a tracker across interpreters
// to follow synchronous host callbacks between them.
//
// Create with NewExecutionTracker, do not copy, and Close when the session ends.
// Interpreter control methods, including RunMain and Abort, do not close it.
type ExecutionTracker struct {
	hook    func(ExecutionPosition)
	options ExecutionOptions

	mu      sync.Mutex
	active  map[uint64]*executionFrame
	visible map[uint64]uint64 // latest sequence until its clear callback returns
	seq     uint64
	closed  atomic.Bool
}

// NewExecutionTracker creates a source-line tracking session.
// Omitting options uses defaults; passing more than one panics.
// Callbacks run synchronously and may overlap, without tracker locks held.
// They must return normally and promptly; events may be retained or queued.
// Use the event's GoroutineID and Sequence to update host state.
// Tracking adds execution overhead.
func NewExecutionTracker(hook func(ExecutionPosition), options ...ExecutionOptions) *ExecutionTracker {
	if len(options) > 1 {
		panic("ixgo: NewExecutionTracker accepts at most one ExecutionOptions")
	}
	var opts ExecutionOptions
	if len(options) == 1 {
		opts = options[0]
	}
	return &ExecutionTracker{
		hook:    hook,
		options: opts,
		active:  make(map[uint64]*executionFrame),
		visible: make(map[uint64]uint64),
	}
}

// Close stops tracking and sends clears synchronously on the closing goroutine,
// including for pending notifications. Execution continues.
// Repeated calls do not wait. In-flight callbacks may still arrive;
// retain sequence watermarks after clears.
// To track again, create a new tracker and recompile.
func (t *ExecutionTracker) Close() {
	t.mu.Lock()
	if t.closed.Swap(true) {
		t.mu.Unlock()
		return
	}
	visible, seq := t.visible, t.seq
	t.seq += uint64(len(visible))
	// No further updates can reach these maps after closing.
	t.active, t.visible = nil, nil
	t.mu.Unlock()
	for goid := range visible {
		seq++
		t.notify(ExecutionPosition{GoroutineID: goid, Clear: true, Sequence: seq})
	}
}

func (t *ExecutionTracker) accepts(pos token.Position) bool {
	return pos.IsValid() && (t.options.Filter == nil || t.options.Filter(pos))
}

func (t *ExecutionTracker) enter(fr *frame) {
	if t.closed.Load() {
		return
	}
	f := fr.executionFrame()
	if f == nil {
		f = new(executionFrame)
		fr.stack[fr.pfn.executionReg()] = f
	}
	goid := GoroutineID()
	t.mu.Lock()
	if !t.closed.Load() {
		*f = executionFrame{parent: t.active[goid], gid: goid}
		if f.parent != nil {
			f.position = f.parent.position
		}
		t.active[goid] = f
	}
	t.mu.Unlock()
}

func (t *ExecutionTracker) setPosition(fr *frame, pos token.Position) {
	// The execution frame is owned by the goroutine running fr.
	f := fr.executionFrame()
	if f == nil || t.closed.Load() {
		return
	}
	old := f.position
	f.position = pos
	if !t.options.ReportRepeated && sameLine(old, pos) {
		return
	}
	t.mu.Lock()
	event := t.updateLocked(f.gid, pos)
	t.mu.Unlock()
	t.notify(event)
}

func (t *ExecutionTracker) leave(fr *frame) {
	// The execution frame is owned by the goroutine running fr.
	f := fr.executionFrame()
	if f == nil {
		return
	}
	parent := f.parent
	var pos token.Position
	if parent != nil {
		pos = parent.position
	}
	same := sameLine(f.position, pos)
	t.mu.Lock()
	var event ExecutionPosition
	if !t.closed.Load() {
		if parent != nil {
			t.active[f.gid] = parent
		} else {
			delete(t.active, f.gid)
		}
		if !same {
			event = t.updateLocked(f.gid, pos)
		}
	}
	t.mu.Unlock()
	// Release parent links before callbacks can reenter the interpreter.
	*f = executionFrame{}
	t.notify(event)
}

// updateLocked records an event under mu unless closed; callers deduplicate.
func (t *ExecutionTracker) updateLocked(goid uint64, pos token.Position) ExecutionPosition {
	if t.closed.Load() {
		return ExecutionPosition{}
	}
	t.seq++
	t.visible[goid] = t.seq
	return ExecutionPosition{GoroutineID: goid, Position: pos, Clear: !pos.IsValid(), Sequence: t.seq}
}

func (t *ExecutionTracker) notify(event ExecutionPosition) {
	if event.Sequence == 0 {
		return
	}
	if t.hook != nil {
		t.hook(event)
	}
	if event.Clear {
		t.mu.Lock()
		// Callbacks may reenter; only clear the sequence they completed.
		if t.visible[event.GoroutineID] == event.Sequence {
			delete(t.visible, event.GoroutineID)
		}
		t.mu.Unlock()
	}
}

// Host callbacks can lose frame.caller, so tracking keeps its own call chain.
// Each frame's state belongs to its executing goroutine.
type executionFrame struct {
	parent   *executionFrame
	gid      uint64
	position token.Position
}

// executionReg is the optional tracking slot after results, arguments and captures.
func (p *function) executionReg() register {
	return register(p.nres + p.narg + p.nenv)
}

func (fr *frame) executionFrame() *executionFrame {
	f, _ := fr.stack[fr.pfn.executionReg()].(*executionFrame)
	return f
}

func sameLine(a, b token.Position) bool {
	return a.Filename == b.Filename && a.Line == b.Line
}
