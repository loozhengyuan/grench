package fs

import (
	"testing"

	"github.com/loozhengyuan/grench/pkg/kv"
)

func TestStorePushPullSequence(t *testing.T) {
	kv.TestStorePushPullSequence(t, func() (kv.Store, error) {
		return NewTemp()
	})
}

func TestStoreClear(t *testing.T) {
	kv.TestStoreClear(t, func() (kv.Store, error) {
		return NewTemp()
	})
}
