package main

import (
	"os"

	"github.com/astroflow/astroflow-go"
	"github.com/astroflow/astroflow-go/log"
)

func main() {

	env := os.Getenv("GO_ENV")
	hostname, _ := os.Hostname()

	log.Config(
		//   astroflow.SetWriter(os.Stderr),
		astroflow.AddFields(
			"service", "api",
			"host", hostname,
			"environment", env,
		),
		astroflow.SetFormatter(astroflow.NewConsoleFormatter()),
	)

	if env == "production" {
		log.Config(
			astroflow.SetFormatter(astroflow.JSONFormatter{}),
			astroflow.SetLevel(astroflow.InfoLevel),
		)
	}

	subLogger := log.With("contextual_field", 42)
	subLogger.Config(
		astroflow.SetFormatter(astroflow.NewCLIFormatter()),
	)

	log.Info("info from logger")
	log.With("field1", "hello world", "field2", 999.99).Info("info from logger with fields")
	subLogger.Debug("debug from sublogger")
	subLogger.Info("info form subLogger")
	subLogger.Warn("warning from sublogger")
	subLogger.Error("error from sublogger")
}
