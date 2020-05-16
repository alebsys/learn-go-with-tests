package main

import "testing"

var assertCorrect = func(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetMinor(t *testing.T) {
	t.Run("C to Am", func(t *testing.T) {
		want := "Am"
		got, _ := getMinor("C")
		assertCorrect(t, got, want)
	})
	t.Run("Unknown major error", func(t *testing.T) {
		want := "Unknown chord"
		_, err := getMinor("Z")
		if err != nil {
			assertCorrect(t, err.Error(), want)
		}
	})
	t.Run("Blank major error", func(t *testing.T) {
		want := "You didn't said chord"
		_, err := getMinor("")
		if err != nil {
			assertCorrect(t, err.Error(), want)
		}
	})
}

func TestGetMajor(t *testing.T) {
	t.Run("Am to C", func(t *testing.T) {
		want := "C"
		got, _ := getMajor("Am")
		assertCorrect(t, got, want)
	})
	t.Run("Unknown minor error", func(t *testing.T) {
		want := "Unknown chord"
		_, err := getMajor("Zm")
		if err != nil {
			assertCorrect(t, err.Error(), want)
		}
	})
	t.Run("Blank minor error", func(t *testing.T) {
		want := "You didn't said chord"
		_, err := getMajor("")
		if err != nil {
			assertCorrect(t, err.Error(), want)
		}
	})
}
