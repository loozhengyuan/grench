// Package pointer provides pointer helper functions.
package pointer

// Int returns the pointer of the int input.
func Int(i int) *int {
	return &i
}

// Int32 returns the pointer of the int32 input.
func Int32(i int32) *int32 {
	return &i
}

// Int64 returns the pointer of the int64 input.
func Int64(i int64) *int64 {
	return &i
}

// Float32 returns the pointer of the float32 input.
func Float32(f float32) *float32 {
	return &f
}

// Float64 returns the pointer of the float64 input.
func Float64(f float64) *float64 {
	return &f
}

// Bool returns the pointer of the bool input.
func Bool(b bool) *bool {
	return &b
}

// String returns the pointer of the string input.
func String(s string) *string {
	return &s
}
