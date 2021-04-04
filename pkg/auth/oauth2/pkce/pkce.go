// Package pkce provides the implementation of the Proof Key for Code Exchange
// specification as defined in RFC7636.
package pkce

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

const randomSequenceOctetLength = 32

// CodeVerifier is a high-entropy, cryptographically random string that
// is used to correlate the authorization request to the token request.
type CodeVerifier struct {
	value string
}

// Value returns the raw CodeVerifier value.
func (v CodeVerifier) Value() string {
	return v.value
}

// CodeChallengeSHA256 returns the code challenge derived from the
// CodeVerifier value that is sent in the authorization request.
func (v CodeVerifier) CodeChallengeSHA256() string {
	h := sha256.New()
	h.Write([]byte(v.Value()))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// NewCodeVerifier returns a new CodeVerifier.
func NewCodeVerifier(b []byte) *CodeVerifier {
	return &CodeVerifier{
		value: base64.RawURLEncoding.EncodeToString(b),
	}
}

// NewCodeVerifierFromRandom is like NewCodeVerifier but generates
// a high-entropy, cryptographically random source value automatically.
func NewCodeVerifierFromRandom() (*CodeVerifier, error) {
	b := make([]byte, randomSequenceOctetLength)
	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("read random: %w", err)
	}
	return NewCodeVerifier(b), nil
}
