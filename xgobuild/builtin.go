package xgobuild

//go:generate go run ../cmd/qexp -outdir ../pkg github.com/qiniu/x/xgo/...
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/qiniu/x/osx
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/qiniu/x/errors
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/qiniu/x/stringutil
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/qiniu/x/stringslice

import (
	_ "github.com/goplus/ixgo/pkg/github.com/qiniu/x/errors"
	_ "github.com/goplus/ixgo/pkg/github.com/qiniu/x/osx"
	_ "github.com/goplus/ixgo/pkg/github.com/qiniu/x/stringslice"
	_ "github.com/goplus/ixgo/pkg/github.com/qiniu/x/stringutil"
	_ "github.com/goplus/ixgo/pkg/github.com/qiniu/x/xgo"
	_ "github.com/goplus/ixgo/pkg/github.com/qiniu/x/xgo/ng"

	_ "github.com/goplus/ixgo/pkg/errors"
	_ "github.com/goplus/ixgo/pkg/fmt"
	_ "github.com/goplus/ixgo/pkg/io"
	_ "github.com/goplus/ixgo/pkg/math"
	_ "github.com/goplus/ixgo/pkg/math/big"
	_ "github.com/goplus/ixgo/pkg/os"
	_ "github.com/goplus/ixgo/pkg/reflect"
	_ "github.com/goplus/ixgo/pkg/strconv"
	_ "github.com/goplus/ixgo/pkg/strings"
)
