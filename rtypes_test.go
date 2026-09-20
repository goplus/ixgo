package ixgo

import (
	"go/constant"
	"go/types"
	"reflect"
	"testing"

	foo "github.com/goplus/ixgo/testdata/foo/v2"
	"github.com/goplus/ixgo/testdata/getpkg"
	unknownpkg "github.com/goplus/ixgo/testdata/unknownpkg"
)

type getpkgFieldHolder struct {
	T getpkg.T
}

type getpkgCurpkgHolder struct {
	A foo.T
	B unknownpkg.U
}

func TestGetPackageLoadsRegistered(t *testing.T) {
	const path = "ixgo.test/getpkgreg"
	RegisterPackage(&Package{
		Name: "getpkgreg",
		Path: path,
		UntypedConsts: map[string]UntypedConst{
			"Answer": {Typ: "untyped int", Value: constant.MakeInt64(42)},
		},
	})
	loader := NewContext(0).Loader.(*TypesLoader)
	pkg := loader.GetPackage(path)
	if !pkg.Complete() {
		t.Fatalf("GetPackage(%q) returned an incomplete placeholder", path)
	}
	if pkg.Scope().Lookup("Answer") == nil {
		t.Fatal("registered const Answer was not installed")
	}
}

func TestGetPackageUnknownIsPlaceholder(t *testing.T) {
	loader := NewContext(0).Loader.(*TypesLoader)
	pkg := loader.GetPackage("ixgo.test/does-not-exist")
	if pkg.Complete() {
		t.Fatal("unknown package should remain a placeholder")
	}
}

func TestGetPackageLoadsRegisteredFromReflectField(t *testing.T) {
	const depPath = "github.com/goplus/ixgo/testdata/getpkg"
	RegisterPackage(&Package{
		Name: "getpkg",
		Path: depPath,
		NamedTypes: map[string]reflect.Type{
			"T": reflect.TypeOf((*getpkg.T)(nil)).Elem(),
		},
	})
	RegisterPackage(&Package{
		Name: "holder",
		Path: "ixgo.test/getpkgholder",
		NamedTypes: map[string]reflect.Type{
			"Holder": reflect.TypeOf((*getpkgFieldHolder)(nil)).Elem(),
		},
	})
	ctx := NewContext(0)
	if _, err := ctx.Loader.Import("ixgo.test/getpkgholder"); err != nil {
		t.Fatal(err)
	}
	pkg, err := ctx.Loader.Import(depPath)
	if err != nil {
		t.Fatal(err)
	}
	if !pkg.Complete() {
		t.Fatal("getpkg should be loaded from the registry, not left as a placeholder")
	}
	obj := pkg.Scope().Lookup("T")
	if obj == nil {
		t.Fatal("missing type T")
	}
	if _, ok := obj.Type().Underlying().(*types.Struct); !ok {
		t.Fatalf("T type = %T, want struct", obj.Type().Underlying())
	}
}

func TestGetPackagePreservesCurpkgAfterNestedImport(t *testing.T) {
	const (
		fooPath     = "github.com/goplus/ixgo/testdata/foo/v2"
		unknownPath = "github.com/goplus/ixgo/testdata/unknownpkg"
		holderPath  = "ixgo.test/curpkgholder"
	)
	RegisterPackage(&Package{
		Name: "foo",
		Path: fooPath,
		NamedTypes: map[string]reflect.Type{
			"T": reflect.TypeOf((*foo.T)(nil)).Elem(),
		},
	})
	RegisterPackage(&Package{
		Name: "holder",
		Path: holderPath,
		Deps: map[string]string{
			fooPath:     "foo",
			unknownPath: "othername",
		},
		NamedTypes: map[string]reflect.Type{
			"Holder": reflect.TypeOf((*getpkgCurpkgHolder)(nil)).Elem(),
		},
	})
	ctx := NewContext(0)
	if _, err := ctx.Loader.Import(holderPath); err != nil {
		t.Fatal(err)
	}
	loader := ctx.Loader.(*TypesLoader)
	pkg := loader.GetPackage(unknownPath)
	if got, want := pkg.Name(), "othername"; got != want {
		t.Fatalf("unknown package name = %q, want %q (curpkg.Deps lost after nested Import)", got, want)
	}
}

func TestGetPackageRegisteredImportErrorFallsBack(t *testing.T) {
	const path = "ixgo.test/badsrc"
	RegisterPackage(&Package{
		Name:   "badsrc",
		Path:   path,
		Source: "package badsrc\nfunc (",
	})
	loader := NewContext(0).Loader.(*TypesLoader)
	pkg := loader.GetPackage(path)
	if pkg.Complete() {
		t.Fatal("failed registered source load should fall back to a placeholder")
	}
	if got, want := pkg.Name(), "badsrc"; got != want {
		t.Fatalf("placeholder name = %q, want %q", got, want)
	}
}
