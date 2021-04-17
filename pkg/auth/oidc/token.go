package oidc

import (
	"time"

	"github.com/loozhengyuan/grench/pkg/auth/oauth2"
)

// TokenResponse represents the response type when querying
// the token endpoint.
type TokenResponse struct {
	oauth2.TokenResponse
	IDToken string `json:"id_token,omitempty"`
}

type IDToken struct {
	Issuer              string    `json:"iss,omitempty"`
	Subject             string    `json:"sub,omitempty"`
	Audience            string    `json:"aud,omitempty"`
	ExpiresAt           time.Time `json:"exp,omitempty"`
	IssuedAt            int       `json:"iat,omitempty"`       // TODO: Use time.Time?
	AuthTime            int       `json:"auth_time,omitempty"` // TODO: Use time.Time?
	Nonce               string    `json:"nonce,omitempty"`
	AuthClassReference  string    `json:"acr,omitempty"`
	AuthMethodReference string    `json:"amr,omitempty"`
	AuthorizedParty     string    `json:"azp,omitempty"`
}
