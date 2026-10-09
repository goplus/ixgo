# Direct funcs

Register a signature by embedding `DirectFuncVal` and binding a method. This example
makes interpreted `func(int) int` values ordinary Go functions:

```go
package myint

import "github.com/goplus/ixgo"

type fn struct{ ixgo.DirectFuncVal }

func (c fn) Int(n int) int {
	return c.Call(n).(int)
}

func init() {
	ixgo.RegisterDirectFuncFor(func(c ixgo.DirectFuncVal) func(int) int {
		return fn{c}.Int
	})
}
```

```go
src := `package main
func main() {
	add := func(x int) int { return x + 1 }
	println(add(41))
}
`
ixgo.RunFile("main.go", src, nil, 0)
```

`func()` is built into ixgo. Import this package for common `func() T` results.

```go
import _ "github.com/goplus/ixgo/directfunc"
```

Registered signatures:

- `func() bool`
- `func() int` / `int8` / `int16` / `int32` / `int64`
- `func() uint` / `uint8` / `uint16` / `uint32` / `uint64` / `uintptr`
- `func() float32` / `float64`
- `func() string`
- `func() error`

Named types (for example `type Action func()`) stay on `reflect.MakeFunc`
unless the host registers them.

A later registration for the same signature replaces an earlier one. Passing a
nil maker to `RegisterDirectFunc`, or a nil bind to
`RegisterDirectFuncFor`, removes the registration. Do not unregister or
replace a signature while function values of that type are still in use.
