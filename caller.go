package astro

import (
	"errors"
	"fmt"
	"runtime"
)

func caller(level int) (string, error) {

	_, file, line, ok := runtime.Caller(level)
	if ok {
		return fmt.Sprintf("%s:%d", file, line), nil
	}
	return "", errors.New("error getting caller")
}
