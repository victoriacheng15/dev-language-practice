package main

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeats a character 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("repeats a character 3 times", func(t *testing.T) {
		got := Repeat("b", 3)
		want := "bbb"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
