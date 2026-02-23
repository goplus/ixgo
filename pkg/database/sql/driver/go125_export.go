// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package driver

import (
	q "database/sql/driver"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "driver",
		Path: "database/sql/driver",
		Deps: map[string]string{
			"context": "context",
			"errors":  "errors",
			"fmt":     "fmt",
			"reflect": "reflect",
			"strconv": "strconv",
			"time":    "time",
		},
		Interfaces: map[string]reflect.Type{
			"ColumnConverter":                reflect.TypeOf((*q.ColumnConverter)(nil)).Elem(),
			"Conn":                           reflect.TypeOf((*q.Conn)(nil)).Elem(),
			"ConnBeginTx":                    reflect.TypeOf((*q.ConnBeginTx)(nil)).Elem(),
			"ConnPrepareContext":             reflect.TypeOf((*q.ConnPrepareContext)(nil)).Elem(),
			"Connector":                      reflect.TypeOf((*q.Connector)(nil)).Elem(),
			"Driver":                         reflect.TypeOf((*q.Driver)(nil)).Elem(),
			"DriverContext":                  reflect.TypeOf((*q.DriverContext)(nil)).Elem(),
			"Execer":                         reflect.TypeOf((*q.Execer)(nil)).Elem(),
			"ExecerContext":                  reflect.TypeOf((*q.ExecerContext)(nil)).Elem(),
			"NamedValueChecker":              reflect.TypeOf((*q.NamedValueChecker)(nil)).Elem(),
			"Pinger":                         reflect.TypeOf((*q.Pinger)(nil)).Elem(),
			"Queryer":                        reflect.TypeOf((*q.Queryer)(nil)).Elem(),
			"QueryerContext":                 reflect.TypeOf((*q.QueryerContext)(nil)).Elem(),
			"Result":                         reflect.TypeOf((*q.Result)(nil)).Elem(),
			"Rows":                           reflect.TypeOf((*q.Rows)(nil)).Elem(),
			"RowsColumnTypeDatabaseTypeName": reflect.TypeOf((*q.RowsColumnTypeDatabaseTypeName)(nil)).Elem(),
			"RowsColumnTypeLength":           reflect.TypeOf((*q.RowsColumnTypeLength)(nil)).Elem(),
			"RowsColumnTypeNullable":         reflect.TypeOf((*q.RowsColumnTypeNullable)(nil)).Elem(),
			"RowsColumnTypePrecisionScale":   reflect.TypeOf((*q.RowsColumnTypePrecisionScale)(nil)).Elem(),
			"RowsColumnTypeScanType":         reflect.TypeOf((*q.RowsColumnTypeScanType)(nil)).Elem(),
			"RowsNextResultSet":              reflect.TypeOf((*q.RowsNextResultSet)(nil)).Elem(),
			"SessionResetter":                reflect.TypeOf((*q.SessionResetter)(nil)).Elem(),
			"Stmt":                           reflect.TypeOf((*q.Stmt)(nil)).Elem(),
			"StmtExecContext":                reflect.TypeOf((*q.StmtExecContext)(nil)).Elem(),
			"StmtQueryContext":               reflect.TypeOf((*q.StmtQueryContext)(nil)).Elem(),
			"Tx":                             reflect.TypeOf((*q.Tx)(nil)).Elem(),
			"Validator":                      reflect.TypeOf((*q.Validator)(nil)).Elem(),
			"Value":                          reflect.TypeOf((*q.Value)(nil)).Elem(),
			"ValueConverter":                 reflect.TypeOf((*q.ValueConverter)(nil)).Elem(),
			"Valuer":                         reflect.TypeOf((*q.Valuer)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"IsolationLevel": reflect.TypeOf((*q.IsolationLevel)(nil)).Elem(),
			"NamedValue":     reflect.TypeOf((*q.NamedValue)(nil)).Elem(),
			"NotNull":        reflect.TypeOf((*q.NotNull)(nil)).Elem(),
			"Null":           reflect.TypeOf((*q.Null)(nil)).Elem(),
			"RowsAffected":   reflect.TypeOf((*q.RowsAffected)(nil)).Elem(),
			"TxOptions":      reflect.TypeOf((*q.TxOptions)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Bool":                      reflect.ValueOf(&q.Bool),
			"DefaultParameterConverter": reflect.ValueOf(&q.DefaultParameterConverter),
			"ErrBadConn":                reflect.ValueOf(&q.ErrBadConn),
			"ErrRemoveArgument":         reflect.ValueOf(&q.ErrRemoveArgument),
			"ErrSkip":                   reflect.ValueOf(&q.ErrSkip),
			"Int32":                     reflect.ValueOf(&q.Int32),
			"ResultNoRows":              reflect.ValueOf(&q.ResultNoRows),
			"String":                    reflect.ValueOf(&q.String),
		},
		Funcs: map[string]reflect.Value{
			"IsScanValue": reflect.ValueOf(q.IsScanValue),
			"IsValue":     reflect.ValueOf(q.IsValue),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
