// export by github.com/goplus/ixgo/cmd/qexp

package token

import (
	q "github.com/goplus/xgo/tpl/token"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/token", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "token",
			Path: "github.com/goplus/xgo/tpl/token",
			Deps: map[string]string{
				"go/token": "token",
				"strconv":  "strconv",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Token": reflect.TypeOf((*q.Token)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"File":     reflect.TypeOf((*q.File)(nil)).Elem(),
				"FileSet":  reflect.TypeOf((*q.FileSet)(nil)).Elem(),
				"Pos":      reflect.TypeOf((*q.Pos)(nil)).Elem(),
				"Position": reflect.TypeOf((*q.Position)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"ForEach":    q.ForEach,
				"NewFileSet": q.NewFileSet,
			},
			TypedConsts: map[string]interface{}{
				"ADD_ASSIGN":     q.ADD_ASSIGN,
				"AND_ASSIGN":     q.AND_ASSIGN,
				"AND_NOT":        q.AND_NOT,
				"AND_NOT_ASSIGN": q.AND_NOT_ASSIGN,
				"ARROW":          q.ARROW,
				"BIDIARROW":      q.BIDIARROW,
				"CHAR":           q.CHAR,
				"COMMENT":        q.COMMENT,
				"DEC":            q.DEC,
				"DEFINE":         q.DEFINE,
				"DRARROW":        q.DRARROW,
				"ELLIPSIS":       q.ELLIPSIS,
				"EOF":            q.EOF,
				"EQ":             q.EQ,
				"FLOAT":          q.FLOAT,
				"GE":             q.GE,
				"IDENT":          q.IDENT,
				"ILLEGAL":        q.ILLEGAL,
				"IMAG":           q.IMAG,
				"INC":            q.INC,
				"INT":            q.INT,
				"LAND":           q.LAND,
				"LE":             q.LE,
				"LOR":            q.LOR,
				"MUL_ASSIGN":     q.MUL_ASSIGN,
				"NE":             q.NE,
				"NoPos":          q.NoPos,
				"OR_ASSIGN":      q.OR_ASSIGN,
				"QUO_ASSIGN":     q.QUO_ASSIGN,
				"RAT":            q.RAT,
				"REM_ASSIGN":     q.REM_ASSIGN,
				"SHL":            q.SHL,
				"SHL_ASSIGN":     q.SHL_ASSIGN,
				"SHR":            q.SHR,
				"SHR_ASSIGN":     q.SHR_ASSIGN,
				"SRARROW":        q.SRARROW,
				"STRING":         q.STRING,
				"SUB_ASSIGN":     q.SUB_ASSIGN,
				"UNIT":           q.UNIT,
				"XOR_ASSIGN":     q.XOR_ASSIGN,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"ADD":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.ADD))},
				"AND":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.AND))},
				"ASSIGN":    {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.ASSIGN))},
				"AT":        {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.AT))},
				"Break":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Break))},
				"COLON":     {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.COLON))},
				"COMMA":     {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.COMMA))},
				"ENV":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.ENV))},
				"GT":        {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.GT))},
				"LBRACE":    {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.LBRACE))},
				"LBRACK":    {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.LBRACK))},
				"LPAREN":    {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.LPAREN))},
				"LT":        {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.LT))},
				"MUL":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.MUL))},
				"NOT":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.NOT))},
				"OR":        {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.OR))},
				"PERIOD":    {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.PERIOD))},
				"QUESTION":  {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.QUESTION))},
				"QUO":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.QUO))},
				"RBRACE":    {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.RBRACE))},
				"RBRACK":    {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.RBRACK))},
				"REM":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.REM))},
				"RPAREN":    {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.RPAREN))},
				"SEMICOLON": {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.SEMICOLON))},
				"SUB":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.SUB))},
				"TILDE":     {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TILDE))},
				"XOR":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.XOR))},
			},
		}
	})
}
