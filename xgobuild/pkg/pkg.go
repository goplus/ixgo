package pkg

//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/qiniu/x/gsh
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/ast
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/parser
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/printer
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/scanner
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/token
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/encoding/...
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/tpl/...
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/dql/fetcher
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/dql/fs
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/dql/golang
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/dql/html
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/dql/json
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/dql/maps
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/dql/reflects
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/dql/xgo
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/dql/xml
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/dql/yaml
//go:generate go run ../../cmd/qexp -code -outdir ../../pkg github.com/goplus/xgo/dql
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/test
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/doc

import (
	_ "github.com/goplus/ixgo/xgobuild/pkg/dql"
	_ "github.com/goplus/ixgo/xgobuild/pkg/encoding"
	_ "github.com/goplus/ixgo/xgobuild/pkg/gsh"
	_ "github.com/goplus/ixgo/xgobuild/pkg/test"
	_ "github.com/goplus/ixgo/xgobuild/pkg/tpl"
	_ "github.com/goplus/ixgo/xgobuild/pkg/xgo"
	_ "github.com/goplus/xgo/dql/html"
	_ "github.com/goplus/xgo/dql/xml"
	_ "github.com/goplus/xgo/dql/yaml"
)
