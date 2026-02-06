package xgofsys

import (
	"bytes"
	"io"
	"io/fs"
	"path"
	"time"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/fsys"
	"github.com/goplus/ixgo/xgobuild"
	"github.com/goplus/xgo/parser"
)

func SetBuildFileSystem(ctx *ixgo.Context, xctx *xgobuild.Context, filesystem parser.FileSystem, parseGoMod bool) error {
	err := fsys.SetBuildFileSystem(ctx, filesystem, parseGoMod)
	if err != nil {
		return err
	}
	readDir := ctx.BuildContext.ReadDir
	xgofiles := make(map[string][]byte)
	ctx.BuildContext.ReadDir = func(dir string) ([]fs.FileInfo, error) {
		files, err := readDir(dir)
		if err != nil {
			return nil, err
		}
		var skip bool
		for _, f := range files {
			if f.Name() == "xgo_autogen.go" {
				skip = true
				break
			}
		}
		if !skip {
			pkg, err := xctx.ParseFSDir(filesystem, dir)
			if err != nil {
				return nil, err
			}
			data, err := pkg.ToSource()
			if err != nil {
				return nil, err
			}
			file := path.Join(dir, "xgo_autogen.go")
			xgofiles[file] = data
			files = append(files, &fileinfo{file, len(data), 0666})
		}
		return files, nil
	}
	openFile := ctx.BuildContext.OpenFile
	ctx.BuildContext.OpenFile = func(path string) (io.ReadCloser, error) {
		if data, ok := xgofiles[path]; ok {
			return io.NopCloser(bytes.NewReader(data)), nil
		}
		return openFile(path)
	}
	return nil
}

type fileinfo struct {
	path string // unique path to the file or directory within a filesystem
	size int
	mode fs.FileMode
}

var _ fs.FileInfo = (*fileinfo)(nil)
var _ fs.DirEntry = (*fileinfo)(nil)

func (i *fileinfo) Name() string               { return path.Base(i.path) }
func (i *fileinfo) Size() int64                { return int64(i.size) }
func (i *fileinfo) Mode() fs.FileMode          { return i.mode }
func (i *fileinfo) Type() fs.FileMode          { return i.mode.Type() }
func (i *fileinfo) ModTime() time.Time         { return time.Time{} }
func (i *fileinfo) IsDir() bool                { return i.mode&fs.ModeDir != 0 }
func (i *fileinfo) Sys() any                   { return nil }
func (i *fileinfo) Info() (fs.FileInfo, error) { return i, nil }
