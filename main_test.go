package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying Hello to someone", func(t *testing.T) {
		got := Hello("Pranav")
		want := "Hello Pranav!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello World! when no name is given", func(t *testing.T) {
		got := Hello("")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
