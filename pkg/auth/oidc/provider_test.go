package oidc

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNewProviderFromJSON(t *testing.T) {
	cases := map[string]struct {
		data string
		want *Provider
	}{
		"field_issuer": {
			data: `{"issuer": "abc"}`,
			want: &Provider{
				Config: ProviderMetadata{
					Issuer: "abc",
				},
			},
		},
		"field_authorization_endpoint": {
			data: `{"authorization_endpoint": "abc"}`,
			want: &Provider{
				Config: ProviderMetadata{
					AuthorizationEndpoint: "abc",
				},
			},
		},
		"field_token_endpoint": {
			data: `{"token_endpoint": "abc"}`,
			want: &Provider{
				Config: ProviderMetadata{
					TokenEndpoint: "abc",
				},
			},
		},
		"field_userinfo_endpoint": {
			data: `{"userinfo_endpoint": "abc"}`,
			want: &Provider{
				Config: ProviderMetadata{
					UserInfoEndpoint: "abc",
				},
			},
		},
		"field_jwks_uri": {
			data: `{"jwks_uri": "abc"}`,
			want: &Provider{
				Config: ProviderMetadata{
					JWKSURI: "abc",
				},
			},
		},
		"field_registration_endpoint": {
			data: `{"registration_endpoint": "abc"}`,
			want: &Provider{
				Config: ProviderMetadata{
					RegistrationEndpoint: "abc",
				},
			},
		},
		"field_scopes_supported": {
			data: `{"scopes_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					ScopesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_response_types_supported": {
			data: `{"response_types_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					ResponseTypesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_response_modes_supported": {
			data: `{"response_modes_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					ResponseModesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_grant_types_supported": {
			data: `{"grant_types_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					GrantTypesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_acr_values_supported": {
			data: `{"acr_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					ACRValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_subject_types_supported": {
			data: `{"subject_types_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					SubjectTypesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_id_token_signing_alg_values_supported": {
			data: `{"id_token_signing_alg_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					IDTokenSigningAlgValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_id_token_encryption_alg_values_supported": {
			data: `{"id_token_encryption_alg_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					IDTokenEncryptionAlgValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_id_token_encryption_enc_values_supported": {
			data: `{"id_token_encryption_enc_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					IDTokenEncryptionEncValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_userinfo_signing_alg_values_supported": {
			data: `{"userinfo_signing_alg_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					UserInfoSigningAlgValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_userinfo_encryption_alg_values_supported": {
			data: `{"userinfo_encryption_alg_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					UserInfoEncryptionAlgValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_userinfo_encryption_enc_values_supported": {
			data: `{"userinfo_encryption_enc_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					UserInfoEncryptionEncValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_request_object_signing_alg_values_supported": {
			data: `{"request_object_signing_alg_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					RequestObjectSigningAlgValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_request_object_encryption_alg_values_supported": {
			data: `{"request_object_encryption_alg_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					RequestObjectEncryptionAlgValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_request_object_encryption_enc_values_supported": {
			data: `{"request_object_encryption_enc_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					RequestObjectEncryptionEncValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_token_endpoint_auth_methods_supported": {
			data: `{"token_endpoint_auth_methods_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					TokenEndpointAuthMethodsSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_token_endpoint_auth_signing_alg_values_supported": {
			data: `{"token_endpoint_auth_signing_alg_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					TokenEndpointAuthSigningAlgValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_display_values_supported": {
			data: `{"display_values_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					DisplayValuesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_claim_types_supported": {
			data: `{"claim_types_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					ClaimTypesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_claims_supported": {
			data: `{"claims_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					ClaimsSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_service_documentation": {
			data: `{"service_documentation": "abc"}`,
			want: &Provider{
				Config: ProviderMetadata{
					ServiceDocumentation: "abc",
				},
			},
		},
		"field_claims_locales_supported": {
			data: `{"claims_locales_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					ClaimsLocalesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_ui_locales_supported": {
			data: `{"ui_locales_supported": ["a", "b", "c"]}`,
			want: &Provider{
				Config: ProviderMetadata{
					UILocalesSupported: []string{"a", "b", "c"},
				},
			},
		},
		"field_claims_parameter_supported": {
			data: `{"claims_parameter_supported": true}`,
			want: &Provider{
				Config: ProviderMetadata{
					ClaimsParameterSupported: true,
				},
			},
		},
		"field_request_parameter_supported": {
			data: `{"request_parameter_supported": true}`,
			want: &Provider{
				Config: ProviderMetadata{
					RequestParameterSupported: true,
				},
			},
		},
		"field_request_uri_parameter_supported": {
			data: `{"request_uri_parameter_supported": true}`,
			want: &Provider{
				Config: ProviderMetadata{
					RequestURIParameterSupported: true,
				},
			},
		},
		"field_require_request_uri_registration": {
			data: `{"require_request_uri_registration": true}`,
			want: &Provider{
				Config: ProviderMetadata{
					RequireRequestURIRegistration: true,
				},
			},
		},
		"field_op_policy_uri": {
			data: `{"op_policy_uri": "abc"}`,
			want: &Provider{
				Config: ProviderMetadata{
					OPPolicyURI: "abc",
				},
			},
		},
		"field_op_tos_uri": {
			data: `{"op_tos_uri": "abc"}`,
			want: &Provider{
				Config: ProviderMetadata{
					OPTOSURI: "abc",
				},
			},
		},
		// TODO: JSON decode errors
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(tc.data)
			p, err := NewProviderFromJSON(r)
			if err != nil {
				t.Fatalf("failed to invoke function: %v", err)
			}
			if g, w := p, tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("provider mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestNewProviderFromDiscoveryEndpoint(t *testing.T) {
	cases := map[string]struct {
		httpStatus  int
		httpHeaders map[string]string
		data        string
		want        *Provider
		err         error
	}{
		"default": {
			httpStatus:  http.StatusOK,
			httpHeaders: map[string]string{"Content-Type": "application/json"},
			data:        `{"issuer": "abc"}`,
			want: &Provider{
				Config: ProviderMetadata{
					Issuer: "abc",
				},
			},
		},
		"error_invalid_content_type": {
			httpStatus:  http.StatusOK,
			httpHeaders: map[string]string{"Content-Type": "text/plain"}, // invalid
			err:         ErrDiscoveryEndpointInvalidContentType,
		},
		"error_invalid_status_code": {
			httpStatus:  http.StatusBadRequest, // invalid
			httpHeaders: map[string]string{"Content-Type": "application/json"},
			err:         ErrDiscoveryEndpointInvalidStatusCode,
		},
		// TODO: HTTP request errors
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, v := range tc.httpHeaders {
					w.Header().Set(k, v)
				}
				w.WriteHeader(tc.httpStatus)
				w.Write([]byte(tc.data))
			}))
			p, err := NewProviderFromDiscoveryEndpoint(srv.URL)
			if !errors.Is(err, tc.err) {
				t.Fatalf("error value mismatch:\ngot:\t%#v\nwant:\t%#v", errors.Unwrap(err), tc.err)
			}
			if g, w := p, tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("provider mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}
