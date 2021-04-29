package kv

import (
	"errors"
)

var (
	// ErrEmptyKey is returned when the key is empty.
	ErrEmptyKey = errors.New("empty key")

	// ErrNotFound is returned when the key does not exist on the key-value store.
	ErrNotFound = errors.New("not found")
)
