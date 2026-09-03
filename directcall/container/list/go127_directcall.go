// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package list

import (
	q "container/list"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("container/list", map[string]ixgo.DirectCallAdapter{
		"(*Element).Next":       method_ptr_Element_Next,
		"(*Element).Prev":       method_ptr_Element_Prev,
		"(*List).Back":          method_ptr_List_Back,
		"(*List).Front":         method_ptr_List_Front,
		"(*List).Init":          method_ptr_List_Init,
		"(*List).InsertAfter":   method_ptr_List_InsertAfter,
		"(*List).InsertBefore":  method_ptr_List_InsertBefore,
		"(*List).Len":           method_ptr_List_Len,
		"(*List).MoveAfter":     method_ptr_List_MoveAfter,
		"(*List).MoveBefore":    method_ptr_List_MoveBefore,
		"(*List).MoveToBack":    method_ptr_List_MoveToBack,
		"(*List).MoveToFront":   method_ptr_List_MoveToFront,
		"(*List).PushBack":      method_ptr_List_PushBack,
		"(*List).PushBackList":  method_ptr_List_PushBackList,
		"(*List).PushFront":     method_ptr_List_PushFront,
		"(*List).PushFrontList": method_ptr_List_PushFrontList,
		"(*List).Remove":        method_ptr_List_Remove,
		"New":                   func_New,
	})
}

func method_ptr_Element_Next(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Element).Next(ixgo.DirectCallArg[*q.Element](ctx, 0)))
}

func method_ptr_Element_Prev(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Element).Prev(ixgo.DirectCallArg[*q.Element](ctx, 0)))
}

func method_ptr_List_Back(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).Back(ixgo.DirectCallArg[*q.List](ctx, 0)))
}

func method_ptr_List_Front(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).Front(ixgo.DirectCallArg[*q.List](ctx, 0)))
}

func method_ptr_List_Init(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).Init(ixgo.DirectCallArg[*q.List](ctx, 0)))
}

func method_ptr_List_InsertAfter(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).InsertAfter(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1), ixgo.DirectCallArg[*q.Element](ctx, 2)))
}

func method_ptr_List_InsertBefore(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).InsertBefore(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1), ixgo.DirectCallArg[*q.Element](ctx, 2)))
}

func method_ptr_List_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).Len(ixgo.DirectCallArg[*q.List](ctx, 0)))
}

func method_ptr_List_MoveAfter(ctx ixgo.DirectCallContext) {
	(*q.List).MoveAfter(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[*q.Element](ctx, 1), ixgo.DirectCallArg[*q.Element](ctx, 2))
}

func method_ptr_List_MoveBefore(ctx ixgo.DirectCallContext) {
	(*q.List).MoveBefore(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[*q.Element](ctx, 1), ixgo.DirectCallArg[*q.Element](ctx, 2))
}

func method_ptr_List_MoveToBack(ctx ixgo.DirectCallContext) {
	(*q.List).MoveToBack(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[*q.Element](ctx, 1))
}

func method_ptr_List_MoveToFront(ctx ixgo.DirectCallContext) {
	(*q.List).MoveToFront(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[*q.Element](ctx, 1))
}

func method_ptr_List_PushBack(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).PushBack(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1)))
}

func method_ptr_List_PushBackList(ctx ixgo.DirectCallContext) {
	(*q.List).PushBackList(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[*q.List](ctx, 1))
}

func method_ptr_List_PushFront(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).PushFront(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1)))
}

func method_ptr_List_PushFrontList(ctx ixgo.DirectCallContext) {
	(*q.List).PushFrontList(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[*q.List](ctx, 1))
}

func method_ptr_List_Remove(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.List).Remove(ixgo.DirectCallArg[*q.List](ctx, 0), ixgo.DirectCallArg[*q.Element](ctx, 1)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New())
}
