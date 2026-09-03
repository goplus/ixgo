# Standard-library direct calls

This directory contains optional direct-call adapters for the standard
library packages exported by `ixgo/pkg`. Import `ixgo/pkg` separately when
the interpreter also needs the package exports, then import the adapters:

```go
import _ "github.com/goplus/ixgo/directcall"
```

Individual adapters can be enabled instead, for example:

```go
import _ "github.com/goplus/ixgo/directcall/strings"
```

The generated files are kept for Go 1.25, 1.26, and 1.27. Regenerate one
version from the repository root with:

```sh
go run ./directcall/gen.go 1.25
go run ./directcall/gen.go 1.26
go run ./directcall/gen.go 1.27
```

The generator expects the matching toolchains at `~/golang/go1.25`,
`~/golang/go1.26`, and `~/golang/go1.27`.

## Excluded packages

Direct-call adapters are intentionally not generated for `log` and
`log/slog`. These packages use `runtime.Caller` internally to determine the
source location of a log call. A direct-call adapter changes that call stack,
so `Llongfile` and `Lshortfile` can report the adapter file instead of the
user's source file. They therefore remain on the source implementation.
