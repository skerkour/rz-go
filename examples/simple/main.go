package main

import (
	"os"

	"github.com/bloom42/rz-go"
)

func main() {

	env := os.Getenv("GO_ENV")
	hostname, _ := os.Hostname()
	_ = hostname

	// log := rz.New(rz.Fields(func(e *rz.Event) {
	// 	e.String("service", "api").
	// 		String("host", hostname).
	// 		String("environment", env)
	// }))
	log := rz.New()
	log2 := log.Config(rz.With(func(e *rz.Event) {
		e.String("service", "api").
			String("host", hostname).
			String("environment", env)
	}))

	// log.Logger = log.Config(
	// 	//   rz.SetWriter(os.Stderr),
	// 	rz.AddFields(
	// 		"service", "api",
	// 		"host", hostname,
	// 		"environment", env,
	// 	),
	// 	rz.SetFormatter(rz.NewConsoleFormatter()),
	// )

	if env == "production" {
		log = log.Config(
			// rz.SetFormatter(rz.JSONFormatter{}),
			rz.Level(rz.InfoLevel),
		)
	}

	// subLogger := log.With("contextual_field", 42)
	// subLogger.Config(
	// 	rz.SetFormatter(rz.NewCLIFormatter()),
	// )

	// otherLogger := rz.NewLogger()
	// otherOtherLogger := otherLogger.With("field", "MyUUID")

	log.Info("info from logger", func(e *rz.Event) {
		e.String("hello", "world")
	})
	log2.Info("info from logger2", func(e *rz.Event) {
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

// func myFunc(logger rz.Logger) {
// 	logger.Info("info from other logger in func")
// }
