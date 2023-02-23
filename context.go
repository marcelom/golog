package golog

// Context is a wrapper to build sub-loggers.
type Context struct {
	logger Logger
}

// Verbosity sets the verbosity of the attached logger.
func (c *Context) Verbosity(verbosity int) *Context {
	c.logger.verbosity = verbosity
	return c
}

// Str sets the context with a string field.
func (c *Context) Str(k, v string) *Context {
	if c.logger.buf == nil {
		c.logger.buf = bufferPool.Get().([]byte)
	}
	c.logger.buf = appendStr(c.logger.buf, k, v)

	return c
}

// Int sets the context with an int field.
func (c *Context) Int(k string, v int) *Context {
	if c.logger.buf == nil {
		c.logger.buf = bufferPool.Get().([]byte)
	}
	c.logger.buf = appendInt(c.logger.buf, k, v)

	return c
}

// Float sets the context with a float field.
func (c *Context) Float(k string, v float64) *Context {
	if c.logger.buf == nil {
		c.logger.buf = bufferPool.Get().([]byte)
	}
	c.logger.buf = appendFloat(c.logger.buf, k, v)

	return c
}

// Bool sets the context with a bool field.
func (c *Context) Bool(k string, v bool) *Context {
	if c.logger.buf == nil {
		c.logger.buf = bufferPool.Get().([]byte)
	}
	c.logger.buf = appendBool(c.logger.buf, k, v)

	return c
}

// Logger returns the logger associated with this Context.
func (c *Context) Logger() *Logger {
	return &c.logger
}
