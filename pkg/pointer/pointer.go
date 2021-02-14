// Package pointer provides pointer helper functions.
package pointer

// IntPtr returns the pointer of the int input.
func IntPtr(i int) *int {
	return &i
}

// Int32Ptr returns the pointer of the int32 input.
func Int32Ptr(i int32) *int32 {
	return &i
}

// Int64Ptr returns the pointer of the int64 input.
func Int64Ptr(i int64) *int64 {
	return &i
}

// Float32Ptr returns the pointer of the float32 input.
func Float32Ptr(f float32) *float32 {
	return &f
}

// Float64Ptr returns the pointer of the float64 input.
func Float64Ptr(f float64) *float64 {
	return &f
}

// BoolPtr returns the pointer of the bool input.
func BoolPtr(b bool) *bool {
	return &b
}

// StringPtr returns the pointer of the string input.
func StringPtr(s string) *string {
	return &s
}
