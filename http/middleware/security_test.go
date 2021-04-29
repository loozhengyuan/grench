package middleware

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPreventCache(t *testing.T) {
	cases := map[string]struct {
		headers map[string]string
		want    http.Header
	}{
		"default": {
			want: http.Header{
				"Cache-Control": []string{"no-cache, no-store"},
				"Expires":       []string{"0"},
				"Pragma":        []string{"no-cache"},
			},
		},
		"override_conflicting": {
			headers: map[string]string{
				"Cache-Control": "something else", // will be overridden
				"Expires":       "something else", // will be overridden
				"Pragma":        "something else", // will be overridden
			},
			want: http.Header{
				"Cache-Control": []string{"no-cache, no-store"},
				"Expires":       []string{"0"},
				"Pragma":        []string{"no-cache"},
			},
		},
		"preserve_unaffected": {
			headers: map[string]string{
				"Easter-Egg": "OK", // will be preserved
			},
			want: http.Header{
				"Easter-Egg":    []string{"OK"},
				"Cache-Control": []string{"no-cache, no-store"},
				"Expires":       []string{"0"},
				"Pragma":        []string{"no-cache"},
			},
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create request
			req, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}

			// Create recorder
			w := httptest.NewRecorder()

			// Set pre-existing headers
			for k, v := range tc.headers {
				w.Header().Set(k, v)
			}

			// Execute request
			handler := PreventCache(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(w, req)
			resp := w.Result()

			// Assert response
			if g, w := resp.Header, tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("http headers mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestClickjackingProtection(t *testing.T) {
	cases := map[string]struct {
		headers map[string]string
		want    http.Header
	}{
		"default": {
			want: http.Header{
				"Content-Security-Policy": []string{"frame-ancestors 'none'"},
				"X-Frame-Options":         []string{"DENY"},
			},
		},
		"override_conflicting": {
			headers: map[string]string{
				"Content-Security-Policy": "something else", // will be overridden
				"X-Frame-Options":         "something else", // will be overridden
			},
			want: http.Header{
				"Content-Security-Policy": []string{"frame-ancestors 'none'"},
				"X-Frame-Options":         []string{"DENY"},
			},
		},
		"preserve_unaffected": {
			headers: map[string]string{
				"Easter-Egg": "OK", // will be preserved
			},
			want: http.Header{
				"Easter-Egg":              []string{"OK"},
				"Content-Security-Policy": []string{"frame-ancestors 'none'"},
				"X-Frame-Options":         []string{"DENY"},
			},
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create request
			req, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}

			// Create recorder
			w := httptest.NewRecorder()

			// Set pre-existing headers
			for k, v := range tc.headers {
				w.Header().Set(k, v)
			}

			// Execute request
			handler := ClickjackingProtection(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(w, req)
			resp := w.Result()

			// Assert response
			if g, w := resp.Header, tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("http headers mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestContentSniffingProtection(t *testing.T) {
	cases := map[string]struct {
		headers map[string]string
		want    http.Header
	}{
		"default": {
			want: http.Header{
				"X-Content-Type-Options": []string{"nosniff"},
			},
		},
		"override_conflicting": {
			headers: map[string]string{
				"X-Content-Type-Options": "something else", // will be overridden
			},
			want: http.Header{
				"X-Content-Type-Options": []string{"nosniff"},
			},
		},
		"preserve_unaffected": {
			headers: map[string]string{
				"Easter-Egg": "OK", // will be preserved
			},
			want: http.Header{
				"Easter-Egg":             []string{"OK"},
				"X-Content-Type-Options": []string{"nosniff"},
			},
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create request
			req, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}

			// Create recorder
			w := httptest.NewRecorder()

			// Set pre-existing headers
			for k, v := range tc.headers {
				w.Header().Set(k, v)
			}

			// Execute request
			handler := ContentSniffingProtection(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(w, req)
			resp := w.Result()

			// Assert response
			if g, w := resp.Header, tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("http headers mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestEnforceHSTS(t *testing.T) {
	cases := map[string]struct {
		headers map[string]string
		want    http.Header
	}{
		"default": {
			want: http.Header{
				"Strict-Transport-Security": []string{"max-age=63072000; includeSubDomains; preload"},
			},
		},
		"override_conflicting": {
			headers: map[string]string{
				"Strict-Transport-Security": "something else", // will be overridden
			},
			want: http.Header{
				"Strict-Transport-Security": []string{"max-age=63072000; includeSubDomains; preload"},
			},
		},
		"preserve_unaffected": {
			headers: map[string]string{
				"Easter-Egg": "OK", // will be preserved
			},
			want: http.Header{
				"Easter-Egg":                []string{"OK"},
				"Strict-Transport-Security": []string{"max-age=63072000; includeSubDomains; preload"},
			},
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create request
			req, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}

			// Create recorder
			w := httptest.NewRecorder()

			// Set pre-existing headers
			for k, v := range tc.headers {
				w.Header().Set(k, v)
			}

			// Execute request
			handler := EnforceHSTS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(w, req)
			resp := w.Result()

			// Assert response
			if g, w := resp.Header, tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("http headers mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestAllowContentTypes(t *testing.T) {
	cases := map[string]struct {
		input   []string
		headers map[string]string
		want    int
	}{
		"empty_input": {
			input:   []string{},          // no input
			headers: map[string]string{}, // no headers
			want:    http.StatusNotAcceptable,
		},
		"unset_header": {
			input: []string{
				"application/json",
			},
			headers: map[string]string{},
			want:    http.StatusNotAcceptable,
		},
		"empty_header": {
			input: []string{
				"application/json",
			},
			headers: map[string]string{
				"Content-Type": "",
			},
			want: http.StatusNotAcceptable,
		},
		"no_match": {
			input: []string{
				"application/json",
			},
			headers: map[string]string{
				"Content-Type": "text/plain",
			},
			want: http.StatusNotAcceptable,
		},
		"match_single_input": {
			input: []string{
				"application/json",
			},
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			want: http.StatusOK,
		},
		"match_multiple_input": {
			input: []string{
				"application/json",
				"application/xml",
			},
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			want: http.StatusOK,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create request
			req, err := http.NewRequest(http.MethodGet, "", nil)
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}

			// Set request headers
			for k, v := range tc.headers {
				req.Header.Set(k, v)
			}

			// Create recorder
			w := httptest.NewRecorder()

			// Execute request
			handler := AllowContentTypes(tc.input)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(w, req)
			resp := w.Result()

			// Assert response
			if got := resp.StatusCode; got != tc.want {
				t.Errorf("http status code mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.want)
			}
		})
	}
}

func TestAllowHTTPMethods(t *testing.T) {
	cases := map[string]struct {
		input  []string
		method string
		want   int
	}{
		"empty_input": {
			input:  []string{}, // no input
			method: "GET",
			want:   http.StatusMethodNotAllowed,
		},
		"no_matching_methods": {
			input: []string{
				http.MethodPost,
				http.MethodDelete,
			},
			method: "GET",
			want:   http.StatusMethodNotAllowed,
		},
		"case_insensitive_match": {
			input: []string{
				http.MethodGet,
			},
			method: "get", // lowercase; should still match
			want:   http.StatusOK,
		},
		"match_single_input": {
			input: []string{
				http.MethodGet,
			},
			method: "GET",
			want:   http.StatusOK,
		},
		"match_multiple_input": {
			input: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodDelete,
			},
			method: "GET",
			want:   http.StatusOK,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create request
			req, err := http.NewRequest(tc.method, "", nil)
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}

			// Create recorder
			w := httptest.NewRecorder()

			// Execute request
			handler := AllowHTTPMethods(tc.input)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(w, req)
			resp := w.Result()

			// Assert response
			if got := resp.StatusCode; got != tc.want {
				t.Errorf("http status code mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.want)
			}
		})
	}
}
