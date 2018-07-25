package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/astroflow/astroflow-go"
	"github.com/astroflow/astroflow-go/log"
)

func main() {
	env := os.Getenv("GO_ENV")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Init(
		astroflow.AddFields(
			"service", "api",
			"host", "abcd",
			"environment", env,
		),
		astroflow.SetFormatter(astroflow.NewConsoleFormatter()),
		astroflow.SetFormatter(astroflow.JSONFormatter{}),
	)

	http.HandleFunc("/", HelloWorld)

	middleware := astroflow.HTTPHandler(log.With())
	err := http.ListenAndServe(":"+port, middleware(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
