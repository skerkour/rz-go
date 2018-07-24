package astro

import (
	"io"
	"time"
)

func SetWriter(writer io.Writer) LoggerOption {
	return func(logger *Logger) error {
		logger.writer = writer
		return nil
	}
}

func SetFormatter(formatter Formatter) LoggerOption {
	return func(logger *Logger) error {
		logger.formatter = formatter
		return nil
	}
}

func SetWith(fields ...interface{}) LoggerOption {
	return func(logger *Logger) error {
		logger.fields = fields
		return nil
	}
}

func SetInsertTimestampField(insert bool) LoggerOption {
	return func(logger *Logger) error {
		logger.insertTimestampField = insert
		return nil
	}
}

func SetLevel(level Level) LoggerOption {
	return func(logger *Logger) error {
		logger.level = level
		return nil
	}
}

func SetTimestampFieldName(fieldName string) LoggerOption {
	return func(logger *Logger) error {
		logger.timestampFieldName = fieldName
		return nil
	}
}

func SetMessageFieldName(fieldName string) LoggerOption {
	return func(logger *Logger) error {
		logger.messageFieldName = fieldName
		return nil
	}
}

func SetLevelFieldName(fieldName string) LoggerOption {
	return func(logger *Logger) error {
		logger.levelFieldName = fieldName
		return nil
	}
}

func SetTimestampFunc(fn func() time.Time) LoggerOption {
	return func(logger *Logger) error {
		logger.timestampFunc = fn
		return nil
	}
}

func AddHook(hook Hook) LoggerOption {
	return func(logger *Logger) error {
		logger.hooks = append(logger.hooks, hook)
		return nil
	}
}
