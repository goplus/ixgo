// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package ast

import (
	q "go/ast"

	"github.com/goplus/ixgo"
	token "go/token"
	io "io"
	reflect "reflect"
)

func init() {
	ixgo.RegisterDirectCalls("go/ast", map[string]ixgo.DirectCallAdapter{
		"(*ArrayType).End":       method_ptr_ArrayType_End,
		"(*ArrayType).Pos":       method_ptr_ArrayType_Pos,
		"(*AssignStmt).End":      method_ptr_AssignStmt_End,
		"(*AssignStmt).Pos":      method_ptr_AssignStmt_Pos,
		"(*BadDecl).End":         method_ptr_BadDecl_End,
		"(*BadDecl).Pos":         method_ptr_BadDecl_Pos,
		"(*BadExpr).End":         method_ptr_BadExpr_End,
		"(*BadExpr).Pos":         method_ptr_BadExpr_Pos,
		"(*BadStmt).End":         method_ptr_BadStmt_End,
		"(*BadStmt).Pos":         method_ptr_BadStmt_Pos,
		"(*BasicLit).End":        method_ptr_BasicLit_End,
		"(*BasicLit).Pos":        method_ptr_BasicLit_Pos,
		"(*BinaryExpr).End":      method_ptr_BinaryExpr_End,
		"(*BinaryExpr).Pos":      method_ptr_BinaryExpr_Pos,
		"(*BlockStmt).End":       method_ptr_BlockStmt_End,
		"(*BlockStmt).Pos":       method_ptr_BlockStmt_Pos,
		"(*BranchStmt).End":      method_ptr_BranchStmt_End,
		"(*BranchStmt).Pos":      method_ptr_BranchStmt_Pos,
		"(*CallExpr).End":        method_ptr_CallExpr_End,
		"(*CallExpr).Pos":        method_ptr_CallExpr_Pos,
		"(*CaseClause).End":      method_ptr_CaseClause_End,
		"(*CaseClause).Pos":      method_ptr_CaseClause_Pos,
		"(*ChanType).End":        method_ptr_ChanType_End,
		"(*ChanType).Pos":        method_ptr_ChanType_Pos,
		"(*CommClause).End":      method_ptr_CommClause_End,
		"(*CommClause).Pos":      method_ptr_CommClause_Pos,
		"(*Comment).End":         method_ptr_Comment_End,
		"(*Comment).Pos":         method_ptr_Comment_Pos,
		"(*CommentGroup).End":    method_ptr_CommentGroup_End,
		"(*CommentGroup).Pos":    method_ptr_CommentGroup_Pos,
		"(*CommentGroup).Text":   method_ptr_CommentGroup_Text,
		"(*CommentMap).Comments": method_ptr_CommentMap_Comments,
		"(*CommentMap).Filter":   method_ptr_CommentMap_Filter,
		"(*CommentMap).String":   method_ptr_CommentMap_String,
		"(*CommentMap).Update":   method_ptr_CommentMap_Update,
		"(*CompositeLit).End":    method_ptr_CompositeLit_End,
		"(*CompositeLit).Pos":    method_ptr_CompositeLit_Pos,
		"(*DeclStmt).End":        method_ptr_DeclStmt_End,
		"(*DeclStmt).Pos":        method_ptr_DeclStmt_Pos,
		"(*DeferStmt).End":       method_ptr_DeferStmt_End,
		"(*DeferStmt).Pos":       method_ptr_DeferStmt_Pos,
		"(*Ellipsis).End":        method_ptr_Ellipsis_End,
		"(*Ellipsis).Pos":        method_ptr_Ellipsis_Pos,
		"(*EmptyStmt).End":       method_ptr_EmptyStmt_End,
		"(*EmptyStmt).Pos":       method_ptr_EmptyStmt_Pos,
		"(*ExprStmt).End":        method_ptr_ExprStmt_End,
		"(*ExprStmt).Pos":        method_ptr_ExprStmt_Pos,
		"(*Field).End":           method_ptr_Field_End,
		"(*Field).Pos":           method_ptr_Field_Pos,
		"(*FieldList).End":       method_ptr_FieldList_End,
		"(*FieldList).NumFields": method_ptr_FieldList_NumFields,
		"(*FieldList).Pos":       method_ptr_FieldList_Pos,
		"(*File).End":            method_ptr_File_End,
		"(*File).Pos":            method_ptr_File_Pos,
		"(*ForStmt).End":         method_ptr_ForStmt_End,
		"(*ForStmt).Pos":         method_ptr_ForStmt_Pos,
		"(*FuncDecl).End":        method_ptr_FuncDecl_End,
		"(*FuncDecl).Pos":        method_ptr_FuncDecl_Pos,
		"(*FuncLit).End":         method_ptr_FuncLit_End,
		"(*FuncLit).Pos":         method_ptr_FuncLit_Pos,
		"(*FuncType).End":        method_ptr_FuncType_End,
		"(*FuncType).Pos":        method_ptr_FuncType_Pos,
		"(*GenDecl).End":         method_ptr_GenDecl_End,
		"(*GenDecl).Pos":         method_ptr_GenDecl_Pos,
		"(*GoStmt).End":          method_ptr_GoStmt_End,
		"(*GoStmt).Pos":          method_ptr_GoStmt_Pos,
		"(*Ident).End":           method_ptr_Ident_End,
		"(*Ident).IsExported":    method_ptr_Ident_IsExported,
		"(*Ident).Pos":           method_ptr_Ident_Pos,
		"(*Ident).String":        method_ptr_Ident_String,
		"(*IfStmt).End":          method_ptr_IfStmt_End,
		"(*IfStmt).Pos":          method_ptr_IfStmt_Pos,
		"(*ImportSpec).End":      method_ptr_ImportSpec_End,
		"(*ImportSpec).Pos":      method_ptr_ImportSpec_Pos,
		"(*IncDecStmt).End":      method_ptr_IncDecStmt_End,
		"(*IncDecStmt).Pos":      method_ptr_IncDecStmt_Pos,
		"(*IndexExpr).End":       method_ptr_IndexExpr_End,
		"(*IndexExpr).Pos":       method_ptr_IndexExpr_Pos,
		"(*IndexListExpr).End":   method_ptr_IndexListExpr_End,
		"(*IndexListExpr).Pos":   method_ptr_IndexListExpr_Pos,
		"(*InterfaceType).End":   method_ptr_InterfaceType_End,
		"(*InterfaceType).Pos":   method_ptr_InterfaceType_Pos,
		"(*KeyValueExpr).End":    method_ptr_KeyValueExpr_End,
		"(*KeyValueExpr).Pos":    method_ptr_KeyValueExpr_Pos,
		"(*LabeledStmt).End":     method_ptr_LabeledStmt_End,
		"(*LabeledStmt).Pos":     method_ptr_LabeledStmt_Pos,
		"(*MapType).End":         method_ptr_MapType_End,
		"(*MapType).Pos":         method_ptr_MapType_Pos,
		"(*ObjKind).String":      method_ptr_ObjKind_String,
		"(*Object).Pos":          method_ptr_Object_Pos,
		"(*Package).End":         method_ptr_Package_End,
		"(*Package).Pos":         method_ptr_Package_Pos,
		"(*ParenExpr).End":       method_ptr_ParenExpr_End,
		"(*ParenExpr).Pos":       method_ptr_ParenExpr_Pos,
		"(*RangeStmt).End":       method_ptr_RangeStmt_End,
		"(*RangeStmt).Pos":       method_ptr_RangeStmt_Pos,
		"(*ReturnStmt).End":      method_ptr_ReturnStmt_End,
		"(*ReturnStmt).Pos":      method_ptr_ReturnStmt_Pos,
		"(*Scope).Insert":        method_ptr_Scope_Insert,
		"(*Scope).Lookup":        method_ptr_Scope_Lookup,
		"(*Scope).String":        method_ptr_Scope_String,
		"(*SelectStmt).End":      method_ptr_SelectStmt_End,
		"(*SelectStmt).Pos":      method_ptr_SelectStmt_Pos,
		"(*SelectorExpr).End":    method_ptr_SelectorExpr_End,
		"(*SelectorExpr).Pos":    method_ptr_SelectorExpr_Pos,
		"(*SendStmt).End":        method_ptr_SendStmt_End,
		"(*SendStmt).Pos":        method_ptr_SendStmt_Pos,
		"(*SliceExpr).End":       method_ptr_SliceExpr_End,
		"(*SliceExpr).Pos":       method_ptr_SliceExpr_Pos,
		"(*StarExpr).End":        method_ptr_StarExpr_End,
		"(*StarExpr).Pos":        method_ptr_StarExpr_Pos,
		"(*StructType).End":      method_ptr_StructType_End,
		"(*StructType).Pos":      method_ptr_StructType_Pos,
		"(*SwitchStmt).End":      method_ptr_SwitchStmt_End,
		"(*SwitchStmt).Pos":      method_ptr_SwitchStmt_Pos,
		"(*TypeAssertExpr).End":  method_ptr_TypeAssertExpr_End,
		"(*TypeAssertExpr).Pos":  method_ptr_TypeAssertExpr_Pos,
		"(*TypeSpec).End":        method_ptr_TypeSpec_End,
		"(*TypeSpec).Pos":        method_ptr_TypeSpec_Pos,
		"(*TypeSwitchStmt).End":  method_ptr_TypeSwitchStmt_End,
		"(*TypeSwitchStmt).Pos":  method_ptr_TypeSwitchStmt_Pos,
		"(*UnaryExpr).End":       method_ptr_UnaryExpr_End,
		"(*UnaryExpr).Pos":       method_ptr_UnaryExpr_Pos,
		"(*ValueSpec).End":       method_ptr_ValueSpec_End,
		"(*ValueSpec).Pos":       method_ptr_ValueSpec_Pos,
		"(CommentMap).Comments":  method_CommentMap_Comments,
		"(CommentMap).Filter":    method_CommentMap_Filter,
		"(CommentMap).String":    method_CommentMap_String,
		"(CommentMap).Update":    method_CommentMap_Update,
		"(ObjKind).String":       method_ObjKind_String,
		"FileExports":            func_FileExports,
		"FilterDecl":             func_FilterDecl,
		"FilterFile":             func_FilterFile,
		"FilterPackage":          func_FilterPackage,
		"Fprint":                 func_Fprint,
		"Inspect":                func_Inspect,
		"IsExported":             func_IsExported,
		"IsGenerated":            func_IsGenerated,
		"MergePackageFiles":      func_MergePackageFiles,
		"NewCommentMap":          func_NewCommentMap,
		"NewIdent":               func_NewIdent,
		"NewObj":                 func_NewObj,
		"NewScope":               func_NewScope,
		"NotNilFilter":           func_NotNilFilter,
		"PackageExports":         func_PackageExports,
		"Preorder":               func_Preorder,
		"PreorderStack":          func_PreorderStack,
		"Print":                  func_Print,
		"SortImports":            func_SortImports,
		"Unparen":                func_Unparen,
		"Walk":                   func_Walk,
	})
}

func method_ptr_ArrayType_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ArrayType).End(ixgo.DirectCallArg[*q.ArrayType](ctx, 0)))
}

func method_ptr_ArrayType_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ArrayType).Pos(ixgo.DirectCallArg[*q.ArrayType](ctx, 0)))
}

func method_ptr_AssignStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AssignStmt).End(ixgo.DirectCallArg[*q.AssignStmt](ctx, 0)))
}

func method_ptr_AssignStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AssignStmt).Pos(ixgo.DirectCallArg[*q.AssignStmt](ctx, 0)))
}

func method_ptr_BadDecl_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BadDecl).End(ixgo.DirectCallArg[*q.BadDecl](ctx, 0)))
}

func method_ptr_BadDecl_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BadDecl).Pos(ixgo.DirectCallArg[*q.BadDecl](ctx, 0)))
}

func method_ptr_BadExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BadExpr).End(ixgo.DirectCallArg[*q.BadExpr](ctx, 0)))
}

func method_ptr_BadExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BadExpr).Pos(ixgo.DirectCallArg[*q.BadExpr](ctx, 0)))
}

func method_ptr_BadStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BadStmt).End(ixgo.DirectCallArg[*q.BadStmt](ctx, 0)))
}

func method_ptr_BadStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BadStmt).Pos(ixgo.DirectCallArg[*q.BadStmt](ctx, 0)))
}

func method_ptr_BasicLit_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BasicLit).End(ixgo.DirectCallArg[*q.BasicLit](ctx, 0)))
}

func method_ptr_BasicLit_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BasicLit).Pos(ixgo.DirectCallArg[*q.BasicLit](ctx, 0)))
}

func method_ptr_BinaryExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BinaryExpr).End(ixgo.DirectCallArg[*q.BinaryExpr](ctx, 0)))
}

func method_ptr_BinaryExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BinaryExpr).Pos(ixgo.DirectCallArg[*q.BinaryExpr](ctx, 0)))
}

func method_ptr_BlockStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BlockStmt).End(ixgo.DirectCallArg[*q.BlockStmt](ctx, 0)))
}

func method_ptr_BlockStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BlockStmt).Pos(ixgo.DirectCallArg[*q.BlockStmt](ctx, 0)))
}

func method_ptr_BranchStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BranchStmt).End(ixgo.DirectCallArg[*q.BranchStmt](ctx, 0)))
}

func method_ptr_BranchStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BranchStmt).Pos(ixgo.DirectCallArg[*q.BranchStmt](ctx, 0)))
}

func method_ptr_CallExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CallExpr).End(ixgo.DirectCallArg[*q.CallExpr](ctx, 0)))
}

func method_ptr_CallExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CallExpr).Pos(ixgo.DirectCallArg[*q.CallExpr](ctx, 0)))
}

func method_ptr_CaseClause_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CaseClause).End(ixgo.DirectCallArg[*q.CaseClause](ctx, 0)))
}

func method_ptr_CaseClause_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CaseClause).Pos(ixgo.DirectCallArg[*q.CaseClause](ctx, 0)))
}

func method_ptr_ChanType_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ChanType).End(ixgo.DirectCallArg[*q.ChanType](ctx, 0)))
}

func method_ptr_ChanType_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ChanType).Pos(ixgo.DirectCallArg[*q.ChanType](ctx, 0)))
}

func method_ptr_CommClause_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommClause).End(ixgo.DirectCallArg[*q.CommClause](ctx, 0)))
}

func method_ptr_CommClause_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommClause).Pos(ixgo.DirectCallArg[*q.CommClause](ctx, 0)))
}

func method_ptr_Comment_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Comment).End(ixgo.DirectCallArg[*q.Comment](ctx, 0)))
}

func method_ptr_Comment_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Comment).Pos(ixgo.DirectCallArg[*q.Comment](ctx, 0)))
}

func method_ptr_CommentGroup_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommentGroup).End(ixgo.DirectCallArg[*q.CommentGroup](ctx, 0)))
}

func method_ptr_CommentGroup_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommentGroup).Pos(ixgo.DirectCallArg[*q.CommentGroup](ctx, 0)))
}

func method_ptr_CommentGroup_Text(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommentGroup).Text(ixgo.DirectCallArg[*q.CommentGroup](ctx, 0)))
}

func method_CommentMap_Comments(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CommentMap.Comments(ixgo.DirectCallArg[q.CommentMap](ctx, 0)))
}

func method_ptr_CommentMap_Comments(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommentMap).Comments(ixgo.DirectCallArg[*q.CommentMap](ctx, 0)))
}

func method_CommentMap_Filter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CommentMap.Filter(ixgo.DirectCallArg[q.CommentMap](ctx, 0), ixgo.DirectCallArg[q.Node](ctx, 1)))
}

func method_ptr_CommentMap_Filter(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommentMap).Filter(ixgo.DirectCallArg[*q.CommentMap](ctx, 0), ixgo.DirectCallArg[q.Node](ctx, 1)))
}

func method_CommentMap_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CommentMap.String(ixgo.DirectCallArg[q.CommentMap](ctx, 0)))
}

func method_ptr_CommentMap_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommentMap).String(ixgo.DirectCallArg[*q.CommentMap](ctx, 0)))
}

func method_CommentMap_Update(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CommentMap.Update(ixgo.DirectCallArg[q.CommentMap](ctx, 0), ixgo.DirectCallArg[q.Node](ctx, 1), ixgo.DirectCallArg[q.Node](ctx, 2)))
}

func method_ptr_CommentMap_Update(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommentMap).Update(ixgo.DirectCallArg[*q.CommentMap](ctx, 0), ixgo.DirectCallArg[q.Node](ctx, 1), ixgo.DirectCallArg[q.Node](ctx, 2)))
}

func method_ptr_CompositeLit_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CompositeLit).End(ixgo.DirectCallArg[*q.CompositeLit](ctx, 0)))
}

func method_ptr_CompositeLit_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CompositeLit).Pos(ixgo.DirectCallArg[*q.CompositeLit](ctx, 0)))
}

func method_ptr_DeclStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DeclStmt).End(ixgo.DirectCallArg[*q.DeclStmt](ctx, 0)))
}

func method_ptr_DeclStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DeclStmt).Pos(ixgo.DirectCallArg[*q.DeclStmt](ctx, 0)))
}

func method_ptr_DeferStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DeferStmt).End(ixgo.DirectCallArg[*q.DeferStmt](ctx, 0)))
}

func method_ptr_DeferStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DeferStmt).Pos(ixgo.DirectCallArg[*q.DeferStmt](ctx, 0)))
}

func method_ptr_Ellipsis_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ellipsis).End(ixgo.DirectCallArg[*q.Ellipsis](ctx, 0)))
}

func method_ptr_Ellipsis_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ellipsis).Pos(ixgo.DirectCallArg[*q.Ellipsis](ctx, 0)))
}

func method_ptr_EmptyStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.EmptyStmt).End(ixgo.DirectCallArg[*q.EmptyStmt](ctx, 0)))
}

func method_ptr_EmptyStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.EmptyStmt).Pos(ixgo.DirectCallArg[*q.EmptyStmt](ctx, 0)))
}

func method_ptr_ExprStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ExprStmt).End(ixgo.DirectCallArg[*q.ExprStmt](ctx, 0)))
}

func method_ptr_ExprStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ExprStmt).Pos(ixgo.DirectCallArg[*q.ExprStmt](ctx, 0)))
}

func method_ptr_Field_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Field).End(ixgo.DirectCallArg[*q.Field](ctx, 0)))
}

func method_ptr_Field_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Field).Pos(ixgo.DirectCallArg[*q.Field](ctx, 0)))
}

func method_ptr_FieldList_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FieldList).End(ixgo.DirectCallArg[*q.FieldList](ctx, 0)))
}

func method_ptr_FieldList_NumFields(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FieldList).NumFields(ixgo.DirectCallArg[*q.FieldList](ctx, 0)))
}

func method_ptr_FieldList_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FieldList).Pos(ixgo.DirectCallArg[*q.FieldList](ctx, 0)))
}

func method_ptr_File_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).End(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Pos(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func func_FileExports(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FileExports(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func func_FilterDecl(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FilterDecl(ixgo.DirectCallArg[q.Decl](ctx, 0), ixgo.DirectCallArg[q.Filter](ctx, 1)))
}

func func_FilterFile(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FilterFile(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[q.Filter](ctx, 1)))
}

func func_FilterPackage(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FilterPackage(ixgo.DirectCallArg[*q.Package](ctx, 0), ixgo.DirectCallArg[q.Filter](ctx, 1)))
}

func method_ptr_ForStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ForStmt).End(ixgo.DirectCallArg[*q.ForStmt](ctx, 0)))
}

func method_ptr_ForStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ForStmt).Pos(ixgo.DirectCallArg[*q.ForStmt](ctx, 0)))
}

func func_Fprint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Fprint(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[*token.FileSet](ctx, 1), ixgo.DirectCallArg[any](ctx, 2), ixgo.DirectCallArg[q.FieldFilter](ctx, 3)))
}

func method_ptr_FuncDecl_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FuncDecl).End(ixgo.DirectCallArg[*q.FuncDecl](ctx, 0)))
}

func method_ptr_FuncDecl_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FuncDecl).Pos(ixgo.DirectCallArg[*q.FuncDecl](ctx, 0)))
}

func method_ptr_FuncLit_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FuncLit).End(ixgo.DirectCallArg[*q.FuncLit](ctx, 0)))
}

func method_ptr_FuncLit_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FuncLit).Pos(ixgo.DirectCallArg[*q.FuncLit](ctx, 0)))
}

func method_ptr_FuncType_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FuncType).End(ixgo.DirectCallArg[*q.FuncType](ctx, 0)))
}

func method_ptr_FuncType_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FuncType).Pos(ixgo.DirectCallArg[*q.FuncType](ctx, 0)))
}

func method_ptr_GenDecl_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.GenDecl).End(ixgo.DirectCallArg[*q.GenDecl](ctx, 0)))
}

func method_ptr_GenDecl_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.GenDecl).Pos(ixgo.DirectCallArg[*q.GenDecl](ctx, 0)))
}

func method_ptr_GoStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.GoStmt).End(ixgo.DirectCallArg[*q.GoStmt](ctx, 0)))
}

func method_ptr_GoStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.GoStmt).Pos(ixgo.DirectCallArg[*q.GoStmt](ctx, 0)))
}

func method_ptr_Ident_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ident).End(ixgo.DirectCallArg[*q.Ident](ctx, 0)))
}

func method_ptr_Ident_IsExported(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ident).IsExported(ixgo.DirectCallArg[*q.Ident](ctx, 0)))
}

func method_ptr_Ident_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ident).Pos(ixgo.DirectCallArg[*q.Ident](ctx, 0)))
}

func method_ptr_Ident_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Ident).String(ixgo.DirectCallArg[*q.Ident](ctx, 0)))
}

func method_ptr_IfStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IfStmt).End(ixgo.DirectCallArg[*q.IfStmt](ctx, 0)))
}

func method_ptr_IfStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IfStmt).Pos(ixgo.DirectCallArg[*q.IfStmt](ctx, 0)))
}

func method_ptr_ImportSpec_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ImportSpec).End(ixgo.DirectCallArg[*q.ImportSpec](ctx, 0)))
}

func method_ptr_ImportSpec_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ImportSpec).Pos(ixgo.DirectCallArg[*q.ImportSpec](ctx, 0)))
}

func method_ptr_IncDecStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IncDecStmt).End(ixgo.DirectCallArg[*q.IncDecStmt](ctx, 0)))
}

func method_ptr_IncDecStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IncDecStmt).Pos(ixgo.DirectCallArg[*q.IncDecStmt](ctx, 0)))
}

func method_ptr_IndexExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IndexExpr).End(ixgo.DirectCallArg[*q.IndexExpr](ctx, 0)))
}

func method_ptr_IndexExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IndexExpr).Pos(ixgo.DirectCallArg[*q.IndexExpr](ctx, 0)))
}

func method_ptr_IndexListExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IndexListExpr).End(ixgo.DirectCallArg[*q.IndexListExpr](ctx, 0)))
}

func method_ptr_IndexListExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IndexListExpr).Pos(ixgo.DirectCallArg[*q.IndexListExpr](ctx, 0)))
}

func func_Inspect(ctx ixgo.DirectCallContext) {
	q.Inspect(ixgo.DirectCallArg[q.Node](ctx, 0), ixgo.DirectCallArg[func(q.Node) bool](ctx, 1))
}

func method_ptr_InterfaceType_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InterfaceType).End(ixgo.DirectCallArg[*q.InterfaceType](ctx, 0)))
}

func method_ptr_InterfaceType_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InterfaceType).Pos(ixgo.DirectCallArg[*q.InterfaceType](ctx, 0)))
}

func func_IsExported(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsExported(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_IsGenerated(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsGenerated(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_KeyValueExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.KeyValueExpr).End(ixgo.DirectCallArg[*q.KeyValueExpr](ctx, 0)))
}

func method_ptr_KeyValueExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.KeyValueExpr).Pos(ixgo.DirectCallArg[*q.KeyValueExpr](ctx, 0)))
}

func method_ptr_LabeledStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LabeledStmt).End(ixgo.DirectCallArg[*q.LabeledStmt](ctx, 0)))
}

func method_ptr_LabeledStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LabeledStmt).Pos(ixgo.DirectCallArg[*q.LabeledStmt](ctx, 0)))
}

func method_ptr_MapType_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MapType).End(ixgo.DirectCallArg[*q.MapType](ctx, 0)))
}

func method_ptr_MapType_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MapType).Pos(ixgo.DirectCallArg[*q.MapType](ctx, 0)))
}

func func_MergePackageFiles(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MergePackageFiles(ixgo.DirectCallArg[*q.Package](ctx, 0), ixgo.DirectCallArg[q.MergeMode](ctx, 1)))
}

func func_NewCommentMap(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCommentMap(ixgo.DirectCallArg[*token.FileSet](ctx, 0), ixgo.DirectCallArg[q.Node](ctx, 1), ixgo.DirectCallArg[[]*q.CommentGroup](ctx, 2)))
}

func func_NewIdent(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewIdent(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_NewObj(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewObj(ixgo.DirectCallArg[q.ObjKind](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_NewScope(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewScope(ixgo.DirectCallArg[*q.Scope](ctx, 0)))
}

func func_NotNilFilter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NotNilFilter(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[reflect.Value](ctx, 1)))
}

func method_ObjKind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ObjKind.String(ixgo.DirectCallArg[q.ObjKind](ctx, 0)))
}

func method_ptr_ObjKind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ObjKind).String(ixgo.DirectCallArg[*q.ObjKind](ctx, 0)))
}

func method_ptr_Object_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Object).Pos(ixgo.DirectCallArg[*q.Object](ctx, 0)))
}

func method_ptr_Package_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).End(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_Package_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Pos(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func func_PackageExports(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PackageExports(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_ParenExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ParenExpr).End(ixgo.DirectCallArg[*q.ParenExpr](ctx, 0)))
}

func method_ptr_ParenExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ParenExpr).Pos(ixgo.DirectCallArg[*q.ParenExpr](ctx, 0)))
}

func func_Preorder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Preorder(ixgo.DirectCallArg[q.Node](ctx, 0)))
}

func func_PreorderStack(ctx ixgo.DirectCallContext) {
	q.PreorderStack(ixgo.DirectCallArg[q.Node](ctx, 0), ixgo.DirectCallArg[[]q.Node](ctx, 1), ixgo.DirectCallArg[func(n q.Node, stack []q.Node) bool](ctx, 2))
}

func func_Print(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Print(ixgo.DirectCallArg[*token.FileSet](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_RangeStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RangeStmt).End(ixgo.DirectCallArg[*q.RangeStmt](ctx, 0)))
}

func method_ptr_RangeStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RangeStmt).Pos(ixgo.DirectCallArg[*q.RangeStmt](ctx, 0)))
}

func method_ptr_ReturnStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ReturnStmt).End(ixgo.DirectCallArg[*q.ReturnStmt](ctx, 0)))
}

func method_ptr_ReturnStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ReturnStmt).Pos(ixgo.DirectCallArg[*q.ReturnStmt](ctx, 0)))
}

func method_ptr_Scope_Insert(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Insert(ixgo.DirectCallArg[*q.Scope](ctx, 0), ixgo.DirectCallArg[*q.Object](ctx, 1)))
}

func method_ptr_Scope_Lookup(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Lookup(ixgo.DirectCallArg[*q.Scope](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Scope_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).String(ixgo.DirectCallArg[*q.Scope](ctx, 0)))
}

func method_ptr_SelectStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SelectStmt).End(ixgo.DirectCallArg[*q.SelectStmt](ctx, 0)))
}

func method_ptr_SelectStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SelectStmt).Pos(ixgo.DirectCallArg[*q.SelectStmt](ctx, 0)))
}

func method_ptr_SelectorExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SelectorExpr).End(ixgo.DirectCallArg[*q.SelectorExpr](ctx, 0)))
}

func method_ptr_SelectorExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SelectorExpr).Pos(ixgo.DirectCallArg[*q.SelectorExpr](ctx, 0)))
}

func method_ptr_SendStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SendStmt).End(ixgo.DirectCallArg[*q.SendStmt](ctx, 0)))
}

func method_ptr_SendStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SendStmt).Pos(ixgo.DirectCallArg[*q.SendStmt](ctx, 0)))
}

func method_ptr_SliceExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SliceExpr).End(ixgo.DirectCallArg[*q.SliceExpr](ctx, 0)))
}

func method_ptr_SliceExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SliceExpr).Pos(ixgo.DirectCallArg[*q.SliceExpr](ctx, 0)))
}

func func_SortImports(ctx ixgo.DirectCallContext) {
	q.SortImports(ixgo.DirectCallArg[*token.FileSet](ctx, 0), ixgo.DirectCallArg[*q.File](ctx, 1))
}

func method_ptr_StarExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StarExpr).End(ixgo.DirectCallArg[*q.StarExpr](ctx, 0)))
}

func method_ptr_StarExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StarExpr).Pos(ixgo.DirectCallArg[*q.StarExpr](ctx, 0)))
}

func method_ptr_StructType_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StructType).End(ixgo.DirectCallArg[*q.StructType](ctx, 0)))
}

func method_ptr_StructType_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StructType).Pos(ixgo.DirectCallArg[*q.StructType](ctx, 0)))
}

func method_ptr_SwitchStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SwitchStmt).End(ixgo.DirectCallArg[*q.SwitchStmt](ctx, 0)))
}

func method_ptr_SwitchStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SwitchStmt).Pos(ixgo.DirectCallArg[*q.SwitchStmt](ctx, 0)))
}

func method_ptr_TypeAssertExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAssertExpr).End(ixgo.DirectCallArg[*q.TypeAssertExpr](ctx, 0)))
}

func method_ptr_TypeAssertExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAssertExpr).Pos(ixgo.DirectCallArg[*q.TypeAssertExpr](ctx, 0)))
}

func method_ptr_TypeSpec_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeSpec).End(ixgo.DirectCallArg[*q.TypeSpec](ctx, 0)))
}

func method_ptr_TypeSpec_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeSpec).Pos(ixgo.DirectCallArg[*q.TypeSpec](ctx, 0)))
}

func method_ptr_TypeSwitchStmt_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeSwitchStmt).End(ixgo.DirectCallArg[*q.TypeSwitchStmt](ctx, 0)))
}

func method_ptr_TypeSwitchStmt_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeSwitchStmt).Pos(ixgo.DirectCallArg[*q.TypeSwitchStmt](ctx, 0)))
}

func method_ptr_UnaryExpr_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnaryExpr).End(ixgo.DirectCallArg[*q.UnaryExpr](ctx, 0)))
}

func method_ptr_UnaryExpr_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnaryExpr).Pos(ixgo.DirectCallArg[*q.UnaryExpr](ctx, 0)))
}

func func_Unparen(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Unparen(ixgo.DirectCallArg[q.Expr](ctx, 0)))
}

func method_ptr_ValueSpec_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ValueSpec).End(ixgo.DirectCallArg[*q.ValueSpec](ctx, 0)))
}

func method_ptr_ValueSpec_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ValueSpec).Pos(ixgo.DirectCallArg[*q.ValueSpec](ctx, 0)))
}

func func_Walk(ctx ixgo.DirectCallContext) {
	q.Walk(ixgo.DirectCallArg[q.Visitor](ctx, 0), ixgo.DirectCallArg[q.Node](ctx, 1))
}
