// Package rand provides helpers for generating cryptographically-secure random values.
package rand

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func New(len int) ([]byte, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("read rand: %w", err)
	}
	return b, nil
}

func NewString(len int) (string, error) {
	b, err := New(len)
	if err != nil {
		return "", fmt.Errorf("new rand: %w", err)
	}
	return hex.EncodeToString(b), nil
}
