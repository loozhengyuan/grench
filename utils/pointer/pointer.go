// Package pointer provides pointer helper functions.
//
// Deprecated: Use the [utils.Coalesce] function instead.
package pointer

import (
	"github.com/loozhengyuan/grench/utils"
)

// Int returns the pointer of the int input.
func Int(i int) *int {
	return utils.Pointer(i)
}

// Int32 returns the pointer of the int32 input.
func Int32(i int32) *int32 {
	return utils.Pointer(i)
}

// Int64 returns the pointer of the int64 input.
func Int64(i int64) *int64 {
	return utils.Pointer(i)
}

// Float32 returns the pointer of the float32 input.
func Float32(f float32) *float32 {
	return utils.Pointer(f)
}

// Float64 returns the pointer of the float64 input.
func Float64(f float64) *float64 {
	return utils.Pointer(f)
}

// Bool returns the pointer of the bool input.
func Bool(b bool) *bool {
	return utils.Pointer(b)
}

// String returns the pointer of the string input.
func String(s string) *string {
	return utils.Pointer(s)
}
