package sql

import (
	_ "database/sql"
	_ "unsafe"

	"database/sql/driver"
)

// Null represents a value that may be null.
// Null implements the [Scanner] interface so
// it can be used as a scan destination:
//
//	var s Null[string]
//	err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&s)
//	...
//	if s.Valid {
//	   // use s.V
//	} else {
//	   // NULL value
//	}
//
// T should be one of the types accepted by [driver.Value].
type Null[T any] struct {
	V     T
	Valid bool
}

func (n *Null[T]) Scan(value any) error {
	if value == nil {
		n.V, n.Valid = *new(T), false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.V, value)
}

func (n Null[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	v := any(n.V)
	// See issue 69728.
	if valuer, ok := v.(driver.Valuer); ok {
		val, err := callValuerValue(valuer)
		if err != nil {
			return val, err
		}
		v = val
	}
	// See issue 69837.
	return driver.DefaultParameterConverter.ConvertValue(v)
}

//go:linkname convertAssign database/sql.convertAssign
func convertAssign(dest, src any) error

//go:linkname callValuerValue database/sql.callValuerValue
func callValuerValue(vr driver.Valuer) (v driver.Value, err error)
