// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package filepath

import (
	q "path/filepath"

	"github.com/goplus/ixgo"
	fs "io/fs"
)

func init() {
	ixgo.RegisterDirectCalls("path/filepath", map[string]ixgo.DirectCallAdapter{
		"Base":       func_Base,
		"Clean":      func_Clean,
		"Dir":        func_Dir,
		"Ext":        func_Ext,
		"FromSlash":  func_FromSlash,
		"HasPrefix":  func_HasPrefix,
		"IsAbs":      func_IsAbs,
		"IsLocal":    func_IsLocal,
		"Join":       func_Join,
		"SplitList":  func_SplitList,
		"ToSlash":    func_ToSlash,
		"VolumeName": func_VolumeName,
		"Walk":       func_Walk,
		"WalkDir":    func_WalkDir,
	})
}

func func_Base(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Base(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Clean(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Clean(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Dir(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Dir(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Ext(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Ext(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_FromSlash(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FromSlash(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_HasPrefix(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HasPrefix(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_IsAbs(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsAbs(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_IsLocal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsLocal(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Join(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Join(ixgo.DirectCallArg[[]string](ctx, 0)...))
}

func func_SplitList(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SplitList(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_ToSlash(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToSlash(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_VolumeName(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.VolumeName(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Walk(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Walk(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[q.WalkFunc](ctx, 1)))
}

func func_WalkDir(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WalkDir(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[fs.WalkDirFunc](ctx, 1)))
}
