// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package template

import (
	q "text/template"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("text/template", map[string]ixgo.DirectCallAdapter{
		"(*ExecError).Error":           method_ptr_ExecError_Error,
		"(*ExecError).Unwrap":          method_ptr_ExecError_Unwrap,
		"(*Template).DefinedTemplates": method_ptr_Template_DefinedTemplates,
		"(*Template).Delims":           method_ptr_Template_Delims,
		"(*Template).Execute":          method_ptr_Template_Execute,
		"(*Template).ExecuteTemplate":  method_ptr_Template_ExecuteTemplate,
		"(*Template).Funcs":            method_ptr_Template_Funcs,
		"(*Template).Lookup":           method_ptr_Template_Lookup,
		"(*Template).Name":             method_ptr_Template_Name,
		"(*Template).New":              method_ptr_Template_New,
		"(*Template).Option":           method_ptr_Template_Option,
		"(*Template).Templates":        method_ptr_Template_Templates,
		"(ExecError).Error":            method_ExecError_Error,
		"(ExecError).Unwrap":           method_ExecError_Unwrap,
		"HTMLEscape":                   func_HTMLEscape,
		"HTMLEscapeString":             func_HTMLEscapeString,
		"HTMLEscaper":                  func_HTMLEscaper,
		"JSEscape":                     func_JSEscape,
		"JSEscapeString":               func_JSEscapeString,
		"JSEscaper":                    func_JSEscaper,
		"Must":                         func_Must,
		"New":                          func_New,
		"URLQueryEscaper":              func_URLQueryEscaper,
	})
}

func method_ExecError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ExecError.Error(ixgo.DirectCallArg[q.ExecError](ctx, 0)))
}

func method_ptr_ExecError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ExecError).Error(ixgo.DirectCallArg[*q.ExecError](ctx, 0)))
}

func method_ExecError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ExecError.Unwrap(ixgo.DirectCallArg[q.ExecError](ctx, 0)))
}

func method_ptr_ExecError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ExecError).Unwrap(ixgo.DirectCallArg[*q.ExecError](ctx, 0)))
}

func func_HTMLEscape(ctx ixgo.DirectCallContext) {
	q.HTMLEscape(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1))
}

func func_HTMLEscapeString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HTMLEscapeString(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_HTMLEscaper(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HTMLEscaper(ixgo.DirectCallArg[[]any](ctx, 0)...))
}

func func_JSEscape(ctx ixgo.DirectCallContext) {
	q.JSEscape(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1))
}

func func_JSEscapeString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.JSEscapeString(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_JSEscaper(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.JSEscaper(ixgo.DirectCallArg[[]any](ctx, 0)...))
}

func func_Must(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Must(ixgo.DirectCallArg[*q.Template](ctx, 0), ixgo.DirectCallArg[error](ctx, 1)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_Template_DefinedTemplates(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Template).DefinedTemplates(ixgo.DirectCallArg[*q.Template](ctx, 0)))
}

func method_ptr_Template_Delims(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Template).Delims(ixgo.DirectCallArg[*q.Template](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_Template_Execute(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Template).Execute(ixgo.DirectCallArg[*q.Template](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[any](ctx, 2)))
}

func method_ptr_Template_ExecuteTemplate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Template).ExecuteTemplate(ixgo.DirectCallArg[*q.Template](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[any](ctx, 3)))
}

func method_ptr_Template_Funcs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Template).Funcs(ixgo.DirectCallArg[*q.Template](ctx, 0), ixgo.DirectCallArg[q.FuncMap](ctx, 1)))
}

func method_ptr_Template_Lookup(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Template).Lookup(ixgo.DirectCallArg[*q.Template](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Template_Name(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Template).Name(ixgo.DirectCallArg[*q.Template](ctx, 0)))
}

func method_ptr_Template_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Template).New(ixgo.DirectCallArg[*q.Template](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Template_Option(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Template).Option(ixgo.DirectCallArg[*q.Template](ctx, 0), ixgo.DirectCallArg[[]string](ctx, 1)...))
}

func method_ptr_Template_Templates(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Template).Templates(ixgo.DirectCallArg[*q.Template](ctx, 0)))
}

func func_URLQueryEscaper(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.URLQueryEscaper(ixgo.DirectCallArg[[]any](ctx, 0)...))
}
