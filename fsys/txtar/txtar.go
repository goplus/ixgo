// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package txtar

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"strings"

	"golang.org/x/tools/txtar"
)

// fileSet is a set of files.
// The zero value for fileSet is an empty set ready to use.
type FileSet struct {
	ProgName string            // default prog name
	Files    []string          // filenames in user-provided order
	M        map[string][]byte // filename -> source
	NoHeader bool              // whether the prog.go entry was implicit
}

// Data returns the content of the named file.
// The fileSet retains ownership of the returned slice.
func (fs *FileSet) Data(filename string) []byte { return fs.M[filename] }

// Num returns the number of files in the set.
func (fs *FileSet) Num() int { return len(fs.M) }

// Contains reports whether fs contains the given filename.
func (fs *FileSet) Contains(filename string) bool {
	_, ok := fs.M[filename]
	return ok
}

// AddFile adds a file to fs. If fs already contains filename, its
// contents are replaced.
func (fs *FileSet) AddFile(filename string, src []byte) {
	had := fs.Contains(filename)
	if fs.M == nil {
		fs.M = make(map[string][]byte)
	}
	fs.M[filename] = src
	if !had {
		fs.Files = append(fs.Files, filename)
	}
}

func (fs *FileSet) Update(filename string, src []byte) {
	if fs.Contains(filename) {
		fs.M[filename] = src
	}
}

func (fs *FileSet) MvFile(source, target string) {
	if fs.M == nil {
		return
	}
	data, ok := fs.M[source]
	if !ok {
		return
	}
	fs.M[target] = data
	delete(fs.M, source)
	for i := range fs.Files {
		if fs.Files[i] == source {
			fs.Files[i] = target
			break
		}
	}
}

// Format returns fs formatted as a txtar archive.
func (fs *FileSet) Format() []byte {
	a := new(txtar.Archive)
	if fs.NoHeader {
		a.Comment = fs.M[fs.ProgName]
	}
	for i, f := range fs.Files {
		if i == 0 && f == fs.ProgName && fs.NoHeader {
			continue
		}
		a.Files = append(a.Files, txtar.File{Name: f, Data: fs.M[f]})
	}
	return txtar.Format(a)
}

// splitFiles splits the user's input program src into 1 or more
// files, splitting it based on boundaries as specified by the "txtar"
// format. It returns an error if any filenames are bogus or
// duplicates. The implicit filename for the txtar comment (the lines
// before any txtar separator line) are named "prog.go". It is an
// error to have an explicit file named "prog.go" in addition to
// having the implicit "prog.go" file (non-empty comment section).
//
// The filenames are validated to only be relative paths, not too
// long, not too deep, not have ".." elements, not have backslashes or
// low ASCII binary characters, and to be in path.Clean canonical
// form.
//
// splitFiles takes ownership of src.
func SplitFiles(src []byte, progName string) (*FileSet, error) {
	fs := &FileSet{ProgName: progName}
	a := txtar.Parse(src)
	if v := bytes.TrimSpace(a.Comment); len(v) > 0 {
		fs.NoHeader = true
		fs.AddFile(fs.ProgName, a.Comment)
	}
	const limitNumFiles = 20 // arbitrary
	numFiles := len(a.Files) + fs.Num()
	if numFiles > limitNumFiles {
		return nil, fmt.Errorf("too many files in txtar archive (%v exceeds limit of %v)", numFiles, limitNumFiles)
	}
	for _, f := range a.Files {
		if len(f.Name) > 200 { // arbitrary limit
			return nil, errors.New("file name too long")
		}
		if strings.IndexFunc(f.Name, isBogusFilenameRune) != -1 {
			return nil, fmt.Errorf("invalid file name %q", f.Name)
		}
		if f.Name != path.Clean(f.Name) || path.IsAbs(f.Name) {
			return nil, fmt.Errorf("invalid file name %q", f.Name)
		}
		parts := strings.Split(f.Name, "/")
		if len(parts) > 10 { // arbitrary limit
			return nil, fmt.Errorf("file name %q too deep", f.Name)
		}
		for _, part := range parts {
			if part == "." || part == ".." {
				return nil, fmt.Errorf("invalid file name %q", f.Name)
			}
		}
		if fs.Contains(f.Name) {
			return nil, fmt.Errorf("duplicate file name %q", f.Name)
		}
		fs.AddFile(f.Name, f.Data)
	}
	return fs, nil
}

// isBogusFilenameRune reports whether r should be rejected if it
// appears in a txtar section's filename.
func isBogusFilenameRune(r rune) bool { return r == '\\' || r < ' ' }
