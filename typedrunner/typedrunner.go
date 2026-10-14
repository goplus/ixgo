/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

// Package typedrunner enables optional scalar and memory typed runners.
//
// Import it for applications that want the optimization:
//
//	import _ "github.com/goplus/ixgo/typedrunner"
package typedrunner

import (
	"go/types"

	"github.com/goplus/ixgo"
	"golang.org/x/tools/go/ssa"
)

func makeTypedRunner(fn *ssa.Function) ixgo.TypedRunner {
	if fn == nil || fn.Signature == nil {
		return nil
	}
	// Select T from any integer or floating-point parameter or result. The core
	// validates the complete signature, body and call graph.
	kind := types.Int
	found := false
	selectKind := func(typ types.Type) {
		if found {
			return
		}
		if pointer, ok := typ.(*types.Pointer); ok {
			typ = pointer.Elem()
		}
		if k, ok := scalarKind(typ); ok {
			kind, found = k, true
		}
	}
	for i := 0; i < fn.Signature.Params().Len(); i++ {
		selectKind(fn.Signature.Params().At(i).Type())
	}
	for i := 0; i < fn.Signature.Results().Len(); i++ {
		selectKind(fn.Signature.Results().At(i).Type())
	}
	switch kind {
	case types.Int:
		return ixgo.MakeTypedRunnerOf[int](fn)
	case types.Int8:
		return ixgo.MakeTypedRunnerOf[int8](fn)
	case types.Int16:
		return ixgo.MakeTypedRunnerOf[int16](fn)
	case types.Int32:
		return ixgo.MakeTypedRunnerOf[int32](fn)
	case types.Int64:
		return ixgo.MakeTypedRunnerOf[int64](fn)
	case types.Uint:
		return ixgo.MakeTypedRunnerOf[uint](fn)
	case types.Uint8:
		return ixgo.MakeTypedRunnerOf[uint8](fn)
	case types.Uint16:
		return ixgo.MakeTypedRunnerOf[uint16](fn)
	case types.Uint32:
		return ixgo.MakeTypedRunnerOf[uint32](fn)
	case types.Uint64:
		return ixgo.MakeTypedRunnerOf[uint64](fn)
	case types.Uintptr:
		return ixgo.MakeTypedRunnerOf[uintptr](fn)
	case types.Float32:
		return ixgo.MakeTypedRunnerOf[float32](fn)
	case types.Float64:
		return ixgo.MakeTypedRunnerOf[float64](fn)
	default:
		return nil
	}
}

func scalarKind(typ types.Type) (types.BasicKind, bool) {
	b, ok := typ.Underlying().(*types.Basic)
	if !ok {
		return 0, false
	}
	switch b.Kind() {
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
		types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64, types.Uintptr,
		types.Float32, types.Float64:
		return b.Kind(), true
	default:
		return 0, false
	}
}

func init() {
	ixgo.RegisterTypedRunnerFactory(makeTypedRunner)
}
