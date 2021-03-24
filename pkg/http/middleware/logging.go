package middleware

import (
	"net/http"
	"time"
)

type Logger interface {
	Log(metadata map[string]interface{})
}

// RequestLogging is a logging middleware that logs the lifecycle of a request.
func RequestLogging(l Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Ensure request goroutine is recovered
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					// TODO: Log something here
				}
			}()

			// Set dependencies
			start := time.Now()

			// Handle request
			next.ServeHTTP(w, r)

			// Log request
			// TODO: Header size, IP, remote IP requires additional destructuring
			// TODO: Status, Size, RequestBodySize, ResponseHeaderSize, ResponseBodySize requires request wrapping
			m := map[string]interface{}{
				"method":     r.Method,
				"url":        r.URL.String(),
				"path":       r.URL.EscapedPath(),
				"user_agent": r.UserAgent(),
				"referer":    r.Referer(),
				"proto":      r.Proto,
				"latency":    time.Since(start),
			}
			l.Log(m)
		})
	}
}
