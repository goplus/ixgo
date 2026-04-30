package mod

import (
	"fmt"

	"golang.org/x/mod/modfile"
)

// ParseModFile parse go.mod
func ParseModFile(file string, data []byte) (*modfile.File, error) {
	f, err := modfile.Parse(file, data, fix)
	if err != nil {
		return nil, fmt.Errorf("failed to parse modfile: %w", err)
	}
	if f.Module == nil {
		return nil, fmt.Errorf("no module declaration in %s", file)
	}
	return f, nil
}

func fix(path, vers string) (resolved string, err error) {
	// do nothing
	return vers, nil
}
