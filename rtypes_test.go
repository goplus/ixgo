package ixgo

import (
	"go/constant"
	"go/types"
	"reflect"
	"testing"

	"github.com/goplus/ixgo/testdata/getpkg"
)

type getpkgFieldHolder struct {
	T getpkg.T
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
