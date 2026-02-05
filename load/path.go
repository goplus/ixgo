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

package load

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

// GetImportPath get import path from dir.
// - lookup go.mod and check modpath+dir+pkgname
// - use go list check if BuildMod == "mod".
func GetImportPath(pkgName string, dir string) (string, error) {
	if BuildMod == "mod" {
		data, err := runGoCommand(dir, "list", "-e", "-mod=mod")
		if err != nil {
			return "", err
		}
		return string(bytes.TrimSuffix(data, []byte{'\n'})), nil
	}
	dir, err := absDir(dir)
	if err != nil {
		return "", err
	}
	if pkgName == "" || pkgName == "main" {
		_, pkgName = filepath.Split(dir)
	}
	mod, found := findModule(dir)
	if !found {
		return pkgName, nil
	}
	data, err := ioutil.ReadFile(mod)
	if err != nil {
		return "", err
	}
	f, err := ParseModFile(mod, data)
	if err != nil {
		return "", err
	}
	root := filepath.Dir(f.Syntax.Name)
	modPath := f.Module.Mod.Path
	dir = filepath.ToSlash(dir)
	root = filepath.ToSlash(root)
	if dir == root {
		return modPath, nil
	}
	_dir, _ := path.Split(dir[len(root)+1:])
	return path.Join(modPath, _dir, pkgName), nil
}

func absDir(dir string) (string, error) {
	if dir == "" {
		dir = "."
	}
	return filepath.Abs(dir)
}

func findModule(dir string) (file string, found bool) {
	for dir != "" {
		file = filepath.Join(dir, "go.mod")
		if fi, e := os.Lstat(file); e == nil && !fi.IsDir() {
			found = true
			return
		}
		if dir, file = filepath.Split(strings.TrimRight(dir, "/\\")); file == "" {
			break
		}
	}
	return
}

// ParseModFile parse go.mod
func ParseModFile(file string, data []byte) (*modfile.File, error) {
	fix := func(path, vers string) (resolved string, err error) {
		// do nothing
		return vers, nil
	}
	f, err := modfile.Parse(file, data, fix)
	if err != nil {
		return nil, fmt.Errorf("parse go.mod error %w", err)
	}
	if f.Module == nil {
		return nil, errors.New("no module declaration in go.mod")
	}
	return f, nil
}
