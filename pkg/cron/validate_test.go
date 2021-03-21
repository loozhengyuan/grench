package cron

import (
	"testing"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		// Standard 5-field pattern
		{input: "* * * * *", want: true},

		// Minute Field
		// Allowed values: 0-59
		// Allowed characters: *,-
		{input: "0 * * * *", want: true},
		{input: "0,59 * * * *", want: true},
		{input: "0-59 * * * *", want: true},

		// Hour Field
		// Allowed values: 0-23
		// Allowed characters: *,-
		{input: "* 0 * * *", want: true},
		{input: "* 0,23 * * *", want: true},
		{input: "* 0-23 * * *", want: true},

		// DayOfMonth Field
		// Allowed values: 1-31
		// Allowed characters: *,-
		{input: "* * 1 * *", want: true},
		{input: "* * 1,31 * *", want: true},
		{input: "* * 1-31 * *", want: true},

		// Month Field
		// Allowed values: 1-12
		// Allowed characters: *,-
		{input: "* * * 0 *", want: true},
		{input: "* * * 1,12 *", want: true},
		{input: "* * * 1-12 *", want: true},

		// DayOfWeek Field
		// Allowed values: 0-6
		// Allowed characters: *,-
		{input: "* * * * 0", want: true},
		{input: "* * * * 0,6", want: true},
		{input: "* * * * 0-6", want: true},

		// Non-standard macros
		{input: "@yearly", want: true},
		{input: "@annually", want: true},
		{input: "@monthly", want: true},
		{input: "@weekly", want: true},
		{input: "@daily", want: true},
		{input: "@midnight", want: true},
		{input: "@hourly", want: true},
		{input: "@reboot", want: true}, // Not applicable
	}
	for _, tc := range cases {
		tc := tc // capture range variable
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()
			got := Validate(tc.input)
			if got != tc.want {
				t.Errorf("got %t want %t", got, tc.want)
			}
		})
	}
}
