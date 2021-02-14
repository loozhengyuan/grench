package middleware

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestSetContentTypeJSON(t *testing.T) {
	cases := map[string]struct {
		headers map[string]string
		want    http.Header
	}{
		"default": {
			want: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"override_conflicting": {
			headers: map[string]string{
				"Content-Type": "something else", // will be overridden
			},
			want: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		"preserve_unaffected": {
			headers: map[string]string{
				"Easter-Egg": "OK", // will be preserved
			},
			want: http.Header{
				"Easter-Egg":   []string{"OK"},
				"Content-Type": []string{"application/json"},
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
			handler := SetContentTypeJSON(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			handler.ServeHTTP(w, req)
			resp := w.Result()

			// Assert response
			if g, w := resp.Header, tc.want; !reflect.DeepEqual(g, w) {
				t.Errorf("http headers mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}
