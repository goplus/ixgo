// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package sql

import (
	q "database/sql"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("database/sql", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "sql",
			Path: "database/sql",
			Deps: map[string]string{
				"bytes":               "bytes",
				"context":             "context",
				"database/sql/driver": "driver",
				"errors":              "errors",
				"fmt":                 "fmt",
				"io":                  "io",
				"maps":                "maps",
				"math/rand/v2":        "rand",
				"reflect":             "reflect",
				"runtime":             "runtime",
				"slices":              "slices",
				"strconv":             "strconv",
				"sync":                "sync",
				"sync/atomic":         "atomic",
				"time":                "time",
				"unicode":             "unicode",
				"unicode/utf8":        "utf8",
				"unsafe":              "unsafe",
			},
			Interfaces: map[string]reflect.Type{
				"Result":  reflect.TypeOf((*q.Result)(nil)).Elem(),
				"Scanner": reflect.TypeOf((*q.Scanner)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"ColumnType":     reflect.TypeOf((*q.ColumnType)(nil)).Elem(),
				"Conn":           reflect.TypeOf((*q.Conn)(nil)).Elem(),
				"DB":             reflect.TypeOf((*q.DB)(nil)).Elem(),
				"DBStats":        reflect.TypeOf((*q.DBStats)(nil)).Elem(),
				"IsolationLevel": reflect.TypeOf((*q.IsolationLevel)(nil)).Elem(),
				"NamedArg":       reflect.TypeOf((*q.NamedArg)(nil)).Elem(),
				"NullBool":       reflect.TypeOf((*q.NullBool)(nil)).Elem(),
				"NullByte":       reflect.TypeOf((*q.NullByte)(nil)).Elem(),
				"NullFloat64":    reflect.TypeOf((*q.NullFloat64)(nil)).Elem(),
				"NullInt16":      reflect.TypeOf((*q.NullInt16)(nil)).Elem(),
				"NullInt32":      reflect.TypeOf((*q.NullInt32)(nil)).Elem(),
				"NullInt64":      reflect.TypeOf((*q.NullInt64)(nil)).Elem(),
				"NullString":     reflect.TypeOf((*q.NullString)(nil)).Elem(),
				"NullTime":       reflect.TypeOf((*q.NullTime)(nil)).Elem(),
				"Out":            reflect.TypeOf((*q.Out)(nil)).Elem(),
				"RawBytes":       reflect.TypeOf((*q.RawBytes)(nil)).Elem(),
				"Row":            reflect.TypeOf((*q.Row)(nil)).Elem(),
				"Rows":           reflect.TypeOf((*q.Rows)(nil)).Elem(),
				"Stmt":           reflect.TypeOf((*q.Stmt)(nil)).Elem(),
				"Tx":             reflect.TypeOf((*q.Tx)(nil)).Elem(),
				"TxOptions":      reflect.TypeOf((*q.TxOptions)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]reflect.Value{
				"ErrConnDone": reflect.ValueOf(&q.ErrConnDone),
				"ErrNoRows":   reflect.ValueOf(&q.ErrNoRows),
				"ErrTxDone":   reflect.ValueOf(&q.ErrTxDone),
			},
			Funcs: map[string]reflect.Value{
				"Drivers":  reflect.ValueOf(q.Drivers),
				"Named":    reflect.ValueOf(q.Named),
				"Open":     reflect.ValueOf(q.Open),
				"OpenDB":   reflect.ValueOf(q.OpenDB),
				"Register": reflect.ValueOf(q.Register),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"LevelDefault":         {Typ: reflect.TypeOf(q.LevelDefault), Value: constant.MakeInt64(int64(q.LevelDefault))},
				"LevelLinearizable":    {Typ: reflect.TypeOf(q.LevelLinearizable), Value: constant.MakeInt64(int64(q.LevelLinearizable))},
				"LevelReadCommitted":   {Typ: reflect.TypeOf(q.LevelReadCommitted), Value: constant.MakeInt64(int64(q.LevelReadCommitted))},
				"LevelReadUncommitted": {Typ: reflect.TypeOf(q.LevelReadUncommitted), Value: constant.MakeInt64(int64(q.LevelReadUncommitted))},
				"LevelRepeatableRead":  {Typ: reflect.TypeOf(q.LevelRepeatableRead), Value: constant.MakeInt64(int64(q.LevelRepeatableRead))},
				"LevelSerializable":    {Typ: reflect.TypeOf(q.LevelSerializable), Value: constant.MakeInt64(int64(q.LevelSerializable))},
				"LevelSnapshot":        {Typ: reflect.TypeOf(q.LevelSnapshot), Value: constant.MakeInt64(int64(q.LevelSnapshot))},
				"LevelWriteCommitted":  {Typ: reflect.TypeOf(q.LevelWriteCommitted), Value: constant.MakeInt64(int64(q.LevelWriteCommitted))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
