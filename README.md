# iXGo The Go/XGo Interpreter

[![Build Status](https://github.com/goplus/ixgo/workflows/Go/badge.svg)](https://github.com/goplus/ixgo/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/goplus/ixgo.svg)](https://pkg.go.dev/github.com/goplus/ixgo)

A fast and fully compatible Go language interpreter.

### Go Version

- Go1.24 ~ Go1.26
- macOS Linux Windows  WebAssembly.

### ABI

- ABI0 stack-based ABI
- ABIInternal [register-based Go calling convention proposal](https://golang.org/design/40724-register-calling)

    - Go1.24 ~ Go1.26: amd64 arm64 ppc64/ppc64le riscv64 loong64


### Generics

- support typeparams
- support alias typeparams

### runtime.GC

`ixgo.Mode ExperimentalSupportGC`

experimental support runtime.GC and runtime.SetFinalizer

### install ixgo command line

```shell
go install github.com/goplus/ixgo/cmd/ixgo@latest
```

unsupport
- runtime.AddCleanup
- iter.Pull/iter.Pull2

### install ixgo command line fully mode

```shell
go install go install -tags linknamefix -ldflags="-checklinkname=0" github.com/goplus/ixgo/cmd/ixgo@latest
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
ixgo verson      # print version
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

### ixgo test unsupport features

- test -fuzz
- test -cover

### ixgo demo

#### The XGo Playground (WebAssembly)

- <https://play.xgo.dev/>
- <https://github.com/goplusjs/play>

#### The XGo REPL Playground (WebAssembly)

- <https://repl.xgo.dev/>
- <https://github.com/goplusjs/repl>

#### ispx

[ispx](https://github.com/goplus/spx/tree/dev/cmd/igox)

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
fields := [
	"engineering",
	"STEM education", 
	"and data science",
]

echo "The XGo language for", fields.join(", ")
`

func main() {
	_, err := ixgo.RunFile("main.xgo", source, nil, 0)
	if err != nil {
		panic(err)
	}
}
```

#### ixgo more demo

<https://github.com/visualfc/ixgo_demo>
