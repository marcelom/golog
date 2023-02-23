package golog

import (
	"strconv"
)

// appendKey takes a buffer and appends a key in the form:
// ' key=' or ' "key"='
func appendKey(b []byte, k string) []byte {
	b = append(b, ' ')
	if needsQuote(k) {
		b = strconv.AppendQuote(b, k)
	} else {
		b = append(b, k...)
	}
	b = append(b, '=')
	return b
}

// appendStr takes a buffer and appends a key then a value as a string.
// The value is always quoted!
func appendStr(b []byte, k, v string) []byte {
	b = appendKey(b, k)
	b = strconv.AppendQuote(b, v)
	return b
}

// appendInt takes a buffer and appends a key then a value as Int.
func appendInt(b []byte, k string, v int) []byte {
	b = appendKey(b, k)
	b = strconv.AppendInt(b, int64(v), 10)
	return b
}

// appendFloat takes a buffer and appends a key then a value as Float.
func appendFloat(b []byte, k string, v float64) []byte {
	b = appendKey(b, k)
	b = strconv.AppendFloat(b, v, 'f', -1, 64)
	return b
}

// appendFloat takes a buffer and appends a key then a value as Float.
func appendBool(b []byte, k string, v bool) []byte {
	b = appendKey(b, k)
	b = strconv.AppendBool(b, v)
	return b
}
