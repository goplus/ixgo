// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package comment

import (
	q "go/doc/comment"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("go/doc/comment", map[string]ixgo.DirectCallAdapter{
		"(*DocLink).DefaultURL": method_ptr_DocLink_DefaultURL,
		"(*Heading).DefaultID":  method_ptr_Heading_DefaultID,
		"(*List).BlankBefore":   method_ptr_List_BlankBefore,
		"(*List).BlankBetween":  method_ptr_List_BlankBetween,
		"(*Parser).Parse":       method_ptr_Parser_Parse,
		"(*Printer).Comment":    method_ptr_Printer_Comment,
		"(*Printer).HTML":       method_ptr_Printer_HTML,
		"(*Printer).Markdown":   method_ptr_Printer_Markdown,
		"(*Printer).Text":       method_ptr_Printer_Text,
	})
}

func method_ptr_DocLink_DefaultURL(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DocLink).DefaultURL(ixgo.DirectCallArg[*q.DocLink](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Heading_DefaultID(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Heading).DefaultID(ixgo.DirectCallArg[*q.Heading](ctx, 0)))
}

func method_ptr_List_BlankBefore(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).BlankBefore(ixgo.DirectCallArg[*q.List](ctx, 0)))
}

func method_ptr_List_BlankBetween(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).BlankBetween(ixgo.DirectCallArg[*q.List](ctx, 0)))
}

func method_ptr_Parser_Parse(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Parser).Parse(ixgo.DirectCallArg[*q.Parser](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Printer_Comment(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Printer).Comment(ixgo.DirectCallArg[*q.Printer](ctx, 0), ixgo.DirectCallArg[*q.Doc](ctx, 1)))
}

func method_ptr_Printer_HTML(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Printer).HTML(ixgo.DirectCallArg[*q.Printer](ctx, 0), ixgo.DirectCallArg[*q.Doc](ctx, 1)))
}

func method_ptr_Printer_Markdown(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Printer).Markdown(ixgo.DirectCallArg[*q.Printer](ctx, 0), ixgo.DirectCallArg[*q.Doc](ctx, 1)))
}

func method_ptr_Printer_Text(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Printer).Text(ixgo.DirectCallArg[*q.Printer](ctx, 0), ixgo.DirectCallArg[*q.Doc](ctx, 1)))
}
