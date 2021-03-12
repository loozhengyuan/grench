// Package types provides the reusable types.
package types

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	// ErrUnhandledType is returned when a particular type is not
	// handled. This typically used for handling the default case
	// in type switches.
	ErrUnhandledType = errors.New("unhandled type")
)

// NullUUID is a type wrapper for a nullable UUID type.
type NullUUID struct {
	String string
	Valid  bool
}

// Ensures that NullUUID implements the driver.Valuer and sql.Scanner interfaces.
var _ driver.Valuer = (*NullUUID)(nil)
var _ sql.Scanner = (*NullUUID)(nil)

// Value returns the NullUUID as a driver.Value.
func (t NullUUID) Value() (driver.Value, error) {
	if !t.Valid {
		return []byte(nil), nil
	}
	return []byte(t.String), nil
}

// Scan scans a value into the typed object.
func (t *NullUUID) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		t.Valid = false
	case []byte:
		t.String = string(v)
		t.Valid = true
		if bytes.Equal(v, []byte{}) {
			t.Valid = false
		}
	case string:
		t.String = v
		t.Valid = true
	default:
		return fmt.Errorf("%w: %T", ErrUnhandledType, v)
	}
	return nil
}

// Ensures that NullUUID implements the json.Marshaler and json.Unmarshaler interfaces.
var _ json.Marshaler = (*NullUUID)(nil)
var _ json.Unmarshaler = (*NullUUID)(nil)

// MarshalJSON encodes the typed object as JSON format.
func (t NullUUID) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte(nil), nil
	}
	return []byte(t.String), nil
}

// UnmarshalJSON decodes the JSON data into the typed object.
func (t *NullUUID) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	t.String = string(data)
	t.Valid = true
	return nil
}
