// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package parse

import (
	q "text/template/parse"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("text/template/parse", map[string]ixgo.DirectCallAdapter{
		"(*ActionNode).Copy":        method_ptr_ActionNode_Copy,
		"(*ActionNode).String":      method_ptr_ActionNode_String,
		"(*BoolNode).Copy":          method_ptr_BoolNode_Copy,
		"(*BoolNode).String":        method_ptr_BoolNode_String,
		"(*BranchNode).Copy":        method_ptr_BranchNode_Copy,
		"(*BranchNode).String":      method_ptr_BranchNode_String,
		"(*BreakNode).Copy":         method_ptr_BreakNode_Copy,
		"(*BreakNode).String":       method_ptr_BreakNode_String,
		"(*ChainNode).Add":          method_ptr_ChainNode_Add,
		"(*ChainNode).Copy":         method_ptr_ChainNode_Copy,
		"(*ChainNode).String":       method_ptr_ChainNode_String,
		"(*CommandNode).Copy":       method_ptr_CommandNode_Copy,
		"(*CommandNode).String":     method_ptr_CommandNode_String,
		"(*CommentNode).Copy":       method_ptr_CommentNode_Copy,
		"(*CommentNode).String":     method_ptr_CommentNode_String,
		"(*ContinueNode).Copy":      method_ptr_ContinueNode_Copy,
		"(*ContinueNode).String":    method_ptr_ContinueNode_String,
		"(*DotNode).Copy":           method_ptr_DotNode_Copy,
		"(*DotNode).String":         method_ptr_DotNode_String,
		"(*DotNode).Type":           method_ptr_DotNode_Type,
		"(*FieldNode).Copy":         method_ptr_FieldNode_Copy,
		"(*FieldNode).String":       method_ptr_FieldNode_String,
		"(*IdentifierNode).Copy":    method_ptr_IdentifierNode_Copy,
		"(*IdentifierNode).SetPos":  method_ptr_IdentifierNode_SetPos,
		"(*IdentifierNode).SetTree": method_ptr_IdentifierNode_SetTree,
		"(*IdentifierNode).String":  method_ptr_IdentifierNode_String,
		"(*IfNode).Copy":            method_ptr_IfNode_Copy,
		"(*ListNode).Copy":          method_ptr_ListNode_Copy,
		"(*ListNode).CopyList":      method_ptr_ListNode_CopyList,
		"(*ListNode).String":        method_ptr_ListNode_String,
		"(*NilNode).Copy":           method_ptr_NilNode_Copy,
		"(*NilNode).String":         method_ptr_NilNode_String,
		"(*NilNode).Type":           method_ptr_NilNode_Type,
		"(*NodeType).Type":          method_ptr_NodeType_Type,
		"(*NumberNode).Copy":        method_ptr_NumberNode_Copy,
		"(*NumberNode).String":      method_ptr_NumberNode_String,
		"(*PipeNode).Copy":          method_ptr_PipeNode_Copy,
		"(*PipeNode).CopyPipe":      method_ptr_PipeNode_CopyPipe,
		"(*PipeNode).String":        method_ptr_PipeNode_String,
		"(*Pos).Position":           method_ptr_Pos_Position,
		"(*RangeNode).Copy":         method_ptr_RangeNode_Copy,
		"(*StringNode).Copy":        method_ptr_StringNode_Copy,
		"(*StringNode).String":      method_ptr_StringNode_String,
		"(*TemplateNode).Copy":      method_ptr_TemplateNode_Copy,
		"(*TemplateNode).String":    method_ptr_TemplateNode_String,
		"(*TextNode).Copy":          method_ptr_TextNode_Copy,
		"(*TextNode).String":        method_ptr_TextNode_String,
		"(*Tree).Copy":              method_ptr_Tree_Copy,
		"(*VariableNode).Copy":      method_ptr_VariableNode_Copy,
		"(*VariableNode).String":    method_ptr_VariableNode_String,
		"(*WithNode).Copy":          method_ptr_WithNode_Copy,
		"(NodeType).Type":           method_NodeType_Type,
		"(Pos).Position":            method_Pos_Position,
		"IsEmptyTree":               func_IsEmptyTree,
		"New":                       func_New,
		"NewIdentifier":             func_NewIdentifier,
	})
}

func method_ptr_ActionNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ActionNode).Copy(ixgo.DirectCallArg[*q.ActionNode](ctx, 0)))
}

func method_ptr_ActionNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ActionNode).String(ixgo.DirectCallArg[*q.ActionNode](ctx, 0)))
}

func method_ptr_BoolNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BoolNode).Copy(ixgo.DirectCallArg[*q.BoolNode](ctx, 0)))
}

func method_ptr_BoolNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BoolNode).String(ixgo.DirectCallArg[*q.BoolNode](ctx, 0)))
}

func method_ptr_BranchNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BranchNode).Copy(ixgo.DirectCallArg[*q.BranchNode](ctx, 0)))
}

func method_ptr_BranchNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BranchNode).String(ixgo.DirectCallArg[*q.BranchNode](ctx, 0)))
}

func method_ptr_BreakNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BreakNode).Copy(ixgo.DirectCallArg[*q.BreakNode](ctx, 0)))
}

func method_ptr_BreakNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BreakNode).String(ixgo.DirectCallArg[*q.BreakNode](ctx, 0)))
}

func method_ptr_ChainNode_Add(ctx ixgo.DirectCallContext) {
	(*q.ChainNode).Add(ixgo.DirectCallArg[*q.ChainNode](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_ChainNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ChainNode).Copy(ixgo.DirectCallArg[*q.ChainNode](ctx, 0)))
}

func method_ptr_ChainNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ChainNode).String(ixgo.DirectCallArg[*q.ChainNode](ctx, 0)))
}

func method_ptr_CommandNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommandNode).Copy(ixgo.DirectCallArg[*q.CommandNode](ctx, 0)))
}

func method_ptr_CommandNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommandNode).String(ixgo.DirectCallArg[*q.CommandNode](ctx, 0)))
}

func method_ptr_CommentNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommentNode).Copy(ixgo.DirectCallArg[*q.CommentNode](ctx, 0)))
}

func method_ptr_CommentNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommentNode).String(ixgo.DirectCallArg[*q.CommentNode](ctx, 0)))
}

func method_ptr_ContinueNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ContinueNode).Copy(ixgo.DirectCallArg[*q.ContinueNode](ctx, 0)))
}

func method_ptr_ContinueNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ContinueNode).String(ixgo.DirectCallArg[*q.ContinueNode](ctx, 0)))
}

func method_ptr_DotNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DotNode).Copy(ixgo.DirectCallArg[*q.DotNode](ctx, 0)))
}

func method_ptr_DotNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DotNode).String(ixgo.DirectCallArg[*q.DotNode](ctx, 0)))
}

func method_ptr_DotNode_Type(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DotNode).Type(ixgo.DirectCallArg[*q.DotNode](ctx, 0)))
}

func method_ptr_FieldNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FieldNode).Copy(ixgo.DirectCallArg[*q.FieldNode](ctx, 0)))
}

func method_ptr_FieldNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FieldNode).String(ixgo.DirectCallArg[*q.FieldNode](ctx, 0)))
}

func method_ptr_IdentifierNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IdentifierNode).Copy(ixgo.DirectCallArg[*q.IdentifierNode](ctx, 0)))
}

func method_ptr_IdentifierNode_SetPos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IdentifierNode).SetPos(ixgo.DirectCallArg[*q.IdentifierNode](ctx, 0), ixgo.DirectCallArg[q.Pos](ctx, 1)))
}

func method_ptr_IdentifierNode_SetTree(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IdentifierNode).SetTree(ixgo.DirectCallArg[*q.IdentifierNode](ctx, 0), ixgo.DirectCallArg[*q.Tree](ctx, 1)))
}

func method_ptr_IdentifierNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IdentifierNode).String(ixgo.DirectCallArg[*q.IdentifierNode](ctx, 0)))
}

func method_ptr_IfNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IfNode).Copy(ixgo.DirectCallArg[*q.IfNode](ctx, 0)))
}

func func_IsEmptyTree(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsEmptyTree(ixgo.DirectCallArg[q.Node](ctx, 0)))
}

func method_ptr_ListNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ListNode).Copy(ixgo.DirectCallArg[*q.ListNode](ctx, 0)))
}

func method_ptr_ListNode_CopyList(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ListNode).CopyList(ixgo.DirectCallArg[*q.ListNode](ctx, 0)))
}

func method_ptr_ListNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ListNode).String(ixgo.DirectCallArg[*q.ListNode](ctx, 0)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[[]map[string]any](ctx, 1)...))
}

func func_NewIdentifier(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewIdentifier(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_NilNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NilNode).Copy(ixgo.DirectCallArg[*q.NilNode](ctx, 0)))
}

func method_ptr_NilNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NilNode).String(ixgo.DirectCallArg[*q.NilNode](ctx, 0)))
}

func method_ptr_NilNode_Type(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NilNode).Type(ixgo.DirectCallArg[*q.NilNode](ctx, 0)))
}

func method_NodeType_Type(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NodeType.Type(ixgo.DirectCallArg[q.NodeType](ctx, 0)))
}

func method_ptr_NodeType_Type(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NodeType).Type(ixgo.DirectCallArg[*q.NodeType](ctx, 0)))
}

func method_ptr_NumberNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NumberNode).Copy(ixgo.DirectCallArg[*q.NumberNode](ctx, 0)))
}

func method_ptr_NumberNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NumberNode).String(ixgo.DirectCallArg[*q.NumberNode](ctx, 0)))
}

func method_ptr_PipeNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PipeNode).Copy(ixgo.DirectCallArg[*q.PipeNode](ctx, 0)))
}

func method_ptr_PipeNode_CopyPipe(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PipeNode).CopyPipe(ixgo.DirectCallArg[*q.PipeNode](ctx, 0)))
}

func method_ptr_PipeNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PipeNode).String(ixgo.DirectCallArg[*q.PipeNode](ctx, 0)))
}

func method_Pos_Position(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pos.Position(ixgo.DirectCallArg[q.Pos](ctx, 0)))
}

func method_ptr_Pos_Position(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pos).Position(ixgo.DirectCallArg[*q.Pos](ctx, 0)))
}

func method_ptr_RangeNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RangeNode).Copy(ixgo.DirectCallArg[*q.RangeNode](ctx, 0)))
}

func method_ptr_StringNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StringNode).Copy(ixgo.DirectCallArg[*q.StringNode](ctx, 0)))
}

func method_ptr_StringNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StringNode).String(ixgo.DirectCallArg[*q.StringNode](ctx, 0)))
}

func method_ptr_TemplateNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TemplateNode).Copy(ixgo.DirectCallArg[*q.TemplateNode](ctx, 0)))
}

func method_ptr_TemplateNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TemplateNode).String(ixgo.DirectCallArg[*q.TemplateNode](ctx, 0)))
}

func method_ptr_TextNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TextNode).Copy(ixgo.DirectCallArg[*q.TextNode](ctx, 0)))
}

func method_ptr_TextNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TextNode).String(ixgo.DirectCallArg[*q.TextNode](ctx, 0)))
}

func method_ptr_Tree_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tree).Copy(ixgo.DirectCallArg[*q.Tree](ctx, 0)))
}

func method_ptr_VariableNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.VariableNode).Copy(ixgo.DirectCallArg[*q.VariableNode](ctx, 0)))
}

func method_ptr_VariableNode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.VariableNode).String(ixgo.DirectCallArg[*q.VariableNode](ctx, 0)))
}

func method_ptr_WithNode_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.WithNode).Copy(ixgo.DirectCallArg[*q.WithNode](ctx, 0)))
}
