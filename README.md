# iXGo The Go/XGo Interpreter

[![Build Status](https://github.com/goplus/ixgo/workflows/Go/badge.svg)](https://github.com/goplus/ixgo/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/goplus/ixgo.svg)](https://pkg.go.dev/github.com/goplus/ixgo)

A fast Go interpreter that supports all Go language features, including generics and generic methods.

### Go Version

- Go 1.25 ~ 1.27

### Compilers

- Go (gc)
- [LLGo](https://github.com/xgo-dev/llgo)

### Platforms

- macOS
- Linux
- Windows
- WebAssembly

### Go ABI

- ABI0 stack-based ABI
- ABIInternal [register-based Go calling convention proposal](https://golang.org/design/40724-register-calling)

    - amd64 arm64 ppc64/ppc64le riscv64 loong64


### Generics

- Supports type parameters (generics)
- Supports generic type aliases
- Supports generic methods (Go 1.27)

### install ixgo command line

```shell
go install github.com/goplus/ixgo/cmd/ixgo@latest
```

### install ixgo export command line
```shell
go install github.com/goplus/ixgo/cmd/qexp@latest
```

### ixgo command

```
ixgo             # ixgo repl mode
ixgo run         # run a Go/XGo package
ixgo build       # compile a Go/XGo package
ixgo test        # test a package
ixgo version     # print version
ixgo export      # export Go package to ixgo builtin package
```

### ixgo run mode
```
Usage: ixgo run [build flags] [package] [arguments...]
  -exp-gc
    	experimental support runtime.GC
  -mod value
    	module download mode to use: readonly, vendor, or mod.
  -ssa
    	print SSA instruction code
  -ssa-trace
    	trace SSA interpreter code
  -tags value
    	a comma-separated list of build tags to consider satisfied during the build
  -v	print the names of packages as they are compiled.
  -x	print the commands.
```

### ixgo repl mode

```shell
ixgo                       # run repl mode, support Go/XGo
ixgo repl                  # run repl mode, support Go/XGo
ixgo repl -go              # run repl mode, disable XGo syntax
```

### ixgo test unsupported features

- test -fuzz
- test -cover

### ifacefuncval (optional)

Alternative to reflectx icall stubs: methods share `makeFuncStub`.
reflectx stores an untagged `*makeFuncImpl` in `addReflectOff` and
sets `tflagIfaceFuncval` (`1<<6`). A patched runtime tags `itab.Fun`
(`makeFuncImpl*|1`) so interface calls unwrap.

Needs a **patched Go 1.25.x, 1.26.x, or 1.27.x** and `-tags goplus.ifacefuncval`.
Supported `GOARCH`: wasm, arm64, amd64, 386.

```
go install github.com/goplus/reflectx/cmd/iface_patch@latest
```
```shell
iface_patch /path/to/go   # 1.25.x, 1.26.x, or 1.27.x
cd /path/to/go/src && ./make.bash
export GOROOT=/path/to/go
export PATH="$GOROOT/bin:$PATH"
go test -tags goplus.ifacefuncval .
# wasip1: GOOS=wasip1 GOARCH=wasm go test -exec wasmtime -tags goplus.ifacefuncval .
```

Without the tag, a patched compiler matches official gc. An unpatched
compiler still accepts the tag; interface method calls then trap.

You can `export GOFLAGS='-tags=goplus.ifacefuncval'`, or prefix a command:
`GOFLAGS='-tags=goplus.ifacefuncval' go build` (also `go run` / `go test`).

See [goplus/reflectx cmd/iface_patch](https://github.com/goplus/reflectx/tree/main/cmd/iface_patch).

### ixgo demo

#### The XGo Playground (WebAssembly)

- <https://play.xgo.dev/>
- <https://github.com/goplusjs/play>

#### The XGo REPL Playground (WebAssembly)

- <https://repl.xgo.dev/>
- <https://github.com/goplusjs/repl>


#### run simple Go source demo

```go
package main

import (
	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/fmt"
)

var source = `
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
}
`

func main() {
	_, err := ixgo.RunFile("main.go", source, nil, 0)
	if err != nil {
		panic(err)
	}
}
```

#### run simple XGo source demo

```go
package main

import (
	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/xgobuild"
)

var source = `
languages := [
	"C",
	"Go",
	"Python",
	"JavaScript",
]

echo "XGo :=", languages.join(" * "), "+ Scratch"
`

func main() {
	_, err := ixgo.RunFile("main.xgo", source, nil, 0)
	if err != nil {
		panic(err)
	}
}
```

#### build linkname mode

```shell
go install -tags linknamefix -ldflags="-checklinkname=0" github.com/goplus/ixgo/cmd/ixgo@latest
```

use runtime linkname for faster performance `iter.Pull/iter.Pull2`

#### ixgo more demo

<https://github.com/visualfc/ixgo_demo>
