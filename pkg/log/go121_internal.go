//go:build go1.21
// +build go1.21

package log

import (
	_ "github.com/goplus/ixgo/pkg/fmt"
	_ "github.com/goplus/ixgo/pkg/io"
	_ "github.com/goplus/ixgo/pkg/log/internal"
	_ "github.com/goplus/ixgo/pkg/os"
	_ "github.com/goplus/ixgo/pkg/runtime"
	_ "github.com/goplus/ixgo/pkg/sync"
	_ "github.com/goplus/ixgo/pkg/sync/atomic"
	_ "github.com/goplus/ixgo/pkg/time"
)
