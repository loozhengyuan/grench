package middleware

import (
	"net/http"

	"github.com/loozhengyuan/grench/pkg/http/header"
	"github.com/loozhengyuan/grench/pkg/http/mime"
)

// SetContentTypeJSON is a middleware that sets the content type of the
// response as JSON.
func SetContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type
		w.Header().Set(header.ContentType.String(), mime.ApplicationJSON.String())
		next.ServeHTTP(w, r)
	})
}
