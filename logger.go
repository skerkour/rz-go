package rz

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// A Logger represents an active logging object that generates lines
// of JSON output to an io.Writer. Each logging operation makes a single
// call to the Writer's Write method. There is no guaranty on access
// serialization to the Writer. If your Writer is not thread safe,
// you may consider a sync wrapper.
type Logger struct {
	writer               LevelWriter
	stack                bool
	caller               bool
	timestamp            bool
	level                LogLevel
	sampler              LogSampler
	context              []byte
	hooks                []LogHook
	timestampFieldName   string
	levelFieldName       string
	messageFieldName     string
	errorFieldName       string
	callerFieldName      string
	callerSkipFrameCount int
	errorStackFieldName  string
	timeFieldFormat      string
	formatter            LogFormatter
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
		writer:               levelWriterAdapter{os.Stdout},
		timestamp:            true,
		timestampFieldName:   DefaultTimestampFieldName,
		levelFieldName:       DefaultLevelFieldName,
		messageFieldName:     DefaultMessageFieldName,
		errorFieldName:       DefaultErrorFieldName,
		callerFieldName:      DefaultCallerFieldName,
		callerSkipFrameCount: DefaultCallerSkipFrameCount,
		errorStackFieldName:  DefaultErrorStackFieldName,
		timeFieldFormat:      DefaultTimeFieldFormat,
	}
	return logger.Config(options...)
}

// Nop returns a disabled logger for which all operation are no-op.
func Nop() Logger {
	return New(Writer(nil), Level(Disabled))
}

// Config apply all the options to the logger
func (l Logger) Config(options ...Option) Logger {
	context := l.context
	l.context = make([]byte, 0, 500)
	if context != nil {
		l.context = append(l.context, context...)
	}
	for _, option := range options {
		option(&l)
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
	e.ch = l.hooks
	e.stack = l.stack
	e.caller = l.caller
	e.timestamp = l.timestamp
	e.timestampFieldName = l.timestampFieldName
	e.levelFieldName = l.levelFieldName
	e.messageFieldName = l.messageFieldName
	e.errorFieldName = l.errorFieldName
	e.callerFieldName = l.callerFieldName
	e.errorStackFieldName = l.errorStackFieldName
	e.timeFieldFormat = l.timeFieldFormat
	if level != NoLevel {
		e.String(l.levelFieldName, level.String())
	}
	if l.context != nil && len(l.context) > 0 {
		e.buf = enc.AppendObjectData(e.buf, l.context)
	}

	if fields != nil {
		fields(e)
	}

	l.writeEvent(e, message, done)
}

func (l *Logger) writeEvent(e *Event, msg string, done func(string)) {
	if len(e.ch) > 0 {
		e.ch[0].Run(e, e.level, msg)
		if len(e.ch) > 1 {
			for _, hook := range e.ch[1:] {
				hook.Run(e, e.level, msg)
			}
		}
	}

	if e.timestamp {
		e.buf = enc.AppendTime(enc.AppendKey(e.buf, e.timestampFieldName), TimestampFunc(), e.timeFieldFormat)
	}
	if msg != "" {
		e.buf = enc.AppendString(enc.AppendKey(e.buf, e.messageFieldName), msg)
	}

	if e.caller {
		_, file, line, ok := runtime.Caller(l.callerSkipFrameCount)
		if ok {
			e.buf = enc.AppendString(enc.AppendKey(e.buf, e.callerFieldName), file+":"+strconv.Itoa(line))
		}
	}
	if done != nil {
		defer done(msg)
	}

	var err error

	if e.level != Disabled {
		e.buf = enc.AppendEndMarker(e.buf)
		e.buf = enc.AppendLineBreak(e.buf)
		if l.formatter != nil {
			e.buf, err = l.formatter(e)
		}
		if e.w != nil {
			_, err = e.w.WriteLevel(e.level, e.buf)
		}
	}

	putEvent(e)

	if err != nil {
		if ErrorHandler != nil {
			ErrorHandler(err)
		} else {
			fmt.Fprintf(os.Stderr, "rz: could not write event: %v\n", err)
		}
	}
}

// should returns true if the log event should be logged.
func (l *Logger) should(lvl LogLevel) bool {
	if lvl < l.level {
		return false
	}
	// if l.sampler != nil && !samplingDisabled() {
	// 	return l.sampler.Sample(lvl)
	// }
	return true
}
