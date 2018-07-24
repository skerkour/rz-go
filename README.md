## Astroflow

[![GoDoc](https://godoc.org/github.com/astroflow/astroflow-go?status.svg)](https://godoc.org/github.com/astroflow/astroflow-go)
[![GitHub release](https://img.shields.io/github/release/astroflow/astroflow-go.svg)](https://github.com/astroflow/astroflow-go/releases)

![Console logging](_docs/example_screenshot.png)

```go
package main

import (
	"os"

	"github.com/astroflow/astroflow-go"
	"github.com/astroflow/astroflow-go/log"
)

func main() {

	env := os.Getenv("GO_ENV")
	hostname, _ := os.Hostname()

	log.Init(
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

	log.Info("info from logger")
	log.With("field1", "hello world", "field2", 999.99).Info("info from logger with fields")
	subLogger.Debug("debug from sublogger")
	subLogger.Warn("warning from sublogger")
	subLogger.Error("error from sublogger")
}
```

## Log
```go
import (
    "github.com/astroflow/astroflow-go/log"
)

log.Init(options ...astro.LoggerOption) error
log.Config(options ...astro.LoggerOption) error
log.With(fields ...interface{}) astro.Logger
log.Debug(message string)
log.Info(message string)
log.Warn(message string)
log.Error(message string)
log.Fatal(message string) // log with the "fatal" level then os.Exit(1)
```

## Configuration

```go
SetWriter(writer io.Writer) // default to os.Stdout
SetFormatter(formatter astro.Formatter) // default to astro.JSONFormatter
SetFields(fields ...interface{})
AddFields(fields ...interface{})
SetInsertTimestampField(insert bool) // default to true
SetLevel(level Level) // default to astro.DebugLevel
SetTimestampFieldName(fieldName string) // default to astro.TimestampFieldName ("timestamp")
SetLevelFieldName(fieldName string) // default to astro.LevelFieldName ("level")
SetTimestampFunc(fn func() time.Time) // default to time.Now().UTC
AddHook(hook astro.Hook)
```
