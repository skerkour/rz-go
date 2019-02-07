package rz

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ConsoleFormatter(ev *Event) ([]byte, error) {
	var evt map[string]interface{}
	var b bytes.Buffer

	d := json.NewDecoder(bytes.NewReader(ev.buf))
	d.UseNumber()
	err := d.Decode(&evt)
	if err != nil {
		return b.Bytes(), fmt.Errorf("cannot decode event: %s", err)
	}
	fmt.Fprintf(&b, "[%v] %s\n", evt[ev.timestampFieldName], evt[ev.messageFieldName])
	return b.Bytes(), nil
}
