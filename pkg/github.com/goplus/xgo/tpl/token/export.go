// export by github.com/goplus/ixgo/cmd/qexp

package token

import (
	q "github.com/goplus/xgo/tpl/token"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ForEach":    reflect.ValueOf(q.ForEach),
			"NewFileSet": reflect.ValueOf(q.NewFileSet),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"ADD_ASSIGN":     {reflect.TypeOf(q.ADD_ASSIGN), constant.MakeInt64(int64(q.ADD_ASSIGN))},
			"AND_ASSIGN":     {reflect.TypeOf(q.AND_ASSIGN), constant.MakeInt64(int64(q.AND_ASSIGN))},
			"AND_NOT":        {reflect.TypeOf(q.AND_NOT), constant.MakeInt64(int64(q.AND_NOT))},
			"AND_NOT_ASSIGN": {reflect.TypeOf(q.AND_NOT_ASSIGN), constant.MakeInt64(int64(q.AND_NOT_ASSIGN))},
			"ARROW":          {reflect.TypeOf(q.ARROW), constant.MakeInt64(int64(q.ARROW))},
			"BIDIARROW":      {reflect.TypeOf(q.BIDIARROW), constant.MakeInt64(int64(q.BIDIARROW))},
			"CHAR":           {reflect.TypeOf(q.CHAR), constant.MakeInt64(int64(q.CHAR))},
			"COMMENT":        {reflect.TypeOf(q.COMMENT), constant.MakeInt64(int64(q.COMMENT))},
			"DEC":            {reflect.TypeOf(q.DEC), constant.MakeInt64(int64(q.DEC))},
			"DEFINE":         {reflect.TypeOf(q.DEFINE), constant.MakeInt64(int64(q.DEFINE))},
			"DRARROW":        {reflect.TypeOf(q.DRARROW), constant.MakeInt64(int64(q.DRARROW))},
			"ELLIPSIS":       {reflect.TypeOf(q.ELLIPSIS), constant.MakeInt64(int64(q.ELLIPSIS))},
			"EOF":            {reflect.TypeOf(q.EOF), constant.MakeInt64(int64(q.EOF))},
			"EQ":             {reflect.TypeOf(q.EQ), constant.MakeInt64(int64(q.EQ))},
			"FLOAT":          {reflect.TypeOf(q.FLOAT), constant.MakeInt64(int64(q.FLOAT))},
			"GE":             {reflect.TypeOf(q.GE), constant.MakeInt64(int64(q.GE))},
			"IDENT":          {reflect.TypeOf(q.IDENT), constant.MakeInt64(int64(q.IDENT))},
			"ILLEGAL":        {reflect.TypeOf(q.ILLEGAL), constant.MakeInt64(int64(q.ILLEGAL))},
			"IMAG":           {reflect.TypeOf(q.IMAG), constant.MakeInt64(int64(q.IMAG))},
			"INC":            {reflect.TypeOf(q.INC), constant.MakeInt64(int64(q.INC))},
			"INT":            {reflect.TypeOf(q.INT), constant.MakeInt64(int64(q.INT))},
			"LAND":           {reflect.TypeOf(q.LAND), constant.MakeInt64(int64(q.LAND))},
			"LE":             {reflect.TypeOf(q.LE), constant.MakeInt64(int64(q.LE))},
			"LOR":            {reflect.TypeOf(q.LOR), constant.MakeInt64(int64(q.LOR))},
			"MUL_ASSIGN":     {reflect.TypeOf(q.MUL_ASSIGN), constant.MakeInt64(int64(q.MUL_ASSIGN))},
			"NE":             {reflect.TypeOf(q.NE), constant.MakeInt64(int64(q.NE))},
			"NoPos":          {reflect.TypeOf(q.NoPos), constant.MakeInt64(int64(q.NoPos))},
			"OR_ASSIGN":      {reflect.TypeOf(q.OR_ASSIGN), constant.MakeInt64(int64(q.OR_ASSIGN))},
			"QUO_ASSIGN":     {reflect.TypeOf(q.QUO_ASSIGN), constant.MakeInt64(int64(q.QUO_ASSIGN))},
			"RAT":            {reflect.TypeOf(q.RAT), constant.MakeInt64(int64(q.RAT))},
			"REM_ASSIGN":     {reflect.TypeOf(q.REM_ASSIGN), constant.MakeInt64(int64(q.REM_ASSIGN))},
			"SHL":            {reflect.TypeOf(q.SHL), constant.MakeInt64(int64(q.SHL))},
			"SHL_ASSIGN":     {reflect.TypeOf(q.SHL_ASSIGN), constant.MakeInt64(int64(q.SHL_ASSIGN))},
			"SHR":            {reflect.TypeOf(q.SHR), constant.MakeInt64(int64(q.SHR))},
			"SHR_ASSIGN":     {reflect.TypeOf(q.SHR_ASSIGN), constant.MakeInt64(int64(q.SHR_ASSIGN))},
			"SRARROW":        {reflect.TypeOf(q.SRARROW), constant.MakeInt64(int64(q.SRARROW))},
			"STRING":         {reflect.TypeOf(q.STRING), constant.MakeInt64(int64(q.STRING))},
			"SUB_ASSIGN":     {reflect.TypeOf(q.SUB_ASSIGN), constant.MakeInt64(int64(q.SUB_ASSIGN))},
			"UNIT":           {reflect.TypeOf(q.UNIT), constant.MakeInt64(int64(q.UNIT))},
			"XOR_ASSIGN":     {reflect.TypeOf(q.XOR_ASSIGN), constant.MakeInt64(int64(q.XOR_ASSIGN))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"ADD":       {"untyped rune", constant.MakeInt64(int64(q.ADD))},
			"AND":       {"untyped rune", constant.MakeInt64(int64(q.AND))},
			"ASSIGN":    {"untyped rune", constant.MakeInt64(int64(q.ASSIGN))},
			"AT":        {"untyped rune", constant.MakeInt64(int64(q.AT))},
			"Break":     {"untyped int", constant.MakeInt64(int64(q.Break))},
			"COLON":     {"untyped rune", constant.MakeInt64(int64(q.COLON))},
			"COMMA":     {"untyped rune", constant.MakeInt64(int64(q.COMMA))},
			"ENV":       {"untyped rune", constant.MakeInt64(int64(q.ENV))},
			"GT":        {"untyped rune", constant.MakeInt64(int64(q.GT))},
			"LBRACE":    {"untyped rune", constant.MakeInt64(int64(q.LBRACE))},
			"LBRACK":    {"untyped rune", constant.MakeInt64(int64(q.LBRACK))},
			"LPAREN":    {"untyped rune", constant.MakeInt64(int64(q.LPAREN))},
			"LT":        {"untyped rune", constant.MakeInt64(int64(q.LT))},
			"MUL":       {"untyped rune", constant.MakeInt64(int64(q.MUL))},
			"NOT":       {"untyped rune", constant.MakeInt64(int64(q.NOT))},
			"OR":        {"untyped rune", constant.MakeInt64(int64(q.OR))},
			"PERIOD":    {"untyped rune", constant.MakeInt64(int64(q.PERIOD))},
			"QUESTION":  {"untyped rune", constant.MakeInt64(int64(q.QUESTION))},
			"QUO":       {"untyped rune", constant.MakeInt64(int64(q.QUO))},
			"RBRACE":    {"untyped rune", constant.MakeInt64(int64(q.RBRACE))},
			"RBRACK":    {"untyped rune", constant.MakeInt64(int64(q.RBRACK))},
			"REM":       {"untyped rune", constant.MakeInt64(int64(q.REM))},
			"RPAREN":    {"untyped rune", constant.MakeInt64(int64(q.RPAREN))},
			"SEMICOLON": {"untyped rune", constant.MakeInt64(int64(q.SEMICOLON))},
			"SUB":       {"untyped rune", constant.MakeInt64(int64(q.SUB))},
			"TILDE":     {"untyped rune", constant.MakeInt64(int64(q.TILDE))},
			"XOR":       {"untyped rune", constant.MakeInt64(int64(q.XOR))},
		},
	})
}
