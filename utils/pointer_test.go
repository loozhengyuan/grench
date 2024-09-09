package utils

import (
	"testing"
)

func TestPointer_Int(t *testing.T) {
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
			got := *Pointer(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestPointer_Int32(t *testing.T) {
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
			got := *Pointer(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestPointer_Int64(t *testing.T) {
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
			got := *Pointer(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestPointer_Float32(t *testing.T) {
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
			got := *Pointer(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestPointer_Float64(t *testing.T) {
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
			got := *Pointer(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestPointer_Bool(t *testing.T) {
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
			got := *Pointer(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestPointer_String(t *testing.T) {
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
			got := *Pointer(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}
