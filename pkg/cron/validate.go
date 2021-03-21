package cron

import (
	"regexp"
)

var cronRegExp = regexp.MustCompile(`^(@(annually|yearly|monthly|weekly|daily|midnight|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7})$`)

// Validate validates a cron expression.
func Validate(str string) bool {
	return cronRegExp.MatchString(str)
}
