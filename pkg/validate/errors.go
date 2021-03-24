package validate

import (
	"strings"
)

// ValidationError is returned when a struct fails any validation
// conditions.
type ValidationError struct {
	Errs []string
}

// Ensures that ValidationError implements the error interface.
var _ error = (*ValidationError)(nil)

// Error returns the string representation of all validation errors.
func (e ValidationError) Error() string {
	return "\n" + strings.Join(e.Errs, "\n")
}
