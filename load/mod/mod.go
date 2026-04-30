package mod

import (
	"errors"
	"fmt"

	"golang.org/x/mod/modfile"
)

// ParseModFile parse go.mod
func ParseModFile(file string, data []byte) (*modfile.File, error) {
	fix := func(path, vers string) (resolved string, err error) {
		// do nothing
		return vers, nil
	}
	f, err := modfile.Parse(file, data, fix)
	if err != nil {
		return nil, fmt.Errorf("parse go.mod %w", err)
	}
	if f.Module == nil {
		return nil, errors.New("no module declaration in go.mod")
	}
	return f, nil
}
