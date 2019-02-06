## Astro

[Make logging great again](https://kerkour.com/post/logging/)

[![GoDoc](https://godoc.org/github.com/bloom42/astro-go?status.svg)](https://godoc.org/github.com/bloom42/astro-go)
[![Build Status](https://travis-ci.org/bloom42/astro-go.svg?branch=master)](https://travis-ci.org/bloom42/astro-go)
[![GitHub release](https://img.shields.io/github/release/bloom42/astro-go.svg)](https://github.com/bloom42/astro-go/releases)

![Console logging](docs/example_screenshot.png)


1. [Quickstart](#quickstart)
2. [Benchmark](#benchmark)
3. [Log usage](#log-usage)
4. [Configuration](#configuration)
5. [HTTPHandler](#httphandler)
6. [Examples](#examples)
7. [Contributing](#contributing)
8. [License](#license)

-------------------

## Quickstart

```go
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

	log.Info("info from logger")
	log.With("field1", "hello world", "field2", 999.99).Info("info from logger with fields")
	subLogger.Debug("debug from sublogger")
	subLogger.Warn("warning from sublogger")
	subLogger.Error("error from sublogger")
}
```

## Benchmarks

```
$ make benchmarks
cd benchmarks && ./run.sh
goos: linux
goarch: amd64
pkg: github.com/bloom42/astro-go/benchmarks
BenchmarkWithoutFields/sirupsen/logrus-4                  300000              4653 ns/op            1137 B/op         24 allocs/op
BenchmarkWithoutFields/bloom42/astro-go-4                1000000              1593 ns/op             832 B/op         13 allocs/op
Benchmark10FieldsContext/sirupsen/logrus-4                100000             21492 ns/op            3261 B/op         50 allocs/op
Benchmark10FieldsContext/bloom42/astro-go-4               300000              4646 ns/op            3196 B/op         17 allocs/op
Benchmark10Fields/sirupsen/logrus-4                        50000             25341 ns/op            4043 B/op         54 allocs/op
Benchmark10Fields/bloom42/astro-go-4                      300000              5067 ns/op            3516 B/op         18 allocs/op
PASS
ok      github.com/bloom42/astro-go/benchmarks  10.970s
```

## Log usage

```go
import (
    "github.com/bloom42/astro-go/log"
)

log.Config(options ...astro.LoggerOption) error
log.With(fields ...interface{}) astro.Logger

// each of the following have it's XXXf companion (e.g. log.Debugf("%s" ,err) ...)
log.Debug(message string)
log.Info(message string)
log.Warn(message string)
log.Error(message string)
log.Fatal(message string) // log with the "fatal" level then os.Exit(1)

log.Track(fields ...interface{}) // log an event without level nor message
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

## HTTPHandler

astro provides an http handler helper to log http requests
```go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bloom42/astro-go"
	"github.com/bloom42/astro-go/log"
)

func main() {
	env := os.Getenv("GO_ENV")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Config(
		astro.AddFields(
			"service", "api",
			"host", "abcd",
			"environment", env,
		),
		astro.SetFormatter(astro.JSONFormatter{}),
	)

	http.HandleFunc("/", HelloWorld)

	middleware := astro.HTTPHandler(log.With())
	err := http.ListenAndServe(":"+port, middleware(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
```


## Examples

See the [examples](https://github.com/bloom42/astro-go/tree/master/examples) folder.


## Contributing

See [https://opensource.bloom.sh/contributing](https://opensource.bloom.sh/contributing)


## License

See `LICENSE.txt` and [https://opensource.bloom.sh/licensing](https://opensource.bloom.sh/licensing)
