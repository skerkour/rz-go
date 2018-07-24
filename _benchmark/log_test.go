package bench

import (
	"errors"
	"io/ioutil"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/astroflow/astro-go"
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
	b.Run("astroflow/astro-go", func(b *testing.B) {
		logger := newAstro()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(_testMessage)
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
	b.Run("astroflow/astro-go", func(b *testing.B) {
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
	b.Run("astroflow/astro-go", func(b *testing.B) {
		logger := newAstro()
		fields := fakeAstroFields()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.With(fields...).Info(_testMessage)
			}
		})
	})
}
