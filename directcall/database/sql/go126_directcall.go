// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package sql

import (
	q "database/sql"

	context "context"
	driver "database/sql/driver"
	"github.com/goplus/ixgo"
	time "time"
)

func init() {
	ixgo.RegisterDirectCalls("database/sql", map[string]ixgo.DirectCallAdapter{
		"(*ColumnType).DatabaseTypeName": method_ptr_ColumnType_DatabaseTypeName,
		"(*ColumnType).Name":             method_ptr_ColumnType_Name,
		"(*ColumnType).ScanType":         method_ptr_ColumnType_ScanType,
		"(*Conn).Close":                  method_ptr_Conn_Close,
		"(*Conn).PingContext":            method_ptr_Conn_PingContext,
		"(*Conn).QueryRowContext":        method_ptr_Conn_QueryRowContext,
		"(*Conn).Raw":                    method_ptr_Conn_Raw,
		"(*DB).Close":                    method_ptr_DB_Close,
		"(*DB).Driver":                   method_ptr_DB_Driver,
		"(*DB).Ping":                     method_ptr_DB_Ping,
		"(*DB).PingContext":              method_ptr_DB_PingContext,
		"(*DB).QueryRow":                 method_ptr_DB_QueryRow,
		"(*DB).QueryRowContext":          method_ptr_DB_QueryRowContext,
		"(*DB).SetConnMaxIdleTime":       method_ptr_DB_SetConnMaxIdleTime,
		"(*DB).SetConnMaxLifetime":       method_ptr_DB_SetConnMaxLifetime,
		"(*DB).SetMaxIdleConns":          method_ptr_DB_SetMaxIdleConns,
		"(*DB).SetMaxOpenConns":          method_ptr_DB_SetMaxOpenConns,
		"(*DB).Stats":                    method_ptr_DB_Stats,
		"(*IsolationLevel).String":       method_ptr_IsolationLevel_String,
		"(*NullBool).Scan":               method_ptr_NullBool_Scan,
		"(*NullByte).Scan":               method_ptr_NullByte_Scan,
		"(*NullFloat64).Scan":            method_ptr_NullFloat64_Scan,
		"(*NullInt16).Scan":              method_ptr_NullInt16_Scan,
		"(*NullInt32).Scan":              method_ptr_NullInt32_Scan,
		"(*NullInt64).Scan":              method_ptr_NullInt64_Scan,
		"(*NullString).Scan":             method_ptr_NullString_Scan,
		"(*NullTime).Scan":               method_ptr_NullTime_Scan,
		"(*Row).Err":                     method_ptr_Row_Err,
		"(*Row).Scan":                    method_ptr_Row_Scan,
		"(*Rows).Close":                  method_ptr_Rows_Close,
		"(*Rows).Err":                    method_ptr_Rows_Err,
		"(*Rows).Next":                   method_ptr_Rows_Next,
		"(*Rows).NextResultSet":          method_ptr_Rows_NextResultSet,
		"(*Rows).Scan":                   method_ptr_Rows_Scan,
		"(*Stmt).Close":                  method_ptr_Stmt_Close,
		"(*Stmt).QueryRow":               method_ptr_Stmt_QueryRow,
		"(*Stmt).QueryRowContext":        method_ptr_Stmt_QueryRowContext,
		"(*Tx).Commit":                   method_ptr_Tx_Commit,
		"(*Tx).QueryRow":                 method_ptr_Tx_QueryRow,
		"(*Tx).QueryRowContext":          method_ptr_Tx_QueryRowContext,
		"(*Tx).Rollback":                 method_ptr_Tx_Rollback,
		"(*Tx).Stmt":                     method_ptr_Tx_Stmt,
		"(*Tx).StmtContext":              method_ptr_Tx_StmtContext,
		"(IsolationLevel).String":        method_IsolationLevel_String,
		"Drivers":                        func_Drivers,
		"Named":                          func_Named,
		"OpenDB":                         func_OpenDB,
		"Register":                       func_Register,
	})
}

func method_ptr_ColumnType_DatabaseTypeName(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ColumnType).DatabaseTypeName(ixgo.DirectCallArg[*q.ColumnType](ctx, 0)))
}

func method_ptr_ColumnType_Name(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ColumnType).Name(ixgo.DirectCallArg[*q.ColumnType](ctx, 0)))
}

func method_ptr_ColumnType_ScanType(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ColumnType).ScanType(ixgo.DirectCallArg[*q.ColumnType](ctx, 0)))
}

func method_ptr_Conn_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).Close(ixgo.DirectCallArg[*q.Conn](ctx, 0)))
}

func method_ptr_Conn_PingContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).PingContext(ixgo.DirectCallArg[*q.Conn](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1)))
}

func method_ptr_Conn_QueryRowContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).QueryRowContext(ixgo.DirectCallArg[*q.Conn](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[[]any](ctx, 3)...))
}

func method_ptr_Conn_Raw(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Conn).Raw(ixgo.DirectCallArg[*q.Conn](ctx, 0), ixgo.DirectCallArg[func(driverConn any) error](ctx, 1)))
}

func method_ptr_DB_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DB).Close(ixgo.DirectCallArg[*q.DB](ctx, 0)))
}

func method_ptr_DB_Driver(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DB).Driver(ixgo.DirectCallArg[*q.DB](ctx, 0)))
}

func method_ptr_DB_Ping(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DB).Ping(ixgo.DirectCallArg[*q.DB](ctx, 0)))
}

func method_ptr_DB_PingContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DB).PingContext(ixgo.DirectCallArg[*q.DB](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1)))
}

func method_ptr_DB_QueryRow(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DB).QueryRow(ixgo.DirectCallArg[*q.DB](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[[]any](ctx, 2)...))
}

func method_ptr_DB_QueryRowContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DB).QueryRowContext(ixgo.DirectCallArg[*q.DB](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[[]any](ctx, 3)...))
}

func method_ptr_DB_SetConnMaxIdleTime(ctx ixgo.DirectCallContext) {
	(*q.DB).SetConnMaxIdleTime(ixgo.DirectCallArg[*q.DB](ctx, 0), ixgo.DirectCallArg[time.Duration](ctx, 1))
}

func method_ptr_DB_SetConnMaxLifetime(ctx ixgo.DirectCallContext) {
	(*q.DB).SetConnMaxLifetime(ixgo.DirectCallArg[*q.DB](ctx, 0), ixgo.DirectCallArg[time.Duration](ctx, 1))
}

func method_ptr_DB_SetMaxIdleConns(ctx ixgo.DirectCallContext) {
	(*q.DB).SetMaxIdleConns(ixgo.DirectCallArg[*q.DB](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_DB_SetMaxOpenConns(ctx ixgo.DirectCallContext) {
	(*q.DB).SetMaxOpenConns(ixgo.DirectCallArg[*q.DB](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_DB_Stats(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DB).Stats(ixgo.DirectCallArg[*q.DB](ctx, 0)))
}

func func_Drivers(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Drivers())
}

func method_IsolationLevel_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsolationLevel.String(ixgo.DirectCallArg[q.IsolationLevel](ctx, 0)))
}

func method_ptr_IsolationLevel_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.IsolationLevel).String(ixgo.DirectCallArg[*q.IsolationLevel](ctx, 0)))
}

func func_Named(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Named(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_NullBool_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NullBool).Scan(ixgo.DirectCallArg[*q.NullBool](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_NullByte_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NullByte).Scan(ixgo.DirectCallArg[*q.NullByte](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_NullFloat64_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NullFloat64).Scan(ixgo.DirectCallArg[*q.NullFloat64](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_NullInt16_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NullInt16).Scan(ixgo.DirectCallArg[*q.NullInt16](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_NullInt32_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NullInt32).Scan(ixgo.DirectCallArg[*q.NullInt32](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_NullInt64_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NullInt64).Scan(ixgo.DirectCallArg[*q.NullInt64](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_NullString_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NullString).Scan(ixgo.DirectCallArg[*q.NullString](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_NullTime_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NullTime).Scan(ixgo.DirectCallArg[*q.NullTime](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func func_OpenDB(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OpenDB(ixgo.DirectCallArg[driver.Connector](ctx, 0)))
}

func func_Register(ctx ixgo.DirectCallContext) {
	q.Register(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[driver.Driver](ctx, 1))
}

func method_ptr_Row_Err(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Row).Err(ixgo.DirectCallArg[*q.Row](ctx, 0)))
}

func method_ptr_Row_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Row).Scan(ixgo.DirectCallArg[*q.Row](ctx, 0), ixgo.DirectCallArg[[]any](ctx, 1)...))
}

func method_ptr_Rows_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rows).Close(ixgo.DirectCallArg[*q.Rows](ctx, 0)))
}

func method_ptr_Rows_Err(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rows).Err(ixgo.DirectCallArg[*q.Rows](ctx, 0)))
}

func method_ptr_Rows_Next(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rows).Next(ixgo.DirectCallArg[*q.Rows](ctx, 0)))
}

func method_ptr_Rows_NextResultSet(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rows).NextResultSet(ixgo.DirectCallArg[*q.Rows](ctx, 0)))
}

func method_ptr_Rows_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rows).Scan(ixgo.DirectCallArg[*q.Rows](ctx, 0), ixgo.DirectCallArg[[]any](ctx, 1)...))
}

func method_ptr_Stmt_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Stmt).Close(ixgo.DirectCallArg[*q.Stmt](ctx, 0)))
}

func method_ptr_Stmt_QueryRow(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Stmt).QueryRow(ixgo.DirectCallArg[*q.Stmt](ctx, 0), ixgo.DirectCallArg[[]any](ctx, 1)...))
}

func method_ptr_Stmt_QueryRowContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Stmt).QueryRowContext(ixgo.DirectCallArg[*q.Stmt](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1), ixgo.DirectCallArg[[]any](ctx, 2)...))
}

func method_ptr_Tx_Commit(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tx).Commit(ixgo.DirectCallArg[*q.Tx](ctx, 0)))
}

func method_ptr_Tx_QueryRow(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tx).QueryRow(ixgo.DirectCallArg[*q.Tx](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[[]any](ctx, 2)...))
}

func method_ptr_Tx_QueryRowContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tx).QueryRowContext(ixgo.DirectCallArg[*q.Tx](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[[]any](ctx, 3)...))
}

func method_ptr_Tx_Rollback(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tx).Rollback(ixgo.DirectCallArg[*q.Tx](ctx, 0)))
}

func method_ptr_Tx_Stmt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tx).Stmt(ixgo.DirectCallArg[*q.Tx](ctx, 0), ixgo.DirectCallArg[*q.Stmt](ctx, 1)))
}

func method_ptr_Tx_StmtContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tx).StmtContext(ixgo.DirectCallArg[*q.Tx](ctx, 0), ixgo.DirectCallArg[context.Context](ctx, 1), ixgo.DirectCallArg[*q.Stmt](ctx, 2)))
}
