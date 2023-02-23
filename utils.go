package golog

import (
	"sync"
)

const (
	DefaultBufferSize       = 1 << 9
	DefaultBufferReturnSize = 1 << 12
)

// bufferPool is a global pool for logging buffers, common to *ALL* loggers.
var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, DefaultBufferSize) // By default you get 'DefaultBufferSize' byte sized buffers.
	},
}

func bufferPut(b []byte) {
	if b != nil {
		if len(b) < DefaultBufferReturnSize { // Only return up to this size
			b = b[:0]
			bufferPool.Put(b) //nolint:staticcheck // persistence is in the underlying array
		}
	}
}

// needsQuote returns true if string outside the range a-z, A-Z, 0-9
func needsQuote(s string) bool {
	for _, ch := range s {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '_' || ch == '-') {
			return true
		}
	}
	return false
}
