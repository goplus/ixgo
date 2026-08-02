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
	"go/types"
	"strings"
	"testing"
)

func TestImportPlannerReusesPathsAndProtectsIdentifiers(t *testing.T) {
	planner := newImportPlanner("generated")
	reflectAlias := planner.add("reflect", "reflect", "reflect", importGroupSupport)
	if got := planner.addPackage(types.NewPackage("reflect", "reflect"), importGroupSupport); got != reflectAlias {
		t.Fatalf("reflect alias = %q; want reused alias %q", got, reflectAlias)
	}

	for _, test := range []struct {
		path string
		name string
	}{
		{path: "example.com/any", name: "any"},
		{path: "example.com/generated", name: "generated"},
		{path: "example.com/init", name: "init"},
	} {
		alias := planner.addPackage(types.NewPackage(test.path, test.name), importGroupSupport)
		if alias == test.name {
			t.Fatalf("import %s shadows protected identifier %q", test.path, alias)
		}
	}

	declarations := strings.Join(planner.declarations(), "\n")
	if got := strings.Count(declarations, `"reflect"`); got != 1 {
		t.Fatalf("reflect import count = %d; want 1\n%s", got, declarations)
	}
}

func TestImportPlannerRequiredAlias(t *testing.T) {
	for _, alias := range []string{"constant", "token"} {
		t.Run(alias, func(t *testing.T) {
			planner := newImportPlanner(alias)
			if _, err := planner.addRequired("go/"+alias, alias, alias, importGroupSupport); err == nil {
				t.Fatalf("required alias %q was silently renamed", alias)
			}
		})
	}

	planner := newImportPlanner()
	alias, err := planner.addRequired("go/constant", "constant", "constant", importGroupSupport)
	if err != nil {
		t.Fatal(err)
	}
	if alias != "constant" {
		t.Fatalf("required alias = %q; want constant", alias)
	}
}

func TestExportPkgPropagatesPlannedSupportAliases(t *testing.T) {
	oldExportAlias := flagExportAlias
	flagExportAlias = true
	t.Cleanup(func() { flagExportAlias = oldExportAlias })

	pkg := &Package{
		Name:       "generated",
		Path:       "example.com/generated",
		NamedTypes: []string{`"Value": ` + reflectImportPlaceholder + `.TypeOf((*int)(nil)).Elem()`},
		Alias:      []string{`"Value": &` + aliasImportPlaceholder + `.Builtin{Typ: "int"}`},
		TypesData:  []byte{},
		Links: []string{
			"func reflect() {}",
			"func alias() {}",
			"func typesdata() {}",
		},
	}
	source, err := exportPkg(pkg, "q", "", nil, "export")
	if err != nil {
		t.Fatal(err)
	}
	generated := string(source)
	for _, want := range []string{
		`reflect2 "reflect"`,
		`alias2 "github.com/goplus/ixgo/alias"`,
		`typesdata2 "github.com/goplus/ixgo/xgobuild/typesdata"`,
		`map[string]reflect2.Type`,
		`reflect2.TypeOf`,
		`map[string]alias2.Type`,
		`&alias2.Builtin`,
		`typesdata2.ImportFunc`,
	} {
		if !strings.Contains(generated, want) {
			t.Fatalf("generated source does not contain %q:\n%s", want, generated)
		}
	}
}

func TestExportPkgReusesConstantSourceImport(t *testing.T) {
	pkg := &Package{
		Name:    "constant",
		Path:    "go/constant",
		usedPkg: true,
		UntypedConsts: []string{
			`"Text": {Typ: "untyped string", Value: ` + constantImportPlaceholder + `.MakeString("constant.value")}`,
		},
	}
	source, err := exportPkg(pkg, "q", "", nil, "export")
	if err != nil {
		t.Fatal(err)
	}
	generated := string(source)
	if got := strings.Count(generated, `q "go/constant"`); got != 1 {
		t.Fatalf("go/constant import count = %d; want 1\n%s", got, generated)
	}
	for _, want := range []string{`q "go/constant"`, `q.MakeString("constant.value")`} {
		if !strings.Contains(generated, want) {
			t.Fatalf("generated source does not contain %q:\n%s", want, generated)
		}
	}
}
