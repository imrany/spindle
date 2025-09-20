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

// corsMiddleware adds CORS headers to every response.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins for now — you can restrict to your domain
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
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

	// Compose middlewares: logging -> CORS -> handler
	handler := loggingMiddleware(corsMiddleware(http.HandlerFunc(scrape.ScrapeHandler)))
	mux.Handle("/scrape", handler)

	portStr := fmt.Sprintf("%s:%d", addr, port)
	slog.Info("Server starting", "addr", addr, "port", port)

	return http.ListenAndServe(portStr, mux)
}
