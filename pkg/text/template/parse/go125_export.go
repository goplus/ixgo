// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package parse

import (
	q "text/template/parse"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("text/template/parse", func() *ixgo.Package {
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"IsEmptyTree":   q.IsEmptyTree,
				"New":           q.New,
				"NewIdentifier": q.NewIdentifier,
				"Parse":         q.Parse,
			},
			TypedConsts: map[string]interface{}{
				"NodeAction":     q.NodeAction,
				"NodeBool":       q.NodeBool,
				"NodeBreak":      q.NodeBreak,
				"NodeChain":      q.NodeChain,
				"NodeCommand":    q.NodeCommand,
				"NodeComment":    q.NodeComment,
				"NodeContinue":   q.NodeContinue,
				"NodeDot":        q.NodeDot,
				"NodeField":      q.NodeField,
				"NodeIdentifier": q.NodeIdentifier,
				"NodeIf":         q.NodeIf,
				"NodeList":       q.NodeList,
				"NodeNil":        q.NodeNil,
				"NodeNumber":     q.NodeNumber,
				"NodePipe":       q.NodePipe,
				"NodeRange":      q.NodeRange,
				"NodeString":     q.NodeString,
				"NodeTemplate":   q.NodeTemplate,
				"NodeText":       q.NodeText,
				"NodeVariable":   q.NodeVariable,
				"NodeWith":       q.NodeWith,
				"ParseComments":  q.ParseComments,
				"SkipFuncCheck":  q.SkipFuncCheck,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
