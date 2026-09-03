// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package regexp

import (
	q "regexp"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("regexp", map[string]ixgo.DirectCallAdapter{
		"(*Regexp).Copy":                       method_ptr_Regexp_Copy,
		"(*Regexp).Expand":                     method_ptr_Regexp_Expand,
		"(*Regexp).ExpandString":               method_ptr_Regexp_ExpandString,
		"(*Regexp).Find":                       method_ptr_Regexp_Find,
		"(*Regexp).FindAll":                    method_ptr_Regexp_FindAll,
		"(*Regexp).FindAllIndex":               method_ptr_Regexp_FindAllIndex,
		"(*Regexp).FindAllString":              method_ptr_Regexp_FindAllString,
		"(*Regexp).FindAllStringIndex":         method_ptr_Regexp_FindAllStringIndex,
		"(*Regexp).FindAllStringSubmatch":      method_ptr_Regexp_FindAllStringSubmatch,
		"(*Regexp).FindAllStringSubmatchIndex": method_ptr_Regexp_FindAllStringSubmatchIndex,
		"(*Regexp).FindAllSubmatch":            method_ptr_Regexp_FindAllSubmatch,
		"(*Regexp).FindAllSubmatchIndex":       method_ptr_Regexp_FindAllSubmatchIndex,
		"(*Regexp).FindIndex":                  method_ptr_Regexp_FindIndex,
		"(*Regexp).FindReaderIndex":            method_ptr_Regexp_FindReaderIndex,
		"(*Regexp).FindReaderSubmatchIndex":    method_ptr_Regexp_FindReaderSubmatchIndex,
		"(*Regexp).FindString":                 method_ptr_Regexp_FindString,
		"(*Regexp).FindStringIndex":            method_ptr_Regexp_FindStringIndex,
		"(*Regexp).FindStringSubmatch":         method_ptr_Regexp_FindStringSubmatch,
		"(*Regexp).FindStringSubmatchIndex":    method_ptr_Regexp_FindStringSubmatchIndex,
		"(*Regexp).FindSubmatch":               method_ptr_Regexp_FindSubmatch,
		"(*Regexp).FindSubmatchIndex":          method_ptr_Regexp_FindSubmatchIndex,
		"(*Regexp).Longest":                    method_ptr_Regexp_Longest,
		"(*Regexp).Match":                      method_ptr_Regexp_Match,
		"(*Regexp).MatchReader":                method_ptr_Regexp_MatchReader,
		"(*Regexp).MatchString":                method_ptr_Regexp_MatchString,
		"(*Regexp).NumSubexp":                  method_ptr_Regexp_NumSubexp,
		"(*Regexp).ReplaceAll":                 method_ptr_Regexp_ReplaceAll,
		"(*Regexp).ReplaceAllFunc":             method_ptr_Regexp_ReplaceAllFunc,
		"(*Regexp).ReplaceAllLiteral":          method_ptr_Regexp_ReplaceAllLiteral,
		"(*Regexp).ReplaceAllLiteralString":    method_ptr_Regexp_ReplaceAllLiteralString,
		"(*Regexp).ReplaceAllString":           method_ptr_Regexp_ReplaceAllString,
		"(*Regexp).ReplaceAllStringFunc":       method_ptr_Regexp_ReplaceAllStringFunc,
		"(*Regexp).Split":                      method_ptr_Regexp_Split,
		"(*Regexp).String":                     method_ptr_Regexp_String,
		"(*Regexp).SubexpIndex":                method_ptr_Regexp_SubexpIndex,
		"(*Regexp).SubexpNames":                method_ptr_Regexp_SubexpNames,
		"(*Regexp).UnmarshalText":              method_ptr_Regexp_UnmarshalText,
		"MustCompile":                          func_MustCompile,
		"MustCompilePOSIX":                     func_MustCompilePOSIX,
		"QuoteMeta":                            func_QuoteMeta,
	})
}

func func_MustCompile(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MustCompile(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_MustCompilePOSIX(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MustCompilePOSIX(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_QuoteMeta(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.QuoteMeta(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_Regexp_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).Copy(ixgo.DirectCallArg[*q.Regexp](ctx, 0)))
}

func method_ptr_Regexp_Expand(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).Expand(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2), ixgo.DirectCallArg[[]byte](ctx, 3), ixgo.DirectCallArg[[]int](ctx, 4)))
}

func method_ptr_Regexp_ExpandString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).ExpandString(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[string](ctx, 3), ixgo.DirectCallArg[[]int](ctx, 4)))
}

func method_ptr_Regexp_Find(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).Find(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Regexp_FindAll(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindAll(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Regexp_FindAllIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindAllIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Regexp_FindAllString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindAllString(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Regexp_FindAllStringIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindAllStringIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Regexp_FindAllStringSubmatch(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindAllStringSubmatch(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Regexp_FindAllStringSubmatchIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindAllStringSubmatchIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Regexp_FindAllSubmatch(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindAllSubmatch(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Regexp_FindAllSubmatchIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindAllSubmatchIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Regexp_FindIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Regexp_FindReaderIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindReaderIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[io.RuneReader](ctx, 1)))
}

func method_ptr_Regexp_FindReaderSubmatchIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindReaderSubmatchIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[io.RuneReader](ctx, 1)))
}

func method_ptr_Regexp_FindString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindString(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Regexp_FindStringIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindStringIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Regexp_FindStringSubmatch(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindStringSubmatch(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Regexp_FindStringSubmatchIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindStringSubmatchIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Regexp_FindSubmatch(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindSubmatch(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Regexp_FindSubmatchIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).FindSubmatchIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Regexp_Longest(ctx ixgo.DirectCallContext) {
	(*q.Regexp).Longest(ixgo.DirectCallArg[*q.Regexp](ctx, 0))
}

func method_ptr_Regexp_Match(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).Match(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Regexp_MatchReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).MatchReader(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[io.RuneReader](ctx, 1)))
}

func method_ptr_Regexp_MatchString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).MatchString(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Regexp_NumSubexp(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).NumSubexp(ixgo.DirectCallArg[*q.Regexp](ctx, 0)))
}

func method_ptr_Regexp_ReplaceAll(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).ReplaceAll(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}

func method_ptr_Regexp_ReplaceAllFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).ReplaceAllFunc(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[func([]byte) []byte](ctx, 2)))
}

func method_ptr_Regexp_ReplaceAllLiteral(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).ReplaceAllLiteral(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}

func method_ptr_Regexp_ReplaceAllLiteralString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).ReplaceAllLiteralString(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_Regexp_ReplaceAllString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).ReplaceAllString(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_Regexp_ReplaceAllStringFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).ReplaceAllStringFunc(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[func(string) string](ctx, 2)))
}

func method_ptr_Regexp_Split(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).Split(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Regexp_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).String(ixgo.DirectCallArg[*q.Regexp](ctx, 0)))
}

func method_ptr_Regexp_SubexpIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).SubexpIndex(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Regexp_SubexpNames(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).SubexpNames(ixgo.DirectCallArg[*q.Regexp](ctx, 0)))
}

func method_ptr_Regexp_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).UnmarshalText(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}
