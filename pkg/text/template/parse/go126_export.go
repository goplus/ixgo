// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package parse

import (
	q "text/template/parse"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
			"NodeAction":     {reflect.TypeOf(q.NodeAction), constant.MakeInt64(int64(q.NodeAction))},
			"NodeBool":       {reflect.TypeOf(q.NodeBool), constant.MakeInt64(int64(q.NodeBool))},
			"NodeBreak":      {reflect.TypeOf(q.NodeBreak), constant.MakeInt64(int64(q.NodeBreak))},
			"NodeChain":      {reflect.TypeOf(q.NodeChain), constant.MakeInt64(int64(q.NodeChain))},
			"NodeCommand":    {reflect.TypeOf(q.NodeCommand), constant.MakeInt64(int64(q.NodeCommand))},
			"NodeComment":    {reflect.TypeOf(q.NodeComment), constant.MakeInt64(int64(q.NodeComment))},
			"NodeContinue":   {reflect.TypeOf(q.NodeContinue), constant.MakeInt64(int64(q.NodeContinue))},
			"NodeDot":        {reflect.TypeOf(q.NodeDot), constant.MakeInt64(int64(q.NodeDot))},
			"NodeField":      {reflect.TypeOf(q.NodeField), constant.MakeInt64(int64(q.NodeField))},
			"NodeIdentifier": {reflect.TypeOf(q.NodeIdentifier), constant.MakeInt64(int64(q.NodeIdentifier))},
			"NodeIf":         {reflect.TypeOf(q.NodeIf), constant.MakeInt64(int64(q.NodeIf))},
			"NodeList":       {reflect.TypeOf(q.NodeList), constant.MakeInt64(int64(q.NodeList))},
			"NodeNil":        {reflect.TypeOf(q.NodeNil), constant.MakeInt64(int64(q.NodeNil))},
			"NodeNumber":     {reflect.TypeOf(q.NodeNumber), constant.MakeInt64(int64(q.NodeNumber))},
			"NodePipe":       {reflect.TypeOf(q.NodePipe), constant.MakeInt64(int64(q.NodePipe))},
			"NodeRange":      {reflect.TypeOf(q.NodeRange), constant.MakeInt64(int64(q.NodeRange))},
			"NodeString":     {reflect.TypeOf(q.NodeString), constant.MakeInt64(int64(q.NodeString))},
			"NodeTemplate":   {reflect.TypeOf(q.NodeTemplate), constant.MakeInt64(int64(q.NodeTemplate))},
			"NodeText":       {reflect.TypeOf(q.NodeText), constant.MakeInt64(int64(q.NodeText))},
			"NodeVariable":   {reflect.TypeOf(q.NodeVariable), constant.MakeInt64(int64(q.NodeVariable))},
			"NodeWith":       {reflect.TypeOf(q.NodeWith), constant.MakeInt64(int64(q.NodeWith))},
			"ParseComments":  {reflect.TypeOf(q.ParseComments), constant.MakeInt64(int64(q.ParseComments))},
			"SkipFuncCheck":  {reflect.TypeOf(q.SkipFuncCheck), constant.MakeInt64(int64(q.SkipFuncCheck))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
