/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
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

// Package unsaferewrite preserves unevaluated unsafe layout operands when
// building SSA. Isolate caller-owned syntax before type checking, then Rewrite
// it with matching type information before constructing SSA packages.
//
// Rewritten calls carry a typed nil pointer followed by a zero marker for
// Sizeof/Alignof or field indices for Offsetof. The instruction compiler must
// decode this internal representation using instantiated Go types.
package unsaferewrite

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"strconv"
)

func importsUnsafe(file *ast.File) bool {
	for _, spec := range file.Imports {
		if path, _ := strconv.Unquote(spec.Path.Value); path == "unsafe" && (spec.Name == nil || spec.Name.Name != "_") {
			return true
		}
	}
	return false
}

// Isolate copies candidate calls and their ancestors in a caller-owned file.
// Other subtrees remain shared. The result is private only for mutations by
// Rewrite; it is not a general-purpose AST clone.
func Isolate(file *ast.File) *ast.File {
	if !importsUnsafe(file) {
		return file
	}
	// Only candidate files need
	// traversal stacks and a replacement table.
	for n := range ast.Preorder(file) {
		if !unsafeLayoutCall(n) {
			continue
		}
		paths := make(map[ast.Node]ast.Node)
		ast.PreorderStack(file, nil, func(n ast.Node, stack []ast.Node) bool {
			if unsafeLayoutCall(n) {
				paths[n] = nil
				for i := len(stack) - 1; i >= 0; i-- {
					if _, marked := paths[stack[i]]; marked {
						break
					}
					paths[stack[i]] = nil
				}
			}
			return true
		})
		return copyLayoutPath(file, paths).(*ast.File)
	}
	return file
}

func unsafeLayoutCall(n ast.Node) bool {
	call, ok := n.(*ast.CallExpr)
	if !ok || len(call.Args) != 1 {
		return false
	}
	switch fun := ast.Unparen(call.Fun).(type) {
	case *ast.SelectorExpr:
		return isUnsafeLayoutName(fun.Sel.Name)
	case *ast.Ident:
		return isUnsafeLayoutName(fun.Name)
	}
	return false
}

func isUnsafeLayoutName(name string) bool {
	return name == "Sizeof" || name == "Alignof" || name == "Offsetof"
}

// Shallow-copy marked nodes and detach a child slice on its first replacement.
// The rewrite replaces CallExpr.Fun and Args wholesale, so the original
// operands, like other unmodified subtrees, can remain shared.
func copyLayoutPath(n ast.Node, paths map[ast.Node]ast.Node) ast.Node {
	if c := paths[n]; c != nil {
		return c
	}
	orig := reflect.ValueOf(n).Elem()
	ptr := reflect.New(orig.Type())
	dst := ptr.Elem()
	dst.Set(orig)
	c := ptr.Interface().(ast.Node)
	paths[n] = c
	child := func(v reflect.Value) ast.Node {
		node, ok := v.Interface().(ast.Node)
		if !ok {
			return nil
		}
		if _, marked := paths[node]; marked {
			return copyLayoutPath(node, paths)
		}
		return nil
	}
	for i := 0; i < dst.NumField(); i++ {
		f := dst.Field(i)
		switch f.Kind() {
		case reflect.Interface, reflect.Pointer:
			if repl := child(f); repl != nil {
				f.Set(reflect.ValueOf(repl))
			}
		case reflect.Slice:
			detached := false
			for j := 0; j < f.Len(); j++ {
				if repl := child(f.Index(j)); repl != nil {
					if !detached {
						dup := reflect.MakeSlice(f.Type(), f.Len(), f.Len())
						reflect.Copy(dup, f)
						f.Set(dup)
						detached = true
					}
					f.Index(j).Set(reflect.ValueOf(repl))
				}
			}
		}
	}
	return c
}

// Rewrite replaces nonconstant Sizeof, Alignof, and Offsetof
// operands after type checking and before SSA construction. Internal calls carry
// a typed nil pointer and constant metadata instead of the unevaluated expression.
func Rewrite(files []*ast.File, info *types.Info) {
	if info == nil {
		return
	}
	for _, file := range files {
		if !importsUnsafe(file) {
			continue
		}
		ast.Inspect(file, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok || len(call.Args) != 1 || info.Types[call].Value != nil {
				return true
			}
			var id *ast.Ident
			switch fun := ast.Unparen(call.Fun).(type) {
			case *ast.SelectorExpr:
				id = fun.Sel
			case *ast.Ident:
				id = fun
			}
			builtin, ok := info.Uses[id].(*types.Builtin)
			if !ok || builtin.Pkg() != types.Unsafe {
				return true
			}
			name := builtin.Name()
			if !isUnsafeLayoutName(name) {
				return true
			}
			operand := call.Args[0]
			receiver := info.TypeOf(operand)
			// Sizeof/Alignof use a zero marker to distinguish this internal
			// type-only representation from ordinary one-argument SSA.
			indices := []int{0}
			if name == "Offsetof" {
				selector, ok := ast.Unparen(operand).(*ast.SelectorExpr)
				if !ok {
					return true
				}
				selection := info.Selections[selector]
				if selection == nil || selection.Kind() != types.FieldVal {
					return true
				}
				receiver = selection.Recv()
				if ptr, ok := receiver.Underlying().(*types.Pointer); ok {
					receiver = ptr.Elem()
				}
				indices = selection.Index()
			}
			ptr := types.NewPointer(receiver)
			placeholder := &ast.Ident{NamePos: operand.Pos(), Name: "nil"}
			info.Uses[placeholder] = types.Universe.Lookup("nil")
			info.Types[placeholder] = types.TypeAndValue{Type: ptr}
			args := make([]ast.Expr, 1, 1+len(indices))
			args[0] = placeholder
			params := make([]*types.Var, 1, 1+len(indices))
			params[0] = types.NewParam(token.NoPos, nil, "", ptr)
			for _, index := range indices {
				arg := &ast.BasicLit{ValuePos: operand.Pos(), Kind: token.INT, Value: strconv.Itoa(index)}
				info.Types[arg] = types.TypeAndValue{Type: types.Typ[types.Uintptr], Value: constant.MakeInt64(int64(index))}
				args = append(args, arg)
				params = append(params, types.NewParam(token.NoPos, nil, "", types.Typ[types.Uintptr]))
			}
			fun := &ast.Ident{NamePos: call.Fun.Pos(), Name: name}
			info.Uses[fun] = builtin
			info.Types[fun] = types.TypeAndValue{Type: types.NewSignatureType(nil, nil, nil,
				types.NewTuple(params...), types.NewTuple(types.NewParam(token.NoPos, nil, "", types.Typ[types.Uintptr])), false)}
			call.Fun = fun
			call.Args = args
			return false
		})
	}
}
