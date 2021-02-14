// Package middleware provides http middleware.
package middleware

import (
	"net/http"
)

// Middleware is the function signature of a HTTP middleware.
type Middleware func(http.Handler) http.Handler
