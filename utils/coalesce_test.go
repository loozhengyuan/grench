package utils

import (
	"testing"
)

func TestCoalesce_Int(t *testing.T) {
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
			got := Coalesce(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func BenchmarkCoalesce_Int(b *testing.B) {
	benchmarks := map[string]struct {
		input []int
	}{
		"best": {
			input: []int{1, 0, 0},
		},
		"worst": {
			input: []int{0, 0, 0},
		},
	}
	for name, bm := range benchmarks {
		bm := bm // capture range variable
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Coalesce(bm.input...)
			}
		})
	}
}

func TestCoalesce_Int32(t *testing.T) {
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
			got := Coalesce(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func BenchmarkCoalesce_Int32(b *testing.B) {
	benchmarks := map[string]struct {
		input []int32
	}{
		"best": {
			input: []int32{1, 0, 0},
		},
		"worst": {
			input: []int32{0, 0, 0},
		},
	}
	for name, bm := range benchmarks {
		bm := bm // capture range variable
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Coalesce(bm.input...)
			}
		})
	}
}

func TestCoalesce_Int64(t *testing.T) {
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
			got := Coalesce(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func BenchmarkCoalesce_Int64(b *testing.B) {
	benchmarks := map[string]struct {
		input []int64
	}{
		"best": {
			input: []int64{1, 0, 0},
		},
		"worst": {
			input: []int64{0, 0, 0},
		},
	}
	for name, bm := range benchmarks {
		bm := bm // capture range variable
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Coalesce(bm.input...)
			}
		})
	}
}

func TestCoalesce_Float32(t *testing.T) {
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
			got := Coalesce(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func BenchmarkCoalesce_Float32(b *testing.B) {
	benchmarks := map[string]struct {
		input []float32
	}{
		"best": {
			input: []float32{1, 0, 0},
		},
		"worst": {
			input: []float32{0, 0, 0},
		},
	}
	for name, bm := range benchmarks {
		bm := bm // capture range variable
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Coalesce(bm.input...)
			}
		})
	}
}

func TestCoalesce_Float64(t *testing.T) {
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
			got := Coalesce(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func BenchmarkCoalesce_Float64(b *testing.B) {
	benchmarks := map[string]struct {
		input []float64
	}{
		"best": {
			input: []float64{1, 0, 0},
		},
		"worst": {
			input: []float64{0, 0, 0},
		},
	}
	for name, bm := range benchmarks {
		bm := bm // capture range variable
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Coalesce(bm.input...)
			}
		})
	}
}

func TestCoalesce_String(t *testing.T) {
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
			got := Coalesce(tc.input...)
			if got != tc.want {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
			}
		})
	}
}

func BenchmarkCoalesce_String(b *testing.B) {
	benchmarks := map[string]struct {
		input []string
	}{
		"best": {
			input: []string{"1", "", ""},
		},
		"worst": {
			input: []string{"", "", ""},
		},
	}
	for name, bm := range benchmarks {
		bm := bm // capture range variable
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Coalesce(bm.input...)
			}
		})
	}
}

// TODO: `interface{}` can only satisfy comparable in Go 1.20
// func TestCoalesce_Interface(t *testing.T) {
// 	cases := map[string]struct {
// 		input []interface{}
// 		want  interface{}
// 	}{
// 		"default": {
// 			input: []interface{}{nil, nil, nil},
// 			want:  nil, // default null value
// 		},
// 		"match_first": {
// 			input: []interface{}{1, "2", true},
// 			want:  1,
// 		},
// 		"match_middle": {
// 			input: []interface{}{nil, "2", true},
// 			want:  "2",
// 		},
// 		"match_last": {
// 			input: []interface{}{nil, nil, true},
// 			want:  true,
// 		},
// 	}
// 	for name, tc := range cases {
// 		tc := tc // capture range variable
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()
// 			got := Coalesce(tc.input...)
// 			if got != tc.want {
// 				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.input)
// 			}
// 		})
// 	}
// }

// TODO: `interface{}` can only satisfy comparable in Go 1.20
// func BenchmarkCoalesce_Interface(b *testing.B) {
// 	benchmarks := map[string]struct {
// 		input []interface{}
// 	}{
// 		"best": {
// 			input: []interface{}{1, nil, nil},
// 		},
// 		"worst": {
// 			input: []interface{}{nil, nil, nil},
// 		},
// 	}
// 	for name, bm := range benchmarks {
// 		bm := bm // capture range variable
// 		b.Run(name, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				Coalesce(bm.input...)
// 			}
// 		})
// 	}
// }
