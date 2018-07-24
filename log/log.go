package log

import (
	"github.com/astroflow/astro-go"
)

var logger = astro.NewLogger()

func Init(options ...astro.LoggerOption) error {
	return logger.Config(options...)
}

func Config(options ...astro.LoggerOption) error {
	return logger.Config(options...)
}

func With(fields ...interface{}) astro.Logger {
	return logger.With(fields...)
}

func Debug(message string) {
	logger.Debug(message)
}

func Info(message string) {
	logger.Info(message)
}

func Warn(message string) {
	logger.Warn(message)
}

func Error(message string) {
	logger.Error(message)
}

func Fatal(message string) {
	logger.Fatal(message)
}

/*

// Msg log a message without level
func Msg(message string) {
	logger.Msg(message)
}

// Track log an event without message and level
func Track() {
	logger.Track()
}
*/
