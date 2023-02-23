package golog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_appendStr(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value string
		want  string
	}{
		{"simple", "key", "value", ` key="value"`},
		{"quotable value", "key", "quotable value", ` key="quotable value"`},
		{"quotable key", "key needs quotes", "value", ` "key needs quotes"="value"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := make([]byte, 0, 100)
			got := appendStr(b, tt.key, tt.value)
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func Benchmark_appendStr(b *testing.B) {
	buf := make([]byte, 0, 100)
	for i := 0; i < b.N; i++ {
		_ = appendStr(buf, "key", "value")
	}
}

func Test_appendInt(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value int
		want  string
	}{
		{"simple", "key", 1, ` key=1`},
		{"quotable key", "key needs quotes", 1, ` "key needs quotes"=1`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := make([]byte, 0, 100)
			got := appendInt(b, tt.key, tt.value)
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func Benchmark_appendInt(b *testing.B) {
	buf := make([]byte, 0, 100)
	for i := 0; i < b.N; i++ {
		_ = appendInt(buf, "key", 1)
	}
}

func Test_appendFloat(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value float64
		want  string
	}{
		{"simple", "key", 1.00001, ` key=1.00001`},
		{"quotable key", "key needs quotes", 2765.8483304, ` "key needs quotes"=2765.8483304`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := make([]byte, 0, 100)
			got := appendFloat(b, tt.key, tt.value)
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func Benchmark_appendFloat(b *testing.B) {
	buf := make([]byte, 0, 100)
	for i := 0; i < b.N; i++ {
		_ = appendFloat(buf, "key", 1.0000001)
	}
}

func Test_appendBool(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value bool
		want  string
	}{
		{"true", "key", true, ` key=true`},
		{"true + quotable key", "key needs quotes", true, ` "key needs quotes"=true`},
		{"false", "key", false, ` key=false`},
		{"false + quotable key", "key needs quotes", false, ` "key needs quotes"=false`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := make([]byte, 0, 100)
			got := appendBool(b, tt.key, tt.value)
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func Benchmark_appendBool(b *testing.B) {
	buf := make([]byte, 0, 100)
	for i := 0; i < b.N; i++ {
		_ = appendBool(buf, "key", true)
	}
}
