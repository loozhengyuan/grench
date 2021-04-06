package oidc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	// ErrDiscoveryEndpointInvalidContentType is returned when the discovery endpoint
	// returns a HTTP response with an invalid Content-Type header value.
	ErrDiscoveryEndpointInvalidContentType = errors.New("discovery endpoint: invalid content type")

	// ErrDiscoveryEndpointInvalidStatusCode is returned when the discovery endpoint
	// returns a HTTP response with an invalid HTTP status code.
	ErrDiscoveryEndpointInvalidStatusCode = errors.New("discovery endpoint: invalid status code")
)

// ProviderMetadata represents the configuration metadata of an OpenID Provider (OP).
//
// https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata
type ProviderMetadata struct {
	Issuer                                     string   `json:"issuer,omitempty"`
	AuthorizationEndpoint                      string   `json:"authorization_endpoint,omitempty"`
	TokenEndpoint                              string   `json:"token_endpoint,omitempty"`
	UserInfoEndpoint                           string   `json:"userinfo_endpoint,omitempty"`
	JWKSURI                                    string   `json:"jwks_uri,omitempty"`
	RegistrationEndpoint                       string   `json:"registration_endpoint,omitempty"`
	ScopesSupported                            []string `json:"scopes_supported,omitempty"`
	ResponseTypesSupported                     []string `json:"response_types_supported,omitempty"`
	ResponseModesSupported                     []string `json:"response_modes_supported,omitempty"`
	GrantTypesSupported                        []string `json:"grant_types_supported,omitempty"`
	ACRValuesSupported                         []string `json:"acr_values_supported,omitempty"`
	SubjectTypesSupported                      []string `json:"subject_types_supported,omitempty"`
	IDTokenSigningAlgValuesSupported           []string `json:"id_token_signing_alg_values_supported,omitempty"`
	IDTokenEncryptionAlgValuesSupported        []string `json:"id_token_encryption_alg_values_supported,omitempty"`
	IDTokenEncryptionEncValuesSupported        []string `json:"id_token_encryption_enc_values_supported,omitempty"`
	UserInfoSigningAlgValuesSupported          []string `json:"userinfo_signing_alg_values_supported,omitempty"`
	UserInfoEncryptionAlgValuesSupported       []string `json:"userinfo_encryption_alg_values_supported,omitempty"`
	UserInfoEncryptionEncValuesSupported       []string `json:"userinfo_encryption_enc_values_supported,omitempty"`
	RequestObjectSigningAlgValuesSupported     []string `json:"request_object_signing_alg_values_supported,omitempty"`
	RequestObjectEncryptionAlgValuesSupported  []string `json:"request_object_encryption_alg_values_supported,omitempty"`
	RequestObjectEncryptionEncValuesSupported  []string `json:"request_object_encryption_enc_values_supported,omitempty"`
	TokenEndpointAuthMethodsSupported          []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	TokenEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported,omitempty"`
	DisplayValuesSupported                     []string `json:"display_values_supported,omitempty"`
	ClaimTypesSupported                        []string `json:"claim_types_supported,omitempty"`
	ClaimsSupported                            []string `json:"claims_supported,omitempty"`
	ServiceDocumentation                       string   `json:"service_documentation,omitempty"`
	ClaimsLocalesSupported                     []string `json:"claims_locales_supported,omitempty"`
	UILocalesSupported                         []string `json:"ui_locales_supported,omitempty"`
	ClaimsParameterSupported                   bool     `json:"claims_parameter_supported,omitempty"`
	RequestParameterSupported                  bool     `json:"request_parameter_supported,omitempty"`
	RequestURIParameterSupported               bool     `json:"request_uri_parameter_supported,omitempty"`
	RequireRequestURIRegistration              bool     `json:"require_request_uri_registration,omitempty"`
	OPPolicyURI                                string   `json:"op_policy_uri,omitempty"`
	OPTOSURI                                   string   `json:"op_tos_uri,omitempty"`
}

// Provider represents an OpenID Provider (OP).
type Provider struct {
	Config ProviderMetadata
}

// NewProviderFromJSON returns a new Provider from JSON data.
func NewProviderFromJSON(r io.Reader) (*Provider, error) {
	var m ProviderMetadata
	if err := json.NewDecoder(r).Decode(&m); err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}
	return &Provider{Config: m}, nil
}

// NewProviderFromDiscoveryEndpoint returns new Provider from a discovery endpoint.
func NewProviderFromDiscoveryEndpoint(url string) (*Provider, error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http request: get %s: %w", url, err)
	}
	defer resp.Body.Close()
	if v := resp.Header.Get("Content-Type"); v != "application/json" {
		return nil, fmt.Errorf("%w: got %s", ErrDiscoveryEndpointInvalidContentType, v)
	}
	if v := resp.StatusCode; v != http.StatusOK {
		return nil, fmt.Errorf("%w: got %d", ErrDiscoveryEndpointInvalidStatusCode, v)
	}
	return NewProviderFromJSON(resp.Body)
}

// NewProviderFromIssuer returns a new Provider from an issuer.
func NewProviderFromIssuer(issuer string) (*Provider, error) {
	u, err := url.Parse(issuer + "/.well-known/openid-configuration")
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}
	return NewProviderFromDiscoveryEndpoint(u.String())
}
