// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package fs

import (
	q "io/fs"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("io/fs", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "fs",
			Path: "io/fs",
			Deps: map[string]string{
				"errors":           "errors",
				"internal/bytealg": "bytealg",
				"internal/oserror": "oserror",
				"io":               "io",
				"path":             "path",
				"slices":           "slices",
				"time":             "time",
				"unicode/utf8":     "utf8",
			},
			Interfaces: map[string]reflect.Type{
				"DirEntry":    reflect.TypeOf((*q.DirEntry)(nil)).Elem(),
				"FS":          reflect.TypeOf((*q.FS)(nil)).Elem(),
				"File":        reflect.TypeOf((*q.File)(nil)).Elem(),
				"FileInfo":    reflect.TypeOf((*q.FileInfo)(nil)).Elem(),
				"GlobFS":      reflect.TypeOf((*q.GlobFS)(nil)).Elem(),
				"ReadDirFS":   reflect.TypeOf((*q.ReadDirFS)(nil)).Elem(),
				"ReadDirFile": reflect.TypeOf((*q.ReadDirFile)(nil)).Elem(),
				"ReadFileFS":  reflect.TypeOf((*q.ReadFileFS)(nil)).Elem(),
				"ReadLinkFS":  reflect.TypeOf((*q.ReadLinkFS)(nil)).Elem(),
				"StatFS":      reflect.TypeOf((*q.StatFS)(nil)).Elem(),
				"SubFS":       reflect.TypeOf((*q.SubFS)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"FileMode":    reflect.TypeOf((*q.FileMode)(nil)).Elem(),
				"PathError":   reflect.TypeOf((*q.PathError)(nil)).Elem(),
				"WalkDirFunc": reflect.TypeOf((*q.WalkDirFunc)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrClosed":     &q.ErrClosed,
				"ErrExist":      &q.ErrExist,
				"ErrInvalid":    &q.ErrInvalid,
				"ErrNotExist":   &q.ErrNotExist,
				"ErrPermission": &q.ErrPermission,
				"SkipAll":       &q.SkipAll,
				"SkipDir":       &q.SkipDir,
			},
			Funcs: map[string]interface{}{
				"FileInfoToDirEntry": q.FileInfoToDirEntry,
				"FormatDirEntry":     q.FormatDirEntry,
				"FormatFileInfo":     q.FormatFileInfo,
				"Glob":               q.Glob,
				"Lstat":              q.Lstat,
				"ReadDir":            q.ReadDir,
				"ReadFile":           q.ReadFile,
				"ReadLink":           q.ReadLink,
				"Stat":               q.Stat,
				"Sub":                q.Sub,
				"ValidPath":          q.ValidPath,
				"WalkDir":            q.WalkDir,
			},
			TypedConsts: map[string]interface{}{
				"ModeAppend":     q.ModeAppend,
				"ModeCharDevice": q.ModeCharDevice,
				"ModeDevice":     q.ModeDevice,
				"ModeDir":        q.ModeDir,
				"ModeExclusive":  q.ModeExclusive,
				"ModeIrregular":  q.ModeIrregular,
				"ModeNamedPipe":  q.ModeNamedPipe,
				"ModePerm":       q.ModePerm,
				"ModeSetgid":     q.ModeSetgid,
				"ModeSetuid":     q.ModeSetuid,
				"ModeSocket":     q.ModeSocket,
				"ModeSticky":     q.ModeSticky,
				"ModeSymlink":    q.ModeSymlink,
				"ModeTemporary":  q.ModeTemporary,
				"ModeType":       q.ModeType,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
