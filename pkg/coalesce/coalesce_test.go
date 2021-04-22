package coalesce

import (
	"testing"
)

func TestInt(t *testing.T) {
	cases := map[string]struct {
		input []int
		want  int
	}{
		"default": {
			input: []int{0, 0, 0},
			want:  0, // default null value
		},
		"match_first": {
			input: []int{1, 2, 3},
			want:  1,
		},
		"match_middle": {
			input: []int{0, 2, 3},
			want:  2,
		},
		"match_last": {
			input: []int{0, 0, 3},
			want:  3,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Int(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	cases := map[string]struct {
		input []int32
		want  int32
	}{
		"default": {
			input: []int32{0, 0, 0},
			want:  0, // default null value
		},
		"match_first": {
			input: []int32{1, 2, 3},
			want:  1,
		},
		"match_middle": {
			input: []int32{0, 2, 3},
			want:  2,
		},
		"match_last": {
			input: []int32{0, 0, 3},
			want:  3,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Int32(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	cases := map[string]struct {
		input []int64
		want  int64
	}{
		"default": {
			input: []int64{0, 0, 0},
			want:  0, // default null value
		},
		"match_first": {
			input: []int64{1, 2, 3},
			want:  1,
		},
		"match_middle": {
			input: []int64{0, 2, 3},
			want:  2,
		},
		"match_last": {
			input: []int64{0, 0, 3},
			want:  3,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Int64(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	cases := map[string]struct {
		input []float32
		want  float32
	}{
		"default": {
			input: []float32{0, 0, 0},
			want:  0, // default null value
		},
		"match_first": {
			input: []float32{1, 2, 3},
			want:  1,
		},
		"match_middle": {
			input: []float32{0, 2, 3},
			want:  2,
		},
		"match_last": {
			input: []float32{0, 0, 3},
			want:  3,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Float32(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	cases := map[string]struct {
		input []float64
		want  float64
	}{
		"default": {
			input: []float64{0, 0, 0},
			want:  0, // default null value
		},
		"match_first": {
			input: []float64{1, 2, 3},
			want:  1,
		},
		"match_middle": {
			input: []float64{0, 2, 3},
			want:  2,
		},
		"match_last": {
			input: []float64{0, 0, 3},
			want:  3,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Float64(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestString(t *testing.T) {
	cases := map[string]struct {
		input []string
		want  string
	}{
		"default": {
			input: []string{"", "", ""},
			want:  "", // default null value
		},
		"match_first": {
			input: []string{"1", "2", "3"},
			want:  "1",
		},
		"match_middle": {
			input: []string{"", "2", "3"},
			want:  "2",
		},
		"match_last": {
			input: []string{"", "", "3"},
			want:  "3",
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := String(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func TestInterface(t *testing.T) {
	cases := map[string]struct {
		input []interface{}
		want  interface{}
	}{
		"default": {
			input: []interface{}{nil, nil, nil},
			want:  nil, // default null value
		},
		"match_first": {
			input: []interface{}{1, "2", true},
			want:  1,
		},
		"match_middle": {
			input: []interface{}{nil, "2", true},
			want:  "2",
		},
		"match_last": {
			input: []interface{}{nil, nil, true},
			want:  true,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Interface(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}
