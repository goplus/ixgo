//go:build go1.22
// +build go1.22

package sql

import (
	"database/sql/driver"
	_ "embed"
	"reflect"
	_ "unsafe"

	"github.com/goplus/ixgo"
)

//go:linkname convertAssign database/sql.convertAssign
func convertAssign(dest, src any) error

//go:embed _patch/sql.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("database/sql.convertAssign", convertAssign)
	ixgo.RegisterExternal("database/sql.callValuerValue", callValuerValue)
	ixgo.RegisterPatch("database/sql", patch_data)
}

var valuerReflectType = reflect.TypeFor[driver.Valuer]()

// callValuerValue returns vr.Value(), with one exception:
// If vr.Value is an auto-generated method on a pointer type and the
// pointer is nil, it would panic at runtime in the panicwrap
// method. Treat it like nil instead.
// Issue 8415.
//
// This is so people can implement driver.Value on value types and
// still use nil pointers to those types to mean nil/NULL, just like
// string/*string.
//
// This function is mirrored in the database/sql package.
func callValuerValue(vr driver.Valuer) (v driver.Value, err error) {
	if rv := reflect.ValueOf(vr); rv.Kind() == reflect.Pointer &&
		rv.IsNil() &&
		rv.Type().Elem().Implements(valuerReflectType) {
		return nil, nil
	}
	return vr.Value()
}
