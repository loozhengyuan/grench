package build

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestInfo_OutputText(t *testing.T) {
	cases := map[string]struct {
		info Info
		want string
	}{
		"default": {
			info: Info{
				App:       "myapp",
				System:    "linux",
				Arch:      "amd64",
				Version:   "v0.0.0",
				Commit:    "dev",
				Timestamp: "1970-01-01T00:00:00Z",
			},
			want: "myapp v0.0.0 linux/amd64 1970-01-01T00:00:00Z dev\n",
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var buf bytes.Buffer
			if err := tc.info.OutputText(&buf); err != nil {
				t.Fatalf("failed to invoke function: %v", err)
			}
			if g, w := buf.String(), tc.want; g != w {
				t.Errorf("output mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestInfo_OutputJSON(t *testing.T) {
	cases := map[string]struct {
		info Info
	}{
		"default": {
			info: Info{
				App:       "myapp",
				System:    "linux",
				Arch:      "amd64",
				Version:   "v0.0.0",
				Commit:    "dev",
				Timestamp: "1970-01-01T00:00:00Z",
			},
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var buf bytes.Buffer
			if err := tc.info.OutputJSON(&buf); err != nil {
				t.Fatalf("failed to invoke function: %v", err)
			}
			var got Info
			if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}
			if g, w := got.App, tc.info.App; g != w {
				t.Errorf("app mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := got.System, tc.info.System; g != w {
				t.Errorf("system mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := got.Arch, tc.info.Arch; g != w {
				t.Errorf("arch mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := got.Version, tc.info.Version; g != w {
				t.Errorf("version mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := got.Commit, tc.info.Commit; g != w {
				t.Errorf("commit mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := got.Timestamp, tc.info.Timestamp; g != w {
				t.Errorf("timestamp mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}
