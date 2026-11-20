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

// Package run implements the “gop run” command.
package run

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/cmd/internal/base"
	"github.com/goplus/ixgo/cmd/internal/flags"
	"github.com/goplus/ixgo/cmd/internal/load"
	"golang.org/x/tools/go/ssa"
)

// -----------------------------------------------------------------------------

// Cmd - ixgo run
var Cmd = &base.Command{
	UsageLine: "ixgo run [build flags] [package] [arguments...]",
	Short:     "run a Go/XGo package",
}

var (
	flag = &Cmd.Flag
)

func init() {
	Cmd.Run = runCmd
	flags.AddBuildFlags(Cmd, flags.OmitModFlag|flags.OmitSSAFlag|flags.OmitSSATraceFlag|
		flags.OmitVFlag|flags.OmitExperimentalGCFlag)
}

func runCmd(cmd *base.Command, args []string) {
	err := flag.Parse(args)
	if err != nil {
		os.Exit(2)
	}
	paths := flag.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}
	args = paths[1:]
	path, _ := filepath.Abs(paths[0])
	isDir, err := load.IsDir(path)
	if err != nil {
		log.Fatalln("input arg check failed:", err)
	}
	var mode ixgo.Mode
	if flags.BuildSSA {
		mode |= ixgo.EnableDumpInstr
	}
	if flags.DebugSSATrace {
		mode |= ixgo.EnableTracing
	}
	if flags.BuildX {
		mode |= ixgo.EnableDumpImports
	}
	if flags.ExperimentalGC {
		mode |= ixgo.ExperimentalSupportGC
	}
	ctx := ixgo.NewContext(mode)
	ctx.BuildContext = flags.BuildContext
	ctx.RunContext = context.TODO()
	var pkg *ssa.Package
	var input string
	if isDir {
		if load.SupportGop && load.IsGopProject(path) {
			err := load.BuildGopDir(ctx, path)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		}
		pkg, err = ctx.LoadDir(path, false)
		input = path
	} else {
		pkg, err = ctx.LoadFile(path, nil)
		input, _ = filepath.Split(path)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	interp, err := ctx.NewInterp(pkg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if flags.BuildV {
		fmt.Println(pkg.Pkg.Path())
	}
	code, err := ctx.RunInterp(interp, input, args)
	if err != nil {
		if e, ok := err.(ixgo.PanicError); ok {
			fmt.Fprintf(os.Stderr, "panic: %v\n\n%s\n", e.Error(), e.Stack())
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(2)
	}
	os.Exit(code)
}
