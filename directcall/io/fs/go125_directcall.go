// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package fs

import (
	q "io/fs"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("io/fs", map[string]ixgo.DirectCallAdapter{
		"(*FileMode).IsDir":     method_ptr_FileMode_IsDir,
		"(*FileMode).IsRegular": method_ptr_FileMode_IsRegular,
		"(*FileMode).Perm":      method_ptr_FileMode_Perm,
		"(*FileMode).String":    method_ptr_FileMode_String,
		"(*FileMode).Type":      method_ptr_FileMode_Type,
		"(*PathError).Error":    method_ptr_PathError_Error,
		"(*PathError).Timeout":  method_ptr_PathError_Timeout,
		"(*PathError).Unwrap":   method_ptr_PathError_Unwrap,
		"(FileMode).IsDir":      method_FileMode_IsDir,
		"(FileMode).IsRegular":  method_FileMode_IsRegular,
		"(FileMode).Perm":       method_FileMode_Perm,
		"(FileMode).String":     method_FileMode_String,
		"(FileMode).Type":       method_FileMode_Type,
		"FileInfoToDirEntry":    func_FileInfoToDirEntry,
		"FormatDirEntry":        func_FormatDirEntry,
		"FormatFileInfo":        func_FormatFileInfo,
		"ValidPath":             func_ValidPath,
		"WalkDir":               func_WalkDir,
	})
}

func func_FileInfoToDirEntry(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FileInfoToDirEntry(ixgo.DirectCallArg[q.FileInfo](ctx, 0)))
}

func method_FileMode_IsDir(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FileMode.IsDir(ixgo.DirectCallArg[q.FileMode](ctx, 0)))
}

func method_ptr_FileMode_IsDir(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileMode).IsDir(ixgo.DirectCallArg[*q.FileMode](ctx, 0)))
}

func method_FileMode_IsRegular(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FileMode.IsRegular(ixgo.DirectCallArg[q.FileMode](ctx, 0)))
}

func method_ptr_FileMode_IsRegular(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileMode).IsRegular(ixgo.DirectCallArg[*q.FileMode](ctx, 0)))
}

func method_FileMode_Perm(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FileMode.Perm(ixgo.DirectCallArg[q.FileMode](ctx, 0)))
}

func method_ptr_FileMode_Perm(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileMode).Perm(ixgo.DirectCallArg[*q.FileMode](ctx, 0)))
}

func method_FileMode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FileMode.String(ixgo.DirectCallArg[q.FileMode](ctx, 0)))
}

func method_ptr_FileMode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileMode).String(ixgo.DirectCallArg[*q.FileMode](ctx, 0)))
}

func method_FileMode_Type(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FileMode.Type(ixgo.DirectCallArg[q.FileMode](ctx, 0)))
}

func method_ptr_FileMode_Type(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileMode).Type(ixgo.DirectCallArg[*q.FileMode](ctx, 0)))
}

func func_FormatDirEntry(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatDirEntry(ixgo.DirectCallArg[q.DirEntry](ctx, 0)))
}

func func_FormatFileInfo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatFileInfo(ixgo.DirectCallArg[q.FileInfo](ctx, 0)))
}

func method_ptr_PathError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PathError).Error(ixgo.DirectCallArg[*q.PathError](ctx, 0)))
}

func method_ptr_PathError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PathError).Timeout(ixgo.DirectCallArg[*q.PathError](ctx, 0)))
}

func method_ptr_PathError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PathError).Unwrap(ixgo.DirectCallArg[*q.PathError](ctx, 0)))
}

func func_ValidPath(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ValidPath(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_WalkDir(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WalkDir(ixgo.DirectCallArg[q.FS](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[q.WalkDirFunc](ctx, 2)))
}
