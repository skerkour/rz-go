package bench

import (
	"errors"
	"io/ioutil"
	"testing"
	"time"

	zl "github.com/astrolib/zerolog"
	"github.com/bloom42/astro-go"
	"github.com/bloom42/astro-go"
	"github.com/sirupsen/logrus"
)

func newDisabledLogrus() *logrus.Logger {
	logger := newLogrus()
	logger.Level = logrus.ErrorLevel
	return logger
}

func newLogrus() *logrus.Logger {
	return &logrus.Logger{
		Out:       ioutil.Discard,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}

func newAstro() astro.Logger {
	logger := astro.NewLogger(
		astro.SetWriter(ioutil.Discard),
	)
	return logger
}

func newZerolog() zerolog.Logger {
	return zerolog.New(ioutil.Discard).With().Timestamp().Logger()
}

func newZl() zl.Logger {
	return zl.New(ioutil.Discard).With().Timestamp().Logger()
}

func newDisabledZerolog() zerolog.Logger {
	return newZerolog().Level(zerolog.Disabled)
}

var _tenInts = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var _tenStrings = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var _tenTimes = []time.Time{time.Now()}

type _testUser struct {
	Username string
	Name     string
	Phone    string
}

var _oneUser = _testUser{
	Username: "lol",
	Name:     "lol2",
	Phone:    "lollol",
}

var _tenUsers = []_testUser{_testUser{}, _testUser{}, _testUser{}, _testUser{}, _testUser{},
	_testUser{}, _testUser{}, _testUser{}, _testUser{}, _testUser{}}
var errExample = errors.New("lolerror")

func fakeLogrusFields() logrus.Fields {
	return logrus.Fields{
		"int":     _tenInts[0],
		"ints":    _tenInts,
		"string":  _tenStrings[0],
		"strings": _tenStrings,
		"time":    _tenTimes[0],
		"times":   _tenTimes,
		"user1":   _oneUser,
		"user2":   _oneUser,
		"users":   _tenUsers,
		"error":   errExample,
	}
}

func fakeAstroFields() []interface{} {
	return []interface{}{
		"int", _tenInts[0],
		"ints", _tenInts,
		"string", _tenStrings[0],
		"strings", _tenStrings,
		"time", _tenTimes[0],
		"times", _tenTimes,
		"user1", _oneUser,
		"user2", _oneUser,
		"users", _tenUsers,
		"error", errExample,
	}
}

func fakeZerologFields(e *zerolog.Event) *zerolog.Event {
	return e.
		Int("int", _tenInts[0]).
		Ints("ints", _tenInts).
		Str("string", _tenStrings[0]).
		Strs("strings", _tenStrings).
		Time("time", _tenTimes[0]).
		Times("times", _tenTimes).
		Interface("user1", _oneUser).
		Interface("user2", _oneUser).
		Interface("users", _tenUsers).
		Err(errExample)
}

func fakeZerologContext(c zerolog.Context) zerolog.Context {
	return c.
		Int("int", _tenInts[0]).
		Ints("ints", _tenInts).
		Str("string", _tenStrings[0]).
		Strs("strings", _tenStrings).
		Time("time", _tenTimes[0]).
		Times("times", _tenTimes).
		Interface("user1", _oneUser).
		Interface("user2", _oneUser).
		Interface("users", _tenUsers).
		Err(errExample)
}

var _testMessage = "hello world"

func BenchmarkWithoutFields(b *testing.B) {
	b.Logf("Logging at a disabled level without any structured context.")
	b.Run("sirupsen/logrus", func(b *testing.B) {
		logger := newLogrus()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(_testMessage)
			}
		})
	})
	b.Run("bloom42/astro-go", func(b *testing.B) {
		logger := newAstro()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(_testMessage)
			}
		})
	})
	b.Run("rs/zerolog", func(b *testing.B) {
		logger := newDisabledZerolog()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info().Msg(_testMessage)
			}
		})
	})
}

func Benchmark10FieldsContext(b *testing.B) {
	b.Run("sirupsen/logrus", func(b *testing.B) {
		logger := newLogrus()
		fields := fakeLogrusFields()
		l := logger.WithFields(fields)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				l.Info(_testMessage)
			}
		})
	})
	b.Run("bloom42/astro-go", func(b *testing.B) {
		logger := newAstro()
		fields := fakeAstroFields()
		l := logger.With(fields...)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				l.Info(_testMessage)
			}
		})
	})
	b.Run("rs/zerolog", func(b *testing.B) {
		logger := fakeZerologContext(newDisabledZerolog().With()).Logger()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info().Msg(_testMessage)
			}
		})
	})
}

func Benchmark10Fields(b *testing.B) {
	b.Run("sirupsen/logrus", func(b *testing.B) {
		logger := newLogrus()
		fields := fakeLogrusFields()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.WithFields(fields).Info(_testMessage)
			}
		})
	})
	b.Run("bloom42/astro-go", func(b *testing.B) {
		logger := newAstro()
		fields := fakeAstroFields()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.With(fields...).Info(_testMessage)
			}
		})
	})
	b.Run("rs/zerolog", func(b *testing.B) {
		logger := newDisabledZerolog()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				fakeZerologFields(logger.Info()).Msg(_testMessage)
			}
		})
	})
}

func BenchmarkZl(b *testing.B) {
	b.Run("rs/zerolog", func(b *testing.B) {
		logger := newZerolog()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Warn().
					Str("Hello", "world").
					Str("Hello2", "world").
					Str("Hello3", "world").
					Str("Hello4", "world").
					Msg(_testMessage)
			}
		})
	})
	b.Run("astrolib/zl", func(b *testing.B) {
		logger := newZl()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Warn(_testMessage, func(e *zl.Event) {
					e.Str("Hello", "world")
					e.Str("Hello2", "world")
					e.Str("Hello3", "world")
					e.Str("Hello4", "world")
				})
			}
		})
	})
}

func BenchmarkZlNoFields(b *testing.B) {
	b.Run("rs/zerolog", func(b *testing.B) {
		logger := newZerolog()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Warn().
					Msg(_testMessage)
			}
		})
	})
	b.Run("astrolib/zl", func(b *testing.B) {
		logger := newZl()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Warn(_testMessage, nil)
			}
		})
	})
}

func BenchmarkZlNoFieldsNoMessage(b *testing.B) {
	b.Run("rs/zerolog", func(b *testing.B) {
		logger := newZerolog()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Warn().
					Msg("")
			}
		})
	})
	b.Run("astrolib/zl", func(b *testing.B) {
		logger := newZl()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Warn("", nil)
			}
		})
	})
}

func BenchmarkZlLotOfFields(b *testing.B) {
	b.Run("rs/zerolog", func(b *testing.B) {
		logger := newZerolog()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Warn().
					Int("int", _tenInts[0]).
					Ints("ints", _tenInts).
					Str("string", _tenStrings[0]).
					Strs("strings", _tenStrings).
					Time("time", _tenTimes[0]).
					Times("times", _tenTimes).
					Interface("user1", _oneUser).
					Interface("user2", _oneUser).
					Interface("users", _tenUsers).
					Err(errExample).
					Msg(_testMessage)
			}
		})
	})
	b.Run("astrolib/zl", func(b *testing.B) {
		logger := newZl()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Warn(_testMessage, func(e *zl.Event) {
					e.Int("int", _tenInts[0])
					e.Ints("ints", _tenInts)
					e.Str("string", _tenStrings[0])
					e.Strs("strings", _tenStrings)
					e.Time("time", _tenTimes[0])
					e.Times("times", _tenTimes)
					e.Interface("user1", _oneUser)
					e.Interface("user2", _oneUser)
					e.Interface("users", _tenUsers)
					e.Err(errExample)
				})
			}
		})
	})
}
