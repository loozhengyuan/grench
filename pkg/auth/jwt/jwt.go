// Package jwt provides the implementation of the JSON Web Token specification
// as defined in RFC7519.
package jwt

// Claims describes the standard claims defined in RFC7519.
type Claims struct {
	Issuer         string `json:"iss,omitempty"`
	Subject        string `json:"sub,omitempty"`
	Audience       string `json:"aud,omitempty"`
	ExpirationTime string `json:"exp,omitempty"`
	NotBefore      string `json:"nbf,omitempty"`
	IssuedAt       string `json:"iat,omitempty"`
	JWTID          string `json:"jti,omitempty"`
}
