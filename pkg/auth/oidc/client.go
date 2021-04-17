package oidc

import (
	"fmt"
	"net/http"
	"net/url"
)

// RequestParameter is a request parameter used in the authentication request.
type RequestParameter string

// Enum types for RequestParameter.
const (
	RequestParameterUnspecified RequestParameter = "" // Default null value
	RequestParameterScope
)

type RequestParameterOption func(*url.Values)

func WithScope(s string) RequestParameterOption {
	return func(v *url.Values) {
		v.Set("scope", s)
	}
}

func WithResponseType(s string) RequestParameterOption {
	return func(v *url.Values) {
		v.Set("response_type", s)
	}
}

func WithClientID(s string) RequestParameterOption {
	return func(v *url.Values) {
		v.Set("client_id", s)
	}
}

func WithRedirectURI(s string) RequestParameterOption {
	return func(v *url.Values) {
		v.Set("redirect_uri", s)
	}
}

// Client represents an OpenID relying party (RP).
type Client struct {
	client *http.Client
}

// AuthURL returns the authorization url.
func (c Client) AuthURL(opts ...RequestParameterOption) string {
	v := &url.Values{}
	for _, opt := range opts {
		opt(v)
	}
	return ""
}

// AuthURLForAuthCodeFlow returns the Authorization URL for the
// Authorization Code Flow.
func (c Client) AuthURLForAuthCodeFlow(opts ...RequestParameterOption) string {
	// REQUIRED: client_id
	// REQUIRED: response_type=code
	// REQUIRED: scope=openid
	// REQUIRED: redirect_uri
	// REQUIRED: state
	return ""
}

// StartAuthorizationCodeFlow starts the Authorization Code flow.
func (c *Client) StartAuthorizationCodeFlow() error {
	// Get authorization url
	authURL := c.AuthURL(
		WithScope("openid"),
		WithResponseType("code"),
		WithClientID(""), // TODO
		WithRedirectURI("http://localhost:8080/auth/callback"), // TODO
	)

	// Get authorization code
	resp, err := c.client.Get(authURL)
	if err != nil {
		return fmt.Errorf("http request: get %s: %w", authURL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http response not ok: got %d %s", resp.StatusCode, resp.Status)
	}
	return nil
}
