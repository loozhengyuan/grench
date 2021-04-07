// Package kv provides interface for interacting with a key-value store.
package kv

import (
	"io"
)

// Store is the interface for interacting with an key-value store.
type Store interface {
	// Push uploads data from a io.Reader object to the key-value store.
	Push(key string, r io.Reader) error

	// Pull downloads data from the store to a io.Writer object.
	Pull(key string, w io.Writer) error

	// Clear removes data associated with a given key.
	Clear(key string) error
}
