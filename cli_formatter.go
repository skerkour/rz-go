package astro

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/json-iterator/go"
)

type CLIFormatter struct {
	NoColor            bool
	TimestampFieldName string
	LevelFieldName     string
	MessageFieldName   string
}

func NewCLIFormatter() CLIFormatter {
	return CLIFormatter{
		TimestampFieldName: TimestampFieldName,
		MessageFieldName:   MessageFieldName,
		LevelFieldName:     LevelFieldName,
		NoColor:            false,
	}
}

func (formatter CLIFormatter) Format(event Event) []byte {
	var ret = new(bytes.Buffer)

	lvlColor := cReset
	level := ""
	if l, ok := event[formatter.LevelFieldName].(string); ok {
		if !formatter.NoColor {
			lvlColor = levelColor(l)
		}
		level = l
	}

	message := ""
	if m, ok := event[formatter.MessageFieldName].(string); ok {
		message = m
	}

	if level != "" {
		ret.WriteString(colorize("• ", lvlColor, !formatter.NoColor))
	}
	ret.WriteString(message)

	fields := make([]string, 0, len(event))
	for field := range event {
		switch field {
		case formatter.TimestampFieldName, formatter.MessageFieldName, formatter.LevelFieldName:
			continue
		}

		fields = append(fields, field)
	}

	sort.Strings(fields)
	for _, field := range fields {
		if needsQuote(field) {
			field = strconv.Quote(field)
		}
		fmt.Fprintf(ret, " %s=", colorize(field, lvlColor, !formatter.NoColor))

		switch value := event[field].(type) {
		case string:
			if len(value) == 0 {
				ret.WriteString("\"\"")
			} else if needsQuote(value) {
				ret.WriteString(strconv.Quote(value))
			} else {
				ret.WriteString(value)
			}
		case time.Time:
			ret.WriteString(value.Format(time.RFC3339))
		default:
			b, err := jsoniter.Marshal(value)
			if err != nil {
				fmt.Fprintf(ret, "[error: %v]", err)
			} else {
				fmt.Fprint(ret, string(b))
			}
		}

	}

	return ret.Bytes()
}
