package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bloom42/rz-go"
	"github.com/bloom42/rz-go/log"
)

func main() {
	env := os.Getenv("GO_ENV")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Config(
		rz.AddFields(
			"service", "api",
			"host", "abcd",
			"environment", env,
		),
		rz.SetFormatter(rz.JSONFormatter{}),
	)

	http.HandleFunc("/", HelloWorld)

	middleware := rz.HTTPHandler(log.With())
	err := http.ListenAndServe(":"+port, middleware(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
