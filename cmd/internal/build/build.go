/*
 Copyright 2021 The GoPlus Authors (goplus.org)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

// Package build implements the “ixgo build” command.
package build

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/cmd/internal/base"
	"github.com/goplus/ixgo/cmd/internal/flags"
	"github.com/goplus/ixgo/cmd/internal/load"
	"golang.org/x/tools/go/ssa"
)

// -----------------------------------------------------------------------------

// Cmd - ixgo build
var Cmd = &base.Command{
	UsageLine: "ixgo build [build flags] [package]",
	Short:     "compile a Go/XGo package",
}

var (
	flag = &Cmd.Flag
)

func init() {
	Cmd.Run = buildCmd
	flags.AddBuildFlags(Cmd, flags.OmitModFlag|flags.OmitSSAFlag|flags.OmitVFlag)
}

func buildCmd(cmd *base.Command, args []string) {
	err := flag.Parse(args)
	if err != nil {
		os.Exit(2)
	}
	paths := flag.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}
	path := paths[0]
	var mode ixgo.Mode
	if flags.BuildSSA {
		mode |= ixgo.EnableDumpInstr
	}
	if flags.BuildX {
		mode |= ixgo.EnableDumpImports
	}
	ctx := ixgo.NewContext(mode)
	ctx.BuildContext = flags.BuildContext
	path, _ = filepath.Abs(path)
	isDir, err := load.IsDir(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	var pkg *ssa.Package
	if isDir {
		if load.SupportGop && load.IsGopProject(path) {
			err := load.BuildGopDir(ctx, path)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		}
		pkg, err = ctx.LoadDir(path, false)
	} else {
		pkg, err = ctx.LoadFile(path, nil)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	_, err = ctx.NewInterp(pkg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if flags.BuildV {
		fmt.Println(pkg.Pkg.Path())
	}
}

// -----------------------------------------------------------------------------
