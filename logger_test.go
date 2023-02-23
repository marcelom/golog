package golog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	for v := 0; v < 100; v++ {
		t.Run(fmt.Sprintf("verbosity=%d", v), func(t *testing.T) {
			got := New(nil, v)

			assert.Equal(t, &Logger{buf: nil, logger: nil, verbosity: v}, got)
		})
	}
}

func TestLogger_With(t *testing.T) {
	tests := []struct {
		name string
		buf  []byte
		want Context
	}{
		{"nil buffer", nil, Context{logger: Logger{}}},
		{"existing buffer", []byte(" k=1"), Context{logger: Logger{buf: []byte(" k=1")}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{buf: tt.buf, logger: nil}
			got := l.With()

			assert.Equal(t, tt.want, *got)
		})
	}
}

func TestLogger_V(t *testing.T) {
	t.Run("lower verbosity", func(t *testing.T) {
		l := &Logger{buf: nil, logger: nil, verbosity: 10}

		got := l.V(5)

		if assert.NotNil(t, got) {
			assert.Same(t, l, got.logger) // event logger is identical
		}
	})
	t.Run("same verbosity", func(t *testing.T) {
		l := &Logger{buf: nil, logger: nil, verbosity: 5}

		got := l.V(5)

		if assert.NotNil(t, got) {
			assert.Same(t, l, got.logger) // event logger is identical
		}
	})
	t.Run("higher verbosity", func(t *testing.T) {
		l := &Logger{buf: nil, logger: nil, verbosity: 5}

		got := l.V(10)

		assert.Nil(t, got)
	})
}
