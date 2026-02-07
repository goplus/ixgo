package txtar

import (
	"io/fs"
	"path"
	"path/filepath"

	"github.com/goplus/ixgo/fsys/txtar/txtarfs"
	"golang.org/x/tools/txtar"
)

// FileSystem represents a file system.
type IFileSystem interface {
	ReadDir(dirname string) ([]fs.DirEntry, error)
	ReadFile(filename string) ([]byte, error)
	Join(elem ...string) string

	// Base returns the last element of path.
	// Trailing path separators are removed before extracting the last element.
	// If the path is empty, Base returns ".".
	// If the path consists entirely of separators, Base returns a single separator.
	Base(filename string) string

	// Abs returns an absolute representation of path.
	Abs(path string) (string, error)
}

func FileSystem(fs *FileSet, filter func(file string) bool) (IFileSystem, error) {
	a := new(txtar.Archive)
	for _, f := range fs.Files {
		if filter != nil && !filter(f) {
			continue
		}
		a.Files = append(a.Files, txtar.File{Name: f, Data: fs.M[f]})
	}
	fsys, err := txtarfs.FS(a)
	if err != nil {
		return nil, err
	}
	return &fileSystem{fs, fsys}, nil
}

func FS(fs *FileSet, filter func(file string) bool) (fs.FS, error) {
	a := new(txtar.Archive)
	for _, f := range fs.Files {
		if filter != nil && !filter(f) {
			continue
		}
		a.Files = append(a.Files, txtar.File{Name: f, Data: fs.M[f]})
	}
	return txtarfs.FS(a)
}

type fileSystem struct {
	Fset *FileSet
	Sys  fs.FS
}

func (f *fileSystem) ReadDir(dirname string) ([]fs.DirEntry, error) {
	return fs.ReadDir(f.Sys, dirname)
}

func (f *fileSystem) ReadFile(filename string) ([]byte, error) {
	return fs.ReadFile(f.Sys, filename)
}

func (f *fileSystem) Join(elem ...string) string {
	return path.Join(elem...)
}

func (f *fileSystem) Base(filename string) string {
	return path.Base(filename)
}

func (f *fileSystem) Abs(path string) (string, error) {
	return filepath.Abs(path)
}
