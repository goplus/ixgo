// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package build

import (
	q "go/build"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("go/build", map[string]ixgo.DirectCallAdapter{
		"(*Context).SrcDirs":            method_ptr_Context_SrcDirs,
		"(*MultiplePackageError).Error": method_ptr_MultiplePackageError_Error,
		"(*NoGoError).Error":            method_ptr_NoGoError_Error,
		"(*Package).IsCommand":          method_ptr_Package_IsCommand,
		"IsLocalImport":                 func_IsLocalImport,
	})
}

func method_ptr_Context_SrcDirs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Context).SrcDirs(ixgo.DirectCallArg[*q.Context](ctx, 0)))
}

func func_IsLocalImport(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsLocalImport(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_MultiplePackageError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MultiplePackageError).Error(ixgo.DirectCallArg[*q.MultiplePackageError](ctx, 0)))
}

func method_ptr_NoGoError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NoGoError).Error(ixgo.DirectCallArg[*q.NoGoError](ctx, 0)))
}

func method_ptr_Package_IsCommand(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).IsCommand(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}
