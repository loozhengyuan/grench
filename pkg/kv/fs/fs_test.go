package fs

import (
	"testing"

	"github.com/loozhengyuan/grench/pkg/kv"
)

func TestStorePushPull(t *testing.T) {
	kv.TestStorePushPull(t, func() (kv.Store, error) {
		return NewTemp()
	})
}

func TestStorePushReaderPullWriter(t *testing.T) {
	kv.TestStorePushReaderPullWriter(t, func() (kv.Store, error) {
		return NewTemp()
	})
}

func TestStorePushError(t *testing.T) {
	kv.TestStorePushError(t, func() (kv.Store, error) {
		return NewTemp()
	})
}

func TestStorePushReaderError(t *testing.T) {
	kv.TestStorePushReaderError(t, func() (kv.Store, error) {
		return NewTemp()
	})
}

func TestStorePullError(t *testing.T) {
	kv.TestStorePullError(t, func() (kv.Store, error) {
		return NewTemp()
	})
}

func TestStorePullWriterError(t *testing.T) {
	kv.TestStorePullWriterError(t, func() (kv.Store, error) {
		return NewTemp()
	})
}

func TestStoreClear(t *testing.T) {
	kv.TestStoreClear(t, func() (kv.Store, error) {
		return NewTemp()
	})
}
func TestStoreClearError(t *testing.T) {
	kv.TestStoreClearError(t, func() (kv.Store, error) {
		return NewTemp()
	})
}
