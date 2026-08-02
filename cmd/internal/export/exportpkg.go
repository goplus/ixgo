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

package export

import (
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

func writeFile(dir string, file string, data []byte) error {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return fmt.Errorf("make dir %v error: %v", dir, err)
	}
	filename := filepath.Join(dir, file)
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		return fmt.Errorf("write file %v error: %v", filename, err)
	}
	return nil
}

func joinList(list []string) string {
	if len(list) == 0 {
		return ""
	}
	sort.Strings(list)
	return "\n\t" + strings.Join(list, ",\n\t") + ",\n"
}

const (
	reflectImportPlaceholder  = "\x00qexp-reflect\x00"
	aliasImportPlaceholder    = "\x00qexp-alias\x00"
	constantImportPlaceholder = "\x00qexp-constant\x00"
	tokenImportPlaceholder    = "\x00qexp-token\x00"
)

func joinGeneratedList(list []string, replacements ...string) string {
	if len(list) == 0 {
		return ""
	}
	replacer := strings.NewReplacer(replacements...)
	rendered := make([]string, len(list))
	for i, item := range list {
		rendered[i] = replacer.Replace(item)
	}
	return joinList(rendered)
}

func packageConstsUseToken(pkg *Package) bool {
	for _, constants := range [][]string{pkg.UntypedConsts, pkg.TypedConsts} {
		for _, constant := range constants {
			if strings.Contains(constant, tokenImportPlaceholder+".") {
				return true
			}
		}
	}
	return false
}

func linkDeclarationNames(links []string) []string {
	var names []string
	for _, link := range links {
		for rest := link; ; {
			index := strings.Index(rest, "func ")
			if index < 0 {
				break
			}
			rest = rest[index+len("func "):]
			end := strings.IndexByte(rest, '(')
			if end < 0 {
				break
			}
			if name := strings.TrimSpace(rest[:end]); name != "" {
				names = append(names, name)
			}
			rest = rest[end+1:]
		}
	}
	return names
}

const EmptyPackage = "empty package"

var (
	errEmptyPackage = errors.New(EmptyPackage)
)

func exportPkg(pkg *Package, sname string, id string, tagList []string, fname string) ([]byte, error) {
	tmpl := template_pkg
	if pkg.IsEmpty() {
		tmpl = template_empty_pkg
	} else if pkg.TypesData != nil {
		tmpl = template_pkg_types
	}
	if len(pkg.Source) > 0 {
		tmpl = template_link_pkg
		if pkg.IsEmpty() {
			tmpl = template_emtpy_link_pkg
		}
	}
	sourcePath, sourceName := pkg.Path, pkg.Name
	if flagCustomPkg != "" {
		pkg.Path = path.Clean(flagCustomPkg)
		pkg.Name = path.Base(pkg.Path)
	}
	if flagExportLazy {
		tmpl = strings.Replace(tmpl, "$INIT", "", 1)
		tmpl = strings.Replace(tmpl, "$IXGO.RegisterPackage(",
			`$IXGO.RegisterPackageLazy("$PKGPATH",func() *$IXGO.Package {$INIT
	return `, 1)
		tmpl = strings.Replace(tmpl, "})", "}\n\t})", 1)
	}

	reserved := append([]string{"source", fname + "TypesData"}, pkg.directCalls.declarationNames()...)
	reserved = append(reserved, linkDeclarationNames(pkg.Links)...)
	imports := newImportPlanner(reserved...)
	if pkg.usedPkg {
		if _, err := imports.addRequired(sourcePath, sourceName, sname, importGroupPackage); err != nil {
			return nil, err
		}
	} else if !flagExportCode {
		imports.addBlank(sourcePath, importGroupPackage)
	}
	reflectAlias := "reflect"
	if !pkg.IsEmpty() {
		reflectAlias = imports.add("reflect", "reflect", "reflect", importGroupSupport)
	}
	var constantAlias, tokenAlias string
	if len(pkg.UntypedConsts) > 0 || len(pkg.TypedConsts) > 0 {
		constantAlias = imports.add("go/constant", "constant", "constant", importGroupSupport)
		if packageConstsUseToken(pkg) {
			tokenAlias = imports.add("go/token", "token", "token", importGroupSupport)
		}
	}
	if len(pkg.Links) > 0 {
		imports.addBlank("unsafe", importGroupSupport)
	}

	var ext, aliasAlias string
	if len(pkg.Alias) != 0 && flagExportAlias {
		aliasAlias = imports.add("github.com/goplus/ixgo/alias", "alias", "alias", importGroupSupport)
		ext = fmt.Sprintf("\nAlias: map[string]%s.Type{%s},", aliasAlias,
			joinGeneratedList(pkg.Alias, aliasImportPlaceholder, aliasAlias))
	}
	runtimeAlias := imports.add("github.com/goplus/ixgo", "ixgo", "ixgo", importGroupRuntime)
	typesDataAlias := ""
	if pkg.TypesData != nil {
		imports.addBlank("embed", importGroupRuntime)
		typesDataAlias = imports.add("github.com/goplus/ixgo/xgobuild/typesdata", "typesdata", "typesdata", importGroupRuntime)
	}
	entries, adapters := pkg.directCalls.render(sname, runtimeAlias, reflectAlias, imports)
	var directCalls string
	if len(entries) != 0 {
		directCalls = fmt.Sprintf("\nDirectCalls: map[string]%s.DirectCallBinding{%s},", runtimeAlias, joinList(entries))
	}
	directCallAdapters := strings.Join(adapters, "\n\n")
	r := strings.NewReplacer("$PKGNAME", pkg.Name,
		"$IMPORTS", strings.Join(imports.declarations(), "\n"),
		"$IXGO", runtimeAlias,
		"$REFLECT", reflectAlias,
		"$TYPESDATA", typesDataAlias,
		"$PKGPATH", pkg.Path,
		"$DEPS", joinList(pkg.Deps),
		"$NAMEDTYPES", joinGeneratedList(pkg.NamedTypes, reflectImportPlaceholder, reflectAlias),
		"$INTERFACES", joinGeneratedList(pkg.Interfaces, reflectImportPlaceholder, reflectAlias),
		"$ALIASTYPES", joinGeneratedList(pkg.AliasTypes, reflectImportPlaceholder, reflectAlias),
		"$VARS", joinGeneratedList(pkg.Vars, reflectImportPlaceholder, reflectAlias),
		"$FUNCS", joinGeneratedList(pkg.Funcs, reflectImportPlaceholder, reflectAlias),
		"$DIRECTCALLS", directCalls,
		"$DIRECTCALLADAPTERS", directCallAdapters,
		"$TYPEDCONSTS", joinGeneratedList(pkg.TypedConsts,
			reflectImportPlaceholder, reflectAlias,
			constantImportPlaceholder, constantAlias,
			tokenImportPlaceholder, tokenAlias),
		"$UNTYPEDCONSTS", joinGeneratedList(pkg.UntypedConsts,
			constantImportPlaceholder, constantAlias,
			tokenImportPlaceholder, tokenAlias),
		"$TAGS", strings.Join(tagList, "\n"),
		"$SOURCE", pkg.Source,
		"$LINKS", strings.Join(pkg.Links, "\n"),
		"$ID", id,
		"$TYPESNAME", fname+"TypesData",
		"$TYPESFILE", fname+".types",
		"$EXT", ext,
		"$INIT", strings.ReplaceAll(pkg.AliasInit, aliasImportPlaceholder, aliasAlias))
	src := r.Replace(tmpl)
	data, err := format.Source([]byte(src))
	if err != nil {
		return nil, fmt.Errorf("format pkg %v error: %v", src, err)
	}
	return data, nil
}

var template_pkg = `// export by github.com/goplus/ixgo/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS
)

func init() {$INIT
	$IXGO.RegisterPackage(&$IXGO.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
		Interfaces: map[string]$REFLECT.Type{$INTERFACES},
		NamedTypes: map[string]$REFLECT.Type{$NAMEDTYPES},
		AliasTypes: map[string]$REFLECT.Type{$ALIASTYPES},
		Vars: map[string]$REFLECT.Value{$VARS},
		Funcs: map[string]$REFLECT.Value{$FUNCS},$DIRECTCALLS
		TypedConsts: map[string]$IXGO.TypedConst{$TYPEDCONSTS},
		UntypedConsts: map[string]$IXGO.UntypedConst{$UNTYPEDCONSTS},$EXT
	})
}

$DIRECTCALLADAPTERS
`

var template_pkg_types = `// export by github.com/goplus/ixgo/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS
)

//go:embed $TYPESFILE
var $TYPESNAME []byte

func init() {$INIT
	$IXGO.RegisterPackage(&$IXGO.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
		Interfaces: map[string]$REFLECT.Type{$INTERFACES},
		NamedTypes: map[string]$REFLECT.Type{$NAMEDTYPES},
		AliasTypes: map[string]$REFLECT.Type{$ALIASTYPES},
		Vars: map[string]$REFLECT.Value{$VARS},
		Funcs: map[string]$REFLECT.Value{$FUNCS},$DIRECTCALLS
		TypedConsts: map[string]$IXGO.TypedConst{$TYPEDCONSTS},
		UntypedConsts: map[string]$IXGO.UntypedConst{$UNTYPEDCONSTS},
		Import: $TYPESDATA.ImportFunc("$PKGPATH", $TYPESNAME),$EXT
	})
}

$DIRECTCALLADAPTERS
`

var template_empty_pkg = `// export by github.com/goplus/ixgo/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS
)

func init() {$INIT
	$IXGO.RegisterPackage(&$IXGO.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},$EXT
	})
}

$DIRECTCALLADAPTERS
`

var template_link_pkg = `// export by github.com/goplus/ixgo/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS
)

func init() {$INIT
	$IXGO.RegisterPackage(&$IXGO.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
		Interfaces: map[string]$REFLECT.Type{$INTERFACES},
		NamedTypes: map[string]$REFLECT.Type{$NAMEDTYPES},
		AliasTypes: map[string]$REFLECT.Type{$ALIASTYPES},
		Vars: map[string]$REFLECT.Value{$VARS},
		Funcs: map[string]$REFLECT.Value{$FUNCS},$DIRECTCALLS
		TypedConsts: map[string]$IXGO.TypedConst{$TYPEDCONSTS},
		UntypedConsts: map[string]$IXGO.UntypedConst{$UNTYPEDCONSTS},
		Source: source,$EXT
	})
}
$LINKS
var source = $SOURCE

$DIRECTCALLADAPTERS
`

var template_emtpy_link_pkg = `// export by github.com/goplus/ixgo/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS
)

func init() {$INIT
	$IXGO.RegisterPackage(&$IXGO.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
		Source: source,$EXT
	})
}
$LINKS
var source = $SOURCE

$DIRECTCALLADAPTERS
`
