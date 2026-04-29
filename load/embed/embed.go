/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package embed

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"strconv"
	_ "unsafe"

	"github.com/visualfc/goembed"
	embedparser "github.com/visualfc/goembed/parser"
)

func buildIdent(name string) string {
	return fmt.Sprintf("__ixgo_embed_%x__", name)
}

var embed_head = `package %v

import (
	"embed"
	"unsafe"
)

func __ixgo_embed_buildFS__(list []struct {
	name string
	data string
	hash [16]byte
}) (f embed.FS) {
	fs := struct {
		files *[]struct {
			name string
			data string
			hash [16]byte
		}
	}{&list}
	return *(*embed.FS)(unsafe.Pointer(&fs))
}
`

// Embed check package embed data
func Embed(bp *build.Package, fset *token.FileSet, files []*ast.File, test bool, xtest bool) (*ast.File, error) {
	var pkgName string
	var err error
	var ems []*goembed.Embed
	if xtest {
		pkgName = bp.Name + "_test"
		ems, err = goembed.CheckEmbed(bp.XTestEmbedPatternPos, fset, files)
		if err != nil {
			return nil, err
		}
	} else {
		pkgName = bp.Name
		ems, err = goembed.CheckEmbed(bp.EmbedPatternPos, fset, files)
		if err != nil {
			return nil, err
		}
		if test {
			tems, err := goembed.CheckEmbed(bp.TestEmbedPatternPos, fset, files)
			if err != nil {
				return nil, err
			}
			if len(tems) > 0 {
				ems = append(ems, tems...)
			}
		}
	}
	if len(ems) == 0 {
		return nil, nil
	}
	r := goembed.NewResolve()
	for _, v := range ems {
		fs, err := r.Load(bp.Dir, fset, v)
		if err != nil {
			return nil, err
		}
		switch v.Kind {
		case goembed.EmbedMaybeAlias:
			// value = Type(data)
			// valid alias string or []byte type used by types.check
			v.Spec.Values = []ast.Expr{
				&ast.CallExpr{
					Fun: v.Spec.Type,
					Args: []ast.Expr{
						&ast.Ident{Name: buildIdent(fs[0].Name),
							NamePos: v.Spec.Names[0].NamePos},
					},
				}}
		case goembed.EmbedBytes:
			// value = []byte(data)
			v.Spec.Values = []ast.Expr{
				&ast.CallExpr{
					Fun:  v.Spec.Type,
					Args: []ast.Expr{ast.NewIdent(buildIdent(fs[0].Name))},
				}}
		case goembed.EmbedString:
			// value = data
			v.Spec.Values = []ast.Expr{ast.NewIdent(buildIdent(fs[0].Name))}
		case goembed.EmbedFiles:
			// value = __ixgo_embed_buildFS__([]struct{name string; data string; hash [16]byte}{...})
			fs = goembed.BuildFS(fs)
			elts := make([]ast.Expr, len(fs), len(fs))
			for i, f := range fs {
				if len(f.Data) == 0 {
					elts[i] = &ast.CompositeLit{
						Elts: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(f.Name)},
							&ast.BasicLit{Kind: token.STRING, Value: `""`},
							&ast.CompositeLit{
								Type: &ast.ArrayType{
									Len: &ast.BasicLit{Kind: token.INT, Value: "16"},
									Elt: ast.NewIdent("byte"),
								},
							},
						},
					}
				} else {
					var hash [16]ast.Expr
					for j, v := range f.Hash {
						hash[j] = &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(int(v))}
					}
					elts[i] = &ast.CompositeLit{
						Elts: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(f.Name)},
							ast.NewIdent(buildIdent(f.Name)),
							&ast.CompositeLit{
								Type: &ast.ArrayType{
									Len: &ast.BasicLit{Kind: token.INT, Value: "16"},
									Elt: ast.NewIdent("byte"),
								},
								Elts: hash[:],
							},
						},
					}
				}
			}
			call := &ast.CallExpr{
				Fun: ast.NewIdent("__ixgo_embed_buildFS__"),
				Args: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.ArrayType{
							Elt: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										&ast.Field{
											Names: []*ast.Ident{ast.NewIdent("name")},
											Type:  ast.NewIdent("string"),
										},
										&ast.Field{
											Names: []*ast.Ident{ast.NewIdent("data")},
											Type:  ast.NewIdent("string"),
										},
										&ast.Field{
											Names: []*ast.Ident{ast.NewIdent("hash")},
											Type: &ast.ArrayType{
												Len: &ast.BasicLit{Kind: token.INT, Value: "16"},
												Elt: ast.NewIdent("byte"),
											},
										},
									},
								},
							},
						},
						Elts: elts,
					},
				},
			}
			v.Spec.Values = []ast.Expr{call}
		}
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, embed_head, pkgName)
	buf.WriteString("\nconst (\n")
	for _, f := range r.Files() {
		if len(f.Data) == 0 {
			fmt.Fprintf(&buf, "\t%v = \"\"\n", buildIdent(f.Name))
		} else {
			fmt.Fprintf(&buf, "\t%v = \"%v\"\n", buildIdent(f.Name), goembed.BytesToHex(f.Data))
		}
	}
	buf.WriteString(")\n\n")
	return parser.ParseFile(fset, "_ixgo_embed_data.go", buf.Bytes(), parser.ParseComments)
}

func EmbedFiles(pkgName string, dir string, fset *token.FileSet, files []*ast.File) (*ast.File, error) {
	embed, err := embedparser.ParseEmbed(fset, files)
	if err != nil {
		return nil, err
	}
	if embed == nil {
		return nil, nil
	}
	bp := &build.Package{
		Name:            pkgName,
		Dir:             dir,
		EmbedPatterns:   embed.Patterns,
		EmbedPatternPos: embed.PatternPos,
	}
	return Embed(bp, fset, files, false, false)
}

type Processor struct {
}

func (Processor) LoadPkg(bp *build.Package, fset *token.FileSet, files []*ast.File, test bool, xtest bool) (*ast.File, error) {
	return Embed(bp, fset, files, test, xtest)
}

func (Processor) LoadFiles(pkgName string, dir string, fset *token.FileSet, files []*ast.File) (*ast.File, error) {
	return EmbedFiles(pkgName, dir, fset, files)
}
