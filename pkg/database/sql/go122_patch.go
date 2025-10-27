//go:build go1.22
// +build go1.22

package sql

import (
	_ "embed"
	_ "unsafe"

	"database/sql/driver"

	"github.com/goplus/ixgo"
)

//go:linkname convertAssign database/sql.convertAssign
func convertAssign(dest, src any) error

//go:linkname callValuerValue database/sql.callValuerValue
func callValuerValue(vr driver.Valuer) (v driver.Value, err error)

//go:embed _patch/sql.go
var patch_data []byte

func init() {
	ixgo.RegisterExternal("database/sql.convertAssign", convertAssign)
	ixgo.RegisterExternal("database/sql.callValuerValue", callValuerValue)
	ixgo.RegisterPatch("database/sql", patch_data)
}
