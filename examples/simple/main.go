package main

import (
	"os"

	"github.com/bloom42/astro-go"
)

func main() {

	env := os.Getenv("GO_ENV")
	hostname, _ := os.Hostname()
	_ = hostname

	// log := astro.New(astro.Fields(func(e *astro.Event) {
	// 	e.String("service", "api").
	// 		String("host", hostname).
	// 		String("environment", env)
	// }))
	log := astro.New()
	log2 := log.Config(astro.With(func(e *astro.Event) {
		e.String("service", "api").
			String("host", hostname).
			String("environment", env)
	}))

	// log.Logger = log.Config(
	// 	//   astro.SetWriter(os.Stderr),
	// 	astro.AddFields(
	// 		"service", "api",
	// 		"host", hostname,
	// 		"environment", env,
	// 	),
	// 	astro.SetFormatter(astro.NewConsoleFormatter()),
	// )

	if env == "production" {
		log = log.Config(
			// astro.SetFormatter(astro.JSONFormatter{}),
			astro.Level(astro.InfoLevel),
		)
	}

	// subLogger := log.With("contextual_field", 42)
	// subLogger.Config(
	// 	astro.SetFormatter(astro.NewCLIFormatter()),
	// )

	// otherLogger := astro.NewLogger()
	// otherOtherLogger := otherLogger.With("field", "MyUUID")

	log.Info("info from logger", func(e *astro.Event) {
		e.String("hello", "world")
	})
	log2.Info("info from logger2", func(e *astro.Event) {
		e.String("hello2", "world2")
	})
	// log.With("field1", "hello world", "field2", 999.99).Info("info from logger with fields")
	// subLogger.Debug("debug from sublogger")
	// subLogger.Info("info form subLogger")
	// subLogger.Warn("warning from sublogger")
	// subLogger.Error("error from sublogger")
	// otherLogger.Info("info from other logger")
	// otherOtherLogger.Info("info from otherOther logger")
	// myFunc(otherLogger)
}

// func myFunc(logger astro.Logger) {
// 	logger.Info("info from other logger in func")
// }
