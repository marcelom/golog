package golog

import (
	"fmt"
	"os"
	"strconv"
)

const (
	callerLevels = 2
	messageField = ` msg=`
	errorField   = ` error=`
)

// Event represents a log event. It is instanced by one of the level methods of
// Logger and finalized by the Msg or Msgf method.
type Event struct {
	buf    []byte
	logger *Logger
	err    error
}

func (e *Event) Enabled() bool {
	return e != nil
}

// Str appends a string field to the Event buffer.
func (e *Event) Str(k, v string) *Event {
	if e == nil {
		return nil
	}
	e.buf = appendStr(e.buf, k, v)
	return e
}

// Int appends an int field to the Event buffer.
func (e *Event) Int(k string, v int) *Event {
	if e == nil {
		return nil
	}
	e.buf = appendInt(e.buf, k, v)
	return e
}

// Float appends a float field to the Event buffer.
func (e *Event) Float(k string, v float64) *Event {
	if e == nil {
		return nil
	}
	e.buf = appendFloat(e.buf, k, v)
	return e
}

// Bool appends a bool field to the Event buffer.
func (e *Event) Bool(k string, v bool) *Event {
	if e == nil {
		return nil
	}
	e.buf = appendBool(e.buf, k, v)
	return e
}

// WithError attaches an error to the Event.
func (e *Event) Err(err error) *Event {
	if e == nil {
		return nil
	}
	e.err = err
	return e
}

// Msg finalizes the event and sends it to the handler.
func (e *Event) Msg(message string) {
	if e == nil {
		return
	}

	b := bufferPool.Get().([]byte)
	defer bufferPut(b)

	if len(message) > 0 {
		b = append(b, messageField...)
		b = strconv.AppendQuote(b, message)
	}
	if e.err != nil {
		b = append(b, errorField...)
		b = strconv.AppendQuote(b, e.err.Error())
	}
	if e.logger.buf != nil {
		b = append(b, e.logger.buf...)
	}
	if e.buf != nil {
		b = append(b, e.buf...)
	}

	if err := e.logger.logger.Output(callerLevels, string(b)); err != nil {
		fmt.Fprintf(os.Stderr, "golog error: unable to log event, err=%v", err) // this should never happen
	}

	bufferPut(e.buf)
}

// Msgf is like Msg, but takes a fmt formatted string.
func (e *Event) Msgf(format string, v ...interface{}) {
	if e == nil {
		return
	}
	e.Msg(fmt.Sprintf(format, v...))
}
