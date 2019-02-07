package rz

import (
	"io"
	"os"
)

// Option is used to configure a logger.
type Option func(logger *Logger)

// Writer update logger's writer.
func Writer(writer io.Writer) Option {
	return func(logger *Logger) {
		if writer == nil {
			writer = os.Stdout
		}
		lw, ok := writer.(LevelWriter)
		if !ok {
			lw = levelWriterAdapter{writer}
		}
		logger.writer = lw
	}
}

// Level update logger's level.
func Level(lvl LogLevel) Option {
	return func(logger *Logger) {
		logger.level = lvl
	}
}

// Sampler update logger's sampler.
func Sampler(sampler LogSampler) Option {
	return func(logger *Logger) {
		logger.sampler = sampler
	}
}

// AddHook appends hook to logger's hook
func AddHook(hook LogHook) Option {
	return func(logger *Logger) {
		logger.hooks = append(logger.hooks, hook)
	}
}

// Hooks replaces logger's hooks
func Hooks(hooks ...LogHook) Option {
	return func(logger *Logger) {
		logger.hooks = hooks
	}
}

// With replaces logger's configuration
func With(fields func(*Event)) Option {
	return func(logger *Logger) {
		if fields != nil {
			e := newEvent(logger.writer, logger.level)
			e.buf = nil
			fields(e)
			logger.stack = e.stack
			logger.caller = e.caller
			logger.timestamp = e.timestamp
			logger.context = enc.AppendObjectData(make([]byte, 0, 500), e.buf)
		}
	}
}

// Stack enable/disable stack in error messages
func Stack(enableStack bool) Option {
	return func(logger *Logger) {
		logger.stack = enableStack
	}
}

// Timestamp enable/disable timestamp logging in error messages
func Timestamp(enableTimestamp bool) Option {
	return func(logger *Logger) {
		logger.timestamp = enableTimestamp
	}
}

// Formatter update logger's formatter
func Formatter(formatter LogFormatter) Option {
	return func(logger *Logger) {
		logger.formatter = formatter
	}
}

// TimestampFieldName update logger's timestampFieldName
func TimestampFieldName(timestampFieldName string) Option {
	return func(logger *Logger) {
		logger.timestampFieldName = timestampFieldName
	}
}

// LevelFieldName update logger's levelFieldName
func LevelFieldName(levelFieldName string) Option {
	return func(logger *Logger) {
		logger.levelFieldName = levelFieldName
	}
}

// MessageFieldName update logger's messageFieldName
func MessageFieldName(messageFieldName string) Option {
	return func(logger *Logger) {
		logger.messageFieldName = messageFieldName
	}
}

// ErrorFieldName update logger's errorFieldName
func ErrorFieldName(errorFieldName string) Option {
	return func(logger *Logger) {
		logger.errorFieldName = errorFieldName
	}
}

// CallerFieldName update logger's callerFieldName
func CallerFieldName(callerFieldName string) Option {
	return func(logger *Logger) {
		logger.callerFieldName = callerFieldName
	}
}

// CallerSkipFrameCount update logger's callerSkipFrameCount
func CallerSkipFrameCount(callerSkipFrameCount int) Option {
	return func(logger *Logger) {
		logger.callerSkipFrameCount = callerSkipFrameCount
	}
}

// ErrorStackFieldName update logger's errorStackFieldName
func ErrorStackFieldName(errorStackFieldName string) Option {
	return func(logger *Logger) {
		logger.errorStackFieldName = errorStackFieldName
	}
}
