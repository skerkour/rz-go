## Astro

```go

import (
    "os"

    "github.com/astroflow/astro-go"
    "github.com/astroflow/astro-go/log"
)


func main() {

    //
    // log.Init(
    //     astro.SetWriter(os.Stderr),
    // )

    if os.Getenv("GO_ENV") == "pduction" {
        log.Config(
            astro.SetFormatter(astro.JSONFormatter{}),
            astro.SetLevel(astro.InfoLevel),
        )
    }

}

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
