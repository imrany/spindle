package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/imrany/spindle/internal/scrape"
)

// loggingMiddleware logs each HTTP request with method, path, query, status and duration.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap ResponseWriter to capture status code
		lrw := &loggingResponseWriter{w, http.StatusOK}
		next.ServeHTTP(lrw, r)

		duration := time.Since(start)
		slog.Info("HTTP request",
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"status", lrw.statusCode,
			"duration", duration,
			"remote", r.RemoteAddr,
		)
	})
}

// loggingResponseWriter captures status code for logging.
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func StartServer(addr string, port int) error {
	mux := http.NewServeMux()

	// Add scrape handler with logging
	mux.Handle("/scrape", loggingMiddleware(http.HandlerFunc(scrape.ScrapeHandler)))

	portStr := fmt.Sprintf("%s:%d", addr, port)
	slog.Info("Server starting", "addr", addr, "port", port)

	return http.ListenAndServe(portStr, mux)
}
