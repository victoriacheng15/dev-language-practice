package main

import "testing"

func TestAdder(t *testing.T) {
	t.Run("adding 2 and 3", func(t *testing.T) {
		x, y := 2, 3
		got := Add(x, y)
		want := 5
		assertCorrectAddition(t, x, y, got, want)
	})

	t.Run("adding -2 and -3", func(t *testing.T) {
		x, y := -2, -3
		got := Add(x, y)
		want := -5
		assertCorrectAddition(t, x, y, got, want)
	})

	t.Run("adding -2 and 3", func(t *testing.T) {
		x, y := -2, 3
		got := Add(x, y)
		want := 1
		assertCorrectAddition(t, x, y, got, want)
	})
}

func assertCorrectAddition(t testing.TB, x, y, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Add(%d, %d): got %d, want %d", x, y, got, want)
	}
}
