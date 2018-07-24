package astro

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/json-iterator/go"
)

type LogfmtFormatter struct{}

func (formatter LogfmtFormatter) Format(event Event) []byte {
	var ret = new(bytes.Buffer)
	fields := make([]string, 0, len(event))
	for field := range event {
		fields = append(fields, field)
	}

	sort.Strings(fields)
	for _, field := range fields {
		if needsQuote(field) {
			field = strconv.Quote(field)
		}
		fmt.Fprintf(ret, " %s=", field)

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
	ret.WriteByte('\n')

	return ret.Bytes()
}

func needsQuote(s string) bool {
	for i := range s {
		if s[i] < 0x20 || s[i] > 0x7e || s[i] == ' ' || s[i] == '\\' || s[i] == '"' {
			return true
		}
	}
	return false
}
