// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package fs

import (
	q "io/fs"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"ErrClosed":     reflect.ValueOf(&q.ErrClosed),
			"ErrExist":      reflect.ValueOf(&q.ErrExist),
			"ErrInvalid":    reflect.ValueOf(&q.ErrInvalid),
			"ErrNotExist":   reflect.ValueOf(&q.ErrNotExist),
			"ErrPermission": reflect.ValueOf(&q.ErrPermission),
			"SkipAll":       reflect.ValueOf(&q.SkipAll),
			"SkipDir":       reflect.ValueOf(&q.SkipDir),
		},
		Funcs: map[string]reflect.Value{
			"FileInfoToDirEntry": reflect.ValueOf(q.FileInfoToDirEntry),
			"FormatDirEntry":     reflect.ValueOf(q.FormatDirEntry),
			"FormatFileInfo":     reflect.ValueOf(q.FormatFileInfo),
			"Glob":               reflect.ValueOf(q.Glob),
			"Lstat":              reflect.ValueOf(q.Lstat),
			"ReadDir":            reflect.ValueOf(q.ReadDir),
			"ReadFile":           reflect.ValueOf(q.ReadFile),
			"ReadLink":           reflect.ValueOf(q.ReadLink),
			"Stat":               reflect.ValueOf(q.Stat),
			"Sub":                reflect.ValueOf(q.Sub),
			"ValidPath":          reflect.ValueOf(q.ValidPath),
			"WalkDir":            reflect.ValueOf(q.WalkDir),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"ModeAppend":     {reflect.TypeOf(q.ModeAppend), constant.MakeInt64(int64(q.ModeAppend))},
			"ModeCharDevice": {reflect.TypeOf(q.ModeCharDevice), constant.MakeInt64(int64(q.ModeCharDevice))},
			"ModeDevice":     {reflect.TypeOf(q.ModeDevice), constant.MakeInt64(int64(q.ModeDevice))},
			"ModeDir":        {reflect.TypeOf(q.ModeDir), constant.MakeInt64(int64(q.ModeDir))},
			"ModeExclusive":  {reflect.TypeOf(q.ModeExclusive), constant.MakeInt64(int64(q.ModeExclusive))},
			"ModeIrregular":  {reflect.TypeOf(q.ModeIrregular), constant.MakeInt64(int64(q.ModeIrregular))},
			"ModeNamedPipe":  {reflect.TypeOf(q.ModeNamedPipe), constant.MakeInt64(int64(q.ModeNamedPipe))},
			"ModePerm":       {reflect.TypeOf(q.ModePerm), constant.MakeInt64(int64(q.ModePerm))},
			"ModeSetgid":     {reflect.TypeOf(q.ModeSetgid), constant.MakeInt64(int64(q.ModeSetgid))},
			"ModeSetuid":     {reflect.TypeOf(q.ModeSetuid), constant.MakeInt64(int64(q.ModeSetuid))},
			"ModeSocket":     {reflect.TypeOf(q.ModeSocket), constant.MakeInt64(int64(q.ModeSocket))},
			"ModeSticky":     {reflect.TypeOf(q.ModeSticky), constant.MakeInt64(int64(q.ModeSticky))},
			"ModeSymlink":    {reflect.TypeOf(q.ModeSymlink), constant.MakeInt64(int64(q.ModeSymlink))},
			"ModeTemporary":  {reflect.TypeOf(q.ModeTemporary), constant.MakeInt64(int64(q.ModeTemporary))},
			"ModeType":       {reflect.TypeOf(q.ModeType), constant.MakeInt64(int64(q.ModeType))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
