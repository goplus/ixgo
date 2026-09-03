// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package zip

import (
	q "archive/zip"

	"github.com/goplus/ixgo"
	io "io"
	fs "io/fs"
	time "time"
)

func init() {
	ixgo.RegisterDirectCalls("archive/zip", map[string]ixgo.DirectCallAdapter{
		"(*FileHeader).FileInfo":         method_ptr_FileHeader_FileInfo,
		"(*FileHeader).ModTime":          method_ptr_FileHeader_ModTime,
		"(*FileHeader).Mode":             method_ptr_FileHeader_Mode,
		"(*FileHeader).SetModTime":       method_ptr_FileHeader_SetModTime,
		"(*FileHeader).SetMode":          method_ptr_FileHeader_SetMode,
		"(*ReadCloser).Close":            method_ptr_ReadCloser_Close,
		"(*Reader).RegisterDecompressor": method_ptr_Reader_RegisterDecompressor,
		"(*Writer).AddFS":                method_ptr_Writer_AddFS,
		"(*Writer).Close":                method_ptr_Writer_Close,
		"(*Writer).Copy":                 method_ptr_Writer_Copy,
		"(*Writer).Flush":                method_ptr_Writer_Flush,
		"(*Writer).RegisterCompressor":   method_ptr_Writer_RegisterCompressor,
		"(*Writer).SetComment":           method_ptr_Writer_SetComment,
		"(*Writer).SetOffset":            method_ptr_Writer_SetOffset,
		"NewWriter":                      func_NewWriter,
		"RegisterCompressor":             func_RegisterCompressor,
		"RegisterDecompressor":           func_RegisterDecompressor,
	})
}

func method_ptr_FileHeader_FileInfo(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileHeader).FileInfo(ixgo.DirectCallArg[*q.FileHeader](ctx, 0)))
}

func method_ptr_FileHeader_ModTime(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileHeader).ModTime(ixgo.DirectCallArg[*q.FileHeader](ctx, 0)))
}

func method_ptr_FileHeader_Mode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileHeader).Mode(ixgo.DirectCallArg[*q.FileHeader](ctx, 0)))
}

func method_ptr_FileHeader_SetModTime(ctx ixgo.DirectCallContext) {
	(*q.FileHeader).SetModTime(ixgo.DirectCallArg[*q.FileHeader](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1))
}

func method_ptr_FileHeader_SetMode(ctx ixgo.DirectCallContext) {
	(*q.FileHeader).SetMode(ixgo.DirectCallArg[*q.FileHeader](ctx, 0), ixgo.DirectCallArg[fs.FileMode](ctx, 1))
}

func func_NewWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriter(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func method_ptr_ReadCloser_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ReadCloser).Close(ixgo.DirectCallArg[*q.ReadCloser](ctx, 0)))
}

func method_ptr_Reader_RegisterDecompressor(ctx ixgo.DirectCallContext) {
	(*q.Reader).RegisterDecompressor(ixgo.DirectCallArg[*q.Reader](ctx, 0), ixgo.DirectCallArg[uint16](ctx, 1), ixgo.DirectCallArg[q.Decompressor](ctx, 2))
}

func func_RegisterCompressor(ctx ixgo.DirectCallContext) {
	q.RegisterCompressor(ixgo.DirectCallArg[uint16](ctx, 0), ixgo.DirectCallArg[q.Compressor](ctx, 1))
}

func func_RegisterDecompressor(ctx ixgo.DirectCallContext) {
	q.RegisterDecompressor(ixgo.DirectCallArg[uint16](ctx, 0), ixgo.DirectCallArg[q.Decompressor](ctx, 1))
}

func method_ptr_Writer_AddFS(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).AddFS(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[fs.FS](ctx, 1)))
}

func method_ptr_Writer_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Close(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Copy(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[*q.File](ctx, 1)))
}

func method_ptr_Writer_Flush(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Flush(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_RegisterCompressor(ctx ixgo.DirectCallContext) {
	(*q.Writer).RegisterCompressor(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[uint16](ctx, 1), ixgo.DirectCallArg[q.Compressor](ctx, 2))
}

func method_ptr_Writer_SetComment(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).SetComment(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Writer_SetOffset(ctx ixgo.DirectCallContext) {
	(*q.Writer).SetOffset(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1))
}
