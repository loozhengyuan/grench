package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRequestID(t *testing.T) {
	cases := map[string]struct {
		input   func() (string, error)
		headers map[string]string
		want    http.Header
	}{
		"default": {
			input: func() (string, error) { return "xxx", nil },
			want: http.Header{
				"X-Request-Id": []string{"xxx"},
			},
		},
		"override_conflicting": {
			input: func() (string, error) { return "xxx", nil },
			headers: map[string]string{
				"X-Request-Id": "something else", // will be overridden
			},
			want: http.Header{
				"X-Request-Id": []string{"xxx"},
			},
		},
		"preserve_unaffected": {
			input: func() (string, error) { return "xxx", nil },
			headers: map[string]string{
				"Easter-Egg": "OK", // will be preserved
			},
			want: http.Header{
				"Easter-Egg":   []string{"OK"},
				"X-Request-Id": []string{"xxx"},
			},
		},
		"skip_on_error": {
			input: func() (string, error) { return "", errors.New("xxx") },
			want:  http.Header{}, // NOTE: Header will not be set on failure!
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
			handler := RequestID(tc.input)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(w, req)
			resp := w.Result()

			// Assert response
			if g, w := resp.Header, tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("http headers mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}
