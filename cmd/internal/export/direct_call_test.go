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
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func testSignature(recv types.Type, params, results []types.Type, variadic bool) *types.Signature {
	vars := func(list []types.Type) *types.Tuple {
		ret := make([]*types.Var, len(list))
		for i, typ := range list {
			ret[i] = types.NewVar(token.NoPos, nil, "", typ)
		}
		return types.NewTuple(ret...)
	}
	var recvVar *types.Var
	if recv != nil {
		recvVar = types.NewVar(token.NoPos, nil, "", recv)
	}
	return types.NewSignatureType(recvVar, nil, nil, vars(params), vars(results), variadic)
}

func addTestFunc(pkg *types.Package, name string, sig *types.Signature) *types.Func {
	fn := types.NewFunc(token.NoPos, pkg, name, sig)
	pkg.Scope().Insert(fn)
	return fn
}

func addTestType(pkg *types.Package, name string) *types.Named {
	obj := types.NewTypeName(token.NoPos, pkg, name, nil)
	named := types.NewNamed(obj, types.NewStruct(nil, nil), nil)
	pkg.Scope().Insert(obj)
	return named
}

func renderTestDirectCalls(output directCallOutput) (entries, declarations, imports []string) {
	planner := newImportPlanner(output.declarationNames()...)
	entries, declarations = output.render("q", "ixgo", "reflect", planner)
	return entries, declarations, planner.declarations()
}

func testRegistryEntries(output directCallOutput) []string {
	entries, _, _ := renderTestDirectCalls(output)
	return entries
}

func testAdapterDeclarations(output directCallOutput) []string {
	_, declarations, _ := renderTestDirectCalls(output)
	return declarations
}

func TestParseDirectCallSelectors(t *testing.T) {
	got, err := parseDirectCallSelectors(" List.At, Warp;Rand__0 ")
	if err != nil {
		t.Fatal(err)
	}
	if want := "List.At,Rand__0,Warp"; strings.Join(got, ",") != want {
		t.Fatalf("selectors = %q; want %q", strings.Join(got, ","), want)
	}
	if _, err := parseDirectCallSelectors("Warp,Warp"); err == nil {
		t.Fatal("duplicate selector was accepted")
	}
}

func TestGenerateDirectCalls(t *testing.T) {
	pkg := types.NewPackage("example.com/host", "host")
	addTestFunc(pkg, "Add", testSignature(nil, []types.Type{types.Typ[types.Int], types.Typ[types.Int]}, []types.Type{types.Typ[types.Int]}, false))

	list := addTestType(pkg, "List")
	anyType := types.Universe.Lookup("any").Type()
	appendMethod := types.NewFunc(token.NoPos, pkg, "Append", testSignature(types.NewPointer(list), []types.Type{anyType}, nil, false))
	list.AddMethod(appendMethod)

	output, err := generateDirectCalls(pkg, "List.Append,Add")
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Join(append(testRegistryEntries(output), testAdapterDeclarations(output)...), "\n")
	want := `"Add": {Target: reflect.ValueOf(q.Add), Adapter: qexpDirectCallFunc_Add}
"(*List).Append": {Target: reflect.ValueOf((*q.List).Append), Adapter: qexpDirectCallMethod_Ptr_List_Append}
func qexpDirectCallFunc_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Add(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}
func qexpDirectCallMethod_Ptr_List_Append(ctx ixgo.DirectCallContext) {
	(*q.List).Append(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[any](ctx, 1))
}`
	if got != want {
		t.Fatalf("generated direct calls differ:\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}

func TestGenerateDirectCallValueMethodAndVariadicFunction(t *testing.T) {
	pkg := types.NewPackage("example.com/host", "host")
	addTestFunc(pkg, "Collect", testSignature(nil, []types.Type{types.Typ[types.String], types.NewSlice(types.Typ[types.Int])}, []types.Type{types.Typ[types.Int]}, true))

	list := addTestType(pkg, "List")
	list.AddMethod(types.NewFunc(token.NoPos, pkg, "Len", testSignature(list, nil, []types.Type{types.Typ[types.Int]}, false)))

	output, err := generateDirectCalls(pkg, "List.Len,Collect")
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Join(testAdapterDeclarations(output), "\n")
	want := `func qexpDirectCallFunc_Collect(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Collect(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[[]int](ctx, 1)...))
}
func qexpDirectCallMethod_List_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.List.Len(ixgo.DirectCallArg[q.List](ctx, 0)))
}
func qexpDirectCallMethod_Ptr_List_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).Len(ixgo.DirectCallArg[*q.List](ctx, 0)))
}`
	if got != want {
		t.Fatalf("generated adapters differ:\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}

func TestGenerateDirectCallDeduplicatesReceiverAlias(t *testing.T) {
	pkg := types.NewPackage("example.com/host", "host")
	list := addTestType(pkg, "List")
	list.AddMethod(types.NewFunc(token.NoPos, pkg, "Len", testSignature(list, nil, []types.Type{types.Typ[types.Int]}, false)))
	aliasName := types.NewTypeName(token.NoPos, pkg, "ListAlias", nil)
	types.NewAlias(aliasName, list)
	pkg.Scope().Insert(aliasName)

	output, err := generateDirectCalls(pkg, "all,ListAlias.Len")
	if err != nil {
		t.Fatal(err)
	}
	if got := len(output.bindings); got != 2 {
		t.Fatalf("bindings = %d; want one value and one pointer binding", got)
	}
	entries := strings.Join(testRegistryEntries(output), "\n")
	for _, key := range []string{
		`"List.Len"`,
		`"(*List).Len"`,
	} {
		if got := strings.Count(entries, key); got != 1 {
			t.Fatalf("binding %s count = %d; want 1\n%s", key, got, entries)
		}
	}
}

func TestGenerateDirectCallWildcards(t *testing.T) {
	pkg := types.NewPackage("example.com/host", "host")
	addTestFunc(pkg, "Add", testSignature(nil, []types.Type{types.Typ[types.Int]}, []types.Type{types.Typ[types.Int]}, false))
	addTestFunc(pkg, "Pair", testSignature(nil, nil, []types.Type{types.Typ[types.Int], types.Typ[types.Int]}, false))
	hidden := addTestType(pkg, "hidden")
	addTestFunc(pkg, "UseHidden", testSignature(nil, []types.Type{hidden}, nil, false))

	list := addTestType(pkg, "List")
	list.AddMethod(types.NewFunc(token.NoPos, pkg, "Append", testSignature(types.NewPointer(list), []types.Type{types.Typ[types.Int]}, nil, false)))
	list.AddMethod(types.NewFunc(token.NoPos, pkg, "Pair", testSignature(types.NewPointer(list), nil, []types.Type{types.Typ[types.Int], types.Typ[types.Int]}, false)))
	value := addTestType(pkg, "Value")
	value.AddMethod(types.NewFunc(token.NoPos, pkg, "Int", testSignature(value, nil, []types.Type{types.Typ[types.Int]}, false)))

	output, err := generateDirectCalls(pkg, "*,List.*")
	if err != nil {
		t.Fatal(err)
	}
	entries := strings.Join(testRegistryEntries(output), "\n")
	for _, want := range []string{`"Add"`, `"(*List).Append"`} {
		if !strings.Contains(entries, want) {
			t.Fatalf("wildcard output does not contain %s:\n%s", want, entries)
		}
	}
	for _, skipped := range []string{"Pair", "UseHidden", "Value.Int"} {
		if strings.Contains(entries, skipped) {
			t.Fatalf("wildcard output unexpectedly contains %q:\n%s", skipped, entries)
		}
	}

	output, err = generateDirectCalls(pkg, "all,List.Append")
	if err != nil {
		t.Fatal(err)
	}
	if got := len(output.bindings); got != 4 {
		t.Fatalf("deduplicated wildcard bindings = %d; want 4", got)
	}
	entries = strings.Join(testRegistryEntries(output), "\n")
	for _, want := range []string{`"Value.Int"`, `"(*Value).Int"`} {
		if !strings.Contains(entries, want) {
			t.Fatalf("all-method wildcard does not contain %s:\n%s", want, entries)
		}
	}
	selectorsOutput, err := generateDirectCalls(pkg, "*,*.*")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := strings.Join(testRegistryEntries(output), "\n"), strings.Join(testRegistryEntries(selectorsOutput), "\n"); got != want {
		t.Fatalf("all selector differs from *,*.*:\n--- all ---\n%s\n--- selectors ---\n%s", got, want)
	}

	if _, err := generateDirectCalls(pkg, "*,Pair"); err == nil {
		t.Fatal("explicit unsupported selector overlapping a wildcard did not fail")
	}
	if _, err := generateDirectCalls(pkg, "Missing.*"); err == nil {
		t.Fatal("wildcard for a missing receiver type did not fail")
	}
}

func TestGenerateDirectCallSkipsNestedUnexportedAlias(t *testing.T) {
	pkg := types.NewPackage("example.com/host", "host")
	aliasName := types.NewTypeName(token.NoPos, pkg, "obj", nil)
	obj := types.NewAlias(aliasName, types.Universe.Lookup("any").Type())
	pkg.Scope().Insert(aliasName)
	addTestFunc(pkg, "Collect", testSignature(nil, []types.Type{types.NewSlice(obj)}, nil, true))

	output, err := generateDirectCalls(pkg, "all")
	if err != nil {
		t.Fatal(err)
	}
	if got := len(output.bindings); got != 0 {
		t.Fatalf("bindings = %d; want 0", got)
	}
	if _, err := generateDirectCalls(pkg, "Collect"); err == nil {
		t.Fatal("explicit selector with an inaccessible nested alias was accepted")
	}
}

func TestGenerateDirectCallImportsArgumentTypes(t *testing.T) {
	host := types.NewPackage("example.com/host", "host")
	dep := types.NewPackage("example.com/data", "data")
	data := types.NewNamed(types.NewTypeName(token.NoPos, dep, "Data", nil), types.NewStruct(nil, nil), nil)
	addTestFunc(host, "Use", testSignature(nil, []types.Type{data}, nil, false))

	output, err := generateDirectCalls(host, "Use")
	if err != nil {
		t.Fatal(err)
	}
	_, _, imports := renderTestDirectCalls(output)
	if len(imports) != 1 || imports[0] != `"example.com/data"` {
		t.Fatalf("imports = %#v; want data import", imports)
	}
	adapter := testAdapterDeclarations(output)[0]
	if !strings.Contains(adapter, "DirectCallArg[data.Data]") {
		t.Fatalf("adapter does not use imported argument type: %s", adapter)
	}
}

func TestGenerateDirectCallRejectsUnsupportedSelector(t *testing.T) {
	pkg := types.NewPackage("example.com/host", "host")
	addTestFunc(pkg, "Pair", testSignature(nil, nil, []types.Type{types.Typ[types.Int], types.Typ[types.Int]}, false))
	dep := types.NewPackage("example.com/dep", "dep")
	foreign := addTestType(dep, "Foreign")
	foreign.AddMethod(types.NewFunc(token.NoPos, dep, "Value", testSignature(foreign, nil, nil, false)))
	aliasName := types.NewTypeName(token.NoPos, pkg, "ForeignAlias", nil)
	types.NewAlias(aliasName, foreign)
	pkg.Scope().Insert(aliasName)

	for _, selector := range []string{"Missing", "Pair", "Type.Missing", "ForeignAlias.Value"} {
		if _, err := generateDirectCalls(pkg, selector); err == nil {
			t.Fatalf("selector %q was accepted", selector)
		}
	}
}

func TestGenerateDirectCallRejectsGenericReceiver(t *testing.T) {
	pkg := types.NewPackage("example.com/host", "host")
	box := addTestType(pkg, "Box")
	constraint := types.NewInterfaceType(nil, nil)
	constraint.Complete()
	typeParam := types.NewTypeParam(types.NewTypeName(token.NoPos, pkg, "T", nil), constraint)
	sig := types.NewSignatureType(
		types.NewVar(token.NoPos, pkg, "", types.NewPointer(box)),
		[]*types.TypeParam{typeParam}, nil,
		types.NewTuple(), types.NewTuple(types.NewVar(token.NoPos, nil, "", typeParam)), false,
	)
	fn := types.NewFunc(token.NoPos, pkg, "Value", sig)

	if err := validateDirectCall(fn); err == nil {
		t.Fatal("method with a generic receiver was accepted")
	}
}

func TestGenerateDirectCallRejectsUnexportedTypeArgument(t *testing.T) {
	pkg := types.NewPackage("example.com/host", "host")
	box := addTestType(pkg, "Box")
	constraint := types.NewInterfaceType(nil, nil)
	constraint.Complete()
	box.SetTypeParams([]*types.TypeParam{
		types.NewTypeParam(types.NewTypeName(token.NoPos, pkg, "T", nil), constraint),
	})
	hidden := addTestType(pkg, "hidden")
	instantiated, err := types.Instantiate(nil, box, []types.Type{hidden}, true)
	if err != nil {
		t.Fatal(err)
	}
	addTestFunc(pkg, "Use", testSignature(nil, []types.Type{instantiated}, nil, false))

	if _, err := generateDirectCalls(pkg, "Use"); err == nil {
		t.Fatal("unexported named type argument was accepted")
	}
}

func TestGenerateDirectCallRejectsUnexportedEmbeddedInterface(t *testing.T) {
	pkg := types.NewPackage("example.com/host", "host")
	hiddenName := types.NewTypeName(token.NoPos, pkg, "hiddenInterface", nil)
	hidden := types.NewNamed(hiddenName, types.NewInterfaceType(nil, nil).Complete(), nil)
	pkg.Scope().Insert(hiddenName)
	param := types.NewInterfaceType(nil, []types.Type{hidden})
	param.Complete()
	addTestFunc(pkg, "Use", testSignature(nil, []types.Type{param}, nil, false))

	if _, err := generateDirectCalls(pkg, "Use"); err == nil {
		t.Fatal("unexported embedded interface was accepted")
	}
}

func TestGenerateDirectCallAdapterNamesDoNotCollide(t *testing.T) {
	pkg := types.NewPackage("example.com/host", "host")
	aUnderscoreB := addTestType(pkg, "A_B")
	aUnderscoreB.AddMethod(types.NewFunc(token.NoPos, pkg, "C", testSignature(aUnderscoreB, nil, nil, false)))
	a := addTestType(pkg, "A")
	a.AddMethod(types.NewFunc(token.NoPos, pkg, "B_C", testSignature(a, nil, nil, false)))

	output, err := generateDirectCalls(pkg, "A_B.C,A.B_C")
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Join(testAdapterDeclarations(output), "\n")
	for _, name := range []string{
		"qexpDirectCallMethod_A_B_C",
		"qexpDirectCallMethod_A_B_C2",
		"qexpDirectCallMethod_Ptr_A_B_C",
		"qexpDirectCallMethod_Ptr_A_B_C2",
	} {
		if !strings.Contains(got, "func "+name+"(") {
			t.Fatalf("generated adapters do not contain %q:\n%s", name, got)
		}
	}
}

func TestGeneratedDirectCallsCompileWithSharedImports(t *testing.T) {
	const fixturePath = "github.com/goplus/ixgo/cmd/internal/export/testdata/directfixture"
	prog := NewProgram(nil)
	if err := prog.Load([]string{fixturePath}); err != nil {
		t.Fatal(err)
	}
	pkg, err := prog.ExportPkg(fixturePath, "q")
	if err != nil {
		t.Fatal(err)
	}
	directCalls, err := generateDirectCalls(prog.prog.Package(fixturePath).Pkg, "Inspect,Value.Number")
	if err != nil {
		t.Fatal(err)
	}
	pkg.directCalls = directCalls

	source, err := exportPkg(pkg, "q", "", nil, "export")
	if err != nil {
		t.Fatal(err)
	}
	file, err := parser.ParseFile(token.NewFileSet(), "export.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	importCounts := make(map[string]int)
	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			t.Fatal(err)
		}
		importCounts[path]++
	}
	for _, path := range []string{"reflect", "go/constant"} {
		if got := importCounts[path]; got != 1 {
			t.Fatalf("generated import %q count = %d; want 1\n%s", path, got, source)
		}
	}
	for _, key := range []string{
		`"Value.Number"`,
		`"(*Value).Number"`,
	} {
		if !strings.Contains(string(source), key) {
			t.Fatalf("generated source does not contain %s\n%s", key, source)
		}
	}
	for _, want := range []string{
		`ixgo.RegisterPackage(&ixgo.Package{`,
		`ixgo.RegisterDirectCalls("` + fixturePath + `", map[string]ixgo.DirectCallBinding{`,
	} {
		if !strings.Contains(string(source), want) {
			t.Fatalf("generated source does not contain %q\n%s", want, source)
		}
	}
	if strings.Contains(string(source), "DirectCalls:") {
		t.Fatalf("generated package still embeds direct calls\n%s", source)
	}

	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate repository root")
	}
	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(testFile), "..", "..", ".."))
	tempDir := t.TempDir()
	goMod := "module github.com/goplus/ixgo/cmd/internal/export/generatedtest\n\n" +
		"go 1.24.0\n\n" +
		"require github.com/goplus/ixgo v0.0.0\n\n" +
		"replace github.com/goplus/ixgo => " + filepath.ToSlash(repoRoot) + "\n"
	if err := os.WriteFile(filepath.Join(tempDir, "go.mod"), []byte(goMod), 0o666); err != nil {
		t.Fatal(err)
	}
	goSum, err := os.ReadFile(filepath.Join(repoRoot, "go.sum"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tempDir, "go.sum"), goSum, 0o666); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tempDir, "export.go"), source, 0o666); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", "-mod=mod", "-run", "^$", ".")
	cmd.Dir = tempDir
	cmd.Env = append(os.Environ(), "GOWORK=off")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated package does not compile: %v\n%s\n--- source ---\n%s", err, output, source)
	}
}

func TestExportPkgsReturnsDirectCallError(t *testing.T) {
	oldDirectCalls, oldExportDir, oldCustomPkg, oldExportCode := flagDirectCalls, flagExportDir, flagCustomPkg, flagExportCode
	flagDirectCalls, flagExportDir = "Missing", t.TempDir()
	t.Cleanup(func() {
		flagDirectCalls, flagExportDir, flagCustomPkg, flagExportCode = oldDirectCalls, oldExportDir, oldCustomPkg, oldExportCode
	})

	if err := ExportPkgs([]string{"github.com/goplus/ixgo/testdata/directcall"}, nil); err == nil {
		t.Fatal("invalid direct-call selector did not fail the export")
	}

	flagDirectCalls, flagCustomPkg = "Counter.Value", "example.com/custom"
	if err := ExportPkgs([]string{"github.com/goplus/ixgo/testdata/directcall"}, nil); err == nil {
		t.Fatal("directcalls combined with pkgpath did not fail the export")
	}

	flagCustomPkg, flagExportCode = "", true
	if err := ExportPkgs([]string{"github.com/goplus/ixgo/testdata/directcall"}, nil); err == nil {
		t.Fatal("directcalls combined with code did not fail the export")
	}
}
