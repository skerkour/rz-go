package rz

import (
	"net"
	"time"
)

// Field functions are used to add fields to events
type Field func(e *Event)

// Discard disables the event
func Discard() Field {
	return func(e *Event) {
		e.Discard()
	}
}

// Stack enables stack trace printing for the error passed to Err().
//
// logger.errorStackMarshaler must be set for this method to do something.
func Stack(enable bool) Field {
	return func(e *Event) {
		e.Stack(enable)
	}
}

// Caller adds the file:line of the caller with the rz.CallerFieldName key.
func Caller(enable bool) Field {
	return func(e *Event) {
		e.Caller(enable)
	}
}

// Fields is a helper function to use a map to set fields using type assertion.
func Fields(fields map[string]interface{}) Field {
	return func(e *Event) {
		e.Fields(fields)
	}
}

// String adds the field key with val as a string to the *Event context.
func String(key, value string) Field {
	return func(e *Event) {
		e.String(key, value)
	}
}

// Strings adds the field key with vals as a []string to the *Event context.
func Strings(key string, value []string) Field {
	return func(e *Event) {
		e.Strings(key, value)
	}
}

// Time adds the field key with t formated as string using rz.TimeFieldFormat.
func Time(key string, value time.Time) Field {
	return func(e *Event) {
		e.Time(key, value)
	}
}

// Times adds the field key with t formated as string using rz.TimeFieldFormat.
func Times(key string, value []time.Time) Field {
	return func(e *Event) {
		e.Times(key, value)
	}
}

// Duration adds the field key with duration d stored as rz.DurationFieldUnit.
// If rz.DurationFieldInteger is true, durations are rendered as integer
// instead of float.
func Duration(key string, value time.Duration) Field {
	return func(e *Event) {
		e.Duration(key, value)
	}
}

// Durations adds the field key with duration d stored as rz.DurationFieldUnit.
// If rz.DurationFieldInteger is true, durations are rendered as integer
// instead of float.
func Durations(key string, value []time.Duration) Field {
	return func(e *Event) {
		e.Durations(key, value)
	}
}

// Object marshals an object that implement the LogObjectMarshaler interface.
func Object(key string, value LogObjectMarshaler) Field {
	return func(e *Event) {
		e.Object(key, value)
	}
}

// EmbedObject marshals an object that implement the LogObjectMarshaler interface.
func EmbedObject(obj LogObjectMarshaler) Field {
	return func(e *Event) {
		e.EmbedObject(obj)
	}
}

// Dict adds the field key with a dict to the event context.
// Use rz.Dict() to create the dictionary.
func Dict(key string, value *Event) Field {
	return func(e *Event) {
		e.Dict(key, value)
	}
}

// Bytes adds the field key with val as a string to the *Event context.
//
// Runes outside of normal ASCII ranges will be hex-encoded in the resulting
// JSON.
func Bytes(key string, value []byte) Field {
	return func(e *Event) {
		e.Bytes(key, value)
	}
}

// Bool adds the field key with i as a bool to the *Event context.
func Bool(key string, value bool) Field {
	return func(e *Event) {
		e.Bool(key, value)
	}
}

// Bools adds the field key with i as a []bool to the *Event context.
func Bools(key string, value []bool) Field {
	return func(e *Event) {
		e.Bools(key, value)
	}
}

// Interface adds the field key with i marshaled using reflection.
func Interface(key string, value interface{}) Field {
	return func(e *Event) {
		e.Interface(key, value)
	}
}

// IP adds IPv4 or IPv6 Address to the event
func IP(key string, value net.IP) Field {
	return func(e *Event) {
		e.IP(key, value)
	}
}

// IPNet adds IPv4 or IPv6 Prefix (address and mask) to the event
func IPNet(key string, value net.IPNet) Field {
	return func(e *Event) {
		e.IPNet(key, value)
	}
}

// HardwareAddr adds HardwareAddr to the event
func HardwareAddr(key string, value net.HardwareAddr) Field {
	return func(e *Event) {
		e.HardwareAddr(key, value)
	}
}

// Timestamp adds the current local time as UNIX timestamp to the *Event context with the
// logger.TimestampFieldName key.
func Timestamp(enable bool) Field {
	return func(e *Event) {
		e.Timestamp(enable)
	}
}

// Error adds the field key with serialized err to the *Event context.
// If err is nil, no field is added.
func Error(key string, value error) Field {
	return func(e *Event) {
		e.Error(key, value)
	}
}

// Err adds the field "error" with serialized err to the *Event context.
// If err is nil, no field is added.
// To customize the key name, uze rz.ErrorFieldName.
//
// If Stack() has been called before and rz.ErrorStackMarshaler is defined,
// the err is passed to ErrorStackMarshaler and the result is appended to the
// rz.ErrorStackFieldName.
func Err(value error) Field {
	return func(e *Event) {
		e.Err(value)
	}
}

// Errors adds the field key with errs as an array of serialized errors to the
// *Event context.
func Errors(key string, value []error) Field {
	return func(e *Event) {
		e.Errors(key, value)
	}
}

// Hex adds the field key with val as a hex string to the *Event context.
func Hex(key string, value []byte) Field {
	return func(e *Event) {
		e.Hex(key, value)
	}
}

// RawJSON adds already encoded JSON to the log line under key.
//
// No sanity check is performed on b; it must not contain carriage returns and
// be valid JSON.
func RawJSON(key string, value []byte) Field {
	return func(e *Event) {
		e.RawJSON(key, value)
	}
}

// Int adds the field key with i as a int to the *Event context.
func Int(key string, value int) Field {
	return func(e *Event) {
		e.Int(key, value)
	}
}

// Ints adds the field key with i as a []int to the *Event context.
func Ints(key string, value []int) Field {
	return func(e *Event) {
		e.Ints(key, value)
	}
}

// Int8 adds the field key with i as a int8 to the *Event context.
func Int8(key string, value int8) Field {
	return func(e *Event) {
		e.Int8(key, value)
	}
}

// Ints8 adds the field key with i as a []int8 to the *Event context.
func Ints8(key string, value []int8) Field {
	return func(e *Event) {
		e.Ints8(key, value)
	}
}

// Int16 adds the field key with i as a int16 to the *Event context.
func Int16(key string, value int16) Field {
	return func(e *Event) {
		e.Int16(key, value)
	}
}

// Ints16 adds the field key with i as a []int16 to the *Event context.
func Ints16(key string, value []int16) Field {
	return func(e *Event) {
		e.Ints16(key, value)
	}
}

// Int32 adds the field key with i as a int32 to the *Event context.
func Int32(key string, value int32) Field {
	return func(e *Event) {
		e.Int32(key, value)
	}
}

// Ints32 adds the field key with i as a []int32 to the *Event context.
func Ints32(key string, value []int32) Field {
	return func(e *Event) {
		e.Ints32(key, value)
	}
}

// Int64 adds the field key with i as a int64 to the *Event context.
func Int64(key string, value int64) Field {
	return func(e *Event) {
		e.Int64(key, value)
	}
}

// Ints64 adds the field key with i as a []int64 to the *Event context.
func Ints64(key string, value []int64) Field {
	return func(e *Event) {
		e.Ints64(key, value)
	}
}

// Uint adds the field key with i as a uint to the *Event context.
func Uint(key string, value uint) Field {
	return func(e *Event) {
		e.Uint(key, value)
	}
}

// Uints adds the field key with i as a []uint to the *Event context.
func Uints(key string, value []uint) Field {
	return func(e *Event) {
		e.Uints(key, value)
	}
}

// Uint8 adds the field key with i as a uint8 to the *Event context.
func Uint8(key string, value uint8) Field {
	return func(e *Event) {
		e.Uint8(key, value)
	}
}

// Uints8 adds the field key with i as a []uint8 to the *Event context.
func Uints8(key string, value []uint8) Field {
	return func(e *Event) {
		e.Uints8(key, value)
	}
}

// Uint16 adds the field key with i as a uint16 to the *Event context.
func Uint16(key string, value uint16) Field {
	return func(e *Event) {
		e.Uint16(key, value)
	}
}

// Uints16 adds the field key with i as a []uint16 to the *Event context.
func Uints16(key string, value []uint16) Field {
	return func(e *Event) {
		e.Uints16(key, value)
	}
}

// Uint32 adds the field key with i as a uint32 to the *Event context.
func Uint32(key string, value uint32) Field {
	return func(e *Event) {
		e.Uint32(key, value)
	}
}

// Uints32 adds the field key with i as a []uint32 to the *Event context.
func Uints32(key string, value []uint32) Field {
	return func(e *Event) {
		e.Uints32(key, value)
	}
}

// Uint64 adds the field key with i as a uint64 to the *Event context.
func Uint64(key string, value uint64) Field {
	return func(e *Event) {
		e.Uint64(key, value)
	}
}

// Uints64 adds the field key with i as a []uint64 to the *Event context.
func Uints64(key string, value []uint64) Field {
	return func(e *Event) {
		e.Uints64(key, value)
	}
}

// Float32 adds the field key with f as a float32 to the *Event context.
func Float32(key string, value float32) Field {
	return func(e *Event) {
		e.Float32(key, value)
	}
}

// Floats32 adds the field key with f as a []float32 to the *Event context.
func Floats32(key string, value []float32) Field {
	return func(e *Event) {
		e.Floats32(key, value)
	}
}

// Float64 adds the field key with f as a float64 to the *Event context.
func Float64(key string, value float64) Field {
	return func(e *Event) {
		e.Float64(key, value)
	}
}

// Floats64 adds the field key with f as a []float64 to the *Event context.
func Floats64(key string, value []float64) Field {
	return func(e *Event) {
		e.Floats64(key, value)
	}
}
