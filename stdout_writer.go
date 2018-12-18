package astro

import (
	"os"
)

type StdoutWriter struct{}

func (writer StdoutWriter) Write(bytes []byte) (int, error) {
	n, err := os.Stdout.Write(bytes)
	if err != nil {
		return n, err
	}
	nn, err := os.Stdout.Write([]byte("\n"))
	return n + nn, err
}
