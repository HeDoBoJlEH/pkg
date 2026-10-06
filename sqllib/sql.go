// Package sqllib provides utilities for layer repository functions.
package sqllib

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type nullableValue interface {
	Scan(value any) error
	Value() (driver.Value, error)
}

// Nullable returns sql.NullValue, contains value if valid, and nil if not.
//
// Great for PATCH requests with using COALESCE(). 
func Nullable(val any) nullableValue {
	switch val := val.(type) {
	case string: 
		return &sql.NullString{String: val, Valid: val != ""}
	case int64:
		return &sql.NullInt64{Int64: val, Valid: val > 0}
	case int32:
		return &sql.NullInt32{Int32: val, Valid: val > 0}
	case int16:
		return &sql.NullInt16{Int16: val, Valid: val > 0}
	case time.Time:
		return &sql.NullTime{Time: val, Valid: !val.IsZero()}
	case float64:
		return &sql.NullFloat64{Float64: val, Valid: val != 0}
	default: 
		return &sql.NullString{Valid: false}
	}
}