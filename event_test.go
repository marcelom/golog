package golog

import (
	"bytes"
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	var nilEvent *Event = nil
	tests := []struct {
		name string
		got  *Event
		want *Event
	}{
		{"Str nil event", nilEvent.Str("k", "abc123"), nil},
		{"Str nil buffer", (&Event{buf: nil}).Str("k", "def456"), &Event{buf: []byte(` k="def456"`)}},
		{"Str pre-existing buffer", (&Event{buf: []byte(` a=1`)}).Str("kk", "abc 123"),
			&Event{buf: []byte(` a=1 kk="abc 123"`)}},
		{"Int nil event", nilEvent.Int("zz", 11), nil},
		{"Int nil buffer", (&Event{buf: nil}).Int("k", 12345), &Event{buf: []byte(` k=12345`)}},
		{"Int pre-existing buffer", (&Event{buf: []byte(` a=1`)}).Int("kk", 54321), &Event{buf: []byte(` a=1 kk=54321`)}},
		{"Float nil event", nilEvent.Float("zz", 1.01), nil},
		{"Float nil buffer", (&Event{buf: nil}).Float("k", 1.01002), &Event{buf: []byte(` k=1.01002`)}},
		{"Float pre-existing buffer", (&Event{buf: []byte(` a=1`)}).Float("kk", 1.010020003),
			&Event{buf: []byte(` a=1 kk=1.010020003`)}},
		{"Bool nil event", nilEvent.Bool("zz", true), nil},
		{"Bool nil buffer", (&Event{buf: nil}).Bool("k", false), &Event{buf: []byte(` k=false`)}},
		{"Bool pre-existing buffer", (&Event{buf: []byte(` a=1`)}).Bool("kk", true), &Event{buf: []byte(` a=1 kk=true`)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.got)
		})
	}
}

func TestEvent_WithError(t *testing.T) {
	err := errors.New("oh noes")
	tests := []struct {
		name  string
		event *Event
		want  *Event
	}{
		{"nil event", nil, nil},
		{"nil error", &Event{}, &Event{err: err}},
		{"pre-existing error", &Event{err: errors.New("another error")}, &Event{err: err}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.want.Err(err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEvent_Msg(t *testing.T) {
	buf := &bytes.Buffer{}
	l := &Logger{logger: log.New(buf, "TEST TEST TEST --", log.Lmsgprefix), verbosity: 0}

	l.V(0).Msg("a message")

	assert.Equal(t, `TEST TEST TEST -- msg="a message"`+"\n", buf.String())
}

func TestEvent_Msg_Nil(t *testing.T) {
	buf := &bytes.Buffer{}
	l := &Logger{logger: log.New(buf, "TEST TEST TEST --", log.Lmsgprefix), verbosity: 0}

	l.V(10).Msg("a message")

	assert.Equal(t, "", buf.String())
}

func TestEvent_Msgf(t *testing.T) {
	buf := &bytes.Buffer{}
	l := &Logger{logger: log.New(buf, "TEST TEST TEST --", log.Lmsgprefix), verbosity: 0}

	message := "message"
	l.V(0).Msgf("a %s", message)

	assert.Equal(t, `TEST TEST TEST -- msg="a message"`+"\n", buf.String())
}

func TestEvent_Msgf_Nil(t *testing.T) {
	buf := &bytes.Buffer{}
	l := &Logger{logger: log.New(buf, "TEST TEST TEST --", log.Lmsgprefix), verbosity: 0}

	message := "message"
	l.V(10).Msgf("a %s", message)

	assert.Equal(t, "", buf.String())
}

func TestEvent_Enabled(t *testing.T) {
	var e *Event

	e = &Event{}
	assert.True(t, e.Enabled())

	e = nil
	assert.False(t, e.Enabled())
}
