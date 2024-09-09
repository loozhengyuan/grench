package utils

// Coalesce returns the first non-zero value from the input.
func Coalesce[T comparable](input ...T) T {
	var x T
	fn := func(v T) bool {
		return v != x
	}
	return CoalesceFunc(fn, input...)
}

// CoalesceFunc returns the first non-zero value from the input.
func CoalesceFunc[T any](fn func(v T) bool, input ...T) T {
	for _, v := range input {
		if fn(v) {
			return v
		}
	}
	return *new(T)
}
