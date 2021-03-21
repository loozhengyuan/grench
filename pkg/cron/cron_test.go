package cron

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  *Expr
	}{
		// Standard 5-field pattern
		{
			name:  "default",
			input: "* * * * *",
			want: &Expr{
				minute:     elementsMinute,
				hour:       elementsHour,
				dayOfMonth: elementsDayOfMonth,
				month:      elementsMonth,
				dayOfWeek:  elementsDayOfWeek,
			},
		},
	}
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := Parse(tc.input)
			if err != nil {
				t.Fatalf("failed to parse input: %v: %s", err, tc.input)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %#v, want %#v", got, tc.want)
			}
		})
	}
}
