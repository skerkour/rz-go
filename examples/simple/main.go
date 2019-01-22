package main

import (
	"os"

	"github.com/bloom42/astro-go"
	"github.com/bloom42/astro-go/log"
)

func main() {

	env := os.Getenv("GO_ENV")
	hostname, _ := os.Hostname()

	log.Config(
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
	subLogger.Config(
		astro.SetFormatter(astro.NewCLIFormatter()),
	)

	otherLogger := astro.NewLogger()

	log.Info("info from logger")
	log.With("field1", "hello world", "field2", 999.99).Info("info from logger with fields")
	subLogger.Debug("debug from sublogger")
	subLogger.Info("info form subLogger")
	subLogger.Warn("warning from sublogger")
	subLogger.Error("error from sublogger")
	otherLogger.Info("info from other logger")
}
