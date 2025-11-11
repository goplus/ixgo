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

package export

import (
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

func writeFile(dir string, file string, data []byte) error {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return fmt.Errorf("make dir %v error: %v", dir, err)
	}
	filename := filepath.Join(dir, file)
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		return fmt.Errorf("write file %v error: %v", filename, err)
	}
	return nil
}

func joinList(list []string) string {
	if len(list) == 0 {
		return ""
	}
	sort.Strings(list)
	return "\n\t" + strings.Join(list, ",\n\t") + ",\n"
}

const EmptyPackage = "empty package"

var (
	errEmptyPackage = errors.New(EmptyPackage)
)

func exportPkg(pkg *Package, sname string, id string, tagList []string) ([]byte, error) {
	var imports []string
	if pkg.usedPkg {
		imports = append(imports, fmt.Sprintf("q %q", pkg.Path))
		imports = append(imports, "")
	} else if !flagExportCode {
		imports = append(imports, fmt.Sprintf("_ %q", pkg.Path))
	}
	if !pkg.IsEmpty() {
		imports = append(imports, `"reflect"`)
	}
	if len(pkg.UntypedConsts) > 0 || len(pkg.TypedConsts) > 0 {
		imports = append(imports, `"go/constant"`)
		var hasToken bool
		for _, c := range pkg.UntypedConsts {
			if strings.Index(c, "token.") >= 0 {
				hasToken = true
				break
			}
		}
		if !hasToken {
			for _, c := range pkg.TypedConsts {
				if strings.Index(c, "token.") >= 0 {
					hasToken = true
					break
				}
			}
		}
		if hasToken {
			imports = append(imports, `"go/token"`)
		}
	}
	tmpl := template_pkg
	if pkg.IsEmpty() {
		tmpl = template_empty_pkg
	}
	if len(pkg.Source) > 0 {
		tmpl = template_link_pkg
		if pkg.IsEmpty() {
			tmpl = template_emtpy_link_pkg
		}
		if len(pkg.Links) > 0 {
			imports = append(imports, `_ "unsafe"`)
		}
	}
	if flagCustomPkg != "" {
		pkg.Path = path.Clean(flagCustomPkg)
		pkg.Name = path.Base(pkg.Path)
	}
	r := strings.NewReplacer("$PKGNAME", pkg.Name,
		"$IMPORTS", strings.Join(imports, "\n"),
		"$PKGPATH", pkg.Path,
		"$DEPS", joinList(pkg.Deps),
		"$NAMEDTYPES", joinList(pkg.NamedTypes),
		"$INTERFACES", joinList(pkg.Interfaces),
		"$ALIASTYPES", joinList(pkg.AliasTypes),
		"$VARS", joinList(pkg.Vars),
		"$FUNCS", joinList(pkg.Funcs),
		"$TYPEDCONSTS", joinList(pkg.TypedConsts),
		"$UNTYPEDCONSTS", joinList(pkg.UntypedConsts),
		"$TAGS", strings.Join(tagList, "\n"),
		"$SOURCE", pkg.Source,
		"$LINKS", strings.Join(pkg.Links, "\n"),
		"$ID", id)
	src := r.Replace(tmpl)
	data, err := format.Source([]byte(src))
	if err != nil {
		return nil, fmt.Errorf("format pkg %v error: %v", src, err)
	}
	return data, nil
}

var template_pkg = `// export by github.com/goplus/ixgo/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
		Interfaces: map[string]reflect.Type{$INTERFACES},
		NamedTypes: map[string]reflect.Type{$NAMEDTYPES},
		AliasTypes: map[string]reflect.Type{$ALIASTYPES},
		Vars: map[string]reflect.Value{$VARS},
		Funcs: map[string]reflect.Value{$FUNCS},
		TypedConsts: map[string]ixgo.TypedConst{$TYPEDCONSTS},
		UntypedConsts: map[string]ixgo.UntypedConst{$UNTYPEDCONSTS},
	})
}
`

var template_empty_pkg = `// export by github.com/goplus/ixgo/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
	})
}
`

var template_link_pkg = `// export by github.com/goplus/ixgo/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
		Interfaces: map[string]reflect.Type{$INTERFACES},
		NamedTypes: map[string]reflect.Type{$NAMEDTYPES},
		AliasTypes: map[string]reflect.Type{$ALIASTYPES},
		Vars: map[string]reflect.Value{$VARS},
		Funcs: map[string]reflect.Value{$FUNCS},
		TypedConsts: map[string]ixgo.TypedConst{$TYPEDCONSTS},
		UntypedConsts: map[string]ixgo.UntypedConst{$UNTYPEDCONSTS},
		Source: source,
	})
}
$LINKS
var source = $SOURCE
`

var template_emtpy_link_pkg = `// export by github.com/goplus/ixgo/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
		Source: source,
	})
}
$LINKS
var source = $SOURCE
`
