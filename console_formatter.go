package astro

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/json-iterator/go"
)

const (
	cReset    = 0
	cBold     = 1
	cRed      = 31
	cGreen    = 32
	cYellow   = 33
	cBlue     = 34
	cMagenta  = 35
	cCyan     = 36
	cGray     = 37
	cDarkGray = 90
)

type ConsoleFormatter struct {
	TimestampFieldName string
	MessageFieldName   string
	LevelFieldName     string
	NoColor            bool
}

func NewConsoleFormatter() ConsoleFormatter {
	return ConsoleFormatter{
		TimestampFieldName: TimestampFieldName,
		MessageFieldName:   MessageFieldName,
		LevelFieldName:     LevelFieldName,
		NoColor:            false,
	}
}

func (formatter ConsoleFormatter) Format(event Event) []byte {
	var ret = new(bytes.Buffer)

	lvlColor := cReset
	level := "????"
	if l, ok := event[formatter.LevelFieldName].(string); ok {
		if !formatter.NoColor {
			lvlColor = levelColor(l)
		}
		level = strings.ToUpper(l)[0:4]
	}

	message := ""
	if m, ok := event[formatter.MessageFieldName].(string); ok {
		message = m
	}

	timestamp := ""
	if t, ok := event[formatter.TimestampFieldName].(time.Time); ok {
		timestamp = t.Format(time.RFC3339)
	}

	if message != "" {
		ret.WriteString(fmt.Sprintf("%-20s |%-4s| %-42s ",
			timestamp,
			colorize(level, lvlColor, !formatter.NoColor),
			message,
		))
	} else {
		ret.WriteString(fmt.Sprintf("%-20s |%-4s|",
			timestamp,
			colorize(level, lvlColor, !formatter.NoColor),
		))
	}

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

func colorize(s interface{}, color int, enabled bool) string {
	if !enabled {
		return fmt.Sprintf("%v", s)
	}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", color, s)
}

func levelColor(level string) int {
	switch level {
	case "debug":
		return cMagenta
	case "info":
		return cCyan
	case "warning":
		return cYellow
	case "error", "fatal":
		return cRed
	default:
		return cReset
	}
}
