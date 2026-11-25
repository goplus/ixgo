//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

// Generate description-built operations from the reference lowering bodies.
// Run with: go run instruction_exec_gen.go
func main() {
	fset := token.NewFileSet()
	var out bytes.Buffer
	out.WriteString("// makeBinaryInstruction builds binary operations from runtime descriptions.\nfunc makeBinaryInstruction(d typedInstruction) (func(*frame), bool) {\n aux := d.Aux.(binaryOpAux)\n ir, ix, iy := d.Result, d.Args[0], d.Args[1]\n kx, ky := kindInvalid, kindInvalid\n if aux.Constant[0] { kx = kindConst }; if aux.Constant[1] { ky = kindConst }\n vx, vy := aux.Values[0], aux.Values[1]\n switch aux.Op {\n")
	ops := map[string]string{"ADD": "ADD", "SUB": "SUB", "MUL": "MUL", "QUO": "QUO", "REM": "REM", "AND": "AND", "OR": "OR", "XOR": "XOR", "ANDNOT": "AND_NOT", "LSS": "LSS", "LEQ": "LEQ", "GTR": "GTR", "GEQ": "GEQ", "EQL": "EQL", "NEQ": "NEQ", "SHL": "SHL", "SHR": "SHR"}
	for _, name := range []string{"binop.go", "binop_x.go", "binop_shift.go"} {
		file, err := parser.ParseFile(fset, name, nil, 0)
		check(err)
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || !strings.HasPrefix(fn.Name.Name, "makeBinOp") {
				continue
			}
			op := ops[strings.TrimPrefix(fn.Name.Name, "makeBinOp")]
			if op == "" {
				panic(fn.Name.Name)
			}
			fmt.Fprintf(&out, "case token.%s:\n", op)
			// The first three statements allocate destination and operand registers.
			body := fn.Body.List[3:]
			for _, stmt := range body {
				ast.Inspect(stmt, func(n ast.Node) bool {
					if _, ok := n.(*ast.FuncLit); ok {
						return false
					}
					if assign, ok := n.(*ast.AssignStmt); ok {
						for j, rhs := range assign.Rhs {
							call, ok := rhs.(*ast.CallExpr)
							if !ok {
								continue
							}
							sel, ok := call.Fun.(*ast.SelectorExpr)
							if !ok || sel.Sel.Name != "preToType" {
								continue
							}
							arg := call.Args[0].(*ast.CallExpr).Fun.(*ast.SelectorExpr).X
							field := "Type"
							if operand, ok := arg.(*ast.SelectorExpr); ok && operand.Sel.Name == "Y" {
								field = "RightType"
							}
							assign.Rhs[j] = &ast.SelectorExpr{X: ast.NewIdent("aux"), Sel: ast.NewIdent(field)}
						}
					}
					if ret, ok := n.(*ast.ReturnStmt); ok {
						ret.Results = append(ret.Results, ast.NewIdent("true"))
					}
					return true
				})
				check(format.Node(&out, fset, stmt))
				out.WriteByte('\n')
			}
		}
	}
	out.WriteString("}\nreturn nil,false\n}\n")
	out.WriteString("// makeUnaryInstruction builds unary operations from runtime descriptions.\nfunc makeUnaryInstruction(d typedInstruction) (func(*frame), bool) {\n aux := d.Aux.(unaryOpAux)\n ir, ix := d.Result, d.Args[0]\n kx := kindInvalid; if aux.Constant {kx=kindConst}; vx:=aux.Value\n switch aux.Op {\n")
	for _, name := range []string{"unop.go", "unop_x.go"} {
		file, err := parser.ParseFile(fset, name, nil, 0)
		check(err)
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}
			op := strings.TrimPrefix(fn.Name.Name, "makeUnOp")
			if op != "SUB" && op != "XOR" && op != "NOT" {
				continue
			}
			fmt.Fprintf(&out, "case token.%s:\n", op)
			for _, stmt := range fn.Body.List[2:] {
				ast.Inspect(stmt, func(n ast.Node) bool {
					if _, ok := n.(*ast.FuncLit); ok {
						return false
					}
					if assign, ok := n.(*ast.AssignStmt); ok {
						for j, rhs := range assign.Rhs {
							if call, ok := rhs.(*ast.CallExpr); ok {
								if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "preToType" {
									assign.Rhs[j] = &ast.SelectorExpr{X: ast.NewIdent("aux"), Sel: ast.NewIdent("Type")}
								}
							}
						}
					}
					if ret, ok := n.(*ast.ReturnStmt); ok {
						ret.Results = append(ret.Results, ast.NewIdent("true"))
					}
					return true
				})
				check(format.Node(&out, fset, stmt))
				out.WriteByte('\n')
			}
		}
	}
	out.WriteString(`
 case token.MUL:
  if aux.Global {
   x:=reflect.ValueOf(aux.Value)
   return func(fr *frame) { elem:=x.Elem();if !elem.IsValid(){panic(fr.runtimeError(fr.pfn.InstrForPC(int(fr.ipc)-1), "invalid memory address or nil pointer dereference"))};fr.setReg(ir,elem.Interface()) },true
  }
  return func(fr *frame){elem:=reflect.ValueOf(fr.reg(ix)).Elem();if !elem.IsValid(){panic(fr.runtimeError(fr.pfn.InstrForPC(int(fr.ipc)-1), "invalid memory address or nil pointer dereference"))};fr.setReg(ir,elem.Interface())},true
 case token.ARROW:
  typ:=aux.Type.Elem()
  if aux.Global {
   x:=reflect.ValueOf(aux.Value)
   if aux.CommaOk {return func(fr *frame){if fr.pfn.Interp.gcEnabled{fr.gcFrames()};v,ok:=x.Recv();if !ok{v=reflect.New(typ).Elem()};fr.setReg(ir,tuple{v.Interface(),ok})},true}
   return func(fr *frame){if fr.pfn.Interp.gcEnabled{fr.gcFrames()};v,ok:=x.Recv();if !ok{v=reflect.New(typ).Elem()};fr.setReg(ir,v.Interface())},true
  }
  if aux.CommaOk {return func(fr *frame){x:=reflect.ValueOf(fr.reg(ix));if fr.pfn.Interp.gcEnabled{fr.gcFrames()};v,ok:=x.Recv();if !ok{v=reflect.New(typ).Elem()};fr.setReg(ir,tuple{v.Interface(),ok})},true}
  return func(fr *frame){x:=reflect.ValueOf(fr.reg(ix));if fr.pfn.Interp.gcEnabled{fr.gcFrames()};v,ok:=x.Recv();if !ok{v=reflect.New(typ).Elem()};fr.setReg(ir,v.Interface())},true
 }
 return nil,false
 }
 `)

	data, err := os.ReadFile("instruction_exec.go")
	check(err)
	s := string(data)
	a := strings.Index(s, "// makeBinaryInstruction builds")
	if a < 0 {
		panic("missing binary builder")
	}
	s = s[:a] + out.String()
	if !strings.Contains(s, "github.com/visualfc/xtype") {
		s = strings.Replace(s, "\"reflect\"", "\"reflect\"\n\"github.com/visualfc/xtype\"", 1)
	}
	data, err = format.Source([]byte(s))
	check(err)
	check(os.WriteFile("instruction_exec.go", data, 0644))
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
