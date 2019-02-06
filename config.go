package astro

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

// // With replaces logger's fields
// func With(f func(*Event)) Option {
// 	return func(logger *Logger) {
// 		l := logger.With(f)
// 		*(&logger) = &l
// 	}
// }

// Stack enable/disable stack in error messages
func Stack(enableStack bool) Option {
	return func(logger *Logger) {
		logger.stack = enableStack
	}
}
