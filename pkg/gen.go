//go:build ignore
// +build ignore

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

package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var (
	gover    = runtime.Version()[:6]
	minorVer int
)

func init() {
	var err error
	minorVer, err = strconv.Atoi(gover[4:])
	if err != nil {
		panic("error version: " + gover)
	}
}

func main() {
	var tags string
	var name string
	var fname string
	switch gover {
	case "go1.25":
		tags = "//+build go1.25"
		name = "go125_export"
		fname = "go125_pkgs.go"
	case "go1.24":
		tags = "//+build go1.24,!go1.25"
		name = "go124_export"
		fname = "go124_pkgs.go"
	case "go1.23":
		tags = "//+build go1.23,!go1.24"
		name = "go123_export"
		fname = "go123_pkgs.go"
	case "go1.22":
		tags = "//+build go1.22,!go1.23"
		name = "go122_export"
		fname = "go122_pkgs.go"
	case "go1.21":
		tags = "//+build go1.21,!go1.22"
		name = "go121_export"
		fname = "go121_pkgs.go"
	case "go1.20":
		tags = "//+build go1.20,!go1.21"
		name = "go120_export"
		fname = "go120_pkgs.go"
	case "go1.19":
		tags = "//+build go1.19,!go1.20"
		name = "go119_export"
		fname = "go119_pkgs.go"
	case "go1.18":
		tags = "//+build go1.18,!go1.19"
		name = "go118_export"
		fname = "go118_pkgs.go"
	case "go1.17":
		tags = "//+build go1.17,!go1.18"
		name = "go117_export"
		fname = "go117_pkgs.go"
	case "go1.16":
		tags = "//+build go1.16,!go1.17"
		name = "go116_export"
		fname = "go116_pkgs.go"
	case "go1.15":
		tags = "//+build go1.15,!go1.16"
		name = "go115_export"
		fname = "go115_pkgs.go"
	case "go1.14":
		tags = "//+build go1.14,!go1.15"
		name = "go114_export"
		fname = "go114_pkgs.go"
	}

	pkgs := stdList()

	log.Println(gover, name, tags)
	log.Println(pkgs)

	cmd := exec.Command("qexp", "-outdir", ".", "-addtags", tags, "-filename", name)
	cmd.Args = append(cmd.Args, normalPkgs(pkgs)...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	// log
	cmd = exec.Command("qexp", "-outdir", ".", "-addtags", tags, "-filename", name, "-code")
	cmd.Args = append(cmd.Args, "log")
	if minorVer >= 21 {
		cmd.Args = append(cmd.Args, "log/slog", "log/internal", "log/slog/internal", "log/slog/internal/buffer")
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// generic pkgs
	gpkgs := genericPkgs(pkgs)
	if len(gpkgs) > 0 {
		cmd = exec.Command("qexp", "-outdir", ".", "-addtags", tags, "-filename", name, "-src")
		cmd.Args = append(cmd.Args, gpkgs...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}
	// syscall
	list := osarchList()
	log.Println(list)
	for _, osarch := range list {
		ar := strings.Split(osarch, "_")
		if len(ar) != 2 {
			continue
		}
		cmd := exec.Command("qexp", "-outdir", ".", "-addtags", tags, "-filename", name, "-contexts", osarch, "syscall")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GOOS="+ar[0])
		cmd.Env = append(cmd.Env, "GOARCH="+ar[1])
		cmd.Env = append(cmd.Env, "GOEXPERIMENT=noregabi")
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		err = makepkg("./"+fname, []string{tags}, pkgs)
		if err != nil {
			panic(err)
		}
	}

	if minorVer >= 23 {
		for _, pkg := range []string{"iter"} {
			log.Printf("export %v: %v patch", pkg, pkg+"/"+fname)
			data, err := os.ReadFile("./_" + strings.Replace(gover, ".", "", -1) + "/" + pkg + "_export.go")
			if err != nil {
				panic(err)
			}
			err = os.WriteFile("./"+pkg+"/"+name+".go", data, 0666)
			if err != nil {
				panic(err)
			}
		}
	}
}

func normalPkgs(std []string) (pkgs []string) {
	for _, pkg := range std {
		switch pkg {
		case "syscall", "log", "log/slog":
			continue
		default:
			if isGeneric(pkg) {
				continue
			}
		}
		pkgs = append(pkgs, pkg)
	}
	return
}

func isGeneric(pkg string) bool {
	switch pkg {
	case "maps", "slices", "cmp", "iter":
		return true
	}
	return false
}

func genericPkgs(std []string) (pkgs []string) {
	for _, pkg := range std {
		if isGeneric(pkg) {
			pkgs = append(pkgs, pkg)
		}
	}
	return
}

func makepkg(fname string, tags []string, std []string) error {
	//_ github.com/goplus/ixgo/pkg
	var pkgs []string
	for _, v := range std {
		if strings.HasPrefix(v, "testing/") {
			continue
		}
		if v == "testing" || v == "log/syslog" || v == "net/rpc" || v == "net/rpc/jsonrpc" {
			continue
		}
		pkgs = append(pkgs, fmt.Sprintf(`_ "github.com/goplus/ixgo/pkg/%v"`, v))
	}
	pkgs = append(pkgs, fmt.Sprintf(`_ "github.com/goplus/ixgo/pkg/syscall"`))
	r := strings.NewReplacer("$TAGS", strings.Join(tags, "\n"), "$PKGS", strings.Join(pkgs, "\t\n"))
	src := r.Replace(tmpl)
	data, err := format.Source([]byte(src))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fname, []byte(data), 0644)
}

var tmpl = `$TAGS

package pkg

import (
	$PKGS
)
`

// skip syscall
func stdList() []string {
	cmd := exec.Command("go", "list", "std")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	var ar []string
	for _, v := range strings.Split(string(data), "\n") {
		if v == "" {
			continue
		}
		if isSkipPkg(v) {
			continue
		}
		ar = append(ar, v)
	}
	return ar
}

func isSkipPkg(pkg string) bool {
	switch pkg {
	case "syscall", "unsafe":
		return true
	case "runtime/cgo", "runtime/race":
		return true
	case "plugin":
		if minorVer >= 23 {
			return true
		}
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

// aix_ppc64 android_386 android_amd64
func checkRegAbi(list []string, ver string) (regabi []string, noregabi []string) {
	for _, v := range list {
		ar := strings.Split(v, "_")
		if len(ar) != 2 {
			continue
		}
		switch ver {
		case "go1.17":
			if ar[1] == "amd64" {
				regabi = append(regabi, v)
				continue
			}
		case "go1.18":
			switch ar[1] {
			case "amd64", "arm64", "ppc64", "ppc64le":
				regabi = append(regabi, v)
				continue
			}
		case "go1.19":
			switch ar[1] {
			case "amd64", "arm64", "ppc64", "ppc64le", "riscv64":
				regabi = append(regabi, v)
				continue
			}
		}
		noregabi = append(noregabi, v)
	}
	return
}

func osarchList() []string {
	//go tool dist list
	cmd := exec.Command("go", "tool", "dist", "list")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	var ar []string
	for _, v := range strings.Split(string(data), "\n") {
		if v == "" {
			continue
		}
		ar = append(ar, strings.Replace(v, "/", "_", 1))
	}
	return ar
}
