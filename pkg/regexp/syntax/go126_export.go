// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package syntax

import (
	q "regexp/syntax"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "syntax",
		Path: "regexp/syntax",
		Deps: map[string]string{
			"slices":       "slices",
			"sort":         "sort",
			"strconv":      "strconv",
			"strings":      "strings",
			"sync":         "sync",
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
			"ClassNL":                  {reflect.TypeOf(q.ClassNL), constant.MakeInt64(int64(q.ClassNL))},
			"DotNL":                    {reflect.TypeOf(q.DotNL), constant.MakeInt64(int64(q.DotNL))},
			"EmptyBeginLine":           {reflect.TypeOf(q.EmptyBeginLine), constant.MakeInt64(int64(q.EmptyBeginLine))},
			"EmptyBeginText":           {reflect.TypeOf(q.EmptyBeginText), constant.MakeInt64(int64(q.EmptyBeginText))},
			"EmptyEndLine":             {reflect.TypeOf(q.EmptyEndLine), constant.MakeInt64(int64(q.EmptyEndLine))},
			"EmptyEndText":             {reflect.TypeOf(q.EmptyEndText), constant.MakeInt64(int64(q.EmptyEndText))},
			"EmptyNoWordBoundary":      {reflect.TypeOf(q.EmptyNoWordBoundary), constant.MakeInt64(int64(q.EmptyNoWordBoundary))},
			"EmptyWordBoundary":        {reflect.TypeOf(q.EmptyWordBoundary), constant.MakeInt64(int64(q.EmptyWordBoundary))},
			"ErrInternalError":         {reflect.TypeOf(q.ErrInternalError), constant.MakeString(string(q.ErrInternalError))},
			"ErrInvalidCharClass":      {reflect.TypeOf(q.ErrInvalidCharClass), constant.MakeString(string(q.ErrInvalidCharClass))},
			"ErrInvalidCharRange":      {reflect.TypeOf(q.ErrInvalidCharRange), constant.MakeString(string(q.ErrInvalidCharRange))},
			"ErrInvalidEscape":         {reflect.TypeOf(q.ErrInvalidEscape), constant.MakeString(string(q.ErrInvalidEscape))},
			"ErrInvalidNamedCapture":   {reflect.TypeOf(q.ErrInvalidNamedCapture), constant.MakeString(string(q.ErrInvalidNamedCapture))},
			"ErrInvalidPerlOp":         {reflect.TypeOf(q.ErrInvalidPerlOp), constant.MakeString(string(q.ErrInvalidPerlOp))},
			"ErrInvalidRepeatOp":       {reflect.TypeOf(q.ErrInvalidRepeatOp), constant.MakeString(string(q.ErrInvalidRepeatOp))},
			"ErrInvalidRepeatSize":     {reflect.TypeOf(q.ErrInvalidRepeatSize), constant.MakeString(string(q.ErrInvalidRepeatSize))},
			"ErrInvalidUTF8":           {reflect.TypeOf(q.ErrInvalidUTF8), constant.MakeString(string(q.ErrInvalidUTF8))},
			"ErrLarge":                 {reflect.TypeOf(q.ErrLarge), constant.MakeString(string(q.ErrLarge))},
			"ErrMissingBracket":        {reflect.TypeOf(q.ErrMissingBracket), constant.MakeString(string(q.ErrMissingBracket))},
			"ErrMissingParen":          {reflect.TypeOf(q.ErrMissingParen), constant.MakeString(string(q.ErrMissingParen))},
			"ErrMissingRepeatArgument": {reflect.TypeOf(q.ErrMissingRepeatArgument), constant.MakeString(string(q.ErrMissingRepeatArgument))},
			"ErrNestingDepth":          {reflect.TypeOf(q.ErrNestingDepth), constant.MakeString(string(q.ErrNestingDepth))},
			"ErrTrailingBackslash":     {reflect.TypeOf(q.ErrTrailingBackslash), constant.MakeString(string(q.ErrTrailingBackslash))},
			"ErrUnexpectedParen":       {reflect.TypeOf(q.ErrUnexpectedParen), constant.MakeString(string(q.ErrUnexpectedParen))},
			"FoldCase":                 {reflect.TypeOf(q.FoldCase), constant.MakeInt64(int64(q.FoldCase))},
			"InstAlt":                  {reflect.TypeOf(q.InstAlt), constant.MakeInt64(int64(q.InstAlt))},
			"InstAltMatch":             {reflect.TypeOf(q.InstAltMatch), constant.MakeInt64(int64(q.InstAltMatch))},
			"InstCapture":              {reflect.TypeOf(q.InstCapture), constant.MakeInt64(int64(q.InstCapture))},
			"InstEmptyWidth":           {reflect.TypeOf(q.InstEmptyWidth), constant.MakeInt64(int64(q.InstEmptyWidth))},
			"InstFail":                 {reflect.TypeOf(q.InstFail), constant.MakeInt64(int64(q.InstFail))},
			"InstMatch":                {reflect.TypeOf(q.InstMatch), constant.MakeInt64(int64(q.InstMatch))},
			"InstNop":                  {reflect.TypeOf(q.InstNop), constant.MakeInt64(int64(q.InstNop))},
			"InstRune":                 {reflect.TypeOf(q.InstRune), constant.MakeInt64(int64(q.InstRune))},
			"InstRune1":                {reflect.TypeOf(q.InstRune1), constant.MakeInt64(int64(q.InstRune1))},
			"InstRuneAny":              {reflect.TypeOf(q.InstRuneAny), constant.MakeInt64(int64(q.InstRuneAny))},
			"InstRuneAnyNotNL":         {reflect.TypeOf(q.InstRuneAnyNotNL), constant.MakeInt64(int64(q.InstRuneAnyNotNL))},
			"Literal":                  {reflect.TypeOf(q.Literal), constant.MakeInt64(int64(q.Literal))},
			"MatchNL":                  {reflect.TypeOf(q.MatchNL), constant.MakeInt64(int64(q.MatchNL))},
			"NonGreedy":                {reflect.TypeOf(q.NonGreedy), constant.MakeInt64(int64(q.NonGreedy))},
			"OneLine":                  {reflect.TypeOf(q.OneLine), constant.MakeInt64(int64(q.OneLine))},
			"OpAlternate":              {reflect.TypeOf(q.OpAlternate), constant.MakeInt64(int64(q.OpAlternate))},
			"OpAnyChar":                {reflect.TypeOf(q.OpAnyChar), constant.MakeInt64(int64(q.OpAnyChar))},
			"OpAnyCharNotNL":           {reflect.TypeOf(q.OpAnyCharNotNL), constant.MakeInt64(int64(q.OpAnyCharNotNL))},
			"OpBeginLine":              {reflect.TypeOf(q.OpBeginLine), constant.MakeInt64(int64(q.OpBeginLine))},
			"OpBeginText":              {reflect.TypeOf(q.OpBeginText), constant.MakeInt64(int64(q.OpBeginText))},
			"OpCapture":                {reflect.TypeOf(q.OpCapture), constant.MakeInt64(int64(q.OpCapture))},
			"OpCharClass":              {reflect.TypeOf(q.OpCharClass), constant.MakeInt64(int64(q.OpCharClass))},
			"OpConcat":                 {reflect.TypeOf(q.OpConcat), constant.MakeInt64(int64(q.OpConcat))},
			"OpEmptyMatch":             {reflect.TypeOf(q.OpEmptyMatch), constant.MakeInt64(int64(q.OpEmptyMatch))},
			"OpEndLine":                {reflect.TypeOf(q.OpEndLine), constant.MakeInt64(int64(q.OpEndLine))},
			"OpEndText":                {reflect.TypeOf(q.OpEndText), constant.MakeInt64(int64(q.OpEndText))},
			"OpLiteral":                {reflect.TypeOf(q.OpLiteral), constant.MakeInt64(int64(q.OpLiteral))},
			"OpNoMatch":                {reflect.TypeOf(q.OpNoMatch), constant.MakeInt64(int64(q.OpNoMatch))},
			"OpNoWordBoundary":         {reflect.TypeOf(q.OpNoWordBoundary), constant.MakeInt64(int64(q.OpNoWordBoundary))},
			"OpPlus":                   {reflect.TypeOf(q.OpPlus), constant.MakeInt64(int64(q.OpPlus))},
			"OpQuest":                  {reflect.TypeOf(q.OpQuest), constant.MakeInt64(int64(q.OpQuest))},
			"OpRepeat":                 {reflect.TypeOf(q.OpRepeat), constant.MakeInt64(int64(q.OpRepeat))},
			"OpStar":                   {reflect.TypeOf(q.OpStar), constant.MakeInt64(int64(q.OpStar))},
			"OpWordBoundary":           {reflect.TypeOf(q.OpWordBoundary), constant.MakeInt64(int64(q.OpWordBoundary))},
			"POSIX":                    {reflect.TypeOf(q.POSIX), constant.MakeInt64(int64(q.POSIX))},
			"Perl":                     {reflect.TypeOf(q.Perl), constant.MakeInt64(int64(q.Perl))},
			"PerlX":                    {reflect.TypeOf(q.PerlX), constant.MakeInt64(int64(q.PerlX))},
			"Simple":                   {reflect.TypeOf(q.Simple), constant.MakeInt64(int64(q.Simple))},
			"UnicodeGroups":            {reflect.TypeOf(q.UnicodeGroups), constant.MakeInt64(int64(q.UnicodeGroups))},
			"WasDollar":                {reflect.TypeOf(q.WasDollar), constant.MakeInt64(int64(q.WasDollar))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
