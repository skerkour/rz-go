## Astro

[Make logging great again](https://kerkour.com/post/logging/)

[![GoDoc](https://godoc.org/github.com/bloom42/astro-go?status.svg)](https://godoc.org/github.com/bloom42/astro-go)
[![Build Status](https://travis-ci.org/bloom42/astro-go.svg?branch=master)](https://travis-ci.org/bloom42/astro-go)
[![GitHub release](https://img.shields.io/github/release/bloom42/astro-go.svg)](https://github.com/bloom42/astro-go/releases)

![Console logging](docs/example_screenshot.png)


1. [Quickstart](#quickstart)
2. [Benchmark](#benchmark)
3. [Configuration](#configuration)
4. [Examples](#examples)
5. [Contributing](#contributing)
6. [License](#license)

-------------------

## Quickstart

```go

```

## Benchmarks

```

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


## Examples

See the [examples](https://github.com/bloom42/astro-go/tree/master/examples) folder.


## Contributing

See [https://opensource.bloom.sh/contributing](https://opensource.bloom.sh/contributing)


## License

See `LICENSE.txt` and [https://opensource.bloom.sh/licensing](https://opensource.bloom.sh/licensing)
