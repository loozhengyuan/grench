package header

import (
	"testing"
)

func TestHeader_String(t *testing.T) {
	cases := map[string]struct {
		header Header
		want   string
	}{
		"accept": {header: Accept, want: "Accept"},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := tc.header.String()
			if g, w := got, tc.want; g != w {
				t.Errorf("string mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}
