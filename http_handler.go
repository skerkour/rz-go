package rz

// import (
// 	"net"
// 	"net/http"
// 	"time"
// )

// type wrapper struct {
// 	http.ResponseWriter
// 	http.Flusher
// 	http.CloseNotifier

// 	written int
// 	status  int
// }

// // WriteHeader wrapper to capture status code.
// func (w *wrapper) WriteHeader(code int) {
// 	w.status = code
// 	w.ResponseWriter.WriteHeader(code)
// }

// // Write wrapper to capture response size.
// func (w *wrapper) Write(b []byte) (int, error) {
// 	n, err := w.ResponseWriter.Write(b)
// 	w.written += n
// 	return n, err
// }

// // Flush implementation.
// func (w *wrapper) Flush() {
// 	if w.Flusher != nil {
// 		w.Flusher.Flush()
// 	}
// }

// // HTTPHandler is a helper middleware to log HTTP requests
// //
// // It's simple enough to be copy pasted if it's does not fit your needs (eg. you want to rename fields)
// func HTTPHandler(logger Logger) func(next http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			start := time.Now()

// 			res := &wrapper{
// 				ResponseWriter: w,
// 				written:        0,
// 				status:         200,
// 			}

// 			if f, ok := w.(http.Flusher); ok {
// 				res.Flusher = f
// 			}

// 			if c, ok := w.(http.CloseNotifier); ok {
// 				res.CloseNotifier = c
// 			}

// 			scheme := "http"
// 			if r.TLS != nil {
// 				scheme = "https"
// 			}
// 			method := r.Method
// 			uri := r.RequestURI
// 			userAgent := r.Header.Get("user-agent")

// 			remote := r.RemoteAddr
// 			host, _, err := net.SplitHostPort(remote)
// 			if err == nil {
// 				remote = host
// 			}

// 			requestLogger := logger.Config((With(func(e *Event) {
// 				e.String("scheme", scheme).
// 					String("host", r.Host).
// 					String("uri", uri).
// 					String("method", method).
// 					String("remote_address", remote).
// 					String("user_agent", userAgent)
// 			})))

// 			next.ServeHTTP(res, r)

// 			durationMs := time.Since(start).Nanoseconds() / 1000000
// 			if durationMs < 1 {
// 				durationMs = 1
// 			}

// 			requestID := ""
// 			if rid, ok := r.Context().Value(HTTPCtxRequestIDKey).(string); ok {
// 				requestID = rid
// 			}

// 			status := res.status
// 			fields := func(e *Event) {
// 				e.Int("status", res.status).
// 					Int("size", res.written).
// 					Int64("duration", durationMs)
// 				if len(requestID) != 0 {
// 					e.String("request_id", requestID)
// 				}
// 			}

// 			message := "access"
// 			switch {
// 			case status < 400:
// 				requestLogger.Info(message, fields)
// 			case status < 500:
// 				requestLogger.Warn(message, fields)
// 			default:
// 				requestLogger.Error(message, fields)
// 			}
// 		})
// 	}
// }
