// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package parse

import (
	q "text/template/parse"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("text/template/parse", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "parse",
			Path: "text/template/parse",
			Deps: map[string]string{
				"bytes":        "bytes",
				"fmt":          "fmt",
				"runtime":      "runtime",
				"strconv":      "strconv",
				"strings":      "strings",
				"unicode":      "unicode",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{
				"Node": reflect.TypeOf((*q.Node)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"ActionNode":     reflect.TypeOf((*q.ActionNode)(nil)).Elem(),
				"BoolNode":       reflect.TypeOf((*q.BoolNode)(nil)).Elem(),
				"BranchNode":     reflect.TypeOf((*q.BranchNode)(nil)).Elem(),
				"BreakNode":      reflect.TypeOf((*q.BreakNode)(nil)).Elem(),
				"ChainNode":      reflect.TypeOf((*q.ChainNode)(nil)).Elem(),
				"CommandNode":    reflect.TypeOf((*q.CommandNode)(nil)).Elem(),
				"CommentNode":    reflect.TypeOf((*q.CommentNode)(nil)).Elem(),
				"ContinueNode":   reflect.TypeOf((*q.ContinueNode)(nil)).Elem(),
				"DotNode":        reflect.TypeOf((*q.DotNode)(nil)).Elem(),
				"FieldNode":      reflect.TypeOf((*q.FieldNode)(nil)).Elem(),
				"IdentifierNode": reflect.TypeOf((*q.IdentifierNode)(nil)).Elem(),
				"IfNode":         reflect.TypeOf((*q.IfNode)(nil)).Elem(),
				"ListNode":       reflect.TypeOf((*q.ListNode)(nil)).Elem(),
				"Mode":           reflect.TypeOf((*q.Mode)(nil)).Elem(),
				"NilNode":        reflect.TypeOf((*q.NilNode)(nil)).Elem(),
				"NodeType":       reflect.TypeOf((*q.NodeType)(nil)).Elem(),
				"NumberNode":     reflect.TypeOf((*q.NumberNode)(nil)).Elem(),
				"PipeNode":       reflect.TypeOf((*q.PipeNode)(nil)).Elem(),
				"Pos":            reflect.TypeOf((*q.Pos)(nil)).Elem(),
				"RangeNode":      reflect.TypeOf((*q.RangeNode)(nil)).Elem(),
				"StringNode":     reflect.TypeOf((*q.StringNode)(nil)).Elem(),
				"TemplateNode":   reflect.TypeOf((*q.TemplateNode)(nil)).Elem(),
				"TextNode":       reflect.TypeOf((*q.TextNode)(nil)).Elem(),
				"Tree":           reflect.TypeOf((*q.Tree)(nil)).Elem(),
				"VariableNode":   reflect.TypeOf((*q.VariableNode)(nil)).Elem(),
				"WithNode":       reflect.TypeOf((*q.WithNode)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"IsEmptyTree":   reflect.ValueOf(q.IsEmptyTree),
				"New":           reflect.ValueOf(q.New),
				"NewIdentifier": reflect.ValueOf(q.NewIdentifier),
				"Parse":         reflect.ValueOf(q.Parse),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"NodeAction":     {Typ: reflect.TypeOf(q.NodeAction), Value: constant.MakeInt64(int64(q.NodeAction))},
				"NodeBool":       {Typ: reflect.TypeOf(q.NodeBool), Value: constant.MakeInt64(int64(q.NodeBool))},
				"NodeBreak":      {Typ: reflect.TypeOf(q.NodeBreak), Value: constant.MakeInt64(int64(q.NodeBreak))},
				"NodeChain":      {Typ: reflect.TypeOf(q.NodeChain), Value: constant.MakeInt64(int64(q.NodeChain))},
				"NodeCommand":    {Typ: reflect.TypeOf(q.NodeCommand), Value: constant.MakeInt64(int64(q.NodeCommand))},
				"NodeComment":    {Typ: reflect.TypeOf(q.NodeComment), Value: constant.MakeInt64(int64(q.NodeComment))},
				"NodeContinue":   {Typ: reflect.TypeOf(q.NodeContinue), Value: constant.MakeInt64(int64(q.NodeContinue))},
				"NodeDot":        {Typ: reflect.TypeOf(q.NodeDot), Value: constant.MakeInt64(int64(q.NodeDot))},
				"NodeField":      {Typ: reflect.TypeOf(q.NodeField), Value: constant.MakeInt64(int64(q.NodeField))},
				"NodeIdentifier": {Typ: reflect.TypeOf(q.NodeIdentifier), Value: constant.MakeInt64(int64(q.NodeIdentifier))},
				"NodeIf":         {Typ: reflect.TypeOf(q.NodeIf), Value: constant.MakeInt64(int64(q.NodeIf))},
				"NodeList":       {Typ: reflect.TypeOf(q.NodeList), Value: constant.MakeInt64(int64(q.NodeList))},
				"NodeNil":        {Typ: reflect.TypeOf(q.NodeNil), Value: constant.MakeInt64(int64(q.NodeNil))},
				"NodeNumber":     {Typ: reflect.TypeOf(q.NodeNumber), Value: constant.MakeInt64(int64(q.NodeNumber))},
				"NodePipe":       {Typ: reflect.TypeOf(q.NodePipe), Value: constant.MakeInt64(int64(q.NodePipe))},
				"NodeRange":      {Typ: reflect.TypeOf(q.NodeRange), Value: constant.MakeInt64(int64(q.NodeRange))},
				"NodeString":     {Typ: reflect.TypeOf(q.NodeString), Value: constant.MakeInt64(int64(q.NodeString))},
				"NodeTemplate":   {Typ: reflect.TypeOf(q.NodeTemplate), Value: constant.MakeInt64(int64(q.NodeTemplate))},
				"NodeText":       {Typ: reflect.TypeOf(q.NodeText), Value: constant.MakeInt64(int64(q.NodeText))},
				"NodeVariable":   {Typ: reflect.TypeOf(q.NodeVariable), Value: constant.MakeInt64(int64(q.NodeVariable))},
				"NodeWith":       {Typ: reflect.TypeOf(q.NodeWith), Value: constant.MakeInt64(int64(q.NodeWith))},
				"ParseComments":  {Typ: reflect.TypeOf(q.ParseComments), Value: constant.MakeInt64(int64(q.ParseComments))},
				"SkipFuncCheck":  {Typ: reflect.TypeOf(q.SkipFuncCheck), Value: constant.MakeInt64(int64(q.SkipFuncCheck))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
