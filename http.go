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

// HTTPHandlerOption are used to configure a HTTPHandler.
type HTTPHandlerOption func(*httpHandler)

// URL is used to updated HTTPHandler's url field name. Set an empty string to disable the field.
func URL(urlFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.urlField = urlFieldName
	}
}

// Message is used to updated HTTPHandler's message field name. Set an empty string to disable it.
func Message(message string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.message = message
	}
}

// Method is used to updated HTTPHandler's method field name. Set an empty string to disable the field.
func Method(methodFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.methodField = methodFieldName
	}
}

// Method is used to updated HTTPHandler's method field name. Set an empty string to disable the field.
func Scheme(schemeFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.schemeField = schemeFieldName
	}
}

// Host is used to updated HTTPHandler's host field name. Set an empty string to disable the field.
func Host(hostFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.hostField = hostFieldName
	}
}

// RemoteAddress is used to updated HTTPHandler's remote address field name. Set an empty string to disable the field.
func RemoteAddress(remoteAddressFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.remoteAddressField = remoteAddressFieldName
	}
}

// UserAgent is used to updated HTTPHandler's user agent field name. Set an empty string to disable the field.
func UserAgent(userAgentFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.userAgentField = userAgentFieldName
	}
}

// Size is used to updated HTTPHandler's size field name. Set an empty string to disable the field.
func Size(sizeFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.sizeField = sizeFieldName
	}
}

// Status is used to updated HTTPHandler's status field name. Set an empty string to disable the field.
func Status(statusFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.statusField = statusFieldName
	}
}

// Duration is used to updated HTTPHandler's duration field name. Set an empty string to disable the field.
func Duration(durationFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.durationField = durationFieldName
	}
}

// RequestID is used to updated HTTPHandler's request ID field name. Set an empty string to disable the field.
func RequestID(requestIDFieldName string) HTTPHandlerOption {
	return func(handler *httpHandler) {
		handler.requestIDField = requestIDFieldName
	}
}

// HTTPHandler is a helper middleware to log HTTP requests
func HTTPHandler(logger Logger, options ...HTTPHandlerOption) func(next http.Handler) http.Handler {
	logger = logger.Config()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// store a copy of the logger
			handler := httpHandler{
				logger:             logger,
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
				handler.logger.Append(func(e *Event) {
					e.String(handler.schemeField, scheme)
				})
			}

			if handler.methodField != "" {
				handler.logger.Append(func(e *Event) {
					e.String(handler.methodField, r.Method)
				})
			}

			if handler.urlField != "" {
				handler.logger.Append(func(e *Event) {
					e.String(handler.urlField, r.RequestURI)
				})
			}

			if handler.hostField != "" {
				handler.logger.Append(func(e *Event) {
					e.String(handler.hostField, r.Host)
				})
			}

			if handler.remoteAddressField != "" {
				remote := r.RemoteAddr
				host, _, err := net.SplitHostPort(remote)
				if err == nil {
					remote = host
				}
				handler.logger.Append(func(e *Event) {
					e.String(handler.remoteAddressField, remote)
				})
			}

			if handler.userAgentField != "" {
				handler.logger.Append(func(e *Event) {
					e.String(handler.userAgentField, r.Header.Get("user-agent"))
				})
			}

			next.ServeHTTP(resWrapper, r)

			if handler.sizeField != "" {
				handler.logger.Append(func(e *Event) {
					e.Int(handler.sizeField, resWrapper.written)
				})
			}

			status := resWrapper.status
			if handler.statusField != "" {
				handler.logger.Append(func(e *Event) {
					e.Int(handler.statusField, status)
				})
			}

			if handler.durationField != "" {
				durationUs := time.Since(start).Nanoseconds() / 1000
				if durationUs < 1 {
					durationUs = 1
				}
				handler.logger.Append(func(e *Event) {
					e.Int64(handler.durationField, durationUs)
				})
			}

			if handler.requestIDField != "" {
				// fmt.Println("in requestID")
				requestID := ""
				if rid, ok := r.Context().Value(HTTPCtxRequestIDKey).(string); ok {
					// fmt.Println("in requestID 2")
					requestID = rid
				}
				handler.logger.Append(func(e *Event) {
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
