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
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goplus/ixgo/cmd/internal/base"
)

var (
	flagExportDir      string
	flagCustomPkg      string
	flagBuildContext   string
	flagCustomTags     string
	flagBuildTags      string
	flagExportFileName string
	flagExportSource   bool
	flagExportCode     bool
)

func init() {
	flag.StringVar(&flagExportDir, "outdir", "", "set export out dir")
	flag.StringVar(&flagCustomPkg, "pkgpath", "", "set custom pkg path")
	flag.StringVar(&flagBuildContext, "contexts", "", "set custom build contexts goos_goarch list. eg \"drawin_amd64 darwin_arm64\"")
	flag.StringVar(&flagCustomTags, "addtags", "", "add custom tags, split by ;")
	flag.StringVar(&flagBuildTags, "tags", "", "a comma-separated list of build tags")
	flag.StringVar(&flagExportFileName, "filename", "export", "set export file name")
	flag.BoolVar(&flagExportSource, "src", false, "export generic source mode")
	flag.BoolVar(&flagExportCode, "code", false, "export full source code mode")
}

// Cmd - ixgo build
var Cmd = &base.Command{
	UsageLine: "ixgo export [flags] [package]",
	Short:     "export Go package to ixgo builtin package",
}

var (
	flag = &Cmd.Flag
)

func init() {
	Cmd.Run = exportCmd
}

func exportCmd(cmd *base.Command, args []string) {
	err := flag.Parse(args)
	if err != nil {
		os.Exit(2)
	}
	args = flag.Args()
	if len(args) == 0 {
		cmd.Usage(os.Stderr)
	}
	if len(args) == 1 {
		pkgs, err := parserPkgs(args[0])
		if err != nil {
			log.Panicln(err)
		}
		if len(pkgs) == 0 {
			log.Panicln("not found packages")
		}
		args = pkgs
	}
	if flagExportFileName == "" {
		flagExportFileName = "export"
	}
	ctxList := parserContextList(flagBuildContext)
	Export(args, ctxList)
}

// pkg/...
func parserPkgs(expr string) ([]string, error) {
	cmds := []string{"list", "-f={{.ImportPath}}={{.Name}}", expr}
	if flagBuildTags != "" {
		cmds = []string{"list", "-tags", flagBuildTags, "-f={{.ImportPath}}={{.Name}}", expr}
	}
	data, err := runGoCommand(cmds...)
	if err != nil {
		return nil, err
	}
	var pkgs []string
	for _, line := range strings.Split(string(data), "\n") {
		pos := strings.Index(line, "=")
		if pos != -1 {
			pkg, name := line[:pos], line[pos+1:]
			if name == "main" {
				continue
			}
			pkgs = append(pkgs, pkg)
		}
	}
	if len(pkgs) <= 1 {
		return pkgs, nil
	}
	var rets []string
	for _, pkg := range pkgs {
		if isSkipPkg(pkg) {
			continue
		}
		rets = append(rets, pkg)
	}
	return rets, nil
}

func runGoCommand(args ...string) (ret []byte, err error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err == nil {
		ret = stdout.Bytes()
	} else if stderr.Len() > 0 {
		err = errors.New(stderr.String())
	}
	return
}

func isSkipPkg(pkg string) bool {
	switch pkg {
	case "unsafe":
		return true
	default:
		if strings.HasPrefix(pkg, "vendor/") {
			return true
		}
		for _, v := range strings.Split(pkg, "/") {
			if v == "internal" {
				return true
			}
		}
	}
	return false
}

func Export(pkgs []string, ctxList []*build.Context) {
	log.Println("process", pkgs)
	if len(ctxList) == 0 {
		ExportPkgs(pkgs, nil)
	} else {
		for _, ctx := range ctxList {
			ExportPkgs(pkgs, ctx)
		}
	}
}

func ExportPkgs(pkgs []string, ctx *build.Context) {
	prog := NewProgram(ctx)
	err := prog.Load(pkgs)
	if err != nil {
		log.Panicln(err)
	}
	for _, pkg := range pkgs {
		if pkg == "unsafe" {
			continue
		}
		fpath, err := ExportPkg(prog, pkg, ctx)
		if err != nil {
			log.Printf("export %v failed: %v\n", pkg, err)
		} else {
			log.Printf("export %v: %v\n", pkg, fpath)
		}
	}
}

func ExportPkg(prog *Program, pkg string, ctx *build.Context) (string, error) {
	e, err := prog.ExportPkg(pkg, "q")
	if err != nil {
		return "", err
	}
	var tags []string
	if flagCustomTags != "" {
		tags = strings.Split(flagCustomTags, ";")
	}
	data, err := exportPkg(e, "q", "", tags)
	if err != nil {
		return "", err
	}
	if flagExportDir == "" {
		fmt.Println(string(data))
		return "stdout", nil
	}
	if flagCustomPkg != "" {
		pkg = flagCustomPkg
	}
	fpath := filepath.Join(flagExportDir, pkg)
	var fname string
	if ctx != nil {
		fname = flagExportFileName + "_" + ctx.GOOS + "_" + ctx.GOARCH + ".go"
	} else {
		fname = flagExportFileName + ".go"
	}
	err = writeFile(fpath, fname, data)
	return filepath.Join(fpath, fname), err
}

func parserContextList(list string) (ctxs []*build.Context) {
	for _, info := range strings.Split(list, " ") {
		info = strings.TrimSpace(info)
		ar := strings.Split(info, "_")
		if len(ar) != 2 {
			continue
		}
		ctx := build.Default
		ctx.GOOS = ar[0]
		ctx.GOARCH = ar[1]
		ctx.BuildTags = strings.Split(flagBuildTags, ",")
		ctxs = append(ctxs, &ctx)
	}
	return
}
