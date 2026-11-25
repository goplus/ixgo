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
	"fmt"
	"go/token"
	"go/types"
	"maps"
	"path/filepath"
	"reflect"
	"runtime"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/tools/go/ssa"
)

func newExecutionRecorder(t *testing.T) *executionRecorder {
	t.Helper()
	recorder := &executionRecorder{
		active: make(map[uint64]token.Position),
		last:   make(map[uint64]uint64),
	}
	t.Cleanup(func() { recorder.check(t) })
	return recorder
}

func newExecutionTest(t *testing.T, mode Mode) (*Context, *ExecutionTracker, *executionRecorder) {
	t.Helper()
	ctx, recorder := NewContext(mode), newExecutionRecorder(t)
	tracker := NewExecutionTracker(recorder.record)
	ctx.SetExecutionTracker(tracker)
	return ctx, tracker, recorder
}

func loadExecution(t *testing.T, ctx *Context, source string) *Interp {
	t.Helper()
	interp, err := ctx.LoadInterp("execution.go", source)
	if err != nil {
		t.Fatal(err)
	}
	return interp
}

func runExecution(t *testing.T, interp *Interp, name string, want Value, args ...Value) {
	t.Helper()
	got, err := interp.RunFunc(name, args...)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("%s = (%v, %v), want (%v, nil)", name, got, err, want)
	}
}

const executionTestSource = `package main
func Root(n int) int { return n + 1 }
`

type executionRecorder struct {
	mu            sync.Mutex
	events        []ExecutionPosition
	active        map[uint64]token.Position
	last          map[uint64]uint64
	errors        []string
	allowRepeated bool
}

func (r *executionRecorder) record(event ExecutionPosition) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Keep the watermark after clear so late events cannot restore a position.
	if event.Sequence <= r.last[event.GoroutineID] {
		return
	}
	r.last[event.GoroutineID] = event.Sequence
	if event.GoroutineID == 0 {
		r.errors = append(r.errors, "event has no goroutine ID")
	}
	previous, exists := r.active[event.GoroutineID]
	if event.Clear {
		// A clear may overtake the goroutine's first position callback.
		delete(r.active, event.GoroutineID)
	} else {
		if !event.Position.IsValid() {
			r.errors = append(r.errors, "position event is invalid")
		}
		if !r.allowRepeated && exists && previous.Filename == event.Position.Filename && previous.Line == event.Position.Line {
			r.errors = append(r.errors, "consecutive notifications for the same source line")
		}
		r.active[event.GoroutineID] = event.Position
	}
	r.events = append(r.events, event)
}

func (r *executionRecorder) snapshot() ([]ExecutionPosition, map[uint64]token.Position) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return slices.Clone(r.events), maps.Clone(r.active)
}

func (r *executionRecorder) check(t *testing.T) {
	t.Helper()
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, err := range r.errors {
		t.Error(err)
	}
	if len(r.active) != 0 {
		t.Errorf("positions remain after execution: %v", r.active)
	}
}

func (r *executionRecorder) checkLines(t *testing.T, filename string, want []int) {
	t.Helper()
	events, _ := r.snapshot()
	var lines []int
	for _, event := range events {
		line := event.Position.Line
		if event.Clear {
			line = 0
		} else if filepath.Base(event.Position.Filename) != filename {
			t.Errorf("source filename = %q, want %s", event.Position.Filename, filename)
		}
		lines = append(lines, line)
	}
	if !reflect.DeepEqual(lines, want) {
		t.Fatalf("source lines = %v, want %v", lines, want)
	}
}

func recvExecution[T any](t *testing.T, ch <-chan T) T {
	t.Helper()
	select {
	case value := <-ch:
		return value
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for interpreted execution")
		var zero T
		return zero
	}
}

func executionGate(t *testing.T) (<-chan struct{}, func()) {
	t.Helper()
	release := make(chan struct{})
	var once sync.Once
	unblock := func() { once.Do(func() { close(release) }) }
	t.Cleanup(unblock)
	return release, unblock
}

func TestNewExecutionTrackerRejectsMultipleOptions(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic for multiple options")
		}
	}()
	NewExecutionTracker(nil, ExecutionOptions{}, ExecutionOptions{})
}

func TestExecutionTrackerNestedPositions(t *testing.T) {
	const source = `package main
func child(n int) int { return n + 1 }
func Root(n int) int {
	n++
	x := child(n)
	return x
}`
	ctx, _, recorder := newExecutionTest(t, 0)
	var firstCaller, caller token.Position
	var enteredChild bool
	ctx.SetDebug(func(info *DebugInfo) {
		if enteredChild {
			return
		}
		pos := info.Position()
		switch pos.Line {
		case 5:
			if !firstCaller.IsValid() {
				firstCaller = pos
			}
			caller = pos
		case 2:
			enteredChild = true
		}
	})
	interp := loadExecution(t, ctx, source)
	runExecution(t, interp, "Root", 42, 40)
	// Child return restores the call line before the next DebugRef.
	recorder.checkLines(t, "execution.go", []int{4, 5, 2, 5, 6, 0})
	if !enteredChild || !caller.IsValid() || firstCaller.Column == caller.Column {
		t.Fatalf("call line did not produce different columns before entering child: first=%v last=%v", firstCaller, caller)
	}
	events, _ := recorder.snapshot()
	if restored := events[3].Position; restored != caller {
		t.Errorf("restored position = %v, want latest caller position %v", restored, caller)
	}
}

func TestExecutionTrackerRecursiveLineAndFrameReuse(t *testing.T) {
	const source = `package main
func Recurse(n int) int { if n == 0 { return n }; return Recurse(n-1) }
`
	for _, mode := range []Mode{0, EnableCachedReg} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			ctx, _, recorder := newExecutionTest(t, mode)
			ctx.SetLeastCallForEnablePool(0)
			interp := loadExecution(t, ctx, source)
			for run := 0; run < 8; run++ {
				runExecution(t, interp, "Recurse", 0, 10)
				events, active := recorder.snapshot()
				if len(active) != 0 || len(events) != 2*(run+1) {
					t.Fatalf("run %d: events = %v, active = %v; want one position and one clear per run", run, events, active)
				}
				first, last := events[len(events)-2], events[len(events)-1]
				if first.Clear || first.Position.Line != 2 || !last.Clear || first.GoroutineID != last.GoroutineID {
					t.Fatalf("run %d: unexpected events %v", run, events[len(events)-2:])
				}
			}
		})
	}
}

func TestExecutionTrackerRepeatedPositions(t *testing.T) {
	const loop = `package main
func Root(n int) int { total := 0; for i := 0; i < n; i++ { total += i }; return total }
`
	for _, test := range []struct {
		name   string
		source string
		want   int
	}{
		{"loop", loop, 6},
		{"recursive", `package main
func Root(n int) int { if n == 0 { return n }; return Root(n-1) + 1 }
`, 4},
	} {
		for _, repeated := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/repeated=%t", test.name, repeated), func(t *testing.T) {
				ctx, recorder := NewContext(EnableCachedReg), newExecutionRecorder(t)
				ctx.SetLeastCallForEnablePool(0)
				recorder.allowRepeated = repeated
				var debug []token.Position
				var events []ExecutionPosition
				ctx.SetDebug(func(info *DebugInfo) { debug = append(debug, info.Position()) })
				ctx.SetExecutionTracker(NewExecutionTracker(func(event ExecutionPosition) {
					events = append(events, event)
					recorder.record(event)
				}, ExecutionOptions{ReportRepeated: repeated}))
				runExecution(t, loadExecution(t, ctx, test.source), "Root", test.want, 4)
				if len(debug) < 2 {
					t.Fatal("source did not produce repeated DebugRefs")
				}
				var positions []token.Position
				var sequence uint64
				var clears int
				for index, event := range events {
					if event.Sequence <= sequence || event.GoroutineID != GoroutineID() {
						t.Fatalf("invalid event ordering or goroutine: %v", events)
					}
					sequence = event.Sequence
					if event.Clear {
						clears++
						if index != len(events)-1 || event.Position != (token.Position{}) {
							t.Fatalf("clear must be last and have no position: %v", events)
						}
					} else {
						positions = append(positions, event.Position)
					}
				}
				want := debug
				if !repeated {
					want = debug[:1]
				}
				// Exact positions also check columns, and exclude spurious restores
				// between recursive calls on the same line.
				if !reflect.DeepEqual(positions, want) || clears != 1 {
					t.Fatalf("positions = %v, clears = %d; want %v and one clear", positions, clears, want)
				}
			})
		}
	}
}

func TestExecutionTrackerUnwindingAndOtherHooks(t *testing.T) {
	tests := []struct {
		name      string
		source    string
		wantError bool
		wantCode  int
	}{
		{
			name: "recovered_named_result",
			source: `package main
func helper() (n int) {
	defer func() { recover(); n = 42 }()
	panic("boom")
}
func main() { if helper() != 42 { panic("result") } }`,
		},
		{
			name: "panic_in_defer",
			source: `package main
func cleanup(n int) { panic(n) }
func helper(n int) { defer cleanup(n); n++ }
func main() { helper(1) }`,
			wantError: true,
		},
		{
			name: "runtime_panic",
			source: `package main
func helper(n int) int { return 42 / n }
func main() { _ = helper(0) }`,
			wantError: true,
		},
		{
			name: "os_exit",
			source: `package main
import "os"
func helper(n int) { os.Exit(n) }
func main() { helper(7) }`,
			wantCode: 7,
		},
		{
			name: "goexit_main",
			source: `package main
import "runtime"
var n int
func cleanup() { n++ }
func main() { defer cleanup(); runtime.Goexit() }`,
			wantError: true,
		},
	}
	for _, repeated := range []bool{false, true} {
		for _, test := range tests {
			t.Run(fmt.Sprintf("%s/repeated=%t", test.name, repeated), func(t *testing.T) {
				ctx, recorder := NewContext(0), newExecutionRecorder(t)
				recorder.allowRepeated = repeated
				var timeline []string
				var debugCalls int
				ctx.SetDebug(func(*DebugInfo) { debugCalls++ })
				ctx.SetFuncHook(nil, func(info *FuncInfo) { timeline = append(timeline, "leave:"+info.Name()) })
				ctx.SetExecutionTracker(NewExecutionTracker(func(event ExecutionPosition) {
					recorder.record(event)
					if event.Clear {
						timeline = append(timeline, "clear")
					}
				}, ExecutionOptions{ReportRepeated: repeated}))
				code, err := ctx.RunFile("main.go", test.source, nil)
				if (err != nil) != test.wantError || (!test.wantError && code != test.wantCode) {
					t.Fatalf("RunFile = (%d, %v), want code %d, error %v", code, err, test.wantCode, test.wantError)
				}
				if debugCalls == 0 {
					t.Error("SetDebug callback did not run alongside execution hook")
				}
				if len(timeline) < 2 || timeline[len(timeline)-2] != "leave:main" || timeline[len(timeline)-1] != "clear" {
					t.Errorf("final position was not cleared after function leave: %v", timeline)
				}
			})
		}
	}
}

func TestExecutionTrackerRestoresAcrossHostCallback(t *testing.T) {
	const source = `package main
func Invoke(callback func())
func callback() { n := 1; _ = n }
func Root() { Invoke(callback) }
`
	ctx, _, recorder := newExecutionTest(t, 0)
	var callbackRan bool
	ctx.RegisterExternal("main.Invoke", func(callback func()) {
		id := GoroutineID()
		_, before := recorder.snapshot()
		caller, ok := before[id]
		if !ok {
			t.Error("caller has no position before the host callback")
		}
		callback()
		callbackRan = true
		_, after := recorder.snapshot()
		if restored, ok := after[id]; !ok || restored.Filename != caller.Filename || restored.Line != caller.Line {
			t.Errorf("position after callback = %v, want caller %v", after, caller)
		}
	})
	interp := loadExecution(t, ctx, source)
	if _, err := interp.RunFunc("Root"); err != nil {
		t.Fatal(err)
	}
	if !callbackRan {
		t.Fatal("host did not invoke the interpreted callback")
	}
	recorder.checkLines(t, "execution.go", []int{4, 3, 4, 0})
}

func TestExecutionTrackerReentrantNotification(t *testing.T) {
	for _, repeated := range []bool{false, true} {
		t.Run(fmt.Sprintf("repeated=%t", repeated), func(t *testing.T) {
			ctx, recorder := NewContext(0), newExecutionRecorder(t)
			recorder.allowRepeated = repeated
			var interp *Interp
			var delayed ExecutionPosition
			ctx.SetExecutionTracker(NewExecutionTracker(func(event ExecutionPosition) {
				if delayed.Sequence == 0 && !event.Clear {
					delayed = event
					runExecution(t, interp, "Nested", 42, 40)
					before, active := recorder.snapshot()
					if len(before) == 0 || active[event.GoroutineID] != event.Position {
						t.Fatalf("reentrant call did not restore %v: events=%v active=%v", event.Position, before, active)
					}
					if restored := before[len(before)-1]; restored.Clear || restored.Sequence <= event.Sequence {
						t.Fatalf("restoration did not advance the sequence: delayed=%v restored=%v", event, restored)
					}
					// The nested call restores the caller before its old notification arrives.
					recorder.record(event)
					if after, _ := recorder.snapshot(); !reflect.DeepEqual(after, before) {
						t.Fatalf("late notification passed the sequence watermark: before=%v after=%v", before, after)
					}
					return
				}
				recorder.record(event)
			}, ExecutionOptions{ReportRepeated: repeated}))
			interp = loadExecution(t, ctx, `package main
func Root(n int) int { return n + 1 }
func Nested(n int) int { return n + 2 }
`)
			runExecution(t, interp, "Root", 42, 41)
			events, active := recorder.snapshot()
			if delayed.Sequence == 0 {
				t.Fatal("notification did not reenter the interpreter")
			}
			var lines []int
			var clears int
			sequence := delayed.Sequence
			for _, event := range events {
				if event.Sequence <= sequence || event.GoroutineID != delayed.GoroutineID {
					t.Fatalf("invalid reentrant event order or goroutine: %v", events)
				}
				sequence = event.Sequence
				line := event.Position.Line
				if event.Clear {
					line = 0
					clears++
				}
				if len(lines) == 0 || lines[len(lines)-1] != line {
					lines = append(lines, line)
				}
			}
			if !reflect.DeepEqual(lines, []int{3, 2, 0}) || clears != 1 || len(active) != 0 {
				t.Fatalf("reentrant execution did not restore and clear: events=%v active=%v", events, active)
			}
		})
	}
}

func TestExecutionTrackerGoroutinesOnSameLine(t *testing.T) {
	ctx, _, recorder := newExecutionTest(t, 0)
	interp := loadExecution(t, ctx, `package main
func Worker(id int, wait func(int)) { wait(id) }
`)
	type arrival struct {
		worker int
		id     uint64
	}
	ready := make(chan arrival, 2)
	first, releaseFirst := executionGate(t)
	second, releaseSecond := executionGate(t)
	release := []<-chan struct{}{first, second}
	done := []chan error{make(chan error, 1), make(chan error, 1)}
	wait := func(worker int) {
		ready <- arrival{worker, GoroutineID()}
		<-release[worker]
	}
	for worker := 0; worker < 2; worker++ {
		go func(worker int) {
			_, err := interp.RunFunc("Worker", worker, wait)
			done[worker] <- err
		}(worker)
	}
	ids := make(map[int]uint64)
	for worker := 0; worker < 2; worker++ {
		arrived := recvExecution(t, ready)
		ids[arrived.worker] = arrived.id
	}
	_, active := recorder.snapshot()
	if len(active) != 2 || active[ids[0]].Line != 2 || active[ids[1]].Line != 2 {
		t.Errorf("both workers must own the same line independently: %v", active)
	}
	releaseFirst()
	if err := recvExecution(t, done[0]); err != nil {
		t.Error(err)
	}
	_, active = recorder.snapshot()
	if len(active) != 1 || active[ids[1]].Line != 2 {
		t.Errorf("first worker cleared the second worker's position: %v", active)
	}
	releaseSecond()
	if err := recvExecution(t, done[1]); err != nil {
		t.Error(err)
	}
}

func TestExecutionTrackerCloseIgnoresLateHostCallback(t *testing.T) {
	ctx, tracker, recorder := newExecutionTest(t, 0)
	ready := make(chan struct{})
	release, unblock := executionGate(t)
	var callbackRan bool
	ctx.RegisterExternal("main.Block", func(callback func()) {
		close(ready)
		<-release
		callback()
		callbackRan = true
	})
	interp := loadExecution(t, ctx, `package main
func Block(callback func())
func callback() { n := 1; _ = n }
func Root() { Block(callback) }
`)
	done := make(chan error, 1)
	go func() {
		_, err := interp.RunFunc("Root")
		done <- err
	}()
	recvExecution(t, ready)
	if _, active := recorder.snapshot(); len(active) != 1 {
		t.Errorf("blocked host call has no source position: %v", active)
	}
	tracker.Close()
	stopped, active := recorder.snapshot()
	if len(active) != 0 {
		t.Errorf("Close did not immediately clear positions: %v", active)
	}
	tracker.Close() // Closing a session twice must be harmless.
	unblock()
	if err := recvExecution(t, done); err != nil {
		t.Error(err)
	}
	if !callbackRan {
		t.Error("Close interrupted the program's host callback")
	}
	finished, _ := recorder.snapshot()
	if len(finished) != len(stopped) {
		t.Errorf("late host callback emitted events after Close: before %v, after %v", stopped, finished)
	}
}

func TestExecutionTrackerRunMainKeepsSessionOpen(t *testing.T) {
	ctx, tracker, recorder := newExecutionTest(t, 0)
	ready, finished := make(chan uint64, 1), make(chan struct{})
	release, unblock := executionGate(t)
	ctx.RegisterExternal("main.Hold", func() {
		ready <- GoroutineID()
		<-release
	})
	var workerID uint64
	var callback func(int) int
	ctx.RegisterExternal("main.Await", func() { workerID = <-ready })
	ctx.RegisterExternal("main.Save", func(fn func(int) int) { callback = fn })
	ctx.RegisterExternal("main.Finished", func() { close(finished) })
	interp := loadExecution(t, ctx, `package main
func Hold()
func Await()
func Save(func(int) int)
func Finished()
func worker() { defer Finished(); Hold() }
func Callback(n int) int { return n + 1 }
func main() { Save(Callback); go worker(); Await() }
`)
	if err := interp.RunInit(); err != nil {
		t.Fatal(err)
	}
	if code, err := interp.RunMain(); code != 0 || err != nil {
		t.Fatalf("RunMain = (%d, %v), want (0, nil)", code, err)
	}
	before, active := recorder.snapshot()
	if len(active) != 1 || active[workerID].Line != 6 {
		t.Errorf("RunMain cleared the blocked worker: %v", active)
	}
	if callback == nil {
		t.Fatal("main did not export its callback")
	}
	if got := callback(41); got != 42 {
		t.Fatalf("callback after RunMain = %d, want 42", got)
	}
	after, active := recorder.snapshot()
	if len(after) != len(before)+2 || after[len(before)].Position.Line != 7 || !after[len(after)-1].Clear {
		t.Errorf("host callback after RunMain was not tracked: %v", after[len(before):])
	}
	if len(active) != 1 || active[workerID].Line != 6 {
		t.Errorf("host callback changed the blocked worker's position: %v", active)
	}
	tracker.Close()
	stopped, active := recorder.snapshot()
	if len(active) != 0 {
		t.Errorf("Close left worker positions visible: %v", active)
	}
	if got := callback(42); got != 43 {
		t.Errorf("callback after Close = %d, want 43", got)
	}
	unblock()
	recvExecution(t, finished)
	if events, _ := recorder.snapshot(); len(events) != len(stopped) {
		t.Errorf("execution emitted events after Close: %v", events[len(stopped):])
	}
}

func TestExecutionTrackerCloseOrdersInFlightNotification(t *testing.T) {
	for _, repeated := range []bool{false, true} {
		for _, pendingClear := range []bool{false, true} {
			t.Run(fmt.Sprintf("repeated=%t/clear=%t", repeated, pendingClear), func(t *testing.T) {
				ctx, recorder := NewContext(0), newExecutionRecorder(t)
				recorder.allowRepeated = repeated
				started := make(chan ExecutionPosition, 1)
				release, unblock := executionGate(t)
				var blocked atomic.Bool
				tracker := NewExecutionTracker(func(event ExecutionPosition) {
					if event.Clear == pendingClear && blocked.CompareAndSwap(false, true) {
						started <- event
						<-release
					}
					recorder.record(event)
				}, ExecutionOptions{ReportRepeated: repeated})
				ctx.SetExecutionTracker(tracker)
				interp := loadExecution(t, ctx, executionTestSource)
				done := make(chan error, 1)
				go func() {
					_, err := interp.RunFunc("Root", 1)
					done <- err
				}()
				pending := recvExecution(t, started)
				before, _ := recorder.snapshot()
				// Close must return without waiting; its clear supersedes the blocked event.
				closed := make(chan struct{})
				go func() {
					tracker.Close()
					close(closed)
				}()
				recvExecution(t, closed)
				stopped, active := recorder.snapshot()
				if len(active) != 0 || len(stopped) != len(before)+1 {
					t.Errorf("Close did not immediately clear the pending notification: events=%v active=%v", stopped, active)
				} else if last := stopped[len(stopped)-1]; !last.Clear || last.Sequence <= pending.Sequence || last.GoroutineID != pending.GoroutineID {
					t.Errorf("Close did not supersede the pending notification: pending=%v clear=%v", pending, last)
				}
				unblock()
				if err := recvExecution(t, done); err != nil {
					t.Error(err)
				}
				events, active := recorder.snapshot()
				if !reflect.DeepEqual(events, stopped) || len(active) != 0 {
					t.Errorf("in-flight notification changed state after Close: stopped=%v events=%v active=%v", stopped, events, active)
				}
			})
		}
	}
}

func TestExecutionTrackerReentrantClear(t *testing.T) {
	for _, repeated := range []bool{false, true} {
		for _, closeInside := range []bool{false, true} {
			t.Run(fmt.Sprintf("repeated=%t/close=%t", repeated, closeInside), func(t *testing.T) {
				ctx, recorder := NewContext(0), newExecutionRecorder(t)
				recorder.allowRepeated = repeated
				var interp *Interp
				var tracker *ExecutionTracker
				var reentered, closing bool
				tracker = NewExecutionTracker(func(event ExecutionPosition) {
					if !event.Clear {
						recorder.record(event)
						return
					}
					switch {
					case !reentered:
						reentered = true
						runExecution(t, interp, "Nested", 42, 40)
					case closeInside && !closing:
						closing = true
						tracker.Close()
					default:
						recorder.record(event)
						if closing {
							tracker.Close()
						}
						return
					}
					before, active := recorder.snapshot()
					if len(active) != 0 || len(before) == 0 || !before[len(before)-1].Clear || before[len(before)-1].Sequence <= event.Sequence {
						t.Errorf("reentrant clear did not supersede the pending clear: pending=%v events=%v active=%v", event, before, active)
					}
					recorder.record(event)
					if after, _ := recorder.snapshot(); !reflect.DeepEqual(after, before) {
						t.Errorf("older clear changed state after reentry: before=%v after=%v", before, after)
					}
				}, ExecutionOptions{ReportRepeated: repeated})
				ctx.SetExecutionTracker(tracker)
				interp = loadExecution(t, ctx, `package main
func Root(n int) int { return n + 1 }
func Nested(n int) int { return n + 2 }
`)
				runExecution(t, interp, "Root", 42, 41)
				if !reentered || closing != closeInside {
					t.Fatalf("clear did not reenter as requested: reentered=%t closing=%t", reentered, closing)
				}
				before, _ := recorder.snapshot()
				tracker.Close()
				if after, _ := recorder.snapshot(); !reflect.DeepEqual(after, before) {
					t.Errorf("completed clear callbacks left positions to clear again: before=%v after=%v", before, after)
				}
			})
		}
	}
}

func TestExecutionTrackerCloseMultipleGoroutines(t *testing.T) {
	for _, repeated := range []bool{false, true} {
		t.Run(fmt.Sprintf("repeated=%t", repeated), func(t *testing.T) {
			const workers = 3
			ctx, recorder := NewContext(0), newExecutionRecorder(t)
			recorder.allowRepeated = repeated
			ready := make(chan struct{}, workers)
			release, unblock := executionGate(t)
			ctx.RegisterExternal("main.Hold", func(int) {
				ready <- struct{}{}
				<-release
			})
			clearing := make(chan struct{})
			finishClear, unblockClear := executionGate(t)
			var first sync.Once
			var mu sync.Mutex
			var raw []ExecutionPosition
			var tracker *ExecutionTracker
			tracker = NewExecutionTracker(func(event ExecutionPosition) {
				mu.Lock()
				raw = append(raw, event)
				mu.Unlock()
				recorder.record(event)
				if event.Clear {
					tracker.Close()
					first.Do(func() {
						close(clearing)
						<-finishClear
					})
				}
			}, ExecutionOptions{ReportRepeated: repeated})
			ctx.SetExecutionTracker(tracker)
			interp := loadExecution(t, ctx, `package main
func Hold(int)
func Worker(n int) int {
	Hold(n)
	return n + 1
}
`)
			done := make(chan error, workers)
			for worker := range workers {
				go func() {
					got, err := interp.RunFunc("Worker", worker)
					if err == nil && got != worker+1 {
						err = fmt.Errorf("Worker(%d) = %v, want %d", worker, got, worker+1)
					}
					done <- err
				}()
			}
			for range workers {
				recvExecution(t, ready)
			}
			before, visible := recorder.snapshot()
			if len(visible) != workers {
				t.Fatalf("workers have no distinct visible positions: %v", visible)
			}
			closed := make(chan struct{})
			go func() { tracker.Close(); close(closed) }()
			recvExecution(t, clearing)
			// Other Close calls and natural returns must finish during a clear callback.
			closedAgain := make(chan struct{})
			go func() { tracker.Close(); tracker.Close(); close(closedAgain) }()
			unblock()
			recvExecution(t, closedAgain)
			for range workers {
				if err := recvExecution(t, done); err != nil {
					t.Error(err)
				}
			}
			mu.Lock()
			during := len(raw)
			mu.Unlock()
			if during != len(before)+1 {
				t.Errorf("returns or repeated Close emitted events during the blocked clear: got %d, want %d", during, len(before)+1)
			}
			unblockClear()
			recvExecution(t, closed)
			mu.Lock()
			events := slices.Clone(raw)
			mu.Unlock()
			if len(events) != len(before)+workers {
				t.Fatalf("Close events = %v, want one clear for each worker", events)
			}
			var sequence uint64
			for _, event := range before {
				sequence = max(sequence, event.Sequence)
			}
			clears := make(map[uint64]int)
			for _, event := range events[len(before):] {
				if _, ok := visible[event.GoroutineID]; !ok || !event.Clear || event.Position != (token.Position{}) || event.Sequence <= sequence {
					t.Fatalf("invalid clear batch: %v", events[len(before):])
				}
				sequence = event.Sequence
				clears[event.GoroutineID]++
			}
			for id := range visible {
				if clears[id] != 1 {
					t.Errorf("goroutine %d has %d clears, want 1", id, clears[id])
				}
			}
			runExecution(t, interp, "Worker", 4, 3)
			mu.Lock()
			unchanged := reflect.DeepEqual(raw, events)
			mu.Unlock()
			if !unchanged {
				t.Error("execution emitted new events after Close")
			}
		})
	}
}

func TestExecutionTrackerCloseCannotBeReopened(t *testing.T) {
	ctx, tracker, recorder := newExecutionTest(t, 0)
	interp := loadExecution(t, ctx, executionTestSource)
	runExecution(t, interp, "Root", 2, 1)
	before, _ := recorder.snapshot()
	tracker.Close()
	interp.Abort()
	if err := interp.RunInit(); err != nil {
		t.Fatal(err)
	}
	runExecution(t, interp, "Root", 42, 41)
	if after, _ := recorder.snapshot(); len(before) != 2 || len(after) != len(before) {
		t.Errorf("RunInit reopened a closed tracker: before=%v after=%v", before, after)
	}

	// New sessions need a new tracker and interpreter.
	next := newExecutionRecorder(t)
	ctx.SetExecutionTracker(NewExecutionTracker(next.record, ExecutionOptions{}))
	fresh := loadExecution(t, ctx, executionTestSource)
	runExecution(t, fresh, "Root", 43, 42)
	next.checkLines(t, "execution.go", []int{2, 0})
}

func TestExecutionTrackerClosedUpdate(t *testing.T) {
	var callbacks int
	tracker := NewExecutionTracker(func(ExecutionPosition) { callbacks++ })
	tracker.Close()

	tracker.mu.Lock()
	event := tracker.updateLocked(1, token.Position{Filename: "closed.go", Line: 1})
	tracker.mu.Unlock()
	tracker.notify(event)
	if event.Sequence != 0 || callbacks != 0 {
		t.Fatalf("closed tracker accepted update: event=%v callbacks=%d", event, callbacks)
	}
}

func TestExecutionTrackerCapturedAtCompilation(t *testing.T) {
	ctx, _, original := newExecutionTest(t, 0)
	replacement := newExecutionRecorder(t)
	interp := loadExecution(t, ctx, executionTestSource)
	ctx.SetExecutionTracker(NewExecutionTracker(replacement.record, ExecutionOptions{}))
	runExecution(t, interp, "Root", 42, 41)
	if events, _ := original.snapshot(); len(events) != 2 {
		t.Errorf("compiled hook events = %v, want a position and clear", events)
	}
	if events, _ := replacement.snapshot(); len(events) != 0 {
		t.Errorf("replacement hook affected an already compiled interpreter: %v", events)
	}
}

func TestExecutionTrackerDisabledPreservesExistingHooks(t *testing.T) {
	ctx := NewContext(0)
	ctx.SetExecutionTracker(nil)
	var entries, leaves, debugCalls int
	ctx.SetFuncHook(func(*FuncInfo) { entries++ }, func(*FuncInfo) { leaves++ })
	ctx.SetDebug(func(*DebugInfo) { debugCalls++ })
	interp := loadExecution(t, ctx, executionTestSource)
	runExecution(t, interp, "Root", 42, 41)
	if entries != 1 || leaves != 1 || debugCalls == 0 {
		t.Errorf("existing hooks changed: entry=%d leave=%d debug=%d", entries, leaves, debugCalls)
	}
}

func TestExecutionTrackerReplPreservesValues(t *testing.T) {
	for _, mode := range []Mode{0, EnableCachedReg} {
		t.Run(fmt.Sprint(mode), func(t *testing.T) {
			ctx, tracker, recorder := newExecutionTest(t, mode)
			repl := NewRepl(ctx)
			repl.SetFileName("execution_repl.go")
			for _, step := range []struct {
				source       string
				want         string
				wantPosition bool
				disabled     bool
			}{
				{`a := 1`, `[]`, true, false},
				{`b := 2`, `[]`, true, false},
				{`a+b`, `[3 int]`, true, false},
				{`func add(n int) int { return n + 10 }`, `[]`, false, false},
				{`add(a+b)`, `[13 int]`, true, false},
				{`a += 4`, `[]`, true, false},
				{`add(a+b)`, `[17 int]`, true, false},
				{`a += 1`, `[]`, false, true},
				{`add(a+b)`, `[18 int]`, false, true},
				{`b += 1`, `[]`, true, false},
				{`add(a+b)`, `[19 int]`, true, false},
			} {
				// Eval recompiles; the tracker keeps one sequence across interpreters.
				if step.disabled {
					ctx.SetExecutionTracker(nil)
				} else {
					ctx.SetExecutionTracker(tracker)
				}
				before, _ := recorder.snapshot()
				_, values, err := repl.Eval(step.source)
				if err != nil {
					t.Fatalf("Eval(%q): %v", step.source, err)
				}
				if got := fmt.Sprint(values); got != step.want {
					t.Fatalf("Eval(%q) = %s, want %s", step.source, got, step.want)
				}
				after, _ := recorder.snapshot()
				events := after[len(before):]
				if step.disabled && len(events) != 0 {
					t.Errorf("Eval(%q) emitted events while tracking was disabled: %v", step.source, events)
				}
				if step.wantPosition && len(events) < 2 {
					t.Errorf("Eval(%q) emitted no position/clear pair: %v", step.source, events)
				}
				if len(events) > 0 && !events[len(events)-1].Clear {
					t.Errorf("Eval(%q) left its final position visible: %v", step.source, events)
				}
				recorder.check(t)
			}
		})
	}
}

func TestExecutionTrackerGoexitInHostGoroutine(t *testing.T) {
	ctx, _, recorder := newExecutionTest(t, 0)
	cleaned := make(chan int, 1)
	ctx.RegisterExternal("main.Finished", func(n int) { cleaned <- n })
	interp := loadExecution(t, ctx, `package main
import "runtime"
func Finished(n int)
func cleanup(n int) { Finished(n) }
func Worker(n int) { defer cleanup(n); runtime.Goexit() }
`)
	done := make(chan struct{})
	go func() {
		// Worker runs off the creating goroutine, exercising native Goexit.
		defer close(done)
		_, _ = interp.RunFunc("Worker", 42)
	}()
	recvExecution(t, done)
	if got := recvExecution(t, cleaned); got != 42 {
		t.Errorf("user defer received %d, want 42", got)
	}
	events, _ := recorder.snapshot()
	if len(events) < 2 || !events[len(events)-1].Clear {
		t.Errorf("Goexit did not clear its final position: %v", events)
	}
	var sawCleanup bool
	for _, event := range events {
		if !event.Clear && event.Position.Line == 4 {
			sawCleanup = true
		}
	}
	if !sawCleanup {
		t.Error("user defer did not report its source position")
	}
}

func TestExecutionTrackerGCAndMultiResultFrameReuse(t *testing.T) {
	ctx, _, recorder := newExecutionTest(t, ExperimentalSupportGC|EnableCachedReg)
	ctx.SetLeastCallForEnablePool(0)
	interp := loadExecution(t, ctx, `package main
import "runtime"
func pair(n int) (int, int) { return n, n + 1 }
func Root(n int) (int, int) {
	runtime.GC()
	total := 0
	for i := 0; i < 3; i++ {
		a, b := pair(n + i)
		total += a + b
	}
	return total, n
}`)
	// Require spare child frame reuse alongside the ordinary pool.
	root := interp.funcs[interp.mainpkg.Func("Root")]
	pair := interp.funcs[interp.mainpkg.Func("pair")]
	if root.childFrameReg < 0 || !pair.frameCacheSafe {
		t.Fatalf("multi-result source did not enable child frame reuse: child register %d, safe %v, clear %v, locals %v", root.childFrameReg, pair.frameCacheSafe, pair.clearRanges, pair.localRefs)
	}
	for run := 0; run < 4; run++ {
		before, _ := recorder.snapshot()
		runExecution(t, interp, "Root", Tuple{6*run + 9, run}, run)
		after, active := recorder.snapshot()
		events := after[len(before):]
		if len(events) == 0 || !events[len(events)-1].Clear || len(active) != 0 {
			t.Fatalf("run %d: final position was not cleared: %v", run, events)
		}
		var restored int
		for index, event := range events {
			if event.Clear || event.Position.Line != 3 {
				continue
			}
			if index+1 >= len(events) || events[index+1].Clear || events[index+1].Position.Line != 8 {
				t.Fatalf("run %d: pair did not restore its caller after GC: %v", run, events)
			}
			restored++
		}
		if restored != 3 {
			t.Fatalf("run %d: observed %d returns after GC, want 3", run, restored)
		}
	}
}

func TestExecutionTrackerFiltersBeforeSavingPosition(t *testing.T) {
	for _, repeated := range []bool{false, true} {
		t.Run(fmt.Sprintf("repeated=%t", repeated), func(t *testing.T) {
			ctx, recorder := NewContext(0), newExecutionRecorder(t)
			recorder.allowRepeated = repeated
			ctx.SetExecutionTracker(NewExecutionTracker(recorder.record, ExecutionOptions{
				Filter:         func(pos token.Position) bool { return filepath.Base(pos.Filename) == "user.go" },
				ReportRepeated: repeated,
			}))
			var libraryDebugs, acceptedDebugs, checks int
			ctx.SetDebug(func(info *DebugInfo) {
				switch filepath.Base(info.Position().Filename) {
				case "library.go":
					libraryDebugs++
				case "user.go":
					acceptedDebugs++
				}
			})
			ctx.RegisterExternal("main.Check", func() {
				checks++
				_, active := recorder.snapshot()
				pos, ok := active[GoroutineID()]
				if !ok || filepath.Base(pos.Filename) != "user.go" || pos.Line != 10 {
					t.Errorf("callback returned to filtered library with position %v, want user.go:10", pos)
				}
			})
			interp := loadExecution(t, ctx, `package main
//line user.go:1
func Check()
func callback() { n := 1; _ = n }
//line library.go:1
func library() { callback(); Check() }
//line user.go:10
func Root() { library() }
`)
			runExecution(t, interp, "Root", nil)
			if checks != 1 || libraryDebugs == 0 {
				t.Errorf("filtered code did not execute or reach SetDebug: checks=%d debug=%d", checks, libraryDebugs)
			}
			if !repeated {
				recorder.checkLines(t, "user.go", []int{10, 2, 10, 0})
				return
			}
			events, _ := recorder.snapshot()
			// One caller restoration and one final clear supplement the accepted refs.
			if len(events) != acceptedDebugs+2 || !events[len(events)-1].Clear {
				t.Fatalf("filtered events = %v, want %d DebugRefs, one restore and one clear", events, acceptedDebugs)
			}
			for _, event := range events {
				if !event.Clear && filepath.Base(event.Position.Filename) != "user.go" {
					t.Errorf("filtered position reported: %v", event)
				}
			}
		})
	}
}

func TestExecutionTrackerRestoresAcrossInterpreters(t *testing.T) {
	outerContext, tracker, recorder := newExecutionTest(t, SupportMultipleInterp)
	innerContext := NewContext(SupportMultipleInterp)
	innerContext.SetExecutionTracker(tracker)
	inner := loadExecution(t, innerContext, `package main
//line inner.go:20
func Root(n int) int { return n + 1 }
`)
	var calls int
	outerContext.RegisterExternal("main.CallInner", func() {
		calls++
		id := GoroutineID()
		_, before := recorder.snapshot()
		caller, ok := before[id]
		if !ok {
			t.Error("outer interpreter has no position before the host call")
		}
		runExecution(t, inner, "Root", 42, 41)
		_, after := recorder.snapshot()
		if restored, ok := after[id]; !ok || restored.Filename != caller.Filename || restored.Line != caller.Line {
			t.Errorf("inner interpreter did not restore its outer caller: got %v, want %v", after, caller)
		}
	})
	outer := loadExecution(t, outerContext, `package main
func CallInner()
func Root() { CallInner() }
`)
	runExecution(t, outer, "Root", nil)
	if calls != 1 {
		t.Fatalf("host calls = %d, want 1", calls)
	}
	events, _ := recorder.snapshot()
	var locations []string
	for _, event := range events {
		location := "clear"
		if !event.Clear {
			location = fmt.Sprintf("%s:%d", filepath.Base(event.Position.Filename), event.Position.Line)
		}
		locations = append(locations, location)
	}
	want := []string{"execution.go:3", "inner.go:20", "execution.go:3", "clear"}
	if !reflect.DeepEqual(locations, want) {
		t.Errorf("shared tracker locations = %v, want %v", locations, want)
	}
}

func TestExecutionTrackerInterpreterLifecycleDoesNotCloseSession(t *testing.T) {
	for _, test := range []struct {
		name string
		run  func(*testing.T, *Context, *Interp)
	}{
		{"abort", func(_ *testing.T, _ *Context, interp *Interp) { interp.Abort() }},
		{"init", func(t *testing.T, _ *Context, interp *Interp) {
			if err := interp.RunInit(); err != nil {
				t.Fatal(err)
			}
		}},
		{"main", func(t *testing.T, _ *Context, interp *Interp) {
			if code, err := interp.RunMain(); code != 0 || err != nil {
				t.Fatalf("RunMain = (%d, %v), want (0, nil)", code, err)
			}
		}},
		{"run_interp", func(t *testing.T, ctx *Context, interp *Interp) {
			if code, err := ctx.RunInterp(interp, "execution.go", nil); code != 0 || err != nil {
				t.Fatalf("RunInterp = (%d, %v), want (0, nil)", code, err)
			}
		}},
		{"release", func(_ *testing.T, _ *Context, interp *Interp) { interp.UnsafeRelease() }},
	} {
		t.Run(test.name, func(t *testing.T) {
			ctx, _, recorder := newExecutionTest(t, 0)
			const source = executionTestSource + "func main() {}\n"
			interp := loadExecution(t, ctx, source)
			runExecution(t, interp, "Root", 2, 1)
			test.run(t, ctx, interp)
			before, _ := recorder.snapshot()
			fresh := loadExecution(t, ctx, source)
			runExecution(t, fresh, "Root", 42, 41)
			after, _ := recorder.snapshot()
			if len(after) != len(before)+2 {
				t.Fatalf("interpreter lifecycle closed its shared tracker: before=%v after=%v", before, after)
			}
			if after[len(before)].Sequence <= before[len(before)-1].Sequence {
				t.Error("sequence was reset by an interpreter lifecycle operation")
			}
		})
	}
}

func TestExecutionMetadataPreservesRecursivePositions(t *testing.T) {
	const source = `package main
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}`
	for _, mode := range []string{"tracker", "late_debug"} {
		t.Run(mode, func(t *testing.T) {
			var want []token.Position
			for _, baseline := range []bool{true, false} {
				ctx := NewContext(0)
				var positions []token.Position
				if mode == "tracker" {
					tracker := NewExecutionTracker(func(event ExecutionPosition) {
						if !event.Clear {
							positions = append(positions, event.Position)
						}
					}, ExecutionOptions{ReportRepeated: true})
					t.Cleanup(tracker.Close)
					ctx.SetExecutionTracker(tracker)
				} else {
					ctx.SetDebug(nil)
				}
				if baseline {
					// A function hook disables memoization for the reference run.
					ctx.SetFuncHook(func(*FuncInfo) {}, nil)
				}
				interp := loadExecution(t, ctx, source)
				t.Cleanup(interp.UnsafeRelease)
				if mode == "late_debug" {
					ctx.SetDebug(func(info *DebugInfo) {
						positions = append(positions, info.Position())
					})
				}
				runExecution(t, interp, "Fib", 5, 5)
				if baseline {
					if len(positions) == 0 {
						t.Fatal("reference run reported no positions")
					}
					want = positions
				} else if !reflect.DeepEqual(positions, want) {
					t.Fatalf("recursive positions = %v, want %v", positions, want)
				}
			}
		})
	}
}

func TestMakeInstrRejectsUnknownInstruction(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("makeInstr accepted an unknown instruction")
		}
	}()
	makeInstr(nil, nil, nil)
}

func TestSetDebugAfterCompilation(t *testing.T) {
	for _, installed := range []bool{false, true} {
		t.Run(fmt.Sprint(installed), func(t *testing.T) {
			ctx := NewContext(0)
			var originalCalls int
			if installed {
				ctx.SetDebug(func(*DebugInfo) { originalCalls++ })
			} else {
				// Keep debug metadata for a callback installed later.
				ctx.SetDebug(nil)
			}
			interp := loadExecution(t, ctx, executionTestSource)
			runExecution(t, interp, "Root", 42, 41)
			if installed && originalCalls == 0 {
				t.Fatal("original callback did not run")
			}
			before := originalCalls
			var replacementCalls int
			var sawParameter, sawNonVariable bool
			ctx.SetDebug(func(info *DebugInfo) {
				replacementCalls++
				variable, value, ok := info.AsVar()
				if ok && variable.Name() == "n" && !info.IsAddr {
					sawParameter = true
					if value != 41 {
						t.Errorf("parameter = %v, want 41", value)
					}
				}
				if !ok {
					sawNonVariable = true
					if variable != nil || value != nil {
						t.Errorf("non-variable DebugRef returned (%v, %v, false)", variable, value)
					}
				}
			})
			runExecution(t, interp, "Root", 42, 41)
			if originalCalls != before || replacementCalls == 0 || !sawParameter || !sawNonVariable {
				t.Fatalf("callback replacement failed: original=%d/%d replacement=%d parameter=%v non-variable=%v", originalCalls, before, replacementCalls, sawParameter, sawNonVariable)
			}
			before = replacementCalls
			ctx.SetDebug(nil)
			runExecution(t, interp, "Root", 42, 41)
			if replacementCalls != before {
				t.Error("cleared debug callback still ran")
			}
		})
	}
}

func TestExecutionFilterPreservesDebugCallbacks(t *testing.T) {
	ctx := NewContext(0)
	var positions, filtered int
	tracker := NewExecutionTracker(func(ExecutionPosition) { positions++ }, ExecutionOptions{
		Filter: func(token.Position) bool { filtered++; return false },
	})
	t.Cleanup(tracker.Close)
	ctx.SetExecutionTracker(tracker)
	interp := loadExecution(t, ctx, executionTestSource)
	var debugCalls int
	var sawParameter bool
	// Late debug callbacks must see values even at filtered positions.
	ctx.SetDebug(func(info *DebugInfo) {
		debugCalls++
		if variable, value, ok := info.AsVar(); ok && variable.Name() == "n" && !info.IsAddr {
			sawParameter = value == 41
		}
	})
	runExecution(t, interp, "Root", 42, 41)
	if filtered == 0 || positions != 0 || debugCalls == 0 || !sawParameter {
		t.Fatalf("filter affected debugging: filtered=%d positions=%d debug=%d parameter=%v", filtered, positions, debugCalls, sawParameter)
	}
}

// Highlighting modes update a synchronized map. events/op counts submissions
// before host deduplication; highlights/op counts visible source-line changes.
// Editor rendering and asynchronous event delivery are excluded.
func BenchmarkExecutionHighlights(b *testing.B) {
	for _, workload := range []struct {
		name string
		args []Value
		want int
	}{
		{"loop", []Value{1000}, 499500},
		{"smallCalls", []Value{1000}, 500500},
		{"recursive", []Value{10}, 55},
		{"concurrent", []Value{4, 250}, 125500},
		{"singleLine", []Value{1000}, 499500},
	} {
		b.Run(workload.name, func(b *testing.B) {
			var hostChanges uint64
			for _, mode := range []string{"baseline", "debug_only", "host_stack", "execution_tracker", "execution_repeated"} {
				b.Run(mode, func(b *testing.B) {
					changes := benchmarkHighlightCase(b, mode, workload.name, workload.args, workload.want)
					if mode == "host_stack" {
						hostChanges = changes
					} else if (mode == "execution_tracker" || mode == "execution_repeated") && hostChanges != 0 && changes != hostChanges {
						b.Fatalf("%s produced %d highlight changes; host stack produced %d", mode, changes, hostChanges)
					}
				})
			}
		})
	}
}

// Returns the warmup count for comparing highlighting modes.
func benchmarkHighlightCase(b *testing.B, mode, name string, args []Value, want int) uint64 {
	b.Helper()
	ctx := NewContext(0)
	ctx.SetLeastCallForEnablePool(0)
	tracked := mode == "execution_tracker" || mode == "execution_repeated"
	h := &benchmarkHighlights{
		positions:    make(map[uint64]token.Position),
		stacks:       make(map[uint64][]token.Position),
		lastSequence: make(map[uint64]uint64),
	}
	switch mode {
	case "debug_only":
		ctx.SetDebug(func(*DebugInfo) {})
	case "host_stack":
		h.attachHostStack(ctx)
	case "execution_tracker", "execution_repeated":
		tracker := NewExecutionTracker(func(event ExecutionPosition) {
			h.mu.Lock()
			h.events++
			if event.Sequence > h.lastSequence[event.GoroutineID] {
				h.lastSequence[event.GoroutineID] = event.Sequence
				h.update(event.GoroutineID, event.Position, event.Clear)
			}
			h.mu.Unlock()
		}, ExecutionOptions{ReportRepeated: mode == "execution_repeated"})
		b.Cleanup(tracker.Close)
		ctx.SetExecutionTracker(tracker)
	}
	interp, err := ctx.LoadInterp("execution_bench.go", highlightBenchSource)
	if err != nil {
		b.Fatal(err)
	}
	b.Cleanup(interp.UnsafeRelease)
	run := func() {
		got, err := interp.RunFunc(name, args...)
		// A result send can precede leave hooks. Join workers in every mode.
		if name == "concurrent" {
			for atomic.LoadInt32(&interp.goroutines) != 1 {
				runtime.Gosched()
			}
		}
		if err != nil || got != want {
			b.Fatalf("got %v, %v; want %d", got, err, want)
		}
		if tracked {
			// All callbacks have returned; release exited workers' watermarks.
			h.mu.Lock()
			clear(h.lastSequence)
			h.mu.Unlock()
		}
	}
	// Compilation, pool warmup, and initial validation are untimed.
	run()
	h.checkEmpty(b)
	changes := h.changes
	h.changes = 0
	h.events = 0
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		run()
	}
	b.StopTimer()
	h.checkEmpty(b)
	if mode == "host_stack" || tracked {
		b.ReportMetric(float64(h.events)/float64(b.N), "events/op")
		b.ReportMetric(float64(h.changes)/float64(b.N), "highlights/op")
	}
	return changes
}

type benchmarkHighlights struct {
	mu           sync.Mutex
	positions    map[uint64]token.Position
	stacks       map[uint64][]token.Position
	lastSequence map[uint64]uint64
	events       uint64
	changes      uint64
}

// update requires h.mu.
func (h *benchmarkHighlights) update(gid uint64, pos token.Position, clear bool) {
	previous, exists := h.positions[gid]
	if clear {
		if exists {
			delete(h.positions, gid)
			h.changes++
		}
		return
	}
	if !exists || previous.Filename != pos.Filename || previous.Line != pos.Line {
		h.positions[gid] = pos
		h.changes++
	}
}

func (h *benchmarkHighlights) attachHostStack(ctx *Context) {
	ctx.SetFuncHook(func(*FuncInfo) {
		gid := GoroutineID()
		h.mu.Lock()
		h.stacks[gid] = append(h.stacks[gid], token.Position{})
		h.mu.Unlock()
	}, func(*FuncInfo) {
		gid := GoroutineID()
		h.mu.Lock()
		defer h.mu.Unlock()
		h.events++ // Each leave submits a restored position or clear.
		stack := h.stacks[gid]
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			delete(h.stacks, gid)
		} else {
			h.stacks[gid] = stack
		}
		// A new function without a position keeps its caller's location.
		for i := len(stack) - 1; i >= 0; i-- {
			if stack[i].IsValid() {
				h.update(gid, stack[i], false)
				return
			}
		}
		h.update(gid, token.Position{}, true)
	})
	ctx.SetDebug(func(info *DebugInfo) {
		pos := info.Position()
		if !pos.IsValid() {
			return
		}
		gid := GoroutineID()
		h.mu.Lock()
		h.events++
		stack := h.stacks[gid]
		stack[len(stack)-1] = pos
		h.update(gid, pos, false)
		h.mu.Unlock()
	})
}

func (h *benchmarkHighlights) checkEmpty(b *testing.B) {
	b.Helper()
	h.mu.Lock()
	defer h.mu.Unlock()
	if len(h.positions) != 0 || len(h.stacks) != 0 {
		b.Fatalf("stale highlights after execution: %d positions, %d stacks", len(h.positions), len(h.stacks))
	}
}

const highlightBenchSource = `package main

func loop(n int) int {
	s := 0
	for i := 0; i < n; i++ {
		s += i
	}
	return s
}

func add(n int) int {
	return n + 1
}

func smallCalls(n int) int {
	s := 0
	for i := 0; i < n; i++ {
		s += add(i)
	}
	return s
}

func recursive(n int) int {
	if n < 2 {
		return n
	}
	return recursive(n - 1) + recursive(n - 2)
}

func concurrent(workers, n int) int {
	results := make(chan int, workers)
	for worker := 0; worker < workers; worker++ {
		go func() {
			results <- smallCalls(n)
		}()
	}
	s := 0
	for worker := 0; worker < workers; worker++ {
		s += <-results
	}
	return s
}

func singleLine(n int) int { s := 0; for i := 0; i < n; i++ { s += i }; return s }
`

func TestExecutionStackHostCallbacks(t *testing.T) {
	const source = `package main
//line Sprite.spx:10
func Root() {
 Host(func() { Child() })
}
func Host(callback func()) {}
//line Stage.spx:20
func Child() { Check() }
func Check() {}
`
	ctx := NewContext(SupportMultipleInterp | EnableCachedReg)
	var latest ExecutionPosition
	var saved ExecutionPosition
	tracker := NewExecutionTracker(func(event ExecutionPosition) { latest = event }, ExecutionOptions{
		TrackStack: true,
		Filter:     func(pos token.Position) bool { return strings.HasSuffix(pos.Filename, ".spx") },
	})
	defer tracker.Close()
	ctx.SetExecutionTracker(tracker)
	ctx.RegisterExternal("main.Host", func(callback func()) {
		if latest.Position.Filename != "Sprite.spx" || latest.Position.Line != 11 {
			t.Fatalf("host position = %v", latest)
		}
		callback()
		if latest.Position.Filename != "Sprite.spx" || latest.Position.Line != 11 {
			t.Fatalf("callback return = %v", latest)
		}
	})
	ctx.RegisterExternal("main.Check", func() {
		saved = latest
		var got []string
		for _, pos := range latest.Stack {
			got = append(got, pos.Filename)
		}
		if !reflect.DeepEqual(got, []string{"Stage.spx", "Sprite.spx", "Sprite.spx"}) {
			t.Fatalf("stack files = %v", got)
		}
		if latest.Stack[0].Line != 20 || latest.Stack[1].Line != 11 || latest.Stack[2].Line != 11 {
			t.Fatalf("stack = %v", latest.Stack)
		}
	})
	interp := loadExecution(t, ctx, source)
	for range 2 {
		runExecution(t, interp, "Root", nil)
		if !latest.Clear || len(latest.Stack) != 0 {
			t.Fatalf("root return = %v", latest)
		}
		// Stack snapshots must survive frame pooling and subsequent invocations.
		if len(saved.Stack) != 3 || saved.Stack[0].Filename != "Stage.spx" {
			t.Fatalf("retained stack = %v", saved.Stack)
		}
	}
}

func TestExecutionStackSameLineRecursion(t *testing.T) {
	const source = `package main
func Root(n int, check func(int)) { check(n); if n > 0 { Root(n-1, check) } }
`
	ctx := NewContext(EnableCachedReg)
	ctx.SetLeastCallForEnablePool(0)
	var latest ExecutionPosition
	var snapshots []ExecutionPosition
	tracker := NewExecutionTracker(func(event ExecutionPosition) { latest = event; snapshots = append(snapshots, event) }, ExecutionOptions{TrackStack: true})
	defer tracker.Close()
	ctx.SetExecutionTracker(tracker)
	interp := loadExecution(t, ctx, source)
	for range 2 {
		runExecution(t, interp, "Root", nil, 3, func(n int) {
			if len(latest.Stack) != 4-n {
				t.Fatalf("depth at n=%d: %v", n, latest.Stack)
			}
		})
		if !latest.Clear {
			t.Fatal("root did not clear")
		}
	}
	if len(snapshots[0].Stack) != 1 {
		t.Fatal("later events mutated an earlier snapshot")
	}
}

func TestExecutionStackWaitsForOwnPosition(t *testing.T) {
	const source = `package main
//line Sprite.spx:10
func Root() { Host(Empty, Bridge) }
func Host(empty func(), bridge func()) {}
//line generated.go:1
func Empty() {}
func Bridge() { Child() }
//line Sprite.spx:20
func Child() { Check() }
func Check() {}
`
	ctx := NewContext(SupportMultipleInterp | EnableCachedReg)
	var latest ExecutionPosition
	tracker := NewExecutionTracker(func(event ExecutionPosition) { latest = event }, ExecutionOptions{
		TrackStack: true,
		Filter:     func(pos token.Position) bool { return strings.HasSuffix(pos.Filename, ".spx") },
	})
	defer tracker.Close()
	ctx.SetExecutionTracker(tracker)
	checked := false
	ctx.RegisterExternal("main.Host", func(empty func(), bridge func()) {
		before := latest
		if before.Clear || len(before.Stack) != 1 {
			t.Fatalf("root position = %v", before)
		}
		empty()
		if latest.Sequence != before.Sequence {
			t.Fatalf("unlocated empty frame reported a change: %v", latest)
		}
		bridge()
		if latest.Clear || !sameLine(latest.Position, before.Position) || len(latest.Stack) != 1 {
			t.Fatalf("return through unlocated frame lost the caller: %v", latest)
		}
	})
	ctx.RegisterExternal("main.Check", func() {
		checked = true
		if len(latest.Stack) != 2 || latest.Stack[0].Line != 20 || latest.Stack[1].Line != 10 {
			t.Fatalf("expected only located child and root frames: %v", latest.Stack)
		}
	})
	runExecution(t, loadExecution(t, ctx, source), "Root", nil)
	if !checked || !latest.Clear || len(latest.Stack) != 0 {
		t.Fatalf("checked=%v, final event=%v", checked, latest)
	}
}

func TestExecutionDescriptionWithoutCompilation(t *testing.T) {
	ctx := NewContext(DisableMemo | ExperimentalSupportGC | EnableGCExecutionPlan)
	ctx.SetLeastCallForEnablePool(0)
	interp, err := ctx.LoadInterp("execution_description.go", `package main
import "runtime"
type counter int
func (c counter) Value() int { return int(c) }
func method() int { var v interface { Value() int } = counter(42); return v.Value() }
func async() int {
 ch := make(chan int, 1)
 go func(n int) { defer close(ch); ch <- n }(42)
 return <-ch
}
func collect() int { p := new(int); *p = 42; runtime.GC(); return *p }
func pair(n int) (int, int) { return n, n+1 }
func sum(n int) int {
 a, b := 0, 1
 for i := 0; i < n; i++ { a, b = b, a+b }
 x, y := pair(a)
 return x+y
}
func recovered() (n int) {
 defer func() { recover(); n = 42 }()
 panic("test")
}
func captured(n int) int { f := func() int { return n+1 }; return f() }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	functions := make(map[string]*function)
	gcBefore := make(map[*function][][]register)
	for fn, p := range interp.funcs {
		gcBefore[p] = make([][]register, len(p.gcRegs))
		for pc := range p.gcRegs {
			gcBefore[p][pc] = slices.Clone(p.gcRegsForPC(pc))
		}
		functions[fn.Name()] = p
		p.functionCompilation = functionCompilation{}
	}
	for _, tc := range []struct {
		name string
		args []value
		want value
	}{
		{"sum", []value{10}, 111},
		{"recovered", nil, 42},
		{"captured", []value{41}, 42},
		{"method", nil, 42},
		{"async", nil, 42},
		{"collect", nil, 42},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := functions[tc.name]
			if p == nil {
				t.Fatal("missing function")
			}
			for i := 0; i < 3; i++ {
				got := interp.callFunction(&frame{}, p, tc.args, nil)
				if !reflect.DeepEqual(got, tc.want) {
					t.Fatalf("got %v, want %v", got, tc.want)
				}
			}
			hasPosition := false
			for pc := range p.Instrs {
				pos := p.PositionForPC(pc)
				hasPosition = hasPosition || pos.IsValid()
			}
			if !hasPosition {
				t.Fatal("missing source metadata")
			}
			if got := runtimeFunc(p).Name(); got == "" {
				t.Fatal("missing runtime function name")
			}
			for pc := range p.gcRegs {
				if got := p.gcRegsForPC(pc); !slices.Equal(got, gcBefore[p][pc]) {
					t.Fatalf("GC registers at PC %d = %v, want %v", pc, got, gcBefore[p][pc])
				}
			}
		})
	}
}

func TestExecutionDescriptionPositionBounds(t *testing.T) {
	p := &function{}
	for _, pc := range []int{-1, 0, 1} {
		pos := p.PositionForPC(pc)
		if p.PosForPC(pc).IsValid() || pos.IsValid() {
			t.Fatalf("unexpected position at %d", pc)
		}
	}
}

func TestExecutionMetadataConcurrent(t *testing.T) {
	interp, err := NewContext(0).LoadInterp("metadata.go", `package main
func outer() int { f := func() int { return 42 }; return f() }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, p := range interp.funcs {
		if p.metadata.runtimeName != "" {
			t.Fatal("runtime name was eagerly formatted")
		}
	}
	type metadataExpectation struct {
		p         *function
		name      string
		autogen   bool
		positions []token.Position
	}
	var expected []metadataExpectation
	for _, p := range interp.funcs {
		name, autogen := formatRuntimeFuncName(p.Fn.String(), isInterfaceThunk(p.Fn.Signature))
		positions := make([]token.Position, len(p.ssaInstrs))
		for pc := range positions {
			positions[pc] = p.Fn.Prog.Fset.Position(p.PosForPC(pc))
		}
		p.functionCompilation = functionCompilation{}
		expected = append(expected, metadataExpectation{p, name, autogen, positions})
	}
	var wg sync.WaitGroup
	for range 16 {
		wg.Go(func() {
			for _, want := range expected {
				if got := runtimeFunc(want.p).Name(); got != want.name {
					t.Errorf("runtime name = %q, want %q", got, want.name)
				}
				if want.p.metadata.autogenerated != want.autogen {
					t.Error("autogenerated flag changed")
				}
				for pc, pos := range want.positions {
					if got := want.p.PositionForPC(pc); got != pos {
						t.Errorf("position at PC %d = %v, want %v", pc, got, pos)
					}
				}
			}
		})
	}
	wg.Wait()
}

func TestRuntimeFunctionNameFormatting(t *testing.T) {
	for _, tc := range []struct {
		name    string
		thunk   bool
		want    string
		autogen bool
	}{
		{"main.outer$1", false, "main.outer.func1", false},
		{"(*main.T).Value$bound", false, "main.(*T).Value-fm", true},
		{"main.Value$thunk", true, "main.Value", true},
		{"main.Value$thunk", false, "main.Value", false},
		{"struct{ X int }.Value$thunk", false, "struct{ X int }.Value", true},
	} {
		name, autogen := formatRuntimeFuncName(tc.name, tc.thunk)
		if name != tc.want || autogen != tc.autogen {
			t.Errorf("%q: got %q/%t, want %q/%t", tc.name, name, autogen, tc.want, tc.autogen)
		}
	}
}

func TestRootFrameConstructor(t *testing.T) {
	fr := rootFrame()
	if fr.block != -1 || fr.valid() {
		t.Fatalf("root frame = block %d, valid %v; want -1, false", fr.block, fr.valid())
	}
}

func TestRootFrameCallersAndGC(t *testing.T) {
	const source = `package main

import "runtime"

var gotNames []string

func work(done chan bool) {
	pcs := make([]uintptr, 16)
	n := runtime.Callers(1, pcs)
	if n == 0 {
		panic("no callers")
	}
	frames := runtime.CallersFrames(pcs[:n])
	names := make([]string, 0, n)
	for {
		f, more := frames.Next()
		if f.Function != "" {
			names = append(names, f.Function)
		}
		if !more {
			break
		}
	}
	runtime.GC()
	gotNames = names
	done <- true
}
func main() { done := make(chan bool, 1); go work(done); <-done }
`
	for _, tc := range []struct {
		name string
		mode Mode
	}{
		{"gc", ExperimentalSupportGC},
		{"gc_cached_registers", ExperimentalSupportGC | EnableCachedReg},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := NewContext(tc.mode)
			interp, err := ctx.LoadInterp("root_callers.go", source)
			if err != nil {
				t.Fatal(err)
			}
			defer interp.UnsafeRelease()
			if _, err := interp.RunFunc("main"); err != nil {
				t.Fatal(err)
			}
			addr, ok := interp.GetVarAddr("gotNames")
			if !ok {
				t.Fatal("missing gotNames global")
			}
			names := *addr.(*[]string)
			workCount := 0
			for _, name := range names {
				if name == "main.work" {
					workCount++
				}
			}
			if workCount != 1 {
				t.Fatalf("call stack = %v; want exactly one main.work frame", names)
			}
		})
	}
}

func TestExecutionDescriptionFallbacks(t *testing.T) {
	interp := loadExecution(t, NewContext(0), "package main\nfunc Root() {}\n")
	for _, tc := range []struct {
		name      string
		wantPanic bool
	}{
		{"init", false},
		{"missing", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fn := interp.mainpkg.Prog.NewFunction(tc.name, types.NewSignatureType(nil, nil, nil, nil, nil, false), "test")
			fn.Pkg = interp.mainpkg
			call := &ssa.CallCommon{Value: fn}
			if tc.wantPanic {
				defer func() {
					if got := recover(); got == nil || !strings.Contains(fmt.Sprint(got), "no code for function") {
						t.Fatalf("panic = %v", got)
					}
				}()
			}
			desc := interp.describeCall(call, 0, nil, nil)
			fnValue, args := desc.prepare(interp, &frame{})
			if got := interp.call(rootFrame(), fnValue, args, nil); got != nil {
				t.Fatalf("init result = %v", got)
			}
		})
	}

	for _, withMethods := range []bool{false, true} {
		t.Run(fmt.Sprintf("methods=%t", withMethods), func(t *testing.T) {
			v := time.Second
			typ := reflect.TypeOf(v)
			if withMethods {
				interp.msets[typ] = map[string]*function{}
			} else {
				delete(interp.msets, typ)
			}
			fr := &frame{stack: []value{v}}
			desc := callDescription{receiver: 0, method: "String"}
			fn, args := desc.prepare(interp, fr)
			if got := interp.call(rootFrame(), fn, args, nil); got != "1s" {
				t.Fatalf("String() = %v", got)
			}
			desc.method = "Missing"
			defer func() {
				if got := recover(); got == nil || !strings.Contains(fmt.Sprint(got), "no code for method") {
					t.Fatalf("panic = %v", got)
				}
			}()
			desc.prepare(interp, fr)
		})
	}
}

func TestExecutionBuiltinDescription(t *testing.T) {
	b := &builtinDescription{name: "println"}
	if b.Name() != "println" || b.String() != "println" || b.Pos() != token.NoPos || b.Parent() != nil {
		t.Fatalf("unexpected builtin description: %v", b)
	}
	interp := loadExecution(t, NewContext(0), "package main\nfunc Root() {}\n")
	if got := interp.call(rootFrame(), b, []value{42}, []bool{true}); got != nil {
		t.Fatalf("println result = %v", got)
	}
}

func TestExecutionMetadataNilInstruction(t *testing.T) {
	interp := loadExecution(t, NewContext(0), "package main\nfunc Root() {}\n")
	p := interp.funcs[interp.mainpkg.Func("Root")]
	p.ssaInstrs = append(p.ssaInstrs, nil)
	p.finishMetadata()
	if got := p.PosForPC(len(p.ssaInstrs) - 1); got != token.NoPos {
		t.Fatalf("nil instruction position = %v", got)
	}
}

func TestGCLivenessInvalidQueries(t *testing.T) {
	rootFrame().gc()
	var plan *gcLivenessPlan
	if got := plan.regsAt(0); got != nil {
		t.Fatalf("nil plan registers = %v", got)
	}
	plan = &gcLivenessPlan{points: []gcLivenessPoint{{block: -1}}}
	for _, pc := range []int{-1, 0, 1} {
		if got := plan.regsAt(pc); got != nil {
			t.Fatalf("registers at %d = %v", pc, got)
		}
	}
}

// Mutating the original SSA operands after compilation detects accidental
// runtime reads that clearing functionCompilation alone would leave reachable.
func TestRuntimeInstructionOperands(t *testing.T) {
	const source = `package main
 type numbers []int
 type counter int
 func (c counter) Value() int { return int(c) }
 func selects() int {
  a, b := make(chan int, 1), make(chan int)
  a <- 42
  select { case v := <-a: return v; case <-b: return 0 }
 }
 func defaultSelect() int { var c chan int; select { case <-c: return 0; default: return 42 } }
 func closedSelect() int { c := make(chan *int); close(c); select { case v, ok := <-c: if v == nil && !ok { return 42 }; return 0 } }
 func sendNilSelect() int {
  c := make(chan *int, 1)
  select { case c <- nil: default: return 0 }
  if <-c == nil { return 42 }; return 0
 }
 func sendClosedSelect() int { c := make(chan int); close(c); select { case c <- 1: return 0; default: return 42 } }
 func nilInterfaceSelect() int {
  c := make(chan interface{}, 1)
  select { case c <- nil: default: return 0 }
  if <-c == nil { return 42 }; return 0
 }
 func slices() int { a := numbers{10, 20, 30, 40}; b := a[1:3:4]; return b[0]+len(b)+cap(b)+17 }
 func omittedBounds() int { a := [3]int{40, 1, 1}; b := a[:]; return b[0]+len(b[1:]) }
 func stringSlice() int { s := "abc"; if s[1:] == "bc" && s[:2] == "ab" && s[3:] == "" { return 42 }; return 0 }
 func slicePanic() int { a := []int{1}; n := 2; return len(a[:n]) }
 func slice3Panic() int { a := []int{1}; n := 2; return len(a[:1:n]) }
 func negativeSlice() int { a := []int{1}; n := -1; return len(a[n:]) }
 func makeSlicePanic() int { n := -1; return len(make([]int,n)) }
 func assertValue() int { var x interface{} = 42; return x.(int) }
 func assertCommaOK() int { var x interface{} = "wrong"; v, ok := x.(int); if !ok && v == 0 { return 42 }; return 0 }
 func assertNil() int { var x interface{}; _, ok := x.(int); if !ok { return 42 }; return 0 }
 func assertTypedNil() int { var x interface{} = (*int)(nil); p, ok := x.(*int); if ok && p == nil { return 42 }; return 0 }
 func assertInterface() int { var x interface{} = counter(42); return x.(interface{ Value() int }).Value() }
 func assertPanic() int { var x interface{} = "wrong"; return x.(int) }
 func assertNilPanic() int { var x interface{}; return x.(int) }
 func assertMethodPanic() int { var x interface{} = 42; return x.(interface{ Value() int }).Value() }
 `
	testGCLivenessModes(t, func(t *testing.T, mode Mode) {
		ctx := NewContext(DisableMemo | ExperimentalSupportGC | mode)
		interp, err := ctx.LoadInterp("runtime_operands.go", source)
		if err != nil {
			t.Fatal(err)
		}
		defer interp.UnsafeRelease()
		for _, p := range interp.funcs {
			for _, instr := range p.ssaInstrs {
				switch v := instr.(type) {
				case *ssa.Select:
					v.Blocking = !v.Blocking
					v.States = nil
				case *ssa.Slice:
					v.Low, v.High, v.Max = nil, nil, nil
				case *ssa.TypeAssert:
					v.CommaOk = !v.CommaOk
					v.AssertedType = nil
				}
			}
			if mode != 0 {
				p.functionCompilation = functionCompilation{}
			}
		}
		for _, name := range []string{"selects", "defaultSelect", "closedSelect", "sendNilSelect", "nilInterfaceSelect", "slices", "omittedBounds", "stringSlice", "assertValue", "assertCommaOK", "assertNil", "assertTypedNil", "assertInterface"} {
			t.Run(name, func(t *testing.T) {
				got, err := interp.RunFunc(name)
				if err != nil || got != 42 {
					t.Fatalf("got %v, %v; want 42", got, err)
				}
			})
		}
		for _, tc := range []struct{ name, message string }{
			{"sendClosedSelect", "send on closed channel"}, {"slicePanic", "slice bounds out of range"}, {"slice3Panic", "slice bounds out of range"}, {"negativeSlice", "slice bounds out of range"}, {"makeSlicePanic", "makeslice: len out of range"},
			{"assertPanic", "interface conversion"}, {"assertNilPanic", "interface conversion"}, {"assertMethodPanic", "missing method Value"},
		} {
			t.Run(tc.name, func(t *testing.T) {
				_, err := interp.RunFunc(tc.name)
				if err == nil || !strings.Contains(err.Error(), tc.message) {
					t.Fatalf("got %v; want %q", err, tc.message)
				}
			})
		}
	})
}

func BenchmarkRuntimeInstructionOperands(b *testing.B) {
	const source = `package main
func slices() int { a := []int{10,20,30,40}; n:=0; for i:=0;i<100;i++ { n+=len(a[1:3:4]) }; return n }
func assertions() int { var x interface{} = 2; n:=0; for i:=0;i<100;i++ { v,ok:=x.(int);if ok { n+=v } };return n }
func selects() int { a:=make(chan int,1);n:=0;for i:=0;i<100;i++ { a<-2;select {case v:=<-a:n+=v;default:} };return n }
`
	for _, name := range []string{"slices", "assertions", "selects"} {
		b.Run(name, func(b *testing.B) {
			interp, err := NewContext(DisableMemo).LoadInterp("operand_bench.go", source)
			if err != nil {
				b.Fatal(err)
			}
			b.Cleanup(interp.UnsafeRelease)
			got, err := interp.RunFunc(name)
			if err != nil || got != 200 {
				b.Fatalf("got %v, %v; want 200", got, err)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				if _, err := interp.RunFunc(name); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func TestPhiDirectAssignmentDependencies(t *testing.T) {
	for _, c := range []struct {
		results []register
		edges   [][]register
		want    bool
	}{
		{nil, nil, false}, {[]register{0}, nil, false},
		{[]register{0}, [][]register{{0, 2}}, true},
		{[]register{0, 1}, [][]register{{2, 3}, {3, 2}}, true},
		{[]register{0, 1}, [][]register{{1}, {2}}, true},
		{[]register{0, 1}, [][]register{{2}, {0}}, false},
		{[]register{0, 1}, [][]register{{1}, {0}}, false},
		{[]register{0, 1}, [][]register{{0}, {0}}, true},
		{[]register{0, 1}, [][]register{{2, 1}, {3, 0}}, false},
	} {
		if got := phiDirectAssignment(c.results, c.edges); got != c.want {
			t.Fatalf("results=%v edges=%v: %t want %t", c.results, c.edges, got, c.want)
		}
	}
}
