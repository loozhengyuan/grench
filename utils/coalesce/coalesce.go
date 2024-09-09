// Package coalesce provides coalesce helper functions.
//
// Deprecated: Use the [utils.Coalesce] function instead.
package coalesce

import (
	"github.com/loozhengyuan/grench/utils"
)

// Int returns the first non-zero value from the int input.
func Int(i ...int) int {
	return utils.Coalesce(i...)
}

// Int32 returns the first non-zero value from the int32 input.
func Int32(i ...int32) int32 {
	return utils.Coalesce(i...)
}

// Int64 returns the first non-zero value from the int64 input.
func Int64(i ...int64) int64 {
	return utils.Coalesce(i...)
}

// Float32 returns the first non-zero value from the float32 input.
func Float32(f ...float32) float32 {
	return utils.Coalesce(f...)
}

// Float64 returns the first non-zero value from the float64 input.
func Float64(f ...float64) float64 {
	return utils.Coalesce(f...)
}

// String returns the first non-empty value from the string input.
func String(s ...string) string {
	return utils.Coalesce(s...)
}

// Interface returns the first non-nil value from the interface{} input.
func Interface(i ...interface{}) interface{} {
	var x interface{}
	return utils.CoalesceFunc(func(v interface{}) bool { return v != x }, i...)
}
