//go:build !wasm
// +build !wasm

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

package list

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

var (
	BuildMod string // BuildMod can be set readonly, vendor, or mod.
)

// ListDriver implement (*ixgo.Context).Lookup use go list
type ListDriver struct {
	init bool
	root string
	pkgs map[string]string // path -> dir
}

// Lookup implement (*ixgo.Context).Lookup
func (d *ListDriver) Lookup(root string, path string) (dir string, found bool) {
	if !d.init || d.root != root {
		d.init = true
		d.root = root
		err := d.Parse(root)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	dir, found = d.pkgs[path]
	if found {
		return
	}
	var list []string
	for k := range d.pkgs {
		if strings.HasPrefix(path, k+"/") {
			list = append(list, k)
		}
	}
	switch len(list) {
	case 0:
	case 1:
		v := list[0]
		dir, found = filepath.Join(d.pkgs[v], path[len(v+"/"):]), true
	default:
		// check path/v2
		sort.Slice(list, func(i, j int) bool {
			return list[i] > list[j]
		})
		v := list[0]
		dir, found = filepath.Join(d.pkgs[v], path[len(v+"/"):]), true
	}
	return
}

// Parse parse deps by go list
func (d *ListDriver) Parse(root string) error {
	args := []string{"list", "-deps", "-e", "-f={{.ImportPath}}={{.Dir}}"}
	if BuildMod != "" {
		args = append(args, "-mod", BuildMod)
	}
	data, err := runGoCommand(root, args...)
	if err != nil {
		return err
	}
	d.pkgs = make(map[string]string)
	for _, line := range strings.Split(string(data), "\n") {
		pos := strings.Index(line, "=")
		if pos != -1 {
			d.pkgs[line[:pos]] = line[pos+1:]
		}
	}
	return nil
}

func runGoCommand(dir string, args ...string) (ret []byte, err error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = dir
	err = cmd.Run()
	if err == nil {
		ret = stdout.Bytes()
	} else if stderr.Len() > 0 {
		err = errors.New(stderr.String())
	}
	return
}
