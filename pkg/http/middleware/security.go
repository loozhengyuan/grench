package middleware

import (
	"net/http"
	"strings"

	"github.com/loozhengyuan/grench/pkg/http/header"
)

// PreventCache is a security middleware that protects sensitive data from
// being cached.
func PreventCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Advises clients not to cache/store the HTTP response.
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control
		w.Header().Set(header.CacheControl.String(), "no-cache, no-store")
		// Set response to expire immediately.
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expires
		w.Header().Set(header.Expires.String(), "0")
		// Deprecated in favour of the Cache-Control header. Preserved for
		// backwards compatibility with HTTP/1.0 clients.
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Pragma
		w.Header().Set(header.Pragma.String(), "no-cache")
		next.ServeHTTP(w, r)
	})
}

// ClickjackingProtection is a security middleware that protects against
// drag-and-drop style clickjacking attacks.
func ClickjackingProtection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevents any domain from framing the content.
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy
		w.Header().Set(header.ContentSecurityPolicy.String(), "frame-ancestors 'none'")
		// Deprecated in favour of `frame-ancestors` directive in the
		// Content-Security-Policy header.
		// NOTE: Although the X-Frame-Options header is deprecated, we still
		// preserve it to support older browsers that do not support the new
		// CSP Level 2 specification. In some browsers, the X-Frame-Options
		// header, though deprecated, may take precendence over the `frame-ancestors`
		// directive in the Content-Security-Policy header.
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
		w.Header().Set(header.XFrameOptions.String(), "DENY")
		next.ServeHTTP(w, r)
	})
}

// ContentSniffingProtection is a security middleware that prevents clients from
// doing content sniffing, which may override the value of the Content-Type header.
func ContentSniffingProtection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevent browsers from performing MIME sniffing and inappropriately
		// interpreting responses as HTML.
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options
		w.Header().Set(header.XContentTypeOptions.String(), "nosniff")
		next.ServeHTTP(w, r)
	})
}

// EnforceHSTS is a security middleware that informs clients that the site
// should only be accessed using HTTPS only, which helps to guard against
// man-in-the-middle attacks.
//
// HSTS, once enabled, cannot be easily reversed. It is therefore important
// to be clear about the impact of using HSTS before enabling it.
func EnforceHSTS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Require connections over HTTPS and protect against spoofed certificates.
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
		w.Header().Set(header.StrictTransportSecurity.String(), "max-age=63072000; includeSubDomains; preload")
		next.ServeHTTP(w, r)
	})
}

// AllowContentTypes is a security middleware that returns a HTTP 406 Not Acceptable
// response if the HTTP request presents an unset or unexpected Content-Type header.
func AllowContentTypes(types []string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ct := r.Header.Get(header.ContentType.String())
			var found bool
			for _, t := range types {
				if strings.EqualFold(t, ct) { // case-insensitive match
					found = true
				}
			}
			if !found {
				w.WriteHeader(http.StatusNotAcceptable)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// AllowHTTPMethods is a security middleware that returns a HTTP 405 Method Not Found
// response if the HTTP request method is not in the list of allowed methods.
func AllowHTTPMethods(methods []string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var found bool
			for _, m := range methods {
				if strings.EqualFold(m, r.Method) { // case-insensitive match
					found = true
				}
			}
			if !found {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
