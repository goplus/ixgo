package fsys_test

import (
	"testing"

	"github.com/goplus/ixgo/xgobuild"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/fsys"
	"github.com/goplus/ixgo/fsys/txtar"
	_ "github.com/goplus/ixgo/pkg/fmt"
)

var data = `
package main

import (
	"play.ground/foo"
)

func main() {
	foo.Bar()
}
-- go.mod --
module play.ground
-- foo/foo.go --
package foo

import "fmt"
import "play.ground/foo/bar"

func Bar() {
	bar.Demo()
	fmt.Println("this is a foo!")
}
-- foo/bar/bar.go --
package bar

import "fmt"

func Demo() {
	fmt.Println("this is a bar!")
}
`
var xgo_data = `
import (
	"play.ground/foo"
)

echo "xgo"
foo.Bar()
-- go.mod --
module play.ground
-- foo/foo.go --
package foo

import "fmt"
import "play.ground/foo/bar"

func Bar() {
	bar.Demo()
	fmt.Println("this is a foo!")
}
-- foo/bar/bar.go --
package bar

import "fmt"

func Demo() {
	fmt.Println("this is a bar!")
}
`

func TestFS(t *testing.T) {
	fset, err := txtar.SplitFiles([]byte(data), "prog.go")
	if err != nil {
		t.Fatal(err)
	}
	ctx := ixgo.NewContext(0)
	fs, err := txtar.FS(fset, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = fsys.SetBuildFS(ctx, fs, true)
	if err != nil {
		t.Fatal(err)
	}

	pkg, err := ctx.LoadDirEx(".", false, func(pkgName string, dir string) (string, bool) {
		return pkgName, true
	})
	if err != nil {
		t.Fatal(err)
	}
	code, _ := ctx.RunPkg(pkg, "", nil)
	if code != 0 {
		t.Fatal("error")
	}
}

func TestFileSystem(t *testing.T) {
	fset, err := txtar.SplitFiles([]byte(data), "prog.go")
	if err != nil {
		t.Fatal(err)
	}
	ctx := ixgo.NewContext(0)
	fs, err := txtar.FileSystem(fset, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = fsys.SetBuildFileSystem(ctx, fs, true)
	if err != nil {
		t.Fatal(err)
	}

	pkg, err := ctx.LoadDirEx(".", false, func(pkgName string, dir string) (string, bool) {
		return pkgName, true
	})
	if err != nil {
		t.Fatal(err)
	}
	code, _ := ctx.RunPkg(pkg, "", nil)
	if code != 0 {
		t.Fatal("error")
	}
}

func TestXGoFileSystem(t *testing.T) {
	fset, err := txtar.SplitFiles([]byte(data), "prog.xgo")
	if err != nil {
		t.Fatal(err)
	}
	ctx := ixgo.NewContext(0)
	fs, err := txtar.FileSystem(fset, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = fsys.SetBuildFileSystem(ctx, fs, true)
	if err != nil {
		t.Fatal(err)
	}

	data, err := xgobuild.BuildFSDir(ctx, fs, ".")
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := ctx.LoadFile("main.go", data)
	if err != nil {
		t.Fatal(err)
	}

	code, _ := ctx.RunPkg(pkg, "", nil)
	if code != 0 {
		t.Fatal("error")
	}
}
