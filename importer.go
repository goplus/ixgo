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

package ixgo

import (
	"fmt"
	"go/importer"
	"go/types"
)

type Importer struct {
	ctx         *Context
	pkgs        map[string]*types.Package
	importing   map[string]bool
	defaultImpl types.Importer
}

func NewImporter(ctx *Context) *Importer {
	return &Importer{
		ctx:         ctx,
		pkgs:        make(map[string]*types.Package),
		importing:   make(map[string]bool),
		defaultImpl: importer.Default(),
	}
}

func (i *Importer) Import(path string) (ret *types.Package, err error) {
	if pkg, ok := i.pkgs[path]; ok {
		return pkg, nil
	}
	if i.importing[path] {
		return nil, fmt.Errorf("cycle importing package %q", path)
	}
	i.importing[path] = true
	defer func() {
		i.importing[path] = false
	}()
	defer func() {
		if err == nil {
			// check patch
			if sp, ok := i.ctx.pkgs[path+"@patch"]; ok {
				if err = sp.Load(); err == nil {
					scope := sp.Package.Scope()
					rscope := ret.Scope()
					names := scope.Names()
					for _, name := range names {
						obj := scope.Lookup(name)
						rscope.Insert(obj)
					}
				}
			}
		}
	}()
	if pkg, err := i.ctx.Loader.Import(path); err == nil && pkg.Complete() {
		i.pkgs[path] = pkg
		return pkg, nil
	}
	if pkg, ok := i.ctx.pkgs[path]; ok {
		if !pkg.Package.Complete() {
			if err := pkg.Load(); err != nil {
				return nil, err
			}
		}
		if !pkg.NoCache {
			i.pkgs[path] = pkg.Package
		}
		return pkg.Package, nil
	}
	if dir, found := i.ctx.lookupPath(path); found {
		pkg, err := i.ctx.addImport(path, dir)
		if err != nil {
			return nil, err
		}
		if err := pkg.Load(); err != nil {
			return nil, err
		}
		return pkg.Package, nil
	}
	return nil, ErrNotFoundPackage
}
