//go:build ignore
// +build ignore

// Command gen generates the optional standard-library direct-call adapters.
// Run it with one of the supported Go toolchains, for example:
//
//	$ go run ./directcall/gen.go 1.25
//
// The selected toolchain is also used by qexp to inspect the standard library.
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

var versions = map[string]struct {
	goBuild  string
	buildTag string
	file     string
}{
	"1.25": {goBuild: "//go:build go1.25 && !go1.26", buildTag: "// +build go1.25,!go1.26", file: "go125_directcall"},
	"1.26": {goBuild: "//go:build go1.26 && !go1.27", buildTag: "// +build go1.26,!go1.27", file: "go126_directcall"},
	"1.27": {goBuild: "//go:build go1.27", buildTag: "// +build go1.27", file: "go127_directcall"},
}

func skipped(pkg string) bool {
	if pkg == "log" || pkg == "log/slog" || pkg == "syscall" || pkg == "unsafe" || pkg == "plugin" {
		return true
	}
	if pkg == "runtime/cgo" || pkg == "runtime/race" {
		return true
	}
	for _, part := range strings.Split(pkg, "/") {
		if part == "internal" || strings.HasPrefix(pkg, "vendor/") {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run ./directcall/gen.go 1.25|1.26|1.27")
		os.Exit(2)
	}
	version := os.Args[1]
	info, ok := versions[version]
	if !ok {
		fmt.Fprintf(os.Stderr, "unsupported Go version %q\n", version)
		os.Exit(2)
	}
	root, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	goBin := filepath.Join(os.Getenv("HOME"), "golang", "go"+version, "bin", "go")
	if _, err := os.Stat(goBin); err != nil {
		panic(fmt.Errorf("Go toolchain not found: %s: %w", goBin, err))
	}

	list := exec.Command(goBin, "list", "std")
	data, err := list.Output()
	if err != nil {
		panic(err)
	}
	var pkgs []string
	for _, pkg := range strings.Fields(string(data)) {
		if !skipped(pkg) {
			pkgs = append(pkgs, pkg)
		}
	}
	sort.Strings(pkgs)

	cmd := exec.Command(goBin, "run", "./cmd/qexp",
		"-outdir", "directcall",
		"-filename", info.file,
		"-addtags", info.buildTag,
		"-directcalls-only",
		"-directcalls", "all")
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "PATH="+filepath.Dir(goBin)+string(os.PathListSeparator)+os.Getenv("PATH"))
	fmt.Printf("generating %s directcalls for %d packages\n", version, len(pkgs))
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	pkgs = filterEmptyPackages(info, pkgs)
	writeAggregate(info, pkgs)
}

func filterEmptyPackages(info struct{ goBuild, buildTag, file string }, pkgs []string) []string {
	active := make([]string, 0, len(pkgs))
	for _, pkg := range pkgs {
		filename := filepath.Join("directcall", pkg, info.file+".go")
		data, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		if bytes.Contains(data, []byte("map[string]ixgo.DirectCallAdapter{}")) {
			if err := os.Remove(filename); err != nil {
				panic(err)
			}
			continue
		}
		active = append(active, pkg)
	}
	return active
}

func writeAggregate(info struct{ goBuild, buildTag, file string }, pkgs []string) {
	var imports []string
	for _, pkg := range pkgs {
		imports = append(imports, fmt.Sprintf("\t_ %q", "github.com/goplus/ixgo/directcall/"+pkg))
	}
	src := info.goBuild + "\n" + info.buildTag + "\n\npackage directcall\n\nimport (\n" + strings.Join(imports, "\n") + "\n)\n"
	if err := os.WriteFile(filepath.Join("directcall", info.file+".go"), []byte(src), 0644); err != nil {
		panic(err)
	}
}
