package cron

import (
	"time"
)

// Iterator provides an iterator wrapper for an Expr.
type Iterator struct {
	expr Expr
	curr time.Time
}

// IteratorOption represents the functional option for Iterator.
type IteratorOption func(*Iterator)

// WithTime sets the time for an Iterator.
func WithTime(t time.Time) IteratorOption {
	return func(i *Iterator) {
		i.curr = t
	}
}

// NewIterator returns a new cron.Iterator object.
func NewIterator(expr Expr, options ...IteratorOption) Iterator {
	iter := Iterator{
		expr: expr,
		curr: time.Now(),
	}
	for _, option := range options {
		option(&iter)
	}
	return iter
}

// Prev returns the previous run time.
func (i *Iterator) Prev() time.Time {
	// TODO: Generate previous from i.expr
	return i.curr
}

// Curr returns the current run time.
func (i *Iterator) Curr() time.Time {
	return i.curr
}

// Next returns the next run time.
func (i *Iterator) Next() time.Time {
	// TODO: Generate next from i.expr
	return i.curr
}
