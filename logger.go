// Package zerolog provides a lightweight logging library dedicated to JSON logging.
//
// A global Logger can be use for simple logging:
//
//     import "github.com/bloom42/astro-go/log"
//
//     log.Info().Msg("hello world")
//     // Output: {"time":1494567715,"level":"info","message":"hello world"}
//
// NOTE: To import the global logger, import the "log" subpackage "github.com/bloom42/astro-go/log".
//
// Fields can be added to log messages:
//
//     log.Info().Str("foo", "bar").Msg("hello world")
//     // Output: {"time":1494567715,"level":"info","message":"hello world","foo":"bar"}
//
// Create logger instance to manage different outputs:
//
//     logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
//     logger.Info().
//            Str("foo", "bar").
//            Msg("hello world")
//     // Output: {"time":1494567715,"level":"info","message":"hello world","foo":"bar"}
//
// Sub-loggers let you chain loggers with additional context:
//
//     sublogger := log.With().Str("component": "foo").Logger()
//     sublogger.Info().Msg("hello world")
//     // Output: {"time":1494567715,"level":"info","message":"hello world","component":"foo"}
//
// Level logging
//
//     zerolog.SetGlobalLevel(zerolog.InfoLevel)
//
//     log.Debug().Msg("filtered out message")
//     log.Info().Msg("routed message")
//
//     if e := log.Debug(); e.Enabled() {
//         // Compute log output only if enabled.
//         value := compute()
//         e.Str("foo": value).Msg("some debug message")
//     }
//     // Output: {"level":"info","time":1494567715,"routed message"}
//
// Customize automatic field names:
//
//     log.TimestampFieldName = "t"
//     log.LevelFieldName = "p"
//     log.MessageFieldName = "m"
//
//     log.Info().Msg("hello world")
//     // Output: {"t":1494567715,"p":"info","m":"hello world"}
//
// Log with no level and message:
//
//     log.Log().Str("foo","bar").Msg("")
//     // Output: {"time":1494567715,"foo":"bar"}
//
// Add contextual fields to global Logger:
//
//     log.Logger = log.With().Str("foo", "bar").Logger()
//
// Sample logs:
//
//     sampled := log.Sample(&zerolog.BasicSampler{N: 10})
//     sampled.Info().Msg("will be logged every 10 messages")
//
// Log with contextual hooks:
//
//     // Create the hook:
//     type SeverityHook struct{}
//
//     func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
//          if level != zerolog.NoLevel {
//              e.Str("severity", level.String())
//          }
//     }
//
//     // And use it:
//     var h SeverityHook
//     log := zerolog.New(os.Stdout).Hook(h)
//     log.Warn().Msg("")
//     // Output: {"level":"warn","severity":"warn"}
//
//
// Caveats
//
// There is no fields deduplication out-of-the-box.
// Using the same key multiple times creates new key in final JSON each time.
//
//     logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
//     logger.Info().
//            Timestamp().
//            Msg("dup")
//     // Output: {"level":"info","time":1494567715,"time":1494567715,"message":"dup"}
//
// However, it’s not a big deal though as JSON accepts dup keys,
// the last one prevails.
package astro

import (
	"os"
)

// A Logger represents an active logging object that generates lines
// of JSON output to an io.Writer. Each logging operation makes a single
// call to the Writer's Write method. There is no guaranty on access
// serialization to the Writer. If your Writer is not thread safe,
// you may consider a sync wrapper.
type Logger struct {
	writer  LevelWriter
	stack   bool
	level   LogLevel
	sampler LogSampler
	context []byte
	hooks   []LogHook
}

// New creates a root logger with given options. If the output writer implements
// the LevelWriter interface, the WriteLevel method will be called instead of the Write
// one. Default writer is os.Stdout
//
// Each logging operation makes a single call to the Writer's Write method. There is no
// guaranty on access serialization to the Writer. If your Writer is not thread safe,
// you may consider using sync wrapper.
func New(options ...Option) Logger {
	logger := Logger{
		writer: levelWriterAdapter{os.Stdout},
	}
	return logger.Config(options...)
}

// Nop returns a disabled logger for which all operation are no-op.
func Nop() Logger {
	return New(Writer(nil), Level(Disabled))
}

// Config apply all the options to the logger
func (l Logger) Config(options ...Option) Logger {
	for _, option := range options {
		option(&l)
	}
	return l
}

// With creates a child logger with the field added to its context.
func (l Logger) With(fields func(*Event)) Logger {
	if fields != nil {
		e := newEvent(nil, NoLevel)
		l.context = enc.AppendObjectData(l.context, e.buf)
	}
	return l
}

// Debug logs a new message with debug level.
func (l *Logger) Debug(message string, fields func(*Event)) {
	l.logEvent(DebugLevel, message, fields, nil)
}

// Info logs a new message with info level.
func (l *Logger) Info(message string, fields func(*Event)) {
	l.logEvent(InfoLevel, message, fields, nil)
}

// Warn logs a new message with warn level.
func (l *Logger) Warn(message string, fields func(*Event)) {
	l.logEvent(WarnLevel, message, fields, nil)
}

// Error logs a message with error level.
func (l *Logger) Error(message string, fields func(*Event)) {
	l.logEvent(ErrorLevel, message, fields, nil)
}

// Fatal logs a new message with fatal level. The os.Exit(1) function
// is then called, which terminates the program immediately.
func (l *Logger) Fatal(message string, fields func(*Event)) {
	l.logEvent(FatalLevel, message, fields, func(msg string) { os.Exit(1) })
}

// Panic logs a new message with panic level. The panic() function
// is then called, which stops the ordinary flow of a goroutine.
func (l *Logger) Panic(message string, fields func(*Event)) {
	l.logEvent(PanicLevel, message, fields, func(msg string) { panic(msg) })
}

// Log logs a new message with no level. Setting GlobalLevel to Disabled
// will still disable events produced by this method.
func (l *Logger) Log(message string, fields func(*Event)) {
	l.logEvent(NoLevel, message, fields, nil)
}

// Write implements the io.Writer interface. This is useful to set as a writer
// for the standard library log.
func (l Logger) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 && p[n-1] == '\n' {
		// Trim CR added by stdlog.
		p = p[0 : n-1]
	}
	l.Log(string(p), nil)
	return
}

func (l *Logger) logEvent(level LogLevel, message string, fields func(*Event), done func(string)) {
	enabled := l.should(level)
	if !enabled {
		return
	}
	e := newEvent(l.writer, level)
	e.done = done
	e.ch = l.hooks
	if level != NoLevel {
		e.String(LevelFieldName, level.String())
	}
	if l.context != nil && len(l.context) > 0 {
		e.buf = enc.AppendObjectData(e.buf, l.context)
	}

	if fields != nil {
		fields(e)
	}
	e.msg(message)
}

func (l *Logger) newEvent(level LogLevel, message string, fields func(*Event), done func(string)) {
	enabled := l.should(level)
	if !enabled {
		return
	}
	e := newEvent(l.writer, level)
	e.done = done
	e.ch = l.hooks
	if level != NoLevel {
		e.String(LevelFieldName, level.String())
	}
	if l.context != nil && len(l.context) > 0 {
		e.buf = enc.AppendObjectData(e.buf, l.context)
	}

	if fields != nil {
		fields(e)
	}
	e.msg(message)
}

// should returns true if the log event should be logged.
func (l *Logger) should(lvl LogLevel) bool {
	if lvl < l.level || lvl < GlobalLevel() {
		return false
	}
	if l.sampler != nil && !samplingDisabled() {
		return l.sampler.Sample(lvl)
	}
	return true
}
