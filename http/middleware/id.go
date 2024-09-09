package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
)

// RequestID generates a request ID for the request chain
// and sets the X-Request-ID HTTP header. If the HTTP header
// already contains a value, this middleware will replace
// the HTTP header value entirely.
func RequestID(fn func() (string, error)) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if id, err := fn(); err == nil { // Only set header if no error
				w.Header().Set("X-Request-ID", id)
			}
			next.ServeHTTP(w, r)
		})
	}
}

// RequestIDRandHex16 returns a hex-encoded, 32-character string derived
// from 16 bytes of randomness. This is designed to satisfy the parameters
// the [RequestID].
func RequestIDRandHex16() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("read random: %w", err)
	}
	return hex.EncodeToString(b), nil
}
