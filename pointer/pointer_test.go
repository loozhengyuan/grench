package pointer

import (
	"testing"
)

func TestInt(t *testing.T) {
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
			got := *Int(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestInt32(t *testing.T) {
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
			got := *Int32(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestInt64(t *testing.T) {
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
			got := *Int64(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
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
			got := *Float32(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
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
			got := *Float64(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestBool(t *testing.T) {
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
			got := *Bool(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestString(t *testing.T) {
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
			got := *String(tc.input) // deference from output
			if got != tc.input {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}
