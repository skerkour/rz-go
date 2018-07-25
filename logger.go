package astroflow

import (
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
}

func getTIme() time.Time {
	return time.Now().UTC()
}

type LoggerOption func(logger *Logger) error

func NewLogger(options ...LoggerOption) Logger {
	logger := Logger{
		hooks:                make([]Hook, 0),
		fields:               make([]interface{}, 0),
		writer:               os.Stdout,
		insertTimestampField: true,
		timestampFieldName:   TimestampFieldName,
		timestampFunc:        getTIme,
		messageFieldName:     MessageFieldName,
		levelFieldName:       LevelFieldName,
		level:                DebugLevel,
		formatter:            JSONFormatter{},
	}

	logger.Config(options...)
	return logger
}

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

	if logger.insertTimestampField {
		data[logger.timestampFieldName] = logger.timestampFunc()
	}

	if len(message) != 0 {
		data[logger.messageFieldName] = message
	}

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

func (logger *Logger) With(fields ...interface{}) Logger {
	newLogger := *logger

	newLogger.fields = append(logger.fields, fields...)

	return newLogger
}

func (logger Logger) Debug(message string) {
	logger.log(DebugLevel, message)
}

func (logger Logger) Info(message string) {
	logger.log(InfoLevel, message)
}

func (logger Logger) Warn(message string) {
	logger.log(WarnLevel, message)
}

func (logger Logger) Error(message string) {
	logger.log(ErrorLevel, message)
}

func (logger Logger) Fatal(message string) {
	logger.log(FatalLevel, message)
	os.Exit(1)
}
