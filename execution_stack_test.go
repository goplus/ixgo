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
	"reflect"
	"strings"
	"testing"
)

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
