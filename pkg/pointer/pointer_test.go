package pointer

import (
	"testing"
)

func TestIntPtr(t *testing.T) {
	cases := map[string]struct {
		input int
	}{
		"default_value":     {},
		"non_default_value": {input: 1},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := *IntPtr(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestInt32Ptr(t *testing.T) {
	cases := map[string]struct {
		input int32
	}{
		"default_value":     {},
		"non_default_value": {input: 1},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := *Int32Ptr(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestInt64Ptr(t *testing.T) {
	cases := map[string]struct {
		input int64
	}{
		"default_value":     {},
		"non_default_value": {input: 1},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := *Int64Ptr(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestFloat32Ptr(t *testing.T) {
	cases := map[string]struct {
		input float32
	}{
		"default_value":     {},
		"non_default_value": {input: 1},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := *Float32Ptr(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestFloat64Ptr(t *testing.T) {
	cases := map[string]struct {
		input float64
	}{
		"default_value":     {},
		"non_default_value": {input: 1},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := *Float64Ptr(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestBoolPtr(t *testing.T) {
	cases := map[string]struct {
		input bool
	}{
		"default_value":     {},
		"non_default_value": {input: true},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := *BoolPtr(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestStringPtr(t *testing.T) {
	cases := map[string]struct {
		input string
	}{
		"default_value":     {},
		"non_default_value": {input: "abc"},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := *StringPtr(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}
