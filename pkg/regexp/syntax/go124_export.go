// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package syntax

import (
	q "regexp/syntax"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("regexp/syntax", func() *ixgo.Package {
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
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"Compile":        reflect.ValueOf(q.Compile),
				"EmptyOpContext": reflect.ValueOf(q.EmptyOpContext),
				"IsWordChar":     reflect.ValueOf(q.IsWordChar),
				"Parse":          reflect.ValueOf(q.Parse),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"ClassNL":                  {Typ: reflect.TypeOf(q.ClassNL), Value: constant.MakeInt64(int64(q.ClassNL))},
				"DotNL":                    {Typ: reflect.TypeOf(q.DotNL), Value: constant.MakeInt64(int64(q.DotNL))},
				"EmptyBeginLine":           {Typ: reflect.TypeOf(q.EmptyBeginLine), Value: constant.MakeInt64(int64(q.EmptyBeginLine))},
				"EmptyBeginText":           {Typ: reflect.TypeOf(q.EmptyBeginText), Value: constant.MakeInt64(int64(q.EmptyBeginText))},
				"EmptyEndLine":             {Typ: reflect.TypeOf(q.EmptyEndLine), Value: constant.MakeInt64(int64(q.EmptyEndLine))},
				"EmptyEndText":             {Typ: reflect.TypeOf(q.EmptyEndText), Value: constant.MakeInt64(int64(q.EmptyEndText))},
				"EmptyNoWordBoundary":      {Typ: reflect.TypeOf(q.EmptyNoWordBoundary), Value: constant.MakeInt64(int64(q.EmptyNoWordBoundary))},
				"EmptyWordBoundary":        {Typ: reflect.TypeOf(q.EmptyWordBoundary), Value: constant.MakeInt64(int64(q.EmptyWordBoundary))},
				"ErrInternalError":         {Typ: reflect.TypeOf(q.ErrInternalError), Value: constant.MakeString(string(q.ErrInternalError))},
				"ErrInvalidCharClass":      {Typ: reflect.TypeOf(q.ErrInvalidCharClass), Value: constant.MakeString(string(q.ErrInvalidCharClass))},
				"ErrInvalidCharRange":      {Typ: reflect.TypeOf(q.ErrInvalidCharRange), Value: constant.MakeString(string(q.ErrInvalidCharRange))},
				"ErrInvalidEscape":         {Typ: reflect.TypeOf(q.ErrInvalidEscape), Value: constant.MakeString(string(q.ErrInvalidEscape))},
				"ErrInvalidNamedCapture":   {Typ: reflect.TypeOf(q.ErrInvalidNamedCapture), Value: constant.MakeString(string(q.ErrInvalidNamedCapture))},
				"ErrInvalidPerlOp":         {Typ: reflect.TypeOf(q.ErrInvalidPerlOp), Value: constant.MakeString(string(q.ErrInvalidPerlOp))},
				"ErrInvalidRepeatOp":       {Typ: reflect.TypeOf(q.ErrInvalidRepeatOp), Value: constant.MakeString(string(q.ErrInvalidRepeatOp))},
				"ErrInvalidRepeatSize":     {Typ: reflect.TypeOf(q.ErrInvalidRepeatSize), Value: constant.MakeString(string(q.ErrInvalidRepeatSize))},
				"ErrInvalidUTF8":           {Typ: reflect.TypeOf(q.ErrInvalidUTF8), Value: constant.MakeString(string(q.ErrInvalidUTF8))},
				"ErrLarge":                 {Typ: reflect.TypeOf(q.ErrLarge), Value: constant.MakeString(string(q.ErrLarge))},
				"ErrMissingBracket":        {Typ: reflect.TypeOf(q.ErrMissingBracket), Value: constant.MakeString(string(q.ErrMissingBracket))},
				"ErrMissingParen":          {Typ: reflect.TypeOf(q.ErrMissingParen), Value: constant.MakeString(string(q.ErrMissingParen))},
				"ErrMissingRepeatArgument": {Typ: reflect.TypeOf(q.ErrMissingRepeatArgument), Value: constant.MakeString(string(q.ErrMissingRepeatArgument))},
				"ErrNestingDepth":          {Typ: reflect.TypeOf(q.ErrNestingDepth), Value: constant.MakeString(string(q.ErrNestingDepth))},
				"ErrTrailingBackslash":     {Typ: reflect.TypeOf(q.ErrTrailingBackslash), Value: constant.MakeString(string(q.ErrTrailingBackslash))},
				"ErrUnexpectedParen":       {Typ: reflect.TypeOf(q.ErrUnexpectedParen), Value: constant.MakeString(string(q.ErrUnexpectedParen))},
				"FoldCase":                 {Typ: reflect.TypeOf(q.FoldCase), Value: constant.MakeInt64(int64(q.FoldCase))},
				"InstAlt":                  {Typ: reflect.TypeOf(q.InstAlt), Value: constant.MakeInt64(int64(q.InstAlt))},
				"InstAltMatch":             {Typ: reflect.TypeOf(q.InstAltMatch), Value: constant.MakeInt64(int64(q.InstAltMatch))},
				"InstCapture":              {Typ: reflect.TypeOf(q.InstCapture), Value: constant.MakeInt64(int64(q.InstCapture))},
				"InstEmptyWidth":           {Typ: reflect.TypeOf(q.InstEmptyWidth), Value: constant.MakeInt64(int64(q.InstEmptyWidth))},
				"InstFail":                 {Typ: reflect.TypeOf(q.InstFail), Value: constant.MakeInt64(int64(q.InstFail))},
				"InstMatch":                {Typ: reflect.TypeOf(q.InstMatch), Value: constant.MakeInt64(int64(q.InstMatch))},
				"InstNop":                  {Typ: reflect.TypeOf(q.InstNop), Value: constant.MakeInt64(int64(q.InstNop))},
				"InstRune":                 {Typ: reflect.TypeOf(q.InstRune), Value: constant.MakeInt64(int64(q.InstRune))},
				"InstRune1":                {Typ: reflect.TypeOf(q.InstRune1), Value: constant.MakeInt64(int64(q.InstRune1))},
				"InstRuneAny":              {Typ: reflect.TypeOf(q.InstRuneAny), Value: constant.MakeInt64(int64(q.InstRuneAny))},
				"InstRuneAnyNotNL":         {Typ: reflect.TypeOf(q.InstRuneAnyNotNL), Value: constant.MakeInt64(int64(q.InstRuneAnyNotNL))},
				"Literal":                  {Typ: reflect.TypeOf(q.Literal), Value: constant.MakeInt64(int64(q.Literal))},
				"MatchNL":                  {Typ: reflect.TypeOf(q.MatchNL), Value: constant.MakeInt64(int64(q.MatchNL))},
				"NonGreedy":                {Typ: reflect.TypeOf(q.NonGreedy), Value: constant.MakeInt64(int64(q.NonGreedy))},
				"OneLine":                  {Typ: reflect.TypeOf(q.OneLine), Value: constant.MakeInt64(int64(q.OneLine))},
				"OpAlternate":              {Typ: reflect.TypeOf(q.OpAlternate), Value: constant.MakeInt64(int64(q.OpAlternate))},
				"OpAnyChar":                {Typ: reflect.TypeOf(q.OpAnyChar), Value: constant.MakeInt64(int64(q.OpAnyChar))},
				"OpAnyCharNotNL":           {Typ: reflect.TypeOf(q.OpAnyCharNotNL), Value: constant.MakeInt64(int64(q.OpAnyCharNotNL))},
				"OpBeginLine":              {Typ: reflect.TypeOf(q.OpBeginLine), Value: constant.MakeInt64(int64(q.OpBeginLine))},
				"OpBeginText":              {Typ: reflect.TypeOf(q.OpBeginText), Value: constant.MakeInt64(int64(q.OpBeginText))},
				"OpCapture":                {Typ: reflect.TypeOf(q.OpCapture), Value: constant.MakeInt64(int64(q.OpCapture))},
				"OpCharClass":              {Typ: reflect.TypeOf(q.OpCharClass), Value: constant.MakeInt64(int64(q.OpCharClass))},
				"OpConcat":                 {Typ: reflect.TypeOf(q.OpConcat), Value: constant.MakeInt64(int64(q.OpConcat))},
				"OpEmptyMatch":             {Typ: reflect.TypeOf(q.OpEmptyMatch), Value: constant.MakeInt64(int64(q.OpEmptyMatch))},
				"OpEndLine":                {Typ: reflect.TypeOf(q.OpEndLine), Value: constant.MakeInt64(int64(q.OpEndLine))},
				"OpEndText":                {Typ: reflect.TypeOf(q.OpEndText), Value: constant.MakeInt64(int64(q.OpEndText))},
				"OpLiteral":                {Typ: reflect.TypeOf(q.OpLiteral), Value: constant.MakeInt64(int64(q.OpLiteral))},
				"OpNoMatch":                {Typ: reflect.TypeOf(q.OpNoMatch), Value: constant.MakeInt64(int64(q.OpNoMatch))},
				"OpNoWordBoundary":         {Typ: reflect.TypeOf(q.OpNoWordBoundary), Value: constant.MakeInt64(int64(q.OpNoWordBoundary))},
				"OpPlus":                   {Typ: reflect.TypeOf(q.OpPlus), Value: constant.MakeInt64(int64(q.OpPlus))},
				"OpQuest":                  {Typ: reflect.TypeOf(q.OpQuest), Value: constant.MakeInt64(int64(q.OpQuest))},
				"OpRepeat":                 {Typ: reflect.TypeOf(q.OpRepeat), Value: constant.MakeInt64(int64(q.OpRepeat))},
				"OpStar":                   {Typ: reflect.TypeOf(q.OpStar), Value: constant.MakeInt64(int64(q.OpStar))},
				"OpWordBoundary":           {Typ: reflect.TypeOf(q.OpWordBoundary), Value: constant.MakeInt64(int64(q.OpWordBoundary))},
				"POSIX":                    {Typ: reflect.TypeOf(q.POSIX), Value: constant.MakeInt64(int64(q.POSIX))},
				"Perl":                     {Typ: reflect.TypeOf(q.Perl), Value: constant.MakeInt64(int64(q.Perl))},
				"PerlX":                    {Typ: reflect.TypeOf(q.PerlX), Value: constant.MakeInt64(int64(q.PerlX))},
				"Simple":                   {Typ: reflect.TypeOf(q.Simple), Value: constant.MakeInt64(int64(q.Simple))},
				"UnicodeGroups":            {Typ: reflect.TypeOf(q.UnicodeGroups), Value: constant.MakeInt64(int64(q.UnicodeGroups))},
				"WasDollar":                {Typ: reflect.TypeOf(q.WasDollar), Value: constant.MakeInt64(int64(q.WasDollar))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
