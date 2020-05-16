package main

import "testing"

// TestChordCircle function for test my app
func TestChordCircle(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("circle of fifths", func(t *testing.T) {
		got := ChordCircle(c)
		want := "C chord have Am parallel chord"
		assertCorrectMessage(t, got, want)
	})
}
