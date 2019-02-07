<p align="center">
  <h3 align="center">rz</h3>
  <p align="center">RIPZap - The fastest structured, leveled JSON logger for Go 📖</p>
</p>

--------

[Make logging great again](https://kerkour.com/post/logging/)

[![GoDoc](https://godoc.org/github.com/bloom42/rz-go?status.svg)](https://godoc.org/github.com/bloom42/rz-go)
[![Build Status](https://travis-ci.org/bloom42/rz-go.svg?branch=master)](https://travis-ci.org/bloom42/rz-go)
[![GitHub release](https://img.shields.io/github/release/bloom42/rz-go.svg)](https://github.com/bloom42/rz-go/releases)
[![Coverage](http://gocover.io/_badge/github.com/rs/zerolog)](http://gocover.io/github.com/rs/zerolog)

![Console logging](docs/example_screenshot.png)

The rz package provides a fast and simple logger dedicated to JSON output.

Rz's API is designed to provide both a great developer experience and stunning [performance](#benchmarks). Its unique chaining API allows zerolog to write JSON (or CBOR) log events by avoiding allocations and reflection.

Uber's [zap](https://godoc.org/go.uber.org/zap) and rs's [zerolog](https://godoc.org/github.com/rs/zerolog) libraries pioneered this approach. Rz is taking this concept to the next level with a simpler to use API and even better performance.

To keep the code base and the API simple, rz focuses on efficient structured logging only. Pretty logging on the console is made possible using the provided (but inefficient) `rz.ConsoleFormatter`.


1. [Quickstart](#quickstart)
2. [Configuration](#configuration)
3. [Examples](#examples)
4. [Benchmarks](#benchmarks)
5. [Contributing](#contributing)
6. [License](#license)

-------------------

## Quickstart

```go
package main

import (
	"os"

	"github.com/bloom42/rz-go"
	"github.com/bloom42/rz-go/log"
)

func main() {

	env := os.Getenv("GO_ENV")
	hostname, _ := os.Hostname()

	// update global logger's context fields
	log.Logger = log.Config(rz.With(func(e *rz.Event) {
		e.String("hostname", hostname).
			String("environment", env)
	}))

	if env == "production" {
		log.Logger = log.Config(rz.Level(rz.InfoLevel))
	}

	log.Info("info from logger", func(e *rz.Event) {
		e.String("hello", "world")
	})
	// {"level":"info","hostname":"","environment":"","hello":"world","timestamp":"2019-02-07T09:30:07Z","message":"info from logger"}
}
```


## Configuration

```go
// Writer update logger's writer.
func Writer(writer io.Writer) LoggerOption {}
// Level update logger's level.
func Level(lvl LogLevel) LoggerOption {}
// Sampler update logger's sampler.
func Sampler(sampler LogSampler) LoggerOption {}
// AddHook appends hook to logger's hook
func AddHook(hook LogHook) LoggerOption {}
// Hooks replaces logger's hooks
func Hooks(hooks ...LogHook) LoggerOption {}
// With replaces logger's context fields
func With(fields func(*Event)) LoggerOption {}
// Stack enable/disable stack in error messages.
func Stack(enableStack bool) LoggerOption {}
// Timestamp enable/disable timestamp logging in error messages.
func Timestamp(enableTimestamp bool) LoggerOption {}
// Caller enable/disable caller field in message messages.
func Caller(enableCaller bool) LoggerOption {}
// Formatter update logger's formatter.
func Formatter(formatter LogFormatter) LoggerOption {}
// TimestampFieldName update logger's timestampFieldName.
func TimestampFieldName(timestampFieldName string) LoggerOption {}
// LevelFieldName update logger's levelFieldName.
func LevelFieldName(levelFieldName string) LoggerOption {}
// MessageFieldName update logger's messageFieldName.
func MessageFieldName(messageFieldName string) LoggerOption {}
// ErrorFieldName update logger's errorFieldName.
func ErrorFieldName(errorFieldName string) LoggerOption {}
// CallerFieldName update logger's callerFieldName.
func CallerFieldName(callerFieldName string) LoggerOption {}
// CallerSkipFrameCount update logger's callerSkipFrameCount.
func CallerSkipFrameCount(callerSkipFrameCount int) LoggerOption {}
// ErrorStackFieldName update logger's errorStackFieldName.
func ErrorStackFieldName(errorStackFieldName string) LoggerOption {}
```


## Examples

See the [examples](https://github.com/bloom42/rz-go/tree/master/examples) folder.


## Benchmarks

```
$ make benchmarks
cd benchmarks && ./run.sh
goos: linux
goarch: amd64
pkg: github.com/bloom42/rz-go/benchmarks
BenchmarkDisabledWithoutFields/sirupsen/logrus-4                100000000               18.0 ns/op            16 B/op          1 allocs/op
BenchmarkDisabledWithoutFields/rs/zerolog-4                     500000000                3.81 ns/op            0 B/op          0 allocs/op
BenchmarkDisabledWithoutFields/bloom42/rz-go-4                  1000000000               2.52 ns/op            0 B/op          0 allocs/op
BenchmarkWithoutFields/sirupsen/logrus-4                          300000              4633 ns/op            1137 B/op         24 allocs/op
BenchmarkWithoutFields/rs/zerolog-4                              5000000               262 ns/op               0 B/op          0 allocs/op
BenchmarkWithoutFields/bloom42/rz-go-4                          20000000                63.1 ns/op             0 B/op          0 allocs/op
Benchmark10Context/sirupsen/logrus-4                              100000             19219 ns/op            3261 B/op         50 allocs/op
Benchmark10Context/rs/zerolog-4                                  5000000               271 ns/op               0 B/op          0 allocs/op
Benchmark10Context/bloom42/rz-go-4                              20000000                67.8 ns/op             0 B/op          0 allocs/op
Benchmark10Fields/sirupsen/logrus-4                               100000             24765 ns/op            4042 B/op         54 allocs/op
Benchmark10Fields/rs/zerolog-4                                    500000              2811 ns/op             640 B/op          6 allocs/op
Benchmark10Fields/bloom42/rz-go-4                                 500000              2615 ns/op             640 B/op          6 allocs/op
Benchmark10Fields10Context/sirupsen/logrus-4                       50000             23952 ns/op            4567 B/op         53 allocs/op
Benchmark10Fields10Context/rs/zerolog-4                           500000              2813 ns/op             640 B/op          6 allocs/op
Benchmark10Fields10Context/bloom42/rz-go-4                        500000              2556 ns/op             640 B/op          6 allocs/op
PASS
ok      github.com/bloom42/rz-go/benchmarks     26.155s
```


## Contributing

See [https://opensource.bloom.sh/contributing](https://opensource.bloom.sh/contributing)


## License

See `LICENSE.txt` and [https://opensource.bloom.sh/licensing](https://opensource.bloom.sh/licensing)
