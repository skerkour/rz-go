package log

import (
	"github.com/astroflow/astroflow-go"
)

var logger = astroflow.NewLogger()

func Logger() astroflow.Logger {
	return logger
}

func Init(options ...astroflow.LoggerOption) error {
	return logger.Config(options...)
}

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
