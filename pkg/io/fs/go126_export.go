// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

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
			"ModeAppend":     {Typ: reflect.TypeOf(q.ModeAppend), Value: constant.MakeInt64(int64(q.ModeAppend))},
			"ModeCharDevice": {Typ: reflect.TypeOf(q.ModeCharDevice), Value: constant.MakeInt64(int64(q.ModeCharDevice))},
			"ModeDevice":     {Typ: reflect.TypeOf(q.ModeDevice), Value: constant.MakeInt64(int64(q.ModeDevice))},
			"ModeDir":        {Typ: reflect.TypeOf(q.ModeDir), Value: constant.MakeInt64(int64(q.ModeDir))},
			"ModeExclusive":  {Typ: reflect.TypeOf(q.ModeExclusive), Value: constant.MakeInt64(int64(q.ModeExclusive))},
			"ModeIrregular":  {Typ: reflect.TypeOf(q.ModeIrregular), Value: constant.MakeInt64(int64(q.ModeIrregular))},
			"ModeNamedPipe":  {Typ: reflect.TypeOf(q.ModeNamedPipe), Value: constant.MakeInt64(int64(q.ModeNamedPipe))},
			"ModePerm":       {Typ: reflect.TypeOf(q.ModePerm), Value: constant.MakeInt64(int64(q.ModePerm))},
			"ModeSetgid":     {Typ: reflect.TypeOf(q.ModeSetgid), Value: constant.MakeInt64(int64(q.ModeSetgid))},
			"ModeSetuid":     {Typ: reflect.TypeOf(q.ModeSetuid), Value: constant.MakeInt64(int64(q.ModeSetuid))},
			"ModeSocket":     {Typ: reflect.TypeOf(q.ModeSocket), Value: constant.MakeInt64(int64(q.ModeSocket))},
			"ModeSticky":     {Typ: reflect.TypeOf(q.ModeSticky), Value: constant.MakeInt64(int64(q.ModeSticky))},
			"ModeSymlink":    {Typ: reflect.TypeOf(q.ModeSymlink), Value: constant.MakeInt64(int64(q.ModeSymlink))},
			"ModeTemporary":  {Typ: reflect.TypeOf(q.ModeTemporary), Value: constant.MakeInt64(int64(q.ModeTemporary))},
			"ModeType":       {Typ: reflect.TypeOf(q.ModeType), Value: constant.MakeInt64(int64(q.ModeType))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
