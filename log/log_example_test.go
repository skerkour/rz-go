// +build !binary_log

package log_test

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/bloom42/rz-go"
	"github.com/bloom42/rz-go/log"
)

// setup would normally be an init() function, however, there seems
// to be something awry with the testing framework when we set the
// global Logger from an init()
func setup() {
	// UNIX Time is faster and smaller than most timestamps
	// If you set rz.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	rz.TimeFieldFormat = ""
	// In order to always output a static time to stdout for these
	// examples to pass, we need to override rz.TimestampFunc
	// and log.Logger globals -- you would not normally need to do this
	rz.TimestampFunc = func() time.Time {
		return time.Date(2008, 1, 8, 17, 5, 05, 0, time.UTC)
	}
	log.Logger = rz.New()
}

// Example of a log with no particular "level"
func ExampleLog() {
	setup()
	log.Log("hello world", nil)

	// Output: {"time":1199811905,"message":"hello world"}
}

// Example of a log at a particular "level" (in this case, "debug")
func ExampleDebug() {
	setup()
	log.Debug("hello world", nil)

	// Output: {"level":"debug","time":1199811905,"message":"hello world"}
}

// Example of a log at a particular "level" (in this case, "info")
func ExampleInfo() {
	setup()
	log.Info("hello world", nil)

	// Output: {"level":"info","time":1199811905,"message":"hello world"}
}

// Example of a log at a particular "level" (in this case, "warn")
func ExampleWarn() {
	setup()
	log.Warn("hello world", nil)

	// Output: {"level":"warn","time":1199811905,"message":"hello world"}
}

// Example of a log at a particular "level" (in this case, "error")
func ExampleError() {
	setup()
	log.Error("hello world", nil)

	// Output: {"level":"error","time":1199811905,"message":"hello world"}
}

// Example of a log at a particular "level" (in this case, "fatal")
func ExampleFatal() {
	setup()
	err := errors.New("A repo man spends his life getting into tense situations")
	service := "myservice"

	log.Fatal(fmt.Sprintf("Cannot start %s", service), func(e *rz.Event) {
		e.Err(err).
			String("service", service)
	})

	// Outputs: {"level":"fatal","time":1199811905,"error":"A repo man spends his life getting into tense situations","service":"myservice","message":"Cannot start myservice"}
}

// TODO: Panic

// This example uses command-line flags to demonstrate various outputs
// depending on the chosen log level.
func Example() {
	setup()
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	rz.SetGlobalLevel(rz.InfoLevel)
	if *debug {
		rz.SetGlobalLevel(rz.DebugLevel)
	}

	log.Debug("This message appears only when log level set to Debug", nil)
	log.Info("This message appears when log level set to Debug or Info", nil)

	// if e := log.Debug(); e.Enabled() {
	// 	// Compute log output only if enabled.
	// 	value := "bar"
	// 	e.Str("foo", value).Msg("some debug message")
	// }

	// Output: {"level":"info","time":1199811905,"message":"This message appears when log level set to Debug or Info"}
}

// TODO: Output

// TODO: With

// TODO: Level

// TODO: Sample

// TODO: Hook

// TODO: WithLevel

// TODO: Ctx
