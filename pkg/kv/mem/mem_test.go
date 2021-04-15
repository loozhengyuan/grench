package mem

import (
	"testing"

	"github.com/loozhengyuan/grench/pkg/kv"
)

func TestStorePushPull(t *testing.T) {
	kv.TestStorePushPull(t, func() (kv.Store, error) {
		return New()
	})
}

func TestStorePushError(t *testing.T) {
	kv.TestStorePushError(t, func() (kv.Store, error) {
		return New()
	})
}

func TestStorePullError(t *testing.T) {
	kv.TestStorePullError(t, func() (kv.Store, error) {
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
