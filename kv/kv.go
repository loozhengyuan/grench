// Package kv provides interface for interacting with a key-value store.
package kv

import (
	"io"
)

// Store is the interface for interacting with an key-value store.
type Store interface {
	// Push pushes data to the key-value store.
	Push(key string, data []byte) error

	// PushReader is like Push but reads the data from r.
	PushReader(key string, r io.Reader) error

	// Pull pulls data from the the key-value store.
	Pull(key string) ([]byte, error)

	// PullWriter is like Pull but writes the data to r.
	PullWriter(key string, w io.Writer) error

	// Clear removes data associated with a given key.
	Clear(key string) error
}
