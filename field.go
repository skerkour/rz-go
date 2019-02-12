package rz

import (
	"net"
	"time"
)

type Field func(e *Event)

func Discard() Field {
	return func(e *Event) {
		e.Discard()
	}
}

type discardField struct{}

func (f *discardField) apply(e *Event) {
	e.level = Disabled
}

// Stack enables stack trace printing for the error passed to Err().
//
// logger.errorStackMarshaler must be set for this method to do something.
func Stack(enable bool) Field {
	return func(e *Event) {
		e.Stack(enable)
	}
}

type stackField struct{}

func (f *stackField) apply(e *Event) {
	e.stack = true
}

// Caller adds the file:line of the caller with the rz.CallerFieldName key.
func Caller() Field {
	return func(e *Event) {
		e.Caller()
	}
}

type callerField struct{}

func (f *callerField) apply(e *Event) {
	e.caller = true
}

func Fields(fields map[string]interface{}) Field {
	return func(e *Event) {
		e.Fields(fields)
	}
}

type filedsField struct {
	fields map[string]interface{}
}

func (f *filedsField) apply(e *Event) {
	e.buf = e.appendFields(e.buf, f.fields)
}

func String(key, value string) Field {
	return func(e *Event) {
		e.String(key, value)
	}
}

// type stringField struct {
// 	key   string
// 	value string
// }

// func (f *stringField) apply(e *Event) {
// 	e.encoder.AppendString(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

func Strings(key string, value []string) Field {
	return func(e *Event) {
		e.Strings(key, value)
	}
}

// type stringsField struct {
// 	key   string
// 	value []string
// }

// func (f *stringsField) apply(e *Event) {
// 	e.encoder.AppendStrings(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Time adds the field key with t formated as string using rz.TimeFieldFormat.
func Time(key string, value time.Time) Field {
	return func(e *Event) {
		e.Time(key, value)
	}
}

// type timeField struct {
// 	key   string
// 	value time.Time
// }

// func (f *timeField) apply(e *Event) {
// 	e.encoder.AppendTime(e.encoder.AppendKey(e.buf, f.key), f.value, e.timeFieldFormat)
// }

// Times adds the field key with t formated as string using rz.TimeFieldFormat.
func Times(key string, value []time.Time) Field {
	return func(e *Event) {
		e.Times(key, value)
	}
}

// type timesField struct {
// 	key   string
// 	value []time.Time
// }

// func (f *timesField) apply(e *Event) {
// 	e.encoder.AppendTimes(e.encoder.AppendKey(e.buf, f.key), f.value, e.timeFieldFormat)
// }

// Duration adds the field key with duration d stored as rz.DurationFieldUnit.
// If rz.DurationFieldInteger is true, durations are rendered as integer
// instead of float.
func Duration(key string, value time.Duration) Field {
	return func(e *Event) {
		e.Duration(key, value)
	}
}

// type durationField struct {
// 	key   string
// 	value time.Duration
// }

// func (f *durationField) apply(e *Event) {
// 	e.encoder.AppendDuration(e.encoder.AppendKey(e.buf, f.key), f.value, DurationFieldUnit, DurationFieldInteger)
// }

// Durations adds the field key with duration d stored as rz.DurationFieldUnit.
// If rz.DurationFieldInteger is true, durations are rendered as integer
// instead of float.
func Durations(key string, value []time.Duration) Field {
	return func(e *Event) {
		e.Durations(key, value)
	}
}

// type durationsField struct {
// 	key   string
// 	value []time.Duration
// }

// func (f *durationsField) apply(e *Event) {
// 	e.encoder.AppendDurations(e.encoder.AppendKey(e.buf, f.key), f.value, DurationFieldUnit, DurationFieldInteger)
// }

// Object marshals an object that implement the LogObjectMarshaler interface.
func Object(key string, value LogObjectMarshaler) Field {
	return func(e *Event) {
		e.Object(key, value)
	}
}

// type objectField struct {
// 	key   string
// 	value LogObjectMarshaler
// }

// func (f *objectField) apply(e *Event) {
// 	e.buf = e.encoder.AppendKey(e.buf, f.key)
// 	e.appendObject(f.value)
// }

// EmbedObject marshals an object that implement the LogObjectMarshaler interface.
func EmbedObject(obj LogObjectMarshaler) Field {
	return func(e *Event) {
		e.EmbedObject(obj)
	}
}

type embedObjectField struct {
	value LogObjectMarshaler
}

func (f *embedObjectField) apply(e *Event) {
	f.value.MarshalRzObject(e)
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

type bytesField struct {
	key   string
	value []byte
}

func (f *bytesField) apply(e *Event) {
	e.buf = enc.AppendBytes(enc.AppendKey(e.buf, f.key), f.value)
}

// Bool adds the field key with i as a bool to the *Event context.
func Bool(key string, value bool) Field {
	return func(e *Event) {
		e.Bool(key, value)
	}
}

// type boolField struct {
// 	key   string
// 	value bool
// }

// func (f *boolField) apply(e *Event) {
// 	e.encoder.AppendBool(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Bools adds the field key with i as a []bool to the *Event context.
func Bools(key string, value []bool) Field {
	return func(e *Event) {
		e.Bools(key, value)
	}
}

// type boolsField struct {
// 	key   string
// 	value []bool
// }

// func (f *boolsField) apply(e *Event) {
// 	e.encoder.AppendBools(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Interface adds the field key with i marshaled using reflection.
func Interface(key string, value interface{}) Field {
	return func(e *Event) {
		e.Interface(key, value)
	}
}

// type interfaceField struct {
// 	key   string
// 	value interface{}
// }

// func (f *interfaceField) apply(e *Event) {
// 	if obj, ok := f.value.(LogObjectMarshaler); ok {
// 		e.object(f.key, obj)
// 	}
// 	e.buf = e.encoder.AppendInterface(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// IP adds IPv4 or IPv6 Address to the event
func IP(key string, value net.IP) Field {
	return func(e *Event) {
		e.IP(key, value)
	}
}

// type ipField struct {
// 	key   string
// 	value net.IP
// }

// func (f *ipField) apply(e *Event) {
// 	e.buf = e.encoder.AppendIPAddr(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// IPNet adds IPv4 or IPv6 Prefix (address and mask) to the event
func IPNet(key string, value net.IPNet) Field {
	return func(e *Event) {
		e.IPNet(key, value)
	}
}

// type ipNetField struct {
// 	key   string
// 	value net.IPNet
// }

// func (f *ipNetField) apply(e *Event) {
// 	e.buf = e.encoder.AppendIPPrefix(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// HardwareAddr adds HardwareAddr to the event
func HardwareAddr(key string, value net.HardwareAddr) Field {
	return func(e *Event) {
		e.MACAddr(key, value)
	}
}

// type harwareAddrField struct {
// 	key   string
// 	value net.HardwareAddr
// }

// func (f *harwareAddrField) apply(e *Event) {
// 	e.buf = e.encoder.AppendMACAddr(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Timestamp adds the current local time as UNIX timestamp to the *Event context with the
// logger.TimestampFieldName key.
func Timestamp(enable bool) Field {
	return func(e *Event) {
		e.Timestamp(enable)
	}
}

// type timestampField struct {
// 	enable bool
// }

// func (f *timestampField) apply(e *Event) {
// 	if e.configEvent && f.enable {
// 		e.timestamp = true
// 	} else if f.enable {
// 		e.timestamp = false
// 		e.buf = e.encoder.AppendTime(e.encoder.AppendKey(e.buf, e.timestampFieldName), e.timestampFunc(), e.timeFieldFormat)
// 	} else {
// 		e.timestamp = false
// 	}
// }

// Error adds the field key with serialized err to the *Event context.
// If err is nil, no field is added.
func Error(key string, value error) Field {
	return func(e *Event) {
		e.Error(key, value)
	}
}

// type errorField struct {
// 	key   string
// 	value error
// }

// func (f *errorField) apply(e *Event) {
// 	switch m := ErrorMarshalFunc(f.value).(type) {
// 	case nil:
// 	case LogObjectMarshaler:
// 		e.object(f.key, m)
// 	case error:
// 		e.string(f.key, m.Error())
// 	case string:
// 		e.string(f.key, m)
// 	default:
// 		e.iinterface(f.key, m)
// 	}
// }

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

func Errors(key string, value []error) Field {
	return func(e *Event) {
		e.Errors(key, value)
	}
}

// type errField struct {
// 	value error
// }

// func (f *errField) apply(e *Event) {
// 	if e.stack && ErrorStackMarshaler != nil {
// 		switch m := ErrorStackMarshaler(f.value).(type) {
// 		case nil:
// 		case LogObjectMarshaler:
// 			e.object(e.errorStackFieldName, m)
// 		case error:
// 			e.string(e.errorStackFieldName, m.Error())
// 		case string:
// 			e.string(e.errorStackFieldName, m)
// 		default:
// 			e.iinterface(e.errorStackFieldName, m)
// 		}
// 	}
// 	e.error(e.errorFieldName, f.value)
// }

// Hex adds the field key with val as a hex string to the *Event context.
func Hex(key string, value []byte) Field {
	return func(e *Event) {
		e.Hex(key, value)
	}
}

// type hexField struct {
// 	key   string
// 	value []byte
// }

// func (f *hexField) apply(e *Event) {
// 	e.buf = e.encoder.AppendHex(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// RawJSON adds already encoded JSON to the log line under key.
//
// No sanity check is performed on b; it must not contain carriage returns and
// be valid JSON.
func RawJSON(key string, value []byte) Field {
	return func(e *Event) {
		e.RawJSON(key, value)
	}
}

// type rawJSONField struct {
// 	key   string
// 	value []byte
// }

// func (f *rawJSONField) apply(e *Event) {
// 	e.buf = e.encoder.AppendHex(e.encoder.AppendKey(e.buf, f.key), f.value)
// 	e.buf = appendJSON(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Dict adds the field key with a dict to the event context.
// Use rz.Dict() to create the dictionary.
// func Dict(key string, fields ...Field) Field {
// 	return &dictField{key, fields}
// 	// dict.buf = enc.AppendEndMarker(dict.buf)
// 	// e.buf = append(enc.AppendKey(e.buf, key), dict.buf...)
// 	// putEvent(dict)
// 	// return e
// }

// type dictField struct {
// 	key   string
// 	value []Field
// }

// func (f *dictField) apply(e *Event) {
// 	dict := newEvent(nil, 0)
// 	for i := range f.value {
// 		f.value[i].apply(dict)
// 	}
// 	dict.buf = e.encoder.AppendEndMarker(dict.buf)
// 	e.buf = append(e.encoder.AppendKey(e.buf, f.key), dict.buf...)
// 	putEvent(dict)
// }

// Int adds the field key with i as a int to the *Event context.
func Int(key string, value int) Field {
	return func(e *Event) {
		e.Int(key, value)
	}
}

// type intField struct {
// 	key   string
// 	value int
// }

// func (f *intField) apply(e *Event) {
// 	e.encoder.AppendInt(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Ints adds the field key with i as a []int to the *Event context.
func Ints(key string, value []int) Field {
	return func(e *Event) {
		e.Ints(key, value)
	}
}

// type intsField struct {
// 	key   string
// 	value []int
// }

// func (f *intsField) apply(e *Event) {
// 	e.encoder.AppendInts(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Int8 adds the field key with i as a int8 to the *Event context.
func Int8(key string, value int8) Field {
	return func(e *Event) {
		e.Int8(key, value)
	}
}

// type int8Field struct {
// 	key   string
// 	value int8
// }

// func (f *int8Field) apply(e *Event) {
// 	e.encoder.AppendInt8(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Ints8 adds the field key with i as a []int8 to the *Event context.
func Ints8(key string, value []int8) Field {
	return func(e *Event) {
		e.Ints8(key, value)
	}
}

// type ints8Field struct {
// 	key   string
// 	value []int8
// }

// func (f *ints8Field) apply(e *Event) {
// 	e.encoder.AppendInts8(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Int16 adds the field key with i as a int16 to the *Event context.
func Int16(key string, value int16) Field {
	return func(e *Event) {
		e.Int16(key, value)
	}
}

// type int16Field struct {
// 	key   string
// 	value int16
// }

// func (f *int16Field) apply(e *Event) {
// 	e.encoder.AppendInt16(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Ints16 adds the field key with i as a []int16 to the *Event context.
func Ints16(key string, value []int16) Field {
	return func(e *Event) {
		e.Ints16(key, value)
	}
}

// type ints16Field struct {
// 	key   string
// 	value []int16
// }

// func (f *ints16Field) apply(e *Event) {
// 	e.encoder.AppendInts16(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Int32 adds the field key with i as a int32 to the *Event context.
func Int32(key string, value int32) Field {
	return func(e *Event) {
		e.Int32(key, value)
	}
}

// type int32Field struct {
// 	key   string
// 	value int32
// }

// func (f *int32Field) apply(e *Event) {
// 	e.encoder.AppendInt32(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Ints32 adds the field key with i as a []int32 to the *Event context.
func Ints32(key string, value []int32) Field {
	return func(e *Event) {
		e.Ints32(key, value)
	}
}

// type ints32Field struct {
// 	key   string
// 	value []int32
// }

// func (f *ints32Field) apply(e *Event) {
// 	e.encoder.AppendInts32(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Int64 adds the field key with i as a int64 to the *Event context.
func Int64(key string, value int64) Field {
	return func(e *Event) {
		e.Int64(key, value)
	}
}

// type int64Field struct {
// 	key   string
// 	value int64
// }

// func (f *int64Field) apply(e *Event) {
// 	e.encoder.AppendInt64(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Ints64 adds the field key with i as a []int64 to the *Event context.
func Ints64(key string, value []int64) Field {
	return func(e *Event) {
		e.Ints64(key, value)
	}
}

// type ints64Field struct {
// 	key   string
// 	value []int64
// }

// func (f *ints64Field) apply(e *Event) {
// 	e.encoder.AppendInts64(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Uint adds the field key with i as a uint to the *Event context.
func Uint(key string, value uint) Field {
	return func(e *Event) {
		e.Uint(key, value)
	}
}

// type uintField struct {
// 	key   string
// 	value uint
// }

// func (f *uintField) apply(e *Event) {
// 	e.encoder.AppendUint(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Uints adds the field key with i as a []uint to the *Event context.
func Uints(key string, value []uint) Field {
	return func(e *Event) {
		e.Uints(key, value)
	}
}

// type uintsField struct {
// 	key   string
// 	value []uint
// }

// func (f *uintsField) apply(e *Event) {
// 	e.encoder.AppendUints(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Uint8 adds the field key with i as a uint8 to the *Event context.
func Uint8(key string, value uint8) Field {
	return func(e *Event) {
		e.Uint8(key, value)
	}
}

// type uint8Field struct {
// 	key   string
// 	value uint8
// }

// func (f *uint8Field) apply(e *Event) {
// 	e.encoder.AppendUint8(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Uints8 adds the field key with i as a []uint8 to the *Event context.
func Uints8(key string, value []uint8) Field {
	return func(e *Event) {
		e.Uints8(key, value)
	}
}

// type uints8Field struct {
// 	key   string
// 	value []uint8
// }

// func (f *uints8Field) apply(e *Event) {
// 	e.encoder.AppendUints8(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Uint16 adds the field key with i as a uint16 to the *Event context.
func Uint16(key string, value uint16) Field {
	return func(e *Event) {
		e.Uint16(key, value)
	}
}

// type uint16Field struct {
// 	key   string
// 	value uint16
// }

// func (f *uint16Field) apply(e *Event) {
// 	e.encoder.AppendUint16(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Uints16 adds the field key with i as a []uint16 to the *Event context.
func Uints16(key string, value []uint16) Field {
	return func(e *Event) {
		e.Uints16(key, value)
	}
}

// type uints16Field struct {
// 	key   string
// 	value []uint16
// }

// func (f *uints16Field) apply(e *Event) {
// 	e.encoder.AppendUints16(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Uint32 adds the field key with i as a uint32 to the *Event context.
func Uint32(key string, value uint32) Field {
	return func(e *Event) {
		e.Uint32(key, value)
	}
}

// type uint32Field struct {
// 	key   string
// 	value uint32
// }

// func (f *uint32Field) apply(e *Event) {
// 	e.encoder.AppendUint32(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Uints32 adds the field key with i as a []uint32 to the *Event context.
func Uints32(key string, value []uint32) Field {
	return func(e *Event) {
		e.Uints32(key, value)
	}
}

// type uints32Field struct {
// 	key   string
// 	value []uint32
// }

// func (f *uints32Field) apply(e *Event) {
// 	e.encoder.AppendUints32(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Uint64 adds the field key with i as a uint64 to the *Event context.
func Uint64(key string, value uint64) Field {
	return func(e *Event) {
		e.Uint64(key, value)
	}
}

// type uint64Field struct {
// 	key   string
// 	value uint64
// }

// func (f *uint64Field) apply(e *Event) {
// 	e.encoder.AppendUint64(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Uints64 adds the field key with i as a []uint64 to the *Event context.
func Uints64(key string, value []uint64) Field {
	return func(e *Event) {
		e.Uints64(key, value)
	}
}

// type uints64Field struct {
// 	key   string
// 	value []uint64
// }

// func (f *uints64Field) apply(e *Event) {
// 	e.encoder.AppendUints64(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Float32 adds the field key with f as a float32 to the *Event context.
func Float32(key string, value float32) Field {
	return func(e *Event) {
		e.Float32(key, value)
	}
}

// type float32Field struct {
// 	key   string
// 	value float32
// }

// func (f *float32Field) apply(e *Event) {
// 	e.buf = e.encoder.AppendFloat32(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Floats32 adds the field key with f as a []float32 to the *Event context.
func Floats32(key string, value []float32) Field {
	return func(e *Event) {
		e.Floats32(key, value)
	}
}

// type floats32Field struct {
// 	key   string
// 	value []float32
// }

// func (f *floats32Field) apply(e *Event) {
// 	e.buf = e.encoder.AppendFloats32(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Float64 adds the field key with f as a float64 to the *Event context.
func Float64(key string, value float64) Field {
	return func(e *Event) {
		e.Float64(key, value)
	}
}

// type float64Field struct {
// 	key   string
// 	value float64
// }

// func (f *float64Field) apply(e *Event) {
// 	e.buf = e.encoder.AppendFloat64(e.encoder.AppendKey(e.buf, f.key), f.value)
// }

// Floats64 adds the field key with f as a []float64 to the *Event context.
func Floats64(key string, value []float64) Field {
	return func(e *Event) {
		e.Floats64(key, value)
	}
}

// type floats64Field struct {
// 	key   string
// 	value []float64
// }

// func (f *floats64Field) apply(e *Event) {
// 	e.buf = e.encoder.AppendFloats64(e.encoder.AppendKey(e.buf, f.key), f.value)
// }
