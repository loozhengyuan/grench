// Package cron provides utilities for cron expressions.
//
// The cron package only supports the standard cron expression.
package cron

import (
	"errors"
	"strings"
)

var (
	// ErrNotImplemented is returned when an expression may be valid but is not implemented.
	ErrNotImplemented = errors.New("not implemented")

	// ErrInvalidExpression is returned when an expression is invalid.
	ErrInvalidExpression = errors.New("invalid expression")
)

// Expr is the representation of a cron expression.
//
// TODO: Decide to use Expr, Expr, or Cron object.
type Expr struct {
	minute     []int
	hour       []int
	dayOfMonth []int
	month      []int
	dayOfWeek  []int
}

// Parse parses a cron expression and returns, if successful,
// an Expr object.
func Parse(expr string) (*Expr, error) {
	return parse(expr)
}

// MustParse is like Parse but panics if the expression cannot
// be parsed.
func MustParse(expr string) *Expr {
	cron, err := parse(expr)
	if err != nil {
		panic(err)
	}
	return cron
}

// TODO: 6th field
// TODO: i-th slash notation
func parse(expr string) (*Expr, error) {
	// Handle non-standard macros
	// TODO: @reboot, @hourly
	switch expr {
	case "@yearly", "@annually":
		expr = "0 0 1 1 *"
	case "@monthly":
		expr = "0 0 1 * *"
	case "@weekly":
		expr = "0 0 * * 0"
	case "@daily", "midnight":
		expr = "0 0 * * *"
	case "@hourly":
		expr = "0 * * * *"
	}
	elems := strings.Split(expr, " ")
	if len(elems) != 5 {
		return nil, ErrInvalidExpression
	}
	cron := &Expr{
		minute:     parseMinute(elems[0]),
		hour:       parseHour(elems[1]),
		dayOfMonth: parseDayOfMonth(elems[2]),
		month:      parseMonth(elems[3]),
		dayOfWeek:  parseDayOfWeek(elems[4]),
	}
	return cron, nil
}

var elementsMinute = []int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 22, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
}

var elementsHour = []int{
	0, 1, 2, 3, 4, 5,
	6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
	18, 19, 20, 21, 22, 23,
}

var elementsDayOfMonth = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
}

var elementsMonth = []int{
	1, 2, 3, 4, 5, 6,
	7, 8, 9, 10, 11, 12,
}

var elementsDayOfWeek = []int{
	0, 1, 2, 3, 4, 5, 6,
}

func parseMinute(expr string) []int {
	// TODO: custom range, selective elements
	switch expr {
	case "*", "0-59":
		return elementsMinute
	}
	return nil
}

func parseHour(expr string) []int {
	// TODO: custom range, selective elements
	switch expr {
	case "*", "0-23":
		return elementsHour
	}
	return nil
}

func parseDayOfMonth(expr string) []int {
	// TODO: custom range, selective elements
	switch expr {
	case "*", "1-31":
		return elementsDayOfMonth
	}
	return nil
}

func parseMonth(expr string) []int {
	// TODO: custom range, selective elements
	switch expr {
	case "*", "1-12":
		return elementsMonth
	}
	return nil
}

func parseDayOfWeek(expr string) []int {
	// TODO: custom range, selective elements
	switch expr {
	case "*", "0-6":
		return elementsDayOfWeek
	}
	return nil
}
