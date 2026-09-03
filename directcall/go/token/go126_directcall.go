// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package token

import (
	q "go/token"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("go/token", map[string]ixgo.DirectCallAdapter{
		"(*File).AddLine":             method_ptr_File_AddLine,
		"(*File).AddLineColumnInfo":   method_ptr_File_AddLineColumnInfo,
		"(*File).AddLineInfo":         method_ptr_File_AddLineInfo,
		"(*File).Base":                method_ptr_File_Base,
		"(*File).End":                 method_ptr_File_End,
		"(*File).Line":                method_ptr_File_Line,
		"(*File).LineCount":           method_ptr_File_LineCount,
		"(*File).LineStart":           method_ptr_File_LineStart,
		"(*File).Lines":               method_ptr_File_Lines,
		"(*File).MergeLine":           method_ptr_File_MergeLine,
		"(*File).Name":                method_ptr_File_Name,
		"(*File).Offset":              method_ptr_File_Offset,
		"(*File).Pos":                 method_ptr_File_Pos,
		"(*File).Position":            method_ptr_File_Position,
		"(*File).PositionFor":         method_ptr_File_PositionFor,
		"(*File).SetLines":            method_ptr_File_SetLines,
		"(*File).SetLinesForContent":  method_ptr_File_SetLinesForContent,
		"(*File).Size":                method_ptr_File_Size,
		"(*FileSet).AddExistingFiles": method_ptr_FileSet_AddExistingFiles,
		"(*FileSet).AddFile":          method_ptr_FileSet_AddFile,
		"(*FileSet).Base":             method_ptr_FileSet_Base,
		"(*FileSet).File":             method_ptr_FileSet_File,
		"(*FileSet).Iterate":          method_ptr_FileSet_Iterate,
		"(*FileSet).Position":         method_ptr_FileSet_Position,
		"(*FileSet).PositionFor":      method_ptr_FileSet_PositionFor,
		"(*FileSet).Read":             method_ptr_FileSet_Read,
		"(*FileSet).RemoveFile":       method_ptr_FileSet_RemoveFile,
		"(*FileSet).Write":            method_ptr_FileSet_Write,
		"(*Pos).IsValid":              method_ptr_Pos_IsValid,
		"(*Position).IsValid":         method_ptr_Position_IsValid,
		"(*Position).String":          method_ptr_Position_String,
		"(*Token).IsKeyword":          method_ptr_Token_IsKeyword,
		"(*Token).IsLiteral":          method_ptr_Token_IsLiteral,
		"(*Token).IsOperator":         method_ptr_Token_IsOperator,
		"(*Token).Precedence":         method_ptr_Token_Precedence,
		"(*Token).String":             method_ptr_Token_String,
		"(Pos).IsValid":               method_Pos_IsValid,
		"(Position).String":           method_Position_String,
		"(Token).IsKeyword":           method_Token_IsKeyword,
		"(Token).IsLiteral":           method_Token_IsLiteral,
		"(Token).IsOperator":          method_Token_IsOperator,
		"(Token).Precedence":          method_Token_Precedence,
		"(Token).String":              method_Token_String,
		"IsExported":                  func_IsExported,
		"IsIdentifier":                func_IsIdentifier,
		"IsKeyword":                   func_IsKeyword,
		"Lookup":                      func_Lookup,
		"NewFileSet":                  func_NewFileSet,
	})
}

func method_ptr_File_AddLine(ctx ixgo.DirectCallContext) {
	(*q.File).AddLine(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_File_AddLineColumnInfo(ctx ixgo.DirectCallContext) {
	(*q.File).AddLineColumnInfo(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[int](ctx, 3), ixgo.DirectCallArg[int](ctx, 4))
}

func method_ptr_File_AddLineInfo(ctx ixgo.DirectCallContext) {
	(*q.File).AddLineInfo(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[int](ctx, 3))
}

func method_ptr_File_Base(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Base(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).End(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Line(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Line(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[q.Pos](ctx, 1)))
}

func method_ptr_File_LineCount(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).LineCount(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_LineStart(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).LineStart(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_File_Lines(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Lines(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_MergeLine(ctx ixgo.DirectCallContext) {
	(*q.File).MergeLine(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_File_Name(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Name(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Offset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Offset(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[q.Pos](ctx, 1)))
}

func method_ptr_File_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Pos(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_File_Position(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Position(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[q.Pos](ctx, 1)))
}

func method_ptr_File_PositionFor(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).PositionFor(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[q.Pos](ctx, 1), ixgo.DirectCallArg[bool](ctx, 2)))
}

func method_ptr_File_SetLines(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).SetLines(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[[]int](ctx, 1)))
}

func method_ptr_File_SetLinesForContent(ctx ixgo.DirectCallContext) {
	(*q.File).SetLinesForContent(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1))
}

func method_ptr_File_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Size(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_FileSet_AddExistingFiles(ctx ixgo.DirectCallContext) {
	(*q.FileSet).AddExistingFiles(ixgo.DirectCallArg[*q.FileSet](ctx, 0), ixgo.DirectCallArg[[]*q.File](ctx, 1)...)
}

func method_ptr_FileSet_AddFile(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileSet).AddFile(ixgo.DirectCallArg[*q.FileSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func method_ptr_FileSet_Base(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileSet).Base(ixgo.DirectCallArg[*q.FileSet](ctx, 0)))
}

func method_ptr_FileSet_File(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileSet).File(ixgo.DirectCallArg[*q.FileSet](ctx, 0), ixgo.DirectCallArg[q.Pos](ctx, 1)))
}

func method_ptr_FileSet_Iterate(ctx ixgo.DirectCallContext) {
	(*q.FileSet).Iterate(ixgo.DirectCallArg[*q.FileSet](ctx, 0), ixgo.DirectCallArg[func(*q.File) bool](ctx, 1))
}

func method_ptr_FileSet_Position(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileSet).Position(ixgo.DirectCallArg[*q.FileSet](ctx, 0), ixgo.DirectCallArg[q.Pos](ctx, 1)))
}

func method_ptr_FileSet_PositionFor(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileSet).PositionFor(ixgo.DirectCallArg[*q.FileSet](ctx, 0), ixgo.DirectCallArg[q.Pos](ctx, 1), ixgo.DirectCallArg[bool](ctx, 2)))
}

func method_ptr_FileSet_Read(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileSet).Read(ixgo.DirectCallArg[*q.FileSet](ctx, 0), ixgo.DirectCallArg[func(any) error](ctx, 1)))
}

func method_ptr_FileSet_RemoveFile(ctx ixgo.DirectCallContext) {
	(*q.FileSet).RemoveFile(ixgo.DirectCallArg[*q.FileSet](ctx, 0), ixgo.DirectCallArg[*q.File](ctx, 1))
}

func method_ptr_FileSet_Write(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FileSet).Write(ixgo.DirectCallArg[*q.FileSet](ctx, 0), ixgo.DirectCallArg[func(any) error](ctx, 1)))
}

func func_IsExported(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsExported(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_IsIdentifier(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsIdentifier(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_IsKeyword(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsKeyword(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Lookup(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Lookup(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_NewFileSet(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewFileSet())
}

func method_Pos_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pos.IsValid(ixgo.DirectCallArg[q.Pos](ctx, 0)))
}

func method_ptr_Pos_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pos).IsValid(ixgo.DirectCallArg[*q.Pos](ctx, 0)))
}

func method_ptr_Position_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Position).IsValid(ixgo.DirectCallArg[*q.Position](ctx, 0)))
}

func method_Position_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Position.String(ixgo.DirectCallArg[q.Position](ctx, 0)))
}

func method_ptr_Position_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Position).String(ixgo.DirectCallArg[*q.Position](ctx, 0)))
}

func method_Token_IsKeyword(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Token.IsKeyword(ixgo.DirectCallArg[q.Token](ctx, 0)))
}

func method_ptr_Token_IsKeyword(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Token).IsKeyword(ixgo.DirectCallArg[*q.Token](ctx, 0)))
}

func method_Token_IsLiteral(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Token.IsLiteral(ixgo.DirectCallArg[q.Token](ctx, 0)))
}

func method_ptr_Token_IsLiteral(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Token).IsLiteral(ixgo.DirectCallArg[*q.Token](ctx, 0)))
}

func method_Token_IsOperator(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Token.IsOperator(ixgo.DirectCallArg[q.Token](ctx, 0)))
}

func method_ptr_Token_IsOperator(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Token).IsOperator(ixgo.DirectCallArg[*q.Token](ctx, 0)))
}

func method_Token_Precedence(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Token.Precedence(ixgo.DirectCallArg[q.Token](ctx, 0)))
}

func method_ptr_Token_Precedence(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Token).Precedence(ixgo.DirectCallArg[*q.Token](ctx, 0)))
}

func method_Token_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Token.String(ixgo.DirectCallArg[q.Token](ctx, 0)))
}

func method_ptr_Token_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Token).String(ixgo.DirectCallArg[*q.Token](ctx, 0)))
}
