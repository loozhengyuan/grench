package oidc

import (
	"github.com/loozhengyuan/grench/auth/oauth2"
)

// TokenResponse represents the response type when querying
// the token endpoint.
type TokenResponse struct {
	oauth2.TokenResponse
	IDToken string `json:"id_token,omitempty"`
}
