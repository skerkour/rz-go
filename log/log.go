// Package log provides a global logger for  rz.
package log

import (
	"context"

	"github.com/bloom42/rz-go"
)

// Logger is the global logger.
var Logger = rz.New()

// Config duplicates the global logger and update it's configuration.
func Config(options ...rz.Option) rz.Logger {
	return Logger.Config(options...)
}

// Debug starts a new message with debug level.
func Debug(message string, fields func(*rz.Event)) {
	Logger.Debug(message, fields)
}

// Info logs a new message with info level.
func Info(message string, fields func(*rz.Event)) {
	Logger.Info(message, fields)
}

// Warn logs a new message with warn level.
func Warn(message string, fields func(*rz.Event)) {
	Logger.Warn(message, fields)
}

// Error logs a message with error level.
func Error(message string, fields func(*rz.Event)) {
	Logger.Error(message, fields)
}

// Fatal logs a new message with fatal level. The os.Exit(1) function
// is then called, which terminates the program immediately.
func Fatal(message string, fields func(*rz.Event)) {
	Logger.Fatal(message, fields)
}

// Panic logs a new message with panic level. The panic() function
// is then called, which stops the ordinary flow of a goroutine.
func Panic(message string, fields func(*rz.Event)) {
	Logger.Panic(message, fields)
}

// Log logs a new message with no level. Setting GlobalLevel to Disabled
// will still disable events produced by this method.
func Log(message string, fields func(*rz.Event)) {
	Logger.Log(message, fields)
}

// FromCtx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func FromCtx(ctx context.Context) *rz.Logger {
	return rz.FromCtx(ctx)
}
