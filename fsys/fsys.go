package fsys

import (
	"bytes"
	"io"
	"io/fs"
	"path"
	"strings"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/load"
)

// FileSystem represents a file system.
type FileSystem interface {
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

func SetBuildFileSystem(ctx *ixgo.Context, fsys FileSystem, parseGoMod bool) error {
	if parseGoMod {
		if data, err := fsys.ReadFile("go.mod"); err == nil {
			f, err := load.ParseModFile("go.mod", data)
			if err != nil {
				return err
			}
			modpath := f.Module.Mod.Path
			ctx.Lookup = func(root, path string) (string, bool) {
				if path == modpath {
					return ".", true
				}
				if strings.HasPrefix(path, modpath+"/") {
					return path[len(modpath+"/"):], true
				}
				return "", false
			}
		}
	}
	ctx.BuildContext.JoinPath = fsys.Join
	alldirs := make(map[string]bool)
	ctx.BuildContext.IsDir = func(path string) bool {
		if path == "." || alldirs[path] {
			return true
		}
		if _, err := fsys.ReadDir(path); err == nil {
			alldirs[path] = true
			return true
		}
		return false
	}
	ctx.BuildContext.ReadDir = func(root string) (infos []fs.FileInfo, err error) {
		dirs, err := fsys.ReadDir(root)
		if err != nil {
			return nil, err
		}
		for _, dir := range dirs {
			if dir.IsDir() {
				alldirs[fsys.Join(root, dir.Name())] = true
				continue
			}
			info, err := dir.Info()
			if err != nil {
				continue
			}
			infos = append(infos, info)
		}
		return
	}
	ctx.BuildContext.OpenFile = func(name string) (io.ReadCloser, error) {
		data, err := fsys.ReadFile(name)
		if err != nil {
			return nil, err
		}
		return io.NopCloser(bytes.NewReader(data)), nil
	}
	return nil
}

func SetBuildFS(ctx *ixgo.Context, fsys fs.FS, parseGoMod bool) error {
	if parseGoMod {
		if file, err := fsys.Open("go.mod"); err == nil {
			data, err := io.ReadAll(file)
			if err != nil {
				return err
			}
			f, err := load.ParseModFile("go.mod", data)
			if err != nil {
				return err
			}
			modpath := f.Module.Mod.Path
			ctx.Lookup = func(root, path string) (string, bool) {
				if path == modpath {
					return ".", true
				}
				if strings.HasPrefix(path, modpath+"/") {
					return path[len(modpath+"/"):], true
				}
				return "", false
			}
		}
	}
	ctx.BuildContext.JoinPath = func(elem ...string) string {
		return path.Join(elem...)
	}
	alldirs := make(map[string]bool)
	ctx.BuildContext.IsDir = func(path string) bool {
		if path == "." || alldirs[path] {
			return true
		}
		if _, err := fs.ReadDir(fsys, path); err == nil {
			alldirs[path] = true
			return true
		}
		return false
	}
	ctx.BuildContext.ReadDir = func(root string) (infos []fs.FileInfo, err error) {
		dirs, err := fs.ReadDir(fsys, root)
		if err != nil {
			return nil, err
		}
		for _, dir := range dirs {
			if dir.IsDir() {
				alldirs[path.Join(root, dir.Name())] = true
				continue
			}
			info, err := dir.Info()
			if err != nil {
				continue
			}
			infos = append(infos, info)
		}
		return
	}
	ctx.BuildContext.OpenFile = func(name string) (io.ReadCloser, error) {
		return fsys.Open(name)
	}
	return nil
}
