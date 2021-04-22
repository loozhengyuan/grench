// Package coalesce provides coalesce helper functions.
package coalesce

// Int returns the first non-zero value from the int input.
func Int(i ...int) int {
	var x int
	for _, v := range i {
		if v != x {
			return v
		}
	}
	return x
}

// Int32 returns the first non-zero value from the int32 input.
func Int32(i ...int32) int32 {
	var x int32
	for _, v := range i {
		if v != x {
			return v
		}
	}
	return x
}

// Int64 returns the first non-zero value from the int64 input.
func Int64(i ...int64) int64 {
	var x int64
	for _, v := range i {
		if v != x {
			return v
		}
	}
	return x
}

// Float32 returns the first non-zero value from the float32 input.
func Float32(f ...float32) float32 {
	var x float32
	for _, v := range f {
		if v != x {
			return v
		}
	}
	return x
}

// Float64 returns the first non-zero value from the float64 input.
func Float64(f ...float64) float64 {
	var x float64
	for _, v := range f {
		if v != x {
			return v
		}
	}
	return x
}

// String returns the first non-empty value from the string input.
func String(s ...string) string {
	var x string
	for _, v := range s {
		if v != x {
			return v
		}
	}
	return x
}

// Interface returns the first non-nil value from the interface{} input.
func Interface(i ...interface{}) interface{} {
	var x interface{}
	for _, v := range i {
		if v != x {
			return v
		}
	}
	return x
}
