// Package fs provides a local filesystem implementation of the kv.Store interface.
package fs

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/loozhengyuan/grench/pkg/kv"
)

type store struct {
	mu   sync.Mutex
	path string
}

var _ kv.Store = (*store)(nil)

func (s *store) Push(key string, r io.Reader) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return kv.ErrEmptyKey
	}
	f, err := os.Create(filepath.Join(s.path, key))
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()
	if _, err := io.Copy(f, r); err != nil {
		return fmt.Errorf("copy io: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close file: %w", err)
	}
	return nil
}

func (s *store) Pull(key string, w io.Writer) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return kv.ErrEmptyKey
	}
	f, err := os.Open(filepath.Join(s.path, key))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("open file: %w", kv.ErrNotFound)
		}
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		return fmt.Errorf("copy io: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close file: %w", err)
	}
	return nil
}

func (s *store) Clear(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return kv.ErrEmptyKey
	}
	if err := os.RemoveAll(filepath.Join(s.path, key)); err != nil {
		return fmt.Errorf("remove file: %w", err)
	}
	return nil
}

// New returns a new local filesystem implementation of a key-value store.
func New(path string) (kv.Store, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("stat file: %w", err)
	}
	return &store{path: path}, nil
}

// NewTemp is like New but creates a temporary directory automatically.
func NewTemp() (kv.Store, error) {
	dir, err := ioutil.TempDir(os.TempDir(), "*")
	if err != nil {
		return nil, fmt.Errorf("new temp dir: %w", err)
	}
	return New(dir)
}
