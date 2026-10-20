/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package ixgo

import "testing"

func TestComposeReceiverParentMissingReceiverState(t *testing.T) {
	for _, tc := range []struct {
		name  string
		state *interpreterMemoState
	}{
		{"empty", &interpreterMemoState{}},
		{"scalar", &interpreterMemoState{memo: &interpreterMemo{}}},
		{"accumulator", &interpreterMemoState{accumulator: &interpreterAccumulatorState{}}},
	} {
		for _, missingParent := range []bool{true, false} {
			side := "child"
			if missingParent {
				side = "parent"
			}
			t.Run(tc.name+"/"+side, func(t *testing.T) {
				receiver := &interpreterMemoState{receiver: &interpreterReceiverState{
					memo:   &interpreterReceiverMemo{},
					affine: []receiverTransform{{a: 2, b: 3}},
				}}
				parent := &frame{memoState: receiver}
				child := &frame{caller: parent, memoState: tc.state}
				if missingParent {
					parent.memoState, child.memoState = child.memoState, parent.memoState
				}
				composeReceiverParent(child, nil)
				if got := receiver.receiver.affine[0]; got != (receiverTransform{a: 2, b: 3}) {
					t.Fatalf("transform changed without receiver state: %#v", got)
				}
			})
		}
	}
}
