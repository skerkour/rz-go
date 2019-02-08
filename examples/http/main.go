package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/bloom42/rz-go"
	"github.com/bloom42/rz-go/log"
	"github.com/go-chi/chi"
)

func main() {
	env := os.Getenv("GO_ENV")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Logger = log.Config(
		rz.With(func(e *rz.Event) {
			e.String("service", "api").
				String("host", "abcd.local").
				String("environment", env)
		}),
	)

	router := chi.NewRouter()

	// replace size field name by latency and disable userAgent logging
	loggingMiddleware := rz.HTTPHandler(log.Logger, rz.Duration("latency"), rz.UserAgent(""))
	router.Use(loggingMiddleware)
	router.Use(requestIDMiddleware)
	router.Use(injectLoggerMiddleware(log.Logger))

	router.Get("/", HelloWorld)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("listening", func(e *rz.Event) { e.Err(err) })
	}
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := "uuid"
		ctx := context.WithValue(r.Context(), rz.HTTPCtxRequestIDKey, requestID)
		w.Header().Set("Request-Id", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func injectLoggerMiddleware(logger rz.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if rid, ok := r.Context().Value(rz.HTTPCtxRequestIDKey).(string); ok {
				logger = logger.Config(rz.With(func(e *rz.Event) {
					e.String("request_id", rid)
				}))
				ctx := logger.ToCtx(r.Context())
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	logger := rz.FromCtx(r.Context())
	logger.Info("hello from GET /", nil)
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
