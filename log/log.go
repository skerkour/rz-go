package log

import (
	"github.com/astroflow/astroflow-go"
)

var logger = astroflow.NewLogger()

func Config(options ...astroflow.LoggerOption) error {
	return logger.Config(options...)
}

func With(fields ...interface{}) astroflow.Logger {
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

// Msg log an event without level
func Msg(message string) {
	logger.Msg(message)
}

// Track log an event without message nor level
func Track(fields ...interface{}) {
	logger.Track(fields...)
}
