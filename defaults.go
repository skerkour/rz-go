package rz

import "time"

var (
	// DefaultTimestampFieldName is the default field name used for the timestamp field.
	DefaultTimestampFieldName = "timestamp"

	// DefaultLevelFieldName is the default field name used for the level field.
	DefaultLevelFieldName = "level"

	// DefaultMessageFieldName is the default field name used for the message field.
	DefaultMessageFieldName = "message"

	// DefaultErrorFieldName is the default field name used for error fields.
	DefaultErrorFieldName = "error"

	// DefaultCallerFieldName is the default field name used for caller field.
	DefaultCallerFieldName = "caller"

	// DefaultCallerSkipFrameCount is the default number of stack frames to skip to find the caller.
	DefaultCallerSkipFrameCount = 3

	// DefaultErrorStackFieldName is the default field name used for error stacks.
	DefaultErrorStackFieldName = "stack"

	// ErrorStackMarshaler extract the stack from err if any.
	ErrorStackMarshaler func(err error) interface{}

	// ErrorMarshalFunc allows customization of global error marshaling
	ErrorMarshalFunc = func(err error) interface{} {
		return err
	}

	// DefaultTimeFieldFormat defines the time format of the Time field type.
	// If set to an empty string, the time is formatted as an UNIX timestamp
	// as integer.
	DefaultTimeFieldFormat = time.RFC3339

	// TimestampFunc defines the function called to generate a timestamp.
	DefaultTimestampFunc = time.Now

	// DurationFieldUnit defines the unit for time.Duration type fields added
	// using the Dur method.
	DurationFieldUnit = time.Microsecond

	// DurationFieldInteger renders Dur fields as integer instead of float if
	// set to true.
	DurationFieldInteger = false

	// ErrorHandler is called whenever zerolog fails to write an event on its
	// output. If not set, an error is printed on the stderr. This handler must
	// be thread safe and non-blocking.
	ErrorHandler func(err error)
)

// var (
// 	gLevel          = new(uint32)
// 	disableSampling = new(uint32)
// )

// TODO: remove
// SetGlobalLevel sets the global override for log level. If this
// // values is raised, all Loggers will use at least this value.
// //
// // To globally disable logs, set GlobalLevel to Disabled.
// func SetGlobalLevel(l LogLevel) {
// 	atomic.StoreUint32(gLevel, uint32(l))
// }

// // GlobalLevel returns the current global log level
// func GlobalLevel() LogLevel {
// 	return LogLevel(atomic.LoadUint32(gLevel))
// }

// // DisableSampling will disable sampling in all Loggers if true.
// func DisableSampling(v bool) {
// 	var i uint32
// 	if v {
// 		i = 1
// 	}
// 	atomic.StoreUint32(disableSampling, i)
// }

// func samplingDisabled() bool {
// 	return atomic.LoadUint32(disableSampling) == 1
// }
