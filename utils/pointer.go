package utils

// Pointer returns the pointer of the input.
func Pointer[T any](v T) *T {
	return &v
}
