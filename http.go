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

type httpHandler struct {
	logger             Logger
	message            string
	urlField           string
	methodField        string
	schemeField        string
	hostField          string
	remoteAddressField string
	userAgentField     string
	sizeField          string
	statusField        string
	durationField      string
	requestIDField     string
}

type HTTPHandlerOption func(*httpHandler)

func URL(urlFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.urlField = urlFieldName
	}
}

func Message(message string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.message = message
	}
}

func Method(methodFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.methodField = methodFieldName
	}
}

func Scheme(schemeFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.schemeField = schemeFieldName
	}
}

func Host(hostFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.hostField = hostFieldName
	}
}

func RemoteAddress(remoteAddressFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.remoteAddressField = remoteAddressFieldName
	}
}

func UserAgent(userAgentFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.userAgentField = userAgentFieldName
	}
}

func Size(sizeFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.sizeField = sizeFieldName
	}
}

func Status(statusFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.statusField = statusFieldName
	}
}

func Duration(durationFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.durationField = durationFieldName
	}
}

func RequestID(requestIDFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.requestIDField = requestIDFieldName
	}
}

// HTTPHandler is a helper middleware to log HTTP requests
func HTTPHandler(logger Logger, options ...HTTPHandlerOption) func(next http.Handler) http.Handler {
	// store a copy of the logger
	handler := httpHandler{
		logger:             logger.Config(),
		message:            "access",
		urlField:           "url",
		methodField:        "method",
		schemeField:        "scheme",
		hostField:          "host",
		remoteAddressField: "remote_address",
		userAgentField:     "user_agent",
		sizeField:          "size",
		statusField:        "status",
		durationField:      "duration",
		requestIDField:     "request_id",
	}
	for _, option := range options {
		option(&handler)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			resWrapper := &responseWrapper{
				ResponseWriter: w,
				written:        0,
				status:         200,
			}

			if f, ok := w.(http.Flusher); ok {
				resWrapper.Flusher = f
			}

			if c, ok := w.(http.CloseNotifier); ok {
				resWrapper.CloseNotifier = c
			}

			if handler.schemeField != "" {
				scheme := "http"
				if r.TLS != nil {
					scheme = "https"
				}
				handler.logger.updateContext(func(e *Event) {
					e.String(handler.schemeField, scheme)
				})
			}

			if handler.methodField != "" {
				handler.logger.updateContext(func(e *Event) {
					e.String(handler.methodField, r.Method)
				})
			}

			if handler.urlField != "" {
				handler.logger.updateContext(func(e *Event) {
					e.String(handler.urlField, r.RequestURI)
				})
			}

			if handler.hostField != "" {
				handler.logger.updateContext(func(e *Event) {
					e.String(handler.hostField, r.Host)
				})
			}

			if handler.remoteAddressField != "" {
				remote := r.RemoteAddr
				host, _, err := net.SplitHostPort(remote)
				if err == nil {
					remote = host
				}
				handler.logger.updateContext(func(e *Event) {
					e.String(handler.remoteAddressField, remote)
				})
			}

			if handler.userAgentField != "" {
				handler.logger.updateContext(func(e *Event) {
					e.String(handler.userAgentField, r.Header.Get("user-agent"))
				})
			}

			next.ServeHTTP(resWrapper, r)

			if handler.sizeField != "" {
				handler.logger.updateContext(func(e *Event) {
					e.Int(handler.sizeField, resWrapper.written)
				})
			}

			status := resWrapper.status
			if handler.statusField != "" {
				handler.logger.updateContext(func(e *Event) {
					e.Int(handler.statusField, status)
				})
			}

			if handler.durationField != "" {
				durationMs := time.Since(start).Nanoseconds() / 1000000
				if durationMs < 1 {
					durationMs = 1
				}
				handler.logger.updateContext(func(e *Event) {
					e.Int64(handler.durationField, durationMs)
				})
			}

			if handler.requestIDField != "" {
				requestID := ""
				if rid, ok := r.Context().Value(HTTPCtxRequestIDKey).(string); ok {
					requestID = rid
				}
				handler.logger.updateContext(func(e *Event) {
					e.String(handler.requestIDField, requestID)
				})
			}

			switch {
			case status < 400:
				handler.logger.Info(handler.message, nil)
			case status < 500:
				handler.logger.Warn(handler.message, nil)
			default:
				handler.logger.Error(handler.message, nil)
			}
		})
	}
}

type responseWrapper struct {
	http.ResponseWriter
	http.Flusher
	http.CloseNotifier

	written int
	status  int
}

// WriteHeader wrapper to capture status code.
func (w *responseWrapper) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// Write wrapper to capture response size.
func (w *responseWrapper) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.written += n
	return n, err
}

// Flush implementation.
func (w *responseWrapper) Flush() {
	if w.Flusher != nil {
		w.Flusher.Flush()
	}
}
