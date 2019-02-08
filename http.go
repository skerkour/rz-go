package rz

import (
	"net"
	"net/http"
	"time"
)

// Key to use when setting the request ID.
type httpCtxKeyRequestID int

// HTTPCtxRequestIDKey is the key that holds the unique request ID in a request context.
const HTTPCtxRequestIDKey httpCtxKeyRequestID = 0

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

// HTTPNewHandler injects logger into requests context.
func HTTPNewHandler(log Logger) func(http.Handler) http.Handler {
	logger := log.Config()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a copy of the logger (including internal context slice)
			// to prevent data race when using UpdateContext.
			start := time.Now()

			sw := statusWriter{ResponseWriter: w}
			r = r.WithContext(logger.ToCtx(r.Context()))
			next.ServeHTTP(&sw, r)

			durationMs := time.Since(start).Nanoseconds() / 1000000
			if durationMs < 1 {
				durationMs = 1
			}

			// status := w.Status
			fields := func(e *Event) {
				e.Int("status", sw.status).
					Int("size", sw.length).
					Int64("duration", durationMs)
			}

			message := "access"
			switch {
			case sw.status < 400:
				logger.Info(message, fields)
			case sw.status < 500:
				logger.Warn(message, fields)
			default:
				logger.Error(message, fields)
			}
		})
	}
}

// HTTPURLHandler adds the requested URL as a field to the context's logger
// using fieldKey as field key.
func HTTPURLHandler(fieldKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := FromCtx(r.Context())
			e := &Event{}
			e.String(fieldKey, r.URL.String())
			log.updateContext(e)
			next.ServeHTTP(w, r)
		})
	}
}

// HTTPMethodHandler adds the request method as a field to the context's logger
// using fieldKey as field key.
func HTTPMethodHandler(fieldKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := FromCtx(r.Context())
			e := &Event{}
			e.String(fieldKey, r.Method)
			log.updateContext(e)
			next.ServeHTTP(w, r)
		})
	}
}

// HTTPRemoteAddrHandler adds the request's remote address as a field to the context's logger
// using fieldKey as field key.
func HTTPRemoteAddrHandler(fieldKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
				log := FromCtx(r.Context())
				e := &Event{}
				e.String(fieldKey, host)
				log.updateContext(e)
			}
			next.ServeHTTP(w, r)
		})
	}
}

// HTTPUserAgentHandler adds the request's user-agent as a field to the context's logger
// using fieldKey as field key.
func HTTPUserAgentHandler(fieldKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ua := r.Header.Get("User-Agent")
			log := FromCtx(r.Context())
			e := &Event{}
			e.String(fieldKey, ua)
			log.updateContext(e)

			next.ServeHTTP(w, r)
		})
	}
}

// HTTPRefererHandler adds the request's referer as a field to the context's logger
// using fieldKey as field key.
func HTTPRefererHandler(fieldKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ref := r.Header.Get("Referer")
			log := FromCtx(r.Context())
			e := &Event{}
			e.String(fieldKey, ref)
			log.updateContext(e)
			next.ServeHTTP(w, r)
		})
	}
}

// HTTPRequestIDHandler adds the request's id as a field to the context's logger
// using fieldKey as field key.
//
// request's id must be present in context with the `HTTPCtxKeyRequestID` key otherwise it's skipped
func HTTPRequestIDHandler(fieldKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if rid, ok := ctx.Value(HTTPCtxRequestIDKey).(string); ok {
				log := FromCtx(r.Context())
				e := &Event{}
				e.String(fieldKey, rid)
				log.updateContext(e)
			}
			next.ServeHTTP(w, r)
		})
	}
}
