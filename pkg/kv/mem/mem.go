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

func (s *store) Push(key string, data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return kv.ErrEmptyKey
	}
	s.data[key] = data
	return nil
}

func (s *store) PushReader(key string, r io.Reader) error {
	var b bytes.Buffer
	if _, err := io.Copy(&b, r); err != nil {
		return fmt.Errorf("copy io: %w", err)
	}
	return s.Push(key, b.Bytes())
}

func (s *store) Pull(key string) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return nil, kv.ErrEmptyKey
	}
	v, ok := s.data[key]
	if !ok {
		return nil, fmt.Errorf("read map: %w", kv.ErrNotFound)
	}
	return v, nil
}

func (s *store) PullWriter(key string, w io.Writer) error {
	v, err := s.Pull(key)
	if err != nil {
		return fmt.Errorf("pull data: %w", err)
	}
	if _, err := io.Copy(w, bytes.NewReader(v)); err != nil {
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
