package log

import (
	"github.com/bloom42/astro-go"
)

var logger = astro.NewLogger(astro.SetCallerLevel(4))

// Config configure the logger
func Config(options ...astro.LoggerOption) error {
	return logger.Config(options...)
}

// With returns a new Logger with the provided fields added
func With(fields ...interface{}) astro.Logger {
	ret := logger.With(fields...)
	ret.Config(astro.SetCallerLevel(5))
	return ret
}

// Debug level message
func Debug(message string) {
	logger.Debug(message)
}

// Debugf level formatted message
func Debugf(format string, a ...interface{}) {
	logger.Debugf(format, a...)
}

// Info level message
func Info(message string) {
	logger.Info(message)
}

// Infof level formatted message
func Infof(format string, a ...interface{}) {
	logger.Infof(format, a...)
}

// Warn warning level message
func Warn(message string) {
	logger.Warn(message)
}

// Warnf warning formatted message
func Warnf(format string, a ...interface{}) {
	logger.Warnf(format, a...)
}

// Error level message
func Error(message string) {
	logger.Error(message)
}

// Errorf error formatted message
func Errorf(format string, a ...interface{}) {
	logger.Errorf(format, a...)
}

// Fatal message, followed by exit(1)
func Fatal(message string) {
	logger.Fatal(message)
}

// Fatalf fatal formatted message, followed by exit(1)
func Fatalf(format string, a ...interface{}) {
	logger.Fatalf(format, a...)
}

// Track log an event without message nor level
func Track(fields ...interface{}) {
	logger.Track(fields...)
}
