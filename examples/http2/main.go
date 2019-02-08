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

	router.Use(rz.HTTPNewHandler(log.Logger))
	router.Use(requestIDMiddleware)
	router.Use(rz.HTTPRequestIDHandler("req_id"))
	router.Use(rz.HTTPUserAgentHandler("ua"))

	router.Get("/", HelloWorld)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("listening", func(e *rz.Event) { e.Err(err) })
	}
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := "uuid"
		ctx := context.WithValue(r.Context(), rz.HTTPCtxKeyRequestID{}, requestID)
		w.Header().Set("Request-Id", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := rz.FromCtx(ctx)
	logger.Info("hello from GET /", nil)
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
