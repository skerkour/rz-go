package astro

import (
	"github.com/json-iterator/go"
)

type JSONFormatter struct{}

func (formatter JSONFormatter) Format(entry Event) []byte {
	serialized, err := jsoniter.Marshal(entry)
	if err != nil {
		return nil
	}
	return serialized
}
