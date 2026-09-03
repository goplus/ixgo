//go:build !windows && !plan9
// +build !windows,!plan9

package directcall

import (
	_ "github.com/goplus/ixgo/directcall/log/syslog"
)
