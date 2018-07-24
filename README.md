## Astro

```go

import (
    "os"

    "github.com/astroflow/astro-go"
    "github.com/astroflow/astro-go/log"
)


func main() {

    // log.Init(
    //     astro.SetWriter(os.Stderr),
    // )

    if os.Getenv("GO_ENV") == "production" {
        log.Config(
            astro.SetFormatter(astro.JSONFormatter{}),
            astro.SetLevel(astro.InfoLevel),
        )
    }

}

```

## Log
```go
import (
    "github.com/astroflow/astro-go/log"
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
SetFormatter(formatter astro.Formatter) // default to astro.ConsoleFormatter
SetWith(fields ...interface{})
SetInsertTimestampField(insert bool) // default to true
SetLevel(level Level) // default to astro.DebugLevel
SetTimestampFieldName(fieldName string) // default to astro.TimestampFieldName ("timestamp")
SetLevelFieldName(fieldName string) // default to astro.LevelFieldName ("level")
SetTimestampFunc(fn func() time.Time) // default to time.Now().UTC
AddHook(hook astro.Hook)
```
