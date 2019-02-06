## Astro

[Make logging great again](https://kerkour.com/post/logging/)

[![GoDoc](https://godoc.org/github.com/bloom42/rz-go?status.svg)](https://godoc.org/github.com/bloom42/rz-go)
[![Build Status](https://travis-ci.org/bloom42/rz-go.svg?branch=master)](https://travis-ci.org/bloom42/rz-go)
[![GitHub release](https://img.shields.io/github/release/bloom42/rz-go.svg)](https://github.com/bloom42/rz-go/releases)

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
SetFormatter(formatter rz.Formatter) // default to rz.JSONFormatter
SetFields(fields ...interface{})
AddFields(fields ...interface{})
SetInsertTimestampField(insert bool) // default to true
SetLevel(level Level) // default to rz.DebugLevel
SetTimestampFieldName(fieldName string) // default to rz.TimestampFieldName ("timestamp")
SetLevelFieldName(fieldName string) // default to rz.LevelFieldName ("level")
SetTimestampFunc(fn func() time.Time) // default to time.Now().UTC
AddHook(hook rz.Hook)
```


## Examples

See the [examples](https://github.com/bloom42/rz-go/tree/master/examples) folder.


## Contributing

See [https://opensource.bloom.sh/contributing](https://opensource.bloom.sh/contributing)


## License

See `LICENSE.txt` and [https://opensource.bloom.sh/licensing](https://opensource.bloom.sh/licensing)
