// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package syslog

import (
	q "log/syslog"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("log/syslog", func() *ixgo.Package {
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
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"Dial":      reflect.ValueOf(q.Dial),
				"New":       reflect.ValueOf(q.New),
				"NewLogger": reflect.ValueOf(q.NewLogger),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"LOG_ALERT":    {Typ: reflect.TypeOf(q.LOG_ALERT), Value: constant.MakeInt64(int64(q.LOG_ALERT))},
				"LOG_AUTH":     {Typ: reflect.TypeOf(q.LOG_AUTH), Value: constant.MakeInt64(int64(q.LOG_AUTH))},
				"LOG_AUTHPRIV": {Typ: reflect.TypeOf(q.LOG_AUTHPRIV), Value: constant.MakeInt64(int64(q.LOG_AUTHPRIV))},
				"LOG_CRIT":     {Typ: reflect.TypeOf(q.LOG_CRIT), Value: constant.MakeInt64(int64(q.LOG_CRIT))},
				"LOG_CRON":     {Typ: reflect.TypeOf(q.LOG_CRON), Value: constant.MakeInt64(int64(q.LOG_CRON))},
				"LOG_DAEMON":   {Typ: reflect.TypeOf(q.LOG_DAEMON), Value: constant.MakeInt64(int64(q.LOG_DAEMON))},
				"LOG_DEBUG":    {Typ: reflect.TypeOf(q.LOG_DEBUG), Value: constant.MakeInt64(int64(q.LOG_DEBUG))},
				"LOG_EMERG":    {Typ: reflect.TypeOf(q.LOG_EMERG), Value: constant.MakeInt64(int64(q.LOG_EMERG))},
				"LOG_ERR":      {Typ: reflect.TypeOf(q.LOG_ERR), Value: constant.MakeInt64(int64(q.LOG_ERR))},
				"LOG_FTP":      {Typ: reflect.TypeOf(q.LOG_FTP), Value: constant.MakeInt64(int64(q.LOG_FTP))},
				"LOG_INFO":     {Typ: reflect.TypeOf(q.LOG_INFO), Value: constant.MakeInt64(int64(q.LOG_INFO))},
				"LOG_KERN":     {Typ: reflect.TypeOf(q.LOG_KERN), Value: constant.MakeInt64(int64(q.LOG_KERN))},
				"LOG_LOCAL0":   {Typ: reflect.TypeOf(q.LOG_LOCAL0), Value: constant.MakeInt64(int64(q.LOG_LOCAL0))},
				"LOG_LOCAL1":   {Typ: reflect.TypeOf(q.LOG_LOCAL1), Value: constant.MakeInt64(int64(q.LOG_LOCAL1))},
				"LOG_LOCAL2":   {Typ: reflect.TypeOf(q.LOG_LOCAL2), Value: constant.MakeInt64(int64(q.LOG_LOCAL2))},
				"LOG_LOCAL3":   {Typ: reflect.TypeOf(q.LOG_LOCAL3), Value: constant.MakeInt64(int64(q.LOG_LOCAL3))},
				"LOG_LOCAL4":   {Typ: reflect.TypeOf(q.LOG_LOCAL4), Value: constant.MakeInt64(int64(q.LOG_LOCAL4))},
				"LOG_LOCAL5":   {Typ: reflect.TypeOf(q.LOG_LOCAL5), Value: constant.MakeInt64(int64(q.LOG_LOCAL5))},
				"LOG_LOCAL6":   {Typ: reflect.TypeOf(q.LOG_LOCAL6), Value: constant.MakeInt64(int64(q.LOG_LOCAL6))},
				"LOG_LOCAL7":   {Typ: reflect.TypeOf(q.LOG_LOCAL7), Value: constant.MakeInt64(int64(q.LOG_LOCAL7))},
				"LOG_LPR":      {Typ: reflect.TypeOf(q.LOG_LPR), Value: constant.MakeInt64(int64(q.LOG_LPR))},
				"LOG_MAIL":     {Typ: reflect.TypeOf(q.LOG_MAIL), Value: constant.MakeInt64(int64(q.LOG_MAIL))},
				"LOG_NEWS":     {Typ: reflect.TypeOf(q.LOG_NEWS), Value: constant.MakeInt64(int64(q.LOG_NEWS))},
				"LOG_NOTICE":   {Typ: reflect.TypeOf(q.LOG_NOTICE), Value: constant.MakeInt64(int64(q.LOG_NOTICE))},
				"LOG_SYSLOG":   {Typ: reflect.TypeOf(q.LOG_SYSLOG), Value: constant.MakeInt64(int64(q.LOG_SYSLOG))},
				"LOG_USER":     {Typ: reflect.TypeOf(q.LOG_USER), Value: constant.MakeInt64(int64(q.LOG_USER))},
				"LOG_UUCP":     {Typ: reflect.TypeOf(q.LOG_UUCP), Value: constant.MakeInt64(int64(q.LOG_UUCP))},
				"LOG_WARNING":  {Typ: reflect.TypeOf(q.LOG_WARNING), Value: constant.MakeInt64(int64(q.LOG_WARNING))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
