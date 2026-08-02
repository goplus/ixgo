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
	"unicode"
)

const (
	importGroupPackage = iota
	importGroupSupport
	importGroupRuntime
)

type plannedImport struct {
	path        string
	packageName string
	alias       string
	group       int
	blank       bool
}

// importPlanner assigns one alias to each import path for the complete
// generated file. Reserving universe and declaration names prevents an import
// from changing the meaning of generated type expressions such as any.
type importPlanner struct {
	byPath map[string]*plannedImport
	used   map[string]bool
}

func newImportPlanner(reserved ...string) *importPlanner {
	p := &importPlanner{
		byPath: make(map[string]*plannedImport),
		used:   make(map[string]bool),
	}
	p.reserve("_", "init")
	p.reserve(types.Universe.Names()...)
	p.reserve(reserved...)
	return p
}

func (p *importPlanner) reserve(names ...string) {
	for _, name := range names {
		if name != "" {
			p.used[name] = true
		}
	}
}

func (p *importPlanner) add(path, packageName, preferredAlias string, group int) string {
	if existing, ok := p.byPath[path]; ok {
		if group < existing.group {
			existing.group = group
		}
		if existing.blank {
			existing.blank = false
			existing.packageName = packageName
			existing.alias = allocateIdentifier(preferredAlias, p.used)
		}
		return existing.alias
	}
	alias := allocateIdentifier(preferredAlias, p.used)
	p.byPath[path] = &plannedImport{
		path:        path,
		packageName: packageName,
		alias:       alias,
		group:       group,
	}
	return alias
}

// addRequired records an import whose qualifier is already embedded in a
// generated expression. It returns an error instead of silently choosing a
// different alias and emitting uncompilable code.
func (p *importPlanner) addRequired(path, packageName, alias string, group int) (string, error) {
	if alias == "" || alias == "_" || alias == "init" || sanitizeIdentifier(alias) != alias ||
		unicode.IsDigit(rune(alias[0])) || token.Lookup(alias).IsKeyword() {
		return "", fmt.Errorf("import %q requires invalid alias %q", path, alias)
	}
	if existing, ok := p.byPath[path]; ok {
		if group < existing.group {
			existing.group = group
		}
		if !existing.blank {
			if existing.alias != alias {
				return "", fmt.Errorf("import %q requires alias %q, but is already named %q", path, alias, existing.alias)
			}
			return existing.alias, nil
		}
		if p.used[alias] {
			return "", fmt.Errorf("import %q requires alias %q, which conflicts with a generated declaration", path, alias)
		}
		p.used[alias] = true
		existing.blank = false
		existing.packageName = packageName
		existing.alias = alias
		return alias, nil
	}
	if p.used[alias] {
		return "", fmt.Errorf("import %q requires alias %q, which conflicts with a generated declaration", path, alias)
	}
	p.used[alias] = true
	p.byPath[path] = &plannedImport{
		path:        path,
		packageName: packageName,
		alias:       alias,
		group:       group,
	}
	return alias, nil
}

func (p *importPlanner) addPackage(pkg *types.Package, group int) string {
	return p.add(pkg.Path(), pkg.Name(), pkg.Name(), group)
}

func (p *importPlanner) addBlank(path string, group int) {
	if existing, ok := p.byPath[path]; ok {
		if group < existing.group {
			existing.group = group
		}
		return
	}
	p.byPath[path] = &plannedImport{path: path, alias: "_", group: group, blank: true}
}

func (p *importPlanner) declarations() []string {
	groups := make(map[int][]*plannedImport)
	var groupIDs []int
	for _, spec := range p.byPath {
		if _, ok := groups[spec.group]; !ok {
			groupIDs = append(groupIDs, spec.group)
		}
		groups[spec.group] = append(groups[spec.group], spec)
	}
	sort.Ints(groupIDs)

	var declarations []string
	for _, group := range groupIDs {
		if len(declarations) != 0 {
			declarations = append(declarations, "")
		}
		specs := groups[group]
		sort.Slice(specs, func(i, j int) bool { return specs[i].path < specs[j].path })
		for _, spec := range specs {
			switch {
			case spec.blank:
				declarations = append(declarations, fmt.Sprintf("_ %q", spec.path))
			case spec.alias == spec.packageName:
				declarations = append(declarations, strconv.Quote(spec.path))
			default:
				declarations = append(declarations, fmt.Sprintf("%s %q", spec.alias, spec.path))
			}
		}
	}
	return declarations
}

func allocateIdentifier(preferred string, used map[string]bool) string {
	base := sanitizeIdentifier(preferred)
	if base == "" || unicode.IsDigit(rune(base[0])) || token.Lookup(base).IsKeyword() {
		base = "qexpPkg"
	}
	if !used[base] {
		used[base] = true
		return base
	}
	for suffix := 2; ; suffix++ {
		name := base + strconv.Itoa(suffix)
		if !used[name] {
			used[name] = true
			return name
		}
	}
}

func sanitizeIdentifier(name string) string {
	var b strings.Builder
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			b.WriteRune(r)
		} else {
			b.WriteByte('_')
		}
	}
	return b.String()
}
