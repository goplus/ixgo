// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package syntax

import (
	q "regexp/syntax"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("regexp/syntax", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "syntax",
			Path: "regexp/syntax",
			Deps: map[string]string{
				"slices":       "slices",
				"sort":         "sort",
				"strconv":      "strconv",
				"strings":      "strings",
				"unicode":      "unicode",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"EmptyOp":   reflect.TypeOf((*q.EmptyOp)(nil)).Elem(),
				"Error":     reflect.TypeOf((*q.Error)(nil)).Elem(),
				"ErrorCode": reflect.TypeOf((*q.ErrorCode)(nil)).Elem(),
				"Flags":     reflect.TypeOf((*q.Flags)(nil)).Elem(),
				"Inst":      reflect.TypeOf((*q.Inst)(nil)).Elem(),
				"InstOp":    reflect.TypeOf((*q.InstOp)(nil)).Elem(),
				"Op":        reflect.TypeOf((*q.Op)(nil)).Elem(),
				"Prog":      reflect.TypeOf((*q.Prog)(nil)).Elem(),
				"Regexp":    reflect.TypeOf((*q.Regexp)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Compile":        q.Compile,
				"EmptyOpContext": q.EmptyOpContext,
				"IsWordChar":     q.IsWordChar,
				"Parse":          q.Parse,
			},
			TypedConsts: map[string]interface{}{
				"ClassNL":                  q.ClassNL,
				"DotNL":                    q.DotNL,
				"EmptyBeginLine":           q.EmptyBeginLine,
				"EmptyBeginText":           q.EmptyBeginText,
				"EmptyEndLine":             q.EmptyEndLine,
				"EmptyEndText":             q.EmptyEndText,
				"EmptyNoWordBoundary":      q.EmptyNoWordBoundary,
				"EmptyWordBoundary":        q.EmptyWordBoundary,
				"ErrInternalError":         q.ErrInternalError,
				"ErrInvalidCharClass":      q.ErrInvalidCharClass,
				"ErrInvalidCharRange":      q.ErrInvalidCharRange,
				"ErrInvalidEscape":         q.ErrInvalidEscape,
				"ErrInvalidNamedCapture":   q.ErrInvalidNamedCapture,
				"ErrInvalidPerlOp":         q.ErrInvalidPerlOp,
				"ErrInvalidRepeatOp":       q.ErrInvalidRepeatOp,
				"ErrInvalidRepeatSize":     q.ErrInvalidRepeatSize,
				"ErrInvalidUTF8":           q.ErrInvalidUTF8,
				"ErrLarge":                 q.ErrLarge,
				"ErrMissingBracket":        q.ErrMissingBracket,
				"ErrMissingParen":          q.ErrMissingParen,
				"ErrMissingRepeatArgument": q.ErrMissingRepeatArgument,
				"ErrNestingDepth":          q.ErrNestingDepth,
				"ErrTrailingBackslash":     q.ErrTrailingBackslash,
				"ErrUnexpectedParen":       q.ErrUnexpectedParen,
				"FoldCase":                 q.FoldCase,
				"InstAlt":                  q.InstAlt,
				"InstAltMatch":             q.InstAltMatch,
				"InstCapture":              q.InstCapture,
				"InstEmptyWidth":           q.InstEmptyWidth,
				"InstFail":                 q.InstFail,
				"InstMatch":                q.InstMatch,
				"InstNop":                  q.InstNop,
				"InstRune":                 q.InstRune,
				"InstRune1":                q.InstRune1,
				"InstRuneAny":              q.InstRuneAny,
				"InstRuneAnyNotNL":         q.InstRuneAnyNotNL,
				"Literal":                  q.Literal,
				"MatchNL":                  q.MatchNL,
				"NonGreedy":                q.NonGreedy,
				"OneLine":                  q.OneLine,
				"OpAlternate":              q.OpAlternate,
				"OpAnyChar":                q.OpAnyChar,
				"OpAnyCharNotNL":           q.OpAnyCharNotNL,
				"OpBeginLine":              q.OpBeginLine,
				"OpBeginText":              q.OpBeginText,
				"OpCapture":                q.OpCapture,
				"OpCharClass":              q.OpCharClass,
				"OpConcat":                 q.OpConcat,
				"OpEmptyMatch":             q.OpEmptyMatch,
				"OpEndLine":                q.OpEndLine,
				"OpEndText":                q.OpEndText,
				"OpLiteral":                q.OpLiteral,
				"OpNoMatch":                q.OpNoMatch,
				"OpNoWordBoundary":         q.OpNoWordBoundary,
				"OpPlus":                   q.OpPlus,
				"OpQuest":                  q.OpQuest,
				"OpRepeat":                 q.OpRepeat,
				"OpStar":                   q.OpStar,
				"OpWordBoundary":           q.OpWordBoundary,
				"POSIX":                    q.POSIX,
				"Perl":                     q.Perl,
				"PerlX":                    q.PerlX,
				"Simple":                   q.Simple,
				"UnicodeGroups":            q.UnicodeGroups,
				"WasDollar":                q.WasDollar,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
