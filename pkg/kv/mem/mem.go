// Package mem provides a in-memory implementation of the kv.Store interface.
package mem

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/loozhengyuan/grench/pkg/kv"
)

type store struct {
	mu   sync.Mutex
	data map[string][]byte
}

var _ kv.Store = (*store)(nil)

func (s *store) PushReader(key string, r io.Reader) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return kv.ErrEmptyKey
	}
	var b bytes.Buffer
	if _, err := io.Copy(&b, r); err != nil {
		return fmt.Errorf("copy io: %w", err)
	}
	s.data[key] = b.Bytes()
	return nil
}

func (s *store) PullWriter(key string, w io.Writer) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return kv.ErrEmptyKey
	}
	v, ok := s.data[key]
	if !ok {
		return fmt.Errorf("read map: %w", kv.ErrNotFound)
	}
	b := bytes.NewReader(v)
	if _, err := io.Copy(w, b); err != nil {
		return fmt.Errorf("copy io: %w", err)
	}
	return nil
}

func (s *store) Clear(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return kv.ErrEmptyKey
	}
	delete(s.data, key)
	return nil
}

// New creates a new in-memory implementation of a key-value store.
func New() (kv.Store, error) {
	return &store{
		data: make(map[string][]byte),
	}, nil
}
