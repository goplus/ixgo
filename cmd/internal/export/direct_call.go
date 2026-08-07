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

package export

import (
	"fmt"
	"go/token"
	"go/types"
	"sort"
	"strconv"
	"strings"
)

type directCallGenerator struct {
	usedNames map[string]bool
}

type directCallBinding struct {
	key           string
	name          string
	receiver      types.Type
	adapterName   string
	argumentTypes []types.Type
	variadic      bool
	hasResult     bool
}

type directCallRenderer struct {
	pkg       *types.Package
	pkgAlias  string
	addImport func(string, string) string
}

func (r *directCallRenderer) typeExpr(typ types.Type) string {
	typ = types.Unalias(typ)
	return types.TypeString(typ, func(pkg *types.Package) string {
		if pkg == r.pkg {
			return r.pkgAlias
		}
		return r.addImport(pkg.Path(), pkg.Name())
	})
}

func (r *directCallRenderer) targetExpr(binding directCallBinding) string {
	if binding.receiver == nil {
		return r.pkgAlias + "." + binding.name
	}
	receiver := r.typeExpr(binding.receiver)
	if _, ok := types.Unalias(binding.receiver).(*types.Pointer); ok {
		receiver = "(" + receiver + ")"
	}
	return receiver + "." + binding.name
}

func (r *directCallRenderer) registryEntry(binding directCallBinding) string {
	return fmt.Sprintf("%q: %s", binding.key, binding.adapterName)
}

func (r *directCallRenderer) adapterDeclaration(binding directCallBinding) string {
	args := make([]string, len(binding.argumentTypes))
	for i, typ := range binding.argumentTypes {
		args[i] = fmt.Sprintf("ixgo.DirectCallArg[%s](ctx, %d)", r.typeExpr(typ), i)
	}
	if binding.variadic {
		args[len(args)-1] += "..."
	}
	call := fmt.Sprintf("%s(%s)", r.targetExpr(binding), strings.Join(args, ", "))
	if binding.hasResult {
		call = "ctx.SetResult(" + call + ")"
	}
	return fmt.Sprintf("func %s(ctx ixgo.DirectCallContext) {\n\t%s\n}", binding.adapterName, call)
}

type directCallOutput struct {
	pkg      *types.Package
	bindings []directCallBinding
}

func (o directCallOutput) isEmpty() bool {
	return len(o.bindings) == 0
}

func (o directCallOutput) render(pkgAlias string, addImport func(string, string) string) (entries, declarations []string) {
	renderer := directCallRenderer{
		pkg:       o.pkg,
		pkgAlias:  pkgAlias,
		addImport: addImport,
	}
	entries = make([]string, len(o.bindings))
	declarations = make([]string, len(o.bindings))
	for i, binding := range o.bindings {
		entries[i] = renderer.registryEntry(binding)
		declarations[i] = renderer.adapterDeclaration(binding)
	}
	return entries, declarations
}

type directCallCandidate struct {
	selector string
	fn       *types.Func
	explicit bool
}

func newDirectCallGenerator() *directCallGenerator {
	return &directCallGenerator{usedNames: make(map[string]bool)}
}

func parseDirectCallSelectors(list string) ([]string, error) {
	selectors := strings.FieldsFunc(list, func(r rune) bool { return r == ',' || r == ';' })
	seen := make(map[string]bool, len(selectors))
	ret := selectors[:0]
	for _, selector := range selectors {
		selector = strings.TrimSpace(selector)
		if selector == "" {
			continue
		}
		if seen[selector] {
			return nil, fmt.Errorf("duplicate direct call %q", selector)
		}
		seen[selector] = true
		ret = append(ret, selector)
	}
	sort.Strings(ret)
	return ret, nil
}

func generateDirectCalls(pkg *types.Package, list string) (directCallOutput, error) {
	selectors, err := parseDirectCallSelectors(list)
	if err != nil {
		return directCallOutput{}, err
	}
	candidates, err := expandDirectCallSelectors(pkg, selectors)
	if err != nil {
		return directCallOutput{}, err
	}
	gen := newDirectCallGenerator()
	output := directCallOutput{pkg: pkg, bindings: make([]directCallBinding, 0, len(candidates))}
	for _, candidate := range candidates {
		if err := validateDirectCall(candidate.fn); err != nil {
			if !candidate.explicit {
				continue
			}
			return directCallOutput{}, fmt.Errorf("direct call %q: %w", candidate.selector, err)
		}
		output.bindings = append(output.bindings, gen.buildBindings(candidate.selector, candidate.fn)...)
	}
	return output, nil
}

func expandDirectCallSelectors(pkg *types.Package, selectors []string) ([]directCallCandidate, error) {
	candidates := make(map[string]directCallCandidate)
	addCandidate := func(selector string, fn *types.Func, explicit bool) {
		key := fn.FullName()
		selector = canonicalDirectCallSelector(fn, selector)
		if candidate, ok := candidates[key]; ok {
			candidate.explicit = candidate.explicit || explicit
			candidates[key] = candidate
			return
		}
		candidates[key] = directCallCandidate{selector: selector, fn: fn, explicit: explicit}
	}
	addAutomatic := func(selector string, fn *types.Func) {
		addCandidate(selector, fn, false)
	}
	for _, selector := range selectors {
		switch {
		case selector == "all":
			addDirectCallFunctions(addAutomatic, pkg)
			addDirectCallNamedTypeMethods(addAutomatic, pkg)
		case selector == "*":
			addDirectCallFunctions(addAutomatic, pkg)
		case selector == "*.*":
			addDirectCallNamedTypeMethods(addAutomatic, pkg)
		case strings.HasSuffix(selector, ".*"):
			typeName := strings.TrimSuffix(selector, ".*")
			named, err := resolveDirectCallReceiver(pkg, selector, typeName)
			if err != nil {
				return nil, err
			}
			addDirectCallMethods(addAutomatic, typeName, named)
		default:
			fn, err := resolveDirectCall(pkg, selector)
			if err != nil {
				return nil, err
			}
			addCandidate(selector, fn, true)
		}
	}
	ret := make([]directCallCandidate, 0, len(candidates))
	for _, candidate := range candidates {
		ret = append(ret, candidate)
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].selector < ret[j].selector })
	return ret, nil
}

func canonicalDirectCallSelector(fn *types.Func, fallback string) string {
	sig, ok := fn.Type().(*types.Signature)
	if !ok || sig.Recv() == nil {
		return fn.Name()
	}
	receiver := types.Unalias(sig.Recv().Type())
	if pointer, ok := receiver.(*types.Pointer); ok {
		receiver = types.Unalias(pointer.Elem())
	}
	if named, ok := receiver.(*types.Named); ok {
		return named.Obj().Name() + "." + fn.Name()
	}
	return fallback
}

func addDirectCallFunctions(add func(string, *types.Func), pkg *types.Package) {
	for _, name := range pkg.Scope().Names() {
		if fn, ok := pkg.Scope().Lookup(name).(*types.Func); ok && token.IsExported(name) {
			add(name, fn)
		}
	}
}

func addDirectCallNamedTypeMethods(add func(string, *types.Func), pkg *types.Package) {
	for _, name := range pkg.Scope().Names() {
		typeName, ok := pkg.Scope().Lookup(name).(*types.TypeName)
		if !ok || typeName.IsAlias() || !token.IsExported(name) {
			continue
		}
		named, ok := typeName.Type().(*types.Named)
		if !ok || named.Obj().Pkg() != pkg {
			continue
		}
		addDirectCallMethods(add, name, named)
	}
}

func addDirectCallMethods(add func(string, *types.Func), typeName string, named *types.Named) {
	if _, ok := named.Underlying().(*types.Interface); ok {
		return
	}
	for i := 0; i < named.NumMethods(); i++ {
		method := named.Method(i)
		if token.IsExported(method.Name()) {
			add(typeName+"."+method.Name(), method)
		}
	}
}

func validateDirectCall(fn *types.Func) error {
	sig, ok := fn.Type().(*types.Signature)
	if !ok {
		return fmt.Errorf("does not have a function signature")
	}
	if hasTypeParam(sig) || sig.RecvTypeParams() != nil {
		return fmt.Errorf("has type parameters")
	}
	if results := sig.Results().Len(); results > 1 {
		return fmt.Errorf("has %d results; at most one is supported", results)
	}
	if recv := sig.Recv(); recv != nil {
		if err := checkDirectCallType(types.Unalias(recv.Type())); err != nil {
			return fmt.Errorf("receiver: %w", err)
		}
	}
	for i := 0; i < sig.Params().Len(); i++ {
		if err := checkDirectCallType(types.Unalias(sig.Params().At(i).Type())); err != nil {
			return fmt.Errorf("argument %d: %w", i, err)
		}
	}
	return nil
}

func resolveDirectCall(pkg *types.Package, selector string) (*types.Func, error) {
	parts := strings.Split(selector, ".")
	switch len(parts) {
	case 1:
		obj := pkg.Scope().Lookup(parts[0])
		fn, ok := obj.(*types.Func)
		if !ok || !token.IsExported(parts[0]) {
			return nil, fmt.Errorf("direct call %q is not an exported package function", selector)
		}
		return fn, nil
	case 2:
		named, err := resolveDirectCallReceiver(pkg, selector, parts[0])
		if err != nil {
			return nil, err
		}
		for i := 0; i < named.NumMethods(); i++ {
			method := named.Method(i)
			if method.Name() == parts[1] && token.IsExported(method.Name()) {
				return method, nil
			}
		}
		return nil, fmt.Errorf("direct call %q is not an exported declared method", selector)
	default:
		return nil, fmt.Errorf("invalid direct call selector %q", selector)
	}
}

func resolveDirectCallReceiver(pkg *types.Package, selector, name string) (*types.Named, error) {
	obj := pkg.Scope().Lookup(name)
	typeName, ok := obj.(*types.TypeName)
	if !ok || !token.IsExported(name) {
		return nil, fmt.Errorf("direct call %q has no exported receiver type", selector)
	}
	named, ok := types.Unalias(typeName.Type()).(*types.Named)
	if !ok {
		return nil, fmt.Errorf("direct call %q receiver is not a defined type", selector)
	}
	if named.Obj().Pkg() != pkg {
		return nil, fmt.Errorf("direct call %q receiver is defined in another package", selector)
	}
	if _, ok := named.Underlying().(*types.Interface); ok {
		return nil, fmt.Errorf("direct call %q is an interface method", selector)
	}
	return named, nil
}

func (g *directCallGenerator) buildBindings(selector string, fn *types.Func) []directCallBinding {
	sig := fn.Type().(*types.Signature)
	receiver := types.Type(nil)
	if recv := sig.Recv(); recv != nil {
		receiver = recv.Type()
	}
	bindings := []directCallBinding{g.newBinding(selector, fn, receiver)}
	if receiver != nil {
		if _, pointer := types.Unalias(receiver).(*types.Pointer); !pointer {
			bindings = append(bindings, g.newBinding(selector, fn, types.NewPointer(receiver)))
		}
	}
	return bindings
}

func (g *directCallGenerator) newBinding(selector string, fn *types.Func, receiver types.Type) directCallBinding {
	sig := fn.Type().(*types.Signature)
	argumentTypes := make([]types.Type, 0, sig.Params().Len()+1)
	if receiver != nil {
		argumentTypes = append(argumentTypes, receiver)
	}
	for i := 0; i < sig.Params().Len(); i++ {
		argumentTypes = append(argumentTypes, sig.Params().At(i).Type())
	}
	key := fn.Name()
	if receiver != nil {
		key = directCallMethodKey(receiver, fn.Name())
	}
	return directCallBinding{
		key:           key,
		name:          fn.Name(),
		receiver:      receiver,
		adapterName:   g.adapterName(selector, receiver),
		argumentTypes: argumentTypes,
		variadic:      sig.Variadic(),
		hasResult:     sig.Results().Len() == 1,
	}
}

func (g *directCallGenerator) adapterName(selector string, receiver types.Type) string {
	var base string
	if receiver == nil {
		base = "func_" + selector
	} else {
		receiverName, name, _ := strings.Cut(selector, ".")
		if _, pointer := types.Unalias(receiver).(*types.Pointer); pointer {
			base = "method_ptr_" + receiverName + "_" + name
		} else {
			base = "method_" + receiverName + "_" + name
		}
	}
	name := base
	for suffix := 2; g.usedNames[name]; suffix++ {
		name = base + strconv.Itoa(suffix)
	}
	g.usedNames[name] = true
	return name
}

// directCallMethodKey returns the method selector format accepted by RegisterDirectCalls.
func directCallMethodKey(receiver types.Type, method string) string {
	name := types.TypeString(receiver, func(*types.Package) string { return "" })
	if _, pointer := types.Unalias(receiver).(*types.Pointer); pointer {
		return "(" + name + ")." + method
	}
	return "(" + name + ")." + method
}

func checkDirectCallType(typ types.Type) error {
	switch typ := typ.(type) {
	case *types.Alias:
		if obj := typ.Obj(); obj.Pkg() != nil && !obj.Exported() {
			return fmt.Errorf("type alias %s is not exported", obj)
		}
		return checkDirectCallType(types.Unalias(typ))
	case *types.Named:
		if obj := typ.Obj(); obj.Pkg() != nil && !obj.Exported() {
			return fmt.Errorf("type %s is not exported", obj)
		}
		for i := 0; i < typ.TypeArgs().Len(); i++ {
			if err := checkDirectCallType(typ.TypeArgs().At(i)); err != nil {
				return err
			}
		}
	case *types.Pointer:
		return checkDirectCallType(typ.Elem())
	case *types.Array:
		return checkDirectCallType(typ.Elem())
	case *types.Slice:
		return checkDirectCallType(typ.Elem())
	case *types.Map:
		if err := checkDirectCallType(typ.Key()); err != nil {
			return err
		}
		return checkDirectCallType(typ.Elem())
	case *types.Chan:
		return checkDirectCallType(typ.Elem())
	case *types.Signature:
		if typ.Recv() != nil {
			if err := checkDirectCallType(typ.Recv().Type()); err != nil {
				return err
			}
		}
		for i := 0; i < typ.Params().Len(); i++ {
			if err := checkDirectCallType(typ.Params().At(i).Type()); err != nil {
				return err
			}
		}
		for i := 0; i < typ.Results().Len(); i++ {
			if err := checkDirectCallType(typ.Results().At(i).Type()); err != nil {
				return err
			}
		}
	case *types.Struct:
		for i := 0; i < typ.NumFields(); i++ {
			if !typ.Field(i).Exported() && typ.Field(i).Pkg() != nil {
				return fmt.Errorf("anonymous struct contains unexported field %s", typ.Field(i).Name())
			}
			if err := checkDirectCallType(typ.Field(i).Type()); err != nil {
				return err
			}
		}
	case *types.Interface:
		for i := 0; i < typ.NumEmbeddeds(); i++ {
			if err := checkDirectCallType(typ.EmbeddedType(i)); err != nil {
				return err
			}
		}
		for i := 0; i < typ.NumExplicitMethods(); i++ {
			method := typ.ExplicitMethod(i)
			if !method.Exported() {
				return fmt.Errorf("interface contains unexported method %s", method.Name())
			}
			if err := checkDirectCallType(method.Type()); err != nil {
				return err
			}
		}
	}
	return nil
}
