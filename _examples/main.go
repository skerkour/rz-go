package main

import (
	"os"

	"github.com/astroflow/astro-go"
	"github.com/astroflow/astro-go/log"
)

func main() {

	env := os.Getenv("GO_ENV")
	hostname, _ := os.Hostname()

	log.Init(
		//   astro.SetWriter(os.Stderr),
		astro.AddFields(
			"service", "api",
			"host", hostname,
			"environment", env,
		),
		astro.SetFormatter(astro.NewConsoleFormatter()),
	)

	if env == "production" {
		log.Config(
			astro.SetFormatter(astro.JSONFormatter{}),
			astro.SetLevel(astro.InfoLevel),
		)
	}

	subLogger := log.With("contextual_field", 42)

	log.Info("info from logger")
	log.With("field1", "hello world", "field2", 999.99).Info("info from logger with fields")
	subLogger.Debug("debug from sublogger")
	subLogger.Warn("warning from sublogger")
	subLogger.Error("error from sublogger")
}
