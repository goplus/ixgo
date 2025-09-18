// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package syslog

import (
	q "log/syslog"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
			"LOG_ALERT":    {reflect.TypeOf(q.LOG_ALERT), constant.MakeInt64(int64(q.LOG_ALERT))},
			"LOG_AUTH":     {reflect.TypeOf(q.LOG_AUTH), constant.MakeInt64(int64(q.LOG_AUTH))},
			"LOG_AUTHPRIV": {reflect.TypeOf(q.LOG_AUTHPRIV), constant.MakeInt64(int64(q.LOG_AUTHPRIV))},
			"LOG_CRIT":     {reflect.TypeOf(q.LOG_CRIT), constant.MakeInt64(int64(q.LOG_CRIT))},
			"LOG_CRON":     {reflect.TypeOf(q.LOG_CRON), constant.MakeInt64(int64(q.LOG_CRON))},
			"LOG_DAEMON":   {reflect.TypeOf(q.LOG_DAEMON), constant.MakeInt64(int64(q.LOG_DAEMON))},
			"LOG_DEBUG":    {reflect.TypeOf(q.LOG_DEBUG), constant.MakeInt64(int64(q.LOG_DEBUG))},
			"LOG_EMERG":    {reflect.TypeOf(q.LOG_EMERG), constant.MakeInt64(int64(q.LOG_EMERG))},
			"LOG_ERR":      {reflect.TypeOf(q.LOG_ERR), constant.MakeInt64(int64(q.LOG_ERR))},
			"LOG_FTP":      {reflect.TypeOf(q.LOG_FTP), constant.MakeInt64(int64(q.LOG_FTP))},
			"LOG_INFO":     {reflect.TypeOf(q.LOG_INFO), constant.MakeInt64(int64(q.LOG_INFO))},
			"LOG_KERN":     {reflect.TypeOf(q.LOG_KERN), constant.MakeInt64(int64(q.LOG_KERN))},
			"LOG_LOCAL0":   {reflect.TypeOf(q.LOG_LOCAL0), constant.MakeInt64(int64(q.LOG_LOCAL0))},
			"LOG_LOCAL1":   {reflect.TypeOf(q.LOG_LOCAL1), constant.MakeInt64(int64(q.LOG_LOCAL1))},
			"LOG_LOCAL2":   {reflect.TypeOf(q.LOG_LOCAL2), constant.MakeInt64(int64(q.LOG_LOCAL2))},
			"LOG_LOCAL3":   {reflect.TypeOf(q.LOG_LOCAL3), constant.MakeInt64(int64(q.LOG_LOCAL3))},
			"LOG_LOCAL4":   {reflect.TypeOf(q.LOG_LOCAL4), constant.MakeInt64(int64(q.LOG_LOCAL4))},
			"LOG_LOCAL5":   {reflect.TypeOf(q.LOG_LOCAL5), constant.MakeInt64(int64(q.LOG_LOCAL5))},
			"LOG_LOCAL6":   {reflect.TypeOf(q.LOG_LOCAL6), constant.MakeInt64(int64(q.LOG_LOCAL6))},
			"LOG_LOCAL7":   {reflect.TypeOf(q.LOG_LOCAL7), constant.MakeInt64(int64(q.LOG_LOCAL7))},
			"LOG_LPR":      {reflect.TypeOf(q.LOG_LPR), constant.MakeInt64(int64(q.LOG_LPR))},
			"LOG_MAIL":     {reflect.TypeOf(q.LOG_MAIL), constant.MakeInt64(int64(q.LOG_MAIL))},
			"LOG_NEWS":     {reflect.TypeOf(q.LOG_NEWS), constant.MakeInt64(int64(q.LOG_NEWS))},
			"LOG_NOTICE":   {reflect.TypeOf(q.LOG_NOTICE), constant.MakeInt64(int64(q.LOG_NOTICE))},
			"LOG_SYSLOG":   {reflect.TypeOf(q.LOG_SYSLOG), constant.MakeInt64(int64(q.LOG_SYSLOG))},
			"LOG_USER":     {reflect.TypeOf(q.LOG_USER), constant.MakeInt64(int64(q.LOG_USER))},
			"LOG_UUCP":     {reflect.TypeOf(q.LOG_UUCP), constant.MakeInt64(int64(q.LOG_UUCP))},
			"LOG_WARNING":  {reflect.TypeOf(q.LOG_WARNING), constant.MakeInt64(int64(q.LOG_WARNING))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
