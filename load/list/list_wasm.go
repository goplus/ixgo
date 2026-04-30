//go:build wasm
// +build wasm

package list

import "errors"

type ListDriver struct{}

func (d *ListDriver) Lookup(root string, path string) (dir string, found bool) {
	return
}

func GetImportPath(pkgName string, dir string) (string, error) {
	return "", errNotImplemented
}

var errNotImplemented = errors.New("not implemented")

var BuildMod string // no-op on WASM
