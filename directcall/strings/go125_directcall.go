// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package strings

import (
	q "strings"

	"github.com/goplus/ixgo"
	unicode "unicode"
)

func init() {
	ixgo.RegisterDirectCalls("strings", map[string]ixgo.DirectCallAdapter{
		"(*Builder).Cap":       method_ptr_Builder_Cap,
		"(*Builder).Grow":      method_ptr_Builder_Grow,
		"(*Builder).Len":       method_ptr_Builder_Len,
		"(*Builder).Reset":     method_ptr_Builder_Reset,
		"(*Builder).String":    method_ptr_Builder_String,
		"(*Builder).WriteByte": method_ptr_Builder_WriteByte,
		"(*Reader).Len":        method_ptr_Reader_Len,
		"(*Reader).Reset":      method_ptr_Reader_Reset,
		"(*Reader).Size":       method_ptr_Reader_Size,
		"(*Reader).UnreadByte": method_ptr_Reader_UnreadByte,
		"(*Reader).UnreadRune": method_ptr_Reader_UnreadRune,
		"(*Replacer).Replace":  method_ptr_Replacer_Replace,
		"Clone":                func_Clone,
		"Compare":              func_Compare,
		"Contains":             func_Contains,
		"ContainsAny":          func_ContainsAny,
		"ContainsFunc":         func_ContainsFunc,
		"ContainsRune":         func_ContainsRune,
		"Count":                func_Count,
		"EqualFold":            func_EqualFold,
		"Fields":               func_Fields,
		"FieldsFunc":           func_FieldsFunc,
		"FieldsFuncSeq":        func_FieldsFuncSeq,
		"FieldsSeq":            func_FieldsSeq,
		"HasPrefix":            func_HasPrefix,
		"HasSuffix":            func_HasSuffix,
		"Index":                func_Index,
		"IndexAny":             func_IndexAny,
		"IndexByte":            func_IndexByte,
		"IndexFunc":            func_IndexFunc,
		"IndexRune":            func_IndexRune,
		"Join":                 func_Join,
		"LastIndex":            func_LastIndex,
		"LastIndexAny":         func_LastIndexAny,
		"LastIndexByte":        func_LastIndexByte,
		"LastIndexFunc":        func_LastIndexFunc,
		"Lines":                func_Lines,
		"Map":                  func_Map,
		"NewReader":            func_NewReader,
		"NewReplacer":          func_NewReplacer,
		"Repeat":               func_Repeat,
		"Replace":              func_Replace,
		"ReplaceAll":           func_ReplaceAll,
		"Split":                func_Split,
		"SplitAfter":           func_SplitAfter,
		"SplitAfterN":          func_SplitAfterN,
		"SplitAfterSeq":        func_SplitAfterSeq,
		"SplitN":               func_SplitN,
		"SplitSeq":             func_SplitSeq,
		"Title":                func_Title,
		"ToLower":              func_ToLower,
		"ToLowerSpecial":       func_ToLowerSpecial,
		"ToTitle":              func_ToTitle,
		"ToTitleSpecial":       func_ToTitleSpecial,
		"ToUpper":              func_ToUpper,
		"ToUpperSpecial":       func_ToUpperSpecial,
		"ToValidUTF8":          func_ToValidUTF8,
		"Trim":                 func_Trim,
		"TrimFunc":             func_TrimFunc,
		"TrimLeft":             func_TrimLeft,
		"TrimLeftFunc":         func_TrimLeftFunc,
		"TrimPrefix":           func_TrimPrefix,
		"TrimRight":            func_TrimRight,
		"TrimRightFunc":        func_TrimRightFunc,
		"TrimSpace":            func_TrimSpace,
		"TrimSuffix":           func_TrimSuffix,
	})
}

func method_ptr_Builder_Cap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Builder).Cap(ixgo.DirectCallArg[*q.Builder](ctx, 0)))
}

func method_ptr_Builder_Grow(ctx ixgo.DirectCallContext) {
	(*q.Builder).Grow(ixgo.DirectCallArg[*q.Builder](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_Builder_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Builder).Len(ixgo.DirectCallArg[*q.Builder](ctx, 0)))
}

func method_ptr_Builder_Reset(ctx ixgo.DirectCallContext) {
	(*q.Builder).Reset(ixgo.DirectCallArg[*q.Builder](ctx, 0))
}

func method_ptr_Builder_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Builder).String(ixgo.DirectCallArg[*q.Builder](ctx, 0)))
}

func method_ptr_Builder_WriteByte(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Builder).WriteByte(ixgo.DirectCallArg[*q.Builder](ctx, 0), ixgo.DirectCallArg[byte](ctx, 1)))
}

func func_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Clone(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Compare(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Contains(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Contains(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_ContainsAny(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ContainsAny(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_ContainsFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ContainsFunc(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[func(rune) bool](ctx, 1)))
}

func func_ContainsRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ContainsRune(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_Count(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Count(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_EqualFold(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EqualFold(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Fields(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Fields(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_FieldsFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FieldsFunc(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[func(rune) bool](ctx, 1)))
}

func func_FieldsFuncSeq(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FieldsFuncSeq(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[func(rune) bool](ctx, 1)))
}

func func_FieldsSeq(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FieldsSeq(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_HasPrefix(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HasPrefix(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_HasSuffix(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HasSuffix(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Index(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Index(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_IndexAny(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IndexAny(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_IndexByte(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IndexByte(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[byte](ctx, 1)))
}

func func_IndexFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IndexFunc(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[func(rune) bool](ctx, 1)))
}

func func_IndexRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IndexRune(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func func_Join(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Join(ixgo.DirectCallArg[[]string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_LastIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LastIndex(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_LastIndexAny(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LastIndexAny(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_LastIndexByte(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LastIndexByte(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[byte](ctx, 1)))
}

func func_LastIndexFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LastIndexFunc(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[func(rune) bool](ctx, 1)))
}

func func_Lines(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Lines(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Map(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Map(ixgo.DirectCallArg[func(rune) rune](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_NewReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReader(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_NewReplacer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReplacer(ixgo.DirectCallArg[[]string](ctx, 0)...))
}

func method_ptr_Reader_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).Len(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Reader_Reset(ctx ixgo.DirectCallContext) {
	(*q.Reader).Reset(ixgo.DirectCallArg[*q.Reader](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_Reader_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).Size(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Reader_UnreadByte(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).UnreadByte(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Reader_UnreadRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).UnreadRune(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func func_Repeat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Repeat(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_Replace(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Replace(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func func_ReplaceAll(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ReplaceAll(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_Replacer_Replace(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Replacer).Replace(ixgo.DirectCallArg[*q.Replacer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Split(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Split(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_SplitAfter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SplitAfter(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_SplitAfterN(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SplitAfterN(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_SplitAfterSeq(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SplitAfterSeq(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_SplitN(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SplitN(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_SplitSeq(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SplitSeq(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Title(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Title(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_ToLower(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToLower(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_ToLowerSpecial(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToLowerSpecial(ixgo.DirectCallArg[unicode.SpecialCase](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_ToTitle(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToTitle(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_ToTitleSpecial(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToTitleSpecial(ixgo.DirectCallArg[unicode.SpecialCase](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_ToUpper(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToUpper(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_ToUpperSpecial(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToUpperSpecial(ixgo.DirectCallArg[unicode.SpecialCase](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_ToValidUTF8(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ToValidUTF8(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Trim(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Trim(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_TrimFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrimFunc(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[func(rune) bool](ctx, 1)))
}

func func_TrimLeft(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrimLeft(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_TrimLeftFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrimLeftFunc(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[func(rune) bool](ctx, 1)))
}

func func_TrimPrefix(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrimPrefix(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_TrimRight(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrimRight(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_TrimRightFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrimRightFunc(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[func(rune) bool](ctx, 1)))
}

func func_TrimSpace(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrimSpace(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_TrimSuffix(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrimSuffix(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}
