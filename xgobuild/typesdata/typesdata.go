package typesdata

import (
	"bytes"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/gcexportdata"
)

func ImportFunc(path string, data []byte) func(fset *token.FileSet, pkgs map[string]*types.Package) (*types.Package, error) {
	return func(fset *token.FileSet, pkgs map[string]*types.Package) (*types.Package, error) {
		return gcexportdata.Read(bytes.NewReader(data), fset, pkgs, path)
	}
}
