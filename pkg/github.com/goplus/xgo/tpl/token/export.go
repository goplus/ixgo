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
			"ADD_ASSIGN":     {Typ: reflect.TypeOf(q.ADD_ASSIGN), Value: constant.MakeInt64(int64(q.ADD_ASSIGN))},
			"AND_ASSIGN":     {Typ: reflect.TypeOf(q.AND_ASSIGN), Value: constant.MakeInt64(int64(q.AND_ASSIGN))},
			"AND_NOT":        {Typ: reflect.TypeOf(q.AND_NOT), Value: constant.MakeInt64(int64(q.AND_NOT))},
			"AND_NOT_ASSIGN": {Typ: reflect.TypeOf(q.AND_NOT_ASSIGN), Value: constant.MakeInt64(int64(q.AND_NOT_ASSIGN))},
			"ARROW":          {Typ: reflect.TypeOf(q.ARROW), Value: constant.MakeInt64(int64(q.ARROW))},
			"BIDIARROW":      {Typ: reflect.TypeOf(q.BIDIARROW), Value: constant.MakeInt64(int64(q.BIDIARROW))},
			"CHAR":           {Typ: reflect.TypeOf(q.CHAR), Value: constant.MakeInt64(int64(q.CHAR))},
			"COMMENT":        {Typ: reflect.TypeOf(q.COMMENT), Value: constant.MakeInt64(int64(q.COMMENT))},
			"DEC":            {Typ: reflect.TypeOf(q.DEC), Value: constant.MakeInt64(int64(q.DEC))},
			"DEFINE":         {Typ: reflect.TypeOf(q.DEFINE), Value: constant.MakeInt64(int64(q.DEFINE))},
			"DRARROW":        {Typ: reflect.TypeOf(q.DRARROW), Value: constant.MakeInt64(int64(q.DRARROW))},
			"ELLIPSIS":       {Typ: reflect.TypeOf(q.ELLIPSIS), Value: constant.MakeInt64(int64(q.ELLIPSIS))},
			"EOF":            {Typ: reflect.TypeOf(q.EOF), Value: constant.MakeInt64(int64(q.EOF))},
			"EQ":             {Typ: reflect.TypeOf(q.EQ), Value: constant.MakeInt64(int64(q.EQ))},
			"FLOAT":          {Typ: reflect.TypeOf(q.FLOAT), Value: constant.MakeInt64(int64(q.FLOAT))},
			"GE":             {Typ: reflect.TypeOf(q.GE), Value: constant.MakeInt64(int64(q.GE))},
			"IDENT":          {Typ: reflect.TypeOf(q.IDENT), Value: constant.MakeInt64(int64(q.IDENT))},
			"ILLEGAL":        {Typ: reflect.TypeOf(q.ILLEGAL), Value: constant.MakeInt64(int64(q.ILLEGAL))},
			"IMAG":           {Typ: reflect.TypeOf(q.IMAG), Value: constant.MakeInt64(int64(q.IMAG))},
			"INC":            {Typ: reflect.TypeOf(q.INC), Value: constant.MakeInt64(int64(q.INC))},
			"INT":            {Typ: reflect.TypeOf(q.INT), Value: constant.MakeInt64(int64(q.INT))},
			"LAND":           {Typ: reflect.TypeOf(q.LAND), Value: constant.MakeInt64(int64(q.LAND))},
			"LE":             {Typ: reflect.TypeOf(q.LE), Value: constant.MakeInt64(int64(q.LE))},
			"LOR":            {Typ: reflect.TypeOf(q.LOR), Value: constant.MakeInt64(int64(q.LOR))},
			"MUL_ASSIGN":     {Typ: reflect.TypeOf(q.MUL_ASSIGN), Value: constant.MakeInt64(int64(q.MUL_ASSIGN))},
			"NE":             {Typ: reflect.TypeOf(q.NE), Value: constant.MakeInt64(int64(q.NE))},
			"NoPos":          {Typ: reflect.TypeOf(q.NoPos), Value: constant.MakeInt64(int64(q.NoPos))},
			"OR_ASSIGN":      {Typ: reflect.TypeOf(q.OR_ASSIGN), Value: constant.MakeInt64(int64(q.OR_ASSIGN))},
			"QUO_ASSIGN":     {Typ: reflect.TypeOf(q.QUO_ASSIGN), Value: constant.MakeInt64(int64(q.QUO_ASSIGN))},
			"RAT":            {Typ: reflect.TypeOf(q.RAT), Value: constant.MakeInt64(int64(q.RAT))},
			"REM_ASSIGN":     {Typ: reflect.TypeOf(q.REM_ASSIGN), Value: constant.MakeInt64(int64(q.REM_ASSIGN))},
			"SHL":            {Typ: reflect.TypeOf(q.SHL), Value: constant.MakeInt64(int64(q.SHL))},
			"SHL_ASSIGN":     {Typ: reflect.TypeOf(q.SHL_ASSIGN), Value: constant.MakeInt64(int64(q.SHL_ASSIGN))},
			"SHR":            {Typ: reflect.TypeOf(q.SHR), Value: constant.MakeInt64(int64(q.SHR))},
			"SHR_ASSIGN":     {Typ: reflect.TypeOf(q.SHR_ASSIGN), Value: constant.MakeInt64(int64(q.SHR_ASSIGN))},
			"SRARROW":        {Typ: reflect.TypeOf(q.SRARROW), Value: constant.MakeInt64(int64(q.SRARROW))},
			"STRING":         {Typ: reflect.TypeOf(q.STRING), Value: constant.MakeInt64(int64(q.STRING))},
			"SUB_ASSIGN":     {Typ: reflect.TypeOf(q.SUB_ASSIGN), Value: constant.MakeInt64(int64(q.SUB_ASSIGN))},
			"UNIT":           {Typ: reflect.TypeOf(q.UNIT), Value: constant.MakeInt64(int64(q.UNIT))},
			"XOR_ASSIGN":     {Typ: reflect.TypeOf(q.XOR_ASSIGN), Value: constant.MakeInt64(int64(q.XOR_ASSIGN))},
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
	})
}
