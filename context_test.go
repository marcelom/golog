package golog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContext_Level(t *testing.T) {
	wantLogger := &Logger{buf: nil, logger: nil, verbosity: 100}
	c := &Context{logger: Logger{buf: nil, logger: nil, verbosity: 0}}

	got := c.Verbosity(100).Logger()
	assert.Equal(t, wantLogger, got)
}

func TestContext_Str(t *testing.T) {
	wantLogger := &Logger{buf: []byte(` k="abc123"`), logger: nil, verbosity: 0}
	c := &Context{logger: Logger{buf: nil, logger: nil, verbosity: 0}}

	got := c.Str("k", "abc123").Logger()
	assert.Equal(t, wantLogger, got)
}

func TestContext_Int(t *testing.T) {
	wantLogger := &Logger{buf: []byte(` k=123`), logger: nil, verbosity: 0}
	c := &Context{logger: Logger{buf: nil, logger: nil, verbosity: 0}}

	got := c.Int("k", 123).Logger()
	assert.Equal(t, wantLogger, got)
}

func TestContext_Float(t *testing.T) {
	wantLogger := &Logger{buf: []byte(` k=1.010203`), logger: nil, verbosity: 0}
	c := &Context{logger: Logger{buf: nil, logger: nil, verbosity: 0}}

	got := c.Float("k", 1.010203).Logger()
	assert.Equal(t, wantLogger, got)
}

func TestContext_Bool(t *testing.T) {
	wantLogger := &Logger{buf: []byte(` k=true`), logger: nil, verbosity: 0}
	c := &Context{logger: Logger{buf: nil, logger: nil, verbosity: 0}}

	got := c.Bool("k", true).Logger()
	assert.Equal(t, wantLogger, got)
}

func TestContext_Logger(t *testing.T) {
	wantLogger := &Logger{buf: []byte(`a buffer...`), logger: nil, verbosity: 0}
	c := &Context{logger: Logger{buf: []byte(`a buffer...`), logger: nil, verbosity: 0}}

	got := c.Logger()
	assert.Equal(t, wantLogger, got)
}
