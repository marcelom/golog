package golog

import (
	"io"
	"log"
	"testing"
)

func Benchmark_Msg(b *testing.B) {
	logLogger := log.New(io.Discard, "", log.LstdFlags)
	gologLogger := New(log.New(io.Discard, "", log.LstdFlags), 0)

	b.Run("log", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			logLogger.Printf("msg=%s", "msg")
		}
	})

	b.Run("golog", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gologLogger.V(0).Msg("msg")
		}
	})
}

func Benchmark_Int(b *testing.B) {
	logLogger := log.New(io.Discard, "", log.LstdFlags)
	gologLogger := New(log.New(io.Discard, "", log.LstdFlags), 0)

	b.Run("log", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			logLogger.Printf("i=%d msg=%s", i, "msg")
		}
	})

	b.Run("golog", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gologLogger.V(0).Int("i", i).Msg("msg")
		}
	})
}
