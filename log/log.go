// Package log provides a global logger for  astro.
package log

import (
	"context"
	"io"
	"os"

	"github.com/bloom42/astro-go"
)

// Logger is the global logger.
var Logger = astro.New(os.Stderr).With().Timestamp().Logger()

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) astro.Logger {
	return Logger.Output(w)
}

// With creates a child logger with the field added to its context.
func With() astro.Context {
	return Logger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func Level(level astro.Level) astro.Logger {
	return Logger.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s astro.Sampler) astro.Logger {
	return Logger.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h astro.Hook) astro.Logger {
	return Logger.Hook(h)
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *astro.Event {
	return Logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *astro.Event {
	return Logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *astro.Event {
	return Logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *astro.Event {
	return Logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *astro.Event {
	return Logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *astro.Event {
	return Logger.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level astro.Level) *astro.Event {
	return Logger.WithLevel(level)
}

// Log starts a new message with no level. Setting  astro.GlobalLevel to
//  astro.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *astro.Event {
	return Logger.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	Logger.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *astro.Logger {
	return astro.Ctx(ctx)
}
