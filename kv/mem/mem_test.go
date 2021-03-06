package mem

import (
	"testing"

	"github.com/loozhengyuan/grench/kv"
)

func TestStorePushPull(t *testing.T) {
	kv.TestStorePushPull(t, func() (kv.Store, error) {
		return New()
	})
}

func TestStorePushReaderPullWriter(t *testing.T) {
	kv.TestStorePushReaderPullWriter(t, func() (kv.Store, error) {
		return New()
	})
}

func TestStorePushError(t *testing.T) {
	kv.TestStorePushError(t, func() (kv.Store, error) {
		return New()
	})
}

func TestStorePushReaderError(t *testing.T) {
	kv.TestStorePushReaderError(t, func() (kv.Store, error) {
		return New()
	})
}

func TestStorePullError(t *testing.T) {
	kv.TestStorePullError(t, func() (kv.Store, error) {
		return New()
	})
}

func TestStorePullWriterError(t *testing.T) {
	kv.TestStorePullWriterError(t, func() (kv.Store, error) {
		return New()
	})
}

func TestStoreClear(t *testing.T) {
	kv.TestStoreClear(t, func() (kv.Store, error) {
		return New()
	})
}
func TestStoreClearError(t *testing.T) {
	kv.TestStoreClearError(t, func() (kv.Store, error) {
		return New()
	})
}
