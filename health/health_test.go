package health

import (
	"reflect"
	"testing"
)

func TestStatus_String(t *testing.T) {
	cases := map[string]struct {
		status Status
		want   string
	}{
		"status_unspecified": {
			status: StatusUnspecified,
			want:   "unspecified",
		},
		"status_pass": {
			status: StatusPass,
			want:   "pass",
		},
		"status_fail": {
			status: StatusFail,
			want:   "fail",
		},
		"out_of_range": {
			status: Status(123),
			want:   "123", // string representation of status int
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if g, w := tc.status.String(), tc.want; g != w {
				t.Errorf("output mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestService_Check(t *testing.T) {
	cases := map[string]struct {
		checks []CheckFunc
		info   []Info
	}{
		"check_single_distinct": {
			checks: []CheckFunc{
				func() Info {
					return Info{
						Component: "db",
						Status:    StatusPass,
					}
				},
			},
			info: []Info{
				{Component: "db", Status: StatusPass},
			},
		},
		"check_multiple_duplicates": {
			checks: []CheckFunc{
				func() Info {
					return Info{
						Component: "db",
						Status:    StatusPass,
					}
				},
				func() Info {
					return Info{
						Component: "db",
						Status:    StatusPass,
					}
				},
				func() Info {
					return Info{
						Component: "db",
						Status:    StatusPass,
					}
				},
			},
			info: []Info{
				{Component: "db", Status: StatusPass},
				{Component: "db", Status: StatusPass},
				{Component: "db", Status: StatusPass},
			},
		},
		"check_multiple_distinct": {
			checks: []CheckFunc{
				func() Info {
					return Info{
						Component: "db",
						Status:    StatusPass,
					}
				},
				func() Info {
					return Info{
						Component: "api",
						Status:    StatusFail,
					}
				},
			},
			info: []Info{
				{Component: "db", Status: StatusPass},
				{Component: "api", Status: StatusFail},
			},
		},
		"check_noop": {},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			chkr := New(tc.checks...)
			got := chkr.Check()
			if g, w := len(got), len(tc.info); g != w {
				t.Fatalf("slice length mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			for i := 0; i < len(tc.info); i++ {
				assertInfoEqual(t, got[i], tc.info[i])
			}
		})
	}
}

func assertInfoEqual(t *testing.T, got, want Info) {
	t.Helper()
	if g, w := got.Component, want.Component; g != w {
		t.Errorf("component mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
	}
	if g, w := got.Status, want.Status; g != w {
		t.Errorf("status mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
	}
	if g, w := got.Message, want.Message; g != w {
		t.Errorf("message mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
	}
	if g, w := got.CheckedAt, want.CheckedAt; g != w {
		t.Errorf("checked timestamp mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
	}
	if g, w := got.Metadata, want.Metadata; !reflect.DeepEqual(g, w) {
		t.Errorf("metadata mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
	}
}
