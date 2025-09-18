### build ixgo for Go version >= 1.23

- go.mod use `golang.org/x/tools v0.30.0` for support range func

- build use `-ldflags="-checklinkname=0"` for support linkname

### Install igxo for Go version >= 1.23
`go install -ldflags="-checklinkname=0" github.com/goplus/ixgo/ixgo@v0.3.1`


### Local development environment

go.work
```
go 1.23

use (
	.
	./..
)
```


