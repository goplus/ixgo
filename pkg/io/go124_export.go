// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package io

import (
	q "io"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("io", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "io",
			Path: "io",
			Deps: map[string]string{
				"errors": "errors",
				"sync":   "sync",
			},
			Interfaces: map[string]reflect.Type{
				"ByteReader":      reflect.TypeOf((*q.ByteReader)(nil)).Elem(),
				"ByteScanner":     reflect.TypeOf((*q.ByteScanner)(nil)).Elem(),
				"ByteWriter":      reflect.TypeOf((*q.ByteWriter)(nil)).Elem(),
				"Closer":          reflect.TypeOf((*q.Closer)(nil)).Elem(),
				"ReadCloser":      reflect.TypeOf((*q.ReadCloser)(nil)).Elem(),
				"ReadSeekCloser":  reflect.TypeOf((*q.ReadSeekCloser)(nil)).Elem(),
				"ReadSeeker":      reflect.TypeOf((*q.ReadSeeker)(nil)).Elem(),
				"ReadWriteCloser": reflect.TypeOf((*q.ReadWriteCloser)(nil)).Elem(),
				"ReadWriteSeeker": reflect.TypeOf((*q.ReadWriteSeeker)(nil)).Elem(),
				"ReadWriter":      reflect.TypeOf((*q.ReadWriter)(nil)).Elem(),
				"Reader":          reflect.TypeOf((*q.Reader)(nil)).Elem(),
				"ReaderAt":        reflect.TypeOf((*q.ReaderAt)(nil)).Elem(),
				"ReaderFrom":      reflect.TypeOf((*q.ReaderFrom)(nil)).Elem(),
				"RuneReader":      reflect.TypeOf((*q.RuneReader)(nil)).Elem(),
				"RuneScanner":     reflect.TypeOf((*q.RuneScanner)(nil)).Elem(),
				"Seeker":          reflect.TypeOf((*q.Seeker)(nil)).Elem(),
				"StringWriter":    reflect.TypeOf((*q.StringWriter)(nil)).Elem(),
				"WriteCloser":     reflect.TypeOf((*q.WriteCloser)(nil)).Elem(),
				"WriteSeeker":     reflect.TypeOf((*q.WriteSeeker)(nil)).Elem(),
				"Writer":          reflect.TypeOf((*q.Writer)(nil)).Elem(),
				"WriterAt":        reflect.TypeOf((*q.WriterAt)(nil)).Elem(),
				"WriterTo":        reflect.TypeOf((*q.WriterTo)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"LimitedReader": reflect.TypeOf((*q.LimitedReader)(nil)).Elem(),
				"OffsetWriter":  reflect.TypeOf((*q.OffsetWriter)(nil)).Elem(),
				"PipeReader":    reflect.TypeOf((*q.PipeReader)(nil)).Elem(),
				"PipeWriter":    reflect.TypeOf((*q.PipeWriter)(nil)).Elem(),
				"SectionReader": reflect.TypeOf((*q.SectionReader)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"Discard":          &q.Discard,
				"EOF":              &q.EOF,
				"ErrClosedPipe":    &q.ErrClosedPipe,
				"ErrNoProgress":    &q.ErrNoProgress,
				"ErrShortBuffer":   &q.ErrShortBuffer,
				"ErrShortWrite":    &q.ErrShortWrite,
				"ErrUnexpectedEOF": &q.ErrUnexpectedEOF,
			},
			Funcs: map[string]interface{}{
				"Copy":             q.Copy,
				"CopyBuffer":       q.CopyBuffer,
				"CopyN":            q.CopyN,
				"LimitReader":      q.LimitReader,
				"MultiReader":      q.MultiReader,
				"MultiWriter":      q.MultiWriter,
				"NewOffsetWriter":  q.NewOffsetWriter,
				"NewSectionReader": q.NewSectionReader,
				"NopCloser":        q.NopCloser,
				"Pipe":             q.Pipe,
				"ReadAll":          q.ReadAll,
				"ReadAtLeast":      q.ReadAtLeast,
				"ReadFull":         q.ReadFull,
				"TeeReader":        q.TeeReader,
				"WriteString":      q.WriteString,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"SeekCurrent": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SeekCurrent))},
				"SeekEnd":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SeekEnd))},
				"SeekStart":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SeekStart))},
			},
		}
	})
}
