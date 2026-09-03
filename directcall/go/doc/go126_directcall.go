// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package doc

import (
	q "go/doc"

	"github.com/goplus/ixgo"
	ast "go/ast"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("go/doc", map[string]ixgo.DirectCallAdapter{
		"(*Package).Filter":   method_ptr_Package_Filter,
		"(*Package).HTML":     method_ptr_Package_HTML,
		"(*Package).Markdown": method_ptr_Package_Markdown,
		"(*Package).Parser":   method_ptr_Package_Parser,
		"(*Package).Printer":  method_ptr_Package_Printer,
		"(*Package).Synopsis": method_ptr_Package_Synopsis,
		"(*Package).Text":     method_ptr_Package_Text,
		"Examples":            func_Examples,
		"IsPredeclared":       func_IsPredeclared,
		"New":                 func_New,
		"Synopsis":            func_Synopsis,
		"ToHTML":              func_ToHTML,
		"ToText":              func_ToText,
	})
}

func func_Examples(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Examples(ixgo.DirectCallArg[[]*ast.File](ctx, 0)...))
}

func func_IsPredeclared(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsPredeclared(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[*ast.Package](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[q.Mode](ctx, 2)))
}

func method_ptr_Package_Filter(ctx ixgo.DirectCallContext) {
	(*q.Package).Filter(ixgo.DirectCallArg[*q.Package](ctx, 0), ixgo.DirectCallArg[q.Filter](ctx, 1))
}

func method_ptr_Package_HTML(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).HTML(ixgo.DirectCallArg[*q.Package](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Package_Markdown(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Markdown(ixgo.DirectCallArg[*q.Package](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Package_Parser(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Parser(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_Package_Printer(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Printer(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_Package_Synopsis(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Synopsis(ixgo.DirectCallArg[*q.Package](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Package_Text(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Text(ixgo.DirectCallArg[*q.Package](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Synopsis(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Synopsis(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_ToHTML(ctx ixgo.DirectCallContext) {
	q.ToHTML(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[map[string]string](ctx, 2))
}

func func_ToText(ctx ixgo.DirectCallContext) {
	q.ToText(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[string](ctx, 3), ixgo.DirectCallArg[int](ctx, 4))
}
