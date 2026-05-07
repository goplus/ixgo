// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package syslog

import (
	q "log/syslog"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("log/syslog", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "syslog",
			Path: "log/syslog",
			Deps: map[string]string{
				"errors":  "errors",
				"fmt":     "fmt",
				"log":     "log",
				"net":     "net",
				"os":      "os",
				"strings": "strings",
				"sync":    "sync",
				"time":    "time",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Priority": reflect.TypeOf((*q.Priority)(nil)).Elem(),
				"Writer":   reflect.TypeOf((*q.Writer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Dial":      q.Dial,
				"New":       q.New,
				"NewLogger": q.NewLogger,
			},
			TypedConsts: map[string]interface{}{
				"LOG_ALERT":    q.LOG_ALERT,
				"LOG_AUTH":     q.LOG_AUTH,
				"LOG_AUTHPRIV": q.LOG_AUTHPRIV,
				"LOG_CRIT":     q.LOG_CRIT,
				"LOG_CRON":     q.LOG_CRON,
				"LOG_DAEMON":   q.LOG_DAEMON,
				"LOG_DEBUG":    q.LOG_DEBUG,
				"LOG_EMERG":    q.LOG_EMERG,
				"LOG_ERR":      q.LOG_ERR,
				"LOG_FTP":      q.LOG_FTP,
				"LOG_INFO":     q.LOG_INFO,
				"LOG_KERN":     q.LOG_KERN,
				"LOG_LOCAL0":   q.LOG_LOCAL0,
				"LOG_LOCAL1":   q.LOG_LOCAL1,
				"LOG_LOCAL2":   q.LOG_LOCAL2,
				"LOG_LOCAL3":   q.LOG_LOCAL3,
				"LOG_LOCAL4":   q.LOG_LOCAL4,
				"LOG_LOCAL5":   q.LOG_LOCAL5,
				"LOG_LOCAL6":   q.LOG_LOCAL6,
				"LOG_LOCAL7":   q.LOG_LOCAL7,
				"LOG_LPR":      q.LOG_LPR,
				"LOG_MAIL":     q.LOG_MAIL,
				"LOG_NEWS":     q.LOG_NEWS,
				"LOG_NOTICE":   q.LOG_NOTICE,
				"LOG_SYSLOG":   q.LOG_SYSLOG,
				"LOG_USER":     q.LOG_USER,
				"LOG_UUCP":     q.LOG_UUCP,
				"LOG_WARNING":  q.LOG_WARNING,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
