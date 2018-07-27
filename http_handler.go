package astroflow

import (
	"fmt"
	"net/http"
	"time"
)

type wrapper struct {
	http.ResponseWriter
	http.Flusher
	http.CloseNotifier

	written int
	status  int
}

// WriteHeader wrapper to capture status code.
func (w *wrapper) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// Write wrapper to capture response size.
func (w *wrapper) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.written += n
	return n, err
}

// Flush implementation.
func (w *wrapper) Flush() {
	if w.Flusher != nil {
		w.Flusher.Flush()
	}
}

func HTTPHandler(logger Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			res := &wrapper{
				ResponseWriter: w,
				written:        0,
				status:         200,
			}

			if f, ok := w.(http.Flusher); ok {
				res.Flusher = f
			}

			if c, ok := w.(http.CloseNotifier); ok {
				res.CloseNotifier = c
			}

			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}
			method := r.Method
			uri := r.RequestURI
			userAgent := r.Header.Get("user-agent")

			requestLogger := logger.With(
				"scheme", scheme,
				"host", r.Host,
				"url", uri,
				"method", method,
				"remote_address", r.RemoteAddr,
				"user_agent", userAgent,
			)

			next.ServeHTTP(res, r)

			latencyMs := time.Since(start).Nanoseconds() / 1000000
			if latencyMs < 1 {
				latencyMs = 1
			}

			status := res.status
			requestLogger = requestLogger.With(
				"status", res.status,
				"size", res.written,
				"latency", latencyMs,
			)

			message := fmt.Sprintf("%d %s %s", status, method, uri)
			switch {
			case status < 400:
				requestLogger.Info(message)
			case status < 500:
				requestLogger.Warn(message)
			default:
				requestLogger.Error(message)
			}
		})
	}
}
