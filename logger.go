package golog

import "log"

// Logger represents a logger with configurable Level and a stdlib log.Logger
type Logger struct {
	buf       []byte
	logger    *log.Logger
	verbosity int
}

// New returns a new Logger with the provided stdlib log.Logger
func New(logger *log.Logger, verbosity int) *Logger {
	return &Logger{
		buf:       nil,
		logger:    logger,
		verbosity: verbosity,
	}
}

// With return a new context, useful for configuring subloggers
func (l *Logger) With() *Context {
	l2 := Logger{logger: l.logger, verbosity: l.verbosity}
	if l.buf != nil {
		l2.buf = append(bufferPool.New().([]byte), l.buf...)
	}
	return &Context{logger: l2}
}

func (l *Logger) V(verbosity int) *Event {
	if verbosity > l.verbosity {
		return nil
	}

	return &Event{
		buf:    bufferPool.New().([]byte),
		logger: l,
		err:    nil,
	}
}
