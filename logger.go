package astro

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	TimestampFieldName = "timestamp"
	MessageFieldName   = "message"
	LevelFieldName     = "level"
)

type Hook func(event Event)

type Logger struct {
	hooks                []Hook
	fields               []interface{}
	writer               io.Writer
	insertTimestampField bool
	timestampFieldName   string
	timestampFunc        func() time.Time
	messageFieldName     string
	levelFieldName       string
	level                Level
	formatter            Formatter
	callerLevel          int
	insertCaller         bool
}

func getTIme() time.Time {
	return time.Now().UTC()
}

type LoggerOption func(logger *Logger) error

// NewLogger returns a new logger with default configuration. Additional can be provided
func NewLogger(options ...LoggerOption) Logger {
	logger := Logger{
		hooks:                make([]Hook, 0),
		fields:               make([]interface{}, 0),
		writer:               StdoutWriter{},
		insertTimestampField: true,
		timestampFieldName:   TimestampFieldName,
		timestampFunc:        getTIme,
		messageFieldName:     MessageFieldName,
		levelFieldName:       LevelFieldName,
		level:                DebugLevel,
		formatter:            JSONFormatter{},
		callerLevel:          3,
		insertCaller:         true,
	}

	logger.Config(options...)
	return logger
}

// Config configure the logger
func (logger *Logger) Config(options ...LoggerOption) error {
	var err error

	for _, option := range options {
		err = option(logger)
		if err != nil {
			return err
		}
	}

	return nil
}

func (logger *Logger) log(level Level, message string) {
	if level < logger.level {
		return
	}

	n := 2
	if logger.insertTimestampField {
		n += 1
	}

	fieldsLen := len(logger.fields)
	data := make(Event, (fieldsLen/2)+n)
	for i := 0; i < fieldsLen; i += 2 {
		data[logger.fields[i].(string)] = logger.fields[i+1]
	}

	caller, err := caller(logger.callerLevel)
	if err == nil {
		data["caller"] = caller
	}

	if logger.insertTimestampField {
		data[logger.timestampFieldName] = logger.timestampFunc()
	}

	if len(message) != 0 {
		data[logger.messageFieldName] = message
	}

	// default case: do not insert level field
	switch level {
	case DebugLevel:
		data[logger.levelFieldName] = "debug"
	case InfoLevel:
		data[logger.levelFieldName] = "info"
	case WarnLevel:
		data[logger.levelFieldName] = "warning"
	case ErrorLevel:
		data[logger.levelFieldName] = "error"
	case FatalLevel:
		data[logger.levelFieldName] = "fatal"
	}

	for _, hook := range logger.hooks {
		hook(data)
	}
	bytes := logger.formatter.Format(data)
	logger.writer.Write(bytes)
}

// With returns a new Logger with the provided fields added
func (logger Logger) With(fields ...interface{}) Logger {
	logger.fields = append(logger.fields, fields...)
	return logger
}

// Debug level message
func (logger Logger) Debug(message string) {
	logger.log(DebugLevel, message)
}

// Debugf level formatted message
func (logger Logger) Debugf(format string, a ...interface{}) {
	logger.log(DebugLevel, fmt.Sprintf(format, a...))
}

// Info level message
func (logger Logger) Info(message string) {
	logger.log(InfoLevel, message)
}

// Infof level formatted message
func (logger Logger) Infof(format string, a ...interface{}) {
	logger.log(InfoLevel, fmt.Sprintf(format, a...))
}

// Warn warning level message
func (logger Logger) Warn(message string) {
	logger.log(WarnLevel, message)
}

// Warnf warning formatted message
func (logger Logger) Warnf(format string, a ...interface{}) {
	logger.log(WarnLevel, fmt.Sprintf(format, a...))
}

// Error level message
func (logger Logger) Error(message string) {
	logger.log(ErrorLevel, message)
}

// Errorf error formatted message
func (logger Logger) Errorf(format string, a ...interface{}) {
	logger.log(ErrorLevel, fmt.Sprintf(format, a...))
}

// Fatal message, followed by exit(1)
func (logger Logger) Fatal(message string) {
	logger.log(FatalLevel, message)
	os.Exit(1)
}

// Fatalf fatal formatted message, followed by exit(1)
func (logger Logger) Fatalf(format string, a ...interface{}) {
	logger.log(FatalLevel, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Track an event without message nor level
func (logger Logger) Track(fields ...interface{}) {
	newLogger := logger.With(fields...)
	newLogger.log(NoneLevel, "")
}
